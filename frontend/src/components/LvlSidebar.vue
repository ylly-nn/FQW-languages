<script setup>
  defineProps({
    levels: {
      type: Array, // массив строк
      required: true,
    },
    modelValue: {
      type: String,
      default: null,
    },
  })

  defineEmits(['update:modelValue'])
</script>

<template>
  <aside class="levels-sidebar">
    <h2 class="levels-title">Уровни</h2>
    <button
      v-for="level in levels"
      :key="level"
      class="level-btn"
      :class="{ active: modelValue === level }"
      @click="$emit('update:modelValue', level)">
      {{ level }}
    </button>
  </aside>
</template>

<style scoped>
  .levels-sidebar {
    width: 160px;
    background: rgba(255, 255, 255, 0.9);
    backdrop-filter: blur(6px);
    border: 1px solid rgba(88, 204, 2, 0.4);
    border-radius: 30px;
    padding: 1.5rem 1rem;
    box-shadow: 0 10px 20px rgba(88, 204, 2, 0.06);
    flex-shrink: 0;
  }

  .levels-title {
    font-size: 1.2rem;
    color: #2c3e50;
    margin-bottom: 1.2rem;
    text-align: center;
    font-weight: 700;
  }

  .level-btn {
    display: block;
    width: 100%;
    padding: 0.8rem 0.5rem;
    margin-bottom: 0.5rem;
    border: none;
    background: transparent;
    border-radius: 20px;
    font-size: 1rem;
    font-weight: 600;
    color: #2c3e50;
    cursor: pointer;
    transition: all 0.2s;
    text-align: center;
  }

  .level-btn:hover {
    background: rgba(88, 204, 2, 0.08);
  }

  .level-btn.active {
    background: rgba(88, 204, 2, 0.15);
    color: #58cc02;
    font-weight: 700;
  }

  /* Адаптив для мобильных */
  @media (max-width: 768px) {
    .levels-sidebar {
      width: 100%;
      display: flex;
      flex-wrap: wrap;
      gap: 0.5rem;
      padding: 0.8rem;
      border-radius: 20px;
    }
    .levels-title {
      width: 100%;
      margin-bottom: 0.5rem;
    }
    .level-btn {
      flex: 1 0 auto;
      padding: 0.6rem 0.8rem;
    }
  }
</style>
