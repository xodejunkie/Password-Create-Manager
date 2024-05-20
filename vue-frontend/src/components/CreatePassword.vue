<template>
  <div class="create-password">
    <h2>Create Password</h2>
    <form @submit.prevent="handleSubmit">
      <div class="form-group">
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="password" required />
      </div>
      <div class="form-group">
        <label for="hash">Hash:</label>
        <select id="hash" v-model="hashType" required>
          <option value="md5">MD5</option>
          <option value="sha1">SHA-1</option>
          <option value="sha256">SHA-256</option>
          <!-- Add other hash types as needed -->
        </select>
      </div>
      <button type="submit">Submit</button>
    </form>
    <!-- Progress Bar -->
    <div v-if="isHashing" class="progress">
      <div class="progress-bar" :style="{ width: progress + '%' }"></div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CreatePassword',
  data() {
    return {
      password: '',
      hashType: 'md5',
      isHashing: false,
      progress: 0
    };
  },
  methods: {
    async handleSubmit() {
      // Reset progress
      this.progress = 0;
      this.isHashing = true;

      // Simulate hashing process
      for (let i = 0; i < 100; i++) {
        this.progress = i;
        await this.sleep(50); // Simulate hashing delay
      }

      // Reset form after hashing
      this.isHashing = false;
      this.progress = 0;
      this.password = '';
    },
    sleep(ms) {
      return new Promise(resolve => setTimeout(resolve, ms));
    }
  }
};
</script>

<style scoped>
.create-password {
  max-width: 400px;
  margin: 0 auto;
}

.form-group {
  margin-bottom: 20px;
}

label {
  display: block;
  margin-bottom: 5px;
}

input[type="password"],
select {
  width: 100%;
  padding: 10px;
  font-size: 16px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

button {
  background-color: #42b983;
  color: white;
  border: none;
  padding: 10px 20px;
  font-size: 16px;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
}

button:hover {
  background-color: #36876c;
}

.progress {
  background-color: #eee;
  border-radius: 5px;
  height: 20px;
  overflow: hidden;
  margin-top: 20px;
}

.progress-bar {
  background-color: #42b983;
  height: 100%;
}
</style>
