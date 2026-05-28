<script setup>
  import { ref, computed } from 'vue'
  import { useRouter } from 'vue-router'
  import BackButton from '@/components/BackButton.vue'
  import data from '@/data/grammar.json'
  import TopicCard from '@/components/LvlTopicCard.vue'
  import Sidebar from '@/components/LvlSidebar.vue'

  const router = useRouter()
  const topics = data.topics

  const levels = ['A1', 'A2', 'B1', 'B2', 'C1', 'C2']
  const selectedLevel = ref(levels[0])

  const filteredTopics = computed(() => {
    if (!selectedLevel.value) return topics
    return topics.filter((t) => t.level === selectedLevel.value)
  })

  const goToTopic = (topicId) => {
    router.push(`/grammar/topics/${topicId}`)
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
        <h1>Темы грамматики</h1>
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
  .empty-message {
    text-align: center;
    color: #777;
    font-style: italic;
    margin-top: 3rem;
  }
  @media (max-width: 768px) {
    .layout {
      flex-direction: column;
      align-items: stretch;
    }
  }
</style>
