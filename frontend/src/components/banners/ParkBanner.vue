<template>
    <section class="container">
        <div class="image">
            <img src="../../assets/images/park_image_1.webp" alt="">
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

        & img {
            width: 100%;
            object-fit: contain;
            border-radius: 20px;
        }
    }
}

@media (max-width: 768px) {
    .container {
        grid-template-columns: repeat(1, minmax(0, 1fr));

        & .text {

            & h3 {
                font-size: 20px;
            }

            & p {
                font-size: 16px;
            }
        }
    }

}
</style>