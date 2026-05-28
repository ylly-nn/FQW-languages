<script setup>
  import { ref, onMounted, computed } from 'vue'
  import { useWordProgress } from '@/composables/useWordProgress'
  import dictionaryData from '@/data/dictionary.json'

  const { progress, getWordPercent, getPercentColor, getAllWordsFromDictionary } = useWordProgress()
  const allWords = ref([])
  const sortBy = ref('percent') // percent, word, total
  const filterPercent = ref('all') // all, high, medium, low

  // Собираем все уникальные слова из словаря (можно также добавить из vocabulary)
  const loadAllWords = () => {
    const wordsSet = new Set()
    dictionaryData.folders.forEach((folder) => {
      folder.words.forEach((w) => wordsSet.add(w.word))
    })
    // Можно также добавить слова из vocabulary.json, если нужно
    allWords.value = Array.from(wordsSet).map((word) => ({
      word,
      percent: getWordPercent(word),
      total: progress.value[word]?.total || 0,
      correct: progress.value[word]?.correct || 0,
    }))
  }

  const filteredWords = computed(() => {
    let filtered = [...allWords.value]
    if (filterPercent.value === 'high') {
      filtered = filtered.filter((w) => w.percent >= 80)
    } else if (filterPercent.value === 'medium') {
      filtered = filtered.filter((w) => w.percent >= 50 && w.percent < 80)
    } else if (filterPercent.value === 'low') {
      filtered = filtered.filter((w) => w.percent < 50)
    }
    if (sortBy.value === 'percent') {
      filtered.sort((a, b) => b.percent - a.percent)
    } else if (sortBy.value === 'word') {
      filtered.sort((a, b) => a.word.localeCompare(b.word))
    } else if (sortBy.value === 'total') {
      filtered.sort((a, b) => b.total - a.total)
    }
    return filtered
  })

  const getColorClass = (percent) => {
    if (percent >= 80) return 'progress-high'
    if (percent >= 50) return 'progress-mid'
    return 'progress-low'
  }

  onMounted(() => {
    loadAllWords()
  })
</script>

<template>
  <div class="progress-root">
    <div class="container">
      <h1>📊 Прогресс изучения слов</h1>
      <div class="controls">
        <div class="filter-group">
          <label>Фильтр:</label>
          <select v-model="filterPercent">
            <option value="all">Все слова</option>
            <option value="high">Знаю хорошо (≥80%)</option>
            <option value="medium">Учу (50–79%)</option>
            <option value="low">Трудно (<50%)</option>
          </select>
        </div>
        <div class="sort-group">
          <label>Сортировка:</label>
          <select v-model="sortBy">
            <option value="percent">По проценту</option>
            <option value="word">По алфавиту</option>
            <option value="total">По числу попыток</option>
          </select>
        </div>
      </div>

      <div class="stats-summary">
        <div class="stat-card">
          <span class="stat-label">Всего слов</span>
          <span class="stat-value">{{ allWords.length }}</span>
        </div>
        <div class="stat-card">
          <span class="stat-label">Изучено (≥80%)</span>
          <span class="stat-value">{{ allWords.filter((w) => w.percent >= 80).length }}</span>
        </div>
        <div class="stat-card">
          <span class="stat-label">Средний %</span>
          <span class="stat-value">
            {{
              allWords.length
                ? Math.round(allWords.reduce((s, w) => s + w.percent, 0) / allWords.length)
                : 0
            }}%
          </span>
        </div>
      </div>

      <div class="words-list">
        <div class="list-header">
          <div>Слово</div>
          <div>Прогресс</div>
          <div>Правильно / Всего</div>
        </div>
        <div v-for="item in filteredWords" :key="item.word" class="list-row">
          <div class="word-cell">{{ item.word }}</div>
          <div class="progress-cell">
            <div class="progress-bar-bg">
              <div
                class="progress-bar-fill"
                :class="getColorClass(item.percent)"
                :style="{ width: item.percent + '%' }"></div>
            </div>
            <span class="percent-text">{{ item.percent }}%</span>
          </div>
          <div class="stats-cell">{{ item.correct }} / {{ item.total }}</div>
        </div>
        <div v-if="filteredWords.length === 0" class="empty">
          Нет слов для отображения. Пройдите тесты в словаре.
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
  .progress-root {
    min-height: 100vh;
    background: #fffdf7;
    padding: 2rem 1rem;
  }
  .container {
    max-width: 1000px;
    margin: 0 auto;
  }
  h1 {
    color: #2c3e50;
    margin-bottom: 2rem;
    text-align: center;
  }
  .controls {
    display: flex;
    gap: 1.5rem;
    justify-content: center;
    margin-bottom: 2rem;
    background: white;
    padding: 1rem;
    border-radius: 28px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  }
  .filter-group,
  .sort-group {
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }
  select {
    padding: 0.5rem 1rem;
    border-radius: 20px;
    border: 2px solid #ddd;
    background: white;
    font-size: 0.9rem;
    outline: none;
    cursor: pointer;
  }
  .stats-summary {
    display: flex;
    gap: 1.5rem;
    justify-content: center;
    margin-bottom: 2rem;
    flex-wrap: wrap;
  }
  .stat-card {
    background: white;
    border-radius: 24px;
    padding: 1rem 2rem;
    text-align: center;
    box-shadow: 0 4px 12px rgba(88, 204, 2, 0.1);
    min-width: 150px;
  }
  .stat-label {
    display: block;
    font-size: 0.9rem;
    color: #666;
    margin-bottom: 0.5rem;
  }
  .stat-value {
    font-size: 1.8rem;
    font-weight: 700;
    color: #2c3e50;
  }
  .words-list {
    background: white;
    border-radius: 28px;
    overflow: hidden;
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.05);
  }
  .list-header,
  .list-row {
    display: grid;
    grid-template-columns: 1fr 2fr 1fr;
    padding: 1rem 1.5rem;
    align-items: center;
  }
  .list-header {
    background: #f7f9fa;
    font-weight: 700;
    color: #2c3e50;
    border-bottom: 1px solid #e0e0e0;
  }
  .list-row {
    border-bottom: 1px solid #f0f0f0;
  }
  .list-row:last-child {
    border-bottom: none;
  }
  .word-cell {
    font-weight: 500;
    color: #2c3e50;
  }
  .progress-cell {
    display: flex;
    align-items: center;
    gap: 1rem;
  }
  .progress-bar-bg {
    flex: 1;
    height: 10px;
    background: #e0e0e0;
    border-radius: 10px;
    overflow: hidden;
  }
  .progress-bar-fill {
    height: 100%;
    border-radius: 10px;
    transition: width 0.2s;
  }
  .progress-bar-fill.progress-high {
    background: #58cc02;
  }
  .progress-bar-fill.progress-mid {
    background: #ffb300;
  }
  .progress-bar-fill.progress-low {
    background: #e74c3c;
  }
  .percent-text {
    min-width: 45px;
    font-weight: 600;
  }
  .stats-cell {
    color: #555;
  }
  .empty {
    text-align: center;
    padding: 2rem;
    color: #888;
  }
</style>
