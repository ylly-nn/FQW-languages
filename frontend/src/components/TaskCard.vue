<script setup>
  import { ref } from 'vue'

  const props = defineProps({
    title: { type: String, required: true },
    description: { type: String, default: '' },
    completedDescription: { type: String, default: 'Задание выполнено!' },
    icon: { type: String, default: '' }, // можно передать эмодзи или текст
  })

  const completed = ref(false)

  const toggle = () => {
    completed.value = !completed.value
  }
</script>

<template>
  <div
    class="task-card"
    :class="{ 'task-completed': completed }"
    @click="toggle"
    role="button"
    tabindex="0"
    @keydown.enter="toggle">
    <div class="task-header">
      <h2 class="task-title">
        <span v-if="icon" class="task-icon">{{ icon }}</span>
        {{ title }}
      </h2>
      <span v-if="completed" class="check-icon">✓</span>
    </div>
    <p class="task-description">
      {{ completed ? completedDescription : description }}
    </p>
    <p class="task-hint">Нажмите, чтобы отметить выполнение</p>
  </div>
</template>

<style scoped>
  .task-card {
    width: 100%;
    max-width: 500px;
    padding: 2rem;
    border-radius: 40px;
    background: #ffffff;
    border: 1px solid rgba(88, 204, 2, 0.6);
    box-shadow:
      0px 20px 40px 0px rgba(88, 204, 2, 0.06),
      0px 10px 20px 0px rgba(0, 0, 0, 0.04);
    cursor: pointer;
    transition: all 0.3s ease;
    user-select: none;
    margin: 0 auto;
  }

  .task-card:hover {
    transform: translateY(-2px);
  }

  .task-completed {
    background: linear-gradient(135deg, rgba(88, 204, 2, 0.08), rgba(76, 174, 1, 0.05));
    border-color: rgba(88, 204, 2, 0.8);
    box-shadow:
      0px 20px 40px 0px rgba(88, 204, 2, 0.1),
      0px 10px 20px 0px rgba(0, 0, 0, 0.06);
  }

  .task-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 0.8rem;
  }

  .task-title {
    font-size: 1.6rem;
    color: #2c3e50;
    font-weight: 700;
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  .task-icon {
    font-size: 1.8rem;
  }

  .check-icon {
    font-size: 2rem;
    color: #58cc02;
    font-weight: bold;
    line-height: 1;
  }

  .task-description {
    font-size: 1rem;
    color: #555;
    line-height: 1.6;
    margin-bottom: 0.8rem;
  }

  .task-hint {
    font-size: 0.8rem;
    color: #aaa;
    font-style: italic;
  }
</style>
