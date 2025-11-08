<template>
    <section class="banner">
        <div class="video">
            <img v-if="isMobile && !videoLoaded" :src="placeholderMobile" alt="Loading..." class="video-placeholder" />

            <video v-if="!isMobile" ref="desktopVideoRef" src="../../assets/video/1.mp4" autoplay muted loop playsinline
                preload="auto"></video>

            <video v-else ref="mobileVideoRef" src="../../assets/video/mob_ferrari.mp4" autoplay muted loop playsinline
                preload="auto" :class="{ 'video-visible': videoLoaded }" @loadeddata="onVideoLoaded"></video>
        </div>
        <div class="block">
            <TextBanner />
            <MainForm />
        </div>
    </section>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import MainForm from '@/components/forms/MainForm.vue'
import TextBanner from '@/components/banners/TextBanner.vue';
import { useScreenSize } from '@/composables/useScreenSize';

const { isMobile } = useScreenSize();
const desktopVideoRef = ref(null);
const mobileVideoRef = ref(null);
const videoLoaded = ref(false);

// Путь к плейсхолдеру для мобилки (замени на свой путь)
const placeholderMobile = ref('../../assets/images/ferrari.png');

// Обработчик загрузки видео
const onVideoLoaded = () => {
    videoLoaded.value = true;
};

// Принудительный запуск видео для iOS
onMounted(() => {
    const videos = document.querySelectorAll('video');
    videos.forEach(video => {
        video.play().catch(err => {
            console.log('Video autoplay prevented:', err);
        });
    });
});
</script>

<style lang="scss" scoped>
.banner {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 40px;
    position: relative;
    overflow: hidden;
}

.video {
    width: 100%;
    height: 100%;
    position: absolute;
    top: 0;
    left: 0;

    & video,
    & .video-placeholder {
        width: 100%;
        height: 100vh;
        object-fit: cover;
        filter: brightness(0.7) blur(5px);
    }

    & video {
        opacity: 1;
    }

    & .video-placeholder {
        position: absolute;
        top: 0;
        left: 0;
        z-index: 1;
    }
}

.block {
    width: 100%;
    position: relative;
    z-index: 2;
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 20px;
    max-width: 1440px;
    margin: 0 auto;
    padding: 80px 20px;
    min-height: 100vh;
    align-items: center;
}

@media (max-width: 768px) {
    .block {
        display: flex;
        flex-direction: column;
        gap: 20px;
        padding: 20px 16px;
        padding-bottom: 40px;
        justify-content: center;
    }

    .video video,
    .video .video-placeholder {
        filter: brightness(0.9) blur(5px);
    }

    .video video {
        opacity: 0;
        transition: opacity 0.5s ease-in-out;
        position: absolute;
        top: 0;
        left: 0;
    }

    .video video.video-visible {
        opacity: 1;
    }
}
</style>