<template>
    <div class="blog-page">
        <div class="container">
            <!-- Hero Section -->
            <div class="blog-hero">
                <h1>Корисні поради</h1>
                <p>Все, що потрібно знати про купівлю та продаж автомобілів</p>
            </div>

            <!-- Categories Filter -->
            <div class="categories">
                <button :class="['category-btn', { active: activeCategory === 'all' }]"
                    @click="filterByCategory('all')">
                    Всі статті
                </button>
                <button :class="['category-btn', { active: activeCategory === 'buying' }]"
                    @click="filterByCategory('buying')">
                    Купівля авто
                </button>
                <button :class="['category-btn', { active: activeCategory === 'selling' }]"
                    @click="filterByCategory('selling')">
                    Продаж авто
                </button>
                <button :class="['category-btn', { active: activeCategory === 'tips' }]"
                    @click="filterByCategory('tips')">
                    Корисні поради
                </button>
            </div>

            <!-- Articles Grid -->
            <div class="articles-grid">
                <article v-for="article in filteredArticles" :key="article.id" class="article-card"
                    @click="openArticle(article.id)">
                    <div class="article-image">
                        <img :src="article.image" :alt="article.title">
                        <div class="article-category">{{ getCategoryName(article.category) }}</div>
                    </div>
                    <div class="article-content">
                        <h3>{{ article.title }}</h3>
                        <p>{{ article.excerpt }}</p>
                        <div class="article-footer">
                            <span class="read-time">
                                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24"
                                    fill="none" stroke="currentColor" stroke-width="2">
                                    <circle cx="12" cy="12" r="10"></circle>
                                    <polyline points="12 6 12 12 16 14"></polyline>
                                </svg>
                                {{ article.readTime }} хв читання
                            </span>
                            <span class="read-more">Читати далі →</span>
                        </div>
                    </div>
                </article>
            </div>

            <!-- Empty State -->
            <div v-if="filteredArticles.length === 0" class="empty-state">
                <p>Статей в цій категорії поки немає</p>
            </div>
        </div>

        <!-- Article Modal -->
        <Teleport to="body">
            <Transition name="modal">
                <div v-if="selectedArticle" class="modal-overlay" @click="closeArticle">
                    <div class="modal-content" @click.stop>
                        <button class="modal-close" @click="closeArticle">×</button>

                        <div class="modal-header">
                            <img :src="selectedArticle.image" :alt="selectedArticle.title">
                            <div class="modal-category">{{ getCategoryName(selectedArticle.category) }}</div>
                        </div>

                        <div class="modal-body">
                            <h2>{{ selectedArticle.title }}</h2>
                            <div class="modal-meta">
                                <span class="read-time">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24"
                                        fill="none" stroke="currentColor" stroke-width="2">
                                        <circle cx="12" cy="12" r="10"></circle>
                                        <polyline points="12 6 12 12 16 14"></polyline>
                                    </svg>
                                    {{ selectedArticle.readTime }} хв читання
                                </span>
                            </div>

                            <div class="article-text" v-html="selectedArticle.content"></div>
                        </div>
                    </div>
                </div>
            </Transition>
        </Teleport>
    </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useSeo } from '@/composables/useSeo';

const { setMeta } = useSeo();

onMounted(() => {
    setMeta({
        title: 'Корисні поради про купівлю та продаж авто - Elegance Auto',
        description: 'Експертні поради про купівлю та продаж автомобілів. Як правильно перевірити вживане авто, оформити документи, не попастися на шахраїв. Все про авто ринок України.',
        keywords: 'поради купівля авто, як купити авто, як продати авто, перевірка вживаного авто, оформлення авто, документи на авто, викуп авто поради, продаж автомобіля Одеса, авто блог',
        url: 'https://eleganceauto.od.ua/blog',
        ogImage: 'https://eleganceauto.od.ua/images/og-blog.jpg'
    });
});
const activeCategory = ref('all')
const selectedArticle = ref(null)

const articles = ref([
    {
        id: 1,
        title: 'Як правильно купити вживаний автомобіль',
        excerpt: 'Покрокова інструкція для безпечної купівлі автомобіля з пробігом. Що перевіряти та на що звертати увагу.',
        category: 'buying',
        readTime: 8,
        image: 'https://images.unsplash.com/photo-1549317661-bd32c8ce0db2?w=800&q=80',
        content: `
            <h3>1. Перевірка документів</h3>
            <p>Перш ніж оглядати автомобіль, обов'язково перевірте всі документи:</p>
            <ul>
                <li>Свідоцтво про реєстрацію транспортного засобу</li>
                <li>Технічний паспорт</li>
                <li>Сервісну книжку</li>
                <li>Договір купівлі-продажу попереднього власника</li>
            </ul>

            <h3>2. Візуальний огляд</h3>
            <p>Уважно оглядайте кузов автомобіля при денному світлі:</p>
            <ul>
                <li>Перевірте рівномірність зазорів між панелями кузова</li>
                <li>Подивіться на стан фарби - чи немає різниці у відтінках</li>
                <li>Перевірте стан шин та дисків</li>
                <li>Оцініть стан фар та ліхтарів</li>
            </ul>

            <h3>3. Перевірка двигуна</h3>
            <p>Двигун - серце автомобіля. Що потрібно перевірити:</p>
            <ul>
                <li>Запустіть двигун на холодну - він має запускатися легко</li>
                <li>Прислухайтеся до звуків роботи двигуна</li>
                <li>Перевірте рівень масла та охолоджуючої рідини</li>
                <li>Подивіться, чи немає підтікань</li>
            </ul>

            <h3>4. Тест-драйв</h3>
            <p>Обов'язково проведіть тест-драйв:</p>
            <ul>
                <li>Перевірте роботу трансмісії</li>
                <li>Оцініть роботу підвіски</li>
                <li>Перевірте гальма</li>
                <li>Протестуйте електроніку та комфорт</li>
            </ul>

            <h3>5. Діагностика на СТО</h3>
            <p>Перед покупкою обов'язково проведіть діагностику у перевіреного майстра:</p>
            <ul>
                <li>Комп'ютерна діагностика всіх систем</li>
                <li>Перевірка товщини лакофарбового покриття</li>
                <li>Огляд підвіски на підйомнику</li>
                <li>Перевірка геометрії кузова</li>
            </ul>

            <h3>6. Перевірка історії авто</h3>
            <p>Скористайтеся онлайн-сервісами для перевірки:</p>
            <ul>
                <li>Участь у ДТП</li>
                <li>Кількість власників</li>
                <li>Пробіг</li>
                <li>Знаходження в заставі</li>
            </ul>

            <p><strong>Пам'ятайте:</strong> краще витратити трохи часу на перевірку, ніж потім шкодувати про поспішне рішення!</p>
        `
    },
    {
        id: 2,
        title: 'Як швидко продати автомобіль за хорошою ціною',
        excerpt: 'Ефективні стратегії продажу вашого авто. Підготовка, ціноутворення та успішні переговори з покупцями.',
        category: 'selling',
        readTime: 6,
        image: 'https://images.unsplash.com/photo-1492144534655-ae79c964c9d7?w=800&q=80',
        content: `
            <h3>1. Підготовка автомобіля</h3>
            <p>Перше враження - найважливіше. Підготуйте авто до продажу:</p>
            <ul>
                <li>Зробіть детальну хімчистку салону</li>
                <li>Відполіруйте кузов</li>
                <li>Усуньте дрібні подряпини та сколи</li>
                <li>Замініть розхідники (свічки, фільтри, масло)</li>
            </ul>

            <h3>2. Правильне ціноутворення</h3>
            <p>Ціна - ключовий фактор швидкого продажу:</p>
            <ul>
                <li>Проаналізуйте ринок - подивіться ціни на аналогічні авто</li>
                <li>Врахуйте реальний стан вашого автомобіля</li>
                <li>Залиште невеликий запас для торгу (5-10%)</li>
                <li>Будьте готові обґрунтувати свою ціну</li>
            </ul>

            <h3>3. Якісні фотографії</h3>
            <p>Гарні фото збільшують інтерес до оголошення:</p>
            <ul>
                <li>Фотографуйте при денному світлі</li>
                <li>Зробіть знімки з усіх ракурсів (мінімум 10 фото)</li>
                <li>Покажіть салон, двигун, багажник</li>
                <li>Сфотографуйте всі дефекти чесно</li>
            </ul>

            <h3>4. Опис оголошення</h3>
            <p>Детальний та чесний опис привертає серйозних покупців:</p>
            <ul>
                <li>Вкажіть всю базову інформацію</li>
                <li>Опишіть комплектацію</li>
                <li>Зазначте історію обслуговування</li>
                <li>Будьте чесні щодо недоліків</li>
            </ul>

            <h3>5. Де розміщувати оголошення</h3>
            <p>Використовуйте популярні майданчики:</p>
            <ul>
                <li>AUTO.RIA</li>
                <li>RST</li>
                <li>OLX</li>
                <li>Соціальні мережі та тематичні групи</li>
            </ul>

            <h3>6. Комунікація з покупцями</h3>
            <p>Як спілкуватися з потенційними покупцями:</p>
            <ul>
                <li>Відповідайте швидко на запити</li>
                <li>Будьте ввічливими та терплячими</li>
                <li>Підготуйтеся до огляду авто</li>
                <li>Погодьтеся на діагностику на СТО покупця</li>
            </ul>

            <h3>7. Безпека угоди</h3>
            <p>Важливі моменти при укладенні угоди:</p>
            <ul>
                <li>Зустрічайтеся в людних місцях</li>
                <li>Перевіряйте документи покупця</li>
                <li>Правильно оформляйте договір</li>
                <li>Отримуйте оплату безпечним способом</li>
            </ul>

            <p><strong>Важливо:</strong> чесність та прозорість - запорука швидкого продажу!</p>
        `
    },
    {
        id: 3,
        title: 'На що звернути увагу при огляді вживаного авто',
        excerpt: 'Детальний чеклист для перевірки автомобіля. Приховані дефекти та як їх виявити.',
        category: 'tips',
        readTime: 10,
        image: 'https://images.unsplash.com/photo-1486496146582-9ffcd0b2b2b7?w=800&q=80',
        content: `
            <h3>Кузов та ЛКП</h3>
            <p>Перевірка стану кузова - перший крок при огляді:</p>
            <ul>
                <li>Використовуйте товщиномір для перевірки товщини фарби</li>
                <li>Подивіться на рівномірність зазорів</li>
                <li>Перевірте стан порогів та арок</li>
                <li>Подивіться на корозію в прихованих місцях</li>
            </ul>

            <h3>Салон</h3>
            <p>Стан салону розкаже багато про власника:</p>
            <ul>
                <li>Зношеність кермо та педалей відповідає пробігу?</li>
                <li>Працюють усі кнопки та перемикачі</li>
                <li>Немає сторонніх запахів (цвілі, затхлості)</li>
                <li>Стан сидінь відповідає заявленому пробігу</li>
            </ul>

            <h3>Двигун</h3>
            <p>Серце автомобіля потребує особливої уваги:</p>
            <ul>
                <li>Чистота моторного відсіку (якщо надто чистий - підозріло)</li>
                <li>Відсутність підтікань масла</li>
                <li>Колір вихлопу при запуску</li>
                <li>Стабільність роботи на холостих</li>
            </ul>

            <h3>Ходова частина</h3>
            <p>Підвіска та гальма - безпека руху:</p>
            <ul>
                <li>Стукає при проїзді нерівностей</li>
                <li>Рівень зношення гальмівних дисків</li>
                <li>Стан амортизаторів</li>
                <li>Люфти в кульових та рульових наконечниках</li>
            </ul>

            <h3>Електроніка</h3>
            <p>Сучасні авто напичкані електронікою:</p>
            <ul>
                <li>Всі лампочки на панелі приладів працюють</li>
                <li>Немає помилок у бортовому комп'ютері</li>
                <li>Працює клімат-контроль</li>
                <li>Функціонують всі допоміжні системи</li>
            </ul>

            <h3>Документи</h3>
            <p>Юридична чистота авто:</p>
            <ul>
                <li>Співпадіння VIN-коду</li>
                <li>Наявність сервісної книжки</li>
                <li>Кількість власників</li>
                <li>Відсутність обтяжень</li>
            </ul>

            <p><strong>Порада:</strong> завжди беріть з собою досвідченого механіка або скористайтеся послугами підбору авто!</p>
        `
    },
    {
        id: 4,
        title: 'Як не потрапити на шахраїв при купівлі авто',
        excerpt: 'Популярні схеми обману покупців автомобілів та як від них захиститися.',
        category: 'buying',
        readTime: 7,
        image: 'https://images.unsplash.com/photo-1449965408869-eaa3f722e40d?w=800&q=80',
        content: `
            <h3>Скручений пробіг</h3>
            <p>Найпоширеніша схема обману:</p>
            <ul>
                <li>Перевіряйте історію обслуговування</li>
                <li>Дивіться на зношеність салону</li>
                <li>Перевірте сервісну книжку</li>
                <li>Використовуйте сервіси перевірки історії</li>
            </ul>

            <h3>Приховані ДТП</h3>
            <p>Як виявити битий автомобіль:</p>
            <ul>
                <li>Різна товщина ЛКП</li>
                <li>Нерівномірні зазори кузова</li>
                <li>Свіжа фарба на деталях</li>
                <li>Перевірка на спеціалізованому СТО</li>
            </ul>

            <h3>Утоплені автомобілі</h3>
            <p>Ознаки авто після затоплення:</p>
            <ul>
                <li>Запах цвілі в салоні</li>
                <li>Іржа в незвичних місцях</li>
                <li>Проблеми з електрикою</li>
                <li>Бруд у важкодоступних місцях</li>
            </ul>

            <h3>Підставні продавці</h3>
            <p>Як не потрапити на перекупів:</p>
            <ul>
                <li>Зустрічайтесь за адресою реєстрації</li>
                <li>Перевіряйте документи власника</li>
                <li>Просіть показати історію ремонтів</li>
                <li>Не вірте занадто вигідним пропозиціям</li>
            </ul>

            <h3>Кредитні та заставні авто</h3>
            <p>Як перевірити юридичну чистоту:</p>
            <ul>
                <li>Запит в реєстр обтяжень</li>
                <li>Перевірка через державні сервіси</li>
                <li>Запит довідки з банків</li>
                <li>Договір з нотаріальним посвідченням</li>
            </ul>

            <h3>Захист при купівлі</h3>
            <p>Як себе захистити:</p>
            <ul>
                <li>Проводьте угоду через нотаріуса</li>
                <li>Використовуйте послуги перевірки авто</li>
                <li>Не давайте завдаток незнайомцям</li>
                <li>Оформлюйте договір правильно</li>
            </ul>

            <p><strong>Пам'ятайте:</strong> якщо пропозиція виглядає надто вигідно - це пастка!</p>
        `
    },
    {
        id: 5,
        title: 'Оформлення купівлі-продажу автомобіля: покрокова інструкція',
        excerpt: 'Всі документи та процедури для легального оформлення угоди купівлі-продажу авто.',
        category: 'selling',
        readTime: 9,
        image: 'https://images.unsplash.com/photo-1450101499163-c8848c66ca85?w=800&q=80',
        content: `
            <h3>Крок 1: Підготовка документів продавця</h3>
            <p>Що потрібно мати при собі продавцю:</p>
            <ul>
                <li>Паспорт громадянина України</li>
                <li>ІПН (ідентифікаційний код)</li>
                <li>Свідоцтво про реєстрацію ТЗ</li>
                <li>Технічний паспорт</li>
            </ul>

            <h3>Крок 2: Складання договору</h3>
            <p>Договір купівлі-продажу повинен містити:</p>
            <ul>
                <li>Дані продавця та покупця</li>
                <li>Повну інформацію про автомобіль (VIN, номери)</li>
                <li>Ціну угоди</li>
                <li>Підписи обох сторін</li>
            </ul>

            <h3>Крок 3: Акт приймання-передачі</h3>
            <p>Обов'язково складіть акт:</p>
            <ul>
                <li>Опис стану автомобіля</li>
                <li>Перелік переданих предметів (ключі, документи)</li>
                <li>Показники лічильника</li>
                <li>Підтвердження передачі грошей</li>
            </ul>

            <h3>Крок 4: Зняття з обліку (якщо потрібно)</h3>
            <p>Процедура в сервісному центрі МВС:</p>
            <ul>
                <li>Заява від власника</li>
                <li>Оригінали документів</li>
                <li>Здача номерних знаків</li>
                <li>Отримання транзитів (за потреби)</li>
            </ul>

            <h3>Крок 5: Оплата</h3>
            <p>Безпечні способи розрахунку:</p>
            <ul>
                <li>Готівка при свідках</li>
                <li>Банківський переказ</li>
                <li>Акредитив</li>
                <li>Розрахунок через нотаріуса</li>
            </ul>

            <h3>Крок 6: Реєстрація на нового власника</h3>
            <p>Що робить покупець:</p>
            <ul>
                <li>Звернення в сервісний центр</li>
                <li>Подання документів</li>
                <li>Сплата послуг реєстрації</li>
                <li>Отримання нових документів</li>
            </ul>

            <h3>Крок 7: Страхування</h3>
            <p>Обов'язкове страхування:</p>
            <ul>
                <li>ОСАЦВ (обов'язкова цивільна відповідальність)</li>
                <li>Зелена картка (для виїзду за кордон)</li>
                <li>Каско (за бажанням)</li>
            </ul>

            <p><strong>Важливо:</strong> у покупця є 10 днів на реєстрацію авто після укладення договору!</p>
        `
    },
    {
        id: 6,
        title: 'Електромобілі: що потрібно знати перед покупкою',
        excerpt: 'Особливості вибору та експлуатації електричних автомобілів. Переваги, недоліки та підводні камені.',
        category: 'tips',
        readTime: 12,
        image: 'https://images.unsplash.com/photo-1593941707882-a5bba14938c7?w=800&q=80',
        content: `
            <h3>Переваги електромобілів</h3>
            <p>Чому варто розглянути електрокар:</p>
            <ul>
                <li>Низька вартість експлуатації</li>
                <li>Екологічність</li>
                <li>Тихий рух</li>
                <li>Високий крутний момент</li>
                <li>Мінімальне обслуговування</li>
            </ul>

            <h3>Недоліки та обмеження</h3>
            <p>Що потрібно враховувати:</p>
            <ul>
                <li>Обмежений запас ходу</li>
                <li>Тривалий час зарядки</li>
                <li>Висока вартість нових моделей</li>
                <li>Деградація батареї з часом</li>
                <li>Недостатня інфраструктура зарядних станцій</li>
            </ul>

            <h3>Перевірка батареї</h3>
            <p>Найважливіший елемент електромобіля:</p>
            <ul>
                <li>SoH (State of Health) - стан здоров'я батареї</li>
                <li>Кількість циклів зарядки</li>
                <li>Реальний запас ходу</li>
                <li>Час повної зарядки</li>
            </ul>

            <h3>Типи зарядок</h3>
            <p>Способи зарядки електромобіля:</p>
            <ul>
                <li>Домашня розетка (220V) - найповільніша</li>
                <li>Wallbox (зарядна станція вдома)</li>
                <li>Громадські зарядні станції</li>
                <li>Швидка зарядка (DC)</li>
            </ul>

            <h3>Вартість експлуатації</h3>
            <p>Економія на електромобілі:</p>
            <ul>
                <li>Вартість електроенергії vs бензин</li>
                <li>Відсутність витрат на масло</li>
                <li>Менше деталей для обслуговування</li>
                <li>Рекуперація енергії при гальмуванні</li>
            </ul>

            <h3>Популярні моделі</h3>
            <p>Електромобілі на ринку України:</p>
            <ul>
                <li>Nissan Leaf - найпопулярніший</li>
                <li>Tesla Model 3, S, X</li>
                <li>Renault Zoe</li>
                <li>BMW i3</li>
                <li>Hyundai Kona Electric</li>
            </ul>

            <h3>Підводні камені при покупці</h3>
            <p>На що звернути увагу:</p>
            <ul>
                <li>Історія експлуатації батареї</li>
                <li>Наявність швидкої зарядки</li>
                <li>Вартість заміни батареї</li>
                <li>Гарантія на батарею</li>
                <li>Наявність сервісу в Україні</li>
            </ul>

            <p><strong>Порада:</strong> перед покупкою обов'язково протестуйте реальний запас ходу в різних умовах!</p>
        `
    }
])

const filteredArticles = computed(() => {
    if (activeCategory.value === 'all') {
        return articles.value
    }
    return articles.value.filter(article => article.category === activeCategory.value)
})

const filterByCategory = (category) => {
    activeCategory.value = category
}

const getCategoryName = (category) => {
    const categories = {
        buying: 'Купівля',
        selling: 'Продаж',
        tips: 'Поради'
    }
    return categories[category] || category
}

const openArticle = (id) => {
    const article = articles.value.find(a => a.id === id)
    if (article) {
        selectedArticle.value = article
        document.body.style.overflow = 'hidden'
    }
}

const closeArticle = () => {
    selectedArticle.value = null
    document.body.style.overflow = ''
}

const handleKeydown = (e) => {
    if (e.key === 'Escape' && selectedArticle.value) {
        closeArticle()
    }
}

onMounted(() => {
    window.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
    window.removeEventListener('keydown', handleKeydown)
    document.body.style.overflow = ''
})
</script>

<style scoped lang="scss">
.blog-page {
    min-height: 100vh;
    padding: 2rem 1rem;
}

.container {
    max-width: 1440px;
    margin: 0 auto;
}

.blog-hero {
    text-align: center;
    margin-bottom: 3rem;
    padding: 2rem 0;

    h1 {
        font-size: 3rem;
        margin-bottom: 1rem;
        color: #FFF;
    }

    p {
        font-size: 1.25rem;
        color: rgba(255, 255, 255, 0.8);
    }
}

.categories {
    display: flex;
    gap: 1rem;
    margin-bottom: 3rem;
    flex-wrap: wrap;
    justify-content: center;
}

.category-btn {
    padding: 0.75rem 1.5rem;
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    border: 2px solid rgba(255, 255, 255, 0.2);
    border-radius: 50px;
    color: #FFF;
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;

    &:hover {
        background: rgba(255, 255, 255, 0.2);
        transform: translateY(-2px);
    }

    &.active {
        background: #007bff;
        border-color: #007bff;
    }
}

.articles-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
    gap: 2rem;
    margin-bottom: 3rem;
}

.article-card {
    background: white;
    border-radius: 15px;
    overflow: hidden;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
    transition: all 0.3s ease;
    cursor: pointer;

    &:hover {
        transform: translateY(-5px);
        box-shadow: 0 8px 30px rgba(0, 0, 0, 0.15);
    }
}

.article-image {
    position: relative;
    height: 200px;
    overflow: hidden;

    img {
        width: 100%;
        height: 100%;
        object-fit: cover;
        transition: transform 0.3s ease;
    }

    &:hover img {
        transform: scale(1.05);
    }
}

.article-category {
    position: absolute;
    top: 1rem;
    right: 1rem;
    background: rgba(0, 123, 255, 0.9);
    backdrop-filter: blur(10px);
    color: white;
    padding: 0.5rem 1rem;
    border-radius: 20px;
    font-size: 0.875rem;
    font-weight: 600;
}

.article-content {
    padding: 1.5rem;

    h3 {
        margin: 0 0 1rem 0;
        font-size: 1.5rem;
        color: #333;
    }

    p {
        color: #666;
        line-height: 1.6;
        margin-bottom: 1rem;
    }
}

.article-footer {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-top: 1rem;
    padding-top: 1rem;
    border-top: 1px solid #eee;
}

.read-time {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    color: #999;
    font-size: 0.875rem;

    svg {
        color: #007bff;
    }
}

.read-more {
    color: #007bff;
    font-weight: 600;
    font-size: 0.875rem;
}

.empty-state {
    text-align: center;
    padding: 3rem;
    color: rgba(255, 255, 255, 0.7);
    font-size: 1.25rem;
}

/* Modal Styles */
.modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.9);
    z-index: 9999;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2rem;
    overflow-y: auto;
}

.modal-content {
    background: white;
    border-radius: 20px;
    max-width: 900px;
    width: 100%;
    max-height: 90vh;
    overflow-y: auto;
    position: relative;
}

.modal-close {
    position: absolute;
    top: 1rem;
    right: 1rem;
    float: right;
    background: rgba(0, 0, 0, 0.7);
    backdrop-filter: blur(10px);
    color: white;
    border: none;
    width: 40px;
    height: 40px;
    border-radius: 50%;
    font-size: 1.5rem;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s ease;
    z-index: 10;

    &:hover {
        background: rgba(0, 0, 0, 0.9);
        transform: rotate(90deg);
    }
}

.modal-header {
    position: relative;
    height: 400px;
    overflow: hidden;

    img {
        width: 100%;
        height: 100%;
        object-fit: cover;
    }
}

.modal-category {
    position: absolute;
    bottom: 2rem;
    left: 2rem;
    background: rgba(0, 123, 255, 0.9);
    backdrop-filter: blur(10px);
    color: white;
    padding: 0.75rem 1.5rem;
    border-radius: 30px;
    font-weight: 600;
}

.modal-body {
    padding: 2rem;

    h2 {
        margin: 0 0 1rem 0;
        font-size: 2rem;
        color: #333;
    }
}

.modal-meta {
    display: flex;
    gap: 1rem;
    margin-bottom: 2rem;
    padding-bottom: 1rem;
    border-bottom: 2px solid #eee;
}

.article-text {
    color: #333;
    line-height: 1.8;

    & * {
        color: #333;
        font-family: "SF Pro Text";
    }

    & ul {
        margin-left: 20px;
    }

    h3 {
        margin: 2rem 0 1rem 0;
        color: #007bff;
        font-size: 1.5rem;
    }

    p {
        margin-bottom: 1rem;
    }

    ul {
        margin: 1rem 0 1rem 2rem;

        li {
            margin-bottom: 0.5rem;
        }
    }

    strong {
        color: #007bff;
    }
}

/* Modal Animation */
.modal-enter-active,
.modal-leave-active {
    transition: all 0.3s ease;
}

.modal-enter-from,
.modal-leave-to {
    opacity: 0;
}

.modal-enter-from .modal-content,
.modal-leave-to .modal-content {
    transform: scale(0.9);
}

@media (max-width: 768px) {
    .blog-hero h1 {
        font-size: 2rem;
    }

    .blog-hero p {
        font-size: 1rem;
    }

    .articles-grid {
        grid-template-columns: 1fr;
    }

    .modal-overlay {
        padding: 0;
    }

    .modal-content {
        max-height: 100vh;
        border-radius: 0;
    }

    .modal-header {
        height: 250px;
    }

    .modal-body {
        padding: 1.5rem;

        h2 {
            font-size: 1.5rem;
        }
    }

    .article-text h3 {
        font-size: 1.25rem;
    }
}
</style>