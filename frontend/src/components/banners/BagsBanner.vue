<template>
    <section class="container">
        <div>
            <img src="../../assets/icons/location.svg" alt="Icon">
            <h3>Availability</h3>
            <p>Diam tincidunt tincidunt erat at semper fermentum. Id ultricies quis</p>
        </div>
        <div>
            <img src="../../assets/icons/car.svg" alt="Icon">
            <h3>Comfort</h3>
            <p>Gravida auctor fermentum morbi vulputate ac egestas orcietium convallis</p>
        </div>
        <div>
            <img src="../../assets/icons/wallet.svg" alt="Icon">
            <h3>Savings</h3>
            <p>Pretium convallis id diam sed commodo vestibulum lobortis volutpat</p>
        </div>
    </section>
</template>
<script setup>

</script>
<style lang="scss" scoped>
.container {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    gap: 20px;
    width: 100%;

    & div {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        gap: 20px;
        width: 100%;
        max-width: 357px;

        & img {
            width: 64px;
            height: 64px;
            object-fit: contain;
        }

        & h3 {
            font-family: Work Sans;
            font-weight: 600;
            font-style: SemiBold;
            font-size: 24px;
            line-height: 100%;
            letter-spacing: 0%;
            text-align: center;
        }

        & p {
            font-family: Inter;
            font-weight: 400;
            font-style: Regular;
            font-size: 16px;
            leading-trim: NONE;
            line-height: 24px;
            letter-spacing: 0%;
            text-align: center;
        }
    }
}
</style>