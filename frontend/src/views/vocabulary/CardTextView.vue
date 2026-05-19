<!-- src/views/vocabulary/CardTextView.vue -->
<script setup>
  import { ref, reactive, computed, onMounted } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import BackButton from '@/components/BackButton.vue'
  import data from '@/data/vocabulary.json'

  const route = useRoute()
  const router = useRouter()

  const lessonId = computed(() => parseInt(route.params.lessonId))
  const lesson = computed(() => data.lessons.find((l) => l.id === lessonId.value))
  const words = computed(() => lesson.value?.words || [])

  const shuffledEnglish = ref([])

  const selected = reactive({})
  const lineStatuses = reactive({})

  // Блокировка интерфейса на время сброса ошибок
  const isChecking = ref(false)

  const allSelected = computed(() => {
    return words.value.length > 0 && Object.keys(selected).length === words.value.length
  })

  function getAvailableOptions(currentLeftIdx) {
    const alreadyTaken = Object.entries(selected)
      .filter(([idx]) => idx != currentLeftIdx)
      .map(([, word]) => word)
    return shuffledEnglish.value.filter((word) => !alreadyTaken.includes(word))
  }

  function selectWord(leftIdx, word) {
    selected[leftIdx] = word
    delete lineStatuses[leftIdx]
  }

  function checkAnswers() {
    if (!allSelected.value || isChecking.value) return

    // 1. Проверяем все пары
    let correctCount = 0
    for (let i = 0; i < words.value.length; i++) {
      if (selected[i] === words.value[i].word) {
        lineStatuses[i] = 'correct'
        correctCount++
      } else {
        lineStatuses[i] = 'wrong'
      }
    }

    // 2. Если всё правильно — завершаем урок
    if (correctCount === words.value.length) {
      const stored = localStorage.getItem('vocabulary_completed')
      const completed = stored ? new Set(JSON.parse(stored)) : new Set()
      completed.add(lessonId.value)
      localStorage.setItem('vocabulary_completed', JSON.stringify([...completed]))

      setTimeout(() => {
        const topicId = lesson.value?.topicId
        if (topicId) {
          router.push(`/vocabulary/topics/${topicId}`)
        } else {
          router.push('/vocabulary/topics')
        }
      }, 600)
      return
    }

    // 3. Есть ошибки — блокируем интерфейс и готовим сброс неправильных
    isChecking.value = true

    // Через 1.2 секунды удаляем неправильные выборы и снимаем блокировку
    setTimeout(() => {
      for (let i = 0; i < words.value.length; i++) {
        if (lineStatuses[i] === 'wrong') {
          delete selected[i] // очищаем выбор
          delete lineStatuses[i] // убираем статус ошибки
        }
      }
      isChecking.value = false
    }, 1200)
  }

  onMounted(() => {
    if (words.value.length > 0) {
      shuffledEnglish.value = [...words.value].map((w) => w.word).sort(() => Math.random() - 0.5)
    }
  })
</script>

<template>
  <div class="root">
    <div class="top-bar">
      <BackButton :to="`/vocabulary/lesson/${lessonId}/0`" label="← К уроку" />
    </div>
    <main class="content">
      <h1>Сопоставьте слова</h1>
      <p class="subtitle">Выберите правильное слово для каждой карточки и нажмите проверить</p>

      <div v-if="words.length > 0" class="cards-grid">
        <div
          v-for="(word, idx) in words"
          :key="idx"
          class="word-card"
          :class="{
            correct: lineStatuses[idx] === 'correct',
            wrong: lineStatuses[idx] === 'wrong',
          }">
          <div class="card-image">
            <img v-if="word.image" :src="word.image" :alt="word.translation" />
            <span v-else class="image-placeholder">🖼️</span>
          </div>
          <div class="card-info">
            <span class="russian-word">{{ word.translation }}</span>
            <select
              v-model="selected[idx]"
              class="word-select"
              :class="{
                correct: lineStatuses[idx] === 'correct',
                wrong: lineStatuses[idx] === 'wrong',
              }"
              :disabled="lineStatuses[idx] === 'correct' || isChecking"
              @change="selectWord(idx, selected[idx])">
              <option :value="null" disabled>Выберите слово</option>
              <option v-for="eng in getAvailableOptions(idx)" :key="eng" :value="eng">
                {{ eng }}
              </option>
            </select>
          </div>
        </div>
      </div>
      <p v-else class="empty-message">Нет слов для этого урока</p>

      <button
        v-if="words.length > 0"
        class="check-btn"
        :disabled="!allSelected || isChecking"
        @click="checkAnswers">
        Проверить
      </button>
    </main>
  </div>
</template>

<style scoped>
  /* Стили без изменений */
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
    max-width: 900px;
    margin-bottom: 2rem;
  }
  .content {
    width: 100%;
    max-width: 900px;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  h1 {
    color: #2c3e50;
    text-align: center;
    margin-bottom: 0.3rem;
  }
  .subtitle {
    text-align: center;
    color: #666;
    margin-bottom: 2rem;
    font-size: 1rem;
  }
  .cards-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 1.5rem;
    width: 100%;
    margin-bottom: 2rem;
  }
  .word-card {
    background: white;
    border: 2px solid #e0e0e0;
    border-radius: 20px;
    padding: 1.2rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
    transition:
      border-color 0.2s,
      box-shadow 0.2s;
  }
  .word-card.correct {
    border-color: #58cc02;
    box-shadow: 0 0 0 3px rgba(88, 204, 2, 0.15);
    background: #f6fff0;
  }
  .word-card.wrong {
    border-color: #e74c3c;
    box-shadow: 0 0 0 3px rgba(231, 76, 60, 0.15);
    background: #fff5f5;
  }
  .card-image {
    width: 120px;
    height: 120px;
    border-radius: 16px;
    overflow: hidden;
    background: #f0f0f0;
    margin-bottom: 0.8rem;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .card-image img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
  .image-placeholder {
    font-size: 3rem;
    opacity: 0.5;
  }
  .card-info {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.5rem;
  }
  .russian-word {
    font-size: 1.2rem;
    font-weight: 600;
    color: #2c3e50;
  }
  .word-select {
    width: 100%;
    max-width: 200px;
    padding: 0.5rem 1rem;
    font-size: 1rem;
    border: 2px solid #ccc;
    border-radius: 12px;
    background: white;
    outline: none;
    transition: border-color 0.2s;
    appearance: auto;
  }
  .word-select.correct {
    border-color: #58cc02;
    background: #f0fff0;
  }
  .word-select.wrong {
    border-color: #e74c3c;
    background: #fff5f5;
  }
  .check-btn {
    margin-top: 1rem;
    background: #58cc02;
    color: white;
    border: none;
    border-radius: 30px;
    padding: 1rem 3rem;
    font-size: 1.3rem;
    font-weight: 600;
    cursor: pointer;
    box-shadow: 0 6px 20px rgba(88, 204, 2, 0.3);
    transition:
      background 0.2s,
      opacity 0.2s;
  }
  .check-btn:disabled {
    opacity: 0.4;
    cursor: not-allowed;
  }
  .empty-message {
    text-align: center;
    color: #777;
    margin-top: 3rem;
  }
</style>
