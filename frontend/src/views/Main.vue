<template>
    <div class="main-page">
        <MainBanner />
    </div>
</template>

<script setup>
import MainBanner from '@/components/banners/MainBanner.vue'
</script>

<style scoped lang="scss">
.main-page {
    width: 100%;
}

.top-section {
    position: relative;
    width: 100%;
    height: 100vh;
    overflow: hidden;
}

.main-banner {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    z-index: 2;
    width: 100%;

    & * {
        font-family: "Zalando Sans SemiExpanded", sans-serif;
    }
}

.video-container {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;

    & video {
        width: 100%;
        height: 100%;
        object-fit: cover;
        z-index: 1;
    }
}
</style>