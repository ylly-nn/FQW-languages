<script setup>
  import { ref, computed } from 'vue'
  import { useRoute } from 'vue-router'
  import AudioPlayer from '@/components/AudioPlayer.vue'
  import BackButton from '@/components/BackButton.vue'
  import data from '@/data/lessons.json'

  const route = useRoute()
  const lessonId = computed(() => parseInt(route.params.id))
  const lesson = computed(() => data.lessons.find((l) => l.id === lessonId.value))

  const selectedAnswers = ref({})
  const showResults = ref(false)

  const selectAnswer = (qIndex, optIndex) => {
    if (!showResults.value) {
      selectedAnswers.value[qIndex] = optIndex
    }
  }
  const checkAnswers = () => {
    showResults.value = true
  }
</script>

<template>
  <div class="root">
    <BackButton to="/dashboard" label="← К заданиям" />
    <div v-if="lesson" class="lesson-container">
      <AudioPlayer :src="lesson.audio" :title="lesson.title" />
      <div class="questions">
        <div v-for="(q, i) in lesson.questions" :key="i" class="question">
          <p>{{ i + 1 }}. {{ q.text }}</p>
          <div class="options">
            <button
              v-for="(opt, j) in q.options"
              :key="j"
              :class="{
                selected: selectedAnswers[i] === j,
                correct: showResults && j === q.answer,
                wrong: showResults && selectedAnswers[i] === j && j !== q.answer,
              }"
              @click="selectAnswer(i, j)">
              {{ opt }}
            </button>
          </div>
        </div>
      </div>
      <button v-if="!showResults" @click="checkAnswers" class="check-btn">Проверить</button>
    </div>
    <div v-else class="not-found">Урок не найден</div>
  </div>
</template>

<style scoped>
  .root {
    min-height: 100vh;
    background: #fffdf7;
    padding: 2rem;
  }
  .lesson-container {
    max-width: 700px;
    margin: 0 auto;
  }
  .questions {
    margin-top: 2rem;
  }
  .question {
    margin-bottom: 1.5rem;
  }
  .options {
    display: flex;
    gap: 0.5rem;
    flex-wrap: wrap;
  }
  .options button {
    background: white;
    border: 1px solid #58cc02;
    border-radius: 20px;
    padding: 0.5rem 1rem;
    cursor: pointer;
    transition: 0.2s;
  }
  .selected {
    background: #58cc02;
    color: white;
  }
  .correct {
    background: #58cc02;
    color: white;
  }
  .wrong {
    background: #e74c3c;
    color: white;
    border-color: #e74c3c;
  }
  .check-btn {
    background: #58cc02;
    color: white;
    border: none;
    padding: 0.8rem 2rem;
    border-radius: 30px;
    font-size: 1.1rem;
    margin-top: 1rem;
    cursor: pointer;
  }
  .not-found {
    text-align: center;
    font-size: 1.5rem;
    color: #777;
    margin-top: 5rem;
  }
</style>
