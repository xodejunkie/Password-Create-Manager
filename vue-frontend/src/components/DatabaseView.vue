<template>
  <div>
    <h2 class="header">Saved Passwords</h2>
    <div class="sidebar">
      <ul>
        <li @click="activeTab = 'generated'; fetchGeneratedPasswords()" :class="{ 'active': activeTab === 'generated' }">Generated Passwords</li>
        <li @click="activeTab = 'created'; fetchCreatedPasswords()" :class="{ 'active': activeTab === 'created' }">Created Passwords</li>
      </ul>
    </div>
    <div class="options-bar">
      <div class="options">
        <div class="filter-icon" @click="toggleFilter" :class="{ 'active': isFilterActive }">
          <i class="fas fa-filter"></i>
          <div class="filter-dropdown" v-show="showFilter">
            <!-- Filter options -->
            <label for="dateRange">Date Range:</label>
            <input type="date" v-model="dateRangeFrom" />
            <input type="date" v-model="dateRangeTo" />
            <br />
            <label for="timeRange">Time Range:</label>
            <input type="time" v-model="timeRangeFrom" />
            <input type="time" v-model="timeRangeTo" />
            <br />
            <label for="hashType">Hash Type:</label>
            <select v-model="hashType">
              <option v-for="(hashOption, index) in hashOptions" :key="index" :value="hashOption.value">{{ hashOption.label }}</option>
            </select>
            <br />
            <label for="sortDirection">Sort Direction:</label>
            <select v-model="sortDirection">
              <option value="asc">Up (A-Z)</option>
              <option value="desc">Down (Z-A)</option>
            </select>
          </div>
        </div>
        <div class="search-icon" @click="toggleSearch">
          <i class="fas fa-search"></i>
        </div>
        <input type="text" class="search-bar" v-show="showSearch" @input="searchPasswords" placeholder="Search...">
      </div>
      <div class="password-list">
        <table>
          <thead>
            <tr>
              <th>Password</th>
              <th>Hash Type</th>
              <th>Created At</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(password, index) in paginatedPasswords" :key="index">
              <td>{{ password.password }}</td>
              <td>{{ password.hash_type }}</td>
              <td>{{ formatDateTime(password.created_at) }}</td>
            </tr>
          </tbody>
        </table>
        <div class="pagination">
          <button @click="prevPage" :disabled="currentPage === 1">Previous</button>
          <span>{{ currentPage }}</span>
          <button @click="nextPage" :disabled="currentPage === totalPages">Next</button>
        </div>
      </div>
      
    </div>
  </div>
</template>

<script>
export default {
  name: 'DatabaseView',
  data() {
    return {
      perPageOptions: [10, 20, 30], // Options for number of entries per page
      perPage: 20, // Default number of entries per page
      sortBy: 'created_at', // Default sort by Date
      sortDirection: 'asc', // Default sort direction
      currentPage: 1, // Current page number
      activeTab: 'generated', // Default to showing generated passwords
      passwords: [], // All fetched passwords
      isFilterActive: false,
      showFilter: false,
      showSearch: false,
      searchText: '', // Store search input value
      dateRangeFrom: '', // Date range from
      dateRangeTo: '', // Date range to
      timeRangeFrom: '', // Time range from
      timeRangeTo: '', // Time range to
      hashType: '', // Hash type selected
      hashOptions: [
        { label: 'MD5', value: 'md5' },
        { label: 'SHA-1', value: 'sha-1' },
        { label: 'SHA-256', value: 'sha-256' },
        { label: 'SHA-512', value: 'sha-512' }
      ]
    };
  },
  computed: {
    paginatedPasswords() {
      const start = (this.currentPage - 1) * this.perPage;
      const end = start + this.perPage;
      return this.passwords.slice(start, end);
    },
    totalPages() {
      return Math.ceil(this.passwords.length / this.perPage);
    }
  },
  methods: {
    toggleFilter() {
      this.showFilter = !this.showFilter;
      this.isFilterActive = this.showFilter;
    },
    toggleSearch() {
      this.showSearch = !this.showSearch;
    },
    async fetchGeneratedPasswords() {
      try {
        const response = await fetch('http://localhost:8080/generated-passwords');
        if (!response.ok) {
          throw new Error('Failed to fetch generated passwords');
        }
        const data = await response.json();
        this.passwords = data;
      } catch (error) {
        console.error(error);
      }
    },
    async fetchCreatedPasswords() {
      try {
        const response = await fetch('http://localhost:8080/passwords');
        if (!response.ok)
        {
          throw new Error('Failed to fetch created passwords');
        }
        const data = await response.json();
        this.passwords = data;
      } catch (error) {
        console.error(error);
      }
    },
    fetchPasswords() {
      if (this.activeTab === 'generated') {
        this.fetchGeneratedPasswords();
      } else if (this.activeTab === 'created') {
        this.fetchCreatedPasswords();
      }
    },
    formatDateTime(dateTimeString) {
      const options = { year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric', second: 'numeric' };
      return new Date(dateTimeString).toLocaleString(undefined, options);
    },
    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
      }
    },
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
      }
    },
    searchPasswords() {
      // Fetch passwords again to ensure the original data is preserved
      this.fetchPasswords();
      
      // Convert search text to lowercase for case-insensitive comparison
      const searchText = this.searchText.toLowerCase();
      
      // Filter passwords based on search text
      this.passwords = this.passwords.filter(password => {
        return (
          password.password.toLowerCase().includes(searchText) ||
          password.hash_type.toLowerCase().includes(searchText) ||
          password.created_at.toLowerCase().includes(searchText)
        );
      });
    }
  },
  mounted() {
    this.fetchGeneratedPasswords(); // Fetch generated passwords by default
  }
};
</script>

<style scoped>
.header {
  font-size: 36px;
  font-weight: bold;
  text-align: center;
  margin-bottom: 20px;
}

.sidebar {
  float: left;
  width: 200px;
  margin-right: 20px;
}

.sidebar ul {
  list-style-type: none;
  padding: 0;
}

.sidebar ul li {
  padding: 10px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.sidebar ul li:hover {
  background-color: #f0f0f0;
}

.sidebar ul li.active {
  background-color: #e0e0e0;
}

.options-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.options {
  display: flex;
  align-items: center;
}

.options-bar .filter-icon,
.options-bar .search-icon {
  cursor: pointer;
}

.filter-icon.active {
  color: blue; /* Change to desired color */
}

.filter-dropdown {
  position: absolute;
  top: 200px; /* Adjust as needed */
  right: 80%;
  background-color: #fff;
  border: 1px solid #ccc;
  padding: 10px;
  border-radius: 5px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  z-index: 1;
}

.search-icon {
  cursor: pointer;
  margin-left: 10px;
}

.search-bar {
  width: 200px;
  margin-left: 10px;
}

.password-list {
  width: calc(100% - 250px); /* Adjust based on the width of the filter dropdown */
}

table {
  width: 100%;
  border-collapse: collapse;
}

th,
td {
  padding: 8px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

th {
  background-color: #f2f2f2;
}

.pagination {
  margin-top: 20px;
  text-align: right;
}

.pagination button {
  margin-left: 10px;
}
</style>

