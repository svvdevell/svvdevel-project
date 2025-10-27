<template>
    <div class="container">
        <CarsGrid />
    </div>
</template>
<script setup>
import CarsGrid from '@/components/Cars/CarsGrid.vue';
import { onMounted } from 'vue';
import { useSeo } from '@/composables/useSeo';

const { setMeta } = useSeo();

onMounted(() => {
    setMeta({
        title: 'Каталог автомобілів - Elegance Auto Одеса | Продаж авто',
        description: 'Купити авто в Одесі після викупу. Великий вибір перевірених автомобілів з повною історією. Авто після trade-in за вігідними цінами. Купить авто в Одессе.',
        keywords: 'каталог авто Одеса, купити авто Одеса, продаж авто Одеса, купить машину Одесса, автомобілі в наявності, б/у авто Одеса, авто после выкупа, trade-in Одеса',
        url: 'https://eleganceauto.od.ua/catalog',
        ogImage: 'https://eleganceauto.od.ua/images/og-catalog.jpg'
    });
});
</script>
<style lang="scss" scoped>
.container {
    width: 100%;

}
</style>
