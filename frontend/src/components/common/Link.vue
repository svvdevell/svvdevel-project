<template>
    <a :href="url" class="link">{{ text }}</a>
</template>
<script setup>

const props = defineProps({
    url: {
        type: String,
        required: true
    },
    text: {
        type: String,
        required: true
    }
})
</script>
<style lang="scss" scoped>
.link {
    background: #aa3535;
    padding: 13px 28px;
    gap: 10px;
    border-radius: 12px;
    font-weight: 600;
    font-style: Semi Bold;
    font-size: 25px;
    line-height: 100%;
    letter-spacing: 0%;
    text-decoration: none;
    cursor: pointer;
    color: #FFF;

    &:hover {
        opacity: 0.7;
    }
}

@media (max-width: 960px) {
    .link {
        font-size: 16px;
    }
}
</style>