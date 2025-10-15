<template>
    <section class="container">
        <div class="image">
            <img src="../../assets/images/park-photo-3.jpg" alt="">
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

</script>
<style lang="scss" scoped>
.container {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 60px 0;

    & .image {
        width: 100%;

        & img {
            width: 100%;
            object-fit: contain;
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

@media (max-width: 768px) {
    .text {
        grid-template-columns: repeat(1, minmax(0, 1fr));

        & .block {
            & p {
                max-width: 100%;
            }
        }
    }
}
</style>