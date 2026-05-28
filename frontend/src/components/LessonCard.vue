<!-- components/LessonCard.vue -->
<script setup>
  defineProps({
    title: { type: String, required: true },
    to: { type: String, default: undefined },
    icon: { type: String, default: '📖' },
    detail: { type: String, default: '' }, // например, "10 слов"
  })
</script>

<template>
  <component :is="to ? 'router-link' : 'div'" :to="to" class="lesson-card">
    <span class="lesson-icon">{{ icon }}</span>
    <div class="lesson-info">
      <h3 class="lesson-title">{{ title }}</h3>
      <p v-if="detail" class="lesson-detail">{{ detail }}</p>
    </div>
  </component>
</template>

<style scoped>
  .lesson-card {
    background: white;
    border: 1px solid rgba(88, 204, 2, 0.3);
    border-radius: 20px;
    padding: 1.5rem;
    display: flex;
    align-items: center;
    gap: 1.5rem;
    cursor: pointer;
    transition:
      transform 0.2s,
      box-shadow 0.2s;
    text-decoration: none;
    color: inherit;
    box-shadow: 0 8px 16px rgba(88, 204, 2, 0.03);
  }
  .lesson-card:hover {
    transform: translateY(-2px);
    border-color: rgba(88, 204, 2, 0.6);
    box-shadow: 0 12px 20px rgba(88, 204, 2, 0.08);
  }
  .lesson-icon {
    font-size: 2.5rem;
  }
  .lesson-title {
    font-size: 1.1rem;
    font-weight: 600;
    margin: 0 0 0.2rem 0;
    color: #2c3e50;
  }
  .lesson-detail {
    font-size: 0.9rem;
    color: #777;
    margin: 0;
  }
</style>
