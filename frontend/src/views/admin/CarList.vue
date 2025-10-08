<template>
    <div class="cars-list-container">
        <div class="header">
            <h1>Автомобілі на продаж</h1>
            <button @click="goToAddCar" class="btn-add">
                + Додати автомобіль
            </button>
        </div>

        <div v-if="loading" class="loading">Завантаження...</div>

        <div v-else-if="cars.length === 0" class="empty-state">
            <p>Немає автомобілів в каталозі</p>
            <button @click="goToAddCar" class="btn-primary">Додати перший автомобіль</button>
        </div>

        <div v-else class="cars-grid">
            <div v-for="car in cars" :key="car.id" class="car-card">
                <div class="car-image">
                    <img v-if="car.imageCount > 0" 
                         :src="`/uploads/car_sale_${car.id}_image_1_*.jpg`" 
                         :alt="`${car.brand} ${car.model}`"
                         @error="handleImageError">
                    <div v-else class="no-image">Немає фото</div>
                    
                    <!-- Бейдж статусу -->
                    <div v-if="car.status && car.status !== 'active'" class="status-badge" :class="`status-${car.status}`">
                        {{ getStatusLabel(car.status) }}
                    </div>
                </div>

                <div class="car-info">
                    <h3>{{ car.brand }} {{ car.model }}</h3>
                    
                    <div class="car-details">
                        <div class="detail-row">
                            <span class="label">Рік:</span>
                            <span class="value">{{ car.year }}</span>
                        </div>
                        <div v-if="car.color" class="detail-row">
                            <span class="label">Колір:</span>
                            <span class="value">{{ car.color }}</span>
                        </div>
                        <div class="detail-row">
                            <span class="label">Пробіг:</span>
                            <span class="value">{{ formatMileage(car.mileage) }} км</span>
                        </div>
                        <div class="detail-row">
                            <span class="label">Паливо:</span>
                            <span class="value">{{ car.fuel }}</span>
                        </div>
                        <div class="detail-row">
                            <span class="label">КПП:</span>
                            <span class="value">{{ car.transmission }}</span>
                        </div>
                        <div class="detail-row">
                            <span class="label">Привід:</span>
                            <span class="value">{{ car.drive }}</span>
                        </div>
                    </div>

                    <p v-if="car.description" class="car-description">
                        {{ truncateText(car.description, 100) }}
                    </p>

                    <div class="car-meta">
                        <span class="photo-count">📷 {{ car.imageCount }} фото</span>
                        <span class="created-date">{{ formatDate(car.createdAt) }}</span>
                    </div>

                    <div class="car-actions">
                        <button @click="viewCar(car.id)" class="btn-view">
                            Детальніше
                        </button>
                        <button @click="editCar(car.id)" class="btn-edit">
                            ✏️ Редагувати
                        </button>
                        <button @click="confirmDelete(car)" class="btn-delete">
                            🗑️ Видалити
                        </button>
                    </div>
                </div>
            </div>
        </div>

        <!-- Пагінація -->
        <div v-if="pagination.pages > 1" class="pagination">
            <button @click="changePage(pagination.page - 1)" 
                    :disabled="pagination.page === 1"
                    class="btn-page">
                ← Назад
            </button>
            
            <span class="page-info">
                Сторінка {{ pagination.page }} з {{ pagination.pages }}
            </span>
            
            <button @click="changePage(pagination.page + 1)" 
                    :disabled="pagination.page === pagination.pages"
                    class="btn-page">
                Вперед →
            </button>
        </div>

        <!-- Модальне вікно підтвердження видалення -->
        <div v-if="deleteModal.show" class="modal-overlay" @click="closeDeleteModal">
            <div class="modal" @click.stop>
                <h3>Видалити автомобіль?</h3>
                <p>
                    Ви впевнені, що хочете видалити 
                    <strong>{{ deleteModal.car?.brand }} {{ deleteModal.car?.model }}</strong>?
                </p>
                <p class="warning">Цю дію неможливо скасувати!</p>
                
                <div class="modal-actions">
                    <button @click="closeDeleteModal" class="btn-secondary" :disabled="deleteModal.deleting">
                        Скасувати
                    </button>
                    <button @click="deleteCar" class="btn-danger" :disabled="deleteModal.deleting">
                        {{ deleteModal.deleting ? 'Видалення...' : 'Видалити' }}
                    </button>
                </div>

                <div v-if="deleteModal.error" class="error-message">
                    {{ deleteModal.error }}
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const cars = ref([])
const loading = ref(true)
const pagination = reactive({
    page: 1,
    limit: 12,
    total: 0,
    pages: 0
})

const deleteModal = reactive({
    show: false,
    car: null,
    deleting: false,
    error: ''
})

// Завантаження списку автомобілів
const loadCars = async (page = 1) => {
    loading.value = true
    
    try {
        const apiUrl = process.env.NODE_ENV === 'production'
            ? '/api/cars-sale?page='
            : 'http://localhost:8001/api/cars-sale?page='
        
        const response = await fetch(`${apiUrl}${page}&limit=${pagination.limit}`)
        
        if (!response.ok) {
            throw new Error('Помилка завантаження даних')
        }

        const result = await response.json()
        cars.value = result.data || []
        
        if (result.pagination) {
            Object.assign(pagination, result.pagination)
        }

    } catch (error) {
        console.error('Error loading cars:', error)
    } finally {
        loading.value = false
    }
}

// Форматування пробігу
const formatMileage = (mileage) => {
    return mileage.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ' ')
}

// Форматування дати
const formatDate = (dateString) => {
    const date = new Date(dateString)
    return date.toLocaleDateString('uk-UA', { 
        year: 'numeric', 
        month: 'long', 
        day: 'numeric' 
    })
}

// Обрізання тексту
const truncateText = (text, maxLength) => {
    if (!text) return ''
    if (text.length <= maxLength) return text
    return text.substring(0, maxLength) + '...'
}

// Отримання назви статусу
const getStatusLabel = (status) => {
    const labels = {
        'sold': 'Продано',
        'new': 'Новинка',
        'sale': 'Знижка',
        'super-price': 'Супер ціна'
    }
    return labels[status] || status
}

// Обробка помилки зображення
const handleImageError = (event) => {
    event.target.src = '/placeholder-car.jpg'
}

// Навігація
const goToAddCar = () => {
    router.push('/admin/add')
}

const viewCar = (id) => {
    router.push(`/admin/detail/${id}`)
}

const editCar = (id) => {
    router.push(`/admin/edit/${id}`)
}

const changePage = (page) => {
    if (page >= 1 && page <= pagination.pages) {
        loadCars(page)
        window.scrollTo({ top: 0, behavior: 'smooth' })
    }
}

// Видалення автомобіля
const confirmDelete = (car) => {
    deleteModal.show = true
    deleteModal.car = car
    deleteModal.error = ''
}

const closeDeleteModal = () => {
    if (!deleteModal.deleting) {
        deleteModal.show = false
        deleteModal.car = null
        deleteModal.error = ''
    }
}

const deleteCar = async () => {
    if (!deleteModal.car) return

    deleteModal.deleting = true
    deleteModal.error = ''

    try {
        const apiUrl = process.env.NODE_ENV === 'production'
            ? '/api/cars-sale'
            : 'http://localhost:8001/api/cars-sale'

        const response = await fetch(`${apiUrl}${deleteModal.car.id}`, {
            method: 'DELETE',
            headers: {
                'Authorization': `Bearer ${authStore.token}`
            }
        })

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({ error: 'Unknown error' }))
            throw new Error(errorData.error || 'Помилка видалення')
        }

        // Видаляємо зі списку
        cars.value = cars.value.filter(c => c.id !== deleteModal.car.id)
        
        // Закриваємо модальне вікно
        closeDeleteModal()

        // Перезавантажуємо список якщо сторінка спорожніла
        if (cars.value.length === 0 && pagination.page > 1) {
            loadCars(pagination.page - 1)
        }

    } catch (error) {
        console.error('Delete error:', error)
        deleteModal.error = error.message || 'Помилка при видаленні автомобіля'
    } finally {
        deleteModal.deleting = false
    }
}

// Завантаження при монтуванні
onMounted(() => {
    loadCars()
})
</script>

<style scoped>
.cars-list-container {
    max-width: 1400px;
    margin: 0 auto;
    padding: 2rem;
}

.header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
}

.header h1 {
    margin: 0;
    color: #333;
}

.btn-add {
    padding: 0.75rem 1.5rem;
    background: #28a745;
    color: white;
    border: none;
    border-radius: 4px;
    font-size: 1rem;
    cursor: pointer;
    transition: background 0.3s ease;
}

.btn-add:hover {
    background: #218838;
}

.loading {
    text-align: center;
    padding: 3rem;
    font-size: 1.2rem;
    color: #666;
}

.empty-state {
    text-align: center;
    padding: 3rem;
}

.empty-state p {
    font-size: 1.2rem;
    color: #666;
    margin-bottom: 1.5rem;
}

.cars-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
    gap: 1.5rem;
    margin-bottom: 2rem;
}

.car-card {
    background: white;
    border-radius: 8px;
    overflow: hidden;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.car-card:hover {
    transform: translateY(-4px);
    box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
}

.car-image {
    position: relative;
    width: 100%;
    height: 220px;
    background: #f5f5f5;
    overflow: hidden;
}

.car-image img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.no-image {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
    color: #999;
    font-size: 1.1rem;
}

.status-badge {
    position: absolute;
    top: 10px;
    right: 10px;
    padding: 0.4rem 0.8rem;
    border-radius: 4px;
    font-size: 0.85rem;
    font-weight: 600;
    color: white;
    text-transform: uppercase;
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

.car-info {
    padding: 1.25rem;
}

.car-info h3 {
    margin: 0 0 1rem 0;
    color: #333;
    font-size: 1.3rem;
}

.car-details {
    margin-bottom: 1rem;
}

.detail-row {
    display: flex;
    justify-content: space-between;
    padding: 0.4rem 0;
    font-size: 0.9rem;
    border-bottom: 1px solid #f0f0f0;
}

.detail-row:last-child {
    border-bottom: none;
}

.detail-row .label {
    color: #666;
    font-weight: 500;
}

.detail-row .value {
    color: #333;
}

.car-description {
    margin: 1rem 0;
    color: #666;
    font-size: 0.9rem;
    line-height: 1.5;
}

.car-meta {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin: 1rem 0;
    padding-top: 1rem;
    border-top: 1px solid #eee;
    font-size: 0.85rem;
    color: #999;
}

.car-actions {
    display: flex;
    gap: 0.5rem;
    margin-top: 1rem;
}

.btn-view,
.btn-edit,
.btn-delete {
    flex: 1;
    padding: 0.6rem;
    border: none;
    border-radius: 4px;
    font-size: 0.9rem;
    cursor: pointer;
    transition: all 0.3s ease;
}

.btn-view {
    background: #007bff;
    color: white;
}

.btn-view:hover {
    background: #0056b3;
}

.btn-edit {
    background: #ffc107;
    color: #333;
}

.btn-edit:hover {
    background: #e0a800;
}

.btn-delete {
    background: #dc3545;
    color: white;
}

.btn-delete:hover {
    background: #c82333;
}

.pagination {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 1rem;
    margin-top: 2rem;
}

.btn-page {
    padding: 0.6rem 1.2rem;
    background: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background 0.3s ease;
}

.btn-page:hover:not(:disabled) {
    background: #0056b3;
}

.btn-page:disabled {
    background: #ccc;
    cursor: not-allowed;
}

.page-info {
    color: #666;
    font-weight: 500;
}

/* Модальне вікно */
.modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
}

.modal {
    background: white;
    padding: 2rem;
    border-radius: 8px;
    max-width: 500px;
    width: 90%;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
}

.modal h3 {
    margin: 0 0 1rem 0;
    color: #333;
}

.modal p {
    color: #666;
    line-height: 1.6;
    margin: 0.5rem 0;
}

.modal .warning {
    color: #dc3545;
    font-weight: 500;
    margin-top: 1rem;
}

.modal-actions {
    display: flex;
    gap: 1rem;
    margin-top: 1.5rem;
}

.btn-primary,
.btn-secondary,
.btn-danger {
    flex: 1;
    padding: 0.75rem;
    border: none;
    border-radius: 4px;
    font-size: 1rem;
    cursor: pointer;
    transition: all 0.3s ease;
}

.btn-primary {
    background: #007bff;
    color: white;
}

.btn-primary:hover:not(:disabled) {
    background: #0056b3;
}

.btn-secondary {
    background: #6c757d;
    color: white;
}

.btn-secondary:hover:not(:disabled) {
    background: #545b62;
}

.btn-danger {
    background: #dc3545;
    color: white;
}

.btn-danger:hover:not(:disabled) {
    background: #c82333;
}

.btn-primary:disabled,
.btn-secondary:disabled,
.btn-danger:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}

.error-message {
    background: #f8d7da;
    color: #721c24;
    padding: 0.75rem;
    border-radius: 4px;
    border: 1px solid #f5c6cb;
    margin-top: 1rem;
    font-size: 0.9rem;
}

@media (max-width: 768px) {
    .cars-list-container {
        padding: 1rem;
    }

    .header {
        flex-direction: column;
        gap: 1rem;
        align-items: stretch;
    }

    .btn-add {
        width: 100%;
    }

    .cars-grid {
        grid-template-columns: 1fr;
    }

    .car-actions {
        flex-direction: column;
    }

    .btn-view,
    .btn-edit,
    .btn-delete {
        width: 100%;
    }

    .pagination {
        flex-direction: column;
        gap: 0.5rem;
    }

    .modal {
        padding: 1.5rem;
    }

    .modal-actions {
        flex-direction: column;
    }
}
</style>