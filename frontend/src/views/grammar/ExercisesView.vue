<script setup>
  import { computed } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import BackButton from '@/components/BackButton.vue'
  import data from '@/data/grammar.json'

  const route = useRoute()
  const router = useRouter()

  const topicId = computed(() => parseInt(route.params.topicId))
  const topic = computed(() => data.topics.find((t) => t.id === topicId.value))

  // Подсчитываем количество упражнений каждого типа
  const exercises = computed(() => {
    const lessons = data.lessons.filter((l) => l.topicId === topicId.value)
    const all = []
    lessons.forEach((lesson) => {
      lesson.exercises.forEach((ex) => all.push(ex))
    })
    return all
  })

  const countsByType = computed(() => {
    const counts = { fill: 0, choose: 0, reorder: 0 }
    exercises.value.forEach((ex) => {
      if (counts[ex.type] !== undefined) counts[ex.type]++
    })
    return counts
  })

  const goToType = (type) => {
    router.push(`/grammar/exercise/${topicId.value}/${type}`)
  }
</script>

<template>
  <div class="root">
    <div class="top-bar">
      <BackButton :to="`/grammar/topics/${topicId}`" label="← К теории" />
    </div>

    <main v-if="topic" class="content">
      <h1 class="page-title">{{ topic.title }} — Задания</h1>

      <div class="type-buttons">
        <button class="type-btn" :disabled="countsByType.fill === 0" @click="goToType('fill')">
          <span class="type-icon">✍️</span>
          <span class="type-label">Вставить слово</span>
          <span class="type-count">{{ countsByType.fill }} заданий</span>
        </button>

        <button class="type-btn" :disabled="countsByType.choose === 0" @click="goToType('choose')">
          <span class="type-icon">🔘</span>
          <span class="type-label">Выбрать ответ</span>
          <span class="type-count">{{ countsByType.choose }} заданий</span>
        </button>

        <button
          class="type-btn"
          :disabled="countsByType.reorder === 0"
          @click="goToType('reorder')">
          <span class="type-icon">🔄</span>
          <span class="type-label">Составить предложение</span>
          <span class="type-count">{{ countsByType.reorder }} заданий</span>
        </button>
      </div>
    </main>

    <div v-else class="not-found">Тема не найдена</div>
  </div>
</template>

<style scoped>
  .root {
    min-height: 100vh;
    background: #fffdf7;
    padding: 2rem 1rem;
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .top-bar {
    width: 100%;
    max-width: 800px;
    margin-bottom: 2rem;
  }

  .content {
    width: 100%;
    max-width: 800px;
    text-align: center;
  }

  .page-title {
    color: #2c3e50;
    margin-bottom: 2.5rem;
  }

  .type-buttons {
    display: flex;
    flex-direction: column;
    gap: 1.2rem;
    align-items: center;
  }

  .type-btn {
    display: flex;
    align-items: center;
    gap: 1rem;
    width: 100%;
    max-width: 450px;
    padding: 1.2rem 1.8rem;
    background: white;
    border: 2px solid rgba(88, 204, 2, 0.4);
    border-radius: 20px;
    font-size: 1.2rem;
    cursor: pointer;
    transition: all 0.2s;
  }

  .type-btn:hover:not(:disabled) {
    border-color: #58cc02;
    box-shadow: 0 8px 20px rgba(88, 204, 2, 0.12);
    transform: translateY(-2px);
  }

  .type-btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
    background: #f5f5f5;
  }

  .type-icon {
    font-size: 2rem;
  }

  .type-label {
    flex: 1;
    font-weight: 600;
    color: #2c3e50;
    text-align: left;
  }

  .type-count {
    background: rgba(88, 204, 2, 0.1);
    color: #58cc02;
    padding: 0.3rem 0.8rem;
    border-radius: 20px;
    font-size: 0.9rem;
    font-weight: 500;
  }

  .not-found {
    text-align: center;
    font-size: 1.5rem;
    color: #777;
    margin-top: 5rem;
  }
</style>
