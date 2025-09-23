<template>
  <div class="app">
    <header>
      <h1>🚀 SVV Devel Project</h1>
      <p>Version: {{ version }}</p>
    </header>
    
    <main>
      <section class="status">
        <h2>Server Status</h2>
        <button @click="checkStatus" :disabled="loading">
          {{ loading ? 'Checking...' : 'Check API Status' }}
        </button>
        <div v-if="status" class="status-result" :class="status.success ? 'success' : 'error'">
          <pre>{{ JSON.stringify(status.data, null, 2) }}</pre>
        </div>
      </section>
    </main>
  </div>
</template>

<script>
import { ref } from 'vue'

export default {
  name: 'App',
  setup() {
    const version = ref('1.0.0')
    const status = ref(null)
    const loading = ref(false)
    
    const checkStatus = async () => {
      loading.value = true
      try {
        const response = await fetch('/api/status')
        const data = await response.json()
        status.value = { success: true, data }
      } catch (error) {
        status.value = { 
          success: false, 
          data: { error: error.message } 
        }
      } finally {
        loading.value = false
      }
    }
    
    return {
      version,
      status,
      loading,
      checkStatus
    }
  }
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.app {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  max-width: 800px;
  margin: 0 auto;
  padding: 40px 20px;
}

header {
  text-align: center;
  margin-bottom: 40px;
}

header h1 {
  color: #2c3e50;
  margin-bottom: 10px;
}

header p {
  color: #7f8c8d;
}

.status {
  background: white;
  padding: 30px;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
}

button {
  background: #3498db;
  color: white;
  padding: 12px 24px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 16px;
  transition: background 0.3s;
}

button:hover:not(:disabled) {
  background: #2980b9;
}

button:disabled {
  background: #bdc3c7;
  cursor: not-allowed;
}

.status-result {
  margin-top: 20px;
  padding: 15px;
  border-radius: 6px;
  font-family: 'Courier New', monospace;
}

.status-result.success {
  background: #d4edda;
  border: 1px solid #c3e6cb;
  color: #155724;
}

.status-result.error {
  background: #f8d7da;
  border: 1px solid #f5c6cb;
  color: #721c24;
}

body {
  background: #ecf0f1;
}
</style>
