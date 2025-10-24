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
        title: 'Каталог автомобілів - Elegance Auto',
        description: 'Великий вибір якісних автомобілів в Одесі',
        url: 'https://eleganceauto.od.ua/catalog'
    });
});
</script>
<style lang="scss" scoped>
.container {
    width: 100%;

}
</style>
