<template>
  <div id="app">
    <header>
      <nav>
        <!-- Logo -->
        <div class="logo">
          <img src="/assets/logo.png" alt="Logo" />
        </div>

        <!-- Hamburger Menu for Mobile -->
        <div class="menu-toggle" @click="toggleMenu">
          <div class="bar"></div>
          <div class="bar"></div>
          <div class="bar"></div>
        </div>

        <!-- Navigation Links -->
        <div class="nav-links" :class="{ 'active': isMenuOpen }">
          <router-link to="/">Home</router-link>
          <router-link to="/create-password">Create Password</router-link>
          <router-link to="/generate-password">Generate Password</router-link>
        <router-link to="/database">View Database</router-link>
          <!-- Add other navigation links as needed -->
        </div>
      </nav>
    </header>
    <main>
      <router-view></router-view>
    </main>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      isMenuOpen: false
    };
  },
  methods: {
    toggleMenu() {
      this.isMenuOpen = !this.isMenuOpen;
    }
  }
};
</script>

<style>
/* Navbar Styles */
header {
  background-color: #42b983;
  padding: 10px;
}

nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo img {
  width: 50px; /* Adjust size as needed */
}

.menu-toggle {
  display: none; /* Hide by default on larger screens */
  cursor: pointer;
}

.bar {
  width: 25px;
  height: 3px;
  background-color: #fff;
  margin: 5px 0;
}

.nav-links {
  display: flex;
}

.nav-links a {
  color: white;
  text-decoration: none;
  margin-left: 20px;
}

/* Active class for mobile menu */
.nav-links.active {
  display: flex;
  flex-direction: column;
  position: absolute;
  top: 70px;
  left: 0;
  width: 100%;
  background-color: #42b983;
}

.nav-links.active a {
  margin: 10px;
}

@media screen and (max-width: 768px) {
  .menu-toggle {
    display: block; /* Show on smaller screens */
  }

  .nav-links {
    display: none; /* Hide on smaller screens */
  }

  .nav-links.active {
    display: flex; /* Show when mobile menu is active */
  }
}
</style>
