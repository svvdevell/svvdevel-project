<template>
  <div class="reviews-page">
    <div class="container">
      <div class="page-header">
        <h1>Відгуки клієнтів</h1>
        <p>Що говорять наші клієнти про роботу автосалону Elegance Auto</p>

        <div class="stats">
          <div class="stat-item">
            <div class="stat-number">{{ averageRating }}</div>
            <div class="stat-label">Середня оцінка</div>
            <div class="stars">
              <span v-for="star in 5" :key="star" class="star"
                :class="{ active: star <= Math.round(averageRating) }">⭐</span>
            </div>
          </div>
          <div class="stat-item">
            <div class="stat-number">{{ reviews.length }}</div>
            <div class="stat-label">Всього відгуків</div>
          </div>
          <div class="stat-item">
            <div class="stat-number">{{ recommendPercent }}%</div>
            <div class="stat-label">Рекомендують нас</div>
          </div>
        </div>
      </div>

      <div class="reviews-controls">
        <div class="filter-buttons">
          <button @click="currentFilter = 'all'" :class="{ active: currentFilter === 'all' }" class="filter-btn">
            Всі відгуки
          </button>
          <button v-for="rating in [5, 4, 3, 2, 1]" :key="rating" @click="currentFilter = rating"
            :class="{ active: currentFilter === rating }" class="filter-btn">
            {{ rating }} ⭐
          </button>
        </div>

        <div class="sort-select">
          <select v-model="sortBy">
            <option value="newest">Спочатку нові</option>
            <option value="oldest">Спочатку старі</option>
            <option value="rating-high">Високі оцінки</option>
            <option value="rating-low">Низькі оцінки</option>
          </select>
        </div>
      </div>

      <div class="reviews-grid">
        <div v-for="review in filteredAndSortedReviews" :key="review.id" class="review-card"
          :class="'rating-' + review.rating">
          <div class="review-header">
            <div class="user-info">
              <div class="avatar">{{ review.name.charAt(0) }}</div>
              <div class="user-details">
                <h3>{{ review.name }}</h3>
                <p class="date">{{ formatDate(review.date) }}</p>
              </div>
            </div>
            <div class="rating">
              <span v-for="star in 5" :key="star" class="star" :class="{ active: star <= review.rating }">⭐</span>
            </div>
          </div>

          <div class="review-body">
            <h4 v-if="review.title">{{ review.title }}</h4>
            <p>{{ review.comment }}</p>

            <div v-if="review.pros.length > 0" class="pros-cons">
              <div class="pros">
                <strong>👍 Плюси:</strong>
                <ul>
                  <li v-for="pro in review.pros" :key="pro">{{ pro }}</li>
                </ul>
              </div>
            </div>

            <div v-if="review.cons.length > 0" class="pros-cons">
              <div class="cons">
                <strong>👎 Мінуси:</strong>
                <ul>
                  <li v-for="con in review.cons" :key="con">{{ con }}</li>
                </ul>
              </div>
            </div>
          </div>

          <div class="review-footer">
            <div class="helpful">
              <button @click="toggleHelpful(review.id)" class="helpful-btn">
                👍 Корисно ({{ review.helpful }})
              </button>
            </div>
            <div class="category">
              <span class="tag">{{ getCategoryName(review.category) }}</span>
            </div>
          </div>
        </div>
      </div>

      <div class="add-review-section">
        <h3>Залишити відгук</h3>
        <button @click="showReviewForm = !showReviewForm" class="btn btn-primary">
          {{ showReviewForm ? 'Сховати форму' : 'Написати відгук' }}
        </button>

        <div v-if="showReviewForm" class="review-form">
          <form @submit.prevent="submitReview">
            <div class="form-row">
              <div class="form-group">
                <input v-model="newReview.name" type="text" placeholder="Ваше ім'я" required>
              </div>
              <div class="form-group">
                <select v-model="newReview.category" required>
                  <option value="">Оберіть категорію</option>
                  <option value="buy">Покупка авто</option>
                  <option value="sell">Продаж авто</option>
                  <option value="service">Сервіс</option>
                  <option value="parts">Запчастини</option>
                </select>
              </div>
            </div>

            <div class="form-group">
              <input v-model="newReview.title" type="text" placeholder="Заголовок відгуку">
            </div>

            <div class="form-group">
              <label>Ваша оцінка:</label>
              <div class="rating-input">
                <span v-for="star in 5" :key="star" class="star" :class="{ active: newReview.rating >= star }"
                  @click="newReview.rating = star">
                  ⭐
                </span>
              </div>
            </div>

            <div class="form-group">
              <textarea v-model="newReview.comment" rows="4" placeholder="Ваш відгук" required></textarea>
            </div>

            <button type="submit" class="btn btn-primary">Відправити відгук</button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Reviews',
  data() {
    return {
      currentFilter: 'all',
      sortBy: 'newest',
      showReviewForm: false,
      newReview: {
        name: '',
        category: '',
        title: '',
        rating: 5,
        comment: ''
      },
      reviews: [
        {
          id: 1,
          name: 'Олександр Петренко',
          rating: 5,
          title: 'Відмінний автосалон!',
          comment: 'Покупав тут Toyota Camry. Дуже задоволений! Менеджер Сергій допоміг вибрати ідеальний варіант під мій бюджет. Всі документи оформили швидко, без ніяких проблем. Автомобіль в відмінному стані, всё як обіцяли. Рекомендую всім!',
          date: '2025-01-18',
          category: 'buy',
          pros: ['Чесні менеджери', 'Швидке оформлення', 'Автомобілі в хорошому стані', 'Адекватні ціни'],
          cons: [],
          helpful: 24
        },
        {
          id: 2,
          name: 'Марина Коваленко',
          rating: 4,
          title: 'Продала авто швидко',
          comment: 'Продавала свій Volkswagen Golf. Оценили авто справедливо, хоча трошки нижче ніж я очікувала. Але зате все честно - сказали відразу про всі недоліки які знайшли. Гроші отримала в день продажі. Процедура зайняла близько 2 годин.',
          date: '2025-01-15',
          category: 'sell',
          pros: ['Чесна оцінка', 'Швидка виплата', 'Прозорість угоди'],
          cons: ['Ціна трошки нижча очікувань'],
          helpful: 18
        },
        {
          id: 3,
          name: 'Іван Мельник',
          rating: 5,
          title: 'Найкращий досвід покупки авто!',
          comment: 'Вже второй раз покупаю тут машину. Цього разу взяв BMW X5. Хлопці молодці - показали всі тонкощі автомобіля, дали тест-драйв на цілий день! Ціна була найкраща в місті. Через місяць після покупки телефонували, питали чи всё добре з авто.',
          date: '2025-01-12',
          category: 'buy',
          pros: ['Довгий тест-драйв', 'Кращі ціни', 'Підтримка після покупки', 'Досвідчені консультанти'],
          cons: [],
          helpful: 31
        },
        {
          id: 4,
          name: 'Наталія Семенова',
          rating: 3,
          title: 'Нормально, але є нюанси',
          comment: 'Продавала Nissan Juke. В цілому все ок, але довго чекала поки приїхав оценщік. Обіцяли що буде через годину, а приїхав тільки через 3. Потім ще документи оформляли довго. Але гроші дали відразу після підписання договору.',
          date: '2025-01-10',
          category: 'sell',
          pros: ['Отримала гроші відразу', 'Документи оформили правильно'],
          cons: ['Довго чекала оценщіка', 'Повільне оформлення'],
          helpful: 12
        },
        {
          id: 5,
          name: 'Андрій Кравчук',
          rating: 5,
          title: 'Супер сервіс і ставлення!',
          comment: 'Шукав Honda Accord не дорого. Знайшов тут ідеальний варіант! Менеджер Оксана дуже професійно підійшла до справи. Показала 3 різні варіанти, розказала все детально. Взяв авто в кредит - оформили за день. Дякую команді!',
          date: '2025-01-08',
          category: 'buy',
          pros: ['Великий вибір', 'Професійні менеджери', 'Допомога з кредитом', 'Детальна консультація'],
          cons: [],
          helpful: 27
        },
        {
          id: 6,
          name: 'Сергій Бондаренко',
          rating: 2,
          title: 'Не дуже задоволений',
          comment: 'Хотів продати свій Ford Focus. Оценили дуже низько, на 20% менше ніж в інших місцях. Кажуть що авто потребує ремонту, хоча я тільки що з СТО після повної діагностики. Не рекомендую для продажі, може для покупки краще.',
          date: '2025-01-05',
          category: 'sell',
          pros: ['Швидка оценка'],
          cons: ['Занижена ціна', 'Необ\'єктивна оценка стану авто'],
          helpful: 8
        },
        {
          id: 7,
          name: 'Юлія Харченко',
          rating: 4,
          title: 'Гарний вибір авто',
          comment: 'Дочка вперше купувала машину, я її супроводжувала. Консультант Михайло дуже терпляче все пояснював, показав декілька бюджетних варіантів. В результаті взяли Hyundai i30. Авто хорошеє, ціна нормальна. Тільки паркінг біля салону маленький.',
          date: '2025-01-03',
          category: 'buy',
          pros: ['Терплячі консультанти', 'Хороші авто', 'Роботa з новачками'],
          cons: ['Маленька парковка'],
          helpful: 15
        },
        {
          id: 8,
          name: 'Василь Григоренко',
          rating: 5,
          title: 'Продав швидко і вигідно',
          comment: 'Потрібно було терміново продати Skoda Octavia через переїзд. Зателефонував, через 40 хвилин приїхав оценщік. Ціну дав справедливу, навіть трошки більше ніж я думав отримати. Все оформили за 1,5 години. Дуже швидко і професійно!',
          date: '2025-01-01',
          category: 'sell',
          pros: ['Швидкий виїзд оценщіка', 'Справедлива ціна', 'Швидке оформлення', 'Професіоналізм'],
          cons: [],
          helpful: 22
        }
      ]
    }
  },
  computed: {
    averageRating() {
      const sum = this.reviews.reduce((acc, review) => acc + review.rating, 0)
      return (sum / this.reviews.length).toFixed(1)
    },

    recommendPercent() {
      const recommendations = this.reviews.filter(review => review.rating >= 4).length
      return Math.round((recommendations / this.reviews.length) * 100)
    },

    filteredAndSortedReviews() {
      let filtered = this.reviews

      // Фільтрація
      if (this.currentFilter !== 'all') {
        filtered = filtered.filter(review => review.rating === this.currentFilter)
      }

      // Сортування
      filtered = [...filtered].sort((a, b) => {
        switch (this.sortBy) {
          case 'newest':
            return new Date(b.date) - new Date(a.date)
          case 'oldest':
            return new Date(a.date) - new Date(b.date)
          case 'rating-high':
            return b.rating - a.rating
          case 'rating-low':
            return a.rating - b.rating
          default:
            return 0
        }
      })

      return filtered
    }
  },

  methods: {
    formatDate(dateString) {
      const date = new Date(dateString)
      return date.toLocaleDateString('uk-UA', {
        day: 'numeric',
        month: 'long',
        year: 'numeric'
      })
    },

    getCategoryName(category) {
      const categories = {
        buy: 'Покупка авто',
        sell: 'Продаж авто',
        service: 'Сервіс',
        parts: 'Запчастини'
      }
      return categories[category] || category
    },

    toggleHelpful(reviewId) {
      const review = this.reviews.find(r => r.id === reviewId)
      if (review) {
        review.helpful++
      }
    },

    submitReview() {
      const review = {
        id: this.reviews.length + 1,
        ...this.newReview,
        date: new Date().toISOString().split('T')[0],
        pros: [],
        cons: [],
        helpful: 0
      }

      this.reviews.unshift(review)

      // Скидання форми
      this.newReview = {
        name: '',
        category: '',
        title: '',
        rating: 5,
        comment: ''
      }

      this.showReviewForm = false
      alert('Дякуємо за ваш відгук!')
    }
  }
}
</script>

<style scoped>
.reviews-page {
  padding: 2rem 0;
  padding-top: 100px;
}

.container {
  max-width: 1440px;
  margin: 0 auto;
  padding: 0 1rem;
}

.page-header {
  text-align: center;
  margin-bottom: 3rem;
}

.page-header h1 {
  color: #2c3e50;
  font-size: 2.5rem;
  margin-bottom: 1rem;
}

.page-header p {
  color: #7f8c8d;
  font-size: 1.2rem;
  margin-bottom: 2rem;
}

.stats {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 2rem;
  max-width: 600px;
  margin: 0 auto;
}

.stat-item {
  background: white;
  padding: 1.5rem;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  text-align: center;
}

.stat-number {
  font-size: 2.5rem;
  font-weight: bold;
  color: #3498db;
  margin-bottom: 0.5rem;
}

.stat-label {
  color: #7f8c8d;
  font-size: 0.9rem;
  margin-bottom: 0.5rem;
}

.stars .star {
  font-size: 1.2rem;
  opacity: 0.3;
}

.stars .star.active {
  opacity: 1;
}

.reviews-controls {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  background: white;
  padding: 1.5rem;
  border-radius: 12px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.filter-buttons {
  display: flex;
  gap: 0.5rem;
}

.filter-btn {
  padding: 0.5rem 1rem;
  border: 2px solid #e1e8ed;
  background: white;
  border-radius: 20px;
  cursor: pointer;
  transition: all 0.3s;
  font-size: 0.9rem;
}

.filter-btn:hover {
  border-color: #3498db;
}

.filter-btn.active {
  background: #3498db;
  color: white;
  border-color: #3498db;
}

.sort-select select {
  padding: 0.5rem 1rem;
  border: 2px solid #e1e8ed;
  border-radius: 8px;
  background: white;
  cursor: pointer;
}

.reviews-grid {
  display: grid;
  gap: 2rem;
  margin-bottom: 3rem;
}

.review-card {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  border-left: 4px solid #e1e8ed;
  transition: transform 0.3s;
}

.review-card:hover {
  transform: translateY(-2px);
}

.review-card.rating-5 {
  border-left-color: #27ae60;
}

.review-card.rating-4 {
  border-left-color: #f39c12;
}

.review-card.rating-3 {
  border-left-color: #e67e22;
}

.review-card.rating-2 {
  border-left-color: #e74c3c;
}

.review-card.rating-1 {
  border-left-color: #c0392b;
}

.review-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 1.5rem;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.avatar {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  background: linear-gradient(135deg, #3498db, #2980b9);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: bold;
  font-size: 1.2rem;
}

.user-details h3 {
  color: #2c3e50;
  margin: 0 0 0.25rem 0;
  font-size: 1.1rem;
}

.date {
  color: #7f8c8d;
  font-size: 0.9rem;
  margin: 0;
}

.rating {
  display: flex;
  gap: 0.25rem;
}

.rating .star {
  font-size: 1.2rem;
  opacity: 0.3;
}

.rating .star.active {
  opacity: 1;
}

.review-body h4 {
  color: #2c3e50;
  margin-bottom: 1rem;
  font-size: 1.2rem;
}

.review-body p {
  color: #555;
  line-height: 1.6;
  margin-bottom: 1.5rem;
}

.pros-cons {
  margin-bottom: 1rem;
}

.pros-cons strong {
  color: #2c3e50;
  display: block;
  margin-bottom: 0.5rem;
}

.pros-cons ul {
  margin: 0;
  padding-left: 1.5rem;
}

.pros-cons li {
  color: #555;
  line-height: 1.5;
  margin-bottom: 0.25rem;
}

.review-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 1rem;
  border-top: 1px solid #e1e8ed;
}

.helpful-btn {
  background: none;
  border: 1px solid #e1e8ed;
  padding: 0.5rem 1rem;
  border-radius: 20px;
  cursor: pointer;
  transition: all 0.3s;
  font-size: 0.9rem;
}

.helpful-btn:hover {
  background: #f8f9fa;
  border-color: #3498db;
}

.tag {
  background: #e8f4fd;
  color: #3498db;
  padding: 0.25rem 0.75rem;
  border-radius: 15px;
  font-size: 0.8rem;
  font-weight: 500;
}

.add-review-section {
  background: white;
  padding: 2.5rem;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  text-align: center;
}

.add-review-section h3 {
  color: #2c3e50;
  margin-bottom: 1.5rem;
  font-size: 1.5rem;
}

.review-form {
  margin-top: 2rem;
  text-align: left;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: #2c3e50;
  font-weight: 500;
}

.form-group input,
.form-group select,
.form-group textarea {
  width: 100%;
  padding: 0.75rem;
  border: 2px solid #e1e8ed;
  border-radius: 8px;
  font-size: 1rem;
  transition: border-color 0.3s;
}

.form-group input:focus,
.form-group select:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #3498db;
}

.rating-input {
  display: flex;
  gap: 0.5rem;
  margin-top: 0.5rem;
}

.rating-input .star {
  font-size: 2rem;
  cursor: pointer;
  opacity: 0.3;
  transition: opacity 0.3s;
}

.rating-input .star.active,
.rating-input .star:hover {
  opacity: 1;
}

.btn {
  padding: 0.8rem 2rem;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s;
}

.btn-primary {
  background: linear-gradient(135deg, #3498db, #2980b9);
  color: white;
}

.btn-primary:hover {
  background: linear-gradient(135deg, #2980b9, #21618c);
  transform: translateY(-2px);
}

@media (max-width: 768px) {
  .stats {
    grid-template-columns: 1fr;
    gap: 1rem;
  }

  .reviews-controls {
    flex-direction: column;
    gap: 1rem;
    align-items: stretch;
  }

  .filter-buttons {
    flex-wrap: wrap;
    justify-content: center;
  }

  .review-header {
    flex-direction: column;
    gap: 1rem;
  }

  .review-footer {
    flex-direction: column;
    gap: 1rem;
    align-items: stretch;
  }

  .form-row {
    grid-template-columns: 1fr;
  }

  .add-review-section {
    padding: 1.5rem;
  }
}
</style>