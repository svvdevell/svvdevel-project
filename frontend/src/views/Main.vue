<template>
    <div class="main-page">
        <MainBanner />
        <div class="wrapper">
            <BagsBanner />
            <ImageBanner />
            <FactsBanner />
            <ParkBanner />
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
import ParkBanner from '@/components/banners/ParkBanner.vue'

import { onMounted } from 'vue';
import { useSeo } from '@/composables/useSeo';

const { setMeta } = useSeo();

onMounted(() => {
    setMeta({
        title: 'Elegance Auto - Викуп та продаж автомобілів в Одесі',
        description: 'Вигідний викуп вашого автомобіля та продаж якісних авто в Одесі',
        keywords: 'викуп авто Одеса, продаж автомобілів Одеса',
        url: 'https://eleganceauto.od.ua/'
    });
});
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
</style>