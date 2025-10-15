<template>
    <div class="add-car-container">
        <div class="form-wrapper">
            <h2>{{ isEditMode ? 'Редагувати автомобіль' : 'Додати автомобіль на продаж' }}</h2>

            <form @submit.prevent="submitForm" enctype="multipart/form-data">
                <!-- Марка та модель -->
                <div class="form-row">
                    <div class="form-group autocomplete-wrapper">
                        <label for="brand">Марка *</label>
                        <input 
                            v-model="form.brand" 
                            type="text" 
                            id="brand" 
                            required
                            placeholder="Почніть вводити марку..."
                            :class="{ error: errors.brand }"
                            @input="onBrandInput"
                            @focus="onBrandFocus"
                            @blur="onBrandBlur"
                            autocomplete="off">
                        
                        <!-- Dropdown зі списком марок -->
                        <div v-if="showBrandDropdown && filteredBrands.length > 0" class="dropdown-list">
                            <div v-for="brand in filteredBrands" 
                                 :key="brand" 
                                 class="dropdown-item"
                                 @mousedown="selectBrand(brand)">
                                {{ brand }}
                            </div>
                        </div>
                        
                        <span v-if="errors.brand" class="error-text">{{ errors.brand }}</span>
                    </div>

                    <div class="form-group">
                        <label for="model">Модель *</label>
                        <input v-model="form.model" type="text" id="model" required placeholder="X5, E-Class, Camry..."
                            :class="{ error: errors.model }">
                        <span v-if="errors.model" class="error-text">{{ errors.model }}</span>
                    </div>
                </div>

                <!-- Рік та колір -->
                <div class="form-row">
                    <div class="form-group">
                        <label for="year">Рік випуску *</label>
                        <input v-model.number="form.year" type="number" id="year" required :min="1950"
                            :max="new Date().getFullYear() + 1" placeholder="2020" :class="{ error: errors.year }">
                        <span v-if="errors.year" class="error-text">{{ errors.year }}</span>
                    </div>

                    <div class="form-group">
                        <label for="color">Колір</label>
                        <input v-model="form.color" type="text" id="color" placeholder="Чорний, Білий, Сріблястий..."
                            :class="{ error: errors.color }">
                        <span v-if="errors.color" class="error-text">{{ errors.color }}</span>
                    </div>
                </div>

                <!-- Паливо та трансмісія -->
                <div class="form-row">
                    <div class="form-group">
                        <label for="fuel">Паливо *</label>
                        <select v-model="form.fuel" id="fuel" required :class="{ error: errors.fuel }">
                            <option value="">Оберіть тип палива</option>
                            <option value="Бензин">Бензин</option>
                            <option value="Дизель">Дизель</option>
                            <option value="Гібрид">Гібрид</option>
                            <option value="Електро">Електро</option>
                            <option value="Газ">Газ</option>
                            <option value="Газ/Бензин">Газ/Бензин</option>
                        </select>
                        <span v-if="errors.fuel" class="error-text">{{ errors.fuel }}</span>
                    </div>

                    <div class="form-group">
                        <label for="transmission">Трансмісія *</label>
                        <select v-model="form.transmission" id="transmission" required
                            :class="{ error: errors.transmission }">
                            <option value="">Оберіть тип трансмісії</option>
                            <option value="Механічна">Механічна</option>
                            <option value="Автоматична">Автоматична</option>
                            <option value="Робот">Робот</option>
                            <option value="Варіатор">Варіатор</option>
                        </select>
                        <span v-if="errors.transmission" class="error-text">{{ errors.transmission }}</span>
                    </div>
                </div>

                <!-- Привід та пробіг -->
                <div class="form-row">
                    <div class="form-group">
                        <label for="drive">Привід *</label>
                        <select v-model="form.drive" id="drive" required :class="{ error: errors.drive }">
                            <option value="">Оберіть тип приводу</option>
                            <option value="Передній">Передній</option>
                            <option value="Задній">Задній</option>
                            <option value="Повний">Повний</option>
                        </select>
                        <span v-if="errors.drive" class="error-text">{{ errors.drive }}</span>
                    </div>

                    <div class="form-group">
                        <label for="mileage">Пробіг (км) *</label>
                        <input v-model.number="form.mileage" type="number" id="mileage" required min="0"
                            placeholder="100000" :class="{ error: errors.mileage }">
                        <span v-if="errors.mileage" class="error-text">{{ errors.mileage }}</span>
                    </div>
                </div>

                <!-- Статус -->
                <div class="form-group">
                    <label for="status">Статус</label>
                    <select v-model="form.status" id="status" :class="{ error: errors.status }">
                        <option value="active">Активний</option>
                        <option value="sold">Продано</option>
                        <option value="new">Новинка</option>
                        <option value="sale">Зі знижкою</option>
                        <option value="super-price">Супер ціна</option>
                    </select>
                    <span v-if="errors.status" class="error-text">{{ errors.status }}</span>
                </div>

                <!-- Опис -->
                <div class="form-group">
                    <label for="description">Опис</label>
                    <textarea v-model="form.description" id="description" rows="4"
                        placeholder="Додаткова інформація про автомобіль..."
                        :class="{ error: errors.description }"></textarea>
                    <span v-if="errors.description" class="error-text">{{ errors.description }}</span>
                </div>

                <!-- Існуючі фотографії (в режимі редагування) -->
                <div v-if="isEditMode && existingImages.length > 0" class="form-group">
                    <label>Поточні фотографії</label>
                    <div class="photos-preview">
                        <div v-for="img in existingImages" :key="img.id" class="photo-preview">
                            <img :src="img.fileUrl" :alt="img.fileName">
                            <button type="button" @click="markImageForDeletion(img.id)" 
                                    :class="['remove-photo', { 'marked-delete': imagesToDelete.includes(img.id) }]">
                                {{ imagesToDelete.includes(img.id) ? '↶' : '×' }}
                            </button>
                        </div>
                    </div>
                </div>

                <!-- Нові фотографії -->
                <div class="form-group">
                    <label for="photos">{{ isEditMode ? 'Додати нові фотографії' : 'Фотографії' }} (до 10 штук)</label>
                    <input ref="fileInput" type="file" id="photos" accept="image/*" multiple @change="handleFileUpload"
                        :class="{ error: errors.photos }">
                    <span v-if="errors.photos" class="error-text">{{ errors.photos }}</span>

                    <div class="file-info">
                        Максимальний розмір файлу: 5MB. Підтримувані формати: JPG, PNG, WebP
                    </div>

                    <!-- Превью нових фотографій -->
                    <div v-if="photoPreviews.length > 0" class="photos-preview">
                        <div v-for="(preview, index) in photoPreviews" :key="`new-${index}`" class="photo-preview">
                            <img :src="preview" :alt="`Нове фото ${index + 1}`">
                            <button type="button" @click="removePhoto(index)" class="remove-photo">×</button>
                            <div class="photo-number">{{ index + 1 }}</div>
                        </div>
                    </div>
                </div>

                <!-- Кнопки -->
                <div class="form-actions">
                    <button v-if="isEditMode" type="button" @click="cancelEdit" class="btn-secondary" 
                            :disabled="isSubmitting">
                        Скасувати
                    </button>
                    <button v-else type="button" @click="resetForm" class="btn-secondary" :disabled="isSubmitting">
                        Очистити
                    </button>
                    <button type="submit" class="btn-primary" :disabled="isSubmitting">
                        <span v-if="!isSubmitting">{{ isEditMode ? 'Зберегти зміни' : 'Додати автомобіль' }}</span>
                        <span v-else>Збереження...</span>
                    </button>
                </div>

                <!-- Повідомлення -->
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
import { ref, reactive, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

// Список марок автомобілів
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

// Режим редагування
const isEditMode = computed(() => !!route.params.id)
const carId = computed(() => route.params.id)

// Автокомпліт для марок
const showBrandDropdown = ref(false)
const brandSearchQuery = ref('')

const filteredBrands = computed(() => {
    if (!brandSearchQuery.value) return carBrands
    const query = brandSearchQuery.value.toLowerCase()
    return carBrands.filter(brand => brand.toLowerCase().includes(query))
})

// Реактивні дані
const form = reactive({
    brand: '',
    model: '',
    year: null,
    color: '',
    fuel: '',
    transmission: '',
    drive: '',
    mileage: null,
    status: 'active',
    description: '',
    photos: []
})

const errors = reactive({
    brand: null,
    model: null,
    year: null,
    color: null,
    fuel: null,
    transmission: null,
    drive: null,
    mileage: null,
    status: null,
    description: null,
    photos: null
})

const isSubmitting = ref(false)
const successMessage = ref('')
const errorMessage = ref('')
const photoPreviews = ref([])
const fileInput = ref(null)
const existingImages = ref([])
const imagesToDelete = ref([])

// Функції для автокомпліту марок
const onBrandInput = (event) => {
    form.brand = event.target.value
    brandSearchQuery.value = event.target.value
    showBrandDropdown.value = true
    errors.brand = null
}

const selectBrand = (brand) => {
    form.brand = brand
    brandSearchQuery.value = brand
    showBrandDropdown.value = false
    errors.brand = null
}

const onBrandFocus = () => {
    brandSearchQuery.value = form.brand
    showBrandDropdown.value = true
}

const onBrandBlur = () => {
    setTimeout(() => {
        showBrandDropdown.value = false
    }, 200)
}

// Завантаження даних автомобіля при редагуванні
onMounted(async () => {
    if (isEditMode.value) {
        await loadCarData()
    }
})

const loadCarData = async () => {
    try {
        const apiUrl = process.env.NODE_ENV === 'production'
            ? '/api/cars-sale'
            : 'http://localhost:8001/api/cars-sale'

        const response = await fetch(`${apiUrl}/${carId.value}`)
        
        if (!response.ok) {
            throw new Error('Не вдалося завантажити дані автомобіля')
        }

        const result = await response.json()
        const carData = result.data

        // Заповнюємо форму
        Object.assign(form, {
            brand: carData.brand,
            model: carData.model,
            year: carData.year,
            color: carData.color || '',
            fuel: carData.fuel,
            transmission: carData.transmission,
            drive: carData.drive,
            mileage: carData.mileage,
            status: carData.status || 'active',
            description: carData.description || '',
            photos: []
        })

        // Зберігаємо існуючі зображення
        existingImages.value = carData.images || []

    } catch (error) {
        console.error('Error loading car data:', error)
        errorMessage.value = 'Помилка завантаження даних автомобіля'
    }
}

// Відмітити зображення для видалення
const markImageForDeletion = (imageId) => {
    const index = imagesToDelete.value.indexOf(imageId)
    if (index > -1) {
        imagesToDelete.value.splice(index, 1)
    } else {
        imagesToDelete.value.push(imageId)
    }
}

// Валідація
const validateForm = () => {
    Object.keys(errors).forEach(key => errors[key] = null)

    let isValid = true

    if (!form.brand.trim()) {
        errors.brand = 'Марка обов\'язкова'
        isValid = false
    } else if (!carBrands.includes(form.brand.trim())) {
        errors.brand = 'Оберіть марку зі списку'
        isValid = false
    }

    if (!form.model.trim()) {
        errors.model = 'Модель обов\'язкова'
        isValid = false
    }

    if (!form.year || form.year < 1950 || form.year > new Date().getFullYear() + 1) {
        errors.year = 'Введіть коректний рік'
        isValid = false
    }

    if (!form.fuel) {
        errors.fuel = 'Оберіть тип палива'
        isValid = false
    }

    if (!form.transmission) {
        errors.transmission = 'Оберіть тип трансмісії'
        isValid = false
    }

    if (!form.drive) {
        errors.drive = 'Оберіть тип приводу'
        isValid = false
    }

    if (form.mileage === null || form.mileage < 0) {
        errors.mileage = 'Введіть коректний пробіг'
        isValid = false
    }

    return isValid
}

// Обробка завантаження файлів
const handleFileUpload = (event) => {
    const files = Array.from(event.target.files)

    if (files.length === 0) return

    const totalImages = existingImages.value.length - imagesToDelete.value.length + files.length
    
    if (totalImages > 10) {
        errors.photos = 'Максимум 10 фотографій всього'
        if (fileInput.value) fileInput.value.value = ''
        return
    }

    const maxSize = 5 * 1024 * 1024 // 5MB
    const allowedTypes = ['image/jpeg', 'image/jpg', 'image/png', 'image/webp', 'image/gif']

    for (const file of files) {
        if (!allowedTypes.includes(file.type.toLowerCase())) {
            errors.photos = 'Дозволені тільки зображення (JPG, PNG, WebP, GIF)'
            if (fileInput.value) fileInput.value.value = ''
            return
        }

        if (file.size > maxSize) {
            errors.photos = `Файл "${file.name}" занадто великий. Максимум 5MB`
            if (fileInput.value) fileInput.value.value = ''
            return
        }
    }

    form.photos = files
    errors.photos = null

    photoPreviews.value = []
    files.forEach(file => {
        const reader = new FileReader()
        reader.onload = (e) => {
            photoPreviews.value.push(e.target.result)
        }
        reader.readAsDataURL(file)
    })
}

// Видалення нового фото
const removePhoto = (index) => {
    const newFiles = Array.from(form.photos).filter((_, i) => i !== index)
    form.photos = newFiles
    photoPreviews.value.splice(index, 1)

    if (fileInput.value) {
        const dt = new DataTransfer()
        newFiles.forEach(file => dt.items.add(file))
        fileInput.value.files = dt.files
    }
}

// Відправка форми
const submitForm = async () => {
    if (!validateForm()) return

    isSubmitting.value = true
    successMessage.value = ''
    errorMessage.value = ''

    try {
        const formData = new FormData()

        formData.append('brand', form.brand.trim())
        formData.append('model', form.model.trim())
        formData.append('year', form.year.toString())
        formData.append('color', form.color.trim())
        formData.append('fuel', form.fuel)
        formData.append('transmission', form.transmission)
        formData.append('drive', form.drive)
        formData.append('mileage', form.mileage.toString())
        formData.append('status', form.status)
        formData.append('description', form.description.trim())

        // Додаємо ID зображень для видалення
        if (isEditMode.value && imagesToDelete.value.length > 0) {
            formData.append('deleteImages', imagesToDelete.value.join(','))
        }

        // Додаємо нові фотографії
        form.photos.forEach(photo => {
            formData.append('images', photo)
        })

        const apiUrl = process.env.NODE_ENV === 'production'
            ? '/api/cars-sale'
            : 'http://localhost:8001/api/cars-sale'

        const endpoint = isEditMode.value ? `/${carId.value}` : ''

        const response = await fetch(`${apiUrl}${endpoint}`, {
            method: isEditMode.value ? 'PUT' : 'POST',
            headers: {
                'Authorization': `Bearer ${authStore.token}`
            },
            body: formData
        })

        if (response.ok) {
            const result = await response.json()
            successMessage.value = isEditMode.value 
                ? 'Автомобіль успішно оновлено!' 
                : 'Автомобіль успішно додано в каталог!'
            
            if (!isEditMode.value) {
                resetForm()
            } else {
                setTimeout(() => {
                    router.push('/admin/list')
                }, 1500)
            }
        } else {
            const errorData = await response.json().catch(() => ({ error: 'Unknown error' }))
            throw new Error(errorData.error || `HTTP ${response.status}`)
        }

    } catch (error) {
        console.error('Submission error:', error)
        errorMessage.value = 'Помилка збереження. Спробуйте ще раз.'
    } finally {
        isSubmitting.value = false
    }
}

// Скасування редагування
const cancelEdit = () => {
    router.push('/admin/list')
}

// Скидання форми
const resetForm = () => {
    Object.assign(form, {
        brand: '',
        model: '',
        year: null,
        color: '',
        fuel: '',
        transmission: '',
        drive: '',
        mileage: null,
        status: 'active',
        description: '',
        photos: []
    })

    Object.keys(errors).forEach(key => errors[key] = null)
    photoPreviews.value = []
    existingImages.value = []
    imagesToDelete.value = []
    successMessage.value = ''
    errorMessage.value = ''
    brandSearchQuery.value = ''
    showBrandDropdown.value = false

    if (fileInput.value) {
        fileInput.value.value = ''
    }
}
</script>

<style scoped>
.add-car-container {
    max-width: 1440px;
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
    position: relative;
}

.autocomplete-wrapper {
    position: relative;
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
    box-sizing: border-box;
    color: #000;
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

/* Dropdown для автокомпліту */
.dropdown-list {
    position: absolute;
    top: calc(100% + 4px);
    left: 0;
    right: 0;
    background: #ffffff;
    border: 2px solid #e0e0e0;
    border-radius: 8px;
    max-height: 250px;
    overflow-y: auto;
    z-index: 1000;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.dropdown-list::-webkit-scrollbar {
    width: 8px;
}

.dropdown-list::-webkit-scrollbar-track {
    background: #f5f5f5;
    border-radius: 4px;
}

.dropdown-list::-webkit-scrollbar-thumb {
    background: #ccc;
    border-radius: 4px;
}

.dropdown-list::-webkit-scrollbar-thumb:hover {
    background: #999;
}

.dropdown-item {
    padding: 12px 16px;
    cursor: pointer;
    font-size: 1rem;
    color: #333;
    transition: background-color 0.2s ease;
}

.dropdown-item:hover {
    background-color: #f0f0f0;
}

.dropdown-item:first-child {
    border-radius: 6px 6px 0 0;
}

.dropdown-item:last-child {
    border-radius: 0 0 6px 6px;
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
    transition: background 0.3s ease;
}

.remove-photo:hover {
    background: #dc3545;
}

.remove-photo.marked-delete {
    background: rgba(255, 193, 7, 0.9);
}

.remove-photo.marked-delete:hover {
    background: #ffc107;
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