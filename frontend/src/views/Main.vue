<template>
    <div class="main-page">
        <section class="top-section">
            <div class="video-container">
                <video src="../media/video/aston.mp4" autoplay muted loop></video>
            </div>
            <div class="main-banner">
                <MainBanner />
            </div>
        </section>
    </div>
</template>

<script setup>
import MainBanner from '../components/MainBanner.vue'

</script>

<style scoped lang="scss">
.main-page {
    width: 100%;
}

.global-wrapper {}

.top-section {
    position: relative;
    width: 100%;
    height: 100%;
}

.main-banner {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%,-50%);
    z-index: 2;

    & * {
        font-family: "Zalando Sans SemiExpanded", sans-serif;
    }
}

.video-container {
    width: 100vw;
    height: 100vh;

    & video {
        z-index: -1;
        width: 100vw;
    }
}
</style>