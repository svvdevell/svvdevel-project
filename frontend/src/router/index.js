import { createRouter, createWebHistory } from 'vue-router'

// Импортируйте ваши компоненты
import Main from '@/views/Main.vue'
import Brands from '@/views/Brands.vue'
import NotFound from '@/views/NotFound.vue'

const routes = [
    {
        path: '/',
        name: 'Main',
        component: Main
    },
    {
        path: '/brands',
        name: 'Brands',
        component: Brands
    },
    {
        path: '/contact',
        name: 'Contact',
        // Lazy loading компонента
        component: () => import('@/views/Contact.vue')
    },
    {
        path: '/form',
        name: 'Form',
        component: () => import('@/views/Form.vue')
    },
    {
        path: '/reviews',
        name: 'Reviews',
        component: () => import('@/views/Reviews.vue')
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: NotFound
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
    // Настройка скролла при переходах между страницами
    scrollBehavior(to, from, savedPosition) {
        if (savedPosition) {
            return savedPosition
        } else {
            return { top: 0 }
        }
    }
})

// Navigation guards (опционально)
router.beforeEach((to, from, next) => {
    // Логика перед переходом на новый маршрут
    console.log(`Переход с ${from.path} на ${to.path}`)
    next()
})

export default router