<template>
    <div class="add-car-container">
        <div class="form-wrapper">
            <h2>Добавить автомобиль в продажу</h2>

            <form @submit.prevent="submitForm" enctype="multipart/form-data">
                <!-- Марка и модель -->
                <div class="form-row">
                    <div class="form-group">
                        <label for="brand">Марка *</label>
                        <input v-model="form.brand" type="text" id="brand" required
                            placeholder="BMW, Mercedes, Toyota..." :class="{ error: errors.brand }">
                        <span v-if="errors.brand" class="error-text">{{ errors.brand }}</span>
                    </div>

                    <div class="form-group">
                        <label for="model">Модель *</label>
                        <input v-model="form.model" type="text" id="model" required placeholder="X5, E-Class, Camry..."
                            :class="{ error: errors.model }">
                        <span v-if="errors.model" class="error-text">{{ errors.model }}</span>
                    </div>
                </div>

                <!-- Год -->
                <div class="form-group">
                    <label for="year">Год выпуска *</label>
                    <input v-model.number="form.year" type="number" id="year" required :min="1950"
                        :max="new Date().getFullYear() + 1" placeholder="2020" :class="{ error: errors.year }">
                    <span v-if="errors.year" class="error-text">{{ errors.year }}</span>
                </div>

                <!-- Топливо и трансмиссия -->
                <div class="form-row">
                    <div class="form-group">
                        <label for="fuel">Топливо *</label>
                        <select v-model="form.fuel" id="fuel" required :class="{ error: errors.fuel }">
                            <option value="">Выберите тип топлива</option>
                            <option value="Бензин">Бензин</option>
                            <option value="Дизель">Дизель</option>
                            <option value="Гибрид">Гибрид</option>
                            <option value="Электро">Электро</option>
                            <option value="Газ">Газ</option>
                            <option value="Газ/Бензин">Газ/Бензин</option>
                        </select>
                        <span v-if="errors.fuel" class="error-text">{{ errors.fuel }}</span>
                    </div>

                    <div class="form-group">
                        <label for="transmission">Трансмиссия *</label>
                        <select v-model="form.transmission" id="transmission" required
                            :class="{ error: errors.transmission }">
                            <option value="">Выберите тип трансмиссии</option>
                            <option value="Механическая">Механическая</option>
                            <option value="Автоматическая">Автоматическая</option>
                            <option value="Робот">Робот</option>
                            <option value="Вариатор">Вариатор</option>
                        </select>
                        <span v-if="errors.transmission" class="error-text">{{ errors.transmission }}</span>
                    </div>
                </div>

                <!-- Привод и пробег -->
                <div class="form-row">
                    <div class="form-group">
                        <label for="drive">Привод *</label>
                        <select v-model="form.drive" id="drive" required :class="{ error: errors.drive }">
                            <option value="">Выберите тип привода</option>
                            <option value="Передний">Передний</option>
                            <option value="Задний">Задний</option>
                            <option value="Полный">Полный</option>
                        </select>
                        <span v-if="errors.drive" class="error-text">{{ errors.drive }}</span>
                    </div>

                    <div class="form-group">
                        <label for="mileage">Пробег (км) *</label>
                        <input v-model.number="form.mileage" type="number" id="mileage" required min="0"
                            placeholder="100000" :class="{ error: errors.mileage }">
                        <span v-if="errors.mileage" class="error-text">{{ errors.mileage }}</span>
                    </div>
                </div>

                <!-- Описание -->
                <div class="form-group">
                    <label for="description">Описание</label>
                    <textarea v-model="form.description" id="description" rows="4"
                        placeholder="Дополнительная информация об автомобиле..."
                        :class="{ error: errors.description }"></textarea>
                    <span v-if="errors.description" class="error-text">{{ errors.description }}</span>
                </div>

                <!-- Фотографии -->
                <div class="form-group">
                    <label for="photos">Фотографии (до 10 штук)</label>
                    <input ref="fileInput" type="file" id="photos" accept="image/*" multiple @change="handleFileUpload"
                        :class="{ error: errors.photos }">
                    <span v-if="errors.photos" class="error-text">{{ errors.photos }}</span>

                    <div class="file-info">
                        Максимальный размер файла: 5MB. Поддерживаемые форматы: JPG, PNG, WebP
                    </div>

                    <!-- Превью фотографий -->
                    <div v-if="photoPreviews.length > 0" class="photos-preview">
                        <div v-for="(preview, index) in photoPreviews" :key="index" class="photo-preview">
                            <img :src="preview" :alt="`Фото ${index + 1}`">
                            <button type="button" @click="removePhoto(index)" class="remove-photo">×</button>
                            <div class="photo-number">{{ index + 1 }}</div>
                        </div>
                    </div>
                </div>

                <!-- Кнопки -->
                <div class="form-actions">
                    <button type="button" @click="resetForm" class="btn-secondary" :disabled="isSubmitting">
                        Очистить
                    </button>
                    <button type="submit" class="btn-primary" :disabled="isSubmitting">
                        <span v-if="!isSubmitting">Добавить автомобиль</span>
                        <span v-else>Сохранение...</span>
                    </button>
                </div>

                <!-- Сообщения -->
                <div v-if="successMessage" class="success-message">
                    {{ successMessage }}
                </div>

                <div v-if="errorMessage" class="error-message">
                    {{ errorMessage }}
                </div>
            </form>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

// Реактивные данные
const form = reactive({
    brand: '',
    model: '',
    year: null,
    fuel: '',
    transmission: '',
    drive: '',
    mileage: null,
    description: '',
    photos: []
})

const errors = reactive({
    brand: null,
    model: null,
    year: null,
    fuel: null,
    transmission: null,
    drive: null,
    mileage: null,
    description: null,
    photos: null
})

const isSubmitting = ref(false)
const successMessage = ref('')
const errorMessage = ref('')
const photoPreviews = ref([])
const fileInput = ref(null)

// Валидация
const validateForm = () => {
    // Очищаем предыдущие ошибки
    Object.keys(errors).forEach(key => errors[key] = null)

    let isValid = true

    // Проверяем обязательные поля
    if (!form.brand.trim()) {
        errors.brand = 'Марка обязательна'
        isValid = false
    }

    if (!form.model.trim()) {
        errors.model = 'Модель обязательна'
        isValid = false
    }

    if (!form.year || form.year < 1950 || form.year > new Date().getFullYear() + 1) {
        errors.year = 'Введите корректный год'
        isValid = false
    }

    if (!form.fuel) {
        errors.fuel = 'Выберите тип топлива'
        isValid = false
    }

    if (!form.transmission) {
        errors.transmission = 'Выберите тип трансмиссии'
        isValid = false
    }

    if (!form.drive) {
        errors.drive = 'Выберите тип привода'
        isValid = false
    }

    if (form.mileage === null || form.mileage < 0) {
        errors.mileage = 'Введите корректный пробег'
        isValid = false
    }

    return isValid
}

// Обработка загрузки файлов
const handleFileUpload = (event) => {
    const files = Array.from(event.target.files)

    if (files.length === 0) return

    if (files.length > 10) {
        errors.photos = 'Максимум 10 фотографий'
        if (fileInput.value) fileInput.value.value = ''
        return
    }

    // Валидация файлов
    const maxSize = 5 * 1024 * 1024 // 5MB
    const allowedTypes = ['image/jpeg', 'image/jpg', 'image/png', 'image/webp', 'image/gif']

    for (const file of files) {
        if (!allowedTypes.includes(file.type.toLowerCase())) {
            errors.photos = 'Разрешены только изображения (JPG, PNG, WebP, GIF)'
            if (fileInput.value) fileInput.value.value = ''
            return
        }

        if (file.size > maxSize) {
            errors.photos = `Файл "${file.name}" слишком большой. Максимум 5MB`
            if (fileInput.value) fileInput.value.value = ''
            return
        }
    }

    // Сохраняем файлы и создаем превью
    form.photos = files
    errors.photos = null

    // Создаем превью
    photoPreviews.value = []
    files.forEach(file => {
        const reader = new FileReader()
        reader.onload = (e) => {
            photoPreviews.value.push(e.target.result)
        }
        reader.readAsDataURL(file)
    })
}

// Удаление фото
const removePhoto = (index) => {
    const newFiles = Array.from(form.photos).filter((_, i) => i !== index)
    form.photos = newFiles
    photoPreviews.value.splice(index, 1)

    // Обновляем input
    if (fileInput.value) {
        const dt = new DataTransfer()
        newFiles.forEach(file => dt.items.add(file))
        fileInput.value.files = dt.files
    }
}

// Отправка формы
const submitForm = async () => {
    if (!validateForm()) return

    isSubmitting.value = true
    successMessage.value = ''
    errorMessage.value = ''

    try {
        const formData = new FormData()

        // Добавляем все поля
        formData.append('brand', form.brand.trim())
        formData.append('model', form.model.trim())
        formData.append('year', form.year.toString())
        formData.append('fuel', form.fuel)
        formData.append('transmission', form.transmission)
        formData.append('drive', form.drive)
        formData.append('mileage', form.mileage.toString())
        formData.append('description', form.description.trim())

        // Добавляем фотографии
        form.photos.forEach(photo => {
            formData.append('images', photo)
        })

        const apiUrl = process.env.NODE_ENV === 'production'
            ? '/api/cars-sale'
            : 'http://localhost:8001/api/cars-sale'

        const response = await fetch(apiUrl, {
            method: 'POST',
            body: formData
        })

        if (response.ok) {
            const result = await response.json()
            successMessage.value = 'Автомобиль успешно добавлен в каталог!'
            resetForm()
        } else {
            const errorData = await response.json().catch(() => ({ error: 'Unknown error' }))
            throw new Error(errorData.error || `HTTP ${response.status}`)
        }

    } catch (error) {
        console.error('Submission error:', error)
        errorMessage.value = 'Ошибка сохранения. Попробуйте еще раз.'
    } finally {
        isSubmitting.value = false
    }
}

// Сброс формы
const resetForm = () => {
    Object.assign(form, {
        brand: '',
        model: '',
        year: null,
        fuel: '',
        transmission: '',
        drive: '',
        mileage: null,
        description: '',
        photos: []
    })

    Object.keys(errors).forEach(key => errors[key] = null)
    photoPreviews.value = []
    successMessage.value = ''
    errorMessage.value = ''

    if (fileInput.value) {
        fileInput.value.value = ''
    }
}
</script>

<style scoped>
.add-car-container {
    max-width: 800px;
    margin: 2rem auto;
    padding: 1rem;
}

.form-wrapper {
    background: white;
    border-radius: 8px;
    padding: 2rem;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

h2 {
    margin: 0 0 2rem 0;
    color: #333;
    text-align: center;
}

.form-row {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 1rem;
}

.form-group {
    margin-bottom: 1.5rem;
}

label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: 500;
    color: #555;
}

input,
select,
textarea {
    width: 100%;
    padding: 0.75rem;
    border: 2px solid #ddd;
    border-radius: 4px;
    font-size: 1rem;
    transition: border-color 0.3s ease;
}

input:focus,
select:focus,
textarea:focus {
    outline: none;
    border-color: #007bff;
}

input.error,
select.error,
textarea.error {
    border-color: #dc3545;
}

.error-text {
    color: #dc3545;
    font-size: 0.875rem;
    margin-top: 0.25rem;
    display: block;
}

.file-info {
    font-size: 0.875rem;
    color: #666;
    margin-top: 0.5rem;
}

.photos-preview {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(100px, 1fr));
    gap: 1rem;
    margin-top: 1rem;
}

.photo-preview {
    position: relative;
    aspect-ratio: 1;
    border: 2px solid #ddd;
    border-radius: 4px;
    overflow: hidden;
}

.photo-preview img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.remove-photo {
    position: absolute;
    top: 4px;
    right: 4px;
    background: rgba(220, 53, 69, 0.9);
    color: white;
    border: none;
    border-radius: 50%;
    width: 24px;
    height: 24px;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 16px;
    line-height: 1;
}

.remove-photo:hover {
    background: #dc3545;
}

.photo-number {
    position: absolute;
    bottom: 4px;
    left: 4px;
    background: rgba(0, 0, 0, 0.7);
    color: white;
    padding: 2px 6px;
    border-radius: 4px;
    font-size: 0.75rem;
}

.form-actions {
    display: flex;
    gap: 1rem;
    justify-content: flex-end;
    margin-top: 2rem;
}

.btn-primary,
.btn-secondary {
    padding: 0.75rem 1.5rem;
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

.btn-primary:disabled {
    background: #6c757d;
    cursor: not-allowed;
}

.btn-secondary {
    background: #6c757d;
    color: white;
}

.btn-secondary:hover:not(:disabled) {
    background: #545b62;
}

.btn-secondary:disabled {
    background: #adb5bd;
    cursor: not-allowed;
}

.success-message {
    background: #d4edda;
    color: #155724;
    padding: 1rem;
    border-radius: 4px;
    border: 1px solid #c3e6cb;
    margin-top: 1rem;
}

.error-message {
    background: #f8d7da;
    color: #721c24;
    padding: 1rem;
    border-radius: 4px;
    border: 1px solid #f5c6cb;
    margin-top: 1rem;
}

@media (max-width: 768px) {
    .add-car-container {
        margin: 1rem;
        padding: 0;
    }

    .form-wrapper {
        padding: 1rem;
    }

    .form-row {
        grid-template-columns: 1fr;
    }

    .form-actions {
        flex-direction: column;
    }

    .photos-preview {
        grid-template-columns: repeat(auto-fill, minmax(80px, 1fr));
    }
}
</style>