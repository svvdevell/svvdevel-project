<template>
    <div class="login-container">
        <div class="login-box">
            <div class="logo">
                <h1>🚗 Elegance Auto</h1>
                <p>Адміністративна панель</p>
            </div>

            <form @submit.prevent="handleLogin" class="login-form">
                <div class="form-group">
                    <label for="username">Логін</label>
                    <input 
                        v-model="username" 
                        type="text" 
                        id="username" 
                        placeholder="Введіть логін"
                        required
                        :disabled="loading"
                        autocomplete="username"
                    >
                </div>

                <div class="form-group">
                    <label for="password">Пароль</label>
                    <div class="password-input">
                        <input 
                            v-model="password" 
                            :type="showPassword ? 'text' : 'password'" 
                            id="password" 
                            placeholder="Введіть пароль"
                            required
                            :disabled="loading"
                            autocomplete="current-password"
                        >
                        <button 
                            type="button" 
                            class="toggle-password" 
                            @click="showPassword = !showPassword"
                            :disabled="loading"
                        >
                            {{ showPassword ? '👁️' : '👁️‍🗨️' }}
                        </button>
                    </div>
                </div>

                <button type="submit" class="btn-login" :disabled="loading">
                    <span v-if="!loading">Увійти</span>
                    <span v-else>Перевірка...</span>
                </button>

                <div v-if="error" class="error-message">
                    {{ error }}
                </div>
            </form>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const username = ref('')
const password = ref('')
const showPassword = ref(false)
const loading = ref(false)
const error = ref('')

const handleLogin = async () => {
    error.value = ''
    loading.value = true

    try {
        const apiUrl = process.env.NODE_ENV === 'production'
            ? '/api/auth/login'
            : 'http://localhost:8001/api/auth/login'

        const response = await fetch(apiUrl, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username: username.value.trim(),
                password: password.value.trim()
            }),
        })

        const data = await response.json()

        if (response.ok && data.status === 'success') {
            // Сохраняем токен
            authStore.setToken(data.token)
            authStore.setUsername(username.value)
            
            // Перенаправляем в админ панель
            router.push('/admin')
        } else {
            error.value = data.message || 'Невірний логін або пароль'
        }

    } catch (err) {
        console.error('Login error:', err)
        error.value = 'Помилка з\'єднання з сервером'
    } finally {
        loading.value = false
    }
}
</script>

<style scoped>
.login-container {
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    padding: 1rem;
}

.login-box {
    background: white;
    border-radius: 12px;
    box-shadow: 0 10px 40px rgba(0, 0, 0, 0.2);
    padding: 3rem;
    width: 100%;
    max-width: 400px;
    animation: slideUp 0.5s ease;
}

@keyframes slideUp {
    from {
        opacity: 0;
        transform: translateY(30px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}

.logo {
    text-align: center;
    margin-bottom: 2rem;
}

.logo h1 {
    margin: 0 0 0.5rem 0;
    color: #333;
    font-size: 2rem;
}

.logo p {
    margin: 0;
    color: #666;
    font-size: 0.95rem;
}

.login-form {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
}

.form-group {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
}

.form-group label {
    font-weight: 500;
    color: #555;
    font-size: 0.95rem;
}

.form-group input {
    padding: 0.875rem;
    border: 2px solid #ddd;
    border-radius: 6px;
    font-size: 1rem;
    transition: border-color 0.3s ease;
}

.form-group input:focus {
    outline: none;
    border-color: #667eea;
}

.form-group input:disabled {
    background: #f5f5f5;
    cursor: not-allowed;
}

.password-input {
    position: relative;
    display: flex;
    align-items: center;
}

.password-input input {
    width: 100%;
    padding-right: 3rem;
}

.toggle-password {
    position: absolute;
    right: 0.75rem;
    background: none;
    border: none;
    cursor: pointer;
    font-size: 1.2rem;
    padding: 0.5rem;
    opacity: 0.6;
    transition: opacity 0.3s;
}

.toggle-password:hover:not(:disabled) {
    opacity: 1;
}

.toggle-password:disabled {
    cursor: not-allowed;
}

.btn-login {
    padding: 1rem;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    border: none;
    border-radius: 6px;
    font-size: 1.05rem;
    font-weight: 600;
    cursor: pointer;
    transition: transform 0.2s, box-shadow 0.2s;
    margin-top: 0.5rem;
}

.btn-login:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 5px 15px rgba(102, 126, 234, 0.4);
}

.btn-login:active:not(:disabled) {
    transform: translateY(0);
}

.btn-login:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}

.error-message {
    background: #fee;
    color: #c33;
    padding: 0.875rem;
    border-radius: 6px;
    border: 1px solid #fcc;
    font-size: 0.9rem;
    text-align: center;
    animation: shake 0.5s;
}

@keyframes shake {
    0%, 100% { transform: translateX(0); }
    25% { transform: translateX(-10px); }
    75% { transform: translateX(10px); }
}

@media (max-width: 480px) {
    .login-box {
        padding: 2rem 1.5rem;
    }

    .logo h1 {
        font-size: 1.5rem;
    }
}
</style>