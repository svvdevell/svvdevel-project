<template>
    <div class="admin-container">
        <header class="admin-header">
            <div class="header-content">
                <h1>🚗 Elegance Auto - Адмін панель</h1>
                <div class="user-info">
                    <span class="username">👤 {{ authStore.username }}</span>
                    <button @click="handleLogout" class="btn-logout">Вийти</button>
                </div>
            </div>
        </header>

        <div class="admin-content">
            <div class="welcome-section">
                <h2>Вітаємо в адміністративній панелі!</h2>
                <p>Оберіть потрібну дію для управління каталогом автомобілів</p>
            </div>

            <div class="admin-cards">
                <div class="admin-card" @click="goTo('/admin/add')">
                    <div class="card-icon">➕</div>
                    <h3>Додати автомобіль</h3>
                    <p>Створити нове оголошення про продаж автомобіля</p>
                    <button class="card-button">Перейти</button>
                </div>

                <div class="admin-card" @click="goTo('/admin/list')">
                    <div class="card-icon">📋</div>
                    <h3>Список автомобілів</h3>
                    <p>Переглянути, редагувати або видалити існуючі автомобілі</p>
                    <button class="card-button">Перейти</button>
                </div>

                <div class="admin-card" @click="goTo('/catalog')">
                    <div class="card-icon">🏪</div>
                    <h3>Публічний каталог</h3>
                    <p>Переглянути каталог як його бачать відвідувачі</p>
                    <button class="card-button">Перейти</button>
                </div>
            </div>

            <div class="stats-section">
                <h2>Статистика</h2>
                <div class="stats-grid">
                    <div class="stat-card">
                        <div class="stat-value">{{ stats.total }}</div>
                        <div class="stat-label">Всього автомобілів</div>
                    </div>
                    <div class="stat-card">
                        <div class="stat-value">{{ stats.active }}</div>
                        <div class="stat-label">Активні</div>
                    </div>
                    <div class="stat-card">
                        <div class="stat-value">{{ stats.sold }}</div>
                        <div class="stat-label">Продано</div>
                    </div>
                    <div class="stat-card">
                        <div class="stat-value">{{ stats.new }}</div>
                        <div class="stat-label">Новинки</div>
                    </div>
                </div>
            </div>

            <div class="quick-actions">
                <h2>Швидкі дії</h2>
                <div class="actions-list">
                    <button @click="goTo('/admin/add')" class="action-btn">
                        ➕ Додати новий автомобіль
                    </button>
                    <button @click="refreshStats" class="action-btn">
                        🔄 Оновити статистику
                    </button>
                    <button @click="goTo('/')" class="action-btn">
                        🏠 На головну сторінку
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const stats = ref({
    total: 0,
    active: 0,
    sold: 0,
    new: 0
})

const loadStats = async () => {
    try {
        const apiUrl = process.env.NODE_ENV || 'http://localhost:8001'
        
        const response = await fetch(`${apiUrl}/api/cars-sale?limit=1000`)
        
        if (response.ok) {
            const data = await response.json()
            const cars = data.data || []
            
            stats.value = {
                total: cars.length,
                active: cars.filter(car => car.status === 'active').length,
                sold: cars.filter(car => car.status === 'sold').length,
                new: cars.filter(car => car.status === 'new').length
            }
        }
    } catch (error) {
        console.error('Error loading stats:', error)
    }
}

const refreshStats = () => {
    loadStats()
}

const goTo = (path) => {
    router.push(path)
}

const handleLogout = async () => {
    try {
        const apiUrl = process.env.NODE_ENV || 'http://localhost:8001'
        
        await fetch(`${apiUrl}/api/auth/logout`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${authStore.token}`
            }
        })
    } catch (error) {
        console.error('Logout error:', error)
    } finally {
        authStore.logout()
        router.push('/login')
    }
}

onMounted(() => {
    loadStats()
})
</script>

<style scoped>
.admin-container {
    min-height: 100vh;
    background: #f5f7fa;
}

.admin-header {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    padding: 1.5rem 2rem;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.header-content {
    max-width: 1400px;
    margin: 0 auto;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.admin-header h1 {
    margin: 0;
    font-size: 1.8rem;
}

.user-info {
    display: flex;
    align-items: center;
    gap: 1rem;
}

.username {
    font-size: 1rem;
    font-weight: 500;
}

.btn-logout {
    padding: 0.6rem 1.2rem;
    background: rgba(255, 255, 255, 0.2);
    color: white;
    border: 2px solid white;
    border-radius: 6px;
    cursor: pointer;
    font-weight: 500;
    transition: all 0.3s ease;
}

.btn-logout:hover {
    background: white;
    color: #667eea;
}

.admin-content {
    max-width: 1400px;
    margin: 0 auto;
    padding: 2rem;
}

.welcome-section {
    text-align: center;
    margin-bottom: 3rem;
}

.welcome-section h2 {
    margin: 0 0 0.5rem 0;
    color: #333;
    font-size: 2rem;
}

.welcome-section p {
    margin: 0;
    color: #666;
    font-size: 1.1rem;
}

.admin-cards {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 2rem;
    margin-bottom: 3rem;
}

.admin-card {
    background: white;
    border-radius: 12px;
    padding: 2rem;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    cursor: pointer;
    transition: all 0.3s ease;
    text-align: center;
}

.admin-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
}

.card-icon {
    font-size: 3rem;
    margin-bottom: 1rem;
}

.admin-card h3 {
    margin: 0 0 1rem 0;
    color: #333;
    font-size: 1.5rem;
}

.admin-card p {
    margin: 0 0 1.5rem 0;
    color: #666;
    line-height: 1.6;
}

.card-button {
    padding: 0.75rem 2rem;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    border: none;
    border-radius: 6px;
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
}

.card-button:hover {
    transform: scale(1.05);
}

.stats-section {
    margin-bottom: 3rem;
}

.stats-section h2 {
    margin: 0 0 1.5rem 0;
    color: #333;
    font-size: 1.8rem;
}

.stats-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 1.5rem;
}

.stat-card {
    background: white;
    border-radius: 12px;
    padding: 2rem;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    text-align: center;
}

.stat-value {
    font-size: 3rem;
    font-weight: 700;
    color: #667eea;
    margin-bottom: 0.5rem;
}

.stat-label {
    color: #666;
    font-size: 1rem;
    font-weight: 500;
}

.quick-actions h2 {
    margin: 0 0 1.5rem 0;
    color: #333;
    font-size: 1.8rem;
}

.actions-list {
    display: flex;
    flex-wrap: wrap;
    gap: 1rem;
}

.action-btn {
    padding: 1rem 1.5rem;
    background: white;
    color: #667eea;
    border: 2px solid #667eea;
    border-radius: 6px;
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
}

.action-btn:hover {
    background: #667eea;
    color: white;
}

@media (max-width: 768px) {
    .admin-header {
        padding: 1rem;
    }

    .header-content {
        flex-direction: column;
        gap: 1rem;
        align-items: flex-start;
    }

    .admin-header h1 {
        font-size: 1.3rem;
    }

    .admin-content {
        padding: 1rem;
    }

    .welcome-section h2 {
        font-size: 1.5rem;
    }

    .admin-cards {
        grid-template-columns: 1fr;
    }

    .stats-grid {
        grid-template-columns: repeat(2, 1fr);
    }

    .actions-list {
        flex-direction: column;
    }

    .action-btn {
        width: 100%;
    }
}
</style>