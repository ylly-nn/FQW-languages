<script setup>
  import { ref, computed, onMounted } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import BackButton from '@/components/BackButton.vue'
  import LessonList from '@/components/LessonList.vue'
  import data from '@/data/audition.json'

  const route = useRoute()
  const router = useRouter()
  const topicId = computed(() => parseInt(route.params.topicId))
  const topic = computed(() => data.topics.find((t) => t.id === topicId.value))

  const allLessons = data.lessons.filter((l) => l.topicId === topicId.value)

  const completedIds = ref(new Set())
  const premiumIds = ref(new Set())

  onMounted(() => {
    const stored = localStorage.getItem('audition_completed')
    if (stored) {
      completedIds.value = new Set(JSON.parse(stored))
    }
    const storedPremium = localStorage.getItem('audition_premium')
    if (storedPremium) {
      premiumIds.value = new Set(JSON.parse(storedPremium))
    }
  })

  const saveCompleted = () => {
    localStorage.setItem('audition_completed', JSON.stringify([...completedIds.value]))
  }
  const savePremium = () => {
    localStorage.setItem('audition_premium', JSON.stringify([...premiumIds.value]))
  }

  const selectLesson = (lesson) => {
    router.push(`/audition/lesson/${lesson.id}`)
  }

  const togglePremium = (lesson) => {
    if (premiumIds.value.has(lesson.id)) {
      premiumIds.value.delete(lesson.id)
    } else {
      premiumIds.value.add(lesson.id)
    }
    premiumIds.value = new Set(premiumIds.value) // реактивность
    savePremium()
  }

  // При возврате с урока можно пометить выполненным, но будем передавать статус через query или глобально
  // Пока просто добавим пример: в будущем можно слушать событие завершения
</script>

<template>
  <div class="root">
    <div class="top-bar">
      <BackButton to="/audition/topics" label="← Темы" />
    </div>

    <div class="content" v-if="topic">
      <!-- Карточка с рисунком (иконка темы) -->
      <div class="topic-header-card">
        <span class="topic-icon">{{ topic.icon }}</span>
        <h1 class="topic-title">{{ topic.title }}</h1>
      </div>

      <!-- Список уроков (отдельно от карточки) -->
      <div class="lessons-list-container">
        <LessonList
          v-if="allLessons.length > 0"
          :lessons="allLessons"
          :completedIds="completedIds"
          :premiumIds="premiumIds"
          @select="selectLesson"
          @toggle-premium="togglePremium" />
        <p v-else class="empty-message">Нет уроков для этой темы</p>
      </div>
    </div>

    <div v-else class="not-found">Тема не найдена</div>
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
    max-width: 1000px;
    margin-bottom: 2rem;
  }

  .content {
    width: 100%;
    max-width: 800px;
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }

  /* Карточка с рисунком */
  .topic-header-card {
    background: white;
    border: 1px solid rgba(88, 204, 2, 0.4);
    border-radius: 40px;
    padding: 2.5rem 2rem;
    text-align: center;
    box-shadow: 0 20px 40px rgba(88, 204, 2, 0.08);
  }

  .topic-icon {
    font-size: 5rem;
    display: block;
    margin-bottom: 0.5rem;
  }

  .topic-title {
    font-size: 2rem;
    color: #2c3e50;
    margin: 0;
  }

  /* Контейнер списка уроков */
  .lessons-list-container {
    background: transparent; /* список не оборачивается в карточку */
  }

  .not-found {
    text-align: center;
    font-size: 1.5rem;
    color: #777;
    margin-top: 5rem;
  }

  .empty-message {
    text-align: center;
    color: #777;
    font-style: italic;
    margin-top: 3rem;
    font-size: 1.2rem;
  }
</style>
