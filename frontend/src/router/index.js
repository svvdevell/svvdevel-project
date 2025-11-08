// router/index.js

import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

// Імпорт компонентів
import Main from '@/views/Main.vue'
import NotFound from '@/views/NotFound.vue'
import CarsCatalog from '@/views/CarsCatalog.vue'
import Login from '@/views/Login.vue'
import CarPage from '@/views/CarPage.vue'
import BlogView from '@/views/BlogView.vue'

const routes = [
    {
        path: '/',
        name: 'Main',
        component: Main,
        meta: {
            title: 'Головна'
        }
    },
    {
        path: '/catalog',
        name: 'CarsCatalog',
        component: CarsCatalog,
        meta: {
            title: 'У продажі!'
        }
    },
    {
        path: '/blog',
        name: 'BlogView',
        component: BlogView,
        meta: {
            title: 'Корисні поради - Elegance Auto'
        }
    },
    {
        path: '/cars/:id',
        name: 'CarPage',
        component: CarPage,
        meta: {
            title: 'Деталі автомобіля'
        }
    },
    {
        path: '/contact',
        name: 'Contact',
        component: () => import('@/views/Contact.vue'),
        meta: {
            title: 'Контакти'
        }
    },
    // Сторінка входу
    {
        path: '/login',
        name: 'Login',
        component: Login,
        meta: {
            title: 'Вхід в адмін панель',
            requiresGuest: true // Тільки для неавторизованих
        }
    },
    // Адміністративні маршрути (захищені)
    {
        path: '/admin',
        name: 'Admin',
        component: () => import('@/views/admin/Admin.vue'),
        meta: {
            title: 'Адмін панель',
            requiresAuth: true
        }
    },
    {
        path: '/admin/add',
        name: 'AddCar',
        component: () => import('@/views/admin/AddCar.vue'),
        meta: {
            title: 'Додати автомобіль',
            requiresAuth: true
        }
    },
    {
        path: '/admin/list',
        name: 'CarsList',
        component: () => import('@/views/admin/CarList.vue'),
        meta: {
            title: 'Список автомобілів',
            requiresAuth: true
        }
    },
    {
        path: '/admin/edit/:id',
        name: 'EditCar',
        component: () => import('@/views/admin/AddCar.vue'),
        meta: {
            title: 'Редагувати автомобіль',
            requiresAuth: true
        }
    },
    {
        path: '/admin/detail/:id',
        name: 'AdminCarDetail',
        component: () => import('@/views/admin/CarDetail.vue'),
        meta: {
            title: 'Деталі автомобіля',
            requiresAuth: true
        }
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: NotFound,
        meta: {
            title: 'Сторінка не знайдена'
        }
    }
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL || '/'),
    routes,
    scrollBehavior(to, from, savedPosition) {
        if (savedPosition) {
            return savedPosition
        } else {
            return { top: 0 }
        }
    }
})

// Navigation guard для захисту маршрутів
router.beforeEach(async (to, from, next) => {
    const authStore = useAuthStore()

    console.log(`Перехід з ${from.path} на ${to.path}`)

    // Оновлюємо заголовок сторінки
    document.title = to.meta.title || 'Elegance Auto'

    // Перевіряємо чи потрібна аутентифікація
    if (to.meta.requiresAuth) {
        if (!authStore.isAuthenticated) {
            // Не авторизований - відправляємо на логін
            next({
                name: 'Login',
                query: { redirect: to.fullPath } // Зберігаємо URL для редіректу після входу
            })
            return
        }

        // Перевіряємо валідність токена
        const isValid = await authStore.verifyToken()

        if (!isValid) {
            // Токен невалідний - відправляємо на логін
            next({
                name: 'Login',
                query: { redirect: to.fullPath }
            })
            return
        }
    }

    // Якщо сторінка тільки для гостей (логін) і користувач авторизований
    if (to.meta.requiresGuest && authStore.isAuthenticated) {
        next({ name: 'Admin' })
        return
    }

    next()
})

export default router