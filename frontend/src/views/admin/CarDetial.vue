<template>
    <div class="car-detail-container">
        <div v-if="loading" class="loading">Загрузка...</div>

        <div v-else-if="error" class="error-state">
            <p>{{ error }}</p>
            <button @click="goBack" class="btn-primary">Вернуться назад</button>
        </div>

        <div v-else-if="car" class="car-detail">
            <!-- Навигация -->
            <div class="navigation">
                <button @click="goBack" class="btn-back">← Назад к списку</button>
                <div class="nav-actions">
                    <button @click="editCar" class="btn-edit">✏️ Редактировать</button>
                    <button @click="confirmDelete" class="btn-delete">🗑️ Удалить</button>
                </div>
            </div>

            <!-- Заголовок -->
            <div class="car-header">
                <h1>{{ car.brand }} {{ car.model }}</h1>
                <div v-if="car.status && car.status !== 'active'" class="status-badge" :class="`status-${car.status}`">
                    {{ getStatusLabel(car.status) }}
                </div>
            </div>

            <!-- Галерея изображений -->
            <div class="image-gallery">
                <div v-if="car.images && car.images.length > 0" class="gallery-main">
                    <img :src="currentImage.fileUrl" :alt="`${car.brand} ${car.model}`" class="main-image">

                    <div v-if="car.images.length > 1" class="gallery-controls">
                        <button @click="prevImage" class="gallery-btn">‹</button>
                        <span class="image-counter">{{ currentImageIndex + 1 }} / {{ car.images.length }}</span>
                        <button @click="nextImage" class="gallery-btn">›</button>
                    </div>
                </div>

                <div v-else class="no-images">
                    <p>Нет фотографий</p>
                </div>

                <!-- Миниатюры -->
                <div v-if="car.images && car.images.length > 1" class="gallery-thumbnails">
                    <div v-for="(image, index) in car.images" :key="image.id" class="thumbnail"
                        :class="{ active: index === currentImageIndex }" @click="currentImageIndex = index">
                        <img :src="image.fileUrl" :alt="`Фото ${index + 1}`">
                    </div>
                </div>
            </div>

            <!-- Основная информация -->
            <div class="car-info-grid">
                <div class="info-section">
                    <h2>Технические характеристики</h2>
                    <div class="info-table">
                        <div class="info-row">
                            <span class="info-label">Год выпуска:</span>
                            <span class="info-value">{{ car.year }}</span>
                        </div>
                        <div v-if="car.color" class="info-row">
                            <span class="info-label">Цвет:</span>
                            <span class="info-value">{{ car.color }}</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">Пробег:</span>
                            <span class="info-value">{{ formatMileage(car.mileage) }} км</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">Тип топлива:</span>
                            <span class="info-value">{{ car.fuel }}</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">Трансмиссия:</span>
                            <span class="info-value">{{ car.transmission }}</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">Привод:</span>
                            <span class="info-value">{{ car.drive }}</span>
                        </div>
                    </div>
                </div>

                <div v-if="car.description" class="info-section">
                    <h2>Описание</h2>
                    <p class="description">{{ car.description }}</p>
                </div>

                <div class="info-section">
                    <h2>Дополнительная информация</h2>
                    <div class="info-table">
                        <div class="info-row">
                            <span class="info-label">Добавлено:</span>
                            <span class="info-value">{{ formatDate(car.createdAt) }}</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">Обновлено:</span>
                            <span class="info-value">{{ formatDate(car.updatedAt) }}</span>
                        </div>
                        <div class="info-row">
                            <span class="info-label">ID объявления:</span>
                            <span class="info-value">#{{ car.id }}</span>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Контактная информация (можно добавить) -->
            <div class="contact-section">
                <h2>Свяжитесь с нами</h2>
                <p>Для получения дополнительной информации или записи на тест-драйв свяжитесь с нашими менеджерами.</p>
                <button class="btn-contact">📞 Связаться</button>
            </div>
        </div>

        <!-- Модальное окно удаления -->
        <div v-if="deleteModal.show" class="modal-overlay" @click="closeDeleteModal">
            <div class="modal" @click.stop>
                <h3>Удалить автомобиль?</h3>
                <p>
                    Вы уверены, что хотите удалить
                    <strong>{{ car?.brand }} {{ car?.model }}</strong>?
                </p>
                <p class="warning">Это действие нельзя отменить!</p>

                <div class="modal-actions">
                    <button @click="closeDeleteModal" class="btn-secondary" :disabled="deleteModal.deleting">
                        Отмена
                    </button>
                    <button @click="deleteCar" class="btn-danger" :disabled="deleteModal.deleting">
                        {{ deleteModal.deleting ? 'Удаление...' : 'Удалить' }}
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
// Замініть ВЕСЬ <script setup> на цей:

import { ref, reactive, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const car = ref(null)
const loading = ref(true)
const error = ref('')
const currentImageIndex = ref(0)

const deleteModal = reactive({
    show: false,
    deleting: false,
    error: ''
})

const currentImage = computed(() => {
    if (car.value?.images && car.value.images.length > 0) {
        return car.value.images[currentImageIndex.value]
    }
    return null
})

// Загрузка данных автомобиля
const loadCarData = async () => {
    loading.value = true
    error.value = ''

    try {
        const carId = route.params.id
        const apiUrl = process.env.NODE_ENV === 'production'
            ? `/api/cars-sale/${carId}`
            : `http://localhost:8001/api/cars-sale/${carId}`

        const response = await fetch(apiUrl)

        if (!response.ok) {
            throw new Error('Автомобиль не найден')
        }

        const result = await response.json()
        car.value = result.data

    } catch (err) {
        console.error('Error loading car:', err)
        error.value = err.message || 'Ошибка загрузки данных'
    } finally {
        loading.value = false
    }
}

// Форматирование
const formatMileage = (mileage) => {
    return mileage.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ' ')
}

const formatDate = (dateString) => {
    const date = new Date(dateString)
    return date.toLocaleDateString('uk-UA', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
    })
}

const getStatusLabel = (status) => {
    const labels = {
        'sold': 'Продано',
        'new': 'Новинка',
        'sale': 'Знижка',
        'super-price': 'Супер ціна'
    }
    return labels[status] || status
}

// Навигация по изображениям
const nextImage = () => {
    if (car.value?.images && car.value.images.length > 1) {
        currentImageIndex.value = (currentImageIndex.value + 1) % car.value.images.length
    }
}

const prevImage = () => {
    if (car.value?.images && car.value.images.length > 1) {
        currentImageIndex.value = currentImageIndex.value === 0
            ? car.value.images.length - 1
            : currentImageIndex.value - 1
    }
}

// Навигация
const goBack = () => {
    router.push('/admin/list')
}

const editCar = () => {
    router.push(`/admin/edit/${route.params.id}`)
}

// Удаление
const confirmDelete = () => {
    deleteModal.show = true
    deleteModal.error = ''
}

const closeDeleteModal = () => {
    if (!deleteModal.deleting) {
        deleteModal.show = false
        deleteModal.error = ''
    }
}

const deleteCar = async () => {
    deleteModal.deleting = true
    deleteModal.error = ''

    try {
        const apiUrl = process.env.NODE_ENV === 'production'
            ? `/api/cars-sale/${route.params.id}`
            : `http://localhost:8001/api/cars-sale/${route.params.id}`

        const response = await fetch(apiUrl, {
            method: 'DELETE',
            headers: {
                'Authorization': `Bearer ${authStore.token}`
            }
        })

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({ error: 'Unknown error' }))
            throw new Error(errorData.error || 'Ошибка удаления')
        }

        // Успешно удалено, возвращаемся к списку
        router.push('/admin/list')

    } catch (err) {
        console.error('Delete error:', err)
        deleteModal.error = err.message || 'Ошибка при удалении автомобиля'
    } finally {
        deleteModal.deleting = false
    }
}

onMounted(() => {
    loadCarData()
})
</script>

<style scoped>
.car-detail-container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 2rem;
    padding-top: 150px;
}

.loading,
.error-state {
    text-align: center;
    padding: 3rem;
}

.loading {
    font-size: 1.2rem;
    color: #666;
}

.error-state p {
    font-size: 1.1rem;
    color: #dc3545;
    margin-bottom: 1.5rem;
}

.navigation {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
}

.nav-actions {
    display: flex;
    gap: 0.5rem;
}

.btn-back,
.btn-edit,
.btn-delete {
    padding: 0.6rem 1.2rem;
    border: none;
    border-radius: 4px;
    font-size: 0.95rem;
    cursor: pointer;
    transition: all 0.3s ease;
}

.btn-back {
    background: #6c757d;
    color: white;
}

.btn-back:hover {
    background: #545b62;
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

.car-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
}

.car-header h1 {
    margin: 0;
    color: #333;
    font-size: 2.5rem;
}

.status-badge {
    padding: 0.6rem 1.2rem;
    border-radius: 4px;
    font-size: 0.9rem;
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

.image-gallery {
    margin-bottom: 2rem;
}

.gallery-main {
    position: relative;
    background: #f5f5f5;
    border-radius: 8px;
    overflow: hidden;
    margin-bottom: 1rem;
}

.main-image {
    width: 100%;
    height: 500px;
    object-fit: contain;
    display: block;
}

.gallery-controls {
    position: absolute;
    bottom: 1rem;
    left: 50%;
    transform: translateX(-50%);
    display: flex;
    align-items: center;
    gap: 1rem;
    background: rgba(0, 0, 0, 0.7);
    padding: 0.5rem 1rem;
    border-radius: 20px;
}

.gallery-btn {
    background: transparent;
    border: none;
    color: white;
    font-size: 2rem;
    cursor: pointer;
    padding: 0 0.5rem;
    transition: opacity 0.3s;
}

.gallery-btn:hover {
    opacity: 0.7;
}

.image-counter {
    color: white;
    font-size: 0.9rem;
}

.gallery-thumbnails {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(100px, 1fr));
    gap: 0.5rem;
}

.thumbnail {
    aspect-ratio: 1;
    border: 2px solid transparent;
    border-radius: 4px;
    overflow: hidden;
    cursor: pointer;
    transition: border-color 0.3s;
}

.thumbnail:hover {
    border-color: #007bff;
}

.thumbnail.active {
    border-color: #007bff;
}

.thumbnail img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.no-images {
    background: #f5f5f5;
    border-radius: 8px;
    padding: 3rem;
    text-align: center;
    color: #999;
}

.car-info-grid {
    display: grid;
    gap: 2rem;
    margin-bottom: 2rem;
}

.info-section {
    background: white;
    border-radius: 8px;
    padding: 1.5rem;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.info-section h2 {
    margin: 0 0 1rem 0;
    color: #333;
    font-size: 1.5rem;
}

.info-table {
    display: grid;
    gap: 0.75rem;
}

.info-row {
    display: flex;
    justify-content: space-between;
    padding: 0.75rem;
    background: #f8f9fa;
    border-radius: 4px;
}

.info-label {
    font-weight: 500;
    color: #666;
}

.info-value {
    color: #333;
    font-weight: 600;
}

.description {
    color: #666;
    line-height: 1.8;
    white-space: pre-wrap;
}

.contact-section {
    background: #007bff;
    color: white;
    border-radius: 8px;
    padding: 2rem;
    text-align: center;
}

.contact-section h2 {
    margin: 0 0 1rem 0;
}

.contact-section p {
    margin: 0 0 1.5rem 0;
    opacity: 0.9;
}

.btn-contact {
    padding: 0.75rem 2rem;
    background: white;
    color: #007bff;
    border: none;
    border-radius: 4px;
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
    transition: transform 0.3s;
}

.btn-contact:hover {
    transform: scale(1.05);
}

/* Модальное окно */
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
    .car-detail-container {
        padding: 1rem;
    }

    .car-header {
        flex-direction: column;
        align-items: flex-start;
        gap: 1rem;
    }

    .car-header h1 {
        font-size: 1.8rem;
    }

    .navigation {
        flex-direction: column;
        gap: 1rem;
    }

    .nav-actions {
        width: 100%;
    }

    .btn-back,
    .btn-edit,
    .btn-delete {
        flex: 1;
    }

    .main-image {
        height: 300px;
    }

    .gallery-thumbnails {
        grid-template-columns: repeat(auto-fill, minmax(80px, 1fr));
    }

    .modal {
        padding: 1.5rem;
    }

    .modal-actions {
        flex-direction: column;
    }
}
</style>