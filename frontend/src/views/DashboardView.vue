<script setup>
  import { ref } from 'vue'
  import { useRouter } from 'vue-router'
  import TaskCard from '@/components/TaskCard.vue'

  const router = useRouter()
  const sidebarOpen = ref(false)

  const toggleSidebar = () => {
    sidebarOpen.value = !sidebarOpen.value
  }

  const menuItems = [
    { id: 'logout', label: 'Выйти', action: () => console.log('Выйти') },
    { id: 'language', label: 'Выбор языка', action: () => router.push('/language') },
    { id: 'dictionary', label: 'Словарь', action: () => router.push('/dictionary') },
    { id: 'texts', label: 'Мои тексты', action: () => console.log('Мои тексты') },
    { id: 'progress', label: 'Прогресс', action: () => console.log('Прогресс') },
    { id: 'settings', label: 'Настройки', action: () => console.log('Настройки') },
  ]
</script>

<template>
  <div class="mosaic-wrap">
    <div id="i8z5xo3zc_0" class="root root--u-i8z5xo3zc root--s1-i8z5xo3zc root--s2-i8z5xo3zc">
      <div
        id="ikc0yythz_0"
        class="container container--u-ikc0yythz container--s1-ikc0yythz container--s2-ikc0yythz">
        <!-- Основной контент: боковое меню (фиксированное) + центральная область -->
        <div class="dashboard-layout">
          <!-- Боковое меню (fixed) -->
          <aside class="sidebar" :class="{ 'sidebar--open': sidebarOpen }">
            <button class="hamburger" @click="toggleSidebar">
              <span v-if="sidebarOpen">✕</span>
              <span v-else>☰</span>
            </button>
            <nav v-show="sidebarOpen" class="sidebar-nav">
              <button
                v-for="item in menuItems"
                :key="item.id"
                class="menu-item"
                @click="item.action">
                {{ item.label }}
              </button>
            </nav>
          </aside>

          <!-- Центральная область: три карточки -->
          <main class="main-content">
            <div class="cards-row">
              <TaskCard
                title="Слова"
                icon="📚"
                to="/vocabulary/topics"
                description="Выучите 10 новых слов и повторите старые."
                completed-description="Слова на сегодня выучены! Отлично!" />
              <TaskCard
                title="Аудирование"
                icon="🎧"
                to="/audition/topics"
                description="Прослушайте аудиоурок и ответьте на вопросы."
                completed-description="Аудирование пройдено! Хорошая работа!" />
              <TaskCard
                title="Правила"
                icon="📖"
                to="/grammar/topics"
                description="Изучите новое грамматическое правило и выполните упражнения."
                completed-description="Правила изучены! Так держать!" />
            </div>
          </main>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
  /* === Базовые стили === */
  .root--u-i8z5xo3zc {
    background-size: auto;
    background-image: linear-gradient(rgba(255, 253, 247, 1) 0%, rgba(255, 253, 247, 1) 100%);
    background-repeat: no-repeat;
    background-position: left 0px top 0px;
    background-attachment: scroll;
    display: flex;
    align-items: center;
    justify-content: center;
    min-height: 100vh;
  }

  .container--u-ikc0yythz {
    min-height: 100vh;
    width: 100%;
    max-width: 1200px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    padding: 5rem 1rem;
  }

  .container--u-iw3pdlkas {
    display: flex;
    justify-content: flex-start;
    align-items: center;
    width: 100%;
    max-width: 1000px;
    padding: 0 1rem;
    margin-bottom: 2rem;
  }

  .dashboard-layout {
    display: flex;
    width: 100%;
    max-width: 1000px;
    gap: 2rem;
    align-items: flex-start;
  }

  /* === Фиксированное боковое меню === */
  .sidebar {
    position: fixed;
    left: 0;
    top: 0;
    bottom: 0;
    width: 40px; /* свёрнутая ширина */
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(5px);
    border-right: 1px solid rgba(88, 204, 2, 0.4);
    box-shadow: 0 10px 20px rgba(88, 204, 2, 0.04);
    padding: 1rem 0.5rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    transition: width 0.3s ease;
    z-index: 100;
    overflow-x: hidden;
    overflow-y: auto;
  }

  .sidebar--open {
    width: 250px; /* развёрнутая ширина */
    align-items: stretch;
    padding: 0.8rem 1rem;
  }

  .hamburger {
    background: none;
    border: none;
    font-size: 1.4rem;
    cursor: pointer;
    color: #2c3e50;
    padding: 0.4rem;
    border-radius: 8px;
    transition:
      background 0.2s,
      color 0.2s;
    align-self: flex-end;
    /* всегда выравниваем вправо, но в развёрнутом меню кнопка остаётся на том же месте */
  }
  .sidebar--open .hamburger {
    margin-bottom: 0.4rem;
  }
  .hamburger:hover {
    background: rgba(88, 204, 2, 0.08);
    color: #58cc02;
  }

  .sidebar-nav {
    display: flex;
    flex-direction: column;
    gap: 0.8rem;
    width: 100%;
  }

  .menu-item {
    display: block;
    width: 100%;
    padding: 0.8rem 1rem;
    border: none;
    background: transparent;
    text-align: left;
    font-size: 1rem;
    color: #2c3e50;
    border-radius: 20px;
    cursor: pointer;
    transition:
      background 0.2s,
      color 0.2s;
    font-weight: 500;
    white-space: nowrap;
  }

  .menu-item:hover {
    background: rgba(88, 204, 2, 0.08);
    color: #58cc02;
  }

  /* Центральная область */
  .main-content {
    flex: 1;
    display: flex;
    justify-content: center;
    align-items: flex-start;
  }

  .cards-row {
    display: flex;
    flex-direction: row;
    gap: 1.5rem;
    width: 100%;
    max-width: 900px;
    justify-content: center;
    flex-wrap: wrap;
  }

  .cards-row > * {
    flex: 1 1 250px; /* будут растягиваться, но не меньше 250px */
    max-width: 1fr;
  }
</style>
