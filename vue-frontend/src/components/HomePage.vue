<template>
  <div class="home-page">
    <div class="illustration">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" width="200" height="200">
        <!-- Background -->
        <rect width="512" height="512" fill="#f0f0f0"/>
        
        <!-- Person -->
        <g fill="#333">
          <path d="M256 0C114.85 0 0 114.85 0 256s114.85 256 256 256 256-114.85 256-256S397.15 0 256 0zm0 480c-123.71 0-224-100.29-224-224S132.29 32 256 32s224 100.29 224 224-100.29 224-224 224z"/>
          <path d="M256 96c-58.88 0-107.24 48.14-107.24 107.24v10.32c0 22.01-17.89 39.9-39.9 39.9s-39.9-17.89-39.9-39.9V203.9C68.76 144.14 122.89 90 189.76 90c47.12 0 91.43 23.7 117.24 63.43 25.81-39.73 70.12-63.43 117.24-63.43 66.87 0 121 54.14 121 121.24v10.32c0 22.01-17.89 39.9-39.9 39.9s-39.9-17.89-39.9-39.9v-10.32C363.24 144.14 314.88 96 256 96zm-64 112c-17.67 0-32-14.33-32-32s14.33-32 32-32 32 14.33 32 32-14.33 32-32 32zm128 224c-38.29 0-73.01-18.98-94.12-50.83 12.11-8.52 21.19-20.45 25.25-35.3H231.87c4.07 14.85 13.15 26.78 25.25 35.3-21.11 31.85-55.83 50.83-94.12 50.83-55.23 0-100-44.77-100-100s44.77-100 100-100c6.59 0 12.81.72 19 1.78V176c0-8.84 7.16-16 16-16s16 7.16 16 16v144.82c6.19-1.07 12.41-1.78 19-1.78 55.23 0 100 44.77 100 100s-44.77 100-100 100z"/>
        </g>

        <!-- Thought Bubble -->
        <g fill="#6c757d">
          <path d="M256 96c-17.67 0-32-14.33-32-32H96c0 53.02 42.98 96 96 96s96-42.98 96-96c0 17.67-14.33 32-32 32z"/>
          <path d="M288 64H96c0 70.58 57.42 128 128 128s128-57.42 128-128H288z"/>
        </g>
        
        <!-- Options or Ideas (Example) -->
        <text x="80" y="300" font-family="Arial, sans-serif" font-size="24" fill="#333">Option 1</text>
        <text x="240" y="360" font-family="Arial, sans-serif" font-size="24" fill="#333">Option 2</text>
        <text x="400" y="420" font-family="Arial, sans-serif" font-size="24" fill="#333">Option 3</text>
      </svg>

    </div>
    <div class="content">
      <h1>Welcome to the Password Manager</h1>
      <p>Select an option from the menu to get started.</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HomePage',
};
</script>

<style scoped>
.home-page {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
}

.illustration {
  margin-bottom: 30px;
}

.content {
  text-align: center;
}

h1 {
  font-size: 36px;
  margin-bottom: 10px;
}

p {
  font-size: 18px;
}
</style>
