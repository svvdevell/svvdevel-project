<template>
    <div v-if="car.status && car.status !== 'active'" class="status-badge" :class="`status-${car.status}`">
        <h6>{{ getStatusLabel(car.status) }}</h6>
    </div>
</template>

<script setup>

const props = defineProps({
    badge: {
        type: String,
        required: false
    }
});
const getStatusLabel = (status) => {
    const labels = {
        'sold': 'Продано',
        'new': 'Новинка',
        'sale': 'Знижка',
        'super-price': 'Супер ціна'
    }
    return labels[status] || status
}
</script>

<style lang="scss" scoped>
.status-badge {
    padding: 0.6rem 1.2rem;
    border-radius: 4px;

    & h6 {
        font-size: 0.9rem;
        font-weight: 600;
        color: white;
        text-transform: uppercase;
    }
}

.status-sold {
    background: #dc3545;
}

.status-new {
    background: #28a745;
}

.status-sale {
    background: #ffc107;
    color: #333;
}

.status-super-price {
    background: #ff6b6b;
}
</style>