<template>
  <div id="app">
    <Header />

    <main class="main-content">
      <router-view></router-view>
    </main>

    <Footer />
  </div>
</template>

<script setup>
import Header from './components/global/Header.vue'
import Footer from './components/global/Footer.vue'

</script>

<style lang="scss">
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

#app {
  background-color: #FFF;

  & * {
    font-family: "Work Sans", sans-serif;
  }
}
</style>