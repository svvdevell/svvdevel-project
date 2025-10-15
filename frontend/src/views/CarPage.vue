<template>
    <div class="car-details-page">
        <div class="container">
            <!-- Back Button -->
            <button @click="goBack" class="back-btn">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none"
                    stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <polyline points="15 18 9 12 15 6"></polyline>
                </svg>
                Назад до каталогу
            </button>

            <!-- Loading -->
            <div v-if="loading" class="loading">
                <div class="loading-spinner"></div>
                <p>Завантажуємо інформацію...</p>
            </div>

            <!-- Error -->
            <div v-if="error" class="error">
                <p>{{ error }}</p>
                <button @click="fetchCarDetails" class="retry-btn">Спробувати знову</button>
            </div>

            <!-- Car Details -->
            <div v-if="!loading && !error && car" class="car-details-content">
                <!-- Header -->
                <div class="car-header">
                    <h1>{{ car.brand }} {{ car.model }} {{ car.year }}</h1>
                    <div class="car-status">
                        <span class="status-badge">{{ car.status === 'active' ? 'Активне' : 'Неактивне' }}</span>
                    </div>
                </div>

                <!-- Images Gallery -->
                <div v-if="car.images && car.images.length > 0" class="car-gallery">
                    <div class="main-image">
                        <img :src="car.images[currentImageIndex].fileUrl" 
                             :alt="`${car.brand} ${car.model}`"
                             @error="handleImageError">
                        
                        <!-- Navigation Arrows -->
                        <button v-if="car.images.length > 1" 
                                @click="previousImage" 
                                class="nav-arrow nav-arrow-left">
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" 
                                 fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" 
                                 stroke-linejoin="round">
                                <polyline points="15 18 9 12 15 6"></polyline>
                            </svg>
                        </button>
                        <button v-if="car.images.length > 1" 
                                @click="nextImage" 
                                class="nav-arrow nav-arrow-right">
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" 
                                 fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" 
                                 stroke-linejoin="round">
                                <polyline points="9 18 15 12 9 6"></polyline>
                            </svg>
                        </button>

                        <!-- Image Counter -->
                        <div class="image-counter">
                            {{ currentImageIndex + 1 }} / {{ car.images.length }}
                        </div>
                    </div>

                    <!-- Thumbnails -->
                    <div v-if="car.images.length > 1" class="thumbnails">
                        <div v-for="(image, index) in car.images" 
                             :key="image.id"
                             :class="['thumbnail', { active: index === currentImageIndex }]"
                             @click="currentImageIndex = index">
                            <img :src="image.fileUrl" :alt="`Фото ${index + 1}`">
                        </div>
                    </div>
                </div>

                <!-- No Images -->
                <div v-else class="no-images">
                    <svg xmlns="http://www.w3.org/2000/svg" width="64" height="64" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
                        <circle cx="8.5" cy="8.5" r="1.5"></circle>
                        <polyline points="21 15 16 10 5 21"></polyline>
                    </svg>
                    <p>Фотографії відсутні</p>
                </div>

                <!-- Car Specifications -->
                <div class="car-specs">
                    <h2>Технічні характеристики</h2>
                    <div class="specs-grid">
                        <div class="spec-item">
                            <span class="spec-label">Марка</span>
                            <span class="spec-value">{{ car.brand }}</span>
                        </div>
                        <div class="spec-item">
                            <span class="spec-label">Модель</span>
                            <span class="spec-value">{{ car.model }}</span>
                        </div>
                        <div class="spec-item">
                            <span class="spec-label">Рік випуску</span>
                            <span class="spec-value">{{ car.year }}</span>
                        </div>
                        <div v-if="car.color" class="spec-item">
                            <span class="spec-label">Колір</span>
                            <span class="spec-value">{{ car.color }}</span>
                        </div>
                        <div class="spec-item">
                            <span class="spec-label">Пробіг</span>
                            <span class="spec-value">{{ formatMileage(car.mileage) }}</span>
                        </div>
                        <div class="spec-item">
                            <span class="spec-label">Паливо</span>
                            <span class="spec-value">{{ car.fuel }}</span>
                        </div>
                        <div class="spec-item">
                            <span class="spec-label">Трансмісія</span>
                            <span class="spec-value">{{ car.transmission }}</span>
                        </div>
                        <div class="spec-item">
                            <span class="spec-label">Привід</span>
                            <span class="spec-value">{{ car.drive }}</span>
                        </div>
                        <div class="spec-item">
                            <span class="spec-label">Дата додавання</span>
                            <span class="spec-value">{{ formatDate(car.createdAt) }}</span>
                        </div>
                    </div>
                </div>

                <!-- Description -->
                <div v-if="car.description" class="car-description-section">
                    <h2>Опис</h2>
                    <p>{{ car.description }}</p>
                </div>

                <!-- Share Section -->
                <div class="share-section">
                    <h3>Поділитися оголошенням</h3>
                    <div class="share-buttons">
                        <button @click="copyLink" class="share-btn">
                            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" 
                                 fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" 
                                 stroke-linejoin="round">
                                <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path>
                                <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path>
                            </svg>
                            {{ linkCopied ? 'Скопійовано!' : 'Копіювати посилання' }}
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()

// Reactive data
const car = ref(null)
const loading = ref(false)
const error = ref('')
const currentImageIndex = ref(0)
const linkCopied = ref(false)

// Methods
const fetchCarDetails = async () => {
    loading.value = true
    error.value = ''

    try {
        const carId = route.params.id

        const apiUrl = process.env.NODE_ENV === 'production'
            ? `/api/cars-sale/${carId}`
            : `http://localhost:8001/api/cars-sale/${carId}`

        const response = await fetch(apiUrl)

        if (!response.ok) {
            throw new Error(`HTTP ${response.status}`)
        }

        const data = await response.json()

        if (data.status === 'success') {
            car.value = data.data
        } else {
            throw new Error(data.error || 'Невідома помилка')
        }
    } catch (err) {
        error.value = `Помилка завантаження: ${err.message}`
        console.error('Fetch error:', err)
    } finally {
        loading.value = false
    }
}

const goBack = () => {
    router.push({ name: 'cars-catalog' })
}

const nextImage = () => {
    if (car.value && car.value.images) {
        currentImageIndex.value = (currentImageIndex.value + 1) % car.value.images.length
    }
}

const previousImage = () => {
    if (car.value && car.value.images) {
        currentImageIndex.value = currentImageIndex.value === 0 
            ? car.value.images.length - 1 
            : currentImageIndex.value - 1
    }
}

const handleImageError = (event) => {
    event.target.src = 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iODAwIiBoZWlnaHQ9IjYwMCIgdmlld0JveD0iMCAwIDgwMCA2MDAiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+CjxyZWN0IHdpZHRoPSI4MDAiIGhlaWdodD0iNjAwIiBmaWxsPSIjZjVmNWY1Ii8+Cjx0ZXh0IHg9IjQwMCIgeT0iMzAwIiBmaWxsPSIjOTk5OTk5IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmb250LWZhbWlseT0iQXJpYWwiIGZvbnQtc2l6ZT0iMjQiIGR5PSIuM2VtIj7QndC10YIg0YTQvtGC0L48L3RleHQ+Cjwvc3ZnPg=='
}

const formatMileage = (mileage) => {
    return new Intl.NumberFormat('uk-UA').format(mileage) + ' км'
}

const formatDate = (dateString) => {
    const date = new Date(dateString)
    return date.toLocaleDateString('uk-UA', {
        year: 'numeric',
        month: 'long',
        day: 'numeric'
    })
}

const copyLink = async () => {
    try {
        await navigator.clipboard.writeText(window.location.href)
        linkCopied.value = true
        setTimeout(() => {
            linkCopied.value = false
        }, 2000)
    } catch (err) {
        console.error('Failed to copy link:', err)
    }
}

// Lifecycle
onMounted(() => {
    fetchCarDetails()
})
</script>

<style scoped>
.car-details-page {
    padding: 2rem;
    min-height: 100vh;
}

.container {
    max-width: 1440px;
    margin: 0 auto;
}

.back-btn {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.75rem 1.5rem;
    background: white;
    border: 1px solid #ddd;
    border-radius: 8px;
    cursor: pointer;
    font-size: 1rem;
    color: #333;
    margin-bottom: 2rem;
    transition: all 0.3s ease;
}

.back-btn:hover {
    background: #f8f9fa;
    border-color: #007bff;
    color: #007bff;
}

.loading,
.error {
    text-align: center;
    padding: 3rem;
    background: white;
    border-radius: 15px;
    margin: 2rem 0;
}

.loading-spinner {
    width: 50px;
    height: 50px;
    border: 4px solid #f3f3f3;
    border-radius: 50%;
    border-top-color: #007bff;
    animation: spin 1s ease-in-out infinite;
    margin: 0 auto 1rem;
}

@keyframes spin {
    to {
        transform: rotate(360deg);
    }
}

.retry-btn {
    background: #007bff;
    color: white;
    border: none;
    padding: 0.75rem 1.5rem;
    border-radius: 8px;
    cursor: pointer;
    font-size: 1rem;
    transition: background 0.3s ease;
}

.retry-btn:hover {
    background: #0056b3;
}

.car-details-content {
    background: white;
    border-radius: 15px;
    padding: 2rem;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.car-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
    padding-bottom: 1rem;
    border-bottom: 2px solid #f0f0f0;
}

.car-header h1 {
    margin: 0;
    font-size: 2rem;
    color: #333;
}

.status-badge {
    padding: 0.5rem 1rem;
    background: #28a745;
    color: white;
    border-radius: 20px;
    font-size: 0.875rem;
    font-weight: 600;
}

.car-gallery {
    margin-bottom: 2rem;
}

.main-image {
    position: relative;
    width: 100%;
    height: 500px;
    border-radius: 10px;
    overflow: hidden;
    background: #f8f9fa;
    margin-bottom: 1rem;
}

.main-image img {
    width: 100%;
    height: 100%;
    object-fit: contain;
}

.nav-arrow {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    background: rgba(0, 0, 0, 0.6);
    backdrop-filter: blur(5px);
    color: white;
    border: none;
    width: 50px;
    height: 50px;
    border-radius: 50%;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s ease;
    z-index: 10;
}

.nav-arrow:hover {
    background: rgba(0, 0, 0, 0.8);
    transform: translateY(-50%) scale(1.1);
}

.nav-arrow-left {
    left: 20px;
}

.nav-arrow-right {
    right: 20px;
}

.image-counter {
    position: absolute;
    bottom: 20px;
    right: 20px;
    background: rgba(0, 0, 0, 0.7);
    backdrop-filter: blur(5px);
    color: white;
    padding: 0.5rem 1rem;
    border-radius: 20px;
    font-size: 0.875rem;
    font-weight: 600;
}

.thumbnails {
    display: flex;
    gap: 1rem;
    overflow-x: auto;
    padding: 0.5rem 0;
}

.thumbnail {
    flex-shrink: 0;
    width: 120px;
    height: 90px;
    border-radius: 8px;
    overflow: hidden;
    cursor: pointer;
    border: 3px solid transparent;
    transition: all 0.3s ease;
    opacity: 0.6;
}

.thumbnail:hover,
.thumbnail.active {
    opacity: 1;
    border-color: #007bff;
}

.thumbnail img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.no-images {
    height: 400px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    background: #f8f9fa;
    border-radius: 10px;
    color: #999;
    margin-bottom: 2rem;
}

.no-images svg {
    margin-bottom: 1rem;
}

.car-specs {
    margin-bottom: 2rem;
}

.car-specs h2 {
    margin-bottom: 1.5rem;
    color: #333;
    font-size: 1.5rem;
}

.specs-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 1.5rem;
}

.spec-item {
    padding: 1.25rem;
    background: #f8f9fa;
    border-radius: 10px;
    border-left: 4px solid #007bff;
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
}

.spec-label {
    color: #666;
    font-size: 0.875rem;
    font-weight: 500;
    text-transform: uppercase;
    letter-spacing: 0.5px;
}

.spec-value {
    color: #333;
    font-size: 1.125rem;
    font-weight: 600;
}

.car-description-section {
    margin-bottom: 2rem;
    padding: 1.5rem;
    background: #f8f9fa;
    border-radius: 10px;
}

.car-description-section h2 {
    margin-bottom: 1rem;
    color: #333;
    font-size: 1.5rem;
}

.car-description-section p {
    color: #666;
    line-height: 1.8;
    font-size: 1rem;
    white-space: pre-wrap;
}

.share-section {
    padding: 1.5rem;
    background: #f8f9fa;
    border-radius: 10px;
}

.share-section h3 {
    margin-bottom: 1rem;
    color: #333;
}

.share-buttons {
    display: flex;
    gap: 1rem;
}

.share-btn {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.75rem 1.5rem;
    background: #007bff;
    color: white;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    font-size: 1rem;
    font-weight: 600;
    transition: all 0.3s ease;
}

.share-btn:hover {
    background: #0056b3;
    transform: translateY(-2px);
}

@media (max-width: 768px) {
    .car-details-page {
        padding: 1rem;
    }

    .car-details-content {
        padding: 1.5rem;
    }

    .car-header {
        flex-direction: column;
        align-items: flex-start;
        gap: 1rem;
    }

    .car-header h1 {
        font-size: 1.5rem;
    }

    .main-image {
        height: 300px;
    }

    .specs-grid {
        grid-template-columns: 1fr;
    }

    .nav-arrow {
        width: 40px;
        height: 40px;
    }

    .nav-arrow-left {
        left: 10px;
    }

    .nav-arrow-right {
        right: 10px;
    }
}
</style>