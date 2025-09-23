<template>
  <div id="app">
    <header>
      <Header />
    </header>

    <main class="main-content">
      <router-view></router-view>
    </main>

    <footer class="footer">
      <p>&copy; 2025 Мой Сайт. Все права защищены.</p>
    </footer>
  </div>
</template>

<script setup>
import Header from './components/Header.vue'
</script>

<style scoped lang="scss">
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

#app {
  font-family: "Montserrat", sans-serif;
}

</style>