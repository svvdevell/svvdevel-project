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

            <!-- Поле модели авто -->
            <div class="input-group"
                :class="{ 'has-error': errors.carModel, 'has-success': !errors.carModel && form.carModel.length > 0 }">
                <input v-model="form.carModel" type="text" placeholder="Модель авто" required class="glass-input"
                    @blur="validateCarModel" @input="clearError('carModel')">
                <div v-if="errors.carModel" class="error-message">{{ errors.carModel }}</div>
            </div>

            <!-- Селект года выпуска -->
            <div class="input-group"
                :class="{ 'has-error': errors.carYear, 'has-success': !errors.carYear && form.carYear }">
                <select v-model="form.carYear" required class="glass-select" @blur="validateCarYear"
                    @change="clearError('carYear')">
                    <option value="" disabled selected>Рік випуску</option>
                    <option v-for="year in years" :key="year" :value="year">{{ year }}</option>
                </select>
                <div v-if="errors.carYear" class="error-message">{{ errors.carYear }}</div>
            </div>

            <!-- Селект трансмиссии -->
            <div class="input-group"
                :class="{ 'has-error': errors.carTrans, 'has-success': !errors.carTrans && form.carTrans }">
                <select v-model="form.carTrans" required class="glass-select" @blur="validateCarTrans"
                    @change="clearError('carTrans')">
                    <option value="" disabled selected>Трансмісія</option>
                    <option value="Механічна">Механічна</option>
                    <option value="Автоматична">Автоматична</option>
                    <option value="Робот">Робот</option>
                    <option value="Варіатор">Варіатор</option>
                </select>
                <div v-if="errors.carTrans" class="error-message">{{ errors.carTrans }}</div>
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

// Генерация годов от 1980 до текущего года
const currentYear = new Date().getFullYear()
const years = ref([])
for (let year = currentYear; year >= 1980; year--) {
    years.value.push(year)
}

// Реактивные данные
const form = reactive({
    name: '',
    carBrand: '',
    carModel: '',
    carYear: '',
    carTrans: '',
    phone: '',
    description: ''
})

const errors = reactive({
    name: null,
    carBrand: null,
    carModel: null,
    carYear: null,
    carTrans: null,
    phone: null,
    description: null
})

const isSubmitting = ref(false)
const successMessage = ref('')

// Computed свойства
const isFormValid = computed(() => {
    return (
        form.name.length >= 2 &&
        form.carBrand.length >= 2 &&
        form.carModel.length >= 2 &&
        form.carYear !== '' &&
        form.carTrans !== '' &&
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

const validateCarYear = () => {
    if (!form.carYear) {
        errors.carYear = 'Оберіть рік випуску авто'
    } else {
        errors.carYear = null
    }
}

const validateCarTrans = () => {
    if (!form.carTrans) {
        errors.carTrans = 'Оберіть тип трансмісії'
    } else {
        errors.carTrans = null
    }
}

const formatPhone = () => {
    let value = form.phone.replace(/\D/g, '')

    if (!value) {
        form.phone = ''
        clearError('phone')
        return
    }

    if (value.startsWith('0')) {
        value = '380' + value.slice(1)
    }

    if (!value.startsWith('380')) {
        value = '380' + value
    }

    value = value.slice(0, 12)

    if (value.length <= 3) {
        form.phone = '+' + value
        clearError('phone')
        return
    }

    let formatted = '+380'
    const phoneDigits = value.slice(3)

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
    if (form.description.trim() && form.description.trim().length > 500) {
        errors.description = 'Опис не може містити більше 500 символів'
    } else {
        errors.description = null
    }
}

const isValidUkrainianPhone = (phone) => {
    const phoneRegex = /^\+380 \(\d{2}\) \d{3}-\d{2}-\d{2}$/
    return phoneRegex.test(phone)
}

const clearError = (field) => {
    errors[field] = null
}

const submitForm = async () => {
    validateName()
    validateCarBrand()
    validateCarModel()
    validateCarYear()
    validateCarTrans()
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
        formData.append('carYear', form.carYear)
        formData.append('carTrans', form.carTrans)
        formData.append('phone', form.phone.trim())
        formData.append('description', form.description.trim())

        const apiUrl = process.env.NODE_ENV === 'production'
            ? '/api/cars'
            : 'http://localhost:8001/api/cars'

        const response = await fetch(apiUrl, {
            method: 'POST',
            body: formData,
        })

        if (response.ok) {
            const result = await response.json()
            successMessage.value = 'Дякуємо! Ваша заявка успішно відправлена.'
            resetForm()

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

        let errorMessage = 'Помилка відправки. Спробуйте ще раз.'

        if (error.message.includes('fetch')) {
            errorMessage = 'Помилка з\'єднання з сервером. Перевірте інтернет-з\'єднання.'
        } else if (error.message.includes('400')) {
            errorMessage = 'Некоректні дані. Перевірте заповнені поля.'
        }

        successMessage.value = errorMessage

        setTimeout(() => {
            if (successMessage.value === errorMessage) {
                successMessage.value = ''
            }
        }, 5000)

    } finally {
        isSubmitting.value = false
    }
}

const resetForm = () => {
    Object.assign(form, {
        name: '',
        carBrand: '',
        carModel: '',
        carYear: '',
        carTrans: '',
        phone: '',
        description: ''
    })
    Object.assign(errors, {
        name: null,
        carBrand: null,
        carModel: null,
        carYear: null,
        carTrans: null,
        phone: null,
        description: null
    })
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
    textarea,
    select {
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
        .glass-textarea,
        .glass-select {
            border-bottom: 1px solid;
            border-color: rgba(231, 76, 60, 0.6);
        }
    }

    &.has-success {

        .glass-input,
        .glass-textarea,
        .glass-select {
            border-bottom: 1px solid;
            border-color: rgba(39, 174, 96, 0.6);
        }
    }
}

.glass-input,
.glass-textarea,
.glass-select {
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

.glass-select {
    cursor: pointer;
    padding-right: 30px;
    appearance: none;
    background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='12' height='12' viewBox='0 0 12 12'%3E%3Cpath fill='%23666' d='M6 9L1 4h10z'/%3E%3C/svg%3E");
    background-repeat: no-repeat;
    background-position: right 10px center;

    option {
        background: #FAFAFA;
        color: #000;
    }

    &:disabled {
        color: #999;
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

@media (max-width: 768px) {
    .form-container {
        padding: 1rem;
    }

    .form-title {
        font-size: 1.5rem;
    }
}
</style>