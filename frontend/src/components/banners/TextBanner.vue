<template>
    <div class="text-banner">
        <h1>Чому обирають нас для швидкого викупу авто?</h1>
        <div class="text-wrap">
            <h3>Гроші одразу</h3>
            <p> – розрахунок в день звернення готівкою або на картку.</p>
        </div>
        <div class="text-wrap">
            <h3>Чесна оцінка</h3>
            <p> – ми пропонуємо реальну ринкову ціну без прихованих платежів.</p>
        </div>
        <div class="text-wrap">
            <h3>Жодних турбот</h3>
            <p> – беремо на себе всі документи та юридичний супровід.</p>
        </div>
        <div class="text-wrap">
            <h3>Просто і швидко</h3>
            <p> – достатньо заповнити заявку, і вже за 15 хвилин ви знаєте ціну.</p>
        </div>
        <div class="links">
            <a href="tel:23423423432">+ (380) 50 08 04 200</a>
            <a href="">Написати нам</a>
        </div>
    </div>
</template>
<script setup>

</script>

<style lang="scss" scoped>
.text-banner {
    display: flex;
    align-items: flex-start;
    flex-direction: column;
    justify-content: center;
    gap: 30px;

    & h1 {
        font-size: clamp(35px, 3vw, 55px);
        font-weight: 500;
        margin: 0;
        text-transform: uppercase;
        color: #FFF;
        text-align: left;
    }

    & .text-wrap {
        display: flex;

        & h3 {
            display: contents;
            font-size: 24px;
            font-weight: 500;
            margin: 0;
            text-align: center;
            color: #FFF;
        }

        & p {
            display: contents;
            text-align: center;
            font-size: 20px;
            font-weight: 400;
            margin: 0;
            color: #FFF;
            font-style: italic;
        }
    }

    & .links {
        display: flex;
        align-items: center;
        gap: 20px;
    }



    & a {
        margin: 0;
        width: 100%;
        max-width: 400px;
        text-align: center;
        background-color: #FFF;
        border-radius: 20px;
        color: #000;
        font-size: 16px;
        font-weight: 300;
        padding: 20px;
        text-wrap: nowrap;
    }
}
</style>