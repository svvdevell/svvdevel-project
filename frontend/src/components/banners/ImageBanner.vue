<template>
    <section class="container">
        <div class="image">
            <Swiper :modules="modules" :slides-per-view="1" :space-between="0" :autoplay="{
                delay: 3000,
                disableOnInteraction: false,
            }" :pagination="{
                clickable: true,
            }" :loop="true" class="swiper-container">
                <SwiperSlide>
                    <img src="../../assets/images/park-photo-3.jpg" alt="Slide 1">
                </SwiperSlide>
                <SwiperSlide>
                    <img src="../../assets/images/park_image_6.jpg" alt="Slide 2">
                </SwiperSlide>
                <SwiperSlide>
                    <img src="../../assets/images/park_image_9.jpg" alt="Slide 3">
                </SwiperSlide>
                <SwiperSlide>
                    <img src="../../assets/images/park_image_5.jpg" alt="Slide 4">
                </SwiperSlide>
                <SwiperSlide>
                    <img src="../../assets/images/park_image_13.jpg" alt="Slide 5">
                </SwiperSlide>
            </Swiper>
        </div>
        <div class="text">
            <div class="block">
                <div>
                    <h5>1</h5>
                    <h4>Викуп будь-яких автомобілів</h4>
                </div>
                <p>Купуємо авто будь-якої марки та року випуску. Працюємо з легковими та комерційними транспортними
                    засобами в будь-якому стані</p>
            </div>
            <div class="block">
                <div>
                    <h5>2</h5>
                    <h4>Оцінка та оплата за 30 хвилин</h4>
                </div>
                <p>Швидка професійна оцінка вашого автомобіля. Миттєва виплата готівкою або переказом на картку одразу
                    після угоди</p>
            </div>
            <div class="block">
                <div>
                    <h5>3</h5>
                    <h4>Максимальна ціна за ваше авто</h4>
                </div>
                <p>Пропонуємо найвищі ціни на ринку автовикупу в Одесі. Аналізуємо ринкову вартість та гарантуємо чесну
                    оцінку</p>
            </div>
            <div class="block">
                <div>
                    <h5>4</h5>
                    <h4>Чесна та прозора угода</h4>
                </div>
                <p>Справедлива ринкова оцінка без прихованих комісій. Повне юридичне супроводження угоди. Безпечне
                    оформлення всіх документів</p>
            </div>
        </div>
    </section>
</template>

<script setup>
import { Swiper, SwiperSlide } from 'swiper/vue';
import { Autoplay, Pagination } from 'swiper/modules';
import 'swiper/css';
import 'swiper/css/pagination';

const modules = [Autoplay, Pagination];
</script>

<style lang="scss" scoped>
.container {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 60px 0;

    & .image {
        width: 100%;

        & .swiper-container {
            border-radius: 20px;
            overflow: hidden;
        }

        & img {
            width: 100%;
            height: 500px;
            object-fit: cover;
            border-radius: 20px;
        }
    }
}

.text {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 20px;

    & .block {
        display: flex;
        flex-direction: column;
        gap: 20px;

        & div {
            display: flex;
            gap: 20px;
            align-items: center;

            & h5 {
                border-radius: 50%;
                background-color: #aa3535;
                font-family: Work Sans;
                font-weight: 600;
                font-style: SemiBold;
                font-size: 16px;
                leading-trim: NONE;
                line-height: 100%;
                letter-spacing: 0%;
                width: 35px;
                height: 35px;
                color: #FFF;
                display: flex;
                align-items: center;
                justify-content: center;
                padding-right: 2px;
            }

            & h4 {
                font-family: Work Sans;
                font-weight: 600;
                font-style: SemiBold;
                font-size: 20px;
                leading-trim: NONE;
                line-height: 100%;
                letter-spacing: 0%;
            }
        }

        & p {
            font-family: Work Sans;
            font-weight: 400;
            font-style: Regular;
            font-size: 16px;
            leading-trim: NONE;
            line-height: 100%;
            letter-spacing: 0%;
            max-width: 85%;
        }
    }
}

/* Стили для пагинации (точек) */
:deep(.swiper-pagination) {
    bottom: 20px !important;
}

:deep(.swiper-pagination-bullet) {
    width: 10px;
    height: 10px;
    background: #fff;
    opacity: 0.5;
    transition: all 0.3s;
}

:deep(.swiper-pagination-bullet-active) {
    background: #aa3535;
    opacity: 1;
    width: 30px;
    border-radius: 5px;
}

@media (max-width: 768px) {
    .container {
        & .image {
            & img {
                height: 300px;
            }
        }
    }

    .text {
        grid-template-columns: repeat(1, minmax(0, 1fr));

        & .block {
            & p {
                max-width: 100%;
            }
        }
    }

    :deep(.swiper-pagination) {
        bottom: 10px !important;
    }
}
</style>