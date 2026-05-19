<script setup>
  import { ref, reactive, computed, onMounted, nextTick } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import BackButton from '@/components/BackButton.vue'
  import data from '@/data/grammar.json'

  const route = useRoute()
  const router = useRouter()

  const topicId = computed(() => parseInt(route.params.topicId))
  const type = computed(() => route.params.type) // 'fill', 'choose', 'reorder'

  // Собираем все упражнения этого типа из всех уроков темы
  const topic = computed(() => data.topics.find((t) => t.id === topicId.value))

  const exercises = computed(() => {
    const lessons = data.lessons.filter((l) => l.topicId === topicId.value)
    const result = []
    lessons.forEach((lesson) => {
      lesson.exercises
        .filter((ex) => ex.type === type.value)
        .forEach((ex, idx) => {
          result.push({ ...ex, lessonTitle: lesson.title })
        })
    })
    return result
  })

  // Состояния для ответов пользователя
  const userAnswers = reactive({}) // ключ – индекс упражнения, значение – строка/индекс/массив
  const fieldStatus = reactive({}) // 'correct' или 'wrong'
  const isChecking = ref(false)

  // Для reorder: перемешанные слова для каждого упражнения
  const reorderedWords = ref({}) // ключ – индекс упражнения, значение – массив слов

  onMounted(() => {
    // Инициализируем reorderedWords и сбрасываем ответы
    exercises.value.forEach((ex, idx) => {
      if (ex.type === 'reorder') {
        reorderedWords.value[idx] = [...ex.words].sort(() => Math.random() - 0.5)
      } else if (ex.type === 'choose') {
        userAnswers[idx] = null
      } else {
        userAnswers[idx] = ''
      }
    })
  })

  // Проверить все упражнения
  function checkAll() {
    if (isChecking.value) return

    // Проверяем каждое упражнение
    exercises.value.forEach((ex, idx) => {
      if (ex.type === 'fill') {
        const userAns = (userAnswers[idx] || '').trim()
        const correctAns = ex.answer.trim()
        fieldStatus[idx] = userAns.toLowerCase() === correctAns.toLowerCase() ? 'correct' : 'wrong'
      } else if (ex.type === 'choose') {
        fieldStatus[idx] = userAnswers[idx] === ex.answer ? 'correct' : 'wrong'
      } else if (ex.type === 'reorder') {
        const userSentence = (reorderedWords.value[idx] || []).join(' ')
        fieldStatus[idx] =
          userSentence.toLowerCase() === ex.answer.toLowerCase() ? 'correct' : 'wrong'
      }
    })

    isChecking.value = true
  }

  // Сбросить неправильные ответы
  function resetWrong() {
    exercises.value.forEach((ex, idx) => {
      if (fieldStatus[idx] === 'wrong') {
        if (ex.type === 'reorder') {
          reorderedWords.value[idx] = [...ex.words].sort(() => Math.random() - 0.5)
        } else if (ex.type === 'choose') {
          userAnswers[idx] = null
        } else {
          userAnswers[idx] = ''
        }
        delete fieldStatus[idx]
      }
    })
    isChecking.value = false
    nextTick(() => {
      // фокус на первый неправильный (опционально)
    })
  }

  // Перемешать слова в reorder
  function shuffleWords(exIdx) {
    reorderedWords.value[exIdx] = [...reorderedWords.value[exIdx]].sort(() => Math.random() - 0.5)
  }
</script>

<template>
  <div class="root">
    <div class="top-bar">
      <BackButton :to="`/grammar/topics/${topicId}/exercises`" label="← К заданиям" />
    </div>

    <main v-if="topic && exercises.length" class="content">
      <h1 class="page-title">
        {{
          type === 'fill'
            ? '✍️ Вставить слово'
            : type === 'choose'
              ? '🔘 Выбрать ответ'
              : '🔄 Составить предложение'
        }}
      </h1>

      <div class="exercises-card">
        <div
          v-for="(ex, idx) in exercises"
          :key="idx"
          class="exercise-item"
          :class="{
            correct: fieldStatus[idx] === 'correct',
            wrong: fieldStatus[idx] === 'wrong',
          }">
          <!-- Fill -->
          <template v-if="ex.type === 'fill'">
            <p class="sentence">{{ ex.sentence }}</p>
            <input
              v-model="userAnswers[idx]"
              class="answer-input"
              :class="fieldStatus[idx]"
              :disabled="fieldStatus[idx] === 'correct'"
              placeholder="Ваш ответ" />
          </template>

          <!-- Choose -->
          <template v-else-if="ex.type === 'choose'">
            <p class="sentence">{{ ex.sentence }}</p>
            <div class="options">
              <button
                v-for="(opt, optIdx) in ex.options"
                :key="optIdx"
                :class="{
                  selected: userAnswers[idx] === optIdx,
                  correct: fieldStatus[idx] === 'correct' && optIdx === ex.answer,
                  wrong:
                    fieldStatus[idx] === 'wrong' &&
                    userAnswers[idx] === optIdx &&
                    optIdx !== ex.answer,
                }"
                :disabled="fieldStatus[idx] === 'correct'"
                @click="userAnswers[idx] = optIdx">
                {{ opt }}
              </button>
            </div>
          </template>

          <!-- Reorder -->
          <template v-else-if="ex.type === 'reorder'">
            <div class="reorder-words">
              <span
                v-for="(word, wordIdx) in reorderedWords[idx]"
                :key="wordIdx"
                class="reorder-word">
                {{ word }}
              </span>
            </div>
            <button
              class="shuffle-btn"
              :disabled="fieldStatus[idx] === 'correct'"
              @click="shuffleWords(idx)">
              Перемешать
            </button>
          </template>
        </div>

        <!-- Кнопки управления -->
        <div class="actions">
          <button v-if="!isChecking" class="check-btn" @click="checkAll">Проверить всё</button>
          <template v-else>
            <button class="retry-btn" @click="resetWrong">Повторить неправильные</button>
            <button class="back-btn" @click="$router.push(`/grammar/topics/${topicId}/exercises`)">
              Вернуться к заданиям
            </button>
          </template>
        </div>
      </div>
    </main>

    <div v-else class="not-found">
      {{ exercises.length === 0 ? 'Нет заданий этого типа' : 'Тема не найдена' }}
    </div>
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
  }

  .page-title {
    text-align: center;
    color: #2c3e50;
    margin-bottom: 2rem;
  }

  .exercises-card {
    background: white;
    border-radius: 30px;
    padding: 2rem;
    box-shadow: 0 10px 30px rgba(88, 204, 2, 0.08);
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }

  .exercise-item {
    border: 2px solid #e0e0e0;
    border-radius: 20px;
    padding: 1.5rem;
    transition:
      border-color 0.2s,
      background 0.2s;
  }

  .exercise-item.correct {
    border-color: #58cc02;
    background: #f6fff0;
  }

  .exercise-item.wrong {
    border-color: #e74c3c;
    background: #fff5f5;
  }

  .sentence {
    font-size: 1.2rem;
    color: #2c3e50;
    margin: 0 0 1rem;
  }

  .answer-input {
    width: 200px;
    padding: 0.7rem 1rem;
    border: 2px solid #ccc;
    border-radius: 12px;
    font-size: 1.1rem;
    outline: none;
  }

  .answer-input.correct {
    border-color: #58cc02;
    background: #f0fff0;
  }

  .answer-input.wrong {
    border-color: #e74c3c;
    background: #fff0f0;
  }

  .options {
    display: flex;
    gap: 0.8rem;
    flex-wrap: wrap;
  }

  .options button {
    background: white;
    border: 2px solid #58cc02;
    border-radius: 20px;
    padding: 0.5rem 1.2rem;
    cursor: pointer;
    transition: 0.2s;
  }

  .options button.selected {
    background: #58cc02;
    color: white;
  }

  .options button.correct {
    background: #58cc02;
    color: white;
  }

  .options button.wrong {
    background: #e74c3c;
    color: white;
    border-color: #e74c3c;
  }

  .options button:disabled {
    opacity: 0.7;
    cursor: default;
  }

  .reorder-words {
    display: flex;
    gap: 0.5rem;
    flex-wrap: wrap;
    margin-bottom: 1rem;
  }

  .reorder-word {
    background: #f0f0f0;
    padding: 0.5rem 1rem;
    border-radius: 12px;
    font-size: 1.1rem;
    user-select: none;
  }

  .shuffle-btn {
    background: none;
    border: 1px dashed #aaa;
    border-radius: 20px;
    padding: 0.4rem 1rem;
    cursor: pointer;
  }

  .actions {
    display: flex;
    gap: 1rem;
    justify-content: center;
    margin-top: 1rem;
  }

  .check-btn,
  .retry-btn,
  .back-btn {
    background: #58cc02;
    color: white;
    border: none;
    border-radius: 30px;
    padding: 0.8rem 2rem;
    font-size: 1.2rem;
    cursor: pointer;
    transition: background 0.2s;
  }

  .retry-btn {
    background: #ff9800;
  }

  .back-btn {
    background: #888;
  }

  .check-btn:hover {
    background: #46a302;
  }
  .retry-btn:hover {
    background: #e68900;
  }
  .back-btn:hover {
    background: #666;
  }

  .not-found {
    text-align: center;
    font-size: 1.5rem;
    color: #777;
    margin-top: 5rem;
  }
</style>
