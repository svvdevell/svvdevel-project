<template>
    <div class="add-car-container">
        <div class="form-wrapper">
            <h2>{{ isEditMode ? 'Редагувати автомобіль' : 'Додати автомобіль на продаж' }}</h2>

            <form @submit.prevent="submitForm" enctype="multipart/form-data">
                <!-- Марка та модель -->
                <div class="form-row">
                    <div class="form-group autocomplete-wrapper">
                        <label for="brand">Марка *</label>
                        <input v-model="form.brand" type="text" id="brand" required
                            placeholder="Почніть вводити марку..." :class="{ error: errors.brand }"
                            @input="onBrandInput" @focus="onBrandFocus" @blur="onBrandBlur" autocomplete="off">

                        <!-- Dropdown зі списком марок -->
                        <div v-if="showBrandDropdown && filteredBrands.length > 0" class="dropdown-list">
                            <div v-for="brand in filteredBrands" :key="brand" class="dropdown-item"
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
                        <label for="volume">Об'єм двигуна *</label>
                        <input v-model.number="form.volume" type="number" id="volume" required min="0" placeholder="3.0"
                            :class="{ error: errors.volume }">
                        <span v-if="errors.volume" class="error-text">{{ errors.volume }}</span>
                    </div>

                    <div class="form-group">
                        <label for="mileage">Пробіг (км) *</label>
                        <input v-model.number="form.mileage" type="number" id="mileage" required min="0"
                            placeholder="100000" :class="{ error: errors.mileage }">
                        <span v-if="errors.mileage" class="error-text">{{ errors.mileage }}</span>
                    </div>

                    <div class="form-group">
                        <label for="price">Ціна ($) *</label>
                        <input v-model.number="form.price" type="number" id="price" required min="3"
                            placeholder="10000$" :class="{ error: errors.price }">
                        <span v-if="errors.price" class="error-text">{{ errors.price }}</span>
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
                        :class="{ error: errors.photos }" :disabled="isProcessingImages">
                    <span v-if="errors.photos" class="error-text">{{ errors.photos }}</span>

                    <div class="file-info">
                        Максимальний розмір файлу: 10MB. Підтримувані формати: JPG, PNG, WebP. Фото будуть автоматично
                        стиснуті.
                    </div>

                    <!-- Статус обробки зображень -->
                    <div v-if="isProcessingImages" class="processing-status">
                        <span>Обробка зображень...</span>
                    </div>

                    <!-- Превью нових фотографій -->
                    <div v-if="photoPreviews.length > 0" class="photos-preview">
                        <div v-for="(preview, index) in photoPreviews" :key="`new-${index}`" class="photo-preview">
                            <img :src="preview" :alt="`Нове фото ${index + 1}`">
                            <button type="button" @click="removePhoto(index)" class="remove-photo">×</button>
                            <div class="photo-number">{{ index + 1 }}</div>
                            <div v-if="form.photos[index]" class="photo-size">{{ formatFileSize(form.photos[index].size)
                                }}</div>
                        </div>
                    </div>
                </div>

                <!-- Кнопки -->
                <div class="form-actions">
                    <button v-if="isEditMode" type="button" @click="cancelEdit" class="btn-secondary"
                        :disabled="isSubmitting || isProcessingImages">
                        Скасувати
                    </button>
                    <button v-else type="button" @click="resetForm" class="btn-secondary"
                        :disabled="isSubmitting || isProcessingImages">
                        Очистити
                    </button>
                    <button type="submit" class="btn-primary" :disabled="isSubmitting || isProcessingImages">
                        <span v-if="!isSubmitting">{{ isEditMode ? 'Зберегти зміни' : 'Додати автомобіль' }}</span>
                        <span v-else>{{ uploadProgress > 0 ? `Збереження... ${uploadProgress}%` : 'Збереження...'
                            }}</span>
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
]

// Змінні для автокомпліту
const showBrandDropdown = ref(false)
const filteredBrands = ref([])

// Змінні для відображення помилок і повідомлень
const errors = reactive({})
const successMessage = ref('')
const errorMessage = ref('')

// Змінні для обробки зображень
const fileInput = ref(null)
const photoPreviews = ref([])
const isSubmitting = ref(false)
const isProcessingImages = ref(false)
const uploadProgress = ref(0)

// Змінні для режиму редагування
const isEditMode = ref(false)
const carIdToEdit = ref(null)
const existingImages = ref([])
const imagesToDelete = ref([])

// Основна форма
const form = reactive({
    brand: '',
    model: '',
    year: new Date().getFullYear(),
    color: '',
    fuel: '',
    transmission: '',
    drive: '',
    volume: '',
    mileage: '',
    price: '',
    status: 'active',
    description: '',
    photos: []
})

// Функция для получения заголовков аутентификации - ИСПРАВЛЕНО!
function getAuthHeaders() {
    // Проверяем наличие разных свойств, которые могут содержать информацию об аутентификации
    if (authStore.token) {
        return {
            'Authorization': `Bearer ${authStore.token}`
        };
    }

    // Если есть метод getHeaders
    if (typeof authStore.getHeaders === 'function') {
        return authStore.getHeaders();
    }

    // Если есть свойство headers
    if (authStore.headers) {
        return authStore.headers;
    }

    // Если у authStore вообще нет никаких методов или свойств для аутентификации,
    // просто возвращаем пустой объект
    return {};
}

// При завантаженні компонента
onMounted(() => {
    // Перевіряємо чи ми в режимі редагування
    if (route.name === 'EditCar' && route.params.id) {
        isEditMode.value = true
        carIdToEdit.value = route.params.id
        loadCarDetails(route.params.id)
    }
})

// Завантаження даних автомобіля для редагування
async function loadCarDetails(carId) {
    try {
        const response = await fetch(`/api/cars-sale/${carId}`, {
            method: 'GET',
            headers: getAuthHeaders() // Исправлено!
        })

        if (!response.ok) {
            throw new Error('Не вдалося завантажити дані автомобіля')
        }

        const data = await response.json()

        if (data && data.car) {
            // Заповнюємо форму
            form.brand = data.car.brand
            form.model = data.car.model
            form.year = data.car.year
            form.color = data.car.color || ''
            form.fuel = data.car.fuel
            form.transmission = data.car.transmission
            form.drive = data.car.drive
            form.volume = data.car.volume
            form.mileage = data.car.mileage
            form.price = data.car.price
            form.status = data.car.status
            form.description = data.car.description || ''

            // Завантажуємо зображення
            if (data.car.images && data.car.images.length > 0) {
                existingImages.value = data.car.images
            }
        }
    } catch (error) {
        console.error('Error loading car details:', error)
        errorMessage.value = 'Помилка завантаження даних автомобіля'
    }
}

// Функції для автокомпліту
function onBrandInput() {
    const query = form.brand.toLowerCase().trim()

    if (query) {
        filteredBrands.value = carBrands.filter(brand =>
            brand.toLowerCase().includes(query)
        ).slice(0, 10) // Показуємо максимум 10 результатів
        showBrandDropdown.value = true
    } else {
        showBrandDropdown.value = false
        filteredBrands.value = []
    }
}

function onBrandFocus() {
    if (form.brand.trim()) {
        onBrandInput()
    }
}

function onBrandBlur() {
    // Затримуємо приховання списку, щоб користувач міг клікнути на варіант
    setTimeout(() => {
        showBrandDropdown.value = false
    }, 200)
}

function selectBrand(brand) {
    form.brand = brand
    showBrandDropdown.value = false
}

// Функція для видалення фото з превью
function removePhoto(index) {
    form.photos.splice(index, 1)
    photoPreviews.value.splice(index, 1)
}

// Функція для позначення фото на видалення (режим редагування)
function markImageForDeletion(imageId) {
    const index = imagesToDelete.value.indexOf(imageId)
    if (index === -1) {
        imagesToDelete.value.push(imageId)
    } else {
        imagesToDelete.value.splice(index, 1)
    }
}

// Скасування редагування
function cancelEdit() {
    router.back()
}

// Очищення форми
function resetForm() {
    Object.keys(form).forEach(key => {
        if (key === 'year') {
            form[key] = new Date().getFullYear()
        } else if (key === 'status') {
            form[key] = 'active'
        } else if (key === 'photos') {
            form[key] = []
        } else {
            form[key] = ''
        }
    })
    photoPreviews.value = []
    if (fileInput.value) {
        fileInput.value.value = ''
    }
    errors.value = {}
    successMessage.value = ''
    errorMessage.value = ''
}

// Валідація форми
function validateForm() {
    errors.brand = !form.brand ? 'Поле обов\'язкове' : ''
    errors.model = !form.model ? 'Поле обов\'язкове' : ''
    errors.year = !form.year ? 'Поле обов\'язкове' : ''
    errors.fuel = !form.fuel ? 'Поле обов\'язкове' : ''
    errors.transmission = !form.transmission ? 'Поле обов\'язкове' : ''
    errors.drive = !form.drive ? 'Поле обов\'язкове' : ''
    errors.volume = !form.volume ? 'Поле обов\'язкове' : ''
    errors.mileage = !form.mileage ? 'Поле обов\'язкове' : ''
    errors.price = !form.price ? 'Поле обов\'язкове' : ''

    // Перевіряємо чи є хоча б одна помилка
    return !Object.values(errors).some(error => error)
}

// Функція для форматування розміру файлу
function formatFileSize(bytes) {
    if (bytes === 0) return '0 Байт'

    const sizes = ['Байт', 'КБ', 'МБ', 'ГБ']
    const i = Math.floor(Math.log(bytes) / Math.log(1024))

    return `${(bytes / Math.pow(1024, i)).toFixed(1)} ${sizes[i]}`
}

/**
 * Функція для стиснення зображення перед завантаженням
 */
async function compressImage(file, maxSizeMB = 1) {
    return new Promise((resolve, reject) => {
        const maxSize = maxSizeMB * 1024 * 1024 // в байтах

        // Якщо файл вже менший за максимальний розмір або не є зображенням - повертаємо як є
        if (file.size <= maxSize || !file.type.startsWith('image/')) {
            console.log(`Файл ${file.name} не потребує стиснення (${(file.size / 1024 / 1024).toFixed(2)}MB)`)
            return resolve(file)
        }

        console.log(`Стискаємо файл ${file.name} (${(file.size / 1024 / 1024).toFixed(2)}MB)`)

        const reader = new FileReader()
        reader.readAsDataURL(file)

        reader.onload = (event) => {
            const img = new Image()
            img.src = event.target.result

            img.onload = () => {
                // Створюємо canvas для стиснення
                const canvas = document.createElement('canvas')

                // Визначаємо розміри
                let width = img.width
                let height = img.height

                // Якщо розміри занадто великі, зменшуємо
                const MAX_WIDTH = 1800
                const MAX_HEIGHT = 1800

                if (width > MAX_WIDTH) {
                    const ratio = width / MAX_WIDTH
                    width = MAX_WIDTH
                    height = Math.round(height / ratio)
                }

                if (height > MAX_HEIGHT) {
                    const ratio = height / MAX_HEIGHT
                    height = MAX_HEIGHT
                    width = Math.round(width / ratio)
                }

                canvas.width = width
                canvas.height = height

                // Малюємо зображення на canvas
                const ctx = canvas.getContext('2d')
                ctx.drawImage(img, 0, 0, width, height)

                // Починаємо з хорошої якості (0.9)
                let quality = 0.9
                const MIN_QUALITY = 0.5 // Не опускаємось нижче цієї якості

                // Функція рекурсивного стиснення зі зменшенням якості
                const compressRecursively = () => {
                    canvas.toBlob(
                        (blob) => {
                            if (blob.size > maxSize && quality > MIN_QUALITY) {
                                // Якщо все ще більший за максимальний розмір - зменшуємо якість
                                quality -= 0.1
                                console.log(`Повторне стиснення, якість: ${quality.toFixed(1)}`)
                                compressRecursively()
                            } else {
                                // Створюємо новий файл із блоба
                                const compressedFile = new File([blob], file.name, {
                                    type: file.type,
                                    lastModified: file.lastModified
                                })

                                console.log(`Файл стиснуто: ${file.size} -> ${compressedFile.size} (${Math.round(compressedFile.size / file.size * 100)}%)`)
                                resolve(compressedFile)
                            }
                        },
                        file.type,
                        quality
                    )
                }

                compressRecursively()
            }

            img.onerror = (error) => {
                console.error('Помилка при завантаженні зображення для стиснення', error)
                resolve(file) // Повертаємо оригінальний файл у разі помилки
            }
        }

        reader.onerror = (error) => {
            console.error('Помилка читання файлу', error)
            resolve(file) // Повертаємо оригінальний файл у разі помилки
        }
    })
}

// Функція для обробки завантажених файлів з оптимізацією
async function handleFileUpload(event) {
    const files = event.target.files
    if (!files || files.length === 0) return

    // Скидаємо помилки
    errors.photos = ''

    // Перевірка кількості файлів
    const maxAllowedFiles = 10 - (isEditMode.value ? existingImages.value.length - imagesToDelete.value.length : 0)
    if (files.length > maxAllowedFiles) {
        errors.photos = `Можна завантажити максимум ${maxAllowedFiles} зображень`
        return
    }

    // Очищаємо поточні файли та превью
    form.photos = []
    photoPreviews.value.splice(0, photoPreviews.value.length)

    // Відображаємо індикатор завантаження
    isProcessingImages.value = true

    try {
        // Обробляємо кожен файл
        for (const file of files) {
            // Перевірка розміру перед стисненням
            if (file.size > 20 * 1024 * 1024) { // 20MB
                errors.photos = `Файл ${file.name} занадто великий (${(file.size / 1024 / 1024).toFixed(2)}MB). Максимальний розмір - 20MB.`
                continue
            }

            // Стискаємо зображення до максимум 2MB
            const compressedFile = await compressImage(file, 2)

            // Додаємо в форму
            form.photos.push(compressedFile)

            // Створюємо превью
            const reader = new FileReader()
            reader.onload = (e) => {
                photoPreviews.value.push(e.target.result)
            }
            reader.readAsDataURL(compressedFile)
        }
    } catch (error) {
        console.error('Помилка при обробці зображень:', error)
        errors.photos = 'Виникла помилка при обробці зображень'
    } finally {
        isProcessingImages.value = false
    }
}

// Функція для відправки форми з оптимізацією завантаження великих файлів
async function submitForm() {
    try {
        // Якщо вже йде відправка - виходимо
        if (isSubmitting.value) return

        // Очищаємо повідомлення
        successMessage.value = ''
        errorMessage.value = ''

        // Встановлюємо прапор відправки
        isSubmitting.value = true
        uploadProgress.value = 0

        // Валідація форми
        const isValid = validateForm()
        if (!isValid) {
            isSubmitting.value = false
            errorMessage.value = 'Будь ласка, заповніть всі обов\'язкові поля.'
            return
        }

        // Стратегія відправки залежить від кількості та розміру файлів
        const totalImagesSize = form.photos.reduce((sum, file) => sum + file.size, 0)

        // Якщо загальний розмір файлів більший за 15MB або файлів більше 5, використовуємо стратегію пакетної відправки
        if (totalImagesSize > 15 * 1024 * 1024 || form.photos.length > 5) {
            console.log(`Використовуємо пакетну стратегію. Загальний розмір: ${(totalImagesSize / 1024 / 1024).toFixed(2)}MB, К-сть файлів: ${form.photos.length}`)

            // Спочатку створюємо/оновлюємо автомобіль без зображень
            const carId = await createOrUpdateCarWithoutImages()

            // Потім завантажуємо зображення пакетами
            await uploadImagesInBatches(carId)
        } else {
            // Звичайна стратегія: відправляємо все одразу
            console.log(`Використовуємо звичайну стратегію. Загальний розмір: ${(totalImagesSize / 1024 / 1024).toFixed(2)}MB, К-сть файлів: ${form.photos.length}`)
            await createOrUpdateCarWithImages()
        }

        // Успішне завершення
        successMessage.value = isEditMode.value ? 'Автомобіль успішно оновлено' : 'Автомобіль успішно додано'

        if (isEditMode.value) {
            setTimeout(() => {
                router.push(`/cars/${carIdToEdit.value}`)
            }, 1500) // Затримка для показу повідомлення
        } else {
            resetForm()
        }
    } catch (error) {
        console.error('Помилка при відправці форми:', error)
        errorMessage.value = `Помилка: ${error.message || 'Не вдалося зберегти дані'}`
    } finally {
        isSubmitting.value = false
        uploadProgress.value = 0
    }
}

// Функція для створення/оновлення автомобіля без зображень
async function createOrUpdateCarWithoutImages() {
    const formData = new FormData()

    // Додаємо всі дані крім зображень
    Object.keys(form).forEach(key => {
        if (key !== 'photos') {
            formData.append(key, form[key])
        }
    })

    // Якщо в режимі редагування, додаємо ID зображень для видалення
    if (isEditMode.value && imagesToDelete.value.length > 0) {
        formData.append('deleteImages', imagesToDelete.value.join(','))
    }

    // Відправляємо запит
    let url = '/api/cars-sale'
    let method = 'POST'

    if (isEditMode.value) {
        url = `/api/cars-sale/${carIdToEdit.value}`
        method = 'PUT'
    }

    const response = await fetch(url, {
        method,
        body: formData,
        headers: getAuthHeaders() // Исправлено!
    })

    if (!response.ok) {
        let errorText = 'Помилка збереження автомобіля'
        try {
            const errorData = await response.json()
            errorText = errorData.error || errorText
        } catch (e) {
            // Ігноруємо помилку парсингу JSON
        }
        throw new Error(errorText)
    }

    const result = await response.json()

    if (result.status !== 'success') {
        throw new Error(result.error || 'Помилка збереження автомобіля')
    }

    // Повертаємо ID створеного/відредагованого автомобіля
    return isEditMode.value ? carIdToEdit.value : result.data.id
}

// Функція для відправки зображень пакетами
async function uploadImagesInBatches(carId) {
    if (!form.photos.length) return

    const id = carId || (isEditMode.value ? carIdToEdit.value : null)
    if (!id) {
        throw new Error('ID автомобіля не знайдено')
    }

    const maxBatchSize = 8 * 1024 * 1024 // 8MB
    const maxBatchFiles = 3 // Максимум 3 файли в пакеті

    // Створюємо пакети зображень
    const batches = []
    let currentBatch = []
    let currentBatchSize = 0

    for (const file of form.photos) {
        // Якщо пакет заповнений (за розміром або кількістю) - починаємо новий
        if (currentBatchSize + file.size > maxBatchSize || currentBatch.length >= maxBatchFiles) {
            if (currentBatch.length > 0) {
                batches.push(currentBatch)
            }
            currentBatch = [file]
            currentBatchSize = file.size
        } else {
            // Додаємо файл в поточний пакет
            currentBatch.push(file)
            currentBatchSize += file.size
        }
    }

    // Додаємо останній пакет, якщо він не порожній
    if (currentBatch.length > 0) {
        batches.push(currentBatch)
    }

    // Відправляємо пакети один за одним
    let successfulBatches = 0

    for (let i = 0; i < batches.length; i++) {
        uploadProgress.value = Math.round((i / batches.length) * 100)
        console.log(`Відправка пакету ${i + 1}/${batches.length}, ${batches[i].length} файлів`)

        const batchFormData = new FormData()

        // Додаємо мінімальні дані для оновлення
        if (isEditMode.value) {
            Object.keys(form).forEach(key => {
                if (key !== 'photos' && key !== 'deleteImages') {
                    batchFormData.append(key, form[key])
                }
            })
        } else {
            // Для нового автомобіля потрібні лише мінімальні дані
            batchFormData.append('brand', form.brand)
            batchFormData.append('model', form.model)
        }

        // Додаємо зображення поточного пакету
        batches[i].forEach(file => {
            batchFormData.append('images', file)
        })

        try {
            // Відправляємо запит на оновлення
            const batchResponse = await fetch(`/api/cars-sale/${id}`, {
                method: 'PUT',
                body: batchFormData,
                headers: getAuthHeaders() // Исправлено!
            })

            if (batchResponse.ok) {
                successfulBatches++
            } else {
                console.error(`Помилка при завантаженні пакету зображень ${i + 1}/${batches.length}`)
            }
        } catch (error) {
            console.error(`Помилка при відправці пакету ${i + 1}:`, error)
        }

        // Невелика затримка між запитами
        await new Promise(resolve => setTimeout(resolve, 300))
    }

    uploadProgress.value = 100

    if (successfulBatches === 0) {
        throw new Error('Не вдалося завантажити жодне зображення')
    }

    if (successfulBatches < batches.length) {
        console.warn(`Увага: завантажено тільки ${successfulBatches} з ${batches.length} пакетів зображень`)
    }
}

// Функція для стандартної відправки всіх даних разом
async function createOrUpdateCarWithImages() {
    const formData = new FormData()

    // Додаємо всі дані форми
    Object.keys(form).forEach(key => {
        if (key !== 'photos') {
            formData.append(key, form[key])
        }
    })

    // Додаємо зображення
    form.photos.forEach(file => {
        formData.append('images', file)
    })

    // Якщо в режимі редагування, додаємо ID зображень для видалення
    if (isEditMode.value && imagesToDelete.value.length > 0) {
        formData.append('deleteImages', imagesToDelete.value.join(','))
    }

    // Відправляємо запит
    let url = '/api/cars-sale'
    let method = 'POST'

    if (isEditMode.value) {
        url = `/api/cars-sale/${carIdToEdit.value}`
        method = 'PUT'
    }

    const response = await fetch(url, {
        method,
        body: formData,
        headers: getAuthHeaders() // Исправлено!
    })

    if (!response.ok) {
        let errorText = 'Помилка збереження автомобіля'
        try {
            const errorData = await response.json()
            errorText = errorData.error || errorText
        } catch (e) {
            // Ігноруємо помилку парсингу JSON
            errorText = `Помилка сервера: ${response.status}`
        }
        throw new Error(errorText)
    }

    const result = await response.json()

    if (result.status !== 'success') {
        throw new Error(result.error || 'Помилка збереження автомобіля')
    }

    return isEditMode.value ? carIdToEdit.value : result.data.id
}
</script>

<style scoped>
.add-car-container {
    max-width: 1200px;
    margin: 2rem auto;
    padding: 0 1rem;
}

.form-wrapper {
    background: #fff;
    border-radius: 8px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    padding: 2rem;
}

h2 {
    margin-top: 0;
    margin-bottom: 1.5rem;
    color: #333;
    font-size: 1.75rem;
}

.form-row {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
    gap: 1.5rem;
    margin-bottom: 1.5rem;
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
    font-weight: 600;
    color: #333;
}

input,
select,
textarea {
    width: 100%;
    padding: 0.75rem 1rem;
    border: 2px solid #e0e0e0;
    border-radius: 4px;
    font-size: 1rem;
    transition: border-color 0.2s ease;
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

.processing-status {
    margin-top: 0.5rem;
    font-size: 0.875rem;
    color: #007bff;
    font-style: italic;
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

.photo-size {
    position: absolute;
    bottom: 4px;
    right: 4px;
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