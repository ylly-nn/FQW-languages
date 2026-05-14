<!-- src/views/vocabulary/TextLessonView.vue -->
<script setup>
  import { ref, reactive, computed, onMounted, nextTick } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import BackButton from '@/components/BackButton.vue'
  import data from '@/data/vocabulary.json'

  const route = useRoute()
  const router = useRouter()

  const lessonId = computed(() => parseInt(route.params.lessonId))
  const wordIndex = computed(() => parseInt(route.params.wordIndex))
  const lesson = computed(() => data.lessons.find((l) => l.id === lessonId.value))
  const word = computed(() => lesson.value?.words?.[wordIndex.value] || null)

  const VOWELS = new Set(['a', 'e', 'i', 'o', 'u'])

  function generateSteps(wordStr, syllables) {
    const letters = wordStr.split('')
    const vowelIndices = []
    const consonantIndices = []

    letters.forEach((ch, i) => {
      if (VOWELS.has(ch.toLowerCase())) {
        vowelIndices.push(i)
      } else {
        consonantIndices.push(i)
      }
    })

    const shuffle = (arr) => [...arr].sort(() => Math.random() - 0.5)

    // 1. По одной гласной (каждая гласная в случайном порядке)
    const vowelSteps = shuffle(vowelIndices).map((idx) => ({ type: 'vowel', indices: [idx] }))

    // 2. Все гласные сразу (если их больше одной)
    const allVowelsStep =
      vowelIndices.length > 1 ? [{ type: 'allVowels', indices: [...vowelIndices] }] : []

    // 3. По одной согласной (каждая согласная в случайном порядке)
    const consonantSteps = shuffle(consonantIndices).map((idx) => ({
      type: 'consonant',
      indices: [idx],
    }))

    // 4. По одному слогу (каждый слог в случайном порядке, с учетом повторений)
    let syllableSteps = []
    if (syllables.length > 1) {
      const syllableIndicesList = []
      let currentPos = 0
      for (const syll of syllables) {
        const start = wordStr.indexOf(syll, currentPos)
        if (start !== -1) {
          const indices = Array.from({ length: syll.length }, (_, i) => start + i)
          syllableIndicesList.push(indices)
          currentPos = start + syll.length
        }
      }
      syllableSteps = shuffle(syllableIndicesList).map((indices) => ({
        type: 'syllable',
        indices,
      }))
    }
    // 5. Всё слово целиком
    const allStep = { type: 'all', indices: letters.map((_, i) => i) }

    return [...vowelSteps, ...allVowelsStep, ...consonantSteps, ...syllableSteps, allStep]
  }
  const currentStep = ref(0)
  const steps = computed(() => {
    if (!word.value) return []
    return generateSteps(word.value.word, word.value.syllables)
  })
  const currentStepData = computed(() => steps.value[currentStep.value] || null)

  const inputValues = reactive({})
  const fieldStatus = reactive({})
  const isChecking = ref(false)
  const allStepsDone = computed(() => currentStep.value >= steps.value.length)

  const inputRefs = ref([])

  function setInputRef(el, index) {
    if (el) {
      inputRefs.value[index] = el
    }
  }

  function resetStep() {
    Object.keys(inputValues).forEach((k) => delete inputValues[k])
    Object.keys(fieldStatus).forEach((k) => delete fieldStatus[k])
    if (currentStepData.value) {
      currentStepData.value.indices.forEach((i) => {
        inputValues[i] = ''
        fieldStatus[i] = null
      })
    }
    nextTick(() => focusFirstAvailableInput())
  }

  function focusFirstAvailableInput() {
    if (!currentStepData.value) return
    const indices = currentStepData.value.indices
    for (const i of indices) {
      const el = inputRefs.value[i]
      if (el && !el.disabled) {
        el.focus()
        break
      }
    }
  }

  function onInputChar(event, currentIndex) {
    const value = event.target.value
    if (value.length === 1) {
      const indices = currentStepData.value?.indices || []
      const currentPos = indices.indexOf(currentIndex)
      if (currentPos !== -1) {
        for (let j = currentPos + 1; j < indices.length; j++) {
          const nextIdx = indices[j]
          const nextEl = inputRefs.value[nextIdx]
          if (nextEl && !nextEl.disabled) {
            nextEl.focus()
            break
          }
        }
      }
    }
  }

  function nextStep() {
    if (currentStep.value < steps.value.length) {
      currentStep.value++
      if (currentStep.value < steps.value.length) {
        resetStep()
      }
    }
  }

  function checkAnswers() {
    if (isChecking.value || !currentStepData.value) return
    const wordLetters = word.value.word.split('')
    let allCorrect = true
    const indices = currentStepData.value.indices

    indices.forEach((i) => {
      const expected = wordLetters[i].toLowerCase()
      const userInput = (inputValues[i] || '').trim().toLowerCase()
      if (userInput === expected) {
        fieldStatus[i] = 'correct'
      } else {
        fieldStatus[i] = 'wrong'
        allCorrect = false
      }
    })

    if (allCorrect) {
      isChecking.value = true
      setTimeout(() => {
        nextStep()
        isChecking.value = false
      }, 500)
    } else {
      isChecking.value = true
      setTimeout(() => {
        indices.forEach((i) => {
          if (fieldStatus[i] === 'wrong') {
            inputValues[i] = ''
            fieldStatus[i] = null
          }
        })
        isChecking.value = false
        nextTick(() => focusFirstAvailableInput())
      }, 1000)
    }
  }

  function goToNextWord() {
    const nextIndex = wordIndex.value + 1
    if (lesson.value && nextIndex < lesson.value.words.length) {
      router.push(`/vocabulary/lesson/${lessonId.value}/${nextIndex}`)
    } else {
      const stored = localStorage.getItem('vocabulary_completed')
      const completed = stored ? new Set(JSON.parse(stored)) : new Set()
      completed.add(lessonId.value)
      localStorage.setItem('vocabulary_completed', JSON.stringify([...completed]))
      const topicId = lesson.value?.topicId
      if (topicId) router.push(`/vocabulary/topics/${topicId}`)
      else router.push('/vocabulary/topics')
    }
  }

  onMounted(resetStep)
</script>

<template>
  <div class="root">
    <div class="top-bar">
      <BackButton :to="`/vocabulary/lesson/${lessonId}/${wordIndex}`" label="← Назад" />
    </div>
    <div v-if="word" class="exercise-container">
      <h2 class="lesson-title">{{ lesson?.title }}</h2>

      <div class="word-image">
        <img v-if="word.image" :src="word.image" :alt="word.word" class="word-img" />
        <span v-else class="image-placeholder">🖼️</span>
      </div>

      <p class="translation">{{ word.translation }}</p>

      <div class="word-row">
        <template v-for="(char, idx) in word.word.split('')" :key="idx">
          <span v-if="!currentStepData || !currentStepData.indices.includes(idx)" class="letter">
            {{ char }}
          </span>
          <input
            v-else
            :ref="(el) => setInputRef(el, idx)"
            v-model="inputValues[idx]"
            class="letter-input"
            :class="{
              correct: fieldStatus[idx] === 'correct',
              wrong: fieldStatus[idx] === 'wrong',
            }"
            maxlength="1"
            :disabled="fieldStatus[idx] === 'correct' || isChecking"
            @input="onInputChar($event, idx)"
            @keyup.enter="checkAnswers" />
        </template>
        <button
          v-if="currentStepData && !isChecking && !allStepsDone"
          class="ok-btn"
          @click="checkAnswers">
          OK
        </button>
        <button v-if="allStepsDone" class="next-btn" @click="goToNextWord">Далее →</button>
      </div>

      <div class="steps-info">Шаг {{ currentStep + 1 }} / {{ steps.length }}</div>
    </div>
    <div v-else class="not-found">Урок не найден</div>
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
    max-width: 700px;
    margin-bottom: 2rem;
  }
  .exercise-container {
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
    margin-bottom: 1rem;
  }
  .word-image {
    width: 200px;
    height: 200px;
    margin: 0 auto 1rem;
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
  .translation {
    font-size: 1.2rem;
    color: #555;
    margin-bottom: 2rem;
  }
  .word-row {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
    gap: 0.5rem;
    margin-bottom: 2rem;
  }
  .letter {
    font-size: 2.5rem;
    font-weight: bold;
    color: #2c3e50;
    min-width: 2rem;
    text-align: center;
  }
  .letter-input {
    width: 2.5rem;
    height: 2.5rem;
    font-size: 2rem;
    text-align: center;
    border: 2px solid #ccc;
    border-radius: 10px;
    outline: none;
  }
  .letter-input.correct {
    border-color: #58cc02;
    background: #f0fff0;
    color: #2c3e50;
  }
  .letter-input.wrong {
    border-color: #e74c3c;
    background: #fff0f0;
    color: #e74c3c;
  }
  .ok-btn,
  .next-btn {
    background: #58cc02;
    color: white;
    border: none;
    border-radius: 30px;
    padding: 0.8rem 2rem;
    font-size: 1.2rem;
    cursor: pointer;
    transition: background 0.2s;
    margin-left: 1rem;
  }
  .ok-btn:hover,
  .next-btn:hover {
    background: #46a302;
  }
  .steps-info {
    color: #888;
    font-size: 0.9rem;
  }
  .not-found {
    text-align: center;
    font-size: 1.5rem;
    color: #777;
    margin-top: 5rem;
  }
</style>
