<template>
    <div class="text-banner">
        <h1>Чому обирають нас для швидкого викупу авто?</h1>
        <h3>Lorem ipsum dolor, sit amet consectetur adipisicing elit. Est ipsum tempore minus odit nihil dolor voluptas
            et ullam odio, excepturi, pariatur magni laborum aut quibusdam ipsa nostrum iusto ab ipsam.</h3>
        <div class="links">
            <Link :url="'tel:23423423432'" :text="'+ (380) 50 08 04 200'" />
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

    & h3 {
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
</style>