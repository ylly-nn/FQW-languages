<script setup>
  import { computed } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import BackButton from '@/components/BackButton.vue'
  import data from '@/data/grammar.json'

  const route = useRoute()
  const router = useRouter()
  const topicId = computed(() => parseInt(route.params.topicId))
  const topic = computed(() => data.topics.find((t) => t.id === topicId.value))

  const goToExercises = () => {
    router.push(`/grammar/topics/${topicId.value}/exercises`)
  }
</script>

<template>
  <div class="root">
    <div class="top-bar">
      <BackButton to="/grammar/topics" label="← Темы" />
    </div>

    <main v-if="topic" class="content">
      <div class="topic-header-card">
        <span class="topic-icon">{{ topic.icon }}</span>
        <h1 class="topic-title">{{ topic.title }}</h1>
      </div>

      <section class="theory-card">
        <h2>Теория</h2>
        <p class="theory-text">{{ topic.theory }}</p>
        <div v-if="topic.examples?.length" class="examples-block">
          <h3>Примеры:</h3>
          <ul>
            <li v-for="(ex, idx) in topic.examples" :key="idx">{{ ex }}</li>
          </ul>
        </div>
      </section>

      <button class="exercises-btn" @click="goToExercises">Перейти к заданиям →</button>
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
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }
  .topic-header-card {
    background: white;
    border: 1px solid rgba(88, 204, 2, 0.4);
    border-radius: 40px;
    padding: 2rem;
    text-align: center;
    box-shadow: 0 20px 40px rgba(88, 204, 2, 0.08);
  }
  .topic-icon {
    font-size: 4rem;
    display: block;
    margin-bottom: 0.5rem;
  }
  .topic-title {
    font-size: 2.2rem;
    color: #2c3e50;
    margin: 0;
  }
  .theory-card {
    background: white;
    border-radius: 30px;
    padding: 2rem;
    border: 1px solid rgba(88, 204, 2, 0.2);
    box-shadow: 0 10px 20px rgba(0, 0, 0, 0.02);
  }
  .theory-card h2 {
    color: #58cc02;
    margin-top: 0;
  }
  .theory-text {
    font-size: 1.05rem;
    line-height: 1.7;
    color: #333;
    white-space: pre-line;
  }
  .examples-block {
    margin-top: 1.5rem;
    background: #f9f9f9;
    border-radius: 16px;
    padding: 1rem 1.5rem;
  }
  .examples-block h3 {
    margin: 0 0 0.5rem;
    color: #2c3e50;
  }
  .examples-block ul {
    margin: 0;
    padding-left: 1.5rem;
  }
  .examples-block li {
    margin-bottom: 0.3rem;
    font-style: italic;
    color: #444;
  }
  .exercises-btn {
    background: #58cc02;
    color: white;
    border: none;
    border-radius: 30px;
    padding: 1rem 2rem;
    font-size: 1.3rem;
    font-weight: 600;
    cursor: pointer;
    box-shadow: 0 8px 20px rgba(88, 204, 2, 0.3);
    transition:
      background 0.2s,
      transform 0.2s;
    align-self: center;
  }
  .exercises-btn:hover {
    background: #46a302;
    transform: translateY(-2px);
  }
  .not-found {
    text-align: center;
    font-size: 1.5rem;
    color: #777;
    margin-top: 5rem;
  }
</style>
