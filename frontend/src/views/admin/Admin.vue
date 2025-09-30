<template>
    <div class="admin-panel">
        <div class="container">
            <h1>Админ панель - Заявки на автомобили</h1>

            <!-- Статистика -->
            <div class="stats-bar">
                <div class="stat">
                    <span class="stat-number">{{ totalRequests }}</span>
                    <span class="stat-label">Всего заявок</span>
                </div>
                <div class="stat">
                    <span class="stat-number">{{ totalImages }}</span>
                    <span class="stat-label">Всего фотографий</span>
                </div>
            </div>

            <!-- Загрузка -->
            <div v-if="loading" class="loading">
                <div class="loading-spinner"></div>
                <p>Загружаем данные...</p>
            </div>

            <!-- Ошибка -->
            <div v-if="error" class="error">
                <p>{{ error }}</p>
                <button @click="fetchRequests" class="retry-btn">Попробовать снова</button>
            </div>

            <!-- Список заявок -->
            <div v-if="!loading && !error" class="requests-list">
                <div v-for="request in requests" :key="request.id" class="request-card">
                    <!-- Основная информация -->
                    <div class="request-header">
                        <div class="request-info">
                            <h3>{{ request.name }}</h3>
                            <p class="car-brand">{{ request.carBrand }}</p>
                            <p class="phone">{{ request.phone }}</p>
                            <p class="date">{{ formatDate(request.createdAt) }}</p>
                        </div>
                        <div class="request-stats">
                            <div class="photo-count">
                                <span class="count">{{ request.imageCount }}</span>
                                <span class="label">фото</span>
                            </div>
                            <button @click="toggleDetails(request.id)" class="details-btn">
                                {{ expandedRequests.has(request.id) ? 'Скрыть' : 'Подробнее' }}
                            </button>
                        </div>
                    </div>

                    <!-- Описание -->
                    <div v-if="request.description" class="request-description">
                        <p><strong>Описание:</strong> {{ request.description }}</p>
                    </div>

                    <!-- Детальная информация (разворачивается) -->
                    <div v-if="expandedRequests.has(request.id)" class="request-details">
                        <!-- Загрузка деталей -->
                        <div v-if="loadingDetails.has(request.id)" class="loading-details">
                            <div class="loading-spinner small"></div>
                            <span>Загружаем фотографии...</span>
                        </div>

                        <!-- Ошибка загрузки деталей -->
                        <div v-if="detailsErrors.has(request.id)" class="details-error">
                            <p>Ошибка загрузки: {{ detailsErrors.get(request.id) }}</p>
                            <button @click="fetchRequestDetails(request.id)" class="retry-btn small">Повторить</button>
                        </div>

                        <!-- Фотографии -->
                        <div v-if="requestDetails.has(request.id)" class="photos-section">
                            <h4>Фотографии ({{ requestDetails.get(request.id).images.length }})</h4>
                            <div v-if="requestDetails.get(request.id).images.length > 0" class="photos-grid">
                                <div v-for="image in requestDetails.get(request.id).images" :key="image.id"
                                    class="photo-item">
                                    <img :src="image.fileUrl" :alt="image.fileName" @click="openImageModal(image)"
                                        @error="handleImageError" />
                                    <div class="photo-info">
                                        <p class="photo-name">{{ image.fileName }}</p>
                                        <p class="photo-size">{{ formatFileSize(image.fileSize) }}</p>
                                    </div>
                                </div>
                            </div>
                            <div v-else class="no-photos">
                                <p>Фотографии не загружены</p>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Пагинация -->
                <div v-if="totalPages > 1" class="pagination">
                    <button @click="changePage(currentPage - 1)" :disabled="currentPage <= 1" class="page-btn">
                        Предыдущая
                    </button>

                    <span class="page-info">
                        Страница {{ currentPage }} из {{ totalPages }}
                    </span>

                    <button @click="changePage(currentPage + 1)" :disabled="currentPage >= totalPages" class="page-btn">
                        Следующая
                    </button>
                </div>
            </div>

            <!-- Модальное окно для просмотра изображения -->
            <div v-if="selectedImage" class="image-modal" @click="closeImageModal">
                <div class="modal-content" @click.stop>
                    <button class="close-btn" @click="closeImageModal">&times;</button>
                    <img :src="selectedImage.fileUrl" :alt="selectedImage.fileName" />
                    <div class="image-info">
                        <h4>{{ selectedImage.fileName }}</h4>
                        <p>Размер: {{ formatFileSize(selectedImage.fileSize) }}</p>
                        <p>Загружено: {{ formatDate(selectedImage.createdAt) }}</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'

// Реактивные данные
const requests = ref([])
const loading = ref(false)
const error = ref('')
const currentPage = ref(1)
const totalRequests = ref(0)
const totalPages = ref(0)

// Для управления развернутыми заявками
const expandedRequests = reactive(new Set())
const loadingDetails = reactive(new Set())
const detailsErrors = reactive(new Map())
const requestDetails = reactive(new Map())

// Модальное окно
const selectedImage = ref(null)

// Вычисляемые свойства
const totalImages = computed(() => {
    return requests.value.reduce((sum, req) => sum + req.imageCount, 0)
})

// Методы
const fetchRequests = async () => {
    loading.value = true
    error.value = ''

    try {
        const response = await fetch(`/api/admin/requests?page=${currentPage.value}&limit=10`)

        if (!response.ok) {
            throw new Error(`HTTP ${response.status}`)
        }

        const data = await response.json()

        if (data.status === 'success') {
            requests.value = data.data
            totalRequests.value = data.pagination.total
            totalPages.value = data.pagination.pages
        } else {
            throw new Error(data.error || 'Неизвестная ошибка')
        }
    } catch (err) {
        error.value = `Ошибка загрузки данных: ${err.message}`
        console.error('Fetch error:', err)
    } finally {
        loading.value = false
    }
}

const toggleDetails = async (requestId) => {
    if (expandedRequests.has(requestId)) {
        // Скрываем детали
        expandedRequests.delete(requestId)
    } else {
        // Показываем детали
        expandedRequests.add(requestId)

        // Загружаем детали если их еще нет
        if (!requestDetails.has(requestId)) {
            await fetchRequestDetails(requestId)
        }
    }
}

const fetchRequestDetails = async (requestId) => {
    loadingDetails.add(requestId)
    detailsErrors.delete(requestId)

    try {
        const response = await fetch(`/api/admin/requests/${requestId}`)

        if (!response.ok) {
            throw new Error(`HTTP ${response.status}`)
        }

        const data = await response.json()

        if (data.status === 'success') {
            requestDetails.set(requestId, data.data)
        } else {
            throw new Error(data.error || 'Неизвестная ошибка')
        }
    } catch (err) {
        detailsErrors.set(requestId, err.message)
        console.error('Fetch details error:', err)
    } finally {
        loadingDetails.delete(requestId)
    }
}

const changePage = (newPage) => {
    if (newPage >= 1 && newPage <= totalPages.value) {
        currentPage.value = newPage
        fetchRequests()
    }
}

const openImageModal = (image) => {
    selectedImage.value = image
}

const closeImageModal = () => {
    selectedImage.value = null
}

const handleImageError = (event) => {
    event.target.src = 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAwIiBoZWlnaHQ9IjE1MCIgdmlld0JveD0iMCAwIDIwMCAxNTAiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+CjxyZWN0IHdpZHRoPSIyMDAiIGhlaWdodD0iMTUwIiBmaWxsPSIjZjVmNWY1Ii8+Cjx0ZXh0IHg9IjEwMCIgeT0iNzUiIGZpbGw9IiM5OTk5OTkiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZvbnQtZmFtaWx5PSJBcmlhbCIgZm9udC1zaXplPSIxNCIgZHk9Ii4zZW0iPtCk0L7RgtC+INC90LUg0L3QsNC50LTQtdC90L48L3RleHQ+Cjwvc3ZnPg=='
}

// Утилиты
const formatDate = (dateString) => {
    const date = new Date(dateString)
    return date.toLocaleString('uk-UA', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit'
    })
}

const formatFileSize = (bytes) => {
    if (bytes === 0) return '0 Bytes'
    const k = 1024
    const sizes = ['Bytes', 'KB', 'MB', 'GB']
    const i = Math.floor(Math.log(bytes) / Math.log(k))
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

// Загружаем данные при монтировании компонента
onMounted(() => {
    fetchRequests()
})
</script>

<style lang="scss" scoped>
.admin-panel {
    padding: 2rem;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    min-height: 100vh;
    font-family: 'Arial', sans-serif;
}

.container {
    max-width: 1200px;
    margin: 0 auto;
}

h1 {
    color: white;
    text-align: center;
    margin-bottom: 2rem;
    font-size: 2.5rem;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
}

.stats-bar {
    display: flex;
    gap: 2rem;
    justify-content: center;
    margin-bottom: 2rem;
}

.stat {
    background: rgba(255, 255, 255, 0.2);
    backdrop-filter: blur(10px);
    padding: 1rem 2rem;
    border-radius: 15px;
    text-align: center;
    color: white;
    border: 1px solid rgba(255, 255, 255, 0.3);
}

.stat-number {
    display: block;
    font-size: 2rem;
    font-weight: bold;
    margin-bottom: 0.5rem;
}

.stat-label {
    font-size: 0.9rem;
    opacity: 0.8;
}

.loading,
.error {
    text-align: center;
    color: white;
    padding: 2rem;
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    border-radius: 15px;
    margin: 2rem 0;
}

.loading-spinner {
    width: 40px;
    height: 40px;
    border: 4px solid rgba(255, 255, 255, 0.3);
    border-radius: 50%;
    border-top-color: white;
    animation: spin 1s ease-in-out infinite;
    margin: 0 auto 1rem;

    &.small {
        width: 20px;
        height: 20px;
        border-width: 2px;
        display: inline-block;
        margin: 0 0.5rem 0 0;
    }
}

@keyframes spin {
    to {
        transform: rotate(360deg);
    }
}

.retry-btn {
    background: rgba(255, 255, 255, 0.2);
    color: white;
    border: 1px solid rgba(255, 255, 255, 0.3);
    padding: 0.5rem 1rem;
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.3s ease;

    &:hover {
        background: rgba(255, 255, 255, 0.3);
    }

    &.small {
        padding: 0.25rem 0.5rem;
        font-size: 0.8rem;
    }
}

.request-card {
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);
    border-radius: 20px;
    padding: 1.5rem;
    margin-bottom: 1.5rem;
    border: 1px solid rgba(255, 255, 255, 0.3);
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.request-header {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    margin-bottom: 1rem;
}

.request-info {
    flex: 1;

    h3 {
        margin: 0 0 0.5rem 0;
        color: #333;
        font-size: 1.3rem;
    }

    .car-brand {
        color: #666;
        font-weight: bold;
        margin: 0.25rem 0;
    }

    .phone {
        color: #007bff;
        margin: 0.25rem 0;
    }

    .date {
        color: #999;
        font-size: 0.9rem;
        margin: 0.25rem 0;
    }
}

.request-stats {
    display: flex;
    align-items: center;
    gap: 1rem;
}

.photo-count {
    text-align: center;

    .count {
        display: block;
        font-size: 1.5rem;
        font-weight: bold;
        color: #007bff;
    }

    .label {
        font-size: 0.8rem;
        color: #666;
    }
}

.details-btn {
    background: #007bff;
    color: white;
    border: none;
    padding: 0.5rem 1rem;
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.3s ease;

    &:hover {
        background: #0056b3;
    }
}

.request-description {
    color: #555;
    font-style: italic;
    margin-bottom: 1rem;
    padding: 1rem;
    background: rgba(0, 123, 255, 0.1);
    border-radius: 10px;
}

.request-details {
    margin-top: 1rem;
    padding-top: 1rem;
    border-top: 1px solid #eee;
}

.loading-details,
.details-error {
    padding: 1rem;
    text-align: center;
    color: #666;
}

.photos-section {
    h4 {
        color: #333;
        margin-bottom: 1rem;
    }
}

.photos-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    gap: 1rem;
}

.photo-item {
    border: 1px solid #ddd;
    border-radius: 10px;
    overflow: hidden;
    transition: transform 0.3s ease;
    cursor: pointer;

    &:hover {
        transform: scale(1.05);
        box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
    }

    img {
        width: 100%;
        height: 150px;
        object-fit: cover;
    }
}

.photo-info {
    padding: 0.5rem;

    .photo-name {
        font-size: 0.8rem;
        color: #333;
        margin: 0 0 0.25rem 0;
        word-break: break-word;
    }

    .photo-size {
        font-size: 0.7rem;
        color: #999;
        margin: 0;
    }
}

.no-photos {
    text-align: center;
    color: #666;
    font-style: italic;
    padding: 2rem;
}

.pagination {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 1rem;
    margin-top: 2rem;
}

.page-btn {
    background: rgba(255, 255, 255, 0.2);
    color: white;
    border: 1px solid rgba(255, 255, 255, 0.3);
    padding: 0.5rem 1rem;
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.3s ease;

    &:hover:not(:disabled) {
        background: rgba(255, 255, 255, 0.3);
    }

    &:disabled {
        opacity: 0.5;
        cursor: not-allowed;
    }
}

.page-info {
    color: white;
    font-weight: bold;
}

.image-modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.8);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
}

.modal-content {
    background: white;
    border-radius: 15px;
    padding: 1rem;
    max-width: 90vw;
    max-height: 90vh;
    position: relative;

    img {
        max-width: 100%;
        max-height: 70vh;
        object-fit: contain;
    }
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

    &:hover {
        color: #333;
    }
}

.image-info {
    margin-top: 1rem;

    h4 {
        margin: 0 0 0.5rem 0;
        color: #333;
    }

    p {
        margin: 0.25rem 0;
        color: #666;
        font-size: 0.9rem;
    }
}

@media (max-width: 768px) {
    .admin-panel {
        padding: 1rem;
    }

    .request-header {
        flex-direction: column;
        gap: 1rem;
    }

    .photos-grid {
        grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    }

    .stats-bar {
        flex-direction: column;
        align-items: center;
    }
}
</style>