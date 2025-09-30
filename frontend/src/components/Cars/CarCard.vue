<template>
    <div class="card">
        <img src="./" alt="">
        <h2>{{ car.brand }} {{ car.model }}</h2>
        <h3>{{ car.year }} {{ car.mileage }} км</h3>
        <div>
            <p>{{ car.fuel }}</p>
            <p>{{ car.transmission }}</p>
        </div>
        <div>
            <p>{{ car.drive }}</p>
            <p></p>
        </div>
        <h4>{{ car.description }}</h4>
    </div>
    <!-- {{ car }} -->
</template>
<script setup>

const props = defineProps({
    car: {
        type: Object,
        required: true
    }
})

</script>