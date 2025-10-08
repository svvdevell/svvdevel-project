<template>
    <section class="container">
        <h2>Наші досягнення в цифрах</h2>
        <p>Ми працюємо чесно та прозоро, тому наші клієнти повертаються до нас знову. Довіра та професіоналізм - основа нашого бізнесу</p>
        <div class="block">
            <div class="text">
                <img src="../../assets/icons/year.svg" alt="">
                <div>
                    <h4>540+</h4>
                    <span>Авто викуплено</span>
                </div>
            </div>
            <div class="text">
                <img src="../../assets/icons/car-orange.svg" alt="">
                <div>
                    <h4>10k+</h4>
                    <span>Задоволених клієнтів</span>
                </div>
            </div>
            <div class="text">
                <img src="../../assets/icons/customers.svg" alt="">
                <div>
                    <h4>15+</h4>
                    <span>Років досвіду</span>
                </div>
            </div>
            <div class="text">
                <img src="../../assets/icons/speed.svg" alt="">
                <div>
                    <h4>20m+</h4>
                    <span>Пройдено кілометрів</span>
                </div>
            </div>
        </div>
    </section>
</template>
<script setup>

</script>
<style lang="scss" scoped>
.container {
    background: #aa3535;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;
    gap: 40px;
    width: 100%;
    background: #aa3535;
    border-radius: 40px;
    margin: 50px 0;
    padding: 78px 72px;

    & h2 {
        font-family: Work Sans;
        font-weight: 700;
        font-style: Bold;
        font-size: 50px;
        leading-trim: NONE;
        line-height: 100%;
        letter-spacing: 0%;
        text-align: center;
        text-transform: capitalize;
        color: #FFF;
    }

    & p {
        font-family: Work Sans;
        font-weight: 400;
        font-style: Regular;
        font-size: 16px;
        leading-trim: NONE;
        line-height: 100%;
        letter-spacing: 0%;
        text-align: center;
        color: #FFF;
        max-width: 700px;
    }
}

.block {
    display: grid;
    grid-template-columns: repeat(4, minmax(0, 1fr));
    gap: 20px;
    width: 100%;

    & .text {
        display: flex;
        align-items: center;
        gap: 40px;
        padding: 16px;
        background-color: #FFF;
        border-radius: 20px;

        & img {
            width: 72px;
            height: 68px;
        }

        & div {
            display: flex;
            flex-direction: column;
            align-items: flex-start;
            gap: 8px;

            & h4 {
                font-family: Work Sans;
                font-weight: 700;
                font-style: Bold;
                font-size: 24px;
                leading-trim: NONE;
                line-height: 100%;
                letter-spacing: 0%;
                text-align: center;
                text-transform: capitalize;
            }

            & span {
                font-family: Work Sans;
                font-weight: 600;
                font-style: SemiBold;
                font-size: 16px;
                leading-trim: NONE;
                line-height: 100%;
                letter-spacing: 0%;
                text-transform: capitalize;
                color: #00000099;
            }
        }
    }
}

@media (max-width: 768px) {
    .container {
        padding: 30px;

        & h2 {
            font-size: 30px;
        }
    }
    .block {
        grid-template-columns: repeat(1, minmax(0, 1fr));
    }
}
</style>