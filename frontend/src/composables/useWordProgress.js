import { ref, onMounted } from 'vue'
import defaultProgress from '@/data/progress.json'

const PROGRESS_KEY = 'word_progress'

export function useWordProgress() {
  const progress = ref({})

  // Загрузка из localStorage, если нет – инициализация из JSON
  const loadProgress = () => {
    const saved = localStorage.getItem(PROGRESS_KEY)
    if (saved) {
      try {
        progress.value = JSON.parse(saved)
      } catch (e) {
        progress.value = { ...defaultProgress }
      }
    } else {
      progress.value = { ...defaultProgress }
    }
  }

  // Сохранение прогресса
  const saveProgress = () => {
    localStorage.setItem(PROGRESS_KEY, JSON.stringify(progress.value))
  }

  // Обновление для одного слова
  const updateWordProgress = (word, isCorrect) => {
    if (!progress.value[word]) {
      progress.value[word] = { correct: 0, total: 0, percent: 0 }
    }
    const p = progress.value[word]
    p.total += 1
    if (isCorrect) p.correct += 1
    p.percent = Math.round((p.correct / p.total) * 100)
    saveProgress()
  }

  // Получить процент слова
  const getWordPercent = (word) => {
    return progress.value[word]?.percent ?? 0
  }

  // Получить цвет для процента
  const getPercentColor = (percent) => {
    if (percent >= 80) return 'green'
    if (percent >= 50) return 'orange'
    return 'red'
  }

  // Загрузить все слова из папок (чтобы показать даже те, у которых нет прогресса)
  const getAllWordsFromDictionary = (dictionary) => {
    const wordsSet = new Set()
    dictionary.folders.forEach((folder) => {
      folder.words.forEach((w) => wordsSet.add(w.word))
    })
    return Array.from(wordsSet)
  }

  onMounted(() => {
    loadProgress()
  })

  return {
    progress,
    loadProgress,
    updateWordProgress,
    getWordPercent,
    getPercentColor,
    getAllWordsFromDictionary,
  }
}
