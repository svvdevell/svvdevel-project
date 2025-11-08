<template>
    <router-link to="/contact" aria-label="Link to Route" class="link-form">
        <img src="../../assets/video/ringing.gif" alt="Ringing Phone icon">
    </router-link>
</template>
<script setup>
import { useRoute } from 'vue-router'
// const route = useRoute()

</script>
<style scoped lang="scss">
.link-form {
    text-decoration: none;
    background-color: #27ae60;
    border-radius: 50%;
    padding: 1rem;
    color: #000;
    position: fixed;
    bottom: 20px;
    right: 20px;
    z-index: 9999;

    & img {
        width: 48px;
        height: 48px;
        object-fit: contain;
    }
}
</style>