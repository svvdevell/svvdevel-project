<template>
  <div id="app">
    <Header />

    <main class="main-content">
      <router-view></router-view>
    </main>
    <LinkForm />
    <Footer />
  </div>
</template>

<script setup>
import Header from './components/global/Header.vue'
import Footer from './components/global/Footer.vue'
import LinkForm from './components/forms/LinkForm.vue'

</script>

<style lang="scss">
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  color: #FFF;
}

#app {
  background: rgba(26, 26, 26, .95);
  // padding: 0 16px;
  // width: 100%;
  // max-width: 1440px;
  // margin: 0 auto;

  & * {
    font-family: 'SF Pro Text';
  }
}

</style>