<template>
    <div class="form-wrapper">
        <div class="form-container-large">
            <div class="form-header">
                <h1 class="form-title-main">Продайте своє авто вигідно</h1>
                <p class="form-subtitle">Заповніть форму і отримайте найкращу ціну за 15 хвилин</p>
            </div>

            <form @submit.prevent="submitForm" novalidate class="form-large">
                <!-- Honeypot -->
                <input type="text" v-model="honeypot" name="website"
                    style="position:absolute;left:-5000px;width:1px;height:1px" tabindex="-1" autocomplete="off"
                    aria-hidden="true">

                <div class="form-grid">
                    <!-- Поле имени -->
                    <div class="input-group-large"
                        :class="{ 'has-error': errors.name, 'has-success': !errors.name && form.name.length > 0 }">
                        <label class="input-label">Ваше ім'я *</label>
                        <input v-model="form.name" type="text" placeholder="Введіть ваше ім'я" required
                            class="input-field" @blur="validateName" @input="clearError('name')" maxlength="30">
                        <div v-if="errors.name" class="error-text">{{ errors.name }}</div>
                    </div>

                    <!-- Поле телефона -->
                    <div class="input-group-large"
                        :class="{ 'has-error': errors.phone, 'has-success': !errors.phone && form.phone.length > 0 }">
                        <label class="input-label">Номер телефону *</label>
                        <input v-model="form.phone" type="tel" placeholder="+380 (XX) XXX-XX-XX" required
                            class="input-field" @input="formatPhone" @blur="validatePhone" maxlength="19">
                        <div v-if="errors.phone" class="error-text">{{ errors.phone }}</div>
                    </div>

                    <!-- Поле марки -->
                    <div class="input-group-large autocomplete-wrapper"
                        :class="{ 'has-error': errors.carBrand, 'has-success': !errors.carBrand && form.carBrand.length > 0 }">
                        <label class="input-label">Марка автомобіля *</label>
                        <input v-model="form.carBrand" type="text" placeholder="Оберіть марку" required
                            class="input-field" @input="onBrandInput" @focus="onBrandFocus" @blur="onBrandBlur"
                            autocomplete="off">

                        <div v-if="showBrandDropdown && filteredBrands.length > 0" class="dropdown-list">
                            <div v-for="brand in filteredBrands" :key="brand" class="dropdown-item"
                                @mousedown="selectBrand(brand)">
                                {{ brand }}
                            </div>
                        </div>

                        <div v-if="errors.carBrand" class="error-text">{{ errors.carBrand }}</div>
                    </div>

                    <!-- Поле модели -->
                    <div class="input-group-large"
                        :class="{ 'has-error': errors.carModel, 'has-success': !errors.carModel && form.carModel.length > 0 }">
                        <label class="input-label">Модель автомобіля *</label>
                        <input v-model="form.carModel" type="text" placeholder="Введіть модель" required
                            class="input-field" @blur="validateCarModel" @input="clearError('carModel')" maxlength="50">
                        <div v-if="errors.carModel" class="error-text">{{ errors.carModel }}</div>
                    </div>

                    <!-- Поле года -->
                    <div class="input-group-large autocomplete-wrapper"
                        :class="{ 'has-error': errors.carYear, 'has-success': !errors.carYear && form.carYear }">
                        <label class="input-label">Рік випуску *</label>
                        <input v-model="form.carYear" type="text" placeholder="Оберіть рік" required class="input-field"
                            @input="onYearInput" @focus="onYearFocus" @blur="onYearBlur" autocomplete="off"
                            maxlength="4">

                        <div v-if="showYearDropdown && filteredYears.length > 0" class="dropdown-list">
                            <div v-for="year in filteredYears.slice(0, 10)" :key="year" class="dropdown-item"
                                @mousedown="selectYear(year)">
                                {{ year }}
                            </div>
                        </div>

                        <div v-if="errors.carYear" class="error-text">{{ errors.carYear }}</div>
                    </div>

                    <!-- Поле трансмиссии -->
                    <div class="input-group-large autocomplete-wrapper"
                        :class="{ 'has-error': errors.carTrans, 'has-success': !errors.carTrans && form.carTrans }">
                        <label class="input-label">Трансмісія *</label>
                        <input v-model="form.carTrans" type="text" placeholder="Оберіть трансмісію" required
                            class="input-field" @input="onTransInput" @focus="onTransFocus" @blur="onTransBlur"
                            autocomplete="off">

                        <div v-if="showTransDropdown && filteredTransmissions.length > 0" class="dropdown-list">
                            <div v-for="trans in filteredTransmissions" :key="trans" class="dropdown-item"
                                @mousedown="selectTrans(trans)">
                                {{ trans }}
                            </div>
                        </div>

                        <div v-if="errors.carTrans" class="error-text">{{ errors.carTrans }}</div>
                    </div>
                </div>

                <!-- Кнопка отправки -->
                <button type="submit" class="submit-button-large" :class="{ 'loading': isSubmitting }"
                    :disabled="isSubmitting || !isFormValid">
                    <span v-if="!isSubmitting">Отримати оцінку авто</span>
                    <span v-else class="spinner"></span>
                </button>

                <!-- Сообщение об успехе -->
                <div v-if="successMessage" class="success-box">
                    {{ successMessage }}
                </div>

                <p class="form-note">* Обов'язкові поля для заповнення</p>
            </form>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'

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

const currentYear = new Date().getFullYear()
const years = ref([])
for (let year = currentYear; year >= 1980; year--) {
    years.value.push(year)
}

const transmissions = [
    'Механічна',
    'Автоматична',
    'Робот',
    'Варіатор',
    'Tiptronic'
]

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
const honeypot = ref('')
const formStartTime = ref(Date.now())

const showBrandDropdown = ref(false)
const brandSearchQuery = ref('')
const showYearDropdown = ref(false)
const yearSearchQuery = ref('')
const showTransDropdown = ref(false)
const transSearchQuery = ref('')

const filteredBrands = computed(() => {
    if (!brandSearchQuery.value) return carBrands
    const query = brandSearchQuery.value.toLowerCase()
    return carBrands.filter(brand => brand.toLowerCase().includes(query))
})

const filteredYears = computed(() => {
    if (!yearSearchQuery.value) return years.value
    const query = yearSearchQuery.value
    return years.value.filter(year => year.toString().includes(query))
})

const filteredTransmissions = computed(() => {
    if (!transSearchQuery.value) return transmissions
    const query = transSearchQuery.value.toLowerCase()
    return transmissions.filter(trans => trans.toLowerCase().includes(query))
})

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

const isValidUkrainianPhone = (phone) => {
    const phoneRegex = /^\+380 \(\d{2}\) \d{3}-\d{2}-\d{2}$/
    return phoneRegex.test(phone)
}

const clearError = (field) => {
    errors[field] = null
}

const submitForm = async () => {
    if (honeypot.value !== '') {
        console.warn('Bot detected: honeypot filled')
        successMessage.value = 'Помилка відправки. Спробуйте пізніше.'
        setTimeout(() => { successMessage.value = '' }, 3000)
        return
    }

    const fillTime = Date.now() - formStartTime.value
    if (fillTime < 3000) {
        console.warn('Bot detected: form filled too fast')
        successMessage.value = 'Помилка відправки. Заповніть форму повільніше.'
        setTimeout(() => { successMessage.value = '' }, 3000)
        return
    }

    const lastSubmitTime = localStorage.getItem('lastFormSubmit')
    if (lastSubmitTime) {
        const timeSinceLastSubmit = Date.now() - parseInt(lastSubmitTime)
        if (timeSinceLastSubmit < 60000) {
            const secondsLeft = Math.ceil((60000 - timeSinceLastSubmit) / 1000)
            successMessage.value = `Зачекайте ${secondsLeft} секунд перед наступною відправкою.`
            setTimeout(() => { successMessage.value = '' }, 3000)
            return
        }
    }

    validateName()
    validateCarBrand()
    validateCarModel()
    validateCarYear()
    validateCarTrans()
    validatePhone()

    if (!isFormValid.value) return

    isSubmitting.value = true
    successMessage.value = ''

    try {
        const params = new URLSearchParams()
        params.append('name', form.name.trim())
        params.append('carBrand', form.carBrand.trim())
        params.append('carModel', form.carModel.trim())
        params.append('carYear', form.carYear.toString())
        params.append('carTrans', form.carTrans)
        params.append('phone', form.phone.trim())
        params.append('description', form.description.trim())
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
            localStorage.setItem('lastFormSubmit', Date.now().toString())
            resetForm()

            setTimeout(() => {
                const successEl = document.querySelector('.success-box')
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

    honeypot.value = ''
    formStartTime.value = Date.now()
}
</script>

<style lang="scss" scoped>
.form-wrapper {
    width: 100%;
    display: flex;
    justify-content: center;
    padding: 60px 20px;
    padding-top: 200px;
    background: #aa3535;
}

.form-container-large {
    width: 100%;
    max-width: 900px;
    background: #ffffff;
    border-radius: 24px;
    padding: 50px 60px;
    box-shadow: 0 20px 60px rgba(0, 0, 0, 0.15);
}

.form-header {
    text-align: center;
    margin-bottom: 50px;
}

.form-title-main {
    font-family: 'Work Sans', sans-serif;
    font-size: 42px;
    font-weight: 700;
    color: #1a1a1a;
    margin-bottom: 16px;
    line-height: 1.2;
}

.form-subtitle {
    font-family: 'Work Sans', sans-serif;
    font-size: 20px;
    font-weight: 400;
    color: #666;
    line-height: 1.5;
}

.form-large {
    display: flex;
    flex-direction: column;
    gap: 30px;
}

.form-grid {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 30px;
}

.input-group-large {
    position: relative;
    display: flex;
    flex-direction: column;
    gap: 10px;

    &.autocomplete-wrapper {
        position: relative;
    }

    &.has-error {
        .input-field {
            border-color: #e74c3c;
            background-color: #fff5f5;
        }
    }

    &.has-success {
        .input-field {
            border-color: #27ae60;
            background-color: #f0fff4;
        }
    }
}

.input-label {
    font-family: 'Work Sans', sans-serif;
    font-size: 16px;
    font-weight: 600;
    color: #333;
    margin-bottom: 4px;
}

.input-field {
    font-family: 'Work Sans', sans-serif;
    font-size: 18px;
    font-weight: 400;
    color: #1a1a1a;
    background: #fafafa;
    border: 2px solid #e0e0e0;
    border-radius: 12px;
    padding: 18px 20px;
    transition: all 0.3s ease;

    &::placeholder {
        color: #999;
    }

    &:focus {
        outline: none;
        border-color: #667eea;
        background: #ffffff;
        box-shadow: 0 0 0 4px rgba(102, 126, 234, 0.1);
    }
}

.dropdown-list {
    position: absolute;
    top: calc(100% + 4px);
    left: 0;
    right: 0;
    background: #ffffff;
    border: 2px solid #e0e0e0;
    border-radius: 12px;
    max-height: 280px;
    overflow-y: auto;
    z-index: 1000;
    box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);

    &::-webkit-scrollbar {
        width: 8px;
    }

    &::-webkit-scrollbar-track {
        background: #f5f5f5;
        border-radius: 10px;
    }

    &::-webkit-scrollbar-thumb {
        background: #ccc;
        border-radius: 10px;

        &:hover {
            background: #999;
        }
    }
}

.dropdown-item {
    padding: 16px 20px;
    cursor: pointer;
    font-family: 'Work Sans', sans-serif;
    font-size: 17px;
    color: #1a1a1a;
    transition: background-color 0.2s ease;

    &:hover {
        background-color: #f5f5f5;
    }

    &:first-child {
        border-radius: 10px 10px 0 0;
    }

    &:last-child {
        border-radius: 0 0 10px 10px;
    }
}

.error-text {
    font-family: 'Work Sans', sans-serif;
    color: #e74c3c;
    font-size: 14px;
    margin-top: 6px;
    font-weight: 500;
}

.submit-button-large {
    width: 100%;
    padding: 22px;
    background: #aa3535;
    border: none;
    border-radius: 14px;
    color: #ffffff;
    font-family: 'Work Sans', sans-serif;
    font-size: 20px;
    font-weight: 700;
    cursor: pointer;
    transition: all 0.3s ease;
    position: relative;
    overflow: hidden;
    margin-top: 20px;

    &:hover:not(:disabled) {
        transform: translateY(-3px);
        box-shadow: 0 12px 30px rgba(102, 126, 234, 0.4);
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

.spinner {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 24px;
    height: 24px;
    border: 3px solid rgba(255, 255, 255, 0.3);
    border-radius: 50%;
    border-top-color: #ffffff;
    animation: spin 1s ease-in-out infinite;
}

@keyframes spin {
    to {
        transform: translate(-50%, -50%) rotate(360deg);
    }
}

.success-box {
    background: linear-gradient(135deg, #27ae60 0%, #2ecc71 100%);
    color: #ffffff;
    padding: 20px;
    border-radius: 14px;
    text-align: center;
    font-family: 'Work Sans', sans-serif;
    font-size: 18px;
    font-weight: 600;
    box-shadow: 0 8px 24px rgba(39, 174, 96, 0.3);
}

.form-note {
    text-align: center;
    font-family: 'Work Sans', sans-serif;
    font-size: 14px;
    color: #999;
    margin-top: -10px;
}

@media (max-width: 768px) {
    .form-wrapper {
        padding: 30px 16px;
    }

    .form-container-large {
        padding: 30px 24px;
    }

    .form-title-main {
        font-size: 32px;
    }

    .form-subtitle {
        font-size: 16px;
    }

    .form-grid {
        grid-template-columns: 1fr;
        gap: 20px;
    }

    .input-field {
        font-size: 16px;
        padding: 14px 16px;
    }

    .submit-button-large {
        font-size: 18px;
        padding: 16px;
    }
}

@media (max-width: 480px) {
    .form-title-main {
        font-size: 28px;
    }

    .form-subtitle {
        font-size: 15px;
    }
}
</style>