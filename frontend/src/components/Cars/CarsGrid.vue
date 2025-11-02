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
                <!-- Custom Select for Fuel -->
                <div class="filter-group">
                    <div class="custom-select" :class="{ open: openSelect === 'fuel' }">
                        <div class="select-trigger" @click="toggleSelect('fuel')">
                            <span>{{ filters.fuel || 'Всі види палива' }}</span>
                            <svg class="arrow" width="12" height="8" viewBox="0 0 12 8">
                                <path d="M1 1L6 6L11 1" stroke="currentColor" stroke-width="2" fill="none" />
                            </svg>
                        </div>
                        <div class="select-options">
                            <div class="select-option" @click="selectOption('fuel', '')">
                                Всі види палива
                            </div>
                            <div class="select-option" @click="selectOption('fuel', 'Бензин')">
                                Бензин
                            </div>
                            <div class="select-option" @click="selectOption('fuel', 'Дизель')">
                                Дизель
                            </div>
                            <div class="select-option" @click="selectOption('fuel', 'Гібрид')">
                                Гібрид
                            </div>
                            <div class="select-option" @click="selectOption('fuel', 'Електро')">
                                Електро
                            </div>
                            <div class="select-option" @click="selectOption('fuel', 'Газ')">
                                Газ
                            </div>
                            <div class="select-option" @click="selectOption('fuel', 'Газ/Бензин')">
                                Газ/Бензин
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Custom Select for Transmission -->
                <div class="filter-group">
                    <div class="custom-select" :class="{ open: openSelect === 'transmission' }">
                        <div class="select-trigger" @click="toggleSelect('transmission')">
                            <span>{{ filters.transmission || 'Всі трансмісії' }}</span>
                            <svg class="arrow" width="12" height="8" viewBox="0 0 12 8">
                                <path d="M1 1L6 6L11 1" stroke="currentColor" stroke-width="2" fill="none" />
                            </svg>
                        </div>
                        <div class="select-options">
                            <div class="select-option" @click="selectOption('transmission', '')">
                                Всі трансмісії
                            </div>
                            <div class="select-option" @click="selectOption('transmission', 'Механічна')">
                                Механічна
                            </div>
                            <div class="select-option" @click="selectOption('transmission', 'Автоматична')">
                                Автоматична
                            </div>
                            <div class="select-option" @click="selectOption('transmission', 'Робот')">
                                Робот
                            </div>
                            <div class="select-option" @click="selectOption('transmission', 'Варіатор')">
                                Варіатор
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Custom Select for Drive -->
                <div class="filter-group">
                    <div class="custom-select" :class="{ open: openSelect === 'drive' }">
                        <div class="select-trigger" @click="toggleSelect('drive')">
                            <span>{{ filters.drive || 'Всі типи приводу' }}</span>
                            <svg class="arrow" width="12" height="8" viewBox="0 0 12 8">
                                <path d="M1 1L6 6L11 1" stroke="currentColor" stroke-width="2" fill="none" />
                            </svg>
                        </div>
                        <div class="select-options">
                            <div class="select-option" @click="selectOption('drive', '')">
                                Всі типи приводу
                            </div>
                            <div class="select-option" @click="selectOption('drive', 'Передній')">
                                Передній
                            </div>
                            <div class="select-option" @click="selectOption('drive', 'Задній')">
                                Задній
                            </div>
                            <div class="select-option" @click="selectOption('drive', 'Повний')">
                                Повний
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Search Input -->
                <div class="filter-group">
                    <input v-model="filters.search" type="text" placeholder="Пошук по марці або моделі..."
                        @input="debouncedSearch" class="search-input">
                </div>

                <!-- Custom Select for Items Per Page -->
                <div class="filter-group">
                    <div class="custom-select" :class="{ open: openSelect === 'itemsPerPage' }">
                        <div class="select-trigger" @click="toggleSelect('itemsPerPage')">
                            <span>{{ itemsPerPage }} авто</span>
                            <svg class="arrow" width="12" height="8" viewBox="0 0 12 8">
                                <path d="M1 1L6 6L11 1" stroke="currentColor" stroke-width="2" fill="none" />
                            </svg>
                        </div>
                        <div class="select-options">
                            <div class="select-option" @click="selectItemsPerPage(6)">
                                6 авто
                            </div>
                            <div class="select-option" @click="selectItemsPerPage(12)">
                                12 авто
                            </div>
                            <div class="select-option" @click="selectItemsPerPage(24)">
                                24 авто
                            </div>
                            <div class="select-option" @click="selectItemsPerPage(48)">
                                48 авто
                            </div>
                        </div>
                    </div>
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
                <CarCard v-for="car in filteredCars" :key="car.id" :car="car" @open-details="goToCarPage" />
            </div>

            <!-- Empty State -->
            <div v-if="!loading && !error && filteredCars.length === 0" class="empty-state">
                <p>Автомобілі не знайдені</p>
                <button @click="clearFilters" class="clear-filters">Показати всі</button>
            </div>

            <!-- Pagination -->
            <div v-if="totalPages > 1" class="pagination">
                <button @click="changePage(currentPage - 1)" :disabled="currentPage <= 1" class="page-btn">
                    Попередня
                </button>

                <div class="page-numbers">
                    <button v-for="page in visiblePages" :key="page" @click="changePage(page)"
                        :class="['page-btn', { active: page === currentPage }]">
                        {{ page }}
                    </button>
                </div>

                <button @click="changePage(currentPage + 1)" :disabled="currentPage >= totalPages" class="page-btn">
                    Наступна
                </button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import CarCard from '@/components/Cars/CarCard.vue'

const router = useRouter()

// Reactive data
const cars = ref([])
const loading = ref(false)
const error = ref('')
const currentPage = ref(1)
const totalCars = ref(0)
const totalPages = ref(0)
const itemsPerPage = ref(12)
const openSelect = ref(null)

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
const toggleSelect = (selectName) => {
    openSelect.value = openSelect.value === selectName ? null : selectName
}

const selectOption = (filterName, value) => {
    filters[filterName] = value
    openSelect.value = null
    applyFilters()
}

const selectItemsPerPage = (value) => {
    itemsPerPage.value = value
    openSelect.value = null
    changeItemsPerPage()
}

const closeAllSelects = (e) => {
    if (!e.target.closest('.custom-select')) {
        openSelect.value = null
    }
}

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

const goToCarPage = (car) => {
    router.push({ name: 'CarPage', params: { id: car.id } })
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

// Lifecycle
onMounted(() => {
    fetchCars()
    document.addEventListener('click', closeAllSelects)
})

onUnmounted(() => {
    document.removeEventListener('click', closeAllSelects)
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
    margin-bottom: 1rem;
    font-size: 2.5rem;
    color: #FFF;
}

.stats {
    color: #FFF;
    font-size: 1.1rem;
}

.filters {
    display: flex;
    gap: 1rem;
    margin-bottom: 2rem;
    flex-wrap: wrap;
    border-radius: 10px;
}

.filter-group {
    flex: 1;
    min-width: 180px;
    position: relative;
}

/* Custom Select Styles */
.custom-select {
    position: relative;
    width: 100%;
    user-select: none;
}

.select-trigger {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0.75rem 1rem;
    background: white;
    border: 2px solid #e0e0e0;
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.3s ease;
    font-size: 1rem;
    color: #333;
}

.select-trigger span {
    color: #333;
}

.select-trigger:hover {
    border-color: #007bff;
}

.custom-select.open .select-trigger {
    border-color: #007bff;
    border-bottom-left-radius: 0;
    border-bottom-right-radius: 0;
}

.select-trigger .arrow {
    transition: transform 0.3s ease;
    color: #666;
}

.custom-select.open .select-trigger .arrow {
    transform: rotate(180deg);
}

.select-options {
    position: absolute;
    top: 100%;
    left: 0;
    right: 0;
    background: white;
    border: 2px solid #007bff;
    border-top: none;
    border-bottom-left-radius: 8px;
    border-bottom-right-radius: 8px;
    max-height: 0;
    overflow: hidden;
    opacity: 0;
    transition: all 0.3s ease;
    z-index: 100;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.custom-select.open .select-options {
    max-height: 300px;
    overflow-y: auto;
    opacity: 1;
}

.select-option {
    padding: 0.75rem 1rem;
    cursor: pointer;
    transition: all 0.2s ease;
    color: #333;
}

.select-option:hover {
    background: #f0f8ff;
    color: #007bff;
}

.select-option:last-child {
    border-bottom-left-radius: 6px;
    border-bottom-right-radius: 6px;
}

/* Search Input */
.search-input {
    width: 100%;
    padding: 0.75rem 1rem;
    border: 2px solid #e0e0e0;
    border-radius: 8px;
    font-size: 1rem;
    color: #333;
    transition: all 0.3s ease;
}

.search-input:focus {
    outline: none;
    border-color: #007bff;
}

.search-input::placeholder {
    color: #999;
}

/* Clear Button */
.clear-filters {
    padding: 0.75rem 1.5rem;
    background: linear-gradient(135deg, #6c757d 0%, #545b62 100%);
    color: white;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    white-space: nowrap;
    transition: all 0.3s ease;
    font-weight: 600;
    box-shadow: 0 2px 8px rgba(108, 117, 125, 0.3);
}

.clear-filters:hover {
    background: linear-gradient(135deg, #545b62 0%, #3d4349 100%);
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(108, 117, 125, 0.4);
}

.clear-filters:active {
    transform: translateY(0);
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

/* Scrollbar for select options */
.select-options::-webkit-scrollbar {
    width: 6px;
}

.select-options::-webkit-scrollbar-track {
    background: #f1f1f1;
    border-radius: 3px;
}

.select-options::-webkit-scrollbar-thumb {
    background: #007bff;
    border-radius: 3px;
}

.select-options::-webkit-scrollbar-thumb:hover {
    background: #0056b3;
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
}
</style>