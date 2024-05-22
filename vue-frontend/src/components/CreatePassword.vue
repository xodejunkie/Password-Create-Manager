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
          <option value="bcrypt">Bcrypt</option>
        </select>
      </div>
      <button type="submit">Submit</button>
    </form>
    <!-- Progress Bar -->
    <div v-if="isHashing" class="progress">
      <div class="progress-bar" :style="{ width: progress + '%' }"></div>
    </div>
    <!-- Notification Alert -->
    <div v-if="showNotification" class="notification" @click="handlePopup">
      Password hashed 🎉
    </div>
    <!-- Popup -->
    <div v-if="showPopup" class="modal">
      <div class="modal-content">
        <span class="close" @click="closePopup">&times;</span>
        <h2>Result</h2>
        <p>This is your <span class="highlight">{{ password }}</span> hashed in <span class="highlight">{{ hashType }}</span>.</p>
        <div class="password-container">
          <input
            type="text"
            :value="showPassword ? hashedPassword : '**********'"
            ref="hashedPasswordInput"
            readonly
          />
          <button @click="togglePasswordView">
            <i :class="eyeIconClass"></i>
          </button>
          <button @click="copyToClipboard">
            <i class="fa fa-copy"></i>
          </button>
        </div>
      </div>
    </div>
    <audio ref="notificationSound" src="/cheers.wav" preload="auto"></audio>
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
      progress: 0,
      showNotification: false,
      showPopup: false,
      hashedPassword: '',
      showPassword: false
    };
  },
  computed: {
    eyeIconClass() {
      return this.showPassword ? 'fa fa-eye-slash' : 'fa fa-eye';
    }
  },
  methods: {
    async handleSubmit() {
      // Reset progress
      this.progress = 0;
      this.isHashing = true;
      this.showNotification = false;
      this.showPopup = false;

      // Simulate hashing process
      for (let i = 0; i < 100; i++) {
        this.progress = i;
        await this.sleep(20);
      }

      // Hash password and display notification when completed
      try {
        const response = await fetch('http://localhost:8080/create-password', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            password: this.password,
            hash_type: this.hashType,
          }),
        });

        if (response.ok) {
          const data = await response.json();
          this.hashedPassword = data.hashedPassword;
          this.showNotification = true;
          this.$refs.notificationSound.play();

          setTimeout(() => {
            this.showNotification = false;
          }, 3000);

          this.showPopup = true;
        } else {
          console.error('Failed to hash password');
        }
      } catch (error) {
        console.error('Error:', error);
      } finally {
        this.isHashing = false;
        this.progress = 100;
        this.password = '';
      }
    },
    sleep(ms) {
      return new Promise(resolve => setTimeout(resolve, ms));
    },
    handlePopup() {
      this.showPopup = true;
      this.showNotification = false;
    },
    closePopup() {
      this.showPopup = false;
    },
    togglePasswordView() {
      this.showPassword = !this.showPassword;
    },
    copyToClipboard() {
      const el = document.createElement('textarea');
      el.value = this.hashedPassword;
      document.body.appendChild(el);
      el.select();
      document.execCommand('copy');
      document.body.removeChild(el);
      alert('Copied to clipboard');
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

.notification {
  position: fixed;
  top: 70px;
  right: 20px;
  background-color: #ff9800;
  color: white;
  padding: 10px 20px;
  border-radius: 5px;
  cursor: pointer;
  transition: opacity 0.3s;
}

/* Updated Modal Styles */
.modal {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: rgba(255, 255, 255, 0.95); /* Transparent white background */
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1); /* Drop shadow */
  z-index: 1000; /* Ensure modal is on top */
  width: 300px; /* Adjust width as needed */
}

.modal-content {
  text-align: left; /* Align text to the left */
}

.modal-content h2 {
  margin-bottom: 10px;
}

/* Rectangular box for hashed password */
.password-container {
  margin-top: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
  padding: 10px;
  background-color: #f9f9f9;
  display: flex; /* */
  align-items: center; /* Vertically center items */
}

.password-container input[type="text"] {
  flex: 1; /* Take up remaining space */
  padding: 5px;
  border: none;
  background-color: transparent;
}

.password-container button {
  margin-left: 10px; /* Add some spacing between the text field and buttons */
  padding: 5px;
  border: none;
  background-color: #42b983;
  color: white;
  border-radius: 5px;
  cursor: pointer;
}

.password-container button:hover {
  background-color: #36876c;
}

/* Close button style */
.close {
  position: absolute;
  top: 10px;
  right: 10px;
  font-size: 20px;
  cursor: pointer;
}

.close:hover {
  color: #ff0000;
}
</style>
