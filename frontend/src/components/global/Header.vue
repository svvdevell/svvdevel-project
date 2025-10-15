<template>
    <header class="navigation" :class="{ 'scrolled': isScrolled }">
        <nav class="nav-container">
            <router-link to="/" class="logo">
                <img src="../../assets/images/logo.png" alt="Elegance Auto">
            </router-link>

            <!-- Desktop Menu -->
            <div class="nav-menu" :class="{ 'active': isMenuOpen }">
                <router-link to="/" class="nav-link" @click="closeMenu">Головна</router-link>
                <!-- <router-link to="/form" class="nav-link" @click="closeMenu">Авто викуп</router-link> -->
                <router-link to="/catalog" class="nav-link" @click="closeMenu">У продажі</router-link>
                <!-- <router-link to="/reviews" class="nav-link" @click="closeMenu">Відгуки</router-link> -->
                <router-link to="/contact" class="nav-link" @click="closeMenu">Контакти</router-link>

                <!-- Social Links in Mobile Menu -->
                <div class="mobile-social">
                    <a href="https://t.me/eleganceauto_odessa" @click="closeMenu">
                        <img src="../../assets/icons/telegram.svg" alt="Telegram">
                    </a>
                    <a href="https://auto.ria.com/uk/pro/seller-4625376.html" @click="closeMenu">
                        <img src="../../assets/icons/auto-ria.png" alt="AutoRia" class="bordered">
                    </a>
                    <a href="https://www.instagram.com/elegance_auto_od/" @click="closeMenu">
                        <img src="../../assets/icons/instagram.svg" alt="Instagram">
                    </a>
                    <a href="https://www.tiktok.com/@elegance_auto_od" @click="closeMenu">
                        <img src="../../assets/icons/whatsup.svg" alt="TikTok">
                    </a>
                </div>
            </div>

            <!-- Desktop Social Links -->
            <div class="contact-links">
                <a href="https://t.me/eleganceauto_odessa">
                    <img src="../../assets/icons/telegram.svg" alt="Telegram">
                </a>
                <a href="https://auto.ria.com/uk/pro/seller-4625376.html">
                    <img src="../../assets/icons/auto-ria.png" alt="AutoRia" class="bordered">
                </a>
                <a href="https://www.instagram.com/elegance_auto_od/">
                    <img src="../../assets/icons/instagram.svg" alt="Instagram">
                </a>
                <a href="https://www.tiktok.com/@elegance_auto_od">
                    <img src="../../assets/icons/whatsup.svg" alt="TikTok">
                </a>
            </div>

            <!-- Burger Button -->
            <button class="burger-btn" @click="toggleMenu" :class="{ 'active': isMenuOpen }" aria-label="Toggle menu">
                <span></span>
                <span></span>
                <span></span>
            </button>
        </nav>
    </header>
</template>

<script setup>
import { ref, watch, onMounted, onUnmounted } from 'vue'
import { useRoute } from 'vue-router'

const isMenuOpen = ref(false)
const isScrolled = ref(false)
const route = useRoute()

const handleScroll = () => {
    // Якщо скрол більше 50px - додаємо клас
    isScrolled.value = window.scrollY > 50
}

const toggleMenu = () => {
    isMenuOpen.value = !isMenuOpen.value
    // Prevent body scroll when menu is open
    if (isMenuOpen.value) {
        document.body.style.overflow = 'hidden'
    } else {
        document.body.style.overflow = ''
    }
}

const closeMenu = () => {
    isMenuOpen.value = false
    document.body.style.overflow = ''
}

// Close menu on route change
watch(route, () => {
    closeMenu()
})

// Додаємо слухач скролу
onMounted(() => {
    window.addEventListener('scroll', handleScroll)
    // Перевіряємо стан при монтуванні
    handleScroll()
})

// Прибираємо слухач при демонтуванні
onUnmounted(() => {
    window.removeEventListener('scroll', handleScroll)
    document.body.style.overflow = ''
})
</script>

<style lang="scss" scoped>
.logo {
    width: 100%;
    max-width: 160px;
    z-index: 1001;

    & img {
        width: 100%;
        max-width: 160px;
    }
}

.navigation {
    width: 100%;
    position: fixed;
    padding: 12px 20px;
    z-index: 1000;
    top: 0;
    left: 0;
    background: rgba(255, 255, 255, 0.05);
    backdrop-filter: blur(15px);
    border-bottom: 1px solid rgba(255, 255, 255, 0.2);
    box-shadow: 0 2px 20px rgba(0, 0, 0, 0.1);
    transition: all 0.3s ease;

    // Стилі коли скролимо
    &.scrolled {
        background: rgba(26, 26, 26, 0.95);
        backdrop-filter: blur(20px);
        border-bottom: 1px solid rgba(255, 255, 255, 0.1);
        box-shadow: 0 4px 30px rgba(0, 0, 0, 0.3);
        padding: 8px 20px;
    }
}

.nav-container {
    max-width: 1400px;
    margin: 0 auto;
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 20px;
}

.nav-menu {
    display: flex;
    align-items: center;
    gap: 32px;
}

.nav-link {
    font-family: 'Work Sans', sans-serif;
    font-weight: 500;
    font-size: 16px;
    line-height: 100%;
    color: #FFF;
    text-decoration: none;
    padding: 8px 0;
    position: relative;
    transition: all 0.3s ease;

    &::after {
        content: '';
        position: absolute;
        bottom: 0;
        left: 0;
        width: 0;
        height: 2px;
        background: #4CAF50;
        transition: width 0.3s ease;
    }

    &:hover {
        color: #4CAF50;

        &::after {
            width: 100%;
        }
    }

    &.router-link-active {
        font-weight: 600;
        color: #4CAF50;

        &::after {
            width: 100%;
        }
    }
}

.contact-links {
    display: flex;
    align-items: center;
    gap: 12px;

    & a {
        width: 40px;
        height: 40px;
        display: flex;
        align-items: center;
        justify-content: center;
        border-radius: 10px;
        background: rgba(255, 255, 255, 0.1);
        transition: all 0.3s ease;

        &:hover {
            transform: translateY(-4px);
            background: rgba(255, 255, 255, 0.2);
        }
    }

    & img {
        width: 24px;
        height: 24px;
        object-fit: contain;
    }
}

.bordered {
    border-radius: 6px;
}

// Burger Button
.burger-btn {
    display: none;
    flex-direction: column;
    justify-content: space-between;
    width: 30px;
    height: 24px;
    background: transparent;
    border: none;
    cursor: pointer;
    padding: 0;
    z-index: 1001;
    position: relative;

    span {
        width: 100%;
        height: 3px;
        background: #FFF;
        border-radius: 2px;
        transition: all 0.3s ease;
        transform-origin: center;
    }

    &.active {
        span:nth-child(1) {
            transform: translateY(10.5px) rotate(45deg);
        }

        span:nth-child(2) {
            opacity: 0;
            transform: translateX(-20px);
        }

        span:nth-child(3) {
            transform: translateY(-10.5px) rotate(-45deg);
        }
    }
}

// Mobile Social Links (hidden by default)
.mobile-social {
    display: none;
}

// Mobile Styles
@media (max-width: 768px) {
    .navigation {
        padding: 10px 16px;
        background: linear-gradient(180deg, #1a1a1a 0%, #0f0f0f 100%);

        &.scrolled {
            padding: 8px 16px;
        }
    }

    .logo {
        max-width: 120px;

        & img {
            max-width: 120px;
        }
    }

    .burger-btn {
        display: flex;
        z-index: 99999;
    }

    .contact-links {
        display: none;
    }

    .nav-menu {
        position: fixed;
        z-index: 9999;
        top: 0;
        right: -100%;
        width: 80%;
        max-width: 350px;
        height: 100vh;
        background: linear-gradient(180deg, #1a1a1a 0%, #0f0f0f 100%);
        flex-direction: column;
        align-items: flex-start;
        gap: 0;
        padding: 80px 30px 30px;
        transition: right 0.4s ease;
        overflow-y: auto;
        box-shadow: -5px 0 30px rgba(0, 0, 0, 0.5);

        &.active {
            right: 0;
        }
    }

    .nav-link {
        width: 100%;
        font-size: 20px;
        padding: 16px 0;
        border-bottom: 1px solid rgba(255, 255, 255, 0.1);

        &::after {
            display: none;
        }

        &:hover {
            padding-left: 12px;
            color: #4CAF50;
        }
    }

    .mobile-social {
        display: flex;
        gap: 16px;
        margin-top: 32px;
        padding-top: 32px;
        border-top: 1px solid rgba(255, 255, 255, 0.1);
        width: 100%;

        a {
            width: 48px;
            height: 48px;
            display: flex;
            align-items: center;
            justify-content: center;
            border-radius: 12px;
            background: rgba(255, 255, 255, 0.1);
            transition: all 0.3s ease;

            &:hover {
                transform: scale(1.1);
                background: rgba(255, 255, 255, 0.2);
            }

            img {
                width: 28px;
                height: 28px;

                &.bordered {
                    border-radius: 8px;
                }
            }
        }
    }
}

// Tablet Styles
@media (max-width: 1024px) {
    .nav-menu {
        gap: 20px;
    }

    .nav-link {
        font-size: 15px;
    }
}

// Small Mobile
@media (max-width: 480px) {
    .nav-menu {
        width: 85%;
    }

    .logo {
        max-width: 100px;

        & img {
            max-width: 100px;
        }
    }
}
</style>