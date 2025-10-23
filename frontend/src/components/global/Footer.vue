<template>
    <footer class="footer">
        <!-- Main Footer Content -->
        <div class="footer-container">
            <!-- Top Section -->
            <div class="footer-top">
                <div class="footer-brand">
                    <a href="" class="logo">
                        <img src="../../assets/images/logo.png" alt="Elegance Auto">
                    </a>
                    <p class="brand-description">
                        Професійний автовикуп в Одесі. Швидка оцінка, чесна ціна, миттєва виплата.
                    </p>
                    <div class="social-links">
                        <a href="https://www.instagram.com/elegance_auto_od/" class="social-link instagram"
                            aria-label="Instagram">
                            <img src="../../assets/icons/instagram.svg" alt="">
                        </a>
                        <a href="https://t.me/eleganceauto_odessa" class="social-link telegram" aria-label="Telegram">
                            <img src="../../assets/icons/telegram.svg" alt="">
                        </a>
                        <a href="https://auto.ria.com/uk/pro/seller-4625376.html" class="social-link autoria"
                            aria-label="Auto.ria">
                            <img src="../../assets/icons/auto-ria.png" alt="" class="bordered-icon">
                        </a>
                    </div>
                </div>

                <!-- Contact Info -->
                <div class="footer-contacts">
                    <h3 class="footer-title">Контакти</h3>
                    <div class="contact-item">
                        <img src="../../assets/icons/phone.svg" alt="" class="contact-icon">
                        <div>
                            <a href="tel:0734080999" class="contact-link">+38 (073) 40 80 999</a>
                        </div>
                    </div>
                    <div class="contact-item">
                        <img src="../../assets/icons/telegram.svg" alt="" class="contact-icon">
                        <div>
                            <a href="https://t.me/trewqxm" class="contact-link">Менеджер у Телеграмі</a>
                            <a href="https://t.me/eleganceauto_odessa" class="contact-link">Наш канал</a>
                        </div>
                    </div>
                </div>

                <!-- Address -->
                <div class="footer-address">
                    <h3 class="footer-title">Наша адреса</h3>
                    <div class="contact-item">
                        <img src="../../assets/icons/geo.svg" alt="" class="contact-icon">
                        <div>
                            <a href="https://www.google.com/maps?q=%D0%90%D0%B2%D1%82%D0%BE%D1%81%D0%B0%D0%BB%D0%BE%D0%BD+%22Elegance+Auto%22"
                                class="contact-link address-link">
                                м. Одеса,<br>
                                пр. Небесної Сотні, 100/1
                            </a>
                            <p class="work-hours">Працюємо щодня: 9:00 - 20:00</p>
                        </div>
                    </div>
                </div>

                <!-- Quick Links -->
                <div class="footer-links">
                    <h3 class="footer-title">Послуги</h3>
                    <router-link to="/contact" class="footer-link">Викуп легкових авто</router-link>
                    <router-link to="/contact" class="footer-link">Викуп позашляховиків</router-link>
                    <router-link to="/contact" class="footer-link">Оцінка авто онлайн</router-link>
                </div>
            </div>

            <!-- Map Section -->
            <div class="footer-map">
                <iframe
                    src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3644.5524686286963!2d30.7046317!3d46.384694599999996!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x40c7cdf97dcee44f%3A0x6a2762df677153ae!2z0JDQstGC0L7RgdCw0LvQvtC9ICJFbGVnYW5jZSBBdXRvIg!5e1!3m2!1sru!2sua!4v1759781170565!5m2!1sru!2sua"
                    width="100%" height="300" style="border:0; border-radius: 16px;" allowfullscreen="" loading="lazy"
                    referrerpolicy="no-referrer-when-downgrade">
                </iframe>
            </div>
        </div>

        <!-- Bottom Bar -->
        <div class="footer-bottom">
            <div class="footer-container">
                <div class="bottom-content">
                    <p class="copyright">
                        © 2025 Elegance Auto. Всі права захищені
                    </p>
                    <p class="developer">
                        Розробка та дизайн: <a href="https://svvdevel.tech/" target="_blank"
                            rel="noopener">svvdevel.tech</a>
                    </p>
                </div>
            </div>
        </div>
    </footer>
</template>

<script setup>
</script>

<style lang="scss" scoped>
.footer {
    width: 100%;
    background: linear-gradient(180deg, #1a1a1a 0%, #0f0f0f 100%);
    color: white;
    padding: 0 16px;
}

.footer-container {
    max-width: 1400px;
    margin: 0 auto;
}

// Top Section
.footer-top {
    display: grid;
    grid-template-columns: 2fr 1.5fr 1.5fr 1fr;
    gap: 60px;
    padding: 80px 0 60px;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);

    @media (max-width: 1024px) {
        grid-template-columns: repeat(2, 1fr);
        gap: 40px;
    }

    @media (max-width: 640px) {
        grid-template-columns: 1fr;
        gap: 40px;
        padding: 60px 0 40px;
    }
}

// Brand Section
.footer-brand {
    .logo {
        display: inline-block;
        margin-bottom: 20px;

        img {
            width: 160px;
            height: auto;
        }
    }

    .brand-description {
        font-size: 16px;
        line-height: 1.6;
        color: rgba(255, 255, 255, 0.7);
        margin-bottom: 24px;
    }
}

.social-links {
    display: flex;
    gap: 12px;

    .social-link {
        width: 48px;
        height: 48px;
        border-radius: 12px;
        background: rgba(255, 255, 255, 0.1);
        display: flex;
        align-items: center;
        justify-content: center;
        transition: all 0.3s ease;

        &:hover {
            transform: translateY(-4px);
            background: rgba(255, 255, 255, 0.15);
        }

        img {
            width: 24px;
            height: 24px;

            &.bordered-icon {
                border-radius: 6px;
            }
        }

        &.instagram:hover {
            background: linear-gradient(135deg, #833AB4, #E1306C, #FD1D1D);
        }

        &.telegram:hover {
            background: #0088cc;
        }

        &.autoria:hover {
            background: #aa3535;
        }
    }
}

// Footer Sections
.footer-title {
    font-family: 'Work Sans', sans-serif;
    font-size: 20px;
    font-weight: 600;
    margin-bottom: 20px;
    color: white;
}

.contact-item {
    display: flex;
    gap: 16px;
    margin-bottom: 24px;

    &:last-child {
        margin-bottom: 0;
    }

    .contact-icon {
        width: 24px;
        height: 24px;
        flex-shrink: 0;
        margin-top: 4px;
    }

    div {
        display: flex;
        flex-direction: column;
        gap: 8px;
    }

    .contact-link {
        color: rgba(255, 255, 255, 0.8);
        text-decoration: none;
        font-size: 16px;
        transition: color 0.3s ease;

        &:hover {
            color: #4CAF50;
        }

        &.address-link {
            line-height: 1.6;
        }
    }

    .work-hours {
        font-size: 14px;
        color: rgba(255, 255, 255, 0.5);
        margin-top: 4px;
    }
}

// Footer Links
.footer-links {
    display: flex;
    flex-direction: column;

    .footer-link {
        color: rgba(255, 255, 255, 0.7);
        text-decoration: none;
        font-size: 15px;
        padding: 8px 0;
        transition: all 0.3s ease;
        position: relative;

        &:hover {
            color: white;
            padding-left: 8px;
        }

        &::before {
            content: '';
            position: absolute;
            left: 0;
            top: 50%;
            transform: translateY(-50%);
            width: 0;
            height: 2px;
            background: #4CAF50;
            transition: width 0.3s ease;
        }

        &:hover::before {
            width: 4px;
        }
    }
}

// Map Section
.footer-map {
    padding: 40px 0;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);

    iframe {
        width: 100%;
    }
}

// Bottom Bar
.footer-bottom {
    background: rgba(0, 0, 0, 0.3);

    .bottom-content {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 24px 0;
        gap: 20px;

        @media (max-width: 640px) {
            flex-direction: column;
            text-align: center;
        }

        .copyright,
        .developer {
            font-size: 14px;
            color: rgba(255, 255, 255, 0.6);
            margin: 0;

            a {
                color: #4CAF50;
                text-decoration: none;
                transition: color 0.3s ease;

                &:hover {
                    color: #66BB6A;
                }
            }
        }
    }
}
</style>