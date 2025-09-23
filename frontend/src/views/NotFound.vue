<template>
  <div class="not-found-page">
    <div class="container">
      <div class="error-content">
        <div class="error-animation">
          <div class="error-number">4</div>
          <div class="error-circle">
            <div class="circle-inner">0</div>
          </div>
          <div class="error-number">4</div>
        </div>
        
        <h1>Страница не найдена</h1>
        <p>К сожалению, запрашиваемая страница не существует или была перемещена.</p>
        
        <div class="suggestions">
          <h3>Возможные причины:</h3>
          <ul>
            <li>Страница была удалена или перемещена</li>
            <li>Неправильно введен адрес</li>
            <li>Ссылка устарела</li>
            <li>У вас нет прав доступа к этой странице</li>
          </ul>
        </div>

        <div class="actions">
          <button @click="goHome" class="btn btn-primary">
            🏠 На главную
          </button>
          <button @click="goBack" class="btn btn-secondary">
            ← Назад
          </button>
          <router-link to="/contact" class="btn btn-outline">
            📞 Связаться с нами
          </router-link>
        </div>

        <div class="popular-pages">
          <h3>Популярные страницы:</h3>
          <div class="page-links">
            <router-link to="/" class="page-link">
              <div class="link-icon">🏠</div>
              <div class="link-text">
                <strong>Главная</strong>
                <span>Начальная страница сайта</span>
              </div>
            </router-link>
            
            <router-link to="/brands" class="page-link">
              <div class="link-icon">🏷️</div>
              <div class="link-text">
                <strong>Бренды</strong>
                <span>Наши партнерские бренды</span>
              </div>
            </router-link>
            
            <router-link to="/reviews" class="page-link">
              <div class="link-icon">⭐</div>
              <div class="link-text">
                <strong>Отзывы</strong>
                <span>Отзывы наших клиентов</span>
              </div>
            </router-link>
            
            <router-link to="/form" class="page-link">
              <div class="link-icon">📝</div>
              <div class="link-text">
                <strong>Анкета</strong>
                <span>Заполните форму</span>
              </div>
            </router-link>
          </div>
        </div>

        <div class="search-section">
          <h3>Поиск по сайту:</h3>
          <div class="search-form">
            <input 
              v-model="searchQuery" 
              type="text" 
              placeholder="Что вы ищете?"
              @keyup.enter="performSearch"
            >
            <button @click="performSearch" class="search-btn">
              🔍 Найти
            </button>
          </div>
        </div>
      </div>

      <div class="error-illustration">
        <div class="astronaut">
          <div class="helmet"></div>
          <div class="body"></div>
          <div class="arms">
            <div class="arm left"></div>
            <div class="arm right"></div>
          </div>
          <div class="legs">
            <div class="leg left"></div>
            <div class="leg right"></div>
          </div>
        </div>
        
        <div class="planets">
          <div class="planet planet1"></div>
          <div class="planet planet2"></div>
          <div class="planet planet3"></div>
        </div>
        
        <div class="stars">
          <div class="star" v-for="star in 20" :key="star" :style="getStarStyle()"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'NotFound',
  data() {
    return {
      searchQuery: ''
    }
  },
  methods: {
    goHome() {
      this.$router.push('/')
    },
    
    goBack() {
      // Если есть история, идем назад, иначе на главную
      if (window.history.length > 1) {
        this.$router.go(-1)
      } else {
        this.$router.push('/')
      }
    },
    
    performSearch() {
      if (this.searchQuery.trim()) {
        // Здесь можно добавить логику поиска
        // Для примера просто показываем alert
        alert(`Поиск: "${this.searchQuery}". Функция поиска будет добавлена позже.`)
        this.searchQuery = ''
      }
    },
    
    getStarStyle() {
      return {
        left: Math.random() * 100 + '%',
        top: Math.random() * 100 + '%',
        animationDelay: Math.random() * 3 + 's',
        animationDuration: (Math.random() * 2 + 2) + 's'
      }
    }
  },
  
  mounted() {
    // Логирование 404 ошибки
    console.log(`404 Error: Page not found - ${this.$route.fullPath}`)
    
    // Можно отправить статистику 404 ошибок
    // analytics.track('404_error', { path: this.$route.fullPath })
  }
}
</script>

<style scoped>
.not-found-page {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  color: white;
  position: relative;
  overflow: hidden;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem 1rem;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 4rem;
  align-items: center;
  z-index: 2;
  position: relative;
}

.error-animation {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 2rem;
}

.error-number {
  font-size: 8rem;
  font-weight: bold;
  color: white;
  text-shadow: 0 0 20px rgba(255, 255, 255, 0.5);
}

.error-circle {
  margin: 0 1rem;
  width: 8rem;
  height: 8rem;
  border: 4px solid white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  animation: rotate 4s linear infinite;
  box-shadow: 0 0 30px rgba(255, 255, 255, 0.3);
}

.circle-inner {
  font-size: 6rem;
  font-weight: bold;
}

@keyframes rotate {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.error-content h1 {
  font-size: 3rem;
  margin-bottom: 1rem;
  text-align: center;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
}

.error-content > p {
  font-size: 1.2rem;
  text-align: center;
  margin-bottom: 2rem;
  opacity: 0.9;
}

.suggestions {
  background: rgba(255, 255, 255, 0.1);
  padding: 1.5rem;
  border-radius: 12px;
  margin-bottom: 2rem;
  backdrop-filter: blur(10px);
}

.suggestions h3 {
  margin-bottom: 1rem;
  font-size: 1.3rem;
}

.suggestions ul {
  list-style: none;
  padding: 0;
}

.suggestions li {
  padding: 0.5rem 0;
  opacity: 0.9;
  position: relative;
  padding-left: 1.5rem;
}

.suggestions li::before {
  content: '•';
  position: absolute;
  left: 0;
  color: #ffd700;
  font-weight: bold;
}

.actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  margin-bottom: 3rem;
  flex-wrap: wrap;
}

.btn {
  padding: 0.8rem 1.5rem;
  border: none;
  border-radius: 25px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s;
  text-decoration: none;
  display: inline-block;
  font-size: 1rem;
}

.btn-primary {
  background: rgba(255, 255, 255, 0.2);
  color: white;
  backdrop-filter: blur(10px);
}

.btn-primary:hover {
  background: rgba(255, 255, 255, 0.3);
  transform: translateY(-2px);
}

.btn-secondary {
  background: transparent;
  color: white;
  border: 2px solid rgba(255, 255, 255, 0.5);
}

.btn-secondary:hover {
  background: rgba(255, 255, 255, 0.1);
  border-color: white;
}

.btn-outline {
  background: transparent;
  color: white;
  border: 2px solid rgba(255, 255, 255, 0.5);
}

.btn-outline:hover {
  background: rgba(255, 255, 255, 0.1);
  border-color: white;
}

.popular-pages {
  margin-bottom: 2rem;
}

.popular-pages h3 {
  margin-bottom: 1rem;
  text-align: center;
  font-size: 1.3rem;
}

.page-links {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 1rem;
}

.page-link {
  display: flex;
  align-items: center;
  gap: 1rem;
  background: rgba(255, 255, 255, 0.1);
  padding: 1rem;
  border-radius: 8px;
  text-decoration: none;
  color: white;
  transition: all 0.3s;
  backdrop-filter: blur(10px);
}

.page-link:hover {
  background: rgba(255, 255, 255, 0.2);
  transform: translateX(5px);
}

.link-icon {
  font-size: 2rem;
  flex-shrink: 0;
}

.link-text strong {
  display: block;
  font-size: 1.1rem;
  margin-bottom: 0.25rem;
}

.link-text span {
  opacity: 0.8;
  font-size: 0.9rem;
}

.search-section {
  background: rgba(255, 255, 255, 0.1);
  padding: 1.5rem;
  border-radius: 12px;
  backdrop-filter: blur(10px);
}

.search-section h3 {
  margin-bottom: 1rem;
  text-align: center;
}

.search-form {
  display: flex;
  gap: 0.5rem;
}

.search-form input {
  flex: 1;
  padding: 0.75rem;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-radius: 25px;
  background: rgba(255, 255, 255, 0.1);
  color: white;
  backdrop-filter: blur(10px);
}

.search-form input::placeholder {
  color: rgba(255, 255, 255, 0.7);
}

.search-form input:focus {
  outline: none;
  border-color: rgba(255, 255, 255, 0.6);
}

.search-btn {
  padding: 0.75rem 1.5rem;
  background: rgba(255, 255, 255, 0.2);
  border: none;
  border-radius: 25px;
  color: white;
  cursor: pointer;
  transition: all 0.3s;
}

.search-btn:hover {
  background: rgba(255, 255, 255, 0.3);
}

/* Иллюстрация космонавта */
.error-illustration {
  position: relative;
  height: 400px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.astronaut {
  position: relative;
  animation: float 6s ease-in-out infinite;
}

@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-20px); }
}

.helmet {
  width: 80px;
  height: 80px;
  background: rgba(255, 255, 255, 0.9);
  border-radius: 50%;
  margin: 0 auto 10px;
  position: relative;
  box-shadow: 0 0 20px rgba(255, 255, 255, 0.5);
}

.helmet::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 60px;
  height: 60px;
  background: rgba(100, 100, 100, 0.3);
  border-radius: 50%;
}

.body {
  width: 60px;
  height: 80px;
  background: rgba(255, 255, 255, 0.8);
  margin: 0 auto;
  border-radius: 30px;
  position: relative;
}

.arms {
  position: absolute;
  top: 100px;
  left: 50%;
  transform: translateX(-50%);
}

.arm {
  width: 20px;
  height: 50px;
  background: rgba(255, 255, 255, 0.8);
  border-radius: 10px;
  position: absolute;
}

.arm.left {
  left: -40px;
  transform: rotate(-30deg);
}

.arm.right {
  right: -40px;
  transform: rotate(30deg);
}

.legs {
  position: absolute;
  top: 170px;
  left: 50%;
  transform: translateX(-50%);
}

.leg {
  width: 20px;
  height: 60px;
  background: rgba(255, 255, 255, 0.8);
  border-radius: 10px;
  position: absolute;
}

.leg.left {
  left: -15px;
}

.leg.right {
  right: -15px;
}

.planets {
  position: absolute;
  width: 100%;
  height: 100%;
}

.planet {
  position: absolute;
  border-radius: 50%;
  opacity: 0.6;
}

.planet1 {
  width: 40px;
  height: 40px;
  background: #ff6b6b;
  top: 20%;
  left: 10%;
  animation: orbit 20s linear infinite;
}

.planet2 {
  width: 60px;
  height: 60px;
  background: #4ecdc4;
  top: 60%;
  right: 10%;
  animation: orbit 25s linear infinite reverse;
}

.planet3 {
  width: 30px;
  height: 30px;
  background: #ffe66d;
  bottom: 20%;
  left: 20%;
  animation: orbit 15s linear infinite;
}

@keyframes orbit {
  from { transform: rotate(0deg) translateX(100px) rotate(0deg); }
  to { transform: rotate(360deg) translateX(100px) rotate(-360deg); }
}

.stars {
  position: absolute;
  width: 100%;
  height: 100%;
  overflow: hidden;
}

.star {
  position: absolute;
  width: 2px;
  height: 2px;
  background: white;
  border-radius: 50%;
  animation: twinkle 3s ease-in-out infinite;
}

@keyframes twinkle {
  0%, 100% { opacity: 0.3; transform: scale(1); }
  50% { opacity: 1; transform: scale(1.5); }
}

@media (max-width: 768px) {
  .container {
    grid-template-columns: 1fr;
    gap: 2rem;
    text-align: center;
  }
  
  .error-number {
    font-size: 6rem;
  }
  
  .error-circle {
    width: 6rem;
    height: 6rem;
  }
  
  .circle-inner {
    font-size: 4rem;
  }
  
  .error-content h1 {
    font-size: 2.5rem;
  }
  
  .actions {
    flex-direction: column;
    align-items: center;
  }
  
  .page-links {
    grid-template-columns: 1fr;
  }
  
  .error-illustration {
    height: 300px;
  }
  
  .astronaut {
    transform: scale(0.8);
  }
}
</style>