<!-- src/views/ErrView.vue -->
<script setup>
  import { useRouter } from 'vue-router'

  const router = useRouter()

  const goHome = () => {
    router.push('/dashboard')
  }
</script>

<template>
  <div class="err-page">
    <div class="err-content">
      <h1 class="err-code">404</h1>
      <p class="err-message">Страница не найдена</p>
      <p class="err-description">
        Возможно, вы перешли по неверной ссылке или страница была удалена.
      </p>
      <button class="btn-home" @click="goHome">На дашборд</button>
    </div>
  </div>
</template>

<style scoped>
  .err-page {
    min-height: 100vh;
    background: linear-gradient(135deg, #fffdf7 0%, #f8f9fa 100%);
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2rem;
  }

  .err-content {
    text-align: center;
    max-width: 480px;
  }

  .err-code {
    font-size: 8rem;
    font-weight: 800;
    color: #58cc02;
    margin: 0;
    line-height: 1;
    text-shadow: 0 4px 10px rgba(88, 204, 2, 0.2);
  }

  .err-message {
    font-size: 2rem;
    font-weight: 700;
    color: #2c3e50;
    margin: 1rem 0 0.5rem;
  }

  .err-description {
    font-size: 1rem;
    color: #666;
    margin-bottom: 2rem;
    line-height: 1.5;
  }

  .btn-home {
    display: inline-block;
    padding: 0.8rem 2rem;
    background-color: #58cc02;
    color: white;
    border: none;
    border-radius: 30px;
    font-size: 1.1rem;
    font-weight: 600;
    cursor: pointer;
    transition:
      background 0.2s,
      transform 0.2s;
    box-shadow: 0 4px 12px rgba(88, 204, 2, 0.3);
  }

  .btn-home:hover {
    background-color: #4da801;
    transform: translateY(-1px);
  }
</style>
