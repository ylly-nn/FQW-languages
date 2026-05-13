<script setup>
  defineProps({
    to: { type: String, default: '/' },
    label: { type: String, default: '← На главную' },
  })
</script>

<template>
  <RouterLink :to="to" class="back-button">
    <span class="text-button">{{ label }}</span>
  </RouterLink>
</template>

<style scoped>
  .back-button {
    width: auto;
    padding: 0.8rem 2rem;
    font-size: 1rem;
    border-radius: 50px;
    background-image: linear-gradient(rgba(255, 255, 255, 1) 0%, rgba(255, 255, 255, 1) 100%);
    border: 1px solid rgba(88, 204, 2, 0.25);
    box-shadow:
      0px 10px 20px 0px rgba(88, 204, 2, 0.04),
      0px 5px 10px 0px rgba(0, 0, 0, 0.06);
    display: inline-flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    transition:
      transform 0.2s,
      background 0.2s,
      box-shadow 0.2s,
      color 0.2s;
    text-decoration: none;
    color: #333;
    font-weight: 500;
  }

  .back-button:hover {
    transform: translateY(-2px);
    background: linear-gradient(135deg, #58cc02, #4cae01);
    color: white;
    box-shadow:
      0 4px 12px rgba(88, 204, 2, 0.3),
      0 6px 16px rgba(88, 204, 2, 0.4);
  }
</style>
