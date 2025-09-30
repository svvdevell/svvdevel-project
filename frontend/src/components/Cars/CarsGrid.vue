<template>
    <section class="cars-grid">
        <CarCard :car="car" v-for="car in cars" />
    </section>
</template>

<script setup>
import { ref } from 'vue';
import CarCard from './CarCard.vue';

const cars = ref([
    {
        "id": 2,
        "brand": "Mercedes",
        "model": "E63",
        "year": 2020,
        "fuel": "Дизель",
        "transmission": "Механическая",
        "drive": "Полный",
        "mileage": 20000,
        "description": "test mers",
        "createdAt": "2025-09-25T13:56:39Z",
        "updatedAt": "2025-09-25T13:56:39Z",
        "imageCount": 3
    },
    {
        "id": 1,
        "brand": "BMW",
        "model": "X5",
        "year": 2006,
        "fuel": "Гибрид",
        "transmission": "Робот",
        "drive": "Задний",
        "mileage": 1022000,
        "description": "test test descrip",
        "createdAt": "2025-09-25T13:56:02Z",
        "updatedAt": "2025-09-25T13:56:02Z",
        "imageCount": 9
    }
]);
</script>
<style lang="scss" scoped>
.cars-grid {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    gap: 20px;
}
</style>