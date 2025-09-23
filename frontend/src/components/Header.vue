<template>
    <nav class="navigation">
        <div class="nav-container">
            <router-link to="/" class="logo"><img src="../media/images/logo.png" alt=""></router-link>

            <div class="nav-menu">
                <router-link to="/" class="nav-link">Главная</router-link>
                <router-link to="/brands" class="nav-link">Бренды</router-link>
                <router-link to="/reviews" class="nav-link">Отзывы</router-link>
                <router-link to="/form" class="nav-link">Форма</router-link>
                <router-link to="/contact" class="nav-link">Контакты</router-link>
            </div>

            <div class="contact-links">
                <a href="">
                    <img src="../media/icons/telegram.svg" alt="">
                </a>
                <a href="">
                    <img src="../media/icons/facebook.svg" alt="">
                </a>
                <a href="">
                    <img src="../media/icons/instagram.svg" alt="">
                </a>
                <a href="">
                    <img src="../media/icons/whatsup.svg" alt="">
                </a>
            </div>
        </div>
    </nav>
</template>

<style lang="scss" scoped>
.logo {
    width: 100%;
    max-width: 160px;

    & img {
        width: 100%;
        max-width: 160px;
    }
}

.navigation {
    width: 100%;
    position: fixed;
    top: 0;
    left: 0;
    z-index: 999;
    padding: 10px;
    padding-top: 0;
}

.nav-container {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 20px;
    padding: 10px 16px;
    border-radius: 20px;
    border-top-left-radius: 0;
    border-top-right-radius: 0;

    background: rgba(255, 255, 255, 0.07);
    backdrop-filter: blur(6px);
    -webkit-backdrop-filter: blur(6px);
    border: 1px solid rgba(255, 255, 255, 0.3);
    box-shadow:
        0 8px 32px rgba(0, 0, 0, 0.1),
        inset 0 1px 0 rgba(255, 255, 255, 0.5),
        inset 0 -1px 0 rgba(255, 255, 255, 0.1),
        inset 0 0 52px 26px rgba(255, 255, 255, 2.6);
    position: relative;
    overflow: hidden;

    ::before {
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        height: 1px;
        background: linear-gradient(90deg,
                transparent,
                rgba(255, 255, 255, 0.8),
                transparent);
    }

    ::after {
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        width: 1px;
        height: 100%;
        background: linear-gradient(180deg,
                rgba(255, 255, 255, 0.8),
                transparent,
                rgba(255, 255, 255, 0.3));
    }
}

.nav-menu {
    display: flex;
    align-items: center;
    gap: 20px;
}

.nav-link {
    font-size: 22px;
    font-weight: 500;
    color: #000;
    text-decoration: none;
    padding: 10px;
    border-radius: 12px;

    &:hover {
        background-color: #000;
        color: #FFF;
    }
}

.contact-links {
    display: flex;
    align-items: center;
    gap: 10px;

    & img {
        width: 30px;
        height: 30px;
        object-fit: fill;
    }
}
</style>