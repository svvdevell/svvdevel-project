<template>
    <div class="car-card">
        <!-- Car Image -->
        <div class="car-image" @click="handleOpenDetails">
            <img v-if="car.imageCount > 0" :src="carImageUrl" :alt="`${car.brand} ${car.model}`"
                @error="handleImageError">
            <div v-else class="no-image">
                <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none"
                    stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
                    <circle cx="8.5" cy="8.5" r="1.5"></circle>
                    <polyline points="21 15 16 10 5 21"></polyline>
                </svg>
                <span>Нет фото</span>
            </div>
            <div v-if="car.imageCount > 1" class="image-count">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                    stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
                    <circle cx="8.5" cy="8.5" r="1.5"></circle>
                    <polyline points="21 15 16 10 5 21"></polyline>
                </svg>
                {{ car.imageCount }}
            </div>
        </div>

        <!-- Car Info -->
        <div class="car-info">
            <h3 class="car-title">{{ car.brand }} {{ car.model }}</h3>

            <div class="car-details">
                <div class="detail-row">
                    <span class="label">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect>
                            <line x1="16" y1="2" x2="16" y2="6"></line>
                            <line x1="8" y1="2" x2="8" y2="6"></line>
                            <line x1="3" y1="10" x2="21" y2="10"></line>
                        </svg>
                        Год:
                    </span>
                    <span class="value">{{ car.year }}</span>
                </div>

                <div class="detail-row">
                    <span class="label">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <circle cx="12" cy="12" r="10"></circle>
                            <polyline points="12 6 12 12 16 14"></polyline>
                        </svg>
                        Пробег:
                    </span>
                    <span class="value">{{ formatMileage(car.mileage) }}</span>
                </div>

                <div class="detail-row">
                    <span class="label">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <path d="M12 2v20M2 12h20"></path>
                        </svg>
                        Топливо:
                    </span>
                    <span class="value">{{ car.fuel }}</span>
                </div>

                <div class="detail-row">
                    <span class="label">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <circle cx="12" cy="12" r="10"></circle>
                            <line x1="12" y1="8" x2="12" y2="12"></line>
                            <line x1="12" y1="16" x2="12.01" y2="16"></line>
                        </svg>
                        КПП:
                    </span>
                    <span class="value">{{ car.transmission }}</span>
                </div>

                <div class="detail-row">
                    <span class="label">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <circle cx="12" cy="12" r="10"></circle>
                            <line x1="2" y1="12" x2="22" y2="12"></line>
                        </svg>
                        Привод:
                    </span>
                    <span class="value">{{ car.drive }}</span>
                </div>
            </div>

            <div v-if="car.description" class="car-description">
                {{ truncateDescription(car.description) }}
            </div>

            <div class="car-footer">
                <span class="date">
                    <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <circle cx="12" cy="12" r="10"></circle>
                        <polyline points="12 6 12 12 16 14"></polyline>
                    </svg>
                    {{ formatDate(car.createdAt) }}
                </span>
                <button @click="handleOpenDetails" class="view-btn">
                    Подробнее
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
import { computed } from 'vue'

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
const handleOpenDetails = () => {
    emit('open-details', props.car)
}

const handleImageError = (event) => {
    event.target.src = 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMzAwIiBoZWlnaHQ9IjIwMCIgdmlld0JveD0iMCAwIDMwMCAyMDAiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+CjxyZWN0IHdpZHRoPSIzMDAiIGhlaWdodD0iMjAwIiBmaWxsPSIjZjVmNWY1Ii8+Cjx0ZXh0IHg9IjE1MCIgeT0iMTAwIiBmaWxsPSIjOTk5OTk5IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmb250LWZhbWlseT0iQXJpYWwiIGZvbnQtc2l6ZT0iMTYiIGR5PSIuM2VtIj7QndC10YIg0YTQvtGC0L48L3RleHQ+Cjwvc3ZnPg=='
}

const formatMileage = (mileage) => {
    return new Intl.NumberFormat('ru-RU').format(mileage) + ' км'
}

const formatDate = (dateString) => {
    const date = new Date(dateString)
    return date.toLocaleDateString('ru-RU')
}

const truncateDescription = (description) => {
    if (!description) return ''
    return description.length > 100 ? description.substring(0, 100) + '...' : description
}

const getImageTimestamp = (createdAt) => {
    return Math.floor(new Date(createdAt).getTime() / 1000)
}
</script>

<style scoped>
.car-card {
    background: white;
    border-radius: 15px;
    overflow: hidden;
    box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
    transition: transform 0.3s ease, box-shadow 0.3s ease;
    display: flex;
    flex-direction: column;
    height: 100%;
}

.car-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

.car-image {
    position: relative;
    height: 220px;
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
}

.car-title {
    color: #1a1a1a;
    margin: 0 0 1.25rem 0;
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

.detail-row .label svg {
    color: #007bff;
    flex-shrink: 0;
}

.detail-row .value {
    color: #1a1a1a;
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
    border-left: 3px solid #007bff;
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

.date svg {
    color: #999;
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
        height: 200px;
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
    .car-image {
        height: 180px;
    }

    .car-title {
        font-size: 1.2rem;
    }
}
</style>