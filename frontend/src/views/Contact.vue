<template>
  <BigForm />
  <FormBanner />
</template>
<script setup>
import BigForm from '@/components/forms/BigForm.vue';
import FormBanner from '@/components/banners/FormBanner.vue';
import { onMounted } from 'vue';
import { useSeo } from '@/composables/useSeo';

const { setMeta } = useSeo();

onMounted(() => {
    setMeta({
        title: 'Контакти - Викуп авто Одеса | Elegance Auto',
        description: 'Терміновий викуп автомобілів в Одесі та області. Телефонуйте зараз для оцінки вашого авто. Виїзд оцінювача безкоштовно. Выкуп авто в Одессе - звоните!',
        keywords: 'викуп авто Одеса контакти, викуп авто телефон Одеса, выкуп авто Одесса телефон, продать машину Одесса, автовыкуп контакты, адреса автосалону Одеса, терміновий викуп авто',
        url: 'https://eleganceauto.od.ua/contact',
        ogImage: 'https://eleganceauto.od.ua/images/og-contact.jpg'
    });
});
</script>