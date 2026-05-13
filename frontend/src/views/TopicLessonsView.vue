<script setup>
  import { ref, computed, onMounted } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import BackButton from '@/components/BackButton.vue'
  import LessonList from '@/components/LessonList.vue'
  import data from '@/data/lessons.json'

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
    <div class="lessons-wrapper" v-if="topic">
      <h1>{{ topic.icon }} {{ topic.title }}</h1>
      <LessonList
        :lessons="allLessons"
        :completedIds="completedIds"
        :premiumIds="premiumIds"
        @select="selectLesson"
        @toggle-premium="togglePremium" />
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
  .lessons-wrapper {
    width: 100%;
    max-width: 800px;
    background: white;
    border: 1px solid rgba(88, 204, 2, 0.6);
    border-radius: 50px;
    padding: 3rem 2.5rem;
    box-shadow:
      0px 20px 40px rgba(88, 204, 2, 0.06),
      0px 10px 20px rgba(0, 0, 0, 0.04);
  }
  h1 {
    text-align: center;
    color: #2c3e50;
    margin-bottom: 2rem;
  }
  .not-found {
    text-align: center;
    font-size: 1.5rem;
    color: #777;
    margin-top: 5rem;
  }
</style>
