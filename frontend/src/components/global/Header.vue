<template>
    <header class="navigation">
        <nav class="nav-container">
            <router-link to="/" class="logo"><img src="../../assets/images/logo.png" alt=""></router-link>

            <div class="nav-menu">
                <router-link to="/" class="nav-link">Головна</router-link>
                <router-link to="/test" class="nav-link">Test</router-link>
                <router-link to="/form" class="nav-link">Авто викуп</router-link>
                <router-link to="/catalog" class="nav-link">У продажі</router-link>
                <router-link to="/reviews" class="nav-link">Відгуки</router-link>
                <router-link to="/contact" class="nav-link">Контакты</router-link>
            </div>
            <!-- TODO add links for social -->
            <div class="contact-links">
                <a href="">
                    <img src="../../assets/icons/telegram.svg" alt="Telegram Icon">
                </a>
                <a href="">
                    <img src="../../assets/icons/facebook.svg" alt="Facebook Icon">
                </a>
                <a href="">
                    <img src="../../assets/icons/instagram.svg" alt="Instagram Icon">
                </a>
                <a href="">
                    <img src="../../assets/icons/whatsup.svg" alt="What's up Icon">
                </a>
            </div>
        </nav>
    </header>
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
    position: sticky;
    padding: 28px 16px;
    background-color: #FFF;
    z-index: 999;
    width: 100%;
    max-width: 1440px;
    margin: 0 auto;
}

.nav-container {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 20px;
}

.nav-menu {
    display: flex;
    align-items: center;
    gap: 20px;
}

.nav-link {
    font-family: "Inter 24pt" !important;
    font-weight: 500;
    font-size: 18px;
    line-height: 100%;
    letter-spacing: 0%;
    color: #000;
    text-decoration: none;
    padding: 0 6px;

    &:hover {
        font-weight: 700;
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

        &:hover {
            scale: (1.5);
        }
    }
}
</style>