<script setup>
  import { ref, reactive, computed, onMounted, watch } from 'vue'
  import defaultData from '@/data/dictionary.json'

  const LOCAL_KEY = 'user_dictionary'

  const dictionary = reactive({ folders: [] })

  onMounted(() => {
    const saved = localStorage.getItem(LOCAL_KEY)
    if (saved) {
      try {
        const parsed = JSON.parse(saved)
        dictionary.folders = parsed.folders || []
      } catch {
        dictionary.folders = defaultData.folders
      }
    } else {
      dictionary.folders = defaultData.folders
    }
  })

  watch(
    dictionary,
    () => localStorage.setItem(LOCAL_KEY, JSON.stringify({ folders: dictionary.folders })),
    { deep: true },
  )

  // Папки
  const selectedFolderId = ref(null)
  const currentFolder = computed(() =>
    dictionary.folders.find((f) => f.id === selectedFolderId.value),
  )

  function selectFolder(id) {
    selectedFolderId.value = id
    displayedCount.value = 15
    showTest.value = false
    showResultModal.value = false
    isSelectMode.value = false
    selectedIndices.value = new Set()
    searchQuery.value = ''
    closeContextMenu()
  }

  function addFolder() {
    const title = prompt('Название папки:')
    if (title && title.trim()) {
      dictionary.folders.push({ id: Date.now().toString(), title: title.trim(), words: [] })
    }
  }

  // Поиск
  const searchQuery = ref('')
  const filteredWords = computed(() => {
    if (!currentFolder.value) return []
    const q = searchQuery.value.trim().toLowerCase()
    if (!q) return currentFolder.value.words
    return currentFolder.value.words.filter(
      (w) =>
        w.word.toLowerCase().includes(q) ||
        w.translation.toLowerCase().includes(q) ||
        w.transcription.toLowerCase().includes(q),
    )
  })

  // Пагинация
  const displayedCount = ref(15)
  function loadMore() {
    displayedCount.value += 15
  }
  watch(selectedFolderId, () => {
    displayedCount.value = 15
  })

  // Тестирование
  const showTest = ref(false)
  const testMode = ref('')
  const testWords = ref([])
  const testIndex = ref(0)
  const testScore = ref(0)
  const testAnswer = ref(null)
  const testInput = ref('')
  const testOptions = ref([])

  // Источник слов для теста
  let testWordsSource = []

  // Результаты теста
  const showResultModal = ref(false)
  const testResults = ref([]) // { word, translation, userAnswer, isCorrect, correctAnswer }

  // Модальное окно выбора количества
  const showCountModal = ref(false)
  const countModalMode = ref('')
  const countModalInput = ref('')
  const countModalError = ref('')
  const testWordCount = ref(0)
  const currentMaxWords = ref(0)

  function openTestSetup(mode, customWords = null) {
    const wordsArray = customWords || currentFolder.value?.words || []
    if (wordsArray.length === 0) {
      alert('Нет слов для тестирования')
      return
    }
    testWordsSource = wordsArray
    testMode.value = mode
    countModalMode.value = mode
    countModalInput.value = ''
    countModalError.value = ''
    currentMaxWords.value = wordsArray.length
    showCountModal.value = true
  }

  function confirmCount() {
    const max = currentMaxWords.value
    const val = countModalInput.value.trim()
    if (val === '') {
      countModalError.value = 'Введите число'
      return
    }
    const num = parseInt(val, 10)
    if (isNaN(num) || num < 1 || num > max) {
      countModalError.value = `Введите число от 1 до ${max}`
      return
    }
    testWordCount.value = num
    showCountModal.value = false
    testResults.value = []
    startTest(testWordsSource, num)
  }

  function startTest(wordsArray, count) {
    const shuffled = [...wordsArray].sort(() => Math.random() - 0.5)
    testWords.value = shuffled.slice(0, count)
    testIndex.value = 0
    testScore.value = 0
    testAnswer.value = null
    testInput.value = ''
    showTest.value = true
    prepareCurrentQuestion()
  }

  function currentTestWord() {
    return testWords.value[testIndex.value] || null
  }

  function prepareCurrentQuestion() {
    if (testMode.value === 'choose') {
      const correct = currentTestWord()
      if (!correct) return
      const others = currentFolder.value.words
        .filter((w) => w.word !== correct.word)
        .sort(() => Math.random() - 0.5)
        .slice(0, 3)
        .map((w) => w.translation)
      testOptions.value = [correct.translation, ...others].sort(() => Math.random() - 0.5)
      testAnswer.value = null
    } else {
      testInput.value = ''
    }
  }

  function submitChoose() {
    const word = currentTestWord()
    if (!word) return
    const selectedTranslation = testOptions.value[testAnswer.value]
    const isCorrect = selectedTranslation === word.translation
    if (isCorrect) testScore.value++
    testResults.value.push({
      word: word.word,
      translation: word.translation,
      userAnswer: selectedTranslation,
      isCorrect,
      correctAnswer: word.translation,
    })
    nextQuestion()
  }

  function submitWrite() {
    const word = currentTestWord()
    if (!word) return
    const userWritten = testInput.value.trim()
    const isCorrect = userWritten.toLowerCase() === word.word.toLowerCase()
    if (isCorrect) testScore.value++
    testResults.value.push({
      word: word.word,
      translation: word.translation,
      userAnswer: userWritten || '(пусто)',
      isCorrect,
      correctAnswer: word.word,
    })
    nextQuestion()
  }

  function nextQuestion() {
    if (testIndex.value < testWords.value.length - 1) {
      testIndex.value++
      prepareCurrentQuestion()
    } else {
      showResultModal.value = true
    }
  }

  function closeResultModal() {
    showResultModal.value = false
    showTest.value = false
    testResults.value = []
  }

  function cancelTest() {
    showTest.value = false
    showResultModal.value = false
  }

  function testSelectedWords(mode) {
    if (!currentFolder.value || selectedIndices.value.size === 0) return
    const indices = [...selectedIndices.value]
    const words = indices.map((i) => currentFolder.value.words[i]).filter(Boolean)
    if (words.length === 0) return
    testMode.value = mode
    testResults.value = []
    startTest(words, words.length)
  }

  // Добавление / удаление слова
  function addWordToFolder() {
    if (!currentFolder.value) return
    const word = prompt('Слово (на иностранном языке):')
    if (!word) return
    const transcription = prompt('Транскрипция:') || ''
    const translation = prompt('Перевод:')
    if (!translation) return
    currentFolder.value.words.push({
      word: word.trim(),
      transcription: transcription.trim(),
      translation: translation.trim(),
    })
  }

  function removeWord(index) {
    currentFolder.value.words.splice(index, 1)
  }

  // Множественный выбор
  const isSelectMode = ref(false)
  const selectedIndices = ref(new Set())

  function toggleSelectMode() {
    isSelectMode.value = !isSelectMode.value
    if (!isSelectMode.value) selectedIndices.value = new Set()
  }

  function toggleWordSelection(globalIndex) {
    if (!isSelectMode.value) return
    const set = new Set(selectedIndices.value)
    if (set.has(globalIndex)) set.delete(globalIndex)
    else set.add(globalIndex)
    selectedIndices.value = set
  }

  function clearSelection() {
    selectedIndices.value = new Set()
  }

  // Копирование / перемещение
  const showFolderPicker = ref(false)
  const folderPickerAction = ref('')
  const folderPickerIndices = ref([])

  function openFolderPicker(action, indices) {
    folderPickerAction.value = action
    folderPickerIndices.value = indices
    showFolderPicker.value = true
  }

  function executeFolderAction(targetFolderId) {
    if (!currentFolder.value) return
    const targetFolder = dictionary.folders.find((f) => f.id === targetFolderId)
    if (!targetFolder || targetFolder.id === currentFolder.value.id) return
    const wordsToProcess = folderPickerIndices.value
      .map((i) => currentFolder.value.words[i])
      .filter(Boolean)
    if (folderPickerAction.value === 'copy') {
      targetFolder.words.push(...wordsToProcess.map((w) => ({ ...w })))
    } else if (folderPickerAction.value === 'move') {
      targetFolder.words.push(...wordsToProcess.map((w) => ({ ...w })))
      const sortedIndices = [...folderPickerIndices.value].sort((a, b) => b - a)
      for (const idx of sortedIndices) {
        currentFolder.value.words.splice(idx, 1)
      }
    }
    showFolderPicker.value = false
    clearSelection()
    isSelectMode.value = false
  }

  // Контекстное меню
  const contextMenu = ref({ visible: false, x: 0, y: 0, wordIndex: null })

  function openContextMenu(event, wordIndex) {
    event.stopPropagation()
    contextMenu.value = { visible: true, x: event.clientX, y: event.clientY, wordIndex }
  }

  function closeContextMenu() {
    contextMenu.value.visible = false
  }

  function contextAction(action) {
    if (!currentFolder.value || contextMenu.value.wordIndex === null) return
    const idx = contextMenu.value.wordIndex
    if (action === 'delete') removeWord(idx)
    else if (action === 'copy' || action === 'move') openFolderPicker(action, [idx])
    closeContextMenu()
  }

  function onWindowClick() {
    if (contextMenu.value.visible) closeContextMenu()
  }
  onMounted(() => window.addEventListener('click', onWindowClick))
</script>

<template>
  <div class="dict-root">
    <aside class="sidebar">
      <h2 class="sidebar-title">📁 Папки</h2>
      <ul class="folder-list">
        <li
          v-for="folder in dictionary.folders"
          :key="folder.id"
          class="folder-item"
          :class="{ active: folder.id === selectedFolderId }"
          @click="selectFolder(folder.id)">
          {{ folder.title }}
          <span class="word-count">{{ folder.words.length }}</span>
        </li>
      </ul>
      <button class="add-folder-btn" @click="addFolder">+ Добавить папку</button>
    </aside>

    <main class="main-area">
      <template v-if="!currentFolder">
        <div class="empty-state">Выберите папку слева</div>
      </template>

      <template v-else>
        <div class="folder-header">
          <h1>{{ currentFolder.title }}</h1>
          <button
            v-if="currentFolder.id !== 'completed'"
            class="add-word-btn"
            @click="addWordToFolder">
            + Добавить слово
          </button>
        </div>

        <div v-if="!isSelectMode || selectedIndices.size === 0" class="test-buttons">
          <button :disabled="currentFolder.words.length === 0" @click="openTestSetup('choose')">
            🧠 Выбрать перевод
          </button>
          <button :disabled="currentFolder.words.length === 0" @click="openTestSetup('write')">
            ✍️ Написать слово
          </button>
        </div>

        <div class="tools-bar">
          <input v-model="searchQuery" class="search-input" placeholder="Поиск по словам..." />
          <button class="select-mode-btn" @click="toggleSelectMode">
            {{ isSelectMode ? 'Отменить выбор' : 'Выбрать слова' }}
          </button>
        </div>

        <div v-if="isSelectMode && selectedIndices.size > 0" class="selection-actions">
          <span class="selection-count">Выбрано: {{ selectedIndices.size }}</span>
          <button class="action-btn" @click="openFolderPicker('copy', [...selectedIndices])">
            Копировать в...
          </button>
          <button
            v-if="currentFolder.id !== 'completed'"
            class="action-btn"
            @click="openFolderPicker('move', [...selectedIndices])">
            Переместить в...
          </button>
          <button
            v-if="currentFolder.id !== 'completed'"
            class="action-btn danger"
            @click="
              () => {
                const sorted = [...selectedIndices].sort((a, b) => b - a)
                for (const idx of sorted) removeWord(idx)
                clearSelection()
                isSelectMode = false
              }
            ">
            Удалить выбранные
          </button>
          <div class="selection-test-actions">
            <button class="action-btn test-action" @click="testSelectedWords('choose')">
              🧠 Выбрать перевод
            </button>
            <button class="action-btn test-action" @click="testSelectedWords('write')">
              ✍️ Написать слово
            </button>
          </div>
        </div>

        <div v-if="!showTest" class="words-section">
          <div v-if="filteredWords.length === 0" class="empty-folder">
            {{ searchQuery ? 'Ничего не найдено' : 'Папка пуста. Добавьте слова.' }}
          </div>
          <div
            v-for="(word, displayIndex) in filteredWords.slice(0, displayedCount)"
            :key="word.word + displayIndex"
            class="word-row"
            :class="{
              selectable: isSelectMode,
              selected: isSelectMode && selectedIndices.has(currentFolder.words.indexOf(word)),
            }"
            @click="isSelectMode && toggleWordSelection(currentFolder.words.indexOf(word))">
            <div class="word-main">
              <span class="word-text">{{ word.word }}</span>
              <span class="word-transcription">{{ word.transcription }}</span>
              <span class="word-translation">{{ word.translation }}</span>
            </div>
            <div class="word-actions">
              <button
                class="context-menu-btn"
                @click.stop="openContextMenu($event, currentFolder.words.indexOf(word))">
                ⋮
              </button>
            </div>
          </div>
          <div v-if="displayedCount < filteredWords.length" class="load-more">
            <button @click="loadMore">
              Показать ещё ({{ filteredWords.length - displayedCount }})
            </button>
          </div>
        </div>

        <div v-else class="test-container">
          <div class="test-header">
            <span>{{ testMode === 'choose' ? 'Выбор перевода' : 'Напишите слово' }}</span>
            <span>{{ testIndex + 1 }} / {{ testWords.length }}</span>
            <button class="cancel-test-btn" @click="cancelTest">✕</button>
          </div>
          <div class="test-question">
            <div v-if="testMode === 'choose'">
              <p class="test-word">{{ currentTestWord()?.word }}</p>
              <div class="options">
                <button
                  v-for="(opt, i) in testOptions"
                  :key="i"
                  :class="{ selected: testAnswer === i }"
                  @click="testAnswer = i">
                  {{ opt }}
                </button>
              </div>
              <button class="submit-btn" :disabled="testAnswer === null" @click="submitChoose">
                Ответить
              </button>
            </div>
            <div v-else>
              <p class="test-word">{{ currentTestWord()?.translation }}</p>
              <input
                v-model="testInput"
                class="test-input"
                placeholder="Введите слово"
                @keyup.enter="submitWrite" />
              <button class="submit-btn" @click="submitWrite">Ответить</button>
            </div>
          </div>
        </div>
      </template>
    </main>

    <!-- Контекстное меню -->
    <div
      v-if="contextMenu.visible"
      class="context-menu"
      :style="{ left: contextMenu.x + 'px', top: contextMenu.y + 'px' }"
      @click.stop>
      <button @click="contextAction('copy')">Копировать в...</button>
      <button v-if="currentFolder?.id !== 'completed'" @click="contextAction('move')">
        Переместить в...
      </button>
      <button
        v-if="currentFolder?.id !== 'completed'"
        class="danger"
        @click="contextAction('delete')">
        Удалить
      </button>
    </div>

    <!-- Модалка выбора папки для копирования/перемещения -->
    <div v-if="showFolderPicker" class="modal-overlay" @click.self="showFolderPicker = false">
      <div class="modal">
        <h3>{{ folderPickerAction === 'copy' ? 'Копировать в...' : 'Переместить в...' }}</h3>
        <ul class="folder-pick-list">
          <li
            v-for="folder in dictionary.folders.filter((f) => f.id !== currentFolder?.id)"
            :key="folder.id"
            @click="executeFolderAction(folder.id)">
            {{ folder.title }}
          </li>
        </ul>
        <button @click="showFolderPicker = false">Отмена</button>
      </div>
    </div>

    <!-- Модальное окно выбора количества -->
    <div v-if="showCountModal" class="modal-overlay" @click.self="showCountModal = false">
      <div class="modal count-modal">
        <h3>Сколько слов тестировать?</h3>
        <p class="count-hint">
          Доступно слов:
          <strong>{{ currentMaxWords }}</strong>
        </p>
        <div class="count-input-row">
          <input
            v-model="countModalInput"
            class="count-input"
            :placeholder="currentMaxWords"
            @keyup.enter="confirmCount" />
        </div>
        <p v-if="countModalError" class="count-error">{{ countModalError }}</p>
        <div class="count-actions">
          <button class="cancel-btn" @click="showCountModal = false">Отмена</button>
          <button class="confirm-btn" @click="confirmCount">Начать тест</button>
        </div>
      </div>
    </div>

    <!-- Модальное окно результатов теста -->
    <div v-if="showResultModal" class="modal-overlay" @click.self="closeResultModal">
      <div class="modal result-modal">
        <h3>Результаты теста</h3>
        <p class="result-score">
          Правильно:
          <strong>{{ testScore }}</strong>
          из {{ testResults.length }}
        </p>
        <div class="result-list">
          <div
            v-for="(item, idx) in testResults"
            :key="idx"
            class="result-item"
            :class="{ correct: item.isCorrect, wrong: !item.isCorrect }">
            <div class="result-word-row">
              <span class="result-word">{{ item.word }}</span>
              <span class="result-translation">({{ item.translation }})</span>
            </div>
            <div class="result-answer">
              Ваш ответ:
              <strong>{{ item.userAnswer }}</strong>
            </div>
            <div v-if="!item.isCorrect" class="result-correct-answer">
              Правильно:
              <strong>{{ item.correctAnswer }}</strong>
            </div>
          </div>
        </div>
        <button class="confirm-btn" @click="closeResultModal">Закрыть</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
  /* Все предыдущие стили плюс новые для результата */
  .dict-root {
    display: flex;
    min-height: 100vh;
    background: #fffdf7;
    font-family: 'Segoe UI', sans-serif;
  }

  .sidebar {
    width: 260px;
    background: white;
    border-right: 1px solid rgba(88, 204, 2, 0.3);
    padding: 1.5rem;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    box-shadow: 2px 0 10px rgba(0, 0, 0, 0.02);
  }

  .sidebar-title {
    color: #2c3e50;
    font-size: 1.3rem;
    margin: 0;
  }

  .folder-list {
    list-style: none;
    padding: 0;
    margin: 0;
    display: flex;
    flex-direction: column;
    gap: 0.4rem;
  }

  .folder-item {
    padding: 0.7rem 1rem;
    border-radius: 12px;
    cursor: pointer;
    display: flex;
    justify-content: space-between;
    align-items: center;
    color: #2c3e50;
    font-weight: 500;
    transition: background 0.2s;
  }

  .folder-item:hover {
    background: rgba(88, 204, 2, 0.06);
  }

  .folder-item.active {
    background: rgba(88, 204, 2, 0.12);
    color: #58cc02;
    font-weight: 600;
  }

  .word-count {
    background: #eee;
    border-radius: 20px;
    padding: 0.1rem 0.6rem;
    font-size: 0.8rem;
    color: #666;
  }

  .add-folder-btn {
    background: none;
    border: 2px dashed #ccc;
    border-radius: 12px;
    padding: 0.6rem;
    color: #888;
    cursor: pointer;
    transition: all 0.2s;
  }

  .add-folder-btn:hover {
    border-color: #58cc02;
    color: #58cc02;
    background: rgba(88, 204, 2, 0.04);
  }

  .main-area {
    flex: 1;
    padding: 2rem;
  }

  .empty-state {
    text-align: center;
    color: #aaa;
    font-size: 1.5rem;
    margin-top: 6rem;
  }

  .folder-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1.5rem;
  }

  .folder-header h1 {
    margin: 0;
    color: #2c3e50;
  }

  .add-word-btn {
    background: #58cc02;
    color: white;
    border: none;
    border-radius: 20px;
    padding: 0.5rem 1.2rem;
    font-weight: 600;
    cursor: pointer;
    transition: background 0.2s;
  }

  .add-word-btn:hover {
    background: #46a302;
  }

  .test-buttons {
    display: flex;
    gap: 1rem;
    margin-bottom: 2rem;
  }

  .test-buttons button {
    background: white;
    border: 2px solid #58cc02;
    border-radius: 20px;
    padding: 0.6rem 1.4rem;
    font-size: 1rem;
    cursor: pointer;
    transition: all 0.2s;
    display: flex;
    align-items: center;
    gap: 0.3rem;
  }

  .test-buttons button:hover:not(:disabled) {
    background: #58cc02;
    color: white;
  }

  .test-buttons button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .tools-bar {
    display: flex;
    gap: 1rem;
    margin-bottom: 1.5rem;
    align-items: center;
  }

  .search-input {
    flex: 1;
    padding: 0.6rem 1rem;
    border: 2px solid #ddd;
    border-radius: 12px;
    font-size: 1rem;
    outline: none;
  }

  .search-input:focus {
    border-color: #58cc02;
  }

  .select-mode-btn {
    background: white;
    border: 2px solid #58cc02;
    border-radius: 20px;
    padding: 0.5rem 1.2rem;
    cursor: pointer;
    font-weight: 600;
    color: #2c3e50;
    transition: all 0.2s;
    white-space: nowrap;
  }

  .select-mode-btn:hover {
    background: #58cc02;
    color: white;
  }

  .words-section {
    background: white;
    border-radius: 24px;
    padding: 1rem;
    box-shadow: 0 10px 25px rgba(88, 204, 2, 0.06);
  }

  .word-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.8rem 1rem;
    border-bottom: 1px solid #f0f0f0;
    transition: background 0.2s;
  }

  .word-row:last-child {
    border-bottom: none;
  }

  .word-main {
    display: flex;
    gap: 1.5rem;
    align-items: baseline;
    flex-wrap: wrap;
  }

  .word-text {
    font-weight: 700;
    font-size: 1.1rem;
    color: #2c3e50;
    min-width: 100px;
  }

  .word-transcription {
    color: #888;
    font-size: 0.95rem;
    font-style: italic;
  }

  .word-translation {
    color: #555;
    font-size: 0.95rem;
  }

  .word-row.selectable {
    cursor: pointer;
  }

  .word-row.selected {
    background: #f0fff0;
    border-left: 4px solid #58cc02;
    padding-left: calc(1rem - 4px);
  }

  .word-actions {
    position: relative;
  }

  .context-menu-btn {
    background: none;
    border: none;
    font-size: 1.4rem;
    cursor: pointer;
    padding: 0 0.3rem;
    color: #888;
    border-radius: 50%;
    transition: background 0.2s;
    line-height: 1;
  }

  .context-menu-btn:hover {
    background: #eee;
    color: #333;
  }

  .context-menu {
    position: fixed;
    background: white;
    border: 1px solid #ddd;
    border-radius: 12px;
    box-shadow: 0 6px 20px rgba(0, 0, 0, 0.12);
    display: flex;
    flex-direction: column;
    z-index: 1000;
    min-width: 180px;
    padding: 0.4rem 0;
  }

  .context-menu button {
    background: none;
    border: none;
    padding: 0.7rem 1.2rem;
    text-align: left;
    font-size: 0.95rem;
    cursor: pointer;
    transition: background 0.2s;
    color: #2c3e50;
  }

  .context-menu button:hover {
    background: #f5f5f5;
  }

  .context-menu button.danger {
    color: #e74c3c;
  }

  .context-menu button.danger:hover {
    background: #ffeaea;
  }

  .selection-actions {
    display: flex;
    gap: 0.8rem;
    align-items: center;
    background: white;
    border-radius: 20px;
    padding: 0.8rem 1.2rem;
    margin-bottom: 1rem;
    box-shadow: 0 4px 12px rgba(88, 204, 2, 0.08);
    flex-wrap: wrap;
  }

  .selection-count {
    font-weight: 600;
    margin-right: auto;
    color: #2c3e50;
  }

  .action-btn {
    background: white;
    border: 2px solid #58cc02;
    border-radius: 16px;
    padding: 0.4rem 1rem;
    font-size: 0.9rem;
    cursor: pointer;
    transition: 0.2s;
    color: #2c3e50;
    font-weight: 500;
  }

  .action-btn:hover {
    background: #58cc02;
    color: white;
  }

  .action-btn.danger {
    border-color: #e74c3c;
    color: #e74c3c;
  }

  .action-btn.danger:hover {
    background: #e74c3c;
    color: white;
  }

  .selection-test-actions {
    display: flex;
    gap: 0.5rem;
    margin-top: 0.5rem;
    width: 100%;
    justify-content: flex-end;
  }

  .test-action {
    background: white;
    border: 2px solid #58cc02;
    color: #2c3e50;
    font-weight: 500;
    transition: 0.2s;
  }

  .test-action:hover {
    background: #58cc02;
    color: white;
  }

  .load-more {
    text-align: center;
    margin-top: 1rem;
  }

  .load-more button {
    background: none;
    border: 1px solid #ddd;
    border-radius: 20px;
    padding: 0.5rem 1.5rem;
    color: #666;
    cursor: pointer;
  }

  .load-more button:hover {
    border-color: #58cc02;
    color: #58cc02;
  }

  .empty-folder {
    text-align: center;
    color: #888;
    padding: 2rem;
  }

  .test-container {
    background: white;
    border-radius: 24px;
    padding: 1.5rem;
    box-shadow: 0 10px 25px rgba(88, 204, 2, 0.06);
  }

  .test-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
    font-weight: 600;
    color: #2c3e50;
  }

  .cancel-test-btn {
    background: none;
    border: none;
    font-size: 1.5rem;
    cursor: pointer;
    color: #aaa;
  }

  .test-question {
    text-align: center;
  }

  .test-word {
    font-size: 2rem;
    font-weight: 700;
    margin-bottom: 1.5rem;
    color: #2c3e50;
  }

  .options {
    display: flex;
    flex-direction: column;
    gap: 0.8rem;
    align-items: center;
    margin-bottom: 1.5rem;
  }

  .options button {
    width: 300px;
    padding: 0.7rem;
    border: 2px solid #ccc;
    border-radius: 16px;
    background: white;
    cursor: pointer;
    transition: all 0.2s;
  }

  .options button.selected {
    border-color: #58cc02;
    background: #f0fff0;
    font-weight: 600;
  }

  .test-input {
    width: 300px;
    padding: 0.7rem;
    border: 2px solid #ccc;
    border-radius: 16px;
    font-size: 1.2rem;
    text-align: center;
    margin-bottom: 1rem;
  }

  .submit-btn {
    background: #58cc02;
    color: white;
    border: none;
    border-radius: 20px;
    padding: 0.7rem 2rem;
    font-size: 1.1rem;
    cursor: pointer;
  }

  .submit-btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .submit-btn:hover:not(:disabled) {
    background: #46a302;
  }

  .modal-overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.3);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 2000;
  }

  .modal {
    background: white;
    border-radius: 24px;
    padding: 2rem;
    min-width: 260px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  }

  .modal h3 {
    margin: 0 0 1.2rem;
    color: #2c3e50;
  }

  .folder-pick-list {
    list-style: none;
    padding: 0;
    margin: 0 0 1.5rem;
    display: flex;
    flex-direction: column;
    gap: 0.6rem;
  }

  .folder-pick-list li {
    padding: 0.6rem 1rem;
    border-radius: 12px;
    cursor: pointer;
    background: #f9f9f9;
    transition: background 0.2s;
  }

  .folder-pick-list li:hover {
    background: rgba(88, 204, 2, 0.08);
    color: #58cc02;
  }

  .modal > button {
    background: none;
    border: 1px solid #ccc;
    border-radius: 12px;
    padding: 0.5rem 1.5rem;
    cursor: pointer;
    width: 100%;
    font-size: 0.95rem;
    transition: background 0.2s;
  }

  .modal > button:hover {
    background: #f5f5f5;
  }

  .count-modal {
    text-align: center;
  }

  .count-hint {
    color: #666;
    margin-bottom: 1.5rem;
  }

  .count-input-row {
    display: flex;
    gap: 0.5rem;
    justify-content: center;
    margin-bottom: 0.8rem;
  }

  .count-input {
    width: 120px;
    padding: 0.6rem;
    border: 2px solid #ddd;
    border-radius: 12px;
    font-size: 1.1rem;
    text-align: center;
    outline: none;
  }

  .count-input:focus {
    border-color: #58cc02;
  }

  .count-error {
    color: #e74c3c;
    font-size: 0.9rem;
    margin-bottom: 1rem;
  }

  .count-actions {
    display: flex;
    gap: 1rem;
    justify-content: center;
  }

  .cancel-btn {
    background: none;
    border: 1px solid #ccc;
    border-radius: 12px;
    padding: 0.5rem 1.5rem;
    cursor: pointer;
  }

  .confirm-btn {
    background: #58cc02;
    color: white;
    border: none;
    border-radius: 12px;
    padding: 0.5rem 1.8rem;
    font-weight: 600;
    cursor: pointer;
  }

  .confirm-btn:hover {
    background: #46a302;
  }

  /* Стили окна результатов */
  .result-modal {
    max-height: 80vh;
    overflow-y: auto;
    min-width: 350px;
  }

  .result-score {
    text-align: center;
    font-size: 1.2rem;
    margin-bottom: 1.5rem;
    color: #2c3e50;
  }

  .result-list {
    display: flex;
    flex-direction: column;
    gap: 0.8rem;
    margin-bottom: 1.5rem;
  }

  .result-item {
    padding: 0.8rem 1rem;
    border-radius: 12px;
    background: #f9f9f9;
  }

  .result-item.correct {
    border-left: 4px solid #58cc02;
    background: #f0fff0;
  }

  .result-item.wrong {
    border-left: 4px solid #e74c3c;
    background: #fff5f5;
  }

  .result-word-row {
    display: flex;
    align-items: baseline;
    gap: 0.5rem;
  }

  .result-word {
    font-weight: 700;
    font-size: 1.1rem;
    color: #2c3e50;
  }

  .result-translation {
    color: #666;
  }

  .result-answer {
    margin-top: 0.4rem;
    font-size: 0.95rem;
  }

  .result-correct-answer {
    color: #e74c3c;
    font-size: 0.95rem;
    font-weight: 500;
  }
</style>
