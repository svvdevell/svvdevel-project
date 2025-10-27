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
        title: 'Авто викуп в Одесі - Elegance Auto | Швидкий викуп автомобілів',
        description: 'Терміновий викуп авто в Одесі за вигідною ціною. Оцінка за 15 хвилин, виплата готівкою. Викупаємо авто в будь-якому стані. Выкуп авто Одесса - быстро, выгодно, надежно.',
        keywords: 'авто викуп Одеса, викуп автомобілів Одеса, терміновий викуп авто, выкуп авто Одесса, продать авто Одесса, скупка автомобилей, автовыкуп, викуп авто готівкою, продаж автомобілів Одеса, автосалон Одеса',
        url: 'https://eleganceauto.od.ua/',
        ogImage: 'https://eleganceauto.od.ua/images/og-home.jpg'
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