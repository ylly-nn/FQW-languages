<script setup>
  import { ref, onMounted, computed } from 'vue'
  import { useRouter } from 'vue-router'
  import TaskCard from '@/components/TaskCard.vue'
  import { authAPI } from '@/api/auth'

  const router = useRouter()
  const sidebarOpen = ref(false)
  const showModal = ref(false)

  // Список языков из переменной окружения
  const nativeLanguages = computed(() => {
    const langString = import.meta.env.VITE_NATIVE_LANGUAGES || 'Русский,English,Deutsch'
    return langString.split(',').map((lang) => lang.trim())
  })

  const selectedLanguage = ref('')

  // Загрузка сохранённого языка
  onMounted(() => {
    const saved = localStorage.getItem('nativeLanguage')
    if (saved && nativeLanguages.value.includes(saved)) {
      selectedLanguage.value = saved
    } else if (nativeLanguages.value.length) {
      selectedLanguage.value = nativeLanguages.value[0] // значение по умолчанию
      localStorage.setItem('nativeLanguage', selectedLanguage.value)
    }
  })

  // Сохранение выбранного языка
  const saveLanguage = (lang) => {
    selectedLanguage.value = lang
    localStorage.setItem('nativeLanguage', lang)
    showModal.value = false
    // Можно также вызвать событие или обновить другие части приложения
    console.log(`Родной язык выбран: ${lang}`)
  }

  // Открытие модального окна
  const openLanguageModal = () => {
    showModal.value = true
  }

  const toggleSidebar = () => {
    sidebarOpen.value = !sidebarOpen.value
  }

  const handleLogout = async () => {
    const refreshToken = localStorage.getItem('refresh_token')
    try {
      if (refreshToken) {
        await authAPI.logout(refreshToken)
      }
    } catch (e) {
      console.error('Ошибка при выходе:', e)
    } finally {
      localStorage.removeItem('access_token')
      localStorage.removeItem('refresh_token')
      localStorage.removeItem('expires_in')
      router.push('/')
    }
  }

  const menuItems = [
    { id: 'logout', label: 'Выйти', action: handleLogout },
    { id: 'language', label: 'Выбор языка', action: () => router.push('/language') },
    { id: 'dictionary', label: 'Словарь', action: () => router.push('/dictionary') },
    { id: 'progress', label: 'Прогресс', action: () => router.push('/progress') },
    { id: 'settings', label: 'Родной язык', action: openLanguageModal },
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
                @click="
                  () => {
                    console.log('Клик по:', item.id)
                    item.action()
                  }
                ">
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

    <!-- Модальное окно выбора родного языка -->
    <Teleport to="body">
      <div v-if="showModal" class="modal-overlay" @click.self="showModal = false">
        <div class="modal-container">
          <div class="modal-header">
            <h3>Выберите родной язык</h3>
            <button class="modal-close" @click="showModal = false">✕</button>
          </div>
          <div class="modal-body">
            <div class="language-list">
              <button
                v-for="lang in nativeLanguages"
                :key="lang"
                class="language-option"
                :class="{ active: selectedLanguage === lang }"
                @click="saveLanguage(lang)">
                {{ lang }}
              </button>
            </div>
          </div>
          <div class="modal-footer">
            <button class="modal-cancel" @click="showModal = false">Отмена</button>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<style scoped>
  /* === Базовые стили (оставлены без изменений, только добавлены стили модалки) === */
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
    width: 40px;
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
    width: 250px;
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
    flex: 1 1 250px;
    max-width: 1fr;
  }

  /* === Стили модального окна (в стиле Duolingo / приложения) === */
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    backdrop-filter: blur(4px);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
  }

  .modal-container {
    background: white;
    width: 90%;
    max-width: 400px;
    border-radius: 28px;
    box-shadow: 0 25px 40px rgba(0, 0, 0, 0.2);
    overflow: hidden;
    animation: modalFadeIn 0.2s ease-out;
  }

  @keyframes modalFadeIn {
    from {
      opacity: 0;
      transform: scale(0.95);
    }
    to {
      opacity: 1;
      transform: scale(1);
    }
  }

  .modal-header {
    padding: 1.25rem 1.5rem;
    background: #fff;
    border-bottom: 1px solid #e5e5e5;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .modal-header h3 {
    margin: 0;
    font-size: 1.25rem;
    font-weight: 700;
    color: #2c3e50;
  }

  .modal-close {
    background: none;
    border: none;
    font-size: 1.5rem;
    cursor: pointer;
    color: #999;
    transition: color 0.2s;
    line-height: 1;
  }

  .modal-close:hover {
    color: #58cc02;
  }

  .modal-body {
    padding: 1.5rem;
  }

  .language-list {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
  }

  .language-option {
    background: #f7f9fa;
    border: 1px solid #e0e0e0;
    border-radius: 16px;
    padding: 0.75rem 1rem;
    font-size: 1rem;
    text-align: left;
    cursor: pointer;
    transition: all 0.2s;
    color: #2c3e50;
    font-weight: 500;
  }

  .language-option:hover {
    background: #e8f5e9;
    border-color: #58cc02;
    transform: translateX(4px);
  }

  .language-option.active {
    background: #58cc02;
    border-color: #58cc02;
    color: white;
  }

  .modal-footer {
    padding: 1rem 1.5rem;
    background: #fafafa;
    display: flex;
    justify-content: flex-end;
    border-top: 1px solid #e5e5e5;
  }

  .modal-cancel {
    background: transparent;
    border: 1px solid #ccc;
    border-radius: 40px;
    padding: 0.5rem 1.25rem;
    font-size: 0.9rem;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.2s;
    color: #2c3e50;
  }

  .modal-cancel:hover {
    background: #f0f0f0;
    border-color: #999;
  }
</style>
