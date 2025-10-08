import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useAuthStore = defineStore('auth', () => {
    // State
    const token = ref(localStorage.getItem('admin_token') || null)
    const username = ref(localStorage.getItem('admin_username') || null)

    // Getters
    const isAuthenticated = computed(() => !!token.value)

    // Actions
    const setToken = (newToken) => {
        token.value = newToken
        localStorage.setItem('admin_token', newToken)
    }

    const setUsername = (newUsername) => {
        username.value = newUsername
        localStorage.setItem('admin_username', newUsername)
    }

    const logout = () => {
        token.value = null
        username.value = null
        localStorage.removeItem('admin_token')
        localStorage.removeItem('admin_username')
    }

    const verifyToken = async () => {
        if (!token.value) {
            return false
        }

        try {
            const apiUrl = process.env.NODE_ENV

            const response = await fetch(`${apiUrl}/api/auth/verify`, {
                headers: {
                    'Authorization': `Bearer ${token.value}`
                }
            })

            if (response.ok) {
                const data = await response.json()
                return data.valid
            }

            // Токен невалідний
            logout()
            return false

        } catch (error) {
            console.error('Token verification error:', error)
            return false
        }
    }

    return {
        token,
        username,
        isAuthenticated,
        setToken,
        setUsername,
        logout,
        verifyToken
    }
})