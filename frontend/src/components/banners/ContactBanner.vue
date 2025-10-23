<template>
    <section class="contacts-section">
        <div class="contacts-container">
            <!-- Hero Block -->
            <div class="hero-block">
                <h2>Продайте авто швидко та вигідно</h2>
                <p>Оцінимо ваш автомобіль протягом 30 хвилин. Виплата готівкою в день звернення</p>
            </div>

            <!-- Messenger Block -->
            <div class="messenger-block">
                <h3>Або напишіть в месенджер</h3>
                <div class="messenger-grid">
                    <a href="https://t.me/trewqxm" class="messenger-card telegram">
                        <img src="../../assets/icons/telegram.svg" alt="">
                        <div>
                            <span class="messenger-title">Telegram</span>
                            <span class="messenger-desc">Швидка відповідь</span>
                        </div>
                    </a>
                    <a href="tel:0734080999" class="messenger-card telegram">
                        <img src="../../assets/icons/phone.svg" alt="">
                        <div>
                            <span class="messenger-title">Lifecell</span>
                            <span class="messenger-desc">073 408 0999</span>
                        </div>
                    </a>
                    <a href="https://t.me/eleganceauto_odessa" class="messenger-card telegram-channel">
                        <img src="../../assets/icons/telegram.svg" alt="">
                        <div>
                            <span class="messenger-title">Telegram канал</span>
                            <span class="messenger-desc">Актуальні пропозиції</span>
                        </div>
                    </a>
                </div>
            </div>

            <!-- Social & Location -->
            <div class="bottom-grid">
                <a href="https://www.instagram.com/elegance_auto_od/" class="social-card instagram">
                    <img src="../../assets/icons/instagram.svg" alt="">
                    <div>
                        <span class="social-title">Instagram</span>
                        <span class="social-desc">Фото нових авто</span>
                    </div>
                </a>

                <a href="https://auto.ria.com/uk/pro/seller-4625376.html" class="social-card autoria">
                    <img src="../../assets/icons/auto-ria.png" alt="" class="bordered">
                    <div>
                        <span class="social-title">Auto.ria</span>
                        <span class="social-desc">Наші оголошення</span>
                    </div>
                </a>

                <a href="https://www.google.com/maps?q=%D0%90%D0%B2%D1%82%D0%BE%D1%81%D0%B0%D0%BB%D0%BE%D0%BD+%22Elegance+Auto%22"
                    class="social-card maps">
                    <img src="../../assets/icons/geo.svg" alt="">
                    <div>
                        <span class="social-title">Наша адреса</span>
                        <span class="social-desc">пр. Небесної Сотні, 100/1</span>
                    </div>
                </a>
            </div>

            <!-- Image Gallery -->
            <div class="map">
                <iframe
                    src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2752.228414344547!2d30.702056776324614!3d46.38469457110396!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x40c7cdf97dcee44f%3A0x6a2762df677153ae!2z0JDQstGC0L7RgdCw0LvQvtC9ICJFbGVnYW5jZSBBdXRvIg!5e0!3m2!1sru!2sua!4v1759929964549!5m2!1sru!2sua"
                    style="border:0;" allowfullscreen="" loading="lazy"
                    referrerpolicy="no-referrer-when-downgrade"></iframe>
            </div>
        </div>
    </section>
</template>

<script setup>
</script>

<style lang="scss" scoped>
.contacts-section {
    width: 100%;
    padding: 50px 0;
}

.contacts-container {
    max-width: 1400px;
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    gap: 40px;
}

// Hero Block
.hero-block {
    text-align: center;
    margin-bottom: 20px;

    h2 {
        font-family: 'Work Sans', sans-serif;
        font-size: 48px;
        font-weight: 700;
        margin-bottom: 16px;
        line-height: 1.2;
    }

    p {
        font-family: 'Work Sans', sans-serif;
        font-size: 20px;
        max-width: 600px;
        margin: 0 auto;
    }
}

// Main Contact Card
.main-contact-card {
    background: white;
    border-radius: 24px;
    padding: 40px;
    box-shadow: 0 4px 24px rgba(0, 0, 0, 0.06);
}

.card-header {
    display: flex;
    align-items: center;
    gap: 20px;
    margin-bottom: 32px;

    .icon-badge {
        width: 64px;
        height: 64px;
        background: linear-gradient(135deg, #4CAF50 0%, #45a049 100%);
        border-radius: 16px;
        display: flex;
        align-items: center;
        justify-content: center;

        img {
            width: 32px;
            height: 32px;
        }
    }

    h3 {
        font-family: 'Work Sans', sans-serif;
        font-size: 28px;
        font-weight: 700;
        color: #000;
        margin: 0;
    }

    p {
        font-size: 16px;
        color: #666;
        margin: 4px 0 0 0;
    }
}

.phone-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 16px;
}

.phone-btn {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 24px;
    background: linear-gradient(135deg, #4CAF50 0%, #45a049 100%);
    border-radius: 16px;
    text-decoration: none;
    transition: all 0.3s ease;
    box-shadow: 0 4px 12px rgba(76, 175, 80, 0.2);

    &:hover {
        transform: translateY(-4px);
        box-shadow: 0 8px 24px rgba(76, 175, 80, 0.3);
    }

    .phone-number {
        font-family: 'Work Sans', sans-serif;
        font-size: 24px;
        font-weight: 700;
        color: white;
        margin-bottom: 4px;
    }

    .phone-label {
        font-size: 14px;
        color: rgba(255, 255, 255, 0.9);
    }
}

// Messenger Block
.messenger-block {
    h3 {
        font-family: 'Work Sans', sans-serif;
        font-size: 24px;
        font-weight: 600;
        margin-bottom: 20px;
        text-align: center;
    }
}

.messenger-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    gap: 20px;
}

.messenger-card {
    display: flex;
    align-items: center;
    gap: 20px;
    padding: 24px 32px;
    background: white;
    border-radius: 20px;
    text-decoration: none;
    transition: all 0.3s ease;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.06);
    border: 2px solid transparent;

    &:hover {
        transform: translateY(-4px);
        box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
    }

    &.telegram {
        &:hover {
            border-color: #0088cc;
        }
    }

    &.telegram-channel {
        &:hover {
            border-color: #0088cc;
        }
    }

    img {
        width: 48px;
        height: 48px;
    }

    .messenger-title {
        display: block;
        font-family: 'Work Sans', sans-serif;
        font-size: 20px;
        font-weight: 600;
        color: #000;
        margin-bottom: 4px;
    }

    .messenger-desc {
        display: block;
        font-size: 14px;
        color: #666;
    }
}

// Bottom Grid
.bottom-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    gap: 20px;
}

.social-card {
    display: flex;
    align-items: center;
    gap: 20px;
    padding: 24px 32px;
    background: white;
    border-radius: 20px;
    text-decoration: none;
    transition: all 0.3s ease;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.06);
    border: 2px solid transparent;

    &:hover {
        transform: translateY(-4px);
        box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
    }

    &.instagram:hover {
        border-color: #E4405F;
    }

    &.autoria:hover {
        border-color: #aa3535;
    }

    &.maps:hover {
        border-color: #4285F4;
    }

    img {
        width: 48px;
        height: 48px;

        &.bordered {
            border-radius: 12px;
        }
    }

    .social-title {
        display: block;
        font-family: 'Work Sans', sans-serif;
        font-size: 18px;
        font-weight: 600;
        color: #000;
        margin-bottom: 4px;
    }

    .social-desc {
        display: block;
        font-size: 14px;
        color: #666;
    }
}

.map {
    width: 100%;
    padding: 24px 32px;
    background: white;
    border-radius: 20px;
    text-decoration: none;
    transition: all 0.3s ease;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.06);
    border: 2px solid transparent;

    & iframe {
        width: 100%;
        min-height: 600px;
    }
}

// Responsive
@media (max-width: 768px) {
    .contacts-section {
        padding: 40px 0px;
    }

    .hero-block {
        h2 {
            font-size: 32px;
        }

        p {
            font-size: 16px;
        }
    }

    .main-contact-card {
        padding: 24px;
    }

    .card-header {
        flex-direction: column;
        text-align: center;

        h3 {
            font-size: 24px;
        }
    }

    .phone-grid {
        grid-template-columns: 1fr;
    }

    .phone-btn {
        .phone-number {
            font-size: 20px;
        }
    }

    .map {
        padding: 10px;

        & iframe {
            min-height: 300px;
        }
    }
}
</style>