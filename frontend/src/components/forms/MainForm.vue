<template>
    <div class="form-inner">
        <div class="form-container">
            <h2 class="form-title">Розкажіть про Ваше авто !</h2>

            <form @submit.prevent="submitForm" novalidate class="form-banner">
                <!-- Honeypot поле (скрытое от пользователей) -->
                <input type="text" v-model="honeypot" name="website"
                    style="position:absolute;left:-5000px;width:1px;height:1px" tabindex="-1" autocomplete="off"
                    aria-hidden="true">

                <!-- Поле имени -->
                <div class="input-group"
                    :class="{ 'has-error': errors.name, 'has-success': !errors.name && form.name.length > 0 }">
                    <input v-model="form.name" type="text" placeholder="Ваше Ім'я" required class="glass-input"
                        @blur="validateName" @input="clearError('name')" maxlength="30">
                    <div v-if="errors.name" class="error-message">{{ errors.name }}</div>
                </div>

                <!-- Поле марки авто с автокомплитом -->
                <div class="input-group autocomplete-group"
                    :class="{ 'has-error': errors.carBrand, 'has-success': !errors.carBrand && form.carBrand.length > 0 }">
                    <input v-model="form.carBrand" type="text" placeholder="Марка авто" required class="glass-input"
                        @input="onBrandInput" @focus="onBrandFocus" @blur="onBrandBlur" autocomplete="off">

                    <!-- Выпадающий список марок -->
                    <div v-if="showBrandDropdown && filteredBrands.length > 0" class="autocomplete-dropdown">
                        <div v-for="brand in filteredBrands" :key="brand" class="autocomplete-item"
                            @mousedown="selectBrand(brand)">
                            {{ brand }}
                        </div>
                    </div>

                    <div v-if="errors.carBrand" class="error-message">{{ errors.carBrand }}</div>
                </div>

                <!-- Поле модели авто -->
                <div class="input-group"
                    :class="{ 'has-error': errors.carModel, 'has-success': !errors.carModel && form.carModel.length > 0 }">
                    <input v-model="form.carModel" type="text" placeholder="Модель авто" required class="glass-input"
                        @blur="validateCarModel" @input="clearError('carModel')" maxlength="50">
                    <div v-if="errors.carModel" class="error-message">{{ errors.carModel }}</div>
                </div>

                <!-- Селект года выпуска с автокомплитом -->
                <div class="input-group autocomplete-group"
                    :class="{ 'has-error': errors.carYear, 'has-success': !errors.carYear && form.carYear }">
                    <input v-model="form.carYear" type="text" placeholder="Рік випуску" required class="glass-input"
                        @input="onYearInput" @focus="onYearFocus" @blur="onYearBlur" autocomplete="off" maxlength="4">

                    <!-- Выпадающий список годов -->
                    <div v-if="showYearDropdown && filteredYears.length > 0" class="autocomplete-dropdown">
                        <div v-for="year in filteredYears.slice(0, 10)" :key="year" class="autocomplete-item"
                            @mousedown="selectYear(year)">
                            {{ year }}
                        </div>
                    </div>

                    <div v-if="errors.carYear" class="error-message">{{ errors.carYear }}</div>
                </div>

                <!-- Селект трансмиссии с автокомплитом -->
                <div class="input-group autocomplete-group"
                    :class="{ 'has-error': errors.carTrans, 'has-success': !errors.carTrans && form.carTrans }">
                    <input v-model="form.carTrans" type="text" placeholder="Трансмісія" required class="glass-input"
                        @input="onTransInput" @focus="onTransFocus" @blur="onTransBlur" autocomplete="off">

                    <!-- Выпадающий список трансмиссий -->
                    <div v-if="showTransDropdown && filteredTransmissions.length > 0" class="autocomplete-dropdown">
                        <div v-for="trans in filteredTransmissions" :key="trans" class="autocomplete-item"
                            @mousedown="selectTrans(trans)">
                            {{ trans }}
                        </div>
                    </div>

                    <div v-if="errors.carTrans" class="error-message">{{ errors.carTrans }}</div>
                </div>

                <!-- Поле телефона -->
                <div class="input-group"
                    :class="{ 'has-error': errors.phone, 'has-success': !errors.phone && form.phone.length > 0 }">
                    <input v-model="form.phone" type="tel" placeholder="+380 (XX) XXX-XX-XX" required
                        class="glass-input" @input="formatPhone" @blur="validatePhone" maxlength="19">
                    <div v-if="errors.phone" class="error-message">{{ errors.phone }}</div>
                </div>

                <!-- Поле описания (необязательное) -->
                <!-- <div class="input-group"
                :class="{ 'has-error': errors.description, 'has-success': !errors.description && form.description.length > 0 }">
                <textarea v-model="form.description" placeholder="Розкажіть трошки про Ваше авто (необов'язково)"
                    class="glass-textarea" rows="4" @blur="validateDescription" @input="clearError('description')"
                    maxlength="500"></textarea>
                <div v-if="errors.description" class="error-message">{{ errors.description }}</div>
                <div v-if="form.description.length > 0" class="char-counter">
                    {{ form.description.length }} / 500
                </div>
            </div> -->

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
    </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'

// Популярные марки автомобилей (расширенный список на основе украинского рынка)
const carBrands = [
    'Acura', 'Alfa Romeo', 'Aston Martin', 'Audi', 'Bentley', 'BMW', 'Brilliance',
    'Buick', 'BYD', 'Cadillac', 'Changan', 'Chery', 'Chevrolet', 'Chrysler',
    'Citroen', 'Cupra', 'Dacia', 'Daewoo', 'Daihatsu', 'Denza', 'Dodge', 'Dongfeng',
    'DS', 'FAW', 'Ferrari', 'Fiat', 'Ford', 'Foton', 'Gac', 'Geely', 'Genesis',
    'GMC', 'Great Wall', 'Haval', 'Honda', 'Hongqi', 'Hummer', 'Hyundai', 'Infiniti',
    'Isuzu', 'Iveco', 'JAC', 'Jaguar', 'Jeep', 'Jetour', 'Kia', 'Lamborghini',
    'Lancia', 'Land Rover', 'Lexus', 'Lifan', 'Lincoln', 'Lotus', 'Mahindra',
    'Maserati', 'Maxus', 'Maybach', 'Mazda', 'McLaren', 'Mercedes-Benz', 'Mercury',
    'MG', 'MINI', 'Mitsubishi', 'Nissan', 'Opel', 'Peugeot', 'Polestar', 'Pontiac',
    'Porsche', 'Proton', 'Ram', 'Ravon', 'Renault', 'Rolls-Royce', 'Rover', 'Saab',
    'Samsung', 'SEAT', 'Skoda', 'Smart', 'SsangYong', 'Subaru', 'Suzuki', 'TATA',
    'Tesla', 'Toyota', 'Volkswagen', 'Volvo', 'Xpeng', 'Zeekr', 'Zotye',
    'Богдан', 'ВАЗ', 'ГАЗ', 'ЗАЗ', 'ЗИЛ', 'ІЖ', 'ЛуАЗ', 'Москвич', 'УАЗ'
].sort()

// Генерация годов от 1980 до текущего года
const currentYear = new Date().getFullYear()
const years = ref([])
for (let year = currentYear; year >= 1980; year--) {
    years.value.push(year)
}

// Типы трансмиссий
const transmissions = [
    'Механічна',
    'Автоматична',
    'Робот',
    'Варіатор',
    'Tiptronic'
]

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

// Защита от ботов
const honeypot = ref('')
const formStartTime = ref(Date.now())
const formToken = ref('')

// Для автокомплита марки
const showBrandDropdown = ref(false)
const brandSearchQuery = ref('')

// Для автокомплита года
const showYearDropdown = ref(false)
const yearSearchQuery = ref('')

// Для автокомплита трансмиссии
const showTransDropdown = ref(false)
const transSearchQuery = ref('')

// Фильтрованный список марок
const filteredBrands = computed(() => {
    if (!brandSearchQuery.value) return carBrands

    const query = brandSearchQuery.value.toLowerCase()
    return carBrands.filter(brand =>
        brand.toLowerCase().includes(query)
    )
})

// Фильтрованный список годов
const filteredYears = computed(() => {
    if (!yearSearchQuery.value) return years.value

    const query = yearSearchQuery.value
    return years.value.filter(year =>
        year.toString().includes(query)
    )
})

// Фильтрованный список трансмиссий
const filteredTransmissions = computed(() => {
    if (!transSearchQuery.value) return transmissions

    const query = transSearchQuery.value.toLowerCase()
    return transmissions.filter(trans =>
        trans.toLowerCase().includes(query)
    )
})

// Обработчики для автокомплита марки
const onBrandInput = (event) => {
    form.carBrand = event.target.value
    brandSearchQuery.value = event.target.value
    showBrandDropdown.value = true
    clearError('carBrand')
}

const selectBrand = (brand) => {
    form.carBrand = brand
    brandSearchQuery.value = brand
    showBrandDropdown.value = false
    clearError('carBrand')
}

const onBrandFocus = () => {
    brandSearchQuery.value = form.carBrand
    showBrandDropdown.value = true
}

const onBrandBlur = () => {
    setTimeout(() => {
        showBrandDropdown.value = false
        validateCarBrand()
    }, 200)
}

// Обработчики для автокомплита года
const onYearInput = (event) => {
    form.carYear = event.target.value
    yearSearchQuery.value = event.target.value
    showYearDropdown.value = true
    clearError('carYear')
}

const selectYear = (year) => {
    form.carYear = year.toString()
    yearSearchQuery.value = year.toString()
    showYearDropdown.value = false
    clearError('carYear')
}

const onYearFocus = () => {
    yearSearchQuery.value = form.carYear
    showYearDropdown.value = true
}

const onYearBlur = () => {
    setTimeout(() => {
        showYearDropdown.value = false
        validateCarYear()
    }, 200)
}

// Обработчики для автокомплита трансмиссии
const onTransInput = (event) => {
    form.carTrans = event.target.value
    transSearchQuery.value = event.target.value
    showTransDropdown.value = true
    clearError('carTrans')
}

const selectTrans = (trans) => {
    form.carTrans = trans
    transSearchQuery.value = trans
    showTransDropdown.value = false
    clearError('carTrans')
}

const onTransFocus = () => {
    transSearchQuery.value = form.carTrans
    showTransDropdown.value = true
}

const onTransBlur = () => {
    setTimeout(() => {
        showTransDropdown.value = false
        validateCarTrans()
    }, 200)
}

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
    } else if (form.name.trim().length > 30) {
        errors.name = 'Ім\'я не може бути довшим за 30 символів'
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
    } else if (!carBrands.includes(form.carBrand.trim())) {
        errors.carBrand = 'Оберіть марку зі списку'
    } else {
        errors.carBrand = null
    }
}

const validateCarModel = () => {
    if (!form.carModel.trim()) {
        errors.carModel = 'Поле "Модель авто" є обов\'язковим'
    } else if (form.carModel.trim().length < 1) {
        errors.carModel = 'Модель авто повинна містити мінімум 1 символ'
    } else if (form.carModel.trim().length > 50) {
        errors.carModel = 'Модель авто не може бути довшою за 50 символів'
    } else {
        errors.carModel = null
    }
}

const validateCarYear = () => {
    if (!form.carYear) {
        errors.carYear = 'Оберіть рік випуску авто'
    } else {
        const year = parseInt(form.carYear)
        if (isNaN(year) || year < 1980 || year > currentYear) {
            errors.carYear = `Рік має бути від 1980 до ${currentYear}`
        } else {
            errors.carYear = null
        }
    }
}

const validateCarTrans = () => {
    if (!form.carTrans) {
        errors.carTrans = 'Оберіть тип трансмісії'
    } else if (!transmissions.includes(form.carTrans.trim())) {
        errors.carTrans = 'Оберіть трансмісію зі списку'
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
    // === ЗАЩИТА ОТ БОТОВ ===

    // 1. Проверка honeypot поля (должно быть пустым)
    if (honeypot.value !== '') {
        console.warn('Bot detected: honeypot filled')
        successMessage.value = 'Помилка відправки. Спробуйте пізніше.'
        setTimeout(() => { successMessage.value = '' }, 3000)
        return
    }

    // 2. Проверка времени заполнения (минимум 3 секунды)
    const fillTime = Date.now() - formStartTime.value
    if (fillTime < 3000) {
        console.warn('Bot detected: form filled too fast')
        successMessage.value = 'Помилка відправки. Заповніть форму повільніше.'
        setTimeout(() => { successMessage.value = '' }, 3000)
        return
    }

    // 3. Проверка, что форма не отправлялась недавно (защита от спама)
    const lastSubmitTime = localStorage.getItem('lastFormSubmit')
    if (lastSubmitTime) {
        const timeSinceLastSubmit = Date.now() - parseInt(lastSubmitTime)
        if (timeSinceLastSubmit < 60000) { // Минимум 1 минута между отправками
            const secondsLeft = Math.ceil((60000 - timeSinceLastSubmit) / 1000)
            successMessage.value = `Зачекайте ${secondsLeft} секунд перед наступною відправкою.`
            setTimeout(() => { successMessage.value = '' }, 3000)
            return
        }
    }

    // === ВАЛИДАЦИЯ ФОРМЫ ===
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
        // Создаем URL-encoded данные
        const params = new URLSearchParams()
        params.append('name', form.name.trim())
        params.append('carBrand', form.carBrand.trim())
        params.append('carModel', form.carModel.trim())
        params.append('carYear', form.carYear.toString())
        params.append('carTrans', form.carTrans)
        params.append('phone', form.phone.trim())
        params.append('description', form.description.trim())

        // Добавляем время заполнения формы для дополнительной проверки на бэкенде
        params.append('_fillTime', fillTime.toString())

        const apiUrl = process.env.NODE_ENV === 'production'
            ? '/api/cars'
            : 'http://localhost:8001/api/cars'

        const response = await fetch(apiUrl, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: params.toString(),
        })

        if (response.ok) {
            const result = await response.json()
            successMessage.value = 'Дякуємо! Ваша заявка успішно відправлена.'

            // Сохраняем время отправки
            localStorage.setItem('lastFormSubmit', Date.now().toString())

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

    // Сбрасываем honeypot и время начала заполнения
    honeypot.value = ''
    formStartTime.value = Date.now()
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
    padding: 20px;
    border-radius: 20px;
    margin-left: auto;
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
    gap: 5px;
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

// Стили для автокомплита
.autocomplete-group {
    position: relative;
}

.autocomplete-dropdown {
    position: absolute;
    top: 100%;
    left: 0;
    right: 0;
    background: #FFFFFF;
    border: 1px solid #E0E0E0;
    border-radius: 12px;
    max-height: 200px;
    overflow-y: auto;
    z-index: 1000;
    margin-top: 4px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);

    /* Кастомный скроллбар */
    &::-webkit-scrollbar {
        width: 6px;
    }

    &::-webkit-scrollbar-track {
        background: #F5F5F5;
        border-radius: 10px;
    }

    &::-webkit-scrollbar-thumb {
        background: #CCCCCC;
        border-radius: 10px;

        &:hover {
            background: #999999;
        }
    }
}

.autocomplete-item {
    padding: 12px 16px;
    cursor: pointer;
    transition: background-color 0.2s ease;
    font-family: Work Sans;
    font-size: 16px;
    color: #000;

    &:hover {
        background-color: #F5F5F5;
    }

    &:first-child {
        border-radius: 12px 12px 0 0;
    }

    &:last-child {
        border-radius: 0 0 12px 12px;
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

    .input-group {
        padding: 10px;

        & input {
            font-size: 12px;
        }
    }
    .form-inner {
        padding: 0 16px;
    }
    .submit-btn {
        padding: 10px;
    }
}
</style>