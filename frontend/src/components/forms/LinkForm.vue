<template>
    <a href="tel:0734080999" class="phone-call-link" aria-label="Link to Route" :class="{ 'is-ringing': isRinging }"
        @mouseenter="isRinging = true" @mouseleave="isRinging = false">
        <div class="phone-icon">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path
                    d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z">
                </path>
            </svg>
        </div>
    </a>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';

const isRinging = ref(false);
let intervalId = null;

// Автоматическая анимация каждые 5 секунд
onMounted(() => {
    intervalId = setInterval(() => {
        isRinging.value = true;
        setTimeout(() => {
            isRinging.value = false;
        }, 2000);
    }, 5000);
});

onUnmounted(() => {
    if (intervalId) {
        clearInterval(intervalId);
    }
});
</script>

<style scoped>
.phone-call-link {
    position: fixed;
    bottom: 30px;
    right: 30px;
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 15px 20px;
    background: #25D366;
    color: white;
    text-decoration: none;
    border-radius: 50px;
    box-shadow: 0 4px 12px rgba(37, 211, 102, 0.4);
    transition: all 0.3s ease;
    z-index: 1000;
    font-weight: 600;
    font-size: 16px;
}

.phone-call-link:hover {
    background: #20BA5A;
    box-shadow: 0 6px 20px rgba(37, 211, 102, 0.6);
    transform: translateY(-2px);
}

.phone-icon {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 24px;
    height: 24px;
}

.phone-icon svg {
    width: 100%;
    height: 100%;
}

/* Анимация звонка */
.phone-call-link.is-ringing .phone-icon {
    animation: ring 0.5s ease-in-out infinite;
}

.phone-call-link.is-ringing {
    animation: pulse 1.5s ease-in-out infinite;
}

@keyframes ring {

    0%,
    100% {
        transform: rotate(0deg);
    }

    10%,
    30% {
        transform: rotate(-15deg);
    }

    20%,
    40% {
        transform: rotate(15deg);
    }

    50% {
        transform: rotate(0deg);
    }
}

@keyframes pulse {

    0%,
    100% {
        box-shadow: 0 4px 12px rgba(37, 211, 102, 0.4);
    }

    50% {
        box-shadow: 0 4px 20px rgba(37, 211, 102, 0.8), 0 0 0 10px rgba(37, 211, 102, 0.3);
    }
}

/* Адаптив для мобильных */
@media (max-width: 768px) {
    .phone-call-link {
        bottom: 20px;
        right: 20px;
        padding: 12px 16px;
    }

    .phone-text {
        display: none;
    }

    .phone-call-link {
        width: 56px;
        height: 56px;
        border-radius: 50%;
        padding: 0;
        justify-content: center;
    }
}
</style>
