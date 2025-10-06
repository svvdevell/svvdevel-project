<template>
    <div class="main-page">
        <MainBanner />
        <div class="wrapper">
            <BagsBanner />
            <ImageBanner />
            <FactsBanner />
            <ContactBanner />
        </div>
    </div>
</template>

<script setup>
import MainBanner from '@/components/banners/MainBanner.vue'
import BagsBanner from '@/components/banners/BagsBanner.vue'
import ImageBanner from '@/components/banners/ImageBanner.vue'
import FactsBanner from '@/components/banners/FactsBanner.vue'
import ContactBanner from '@/components/banners/ContactBanner.vue'


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

.wrapper {
    width: 100%;
    max-width: 1440px;
    margin: 0 auto;
    padding: 0 16px;
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