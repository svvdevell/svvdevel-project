<template>
    <div class="car-details-page">
        <div class="container">
            <!-- Back Button -->
            <button @click="goBack" class="back-btn">
                <img src="../assets/icons/back.png" alt="">
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
                        <Badge :status="car.status" />
                    </div>
                </div>

                <!-- Images Gallery -->
                <div v-if="car.images && car.images.length > 0" class="car-gallery">
                    <div class="main-image" @click="openLightbox(currentImageIndex)">
                        <img :src="car.images[currentImageIndex].fileUrl" :alt="`${car.brand} ${car.model}`"
                            @error="handleImageError">

                        <!-- Zoom Icon -->
                        <div class="zoom-icon">
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <circle cx="11" cy="11" r="8"></circle>
                                <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
                                <line x1="11" y1="8" x2="11" y2="14"></line>
                                <line x1="8" y1="11" x2="14" y2="11"></line>
                            </svg>
                        </div>

                        <!-- Navigation Arrows -->
                        <button v-if="car.images.length > 1" @click.stop="previousImage"
                            class="nav-arrow nav-arrow-left">
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <polyline points="15 18 9 12 15 6"></polyline>
                            </svg>
                        </button>
                        <button v-if="car.images.length > 1" @click.stop="nextImage" class="nav-arrow nav-arrow-right">
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
                        <div v-for="(image, index) in car.images" :key="image.id"
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
                    <div class="car-specs_header">
                        <h2>Технічні характеристики</h2>
                        <div>
                            <p>{{ formatPrice(car.price) }}</p>
                        </div>
                    </div>
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
                        <div class="spec-item" v-if="car.fuel != 'Електро'">
                            <span class="spec-label">Об`єм двигуна</span>
                            <span class="spec-value">{{ formatEngineVolume(car.volume) }}</span>
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
                    <div class="share-buttons">
                        <a href="tel:0734080999" class="share-btn" style="background-color: #27ae60; margin-top: 20px; text-decoration: none;">
                            Зателефонувати
                        </a>
                    </div>
                </div>
            </div>
        </div>

        <!-- Lightbox Modal -->
        <Teleport to="body">
            <Transition name="lightbox">
                <div v-if="isLightboxOpen" class="lightbox" @click="closeLightbox">
                    <button class="lightbox-close" @click="closeLightbox">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <line x1="18" y1="6" x2="6" y2="18"></line>
                            <line x1="6" y1="6" x2="18" y2="18"></line>
                        </svg>
                    </button>

                    <div class="lightbox-content" @click.stop @mousedown="handleDragStart" @mousemove="handleDragMove"
                        @mouseup="handleDragEnd" @mouseleave="handleDragEnd" @touchstart="handleTouchStart"
                        @touchmove="handleTouchMove" @touchend="handleTouchEnd">
                        <div class="lightbox-image-wrapper" :style="{ transform: `translateX(${dragOffset}px)` }">
                            <img :src="car.images[lightboxIndex].fileUrl" :alt="`${car.brand} ${car.model}`">
                        </div>

                        <!-- Navigation -->
                        <button v-if="car.images.length > 1" @click.stop="previousLightboxImage"
                            class="lightbox-nav lightbox-nav-left">
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <polyline points="15 18 9 12 15 6"></polyline>
                            </svg>
                        </button>
                        <button v-if="car.images.length > 1" @click.stop="nextLightboxImage"
                            class="lightbox-nav lightbox-nav-right">
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round">
                                <polyline points="9 18 15 12 9 6"></polyline>
                            </svg>
                        </button>

                        <!-- Counter -->
                        <div class="lightbox-counter">
                            {{ lightboxIndex + 1 }} / {{ car.images.length }}
                        </div>
                    </div>
                </div>
            </Transition>
        </Teleport>
    </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useSeo } from '@/composables/useSeo';
const { setMeta, setStructuredData } = useSeo();
import Badge from '@/components/common/Badge.vue'
import { useHelpers } from '@/composables/useHelpers'

const route = useRoute()
const router = useRouter()

// Reactive data
const car = ref(null)
const loading = ref(false)
const error = ref('')
const currentImageIndex = ref(0)
const linkCopied = ref(false)
const isLightboxOpen = ref(false)
const lightboxIndex = ref(0)
const { formatDate, formatPrice, formatMileage, formatEngineVolume } = useHelpers()

// Drag/Swipe functionality
const isDragging = ref(false)
const dragStartX = ref(0)
const dragOffset = ref(0)
const dragThreshold = 50 // Minimum distance to trigger swipe

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
        car.value = data.data

        const title = `${car.value.brand} ${car.value.model} ${car.value.year} - ${car.value.price.toLocaleString('uk-UA')} грн`;
        const description = `${car.value.brand} ${car.value.model} ${car.value.year} року, ${car.value.color}, ${car.value.fuel}, ${car.value.transmission}, пробіг ${car.value.mileage.toLocaleString('uk-UA')} км. Ціна: ${car.value.price.toLocaleString('uk-UA')} грн. ${car.value.description || ''}`.slice(0, 160);
        const keywords = `${car.value.brand} ${car.value.model} купити, ${car.value.brand} ${car.value.model} Одеса, ${car.value.brand} ціна, авто ${car.value.year}`;

        setMeta({
            title,
            description,
            keywords,
            url: `https://eleganceauto.od.ua/cars/${car.value.id}`,
            ogImage: car.value.images?.[0] ? `https://eleganceauto.od.ua${car.value.images[0].fileUrl}` : 'https://eleganceauto.od.ua/images/og-default.jpg'
        });

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
    router.push({ name: 'CarsCatalog' })
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
    event.target.src = 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iODAwIiBoZWlnaHQ9IjYwMCIgdmlld0JveD0iMCAwIDgwMCA2MDAiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+CjxyZWN0IHdpZHRoPSI4MDAiIGhlaWdodD0iNjAwIiBmaWxsPSIjZjVmNWY1Ii8+Cjx0ZXh0IHg9IjQwMCIgeT0iMzAwIiBmaWxsPSIjOTk5OTk5IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmb250LWZhbWlseT0iQXJpYWwiIGZvbnQtc2l6ZT0iMjQiIGR5PSIuM2VtIj7QndC10YIg0YTQvtGC0L88L3RleHQ+Cjwvc3ZnPg=='
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

// Lightbox functions
const openLightbox = (index) => {
    lightboxIndex.value = index
    isLightboxOpen.value = true
    document.body.style.overflow = 'hidden'
}

const closeLightbox = () => {
    isLightboxOpen.value = false
    document.body.style.overflow = ''
    dragOffset.value = 0
}

const nextLightboxImage = () => {
    if (car.value && car.value.images) {
        lightboxIndex.value = (lightboxIndex.value + 1) % car.value.images.length
    }
}

const previousLightboxImage = () => {
    if (car.value && car.value.images) {
        lightboxIndex.value = lightboxIndex.value === 0
            ? car.value.images.length - 1
            : lightboxIndex.value - 1
    }
}

// Mouse drag handlers
const handleDragStart = (e) => {
    isDragging.value = true
    dragStartX.value = e.clientX
}

const handleDragMove = (e) => {
    if (!isDragging.value) return
    const currentX = e.clientX
    dragOffset.value = currentX - dragStartX.value
}

const handleDragEnd = () => {
    if (!isDragging.value) return

    if (dragOffset.value < -dragThreshold) {
        nextLightboxImage()
    } else if (dragOffset.value > dragThreshold) {
        previousLightboxImage()
    }

    isDragging.value = false
    dragOffset.value = 0
}

// Touch handlers
const handleTouchStart = (e) => {
    dragStartX.value = e.touches[0].clientX
}

const handleTouchMove = (e) => {
    const currentX = e.touches[0].clientX
    dragOffset.value = currentX - dragStartX.value
}

const handleTouchEnd = () => {
    if (dragOffset.value < -dragThreshold) {
        nextLightboxImage()
    } else if (dragOffset.value > dragThreshold) {
        previousLightboxImage()
    }

    dragOffset.value = 0
}

// Keyboard navigation
const handleKeydown = (e) => {
    if (!isLightboxOpen.value) return

    if (e.key === 'Escape') {
        closeLightbox()
    } else if (e.key === 'ArrowLeft') {
        previousLightboxImage()
    } else if (e.key === 'ArrowRight') {
        nextLightboxImage()
    }
}

// Lifecycle
onMounted(() => {
    fetchCarDetails()
    window.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
    window.removeEventListener('keydown', handleKeydown)
    document.body.style.overflow = ''
})
</script>

<style scoped lang="scss">
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
    padding: 10px;
    background: white;
    border: 1px solid #ddd;
    border-radius: 8px;
    cursor: pointer;
    font-size: 1rem;
    color: #333;
    margin-bottom: 2rem;
    transition: all 0.3s ease;
}

.back-btn img {
    width: 20px;
}

.back-btn:hover {
    opacity: 0.7;
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

.car-details-content {}

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
    margin-bottom: 1rem;
    cursor: pointer;
    transition: transform 0.3s ease;

    &:hover {
        transform: scale(1.01);

        .zoom-icon {
            opacity: 1;
        }
    }
}

.main-image img {
    width: 100%;
    height: 100%;
    object-fit: contain;
}

.zoom-icon {
    position: absolute;
    top: 20px;
    right: 20px;
    background: rgba(0, 0, 0, 0.6);
    backdrop-filter: blur(10px);
    color: white;
    width: 50px;
    height: 50px;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    opacity: 0;
    transition: opacity 0.3s ease;
    pointer-events: none;
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
    width: 70px;
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
                font-size: 2rem;
                font-weight: 600;
                line-height: 1.3;
            }
        }
    }
}

.car-specs h2 {
    font-size: 1.5rem;
}

.specs-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 1.5rem;
}

.spec-item {
    padding: 15px;
    background: #f8f9fa;
    border-radius: 10px;
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

/* Lightbox Styles */
.lightbox {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.95);
    z-index: 9999;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2rem;
}

.lightbox-close {
    position: absolute;
    top: 20px;
    right: 20px;
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    border: 2px solid rgba(255, 255, 255, 0.2);
    color: white;
    width: 45px;
    height: 45px;
    border-radius: 50%;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s ease;
    z-index: 10001;

    &:hover {
        background: rgba(255, 255, 255, 0.2);
        transform: rotate(90deg);
    }
}

.lightbox-content {
    position: relative;
    max-width: 90vw;
    max-height: 90vh;
    display: flex;
    align-items: center;
    justify-content: center;
    user-select: none;
    cursor: grab;

    &:active {
        cursor: grabbing;
    }
}

.lightbox-image-wrapper {
    transition: transform 0.3s ease;

    img {
        max-width: 90vw;
        max-height: 90vh;
        object-fit: contain;
        border-radius: 10px;
        box-shadow: 0 20px 60px rgba(0, 0, 0, 0.5);
        pointer-events: none;
    }
}

.lightbox-nav {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    border: 2px solid rgba(255, 255, 255, 0.2);
    color: white;
    width: 45px;
    height: 45px;
    border-radius: 50%;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s ease;
    z-index: 10001;

    &:hover {
        background: rgba(255, 255, 255, 0.2);
        transform: translateY(-50%) scale(1.1);
    }
}

.lightbox-nav-left {
    left: 20px;
}

.lightbox-nav-right {
    right: 20px;
}

.lightbox-counter {
    position: absolute;
    bottom: -60px;
    left: 50%;
    transform: translateX(-50%);
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    border: 2px solid rgba(255, 255, 255, 0.2);
    color: white;
    padding: 0.75rem 1.5rem;
    border-radius: 30px;
    font-size: 1rem;
    font-weight: 600;
}

/* Lightbox Transitions */
.lightbox-enter-active,
.lightbox-leave-active {
    transition: all 0.3s ease;
}

.lightbox-enter-from,
.lightbox-leave-to {
    opacity: 0;
}

.lightbox-enter-from .lightbox-content,
.lightbox-leave-to .lightbox-content {
    transform: scale(0.9);
}

.lightbox-enter-active .lightbox-content,
.lightbox-leave-active .lightbox-content {
    transition: transform 0.3s ease;
}

@media (max-width: 768px) {
    .car-details-page {
        padding: 1rem;
    }

    .car-details-content {}

    .car-header {
        flex-direction: column;
        align-items: flex-start;
        gap: 1rem;
    }

    .car-header h1 {
        font-size: 1.5rem;
    }

    .main-image {
        height: 400px;
    }

    .specs-grid {
        grid-template-columns: 1fr 1fr;
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

    .lightbox {
        padding: 1rem;
    }

    .lightbox-close {
        top: 10px;
        right: 10px;
        width: 40px;
        height: 40px;
    }

    .lightbox-nav {
        width: 30px;
        height: 30px;
    }

    .lightbox-nav-left {
        left: 10px;
    }

    .lightbox-nav-right {
        right: 10px;
    }

    .lightbox-counter {
        bottom: -50px;
        padding: 0.5rem 1rem;
        font-size: 0.875rem;
    }
}
</style>