<template>
    <section class="banner">
        <!-- <h1>Швидкий викуп авто</h1>
        <h3>Швидко, зручно, надійно </h3>
        <div class="links">
            <a href="tel:23423423432">+ (380) 50 08 04 200</a>
            <a href="">Написати нам</a>
        </div>
        <p>Оцінимо Ваш автомобіль по телефону і запропонуємо найкращу ціну в Одесі!</p> -->
        <FormInput />
    </section>
</template>

<script setup>
import FormInput from '../components/FormInput.vue'

</script>

<style lang="scss" scoped>
.banner {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 40px;
    align-items: center;
    padding: 30px;
    // backdrop-filter: blur(8px);
    // box-shadow: 20px 20px 40px -6px rgba(0,0,0,0.2);
    // border-radius: 20px;

    & h1 {
        font-size: 72px;
        font-weight: 500;
        margin: 0;
        text-transform: uppercase;
    }

    & h3 {
        font-size: 40px;
        font-weight: 400;
        margin: 0;
    }

    & .links {
        display: flex;
        align-items: center;
        gap: 20px;

    }

    & p {
        font-size: 16px;
        font-weight: 400;
        margin: 0;
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