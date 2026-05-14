<script setup>
  import { ref, computed } from 'vue'
  import { useRouter } from 'vue-router'
  import BackButton from '@/components/BackButton.vue'
  import data from '@/data/audition.json'
  import TopicCard from '@/components/LvlTopicCard.vue'
  import Sidebar from '@/components/LvlSidebar.vue'

  const router = useRouter()
  const topics = data.topics

  const levels = ['A1', 'A2', 'B1', 'B2', 'C1', 'C2']

  // По умолчанию можно выбрать первый уровень или оставить null (показать все)
  const selectedLevel = ref(levels[0] || null)

  // Фильтрованные темы
  const filteredTopics = computed(() => {
    if (!selectedLevel.value) return topics
    return topics.filter((t) => t.level === selectedLevel.value)
  })

  const goToTopic = (topicId) => {
    router.push(`/audition/topics/${topicId}`)
  }
</script>

<template>
  <div class="root">
    <div class="top-bar">
      <BackButton to="/dashboard" label="← Дашборд" />
    </div>

    <div class="layout">
      <Sidebar v-model="selectedLevel" :levels="levels" />

      <main class="topics-wrapper">
        <h1>Темы аудирования</h1>
        <div v-if="filteredTopics.length > 0" class="topics-grid">
          <TopicCard
            v-for="topic in filteredTopics"
            :key="topic.id"
            :topic="topic"
            @click="goToTopic" />
        </div>
        <p v-else class="empty-message">Нет тем для этого уровня</p>
      </main>
    </div>
  </div>
</template>

<style scoped>
  .root {
    min-height: 100vh;
    background: linear-gradient(rgba(255, 253, 247, 1) 0%, rgba(255, 253, 247, 1) 100%);
    padding: 2rem 1rem;
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .top-bar {
    width: 100%;
    max-width: 1100px;
    margin-bottom: 2rem;
  }

  .layout {
    display: flex;
    gap: 2rem;
    width: 100%;
    max-width: 1100px;
    align-items: flex-start;
  }

  /* Боковая панель уровней */
  .levels-sidebar {
    width: 160px;
    background: rgba(255, 255, 255, 0.9);
    backdrop-filter: blur(6px);
    border: 1px solid rgba(88, 204, 2, 0.4);
    border-radius: 30px;
    padding: 1.5rem 1rem;
    box-shadow: 0 10px 20px rgba(88, 204, 2, 0.06);
    flex-shrink: 0;
  }

  .levels-title {
    font-size: 1.2rem;
    color: #2c3e50;
    margin-bottom: 1.2rem;
    text-align: center;
    font-weight: 700;
  }

  .level-btn {
    display: block;
    width: 100%;
    padding: 0.8rem 0.5rem;
    margin-bottom: 0.5rem;
    border: none;
    background: transparent;
    border-radius: 20px;
    font-size: 1rem;
    font-weight: 600;
    color: #2c3e50;
    cursor: pointer;
    transition: all 0.2s;
    text-align: center;
  }

  .level-btn:hover {
    background: rgba(88, 204, 2, 0.08);
  }

  .level-btn.active {
    background: rgba(88, 204, 2, 0.15);
    color: #58cc02;
    font-weight: 700;
  }

  /* Темы */
  .topics-wrapper {
    flex: 1;
  }

  h1 {
    text-align: center;
    color: #2c3e50;
    margin-bottom: 2rem;
    margin-top: 0;
  }

  .topics-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    gap: 1.5rem;
  }

  .topic-card {
    background: white;
    border: 1px solid rgba(88, 204, 2, 0.4);
    border-radius: 30px;
    padding: 2rem 1rem;
    text-align: center;
    cursor: pointer;
    transition:
      transform 0.2s,
      box-shadow 0.2s;
    box-shadow: 0 10px 20px rgba(88, 204, 2, 0.04);
  }

  .topic-card:hover {
    transform: translateY(-4px);
    border-color: rgba(88, 204, 2, 0.7);
    box-shadow: 0 12px 24px rgba(88, 204, 2, 0.1);
  }

  .topic-icon {
    font-size: 3rem;
    display: block;
    margin-bottom: 0.5rem;
  }

  .topic-title {
    font-size: 1.2rem;
    font-weight: 600;
    color: #2c3e50;
  }

  .empty-message {
    text-align: center;
    color: #777;
    font-style: italic;
    margin-top: 3rem;
  }

  /* Адаптация для маленьких экранов */
  @media (max-width: 768px) {
    .layout {
      flex-direction: column;
      align-items: stretch;
    }
    .levels-sidebar {
      width: 100%;
      display: flex;
      flex-wrap: wrap;
      gap: 0.5rem;
      padding: 0.8rem;
      border-radius: 20px;
    }
    .levels-title {
      width: 100%;
      margin-bottom: 0.5rem;
    }
    .level-btn {
      flex: 1 0 auto;
      padding: 0.6rem 0.8rem;
    }
  }
</style>
