<!-- components/TopicCard.vue -->
<script setup>
  defineProps({
    icon: { type: String, default: '' },
    title: { type: String, required: true },
    to: { type: String, default: undefined },
  })
</script>

<template>
  <component :is="to ? 'router-link' : 'div'" :to="to" class="topic-card">
    <span class="topic-icon">{{ icon }}</span>
    <span class="topic-title">{{ title }}</span>
  </component>
</template>

<style scoped>
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
    display: block;
    text-decoration: none;
    color: inherit;
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
