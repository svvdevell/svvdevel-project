<template>
    <section class="container">
        <div>
            <img src="../../assets/icons/location.svg" alt="Icon">
            <h3>Доступність</h3>
            <p>Виїзд у будь-яку точку Одеси та області. Оцінка та оформлення на місці</p>
        </div>
        <div>
            <img src="../../assets/icons/car.svg" alt="Icon">
            <h3>Комфорт</h3>
            <p>Швидка та зручна процедура продажу без зайвих паперів та очікування</p>
        </div>
        <div>
            <img src="../../assets/icons/wallet.svg" alt="Icon">
            <h3>Вигода</h3>
            <p>Справедлива ціна за ваш автомобіль з миттєвою виплатою готівкою</p>
        </div>
    </section>
</template>
<script setup>

</script>
<style lang="scss" scoped>
.container {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    gap: 20px;
    width: 100%;
    padding: 50px 0;

    & div {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        gap: 20px;
        width: 100%;
        max-width: 357px;

        & img {
            width: 64px;
            height: 64px;
            object-fit: contain;
        }

        & h3 {
            font-family: Work Sans;
            font-weight: 600;
            font-style: SemiBold;
            font-size: 24px;
            line-height: 100%;
            letter-spacing: 0%;
            text-align: center;
        }

        & p {
            font-weight: 400;
            font-style: Regular;
            font-size: 16px;
            leading-trim: NONE;
            line-height: 24px;
            letter-spacing: 0%;
            text-align: center;
        }
    }
}

@media (max-width: 768px) {
    .container {
        grid-template-columns: repeat(1, minmax(0, 1fr));
        align-items: center;
    }
}
</style>