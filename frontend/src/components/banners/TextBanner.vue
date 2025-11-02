<template>
    <div class="text-banner">
        <h1>Чому обирають нас для швидкого викупу авто?</h1>
        <h2>Викупимо ваш автомобіль за найвищою ринковою ціною протягом 1 дня. Працюємо з будь-якими марками та моделями, в будь-якому стані. Оцінка безкоштовна, гроші отримаєте відразу після укладення угоди. Всі документи оформимо самостійно.</h2>
        <div class="links">
            <Link :url="'tel:0734080999'" :text="'+38 (073) 40 80 999'" />
        </div>
    </div>
</template>
<script setup>
import Link from '@/components/common/Link.vue';

</script>

<style lang="scss" scoped>
.text-banner {
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 30px;
    max-width: 686px;
    width: 100%;

    & h1 {
        font-family: Work Sans;
        font-weight: 700;
        font-style: Bold;
        font-size: 60px;
        line-height: 100%;
        letter-spacing: 0%;
        color: #FFF;
    }

    & h2 {
        font-family: Work Sans;
        font-weight: 400;
        font-style: Regular;
        font-size: 16px;
        line-height: 100%;
        letter-spacing: 0%;
        color: #FFF;
    }

    & .links {
        display: flex;
        align-items: center;
        gap: 20px;
    }
}

@media (max-width: 768px) {
    .text-banner {
        gap: 30px;
        align-items: center;
        padding: 0 16px;

        & h1 {
            text-align: center;
            font-size: 24px;
        }

        & h3 {
            text-align: center;
            font-size: 14px;
        }

    }
}
</style>