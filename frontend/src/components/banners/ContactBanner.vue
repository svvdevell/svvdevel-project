<template>
    <section class="container">
        <div class="image">
            <img src="../../assets/images/park-photo-3.jpg" alt="">
        </div>
        <div class="text">
            <h2>Download mobile app</h2>
            <p>Imperdiet ut tristique viverra nunc. Ultrices orci vel auctor cursus turpis nibh placerat massa.
                Fermentum urna ut at et in. Turpis aliquet cras hendrerit enim condimentum. Condimentum interdum risus
                bibendum urna</p>
            <h3>Дзвони!</h3>
            <div class="links">
                <a href="tel:0734080999"><img src="../../assets/icons/phone.svg" alt="">+38 (073) 40 80 999</a>
                <a href="tel:0631653737"><img src="../../assets/icons/phone.svg" alt="">+38 (063) 16 53 737</a>
                <a href="tel:0963313331"><img src="../../assets/icons/phone.svg" alt="">+38 (096) 33 13 331</a>
            </div>
            <h3>Пиши!</h3>
            <div class="links">
                <a href="https://t.me/trewqxm"><img src="../../assets/icons/telegram.svg" alt="">Менеджер у Телеграмі</a>
                <a href="t.me/eleganceauto_odessa"><img src="../../assets/icons/telegram.svg" alt="">Наш Телеграм канал</a>
            </div>
            <h3>Дивись!</h3>
            <div class="links">
                <a href="https://www.instagram.com/elegance_auto_od/"><img src="../../assets/icons/instagram.svg" alt="">Instagram</a>
                <a href="https://auto.ria.com/uk/pro/seller-4625376.html"><img src="../../assets/icons/auto-ria.png" alt="" class="bordered">Auto Ria</a>
                <a href="https://www.google.com/maps?q=%D0%90%D0%B2%D1%82%D0%BE%D1%81%D0%B0%D0%BB%D0%BE%D0%BD+%22Elegance+Auto%22,+%D0%BF%D1%80%D0%BE%D1%81%D0%BF%D0%B5%D0%BA%D1%82+%D0%9D%D0%B5%D0%B1%D0%B5%D1%81%D0%BD%D0%BE%D1%97+%D0%A1%D0%BE%D1%82%D0%BD%D1%96,+100%5C1,+%D0%9E%D0%B4%D0%B5%D1%81%D0%B0,+%D0%9E%D0%B4%D0%B5%D1%81%D1%8C%D0%BA%D0%B0+%D0%BE%D0%B1%D0%BB%D0%B0%D1%81%D1%82%D1%8C,+65000&ftid=0x40c7cdf97dcee44f:0x6a2762df677153ae&entry=gps&lucs=,94224825,94227247,94227248,94231188,47071704,47069508,94218641,94203019,47084304,94208458,94208447&g_ep=CAISEjI1LjE0LjEuNzQyMzk4MTI0MBgAIIgnKmMsOTQyMjQ4MjUsOTQyMjcyNDcsOTQyMjcyNDgsOTQyMzExODgsNDcwNzE3MDQsNDcwNjk1MDgsOTQyMTg2NDEsOTQyMDMwMTksNDcwODQzMDQsOTQyMDg0NTgsOTQyMDg0NDdCAlVB&skid=c399e3a0-b7c8-4cde-b6a7-fd7cfce4551f&g_st=com.google.maps.preview.copy">
                    <img src="../../assets/icons/geo.svg" alt=""> Google Maps</a>
            </div>
        </div>
    </section>
</template>
<script setup>

</script>
<style lang="scss" scoped>

.bordered {
    border-radius: 50%;
}

.container {
    display: flex;
    flex-direction: column;
    width: 100%;
    gap: 40px;
    align-items: center;

    & .text {
        display: flex;
        flex-direction: column;
        gap: 40px;
        width: 100%;
        align-items: center;

        & h2 {
            font-family: Work Sans;
            font-weight: 700;
            font-style: Bold;
            font-size: 50px;
            leading-trim: NONE;
            line-height: 100%;
            letter-spacing: 0%;
        }

        & h3 {
            font-family: Work Sans;
            font-weight: 700;
            font-style: Bold;
            font-size: 50px;
            leading-trim: NONE;
            line-height: 100%;
            letter-spacing: 0%;
            text-transform: uppercase;
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
            text-align: center;
        }
    }

    & .image {
        width: 100%;

        & img {
            width: 100%;
            object-fit: contain;
            border-radius: 20px;
        }
    }
}

.links {
    display: flex;
    align-items: center;
    gap: 20px;
    justify-content: space-between;
    width: 100%;

    & a {
        width: 100%;
        background: #aa3535;
        padding: 13px 28px;
        gap: 10px;
        border-radius: 12px;
        font-family: Inter;
        font-weight: 600;
        font-style: Semi Bold;
        font-size: 16px;
        line-height: 100%;
        letter-spacing: 0%;
        text-decoration: none;
        cursor: pointer;
        color: #FFF;
        text-align: center;
        display: flex;
        align-items: center;
        gap: 20px;

        & img {
            width: 64px;
            height: 64px;
        }
    }
}
</style>