<template>
    <section class="banner">
        <div class="video">
            <video src="../../assets/video/1.mp4"></video>
        </div>
        <div class="block">
            <TextBanner />
            <MainForm />
        </div>
    </section>
</template>

<script setup>
import MainForm from '@/components/forms/MainForm.vue'
import TextBanner from '@/components/banners/TextBanner.vue';

</script>

<style lang="scss" scoped>
.banner {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 40px;
    width: 100%;
    // background: #aa3535;
    // border-radius: 40px;
    // margin: 50px 0;
    // padding: 78px 72px;
    position: relative;
    overflow: hidden;
}

.video {
    width: 100%;
    height: 100%;

    & video {
        width: 100vw;
        height: 100vh;
        object-fit: cover;
        filter: brightness(0.7) blur(5px);
    }
}

.block {
    width: 100%;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 20px;
    max-width: 1440px;
}
</style>