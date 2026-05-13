<script setup>
  import { useRouter } from 'vue-router'
  import BackButton from '@/components/BackButton.vue'
  import data from '@/data/lessons.json'

  const router = useRouter()
  const topics = data.topics

  const goToTopic = (topicId) => {
    router.push(`/audition/topics/${topicId}`)
  }
</script>

<template>
  <div class="root">
    <div class="top-bar">
      <BackButton to="/dashboard" label="← Дашборд" />
    </div>
    <div class="topics-wrapper">
      <h1>Темы аудирования</h1>
      <div class="topics-grid">
        <div
          v-for="topic in topics"
          :key="topic.id"
          class="topic-card"
          @click="goToTopic(topic.id)">
          <span class="topic-icon">{{ topic.icon }}</span>
          <span class="topic-title">{{ topic.title }}</span>
        </div>
      </div>
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
    max-width: 1000px;
    margin-bottom: 2rem;
  }
  .topics-wrapper {
    width: 100%;
    max-width: 1000px;
  }
  h1 {
    text-align: center;
    color: #2c3e50;
    margin-bottom: 2rem;
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
</style>
