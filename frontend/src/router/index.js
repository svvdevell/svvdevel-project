import { createRouter, createWebHistory } from 'vue-router'

// Импортируйте ваши компоненты
import Main from '@/views/Main.vue'
import NotFound from '@/views/NotFound.vue'
import AddCar from '@/views/admin/AddCar.vue'
import CarsCatalog from '@/views/CarsCatalog.vue'


const routes = [
    {
        path: '/',
        name: 'Main',
        component: Main
    },
    {
        path: '/test',
        name: 'Test',
        component: () => import('@/views/Test.vue')
    },
    {
        path: '/catalog',
        name: 'CarsCatalog',
        component: CarsCatalog
    },
    {
        path: '/admin/add',
        name: 'AddCar',
        component: AddCar
    },
    {
        path: '/contact',
        name: 'Contact',
        component: () => import('@/views/Contact.vue')
    },
    {
        path: '/admin',
        name: 'Admin',
        component: () => import('@/views/admin/Admin.vue')
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: NotFound
    }
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL || '/'), // Используйте import.meta.env для Vite
    routes,
    scrollBehavior(to, from, savedPosition) {
        if (savedPosition) {
            return savedPosition
        } else {
            return { top: 0 }
        }
    }
})

router.beforeEach((to, from, next) => {
    console.log(`Переход с ${from.path} на ${to.path}`)
    next()
})

export default router