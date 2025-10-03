<template>
    <div class="form-container">
        <h2 class="form-title">Розкажіть про Ваше авто !</h2>

        <form @submit.prevent="submitForm" novalidate class="form-banner">
            <!-- Поле имени -->
            <div class="input-group"
                :class="{ 'has-error': errors.name, 'has-success': !errors.name && form.name.length > 0 }">
                <input v-model="form.name" type="text" placeholder="Ваше Ім'я" required class="glass-input"
                    @blur="validateName" @input="clearError('name')">
                <div v-if="errors.name" class="error-message">{{ errors.name }}</div>
            </div>

            <!-- Поле марки авто -->
            <div class="input-group"
                :class="{ 'has-error': errors.carBrand, 'has-success': !errors.carBrand && form.carBrand.length > 0 }">
                <input v-model="form.carBrand" type="text" placeholder="Марка авто" required class="glass-input"
                    @blur="validateCarBrand" @input="clearError('carBrand')">
                <div v-if="errors.carBrand" class="error-message">{{ errors.carBrand }}</div>
            </div>

            <div class="input-group"
                :class="{ 'has-error': errors.carModel, 'has-success': !errors.carModel && form.carModel.length > 0 }">
                <input v-model="form.carModel" type="text" placeholder="Марка авто" required class="glass-input"
                    @blur="validateCarModel" @input="clearError('carModel')">
                <div v-if="errors.carModel" class="error-message">{{ errors.carModeld }}</div>
            </div>

            <!-- Поле телефона -->
            <div class="input-group"
                :class="{ 'has-error': errors.phone, 'has-success': !errors.phone && form.phone.length > 0 }">
                <input v-model="form.phone" type="tel" placeholder="+380 (XX) XXX-XX-XX" required class="glass-input"
                    @input="formatPhone" @blur="validatePhone" maxlength="19">
                <div v-if="errors.phone" class="error-message">{{ errors.phone }}</div>
            </div>

            <!-- Поле описания (необязательное) -->
            <div class="input-group"
                :class="{ 'has-error': errors.description, 'has-success': !errors.description && form.description.length > 0 }">
                <textarea v-model="form.description" placeholder="Розкажіть трошки про Ваше авто (необов'язково)"
                    class="glass-textarea" rows="4" @blur="validateDescription"
                    @input="clearError('description')"></textarea>
                <div v-if="errors.description" class="error-message">{{ errors.description }}</div>
            </div>

            <!-- Поле загрузки фото (необязательное) -->
            <div class="input-group file-container"
                :class="{ 'has-error': errors.photos, 'has-success': form.photos.length > 0 }">
                <div class="file-upload-wrapper">
                    <input ref="fileInput" type="file" accept="image/*" @change="handleFileUpload" class="file-input"
                        id="photo-upload" multiple>
                    <label for="photo-upload" class="file-label">
                        <span class="file-text">
                            Можете надати фото (необов'язково, максимум 10)
                        </span>
                    </label>
                </div>
                <div v-if="errors.photos" class="error-message">{{ errors.photos }}</div>

                <!-- Превью фото -->
                <div v-if="photoPreviews.length > 0" class="photos-preview">
                    <div v-for="(preview, index) in photoPreviews" :key="index" class="photo-preview-item">
                        <img :src="preview" :alt="`Фото ${index + 1}`">
                        <button type="button" @click="removePhoto(index)" class="remove-photo">✕</button>
                    </div>
                </div>
            </div>

            <!-- Кнопка отправки -->
            <button type="submit" class="submit-btn" :class="{ 'loading': isSubmitting }"
                :disabled="isSubmitting || !isFormValid">
                <span v-if="!isSubmitting">Надіслати нам</span>
                <span v-else class="loading-spinner"></span>
            </button>

            <!-- Сообщение об успехе -->
            <div v-if="successMessage" class="success-message">
                {{ successMessage }}
            </div>
        </form>
    </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'

// Реактивные данные
const form = reactive({
    name: '',
    carBrand: '',
    phone: '',
    description: '',
    carModel: '',
    photos: []
})

const errors = reactive({
    name: null,
    carBrand: null,
    carModel: null,
    phone: null,
    description: null,
    photos: null
})

const isSubmitting = ref(false)
const successMessage = ref('')
const photoPreviews = ref([])
const fileInput = ref(null)

// Computed свойства
const isFormValid = computed(() => {
    return (
        form.name.length >= 2 &&
        form.carBrand.length >= 2 &&
        form.carModel.length >= 2 &&
        isValidUkrainianPhone(form.phone) &&
        !Object.values(errors).some(error => error !== null)
    )
})

// Методы валидации
const validateName = () => {
    if (!form.name.trim()) {
        errors.name = 'Поле "Ім\'я" є обов\'язковим'
    } else if (form.name.trim().length < 2) {
        errors.name = 'Ім\'я повинно містити мінімум 2 символи'
    } else if (!/^[а-яА-ЯёЁіІїЇєЄa-zA-Z\s\-\']+$/.test(form.name)) {
        errors.name = 'Ім\'я може містити тільки букви, пробіли, дефіси та апострофи'
    } else {
        errors.name = null
    }
}

const validateCarBrand = () => {
    if (!form.carBrand.trim()) {
        errors.carBrand = 'Поле "Марка авто" є обов\'язковим'
    } else if (form.carBrand.trim().length < 2) {
        errors.carBrand = 'Марка авто повинна містити мінімум 2 символи'
    } else {
        errors.carBrand = null
    }
}

const validateCarModel = () => {
    if (!form.carModel.trim()) {
        errors.carModel = 'Поле "Модель авто" є обов\'язковим'
    } else if (form.carModel.trim().length < 2) {
        errors.carModel = 'Модель авто повинна містити мінімум 2 символи'
    } else {
        errors.carModel = null
    }
}

const formatPhone = () => {
    let value = form.phone.replace(/\D/g, '')

    // Если поле пустое, оставляем пустым
    if (!value) {
        form.phone = ''
        clearError('phone')
        return
    }

    // Если начинается с 0, заменяем на 380
    if (value.startsWith('0')) {
        value = '380' + value.slice(1)
    }

    // Если не начинается с 380, добавляем 380
    if (!value.startsWith('380')) {
        value = '380' + value
    }

    // Ограничиваем до 12 цифр (380 + 9 цифр)
    value = value.slice(0, 12)

    // Если осталось меньше 4 цифр (380 + минимум 1), показываем только то что есть
    if (value.length <= 3) {
        form.phone = '+' + value
        clearError('phone')
        return
    }

    // Форматируем номер поэтапно
    let formatted = '+380'
    const phoneDigits = value.slice(3) // Цифры после 380

    if (phoneDigits.length >= 1) {
        formatted += ' (' + phoneDigits.slice(0, 2)

        if (phoneDigits.length > 2) {
            formatted += ') ' + phoneDigits.slice(2, 5)

            if (phoneDigits.length > 5) {
                formatted += '-' + phoneDigits.slice(5, 7)

                if (phoneDigits.length > 7) {
                    formatted += '-' + phoneDigits.slice(7, 9)
                }
            }
        }
    }

    form.phone = formatted
    clearError('phone')
}

const validatePhone = () => {
    if (!form.phone.trim()) {
        errors.phone = 'Поле "Телефон" є обов\'язковим'
    } else if (!isValidUkrainianPhone(form.phone)) {
        errors.phone = 'Введіть коректний номер телефону'
    } else {
        errors.phone = null
    }
}

const validateDescription = () => {
    // Описание теперь необязательное
    if (form.description.trim() && form.description.trim().length > 500) {
        errors.description = 'Опис не може містити більше 500 символів'
    } else {
        errors.description = null
    }
}

// Проверка украинского номера
const isValidUkrainianPhone = (phone) => {
    const phoneRegex = /^\+380 \(\d{2}\) \d{3}-\d{2}-\d{2}$/
    return phoneRegex.test(phone)
}


const removePhoto = (index) => {
    form.photos.splice(index, 1)
    photoPreviews.value.splice(index, 1)

    // Очищаем input, чтобы можно было загрузить те же файлы снова
    if (fileInput.value) {
        fileInput.value.value = ''
    }

    // Убираем ошибку если фото удалены
    if (form.photos.length === 0) {
        errors.photos = null
    }
}

// Очистка ошибки
const clearError = (field) => {
    errors[field] = null
}

// Отправка формы
// Improved submitForm method for the Vue component
const submitForm = async () => {
    validateName()
    validateCarBrand()
    validateCarModel()
    validatePhone()
    validateDescription()

    if (!isFormValid.value) return

    isSubmitting.value = true
    successMessage.value = ''

    try {
        const formData = new FormData()
        formData.append('name', form.name.trim())
        formData.append('carBrand', form.carBrand.trim())
        formData.append('carModel', form.carModel.trim())
        formData.append('phone', form.phone.trim())
        formData.append('description', form.description.trim())

        // Add photos
        form.photos.forEach((photo) => {
            formData.append('images', photo)
        })

        // Get the correct API URL based on environment
        const apiUrl = process.env.NODE_ENV === 'production'
            ? '/api/cars'
            : 'http://localhost:8001/api/cars'

        const response = await fetch(apiUrl, {
            method: 'POST',
            body: formData,
            // Don't set Content-Type header - let browser set it with boundary for multipart
        })

        if (response.ok) {
            const result = await response.json()
            successMessage.value = 'Дякуємо! Ваша заявка успішно відправлена.'
            resetForm()

            // Optional: scroll to success message
            setTimeout(() => {
                const successEl = document.querySelector('.success-message')
                if (successEl) {
                    successEl.scrollIntoView({ behavior: 'smooth', block: 'center' })
                }
            }, 100)

        } else {
            const errorData = await response.json().catch(() => ({ error: 'Unknown error' }))
            throw new Error(errorData.error || `HTTP ${response.status}`)
        }
    } catch (error) {
        console.error('Submission error:', error)

        // More specific error messages
        let errorMessage = 'Помилка відправки. Спробуйте ще раз.'

        if (error.message.includes('fetch')) {
            errorMessage = 'Помилка з\'єднання з сервером. Перевірте інтернет-з\'єднання.'
        } else if (error.message.includes('413')) {
            errorMessage = 'Файли занадто великі. Зменшіть розмір фото.'
        } else if (error.message.includes('400')) {
            errorMessage = 'Некоректні дані. Перевірте заповнені поля.'
        }

        successMessage.value = errorMessage

        // Auto-clear error message after 5 seconds
        setTimeout(() => {
            if (successMessage.value === errorMessage) {
                successMessage.value = ''
            }
        }, 5000)

    } finally {
        isSubmitting.value = false
    }
}

// Improved file validation
const handleFileUpload = (event) => {
    const files = Array.from(event.target.files)

    if (files.length === 0) return

    // Check total number of photos
    const totalPhotos = form.photos.length + files.length
    if (totalPhotos > 10) {
        errors.photos = `Можна завантажити максимум 10 фото. Зараз вибрано ${files.length}, а вже є ${form.photos.length}`
        // Clear file input
        if (fileInput.value) {
            fileInput.value.value = ''
        }
        return
    }

    // Validate each file
    const validFiles = []
    const validPreviews = []
    const maxSize = 5 * 1024 * 1024 // 5MB
    const allowedTypes = ['image/jpeg', 'image/jpg', 'image/png', 'image/webp', 'image/gif']

    for (const file of files) {
        // Check file type
        if (!allowedTypes.includes(file.type.toLowerCase())) {
            errors.photos = 'Дозволені формати: JPG, PNG, WebP, GIF'
            if (fileInput.value) {
                fileInput.value.value = ''
            }
            return
        }

        // Check file size
        if (file.size > maxSize) {
            const fileSizeMB = (file.size / (1024 * 1024)).toFixed(1)
            errors.photos = `Файл "${file.name}" занадто великий (${fileSizeMB}MB). Максимальний розмір: 5MB`
            if (fileInput.value) {
                fileInput.value.value = ''
            }
            return
        }

        validFiles.push(file)
    }

    // Add valid files
    form.photos.push(...validFiles)
    errors.photos = null

    // Create previews for new files
    validFiles.forEach(file => {
        const reader = new FileReader()
        reader.onload = (e) => {
            photoPreviews.value.push(e.target.result)
        }
        reader.onerror = () => {
            console.error('Error reading file:', file.name)
        }
        reader.readAsDataURL(file)
    })
}

const resetForm = () => {
    Object.assign(form, {
        name: '',
        carBrand: '',
        carModel: '',
        phone: '',
        description: '',
        photos: []
    })
    Object.assign(errors, {
        name: null,
        carBrand: null,
        carModel: '',
        phone: null,
        description: null,
        photos: null
    })
    photoPreviews.value = []
    if (fileInput.value) {
        fileInput.value.value = ''
    }
}
</script>

<style lang="scss" scoped>
.form-container {
    width: 100%;
    max-width: 520px;
    background: #FFFFFF;
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 40px;
    border-radius: 20px;
}


.form-title {
    font-family: Work Sans;
    font-weight: 600;
    font-style: SemiBold;
    font-size: 28px;
    line-height: 100%;
    letter-spacing: 0%;
    text-align: center;
    color: #000;
}

.file-container {
    flex-direction: column;
}

.form-banner {
    display: flex;
    flex-direction: column;
    gap: 20px;
    width: 100%;
}

.input-group {
    position: relative;
    display: flex;
    flex-direction: column;
    background: #FAFAFA;
    border-radius: 12px;
    padding: 16px;

    & input,
    textarea {
        font-family: Work Sans;
        font-weight: 400;
        font-style: Regular;
        font-size: 16px;
        line-height: 20px;
        letter-spacing: 0%;
        color: #000;
        background: #FAFAFA;
        border: none;
        max-height: 150px;
    }

    &.has-error {

        .glass-input,
        .glass-textarea {
            border-bottom: 1px solid;
            border-color: rgba(231, 76, 60, 0.6);
        }

        .textarea-icon {
            color: #e74c3c;
        }
    }

    &.has-success {

        .glass-input,
        .glass-textarea {
            border-bottom: 1px solid;
            border-color: rgba(39, 174, 96, 0.6);
        }

        .textarea-icon {
            color: #27ae60;
        }
    }
}

.glass-input,
.glass-textarea {
    width: 100%;
    color: #000;
    font-size: 1rem;
    transition: all 0.3s ease;

    &::placeholder {
        color: #666;
    }

    &:focus {
        outline: none;
        border-color: rgba(255, 255, 255, 0.6);
        box-shadow: 0 0 0 3px rgba(255, 255, 255, 0.2);
        background: rgba(255, 255, 255, 0.3);
    }
}

.glass-textarea {
    resize: vertical;
    min-height: 100px;
    font-family: inherit;
}


.error-message {
    color: #ff6b6b;
    font-size: 0.85rem;
    margin-top: 0.5rem;
    padding: 10px;
}

// Стили для загрузки файлов
.file-upload-wrapper {
    position: relative;
    display: contents;
}

.file-input {
    position: absolute;
    opacity: 0;
    width: 0;
    height: 0;
}

.file-label {
    display: flex;
    align-items: center;
    cursor: pointer;
    transition: all 0.3s ease;

    &:hover {
        opacity: 0.6;
    }
}

.file-icon {
    font-size: 1.5rem;
    margin-right: 0.8rem;
    color: rgba(255, 255, 255, 0.8);
}

.file-text {
    font-family: Work Sans;
    font-weight: 400;
    font-style: Regular;
    font-size: 16px;
    line-height: 20px;
    letter-spacing: 0%;
    color: #000;
}

.photos-preview {
    display: flex;
    flex-wrap: wrap;
    gap: 0.5rem;
    margin-top: 1rem;
}

.photo-preview-item {
    position: relative;
    width: 50px;
    height: 50px;

    img {
        width: 100%;
        height: 100%;
        object-fit: cover;
        border-radius: 10px;
    }

    .remove-photo {
        position: absolute;
        z-index: 5;
        top: -5px;
        right: -5px;
        background: rgba(231, 76, 60, 0.9);
        color: white;
        border: none;
        border-radius: 50%;
        width: 18px;
        height: 18px;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 0.7rem;
        backdrop-filter: blur(5px);
        transition: all 0.2s ease;

        &:hover {
            background: rgba(231, 76, 60, 1);
            transform: scale(1.1);
        }
    }
}

.submit-btn {
    width: 100%;
    padding: 1rem;
    background-color: #27ae60;
    backdrop-filter: blur(10px);
    border: 2px solid rgba(255, 255, 255, 0.4);
    border-radius: 15px;
    color: rgba(255, 255, 255, 0.9);
    font-size: 1.1rem;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
    position: relative;
    overflow: hidden;

    &:hover:not(:disabled) {
        transform: translateY(-2px);
        opacity: 0.6;
    }

    &:disabled {
        opacity: 0.6;
        cursor: not-allowed;
        transform: none;
    }

    &.loading {
        color: transparent;
    }
}

.loading-spinner {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 20px;
    height: 20px;
    border: 2px solid rgba(255, 255, 255, 0.3);
    border-radius: 50%;
    border-top-color: rgba(255, 255, 255, 0.9);
    animation: spin 1s ease-in-out infinite;
}

@keyframes spin {
    to {
        transform: translate(-50%, -50%) rotate(360deg);
    }
}

.success-message {
    background: rgba(39, 174, 96, 0.2);
    color: rgba(255, 255, 255, 0.9);
    padding: 1rem;
    border-radius: 12px;
    text-align: center;
    margin-top: 1rem;
    border: 1px solid rgba(39, 174, 96, 0.4);
    backdrop-filter: blur(5px);
}

// Декоративные элементы
.decoration-circles {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    overflow: hidden;
    z-index: 1;
}

.circle {
    position: absolute;
    border-radius: 50%;
    opacity: 0.1;
    animation: float 6s ease-in-out infinite;

    &.circle-1 {
        width: 200px;
        height: 200px;
        background: rgba(255, 255, 255, 0.3);
        top: 10%;
        left: -5%;
        animation-delay: 0s;
    }

    &.circle-2 {
        width: 150px;
        height: 150px;
        background: rgba(255, 255, 255, 0.2);
        top: 60%;
        right: -10%;
        animation-delay: 2s;
    }

    &.circle-3 {
        width: 100px;
        height: 100px;
        background: rgba(255, 255, 255, 0.25);
        bottom: 20%;
        left: 20%;
        animation-delay: 4s;
    }
}

@keyframes float {

    0%,
    100% {
        transform: translateY(0px);
    }

    50% {
        transform: translateY(-20px);
    }
}

// Адаптивность
@media (max-width: 768px) {
    .form-container {
        padding: 1rem;
    }

    .glass-form {
        padding: 1.5rem;
    }

    .form-title {
        font-size: 1.5rem;
    }
}
</style>