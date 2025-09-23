<template>
  <div id="app">
    <!-- Навигация -->
    <Header />

    <!-- Основное содержимое -->
    <main class="main-content">
      <!-- Здесь отображаются компоненты в зависимости от маршрута -->
      <router-view></router-view>
    </main>

    <!-- Футер -->
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