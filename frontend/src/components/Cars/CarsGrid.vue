<template>
    <div class="cars-catalog">
        <div class="container">
            <!-- Header -->
            <div class="catalog-header">
                <h1>Автомобілі в продажу</h1>
                <div class="stats">
                    <span class="total-count">Всього: {{ totalCars }}</span>
                </div>
            </div>

            <!-- Filters -->
            <div class="filters">
                <div class="filter-group">
                    <select v-model="filters.fuel" @change="applyFilters">
                        <option value="">Всі види палива</option>
                        <option value="Бензин">Бензин</option>
                        <option value="Дизель">Дизель</option>
                        <option value="Гібрид">Гібрид</option>
                        <option value="Електро">Електро</option>
                        <option value="Газ">Газ</option>
                        <option value="Газ/Бензин">Газ/Бензин</option>
                    </select>
                </div>

                <div class="filter-group">
                    <select v-model="filters.transmission" @change="applyFilters">
                        <option value="">Всі трансмісії</option>
                        <option value="Механічна">Механічна</option>
                        <option value="Автоматична">Автоматична</option>
                        <option value="Робот">Робот</option>
                        <option value="Варіатор">Варіатор</option>
                    </select>
                </div>

                <div class="filter-group">
                    <select v-model="filters.drive" @change="applyFilters">
                        <option value="">Всі типи приводу</option>
                        <option value="Передній">Передній</option>
                        <option value="Задній">Задній</option>
                        <option value="Повний">Повний</option>
                    </select>
                </div>

                <div class="filter-group">
                    <input v-model="filters.search" 
                           type="text" 
                           placeholder="Пошук по марці або моделі..."
                           @input="debouncedSearch">
                </div>

                <div class="filter-group">
                    <select v-model="itemsPerPage" @change="changeItemsPerPage">
                        <option :value="6">6 авто</option>
                        <option :value="12">12 авто</option>
                        <option :value="24">24 авто</option>
                        <option :value="48">48 авто</option>
                    </select>
                </div>

                <button @click="clearFilters" class="clear-filters">Очистити</button>
            </div>

            <!-- Loading -->
            <div v-if="loading" class="loading">
                <div class="loading-spinner"></div>
                <p>Завантажуємо автомобілі...</p>
            </div>

            <!-- Error -->
            <div v-if="error" class="error">
                <p>{{ error }}</p>
                <button @click="fetchCars" class="retry-btn">Спробувати знову</button>
            </div>

            <!-- Cars Grid -->
            <div v-if="!loading && !error" class="cars-grid">
                <CarCard 
                    v-for="car in filteredCars" 
                    :key="car.id" 
                    :car="car"
                    @open-details="openCarModal"
                />
            </div>

            <!-- Empty State -->
            <div v-if="!loading && !error && filteredCars.length === 0" class="empty-state">
                <p>Автомобілі не знайдені</p>
                <button @click="clearFilters" class="clear-filters">Показати всі</button>
            </div>

            <!-- Pagination -->
            <div v-if="totalPages > 1" class="pagination">
                <button @click="changePage(currentPage - 1)" 
                        :disabled="currentPage <= 1" 
                        class="page-btn">
                    Попередня
                </button>

                <div class="page-numbers">
                    <button v-for="page in visiblePages" 
                            :key="page" 
                            @click="changePage(page)"
                            :class="['page-btn', { active: page === currentPage }]">
                        {{ page }}
                    </button>
                </div>

                <button @click="changePage(currentPage + 1)" 
                        :disabled="currentPage >= totalPages" 
                        class="page-btn">
                    Наступна
                </button>
            </div>
        </div>

        <!-- Car Detail Modal -->
        <div v-if="selectedCar" class="car-modal" @click="closeCarModal">
            <div class="modal-content" @click.stop>
                <button class="close-btn" @click="closeCarModal">&times;</button>

                <div class="modal-header">
                    <h2>{{ selectedCar.brand }} {{ selectedCar.model }} {{ selectedCar.year }}</h2>
                </div>

                <div class="modal-body">
                    <!-- Images -->
                    <div v-if="selectedCarImages.length > 0" class="car-images">
                        <div class="main-image">
                            <img :src="selectedCarImages[currentImageIndex].fileUrl" 
                                 :alt="selectedCar.brand">
                        </div>
                        <div v-if="selectedCarImages.length > 1" class="image-thumbnails">
                            <img v-for="(image, index) in selectedCarImages" 
                                 :key="image.id" 
                                 :src="image.fileUrl"
                                 :alt="`Фото ${index + 1}`" 
                                 :class="{ active: index === currentImageIndex }"
                                 @click="currentImageIndex = index">
                        </div>
                    </div>

                    <!-- Car Details -->
                    <div class="car-full-details">
                        <div class="details-grid">
                            <div class="detail-item">
                                <span class="label">Марка:</span>
                                <span class="value">{{ selectedCar.brand }}</span>
                            </div>
                            <div class="detail-item">
                                <span class="label">Модель:</span>
                                <span class="value">{{ selectedCar.model }}</span>
                            </div>
                            <div class="detail-item">
                                <span class="label">Рік випуску:</span>
                                <span class="value">{{ selectedCar.year }}</span>
                            </div>
                            <div v-if="selectedCar.color" class="detail-item">
                                <span class="label">Колір:</span>
                                <span class="value">{{ selectedCar.color }}</span>
                            </div>
                            <div class="detail-item">
                                <span class="label">Пробіг:</span>
                                <span class="value">{{ formatMileage(selectedCar.mileage) }}</span>
                            </div>
                            <div class="detail-item">
                                <span class="label">Паливо:</span>
                                <span class="value">{{ selectedCar.fuel }}</span>
                            </div>
                            <div class="detail-item">
                                <span class="label">Трансмісія:</span>
                                <span class="value">{{ selectedCar.transmission }}</span>
                            </div>
                            <div class="detail-item">
                                <span class="label">Привід:</span>
                                <span class="value">{{ selectedCar.drive }}</span>
                            </div>
                            <div class="detail-item">
                                <span class="label">Дата додавання:</span>
                                <span class="value">{{ formatDate(selectedCar.createdAt) }}</span>
                            </div>
                        </div>

                        <div v-if="selectedCar.description" class="full-description">
                            <h4>Опис:</h4>
                            <p>{{ selectedCar.description }}</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import CarCard from '../Cars/CarCard.vue'

// Reactive data
const cars = ref([])
const loading = ref(false)
const error = ref('')
const currentPage = ref(1)
const totalCars = ref(0)
const totalPages = ref(0)
const itemsPerPage = ref(12)
const selectedCar = ref(null)
const selectedCarImages = ref([])
const currentImageIndex = ref(0)

const filters = reactive({
    fuel: '',
    transmission: '',
    drive: '',
    search: ''
})

let searchTimeout = null

// Computed
const filteredCars = computed(() => {
    let result = cars.value

    if (filters.fuel) {
        result = result.filter(car => car.fuel === filters.fuel)
    }

    if (filters.transmission) {
        result = result.filter(car => car.transmission === filters.transmission)
    }

    if (filters.drive) {
        result = result.filter(car => car.drive === filters.drive)
    }

    if (filters.search) {
        const search = filters.search.toLowerCase()
        result = result.filter(car =>
            car.brand.toLowerCase().includes(search) ||
            car.model.toLowerCase().includes(search)
        )
    }

    return result
})

const visiblePages = computed(() => {
    const pages = []
    const start = Math.max(1, currentPage.value - 2)
    const end = Math.min(totalPages.value, currentPage.value + 2)

    for (let i = start; i <= end; i++) {
        pages.push(i)
    }

    return pages
})

// Methods
const fetchCars = async () => {
    loading.value = true
    error.value = ''

    try {
        const params = new URLSearchParams({
            page: currentPage.value,
            limit: itemsPerPage.value
        })

        const apiUrl = process.env.NODE_ENV === 'production'
            ? `/api/cars-sale?${params}`
            : `http://localhost:8001/api/cars-sale?${params}`

        const response = await fetch(apiUrl)

        if (!response.ok) {
            throw new Error(`HTTP ${response.status}`)
        }

        const data = await response.json()

        if (data.status === 'success') {
            cars.value = data.data || []
            totalCars.value = data.pagination?.total || 0
            totalPages.value = data.pagination?.pages || 0
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

const fetchCarDetails = async (carId) => {
    try {
        const apiUrl = process.env.NODE_ENV === 'production'
            ? `/api/cars-sale/${carId}`
            : `http://localhost:8001/api/cars-sale/${carId}`

        const response = await fetch(apiUrl)

        if (response.ok) {
            const data = await response.json()
            if (data.status === 'success') {
                selectedCarImages.value = data.data.images || []
                currentImageIndex.value = 0
            }
        }
    } catch (err) {
        console.error('Error fetching car details:', err)
    }
}

const openCarModal = async (car) => {
    selectedCar.value = car
    await fetchCarDetails(car.id)
}

const closeCarModal = () => {
    selectedCar.value = null
    selectedCarImages.value = []
    currentImageIndex.value = 0
}

const changePage = (newPage) => {
    if (newPage >= 1 && newPage <= totalPages.value) {
        currentPage.value = newPage
        fetchCars()
        window.scrollTo({ top: 0, behavior: 'smooth' })
    }
}

const changeItemsPerPage = () => {
    currentPage.value = 1
    fetchCars()
}

const applyFilters = () => {
    // Фільтри застосовуються автоматично через computed filteredCars
}

const debouncedSearch = () => {
    clearTimeout(searchTimeout)
    searchTimeout = setTimeout(() => {
        applyFilters()
    }, 500)
}

const clearFilters = () => {
    filters.fuel = ''
    filters.transmission = ''
    filters.drive = ''
    filters.search = ''
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



// Lifecycle
onMounted(() => {
    fetchCars()
})
</script>

<style scoped>
.cars-catalog {
    padding: 2rem;
    min-height: 100vh;
}

.container {
    max-width: 1440px;
    margin: 0 auto;
}

.catalog-header {
    text-align: center;
    margin-bottom: 2rem;
}

.catalog-header h1 {
    color: #333;
    margin-bottom: 1rem;
    font-size: 2.5rem;
}

.stats {
    color: #666;
    font-size: 1.1rem;
}

.filters {
    display: flex;
    gap: 1rem;
    margin-bottom: 2rem;
    flex-wrap: wrap;
    background: white;
    padding: 1rem;
    border-radius: 10px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.filter-group {
    flex: 1;
    min-width: 180px;
}

.filter-group select,
.filter-group input {
    width: 100%;
    padding: 0.75rem;
    border: 1px solid #ddd;
    border-radius: 5px;
    font-size: 1rem;
}

.clear-filters {
    padding: 0.75rem 1.5rem;
    background: #6c757d;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    white-space: nowrap;
    transition: background 0.3s ease;
}

.clear-filters:hover {
    background: #545b62;
}

.loading,
.error {
    text-align: center;
    padding: 2rem;
    background: white;
    border-radius: 10px;
    margin: 2rem 0;
}

.loading-spinner {
    width: 40px;
    height: 40px;
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
    padding: 0.5rem 1rem;
    border-radius: 5px;
    cursor: pointer;
}

.retry-btn:hover {
    background: #0056b3;
}

.cars-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
    gap: 2rem;
    margin: 2rem 0;
}

.empty-state {
    text-align: center;
    padding: 3rem;
    background: white;
    border-radius: 10px;
    color: #666;
}

.pagination {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 1rem;
    margin: 2rem 0;
}

.page-numbers {
    display: flex;
    gap: 0.5rem;
}

.page-btn {
    padding: 0.5rem 1rem;
    border: 1px solid #ddd;
    background: white;
    color: #007bff;
    border-radius: 5px;
    cursor: pointer;
    min-width: 40px;
    transition: all 0.3s ease;
}

.page-btn:hover:not(:disabled) {
    background: #007bff;
    color: white;
}

.page-btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}

.page-btn.active {
    background: #007bff;
    color: white;
}

.car-modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.8);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
    padding: 2rem;
}

.modal-content {
    background: white;
    border-radius: 15px;
    width: 100%;
    max-width: 800px;
    max-height: 90vh;
    overflow-y: auto;
    position: relative;
}

.close-btn {
    position: absolute;
    top: 1rem;
    right: 1rem;
    background: none;
    border: none;
    font-size: 2rem;
    cursor: pointer;
    color: #666;
    z-index: 1001;
    width: 40px;
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 50%;
    transition: all 0.3s ease;
}

.close-btn:hover {
    background: #f0f0f0;
    color: #333;
}

.modal-header {
    padding: 2rem 2rem 1rem;
    border-bottom: 1px solid #eee;
}

.modal-header h2 {
    margin: 0;
    color: #333;
}

.modal-body {
    padding: 2rem;
}

.car-images {
    margin-bottom: 2rem;
}

.main-image {
    margin-bottom: 1rem;
}

.main-image img {
    width: 100%;
    max-height: 400px;
    object-fit: contain;
    border-radius: 10px;
}

.image-thumbnails {
    display: flex;
    gap: 0.5rem;
    overflow-x: auto;
    padding: 0.5rem 0;
}

.image-thumbnails img {
    width: 80px;
    height: 60px;
    object-fit: cover;
    border-radius: 5px;
    cursor: pointer;
    opacity: 0.7;
    border: 2px solid transparent;
    transition: all 0.3s ease;
}

.image-thumbnails img:hover,
.image-thumbnails img.active {
    opacity: 1;
    border-color: #007bff;
}

.details-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 1rem;
    margin-bottom: 2rem;
}

.detail-item {
    padding: 1rem;
    background: #f8f9fa;
    border-radius: 8px;
}

.detail-item .label {
    display: block;
    color: #666;
    font-size: 0.875rem;
    margin-bottom: 0.25rem;
}

.detail-item .value {
    display: block;
    color: #333;
    font-weight: 600;
    font-size: 1rem;
}

.full-description {
    padding: 1rem;
    background: #f8f9fa;
    border-radius: 8px;
}

.full-description h4 {
    margin: 0 0 0.5rem 0;
    color: #333;
}

.full-description p {
    margin: 0;
    color: #666;
    line-height: 1.5;
    white-space: pre-wrap;
}

@media (max-width: 768px) {
    .cars-catalog {
        padding: 1rem;
    }

    .cars-grid {
        grid-template-columns: 1fr;
    }

    .filters {
        flex-direction: column;
    }

    .filter-group {
        min-width: auto;
    }

    .car-modal {
        padding: 1rem;
    }

    .modal-body {
        padding: 1rem;
    }

    .details-grid {
        grid-template-columns: 1fr;
    }
}
</style>