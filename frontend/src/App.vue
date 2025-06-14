<template>
  <div id="app">
    <header class="app-header">
      <h1>이슈 관리 시스템</h1>
    </header>
    <main class="app-container">
      <router-view />
    </main>
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>

<style scoped>
body {
  font-family: "Noto Sans KR", sans-serif;
  background-color: #f8f9fa;
  color: #333;
  margin: 0;
  padding: 0;
}

.container {
  max-width: 800px;
  width: 100%;
  background: white;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0px 2px 8px rgba(0, 0, 0, 0.1);
  margin: 20px auto;
}

button {
  background: #007aff;
  color: white;
  padding: 10px 15px;
  border: none;
  cursor: pointer;
  border-radius: 5px;
  margin-top: 10px;
  transition: background 0.3s;
}

button:hover {
  background: #005fcc;
}

button:disabled {
  background: #ccc;
}

select, input, textarea {
  width: 100%;
  padding: 10px;
  margin-top: 5px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

label {
  font-weight: bold;
  display: block;
  margin-top: 10px;
}
</style>