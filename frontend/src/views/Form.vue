<template>
  <div class="form-page">
    <div class="container">
      <div class="page-header">
        <h1>Анкета</h1>
        <p>Заполните форму для участия в программе</p>
      </div>

      <div class="form-container">
        <form @submit.prevent="submitForm" class="survey-form">
          <!-- Личная информация -->
          <div class="form-section">
            <h3>📋 Личная информация</h3>
            <div class="form-row">
              <div class="form-group">
                <label>Имя *</label>
                <input v-model="form.firstName" type="text" required>
              </div>
              <div class="form-group">
                <label>Фамилия *</label>
                <input v-model="form.lastName" type="text" required>
              </div>
            </div>
            <div class="form-row">
              <div class="form-group">
                <label>Email *</label>
                <input v-model="form.email" type="email" required>
              </div>
              <div class="form-group">
                <label>Телефон</label>
                <input v-model="form.phone" type="tel">
              </div>
            </div>
            <div class="form-group">
              <label>Дата рождения</label>
              <input v-model="form.birthDate" type="date">
            </div>
          </div>

          <!-- Интересы -->
          <div class="form-section">
            <h3>🎯 Интересы и предпочтения</h3>
            <div class="form-group">
              <label>Сфера деятельности</label>
              <select v-model="form.occupation">
                <option value="">Выберите сферу</option>
                <option value="it">IT и технологии</option>
                <option value="business">Бизнес и финансы</option>
                <option value="education">Образование</option>
                <option value="healthcare">Здравоохранение</option>
                <option value="creative">Творчество и дизайн</option>
                <option value="other">Другое</option>
              </select>
            </div>
            
            <div class="form-group">
              <label>Интересующие категории товаров</label>
              <div class="checkbox-group">
                <label class="checkbox-item">
                  <input type="checkbox" v-model="form.interests" value="tech">
                  <span>Технологии</span>
                </label>
                <label class="checkbox-item">
                  <input type="checkbox" v-model="form.interests" value="fashion">
                  <span>Мода</span>
                </label>
                <label class="checkbox-item">
                  <input type="checkbox" v-model="form.interests" value="food">
                  <span>Еда и напитки</span>
                </label>
                <label class="checkbox-item">
                  <input type="checkbox" v-model="form.interests" value="home">
                  <span>Дом и быт</span>
                </label>
                <label class="checkbox-item">
                  <input type="checkbox" v-model="form.interests" value="sport">
                  <span>Спорт</span>
                </label>
                <label class="checkbox-item">
                  <input type="checkbox" v-model="form.interests" value="travel">
                  <span>Путешествия</span>
                </label>
              </div>
            </div>

            <div class="form-group">
              <label>Как часто вы делаете покупки онлайн?</label>
              <div class="radio-group">
                <label class="radio-item">
                  <input type="radio" v-model="form.shoppingFrequency" value="daily">
                  <span>Ежедневно</span>
                </label>
                <label class="radio-item">
                  <input type="radio" v-model="form.shoppingFrequency" value="weekly">
                  <span>Еженедельно</span>
                </label>
                <label class="radio-item">
                  <input type="radio" v-model="form.shoppingFrequency" value="monthly">
                  <span>Ежемесячно</span>
                </label>
                <label class="radio-item">
                  <input type="radio" v-model="form.shoppingFrequency" value="rarely">
                  <span>Редко</span>
                </label>
              </div>
            </div>
          </div>

          <!-- Обратная связь -->
          <div class="form-section">
            <h3>💬 Обратная связь</h3>
            <div class="form-group">
              <label>Оценка нашего сайта</label>
              <div class="rating-group">
                <span 
                  v-for="star in 5" 
                  :key="star"
                  class="star"
                  :class="{ active: form.rating >= star }"
                  @click="form.rating = star"
                >
                  ⭐
                </span>
              </div>
            </div>
            
            <div class="form-group">
              <label>Комментарии и пожелания</label>
              <textarea 
                v-model="form.comments" 
                rows="4" 
                placeholder="Поделитесь своими впечатлениями или предложениями..."
              ></textarea>
            </div>

            <div class="form-group">
              <label class="checkbox-item">
                <input type="checkbox" v-model="form.newsletter" value="yes">
                <span>Я хочу получать новости и специальные предложения</span>
              </label>
            </div>

            <div class="form-group">
              <label class="checkbox-item">
                <input type="checkbox" v-model="form.agreement" required>
                <span>Я согласен с обработкой персональных данных *</span>
              </label>
            </div>
          </div>

          <div class="form-actions">
            <button type="button" @click="resetForm" class="btn btn-secondary">
              Очистить форму
            </button>
            <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
              {{ isSubmitting ? 'Отправка...' : 'Отправить анкету' }}
            </button>
          </div>

          <div v-if="submitMessage" class="submit-message" :class="submitStatus">
            {{ submitMessage }}
          </div>
        </form>

        <!-- Превью данных -->
        <div v-if="showPreview" class="form-preview">
          <h3>Превью данных</h3>
          <pre>{{ JSON.stringify(form, null, 2) }}</pre>
          <button @click="showPreview = false" class="btn btn-small">Скрыть</button>
        </div>
        <button @click="showPreview = !showPreview" class="preview-toggle">
          {{ showPreview ? 'Скрыть' : 'Показать' }} данные формы
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Form',
  data() {
    return {
      form: {
        firstName: '',
        lastName: '',
        email: '',
        phone: '',
        birthDate: '',
        occupation: '',
        interests: [],
        shoppingFrequency: '',
        rating: 0,
        comments: '',
        newsletter: false,
        agreement: false
      },
      isSubmitting: false,
      submitMessage: '',
      submitStatus: '',
      showPreview: false
    }
  },
  methods: {
    async submitForm() {
      if (!this.form.agreement) {
        alert('Необходимо согласиться с обработкой персональных данных')
        return
      }

      this.isSubmitting = true
      this.submitMessage = ''
      
      try {
        // Имитация отправки формы
        await new Promise(resolve => setTimeout(resolve, 2000))
        
        this.submitMessage = 'Анкета успешно отправлена! Спасибо за участие.'
        this.submitStatus = 'success'
        
        // Сброс формы через 3 секунды
        setTimeout(() => {
          this.resetForm()
          this.submitMessage = ''
        }, 3000)
        
      } catch (error) {
        this.submitMessage = 'Произошла ошибка при отправке анкеты. Попробуйте еще раз.'
        this.submitStatus = 'error'
      } finally {
        this.isSubmitting = false
      }
    },
    
    resetForm() {
      this.form = {
        firstName: '',
        lastName: '',
        email: '',
        phone: '',
        birthDate: '',
        occupation: '',
        interests: [],
        shoppingFrequency: '',
        rating: 0,
        comments: '',
        newsletter: false,
        agreement: false
      }
    }
  }
}
</script>

<style scoped>
.form-page {
  padding: 2rem 0;
  background: #f8f9fa;
  min-height: 100vh;
}

.container {
  max-width: 800px;
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

.form-container {
  position: relative;
}

.survey-form {
  background: white;
  padding: 2.5rem;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.form-section {
  margin-bottom: 3rem;
  padding-bottom: 2rem;
  border-bottom: 1px solid #e1e8ed;
}

.form-section:last-of-type {
  border-bottom: none;
}

.form-section h3 {
  color: #2c3e50;
  margin-bottom: 1.5rem;
  font-size: 1.3rem;
  display: flex;
  align-items: center;
  gap: 0.5rem;
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

.checkbox-group,
.radio-group {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
  margin-top: 0.5rem;
}

.checkbox-item,
.radio-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
  padding: 0.5rem;
  border-radius: 6px;
  transition: background-color 0.3s;
}

.checkbox-item:hover,
.radio-item:hover {
  background-color: #f8f9fa;
}

.checkbox-item input,
.radio-item input {
  width: auto;
  margin: 0;
}

.rating-group {
  display: flex;
  gap: 0.5rem;
  margin-top: 0.5rem;
}

.star {
  font-size: 2rem;
  cursor: pointer;
  opacity: 0.3;
  transition: opacity 0.3s;
}

.star.active,
.star:hover {
  opacity: 1;
}

.form-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  margin-top: 2rem;
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

.btn-primary:hover:not(:disabled) {
  background: linear-gradient(135deg, #2980b9, #21618c);
  transform: translateY(-2px);
}

.btn-secondary {
  background: #95a5a6;
  color: white;
}

.btn-secondary:hover {
  background: #7f8c8d;
}

.btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
  transform: none;
}

.submit-message {
  margin-top: 1.5rem;
  padding: 1rem;
  border-radius: 8px;
  font-weight: 500;
  text-align: center;
}

.submit-message.success {
  background: #d4edda;
  color: #155724;
}

.submit-message.error {
  background: #f8d7da;
  color: #721c24;
}

.form-preview {
  position: fixed;
  top: 50%;
  right: 2rem;
  transform: translateY(-50%);
  background: white;
  padding: 1rem;
  border-radius: 8px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
  max-width: 300px;
  max-height: 400px;
  overflow-y: auto;
  z-index: 1000;
}

.form-preview h3 {
  margin-bottom: 1rem;
  color: #2c3e50;
}

.form-preview pre {
  font-size: 0.8rem;
  white-space: pre-wrap;
  word-break: break-all;
}

.preview-toggle {
  position: fixed;
  bottom: 2rem;
  right: 2rem;
  background: #3498db;
  color: white;
  border: none;
  padding: 0.8rem 1.5rem;
  border-radius: 25px;
  cursor: pointer;
  box-shadow: 0 4px 15px rgba(52, 152, 219, 0.3);
  z-index: 999;
}

.btn-small {
  padding: 0.5rem 1rem;
  font-size: 0.9rem;
}

@media (max-width: 768px) {
  .form-row {
    grid-template-columns: 1fr;
  }
  
  .survey-form {
    padding: 1.5rem;
  }
  
  .form-actions {
    flex-direction: column;
  }
  
  .checkbox-group,
  .radio-group {
    grid-template-columns: 1fr;
  }
  
  .form-preview {
    position: relative;
    top: auto;
    right: auto;
    transform: none;
    margin-top: 2rem;
    max-width: none;
  }
  
  .preview-toggle {
    position: relative;
    bottom: auto;
    right: auto;
    margin-top: 1rem;
    width: 100%;
  }
}
</style>