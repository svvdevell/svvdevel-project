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
                    <img src="../../assets/images/salon.jpg" alt="Slide 1">
                </SwiperSlide>
                <SwiperSlide>
                    <img src="../../assets/images/park_image_12" alt="Slide 2">
                </SwiperSlide>
                <SwiperSlide>
                    <img src="../../assets/images/park_image_13" alt="Slide 3">
                </SwiperSlide>
                <SwiperSlide>
                    <img src="../../assets/images/park_image_7" alt="Slide 4">
                </SwiperSlide>
                <SwiperSlide>
                    <img src="../../assets/images/park_image_11" alt="Slide 5">
                </SwiperSlide>
                <SwiperSlide>
                    <img src="../../assets/images/park_image_8" alt="Slide 6">
                </SwiperSlide>
                <SwiperSlide>
                    <img src="../../assets/images/park_image_10" alt="Slide 7">
                </SwiperSlide>
            </Swiper>
        </div>
        <div class="text">
            <h3>Швидке оформлення</h3>
            <p>Всі документи готуємо протягом 1 години. Мінімум паперової тяганини та бюрократії</p>

            <h3>Виїзд по місту</h3>
            <p>Наш фахівець приїде в будь-яку точку Одеси для огляду та оцінки автомобіля</p>

            <h3>Викуп за готівку</h3>
            <p>Миттєва оплата готівкою або переказом на картку одразу після укладення угоди</p>

            <h3>Без прихованих комісій</h3>
            <p>Озвучена ціна - фінальна. Ніяких додаткових платежів та несподіванок при оформленні</p>

            <h3>Будь-які марки авто</h3>
            <p>Викуповуємо легкові, позашляховики, мінівени та комерційний транспорт всіх брендів</p>

            <h3>Досвідчені експерти</h3>
            <p>Наша команда має понад 15 років досвіду роботи на автомобільному ринку Одеси</p>

            <h3>Чесна оцінка</h3>
            <p>Визначаємо реальну ринкову вартість на основі стану авто та актуальних цін</p>

            <h3>Допомога з документами</h3>
            <p>Беремо на себе всі питання з переоформленням та зняттям автомобіля з обліку</p>
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
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 50px;
    padding: 50px 0;

    & .text {
        display: flex;
        flex-direction: column;
        gap: 20px;
        justify-content: center;

        & h3 {
            font-weight: 600;
            font-style: SemiBold;
            font-size: 26px;
            leading-trim: NONE;
            line-height: 100%;
            letter-spacing: 0%;
        }

        & p {
            font-weight: 400;
            font-style: Regular;
            font-size: 20px;
            leading-trim: NONE;
            line-height: 100%;
            letter-spacing: 0%;
            color: rgb(255 255 255 / 72%);
        }
    }

    & .image {
        width: 100%;

        & .swiper-container {
            border-radius: 20px;
            overflow: hidden;
        }

        & img {
            width: 100%;
            height: 600px;
            object-fit: cover;
            border-radius: 20px;
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
        grid-template-columns: repeat(1, minmax(0, 1fr));

        & .image {
            & img {
                height: 300px;
            }
        }

        & .text {
            & h3 {
                font-size: 20px;
            }

            & p {
                font-size: 16px;
            }
        }
    }

    :deep(.swiper-pagination) {
        bottom: 10px !important;
    }
}
</style>