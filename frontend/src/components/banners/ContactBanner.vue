<template>
    <section class="container">
        <div class="text">
            <h2>Download mobile app</h2>
            <p>Imperdiet ut tristique viverra nunc. Ultrices orci vel auctor cursus turpis nibh placerat massa.
                Fermentum urna ut at et in. Turpis aliquet cras hendrerit enim condimentum. Condimentum interdum risus
                bibendum urna</p>
            <div class="links">
                <a href=""></a>
                <a href=""></a>
            </div>
        </div>
        <div class="image">
            <img src="../../assets/images/phones.png" alt="">
        </div>
    </section>
</template>
<script setup>

</script>
<style lang="scss" scoped>
.container {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 40px;
    align-items: center;

    & .text {
        display: flex;
        flex-direction: column;
        gap: 40px;

        & h2 {
            font-family: Work Sans;
            font-weight: 700;
            font-style: Bold;
            font-size: 50px;
            leading-trim: NONE;
            line-height: 100%;
            letter-spacing: 0%;
        }

        & p {
            font-family: Work Sans;
            font-weight: 400;
            font-style: Regular;
            font-size: 16px;
            leading-trim: NONE;
            line-height: 24px;
            letter-spacing: 0%;
            color: #00000099;
            max-width: 536px;
        }
    }

    & .image {
        width: 100%;

        & img {
            width: 100%;
        }
    }
}
</style>