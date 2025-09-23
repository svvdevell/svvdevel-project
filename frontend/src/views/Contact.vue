<template>
  <div class="contact-page">
    <div class="container">
      <div class="page-header">
        <h1>Контакты</h1>
        <p>Свяжитесь с нами удобным способом</p>
      </div>

      <div class="contact-content">
        <div class="contact-info">
          <div class="info-card">
            <div class="icon">📞</div>
            <h3>Телефон</h3>
            <p>+7 (999) 123-45-67</p>
            <p>+7 (999) 123-45-68</p>
          </div>

          <div class="info-card">
            <div class="icon">✉️</div>
            <h3>Email</h3>
            <p>info@company.com</p>
            <p>support@company.com</p>
          </div>

          <div class="info-card">
            <div class="icon">📍</div>
            <h3>Адрес</h3>
            <p>г. Москва</p>
            <p>ул. Примерная, д. 123</p>
          </div>

          <div class="info-card">
            <div class="icon">🕒</div>
            <h3>Режим работы</h3>
            <p>Пн-Пт: 9:00 - 18:00</p>
            <p>Сб-Вс: 10:00 - 16:00</p>
          </div>
        </div>

        <div class="contact-form">
          <h3>Оставьте сообщение</h3>
          <form @submit.prevent="submitForm">
            <div class="form-group">
              <input
                v-model="form.name"
                type="text"
                placeholder="Ваше имя"
                required
              >
            </div>
            <div class="form-group">
              <input
                v-model="form.email"
                type="email"
                placeholder="Email"
                required
              >
            </div>
            <div class="form-group">
              <input
                v-model="form.phone"
                type="tel"
                placeholder="Телефон (необязательно)"
              >
            </div>
            <div class="form-group">
              <select v-model="form.subject" required>
                <option value="">Выберите тему</option>
                <option value="general">Общие вопросы</option>
                <option value="support">Техподдержка</option>
                <option value="sales">Продажи</option>
                <option value="partnership">Партнерство</option>
              </select>
            </div>
            <div class="form-group">
              <textarea
                v-model="form.message"
                placeholder="Ваше сообщение"
                rows="5"
                required
              ></textarea>
            </div>
            <button type="submit" class="submit-btn" :disabled="isSubmitting">
              {{ isSubmitting ? 'Отправка...' : 'Отправить сообщение' }}
            </button>
          </form>

          <div v-if="submitMessage" class="submit-message" :class="submitStatus">
            {{ submitMessage }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Contact',
  data() {
    return {
      form: {
        name: '',
        email: '',
        phone: '',
        subject: '',
        message: ''
      },
      isSubmitting: false,
      submitMessage: '',
      submitStatus: ''
    }
  },
  methods: {
    async submitForm() {
      this.isSubmitting = true
      this.submitMessage = ''
      
      try {
        // Имитация отправки формы
        await new Promise(resolve => setTimeout(resolve, 2000))
        
        this.submitMessage = 'Сообщение успешно отправлено! Мы свяжемся с вами в ближайшее время.'
        this.submitStatus = 'success'
        
        // Очистка формы
        this.form = {
          name: '',
          email: '',
          phone: '',
          subject: '',
          message: ''
        }
        
        // Скрытие сообщения через 5 секунд
        setTimeout(() => {
          this.submitMessage = ''
        }, 5000)
        
      } catch (error) {
        this.submitMessage = 'Произошла ошибка при отправке сообщения. Попробуйте еще раз.'
        this.submitStatus = 'error'
      } finally {
        this.isSubmitting = false
      }
    }
  }
}
</script>

<style scoped>
.contact-page {
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

.contact-content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 4rem;
  align-items: start;
}

.contact-info {
  display: grid;
  gap: 1.5rem;
}

.info-card {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  text-align: center;
  border: 1px solid #e1e8ed;
}

.info-card .icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.info-card h3 {
  color: #2c3e50;
  margin-bottom: 1rem;
  font-size: 1.3rem;
}

.info-card p {
  color: #7f8c8d;
  margin: 0.25rem 0;
  line-height: 1.5;
}

.contact-form {
  background: white;
  padding: 2.5rem;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  border: 1px solid #e1e8ed;
}

.contact-form h3 {
  color: #2c3e50;
  margin-bottom: 2rem;
  font-size: 1.5rem;
  text-align: center;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group input,
.form-group select,
.form-group textarea {
  width: 100%;
  padding: 0.75rem 1rem;
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

.form-group textarea {
  resize: vertical;
  min-height: 120px;
}

.submit-btn {
  width: 100%;
  padding: 1rem;
  background: linear-gradient(135deg, #3498db, #2980b9);
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s;
}

.submit-btn:hover:not(:disabled) {
  background: linear-gradient(135deg, #2980b9, #21618c);
  transform: translateY(-2px);
}

.submit-btn:disabled {
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
  border: 1px solid #c3e6cb;
}

.submit-message.error {
  background: #f8d7da;
  color: #721c24;
  border: 1px solid #f5c6cb;
}

@media (max-width: 768px) {
  .contact-content {
    grid-template-columns: 1fr;
    gap: 2rem;
  }
  
  .contact-form {
    padding: 1.5rem;
  }
}
</style>