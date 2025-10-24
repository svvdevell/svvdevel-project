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
    title: 'Контакти - Elegance Auto',
    description: 'Зв\'яжіться з нами для викупу автомобіля',
    url: 'https://eleganceauto.od.ua/contact'
  });
});
</script>