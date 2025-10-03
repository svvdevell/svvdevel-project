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
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 40px;
    width: 100%;
    background: #aa3535;
    border-radius: 40px;
    margin: 50px 0;
    padding: 78px 72px;

}
</style>