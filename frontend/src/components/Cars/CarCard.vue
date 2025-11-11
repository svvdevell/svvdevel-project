<template>
    <div class="car-card" @click="handleOpenDetails(car)" :class="{ 'selled': car.status === 'Продано' }">
        <!-- Car Image -->
        <div class="car-image" @click="handleOpenDetails">
            <img v-if="car.imageCount > 0" :src="car?.images[0]?.fileUrl" :alt="`${car.brand} ${car.model}`"
                @error="handleImageError">
            <div v-else class="no-image">
                <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none"
                    stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
                    <circle cx="8.5" cy="8.5" r="1.5"></circle>
                    <polyline points="21 15 16 10 5 21"></polyline>
                </svg>
                <span>Немає фото</span>
            </div>
        </div>

        <div class="status-badgex">
            <Badge :status="car.status" />
        </div>

        <!-- Car Info -->
        <div class="car-info">
            <div class="car-info_header">
                <h3 class="car-title">{{ car.brand }} {{ car.model }}</h3>
                <div>
                    <p>{{ formatPrice(car.price) }}</p>
                </div>
            </div>

            <div class="car-details">
                <div class="detail-row">
                    <span class="label">
                        <img src="../../assets/icons/car.png" alt="">
                        Рік:
                    </span>
                    <span class="value">{{ car.year }}</span>
                </div>

                <div class="detail-row">
                    <span class="label">
                        <img src="../../assets/icons/speedometer.png" alt="">
                        Пробіг:
                    </span>
                    <span class="value">{{ formatMileage(car.mileage) }}</span>
                </div>

                <div class="detail-row">
                    <span class="label">
                        <img src="../../assets/icons/engine.png" alt="">
                        Паливо:
                    </span>
                    <span class="value">{{ car.fuel }}</span>
                </div>

                <div class="detail-row" v-if="car.fuel != 'Електро'">
                    <span class="label">
                        <img src="../../assets/icons/volume.png" alt="">
                        Об'єм двигуна:
                    </span>
                    <span class="value">{{ formatEngineVolume(car.volume) }}</span>
                </div>

                <div class="detail-row">
                    <span class="label">
                        <img src="../../assets/icons/manual-transmission.png" alt="">
                        КПП:
                    </span>
                    <span class="value">{{ car.transmission }}</span>
                </div>

                <div class="detail-row">
                    <span class="label">
                        <img src="../../assets/icons/suspension.png" alt="">
                        Привід:
                    </span>
                    <span class="value">{{ car.drive }}</span>
                </div>

                <div class="detail-row">
                    <span class="label">
                        <img src="../../assets/icons/price.png" alt="">
                        Ціна:
                    </span>
                    <span class="value">{{ formatPrice(car.price) }}</span>
                </div>
            </div>

            <div v-if="car.description" class="car-description">
                {{ truncateDescription(car.description) }}
            </div>

            <div class="car-footer">
                <span class="date">
                    <img src="../../assets/icons/calendar.png" alt="">
                    {{ formatDate(car.createdAt) }}
                </span>
                <button class="view-btn">
                    Детальніше
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <polyline points="9 18 15 12 9 6"></polyline>
                    </svg>
                </button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { computed, ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Badge from '@/components/common/Badge.vue'
import { useHelpers } from '@/composables/useHelpers'
const { formatDate, formatPrice, formatMileage, formatEngineVolume } = useHelpers()
const router = useRouter()

const props = defineProps({
    car: {
        type: Object,
        required: true
    }
})

const emit = defineEmits(['open-details'])

// Computed
const carImageUrl = computed(() => {
    if (props.car.imageCount > 0) {
        const timestamp = getImageTimestamp(props.car.createdAt)
        return `/uploads/car_sale_${props.car.id}_image_1_${timestamp}.jpg`
    }
    return ''
})

// Methods
const handleOpenDetails = (car) => {
    router.push({ name: 'CarPage', params: { id: car.id } })
}

const handleImageError = (event) => {
    event.target.src = 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMzAwIiBoZWlnaHQ9IjIwMCIgdmlld0JveD0iMCAwIDMwMCAyMDAiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+CjxyZWN0IHdpZHRoPSIzMDAiIGhlaWdodD0iMjAwIiBmaWxsPSIjZjVmNWY1Ii8+Cjx0ZXh0IHg9IjE1MCIgeT0iMTAwIiBmaWxsPSIjOTk5OTk5IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmb250LWZhbWlseT0iQXJpYWwiIGZvbnQtc2l6ZT0iMTYiIGR5PSIuM2VtIj7QndC10YIg0YTQvtGC0L48L3RleHQ+Cjwvc3ZnPg=='
}

const truncateDescription = (description) => {
    if (!description) return ''
    return description.length > 100 ? description.substring(0, 100) + '...' : description
}

const getImageTimestamp = (createdAt) => {
    return Math.floor(new Date(createdAt).getTime() / 1000)
}
</script>

<style scoped lang="scss">
.selled {
    opacity: 0.4;
}

.car-card {
    // background: white;
    border-radius: 15px;
    overflow: hidden;
    // box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
    transition: transform 0.3s ease, box-shadow 0.3s ease;
    display: flex;
    flex-direction: column;
    height: 100%;
    position: relative;

    background: rgba(255, 255, 255, 0.4);
    backdrop-filter: blur(6px);
    -webkit-backdrop-filter: blur(6px);
    border-radius: 20px;
    border: 1px solid rgba(255, 255, 255, 0.3);
    box-shadow:
        0 8px 32px rgba(0, 0, 0, 0.1),
        inset 0 1px 0 rgba(255, 255, 255, 0.5),
        inset 0 -1px 0 rgba(255, 255, 255, 0.1),
        inset 0 0 4px 2px rgba(255, 255, 255, 0.2);
    position: relative;
    overflow: hidden;
}


.car-card::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 1px;
    background: linear-gradient(90deg,
            transparent,
            rgba(255, 255, 255, 0.8),
            transparent);
}

.car-card::after {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 1px;
    height: 100%;
    background: linear-gradient(180deg,
            rgba(255, 255, 255, 0.8),
            transparent,
            rgba(255, 255, 255, 0.3));
}


.status-badgex {
    position: absolute;
    top: 10px;
    left: 10px;
}

.car-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

.car-image {
    position: relative;
    height: 300px;
    cursor: pointer;
    overflow: hidden;
    background: #f8f9fa;
}

.car-image img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform 0.3s ease;
}

.car-image:hover img {
    transform: scale(1.05);
}

.no-image {
    width: 100%;
    height: 100%;
    background: linear-gradient(135deg, #f8f9fa 0%, #e9ecef 100%);
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    color: #6c757d;
    font-size: 1rem;
    gap: 0.5rem;
}

.no-image svg {
    opacity: 0.5;
}

.image-count {
    position: absolute;
    top: 12px;
    right: 12px;
    background: rgba(0, 0, 0, 0.75);
    backdrop-filter: blur(5px);
    color: white;
    padding: 0.4rem 0.75rem;
    border-radius: 20px;
    font-size: 0.875rem;
    font-weight: 600;
    display: flex;
    align-items: center;
    gap: 0.35rem;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
}

.image-count svg {
    width: 14px;
    height: 14px;
}

.car-info {
    padding: 1.5rem;
    display: flex;
    flex-direction: column;
    flex: 1;

    &_header {
        display: flex;
        gap: 20px;
        justify-content: space-between;
        align-items: center;
        width: 100%;
        margin-bottom: 20px;

        & div {
            display: flex;
            gap: 5px;
            align-items: center;

            & img {
                width: 25px;
                height: 25px;
                object-fit: contain;
            }

            & p {
                color: #4caf50;
                margin: 0;
                font-size: 1rem;
                font-weight: 600;
                line-height: 1.3;
            }
        }
    }
}

.car-title {
    color: #000;
    margin: 0;
    font-size: 1.5rem;
    font-weight: 600;
    line-height: 1.3;
}

.car-details {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
    margin-bottom: 1rem;
}

.detail-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.5rem 0;
    border-bottom: 1px solid #f0f0f0;
}

.detail-row:last-child {
    border-bottom: none;
}

.detail-row .label {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    color: #666;
    font-weight: 500;
    font-size: 0.95rem;
}

.detail-row .label img {
    flex-shrink: 0;
    object-fit: fill;
    width: 20px;
    object-fit: contain;
}

.detail-row .value {
    color: #000;
    font-weight: 600;
    font-size: 0.95rem;
}

.car-description {
    color: #666;
    font-size: 0.95rem;
    line-height: 1.5;
    margin-bottom: 1rem;
    padding: 1rem;
    background: #f8f9fa;
    border-radius: 8px;
}

.car-footer {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding-top: 1rem;
    border-top: 2px solid #f0f0f0;
    margin-top: auto;
}

.date {
    display: flex;
    align-items: center;
    gap: 0.4rem;
    color: #999;
    font-size: 0.875rem;
}

.date img {
    flex-shrink: 0;
    object-fit: fill;
    width: 20px;
    object-fit: contain;
}

.view-btn {
    background: linear-gradient(135deg, #007bff 0%, #0056b3 100%);
    color: white;
    border: none;
    padding: 0.6rem 1.25rem;
    border-radius: 8px;
    cursor: pointer;
    font-weight: 600;
    font-size: 0.95rem;
    display: flex;
    align-items: center;
    gap: 0.4rem;
    transition: all 0.3s ease;
    box-shadow: 0 2px 8px rgba(0, 123, 255, 0.3);
}

.view-btn:hover {
    background: linear-gradient(135deg, #0056b3 0%, #004085 100%);
    box-shadow: 0 4px 12px rgba(0, 123, 255, 0.4);
    transform: translateY(-2px);
}

.view-btn svg {
    transition: transform 0.3s ease;
}

.view-btn:hover svg {
    transform: translateX(3px);
}

@media (max-width: 768px) {
    .car-image {
        height: 250px;
    }

    .car-info {
        padding: 1.25rem;
    }

    .car-title {
        font-size: 1.3rem;
    }

    .detail-row {
        padding: 0.4rem 0;
    }

    .detail-row .label,
    .detail-row .value {
        font-size: 0.9rem;
    }

    .car-footer {
        flex-direction: column;
        gap: 1rem;
        align-items: stretch;
    }

    .view-btn {
        width: 100%;
        justify-content: center;
    }

    .date {
        justify-content: center;
    }
}

@media (max-width: 480px) {
    .car-title {
        font-size: 1.2rem;
    }
}
</style>