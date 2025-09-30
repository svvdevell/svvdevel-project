<template>
    <section class="banner">
        <TextBanner />
        <MainForm />
    </section>
</template>

<script setup>
import MainForm from '@/components/forms/MainForm.vue'
import TextBanner from '@/components/banners/TextBanner.vue';

</script>

<style lang="scss" scoped>
.banner {
    width: 100%;
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 40px;
    width: 100%;
    padding: 30px;
}
</style>