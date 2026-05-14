<!-- src/views/vocabulary/LessonView.vue -->
<script setup>
  import { ref, computed } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import BackButton from '@/components/BackButton.vue'
  import data from '@/data/vocabulary.json'

  const route = useRoute()
  const router = useRouter()

  const lessonId = computed(() => parseInt(route.params.lessonId))
  const wordIndex = computed(() => parseInt(route.params.wordIndex) || 0)
  const lesson = computed(() => data.lessons.find((l) => l.id === lessonId.value))
  const words = computed(() => lesson.value?.words || [])
  const currentWord = computed(() => words.value[wordIndex.value] || null)
  const totalWords = computed(() => words.value.length)

  // Для любого слова переходим к текстовому упражнению
  const goToTextExercise = () => {
    if (!currentWord.value) return
    router.push(`/vocabulary/lesson/${lessonId.value}/${wordIndex.value}/text`)
  }
</script>

<template>
  <div class="root">
    <div class="top-bar">
      <BackButton :to="`/vocabulary/topics/${lesson?.topicId}`" label="← К урокам" />
    </div>
    <div v-if="lesson" class="lesson-container">
      <h1 class="lesson-title">{{ lesson.title }}</h1>
      <div v-if="currentWord" class="word-card">
        <div class="word-image">
          <img v-if="currentWord.image" :src="currentWord.image" class="word-img" />
          <span v-else class="image-placeholder">🖼️</span>
        </div>
        <div class="word-main">
          <h2 class="word-text">{{ currentWord.word }}</h2>
          <p class="word-translation">{{ currentWord.translation }}</p>
          <p class="word-transcription">{{ currentWord.transcription }}</p>
        </div>
        <div class="audio-placeholder">
          <button v-if="currentWord.audio" class="audio-btn active">🔊 Прослушать</button>
          <button v-else class="audio-btn" disabled>🔊 Аудио (скоро)</button>
        </div>
        <div class="progress">{{ wordIndex + 1 }} / {{ totalWords }}</div>
        <button class="next-btn" @click="goToTextExercise">Далее →</button>
      </div>
      <p v-else class="empty-message">В этом уроке нет слов</p>
    </div>
    <div v-else class="not-found">Урок не найден</div>
  </div>
</template>

<style scoped>
  /* Все стили остаются прежними, как у вас были */
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
    max-width: 700px;
    margin-bottom: 2rem;
  }
  .lesson-container {
    width: 100%;
    max-width: 600px;
    background: white;
    border: 1px solid rgba(88, 204, 2, 0.4);
    border-radius: 40px;
    padding: 2.5rem 2rem;
    box-shadow: 0 20px 40px rgba(88, 204, 2, 0.08);
    text-align: center;
  }
  .lesson-title {
    color: #2c3e50;
    margin-bottom: 2rem;
  }
  .word-card {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 1.5rem;
  }
  .word-image {
    width: 200px;
    height: 200px;
    background: #f0f0f0;
    border-radius: 20px;
    display: flex;
    align-items: center;
    justify-content: center;
    overflow: hidden;
  }
  .word-img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
  .image-placeholder {
    font-size: 4rem;
    opacity: 0.5;
  }
  .word-main {
    text-align: center;
  }
  .word-text {
    font-size: 2.5rem;
    color: #2c3e50;
    margin: 0;
  }
  .word-translation {
    font-size: 1.5rem;
    color: #555;
    margin: 0.3rem 0 0;
  }
  .word-transcription {
    font-size: 1.1rem;
    color: #888;
    margin: 0.2rem 0 0;
    font-style: italic;
  }
  .audio-placeholder {
    margin: 0.5rem 0;
  }
  .audio-btn {
    background: #f5f5f5;
    border: 1px dashed #aaa;
    border-radius: 20px;
    padding: 0.5rem 1.5rem;
    font-size: 1rem;
    cursor: pointer;
    opacity: 0.7;
  }
  .audio-btn.active {
    background: #fff;
    border: 2px solid #58cc02;
    color: #2c3e50;
    cursor: pointer;
    opacity: 1;
    font-weight: 600;
  }
  .audio-btn.active:hover {
    background: #f0fff0;
  }
  .audio-btn:disabled {
    cursor: not-allowed;
  }
  .progress {
    color: #777;
    font-size: 0.9rem;
  }
  .next-btn {
    background: #58cc02;
    color: white;
    border: none;
    border-radius: 30px;
    padding: 0.8rem 2.5rem;
    font-size: 1.2rem;
    cursor: pointer;
    transition: background 0.2s;
  }
  .next-btn:hover {
    background: #46a302;
  }
  .empty-message {
    color: #777;
    font-style: italic;
  }
  .not-found {
    text-align: center;
    font-size: 1.5rem;
    color: #777;
    margin-top: 5rem;
  }
</style>
