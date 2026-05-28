<script setup>
  import { computed } from 'vue'

  const props = defineProps({
    lessons: { type: Array, required: true },
    completedIds: { type: Set, default: () => new Set() },
    premiumIds: { type: Set, default: () => new Set() }, // для «подписанных»
  })

  const emit = defineEmits(['select', 'toggle-premium'])

  const lessonStatus = (lesson) => {
    if (props.completedIds.has(lesson.id)) return 'completed'
    if (props.premiumIds.has(lesson.id)) return 'premium'
    return 'new'
  }
</script>

<template>
  <div class="lesson-list">
    <div
      v-for="lesson in lessons"
      :key="lesson.id"
      class="lesson-item"
      :class="lessonStatus(lesson)"
      @click="emit('select', lesson)">
      <div class="lesson-left">
        <span class="status-icon">
          <template v-if="lessonStatus(lesson) === 'completed'">✓</template>
          <template v-else-if="lessonStatus(lesson) === 'premium'">⭐</template>
          <template v-else>●</template>
        </span>
        <span class="lesson-title">{{ lesson.title }}</span>
      </div>
      <button
        v-if="lessonStatus(lesson) === 'premium'"
        class="premium-btn"
        @click.stop="emit('toggle-premium', lesson)">
        <template v-if="props.premiumIds.has(lesson.id)">Отписаться</template>
        <template v-else>Подписаться</template>
      </button>
    </div>
  </div>
</template>

<style scoped>
  .lesson-list {
    display: flex;
    flex-direction: column;
    gap: 0.8rem;
  }
  .lesson-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 1.2rem;
    border-radius: 20px;
    background: white;
    border: 1px solid rgba(88, 204, 2, 0.3);
    cursor: pointer;
    transition: all 0.2s;
  }
  .lesson-item:hover {
    border-color: rgba(88, 204, 2, 0.7);
    box-shadow: 0 4px 12px rgba(88, 204, 2, 0.1);
  }
  .lesson-item.completed {
    background: rgba(88, 204, 2, 0.05);
    border-color: rgba(88, 204, 2, 0.5);
  }
  .lesson-item.premium {
    border-color: gold;
    background: rgba(255, 215, 0, 0.05);
  }
  .lesson-left {
    display: flex;
    align-items: center;
    gap: 0.8rem;
  }
  .status-icon {
    font-size: 1.2rem;
    width: 1.5rem;
    text-align: center;
    color: #58cc02;
  }
  .premium .status-icon {
    color: gold;
  }
  .new .status-icon {
    color: #ccc;
  }
  .lesson-title {
    font-size: 1rem;
    color: #2c3e50;
  }
  .premium-btn {
    background: rgba(88, 204, 2, 0.1);
    border: 1px solid rgba(88, 204, 2, 0.4);
    border-radius: 15px;
    padding: 0.4rem 1rem;
    font-size: 0.85rem;
    cursor: pointer;
    transition: background 0.2s;
  }
  .premium-btn:hover {
    background: rgba(88, 204, 2, 0.2);
  }
</style>
