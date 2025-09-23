<template>
  <div class="brands-page">
    <div class="container">
      <div class="page-header">
        <h1>Наши бренды</h1>
        <p>Познакомьтесь с брендами, которые мы представляем</p>
      </div>

      <div class="brands-grid">
        <div 
          v-for="brand in brands" 
          :key="brand.id"
          class="brand-card"
          @click="selectBrand(brand)"
          :class="{ active: selectedBrand?.id === brand.id }"
        >
          <div class="brand-logo">
            {{ brand.logo }}
          </div>
          <h3>{{ brand.name }}</h3>
          <p>{{ brand.description }}</p>
          <div class="brand-stats">
            <span class="stat">
              <strong>{{ brand.products }}</strong> товаров
            </span>
            <span class="stat rating">
              ⭐ {{ brand.rating }}
            </span>
          </div>
        </div>
      </div>

      <div v-if="selectedBrand" class="brand-details">
        <h2>{{ selectedBrand.name }}</h2>
        <div class="details-content">
          <div class="detail-section">
            <h4>О бренде</h4>
            <p>{{ selectedBrand.fullDescription }}</p>
          </div>
          <div class="detail-section">
            <h4>Особенности</h4>
            <ul>
              <li v-for="feature in selectedBrand.features" :key="feature">
                {{ feature }}
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Brands',
  data() {
    return {
      selectedBrand: null,
      brands: [
        {
          id: 1,
          name: 'TechPro',
          logo: '💻',
          description: 'Инновационные технологические решения',
          products: 156,
          rating: 4.8,
          fullDescription: 'TechPro - ведущий поставщик технологических решений для бизнеса. Компания специализируется на разработке программного обеспечения и предоставлении IT-услуг.',
          features: [
            'Современные технологии',
            'Высокое качество продукции',
            '24/7 техническая поддержка',
            'Гибкие тарифные планы'
          ]
        },
        {
          id: 2,
          name: 'StyleMart',
          logo: '👗',
          description: 'Модная одежда и аксессуары',
          products: 324,
          rating: 4.6,
          fullDescription: 'StyleMart - это современный бренд модной одежды, который сочетает в себе стиль, качество и доступность.',
          features: [
            'Актуальные тренды моды',
            'Качественные материалы',
            'Широкий размерный ряд',
            'Быстрая доставка'
          ]
        },
        {
          id: 3,
          name: 'FoodDelight',
          logo: '🍕',
          description: 'Качественные продукты питания',
          products: 89,
          rating: 4.9,
          fullDescription: 'FoodDelight предлагает широкий ассортимент качественных продуктов питания от проверенных производителей.',
          features: [
            'Свежие продукты',
            'Органические ингредиенты',
            'Быстрая доставка',
            'Контроль качества'
          ]
        },
        {
          id: 4,
          name: 'HomeComfort',
          logo: '🏠',
          description: 'Товары для дома и уюта',
          products: 267,
          rating: 4.7,
          fullDescription: 'HomeComfort - ваш надежный партнер в создании уютного и комфортного дома.',
          features: [
            'Качественная мебель',
            'Стильный декор',
            'Доступные цены',
            'Сборка и установка'
          ]
        }
      ]
    }
  },
  methods: {
    selectBrand(brand) {
      this.selectedBrand = this.selectedBrand?.id === brand.id ? null : brand
    }
  }
}
</script>

<style scoped>
.brands-page {
  padding: 2rem 0;
}

.container {
  max-width: 1200px;
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
}

.brands-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  margin-bottom: 3rem;
}

.brand-card {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  border: 2px solid transparent;
  cursor: pointer;
  transition: all 0.3s ease;
}

.brand-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

.brand-card.active {
  border-color: #3498db;
  box-shadow: 0 8px 25px rgba(52, 152, 219, 0.2);
}

.brand-logo {
  font-size: 4rem;
  text-align: center;
  margin-bottom: 1rem;
}

.brand-card h3 {
  color: #2c3e50;
  font-size: 1.5rem;
  margin-bottom: 0.5rem;
  text-align: center;
}

.brand-card p {
  color: #7f8c8d;
  text-align: center;
  margin-bottom: 1.5rem;
  line-height: 1.5;
}

.brand-stats {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 1rem;
  border-top: 1px solid #ecf0f1;
}

.stat {
  color: #34495e;
  font-size: 0.9rem;
}

.rating {
  color: #f39c12;
}

.brand-details {
  background: linear-gradient(135deg, #74b9ff 0%, #0984e3 100%);
  color: white;
  padding: 3rem;
  border-radius: 12px;
  animation: fadeIn 0.3s ease-in-out;
}

.brand-details h2 {
  font-size: 2rem;
  margin-bottom: 2rem;
  text-align: center;
}

.details-content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 3rem;
}

.detail-section h4 {
  font-size: 1.3rem;
  margin-bottom: 1rem;
  color: #dff9fb;
}

.detail-section p {
  line-height: 1.6;
  opacity: 0.9;
}

.detail-section ul {
  list-style: none;
  padding: 0;
}

.detail-section li {
  padding: 0.5rem 0;
  opacity: 0.9;
  position: relative;
  padding-left: 1.5rem;
}

.detail-section li::before {
  content: '✓';
  position: absolute;
  left: 0;
  color: #00b894;
  font-weight: bold;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (max-width: 768px) {
  .brands-grid {
    grid-template-columns: 1fr;
  }
  
  .details-content {
    grid-template-columns: 1fr;
    gap: 2rem;
  }
  
  .brand-details {
    padding: 2rem;
  }
}
</style>