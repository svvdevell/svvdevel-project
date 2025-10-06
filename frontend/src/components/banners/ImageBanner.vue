<template>
    <section class="container">
        <div class="image">
            <img src="../../assets/images/park-photo-3.jpg" alt="">
        </div>
        <div class="text">
            <div class="block">
                <div>
                    <h5>1</h5>
                    <h4>Erat at semper</h4>
                </div>
                <p>Non amet fermentum est in enim at sit ullamcorper. Sit elementum rhoncus nullam feugiat. Risus sem
                    fermentum</p>
            </div>
            <div class="block">
                <div>
                    <h5>2</h5>
                    <h4>Urna nec vivamus risus duis arcu </h4>
                </div>
                <p>Aliquam adipiscing velit semper morbi. Purus non eu cursus porttitor tristique et gravida. Quis nunc
                    interdum gravida ullamcorper</p>
            </div>
            <div class="block">
                <div>
                    <h5>3</h5>
                    <h4>Lobortis euismod imperdiet tempus </h4>
                </div>
                <p>Viverra scelerisque mauris et nullam molestie et. Augue adipiscing praesent nisl cras nunc luctus
                    viverra nisi</p>
            </div>
            <div class="block">
                <div>
                    <h5>4</h5>
                    <h4>Cras nulla aliquet nam eleifend amet et</h4>
                </div>
                <p>Aliquam adipiscing velit semper morbi. Purus non eu cursus porttitor tristique et gravida. Quis nunc
                    interdum gravida ullamcorper sed integer. Quisque eleifend tincidunt vulputate libero </p>
            </div>
        </div>
    </section>
</template>
<script setup>

</script>
<style lang="scss" scoped>
.container {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 60px 0;

    & .image {
        width: 100%;

        & img {
            width: 100%;
            object-fit: contain;
            border-radius: 20px;
        }
    }
}

.text {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 20px;

    & .block {
        display: flex;
        flex-direction: column;
        gap: 20px;

        & div {
            display: flex;
            gap: 20px;
            align-items: center;

            & h5 {
                border-radius: 50%;
                background-color: #aa3535;
                font-family: Work Sans;
                font-weight: 600;
                font-style: SemiBold;
                font-size: 16px;
                leading-trim: NONE;
                line-height: 100%;
                letter-spacing: 0%;
                width: 35px;
                height: 35px;
                color: #FFF;
                display: flex;
                align-items: center;
                justify-content: center;
                padding-right: 2px;
            }

            & h4 {
                font-family: Work Sans;
                font-weight: 600;
                font-style: SemiBold;
                font-size: 20px;
                leading-trim: NONE;
                line-height: 100%;
                letter-spacing: 0%;
            }
        }

        & p {
            font-family: Work Sans;
            font-weight: 400;
            font-style: Regular;
            font-size: 16px;
            leading-trim: NONE;
            line-height: 100%;
            letter-spacing: 0%;
            color: #00000099;
        }
    }
}
</style>