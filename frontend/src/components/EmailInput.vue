<script setup>
  import { ref, watch } from 'vue'

  const props = defineProps({
    label: { type: String, default: '' },
    type: { type: String, default: 'text' },
    placeholder: { type: String, default: '' },
    modelValue: { type: [String, Number], default: '' },
  })

  const emit = defineEmits(['update:modelValue'])

  const errorMessage = ref('')

  const validateEmail = (value) => {
    // Проверяем только если тип email и значение не пустое
    if (props.type !== 'email') {
      errorMessage.value = ''
      return
    }

    if (!value) {
      errorMessage.value = '' // пустое – нет ошибки формата (required добавит родитель)
      return
    }

    const trimmed = value.trim()
    if (trimmed.length < 5) {
      errorMessage.value = 'Email слишком короткий (минимум 5 символов)'
      return
    }
    if (trimmed.length > 254) {
      errorMessage.value = 'Email слишком длинный (максимум 254 символа)'
      return
    }

    // Основная регулярка: локальная часть и домен
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    if (!emailRegex.test(trimmed)) {
      errorMessage.value = 'Некорректный формат email'
      return
    }

    // Дополнительно можно проверить наличие точки после @ и т.д. – базовая регулярка уже это делает
    errorMessage.value = ''
  }

  // Валидация при потере фокуса
  const onBlur = (event) => {
    validateEmail(event.target.value)
  }

  // При изменении модели извне (например, v-model) очищаем ошибку
  watch(
    () => props.modelValue,
    (newVal) => {
      if (errorMessage.value) {
        validateEmail(newVal)
      }
    },
  )
</script>

<template>
  <div class="form-group">
    <label v-if="label" :for="$attrs.id || label">{{ label }}</label>
    <input
      :type="type"
      :value="modelValue"
      :placeholder="placeholder"
      class="form-input"
      :class="{ 'input--error': errorMessage }"
      v-bind="$attrs"
      @input="emit('update:modelValue', $event.target.value)"
      @blur="onBlur" />
    <p v-if="errorMessage" class="error-text">{{ errorMessage }}</p>
  </div>
</template>

<style scoped>
  /* ... существующие стили ... */
  /* добавляем стиль ошибки */
  .form-group {
    text-align: left;
    width: 100%;
  }

  .form-group label {
    display: block;
    margin-bottom: 0.4rem;
    font-size: 0.95rem;
    color: #444;
    font-weight: 500;
  }

  .form-input {
    width: 100%;
    padding: 0.8rem 1rem;
    font-size: 1rem;
    border: 1px solid rgba(88, 204, 2, 0.4);
    border-radius: 20px;
    background: #fdfdf5;
    outline: none;
    transition:
      border-color 0.2s,
      box-shadow 0.2s;
    box-sizing: border-box;
    font-family: inherit;
    color: #333;
  }

  .form-input:focus {
    border-color: rgba(88, 204, 2, 0.8);
    box-shadow: 0 0 0 3px rgba(88, 204, 2, 0.1);
  }

  .form-input::placeholder {
    color: #aaa;
  }

  /* Новый класс для ошибочного состояния */
  .input--error {
    border-color: #e74c3c !important;
    box-shadow: 0 0 0 3px rgba(231, 76, 60, 0.1) !important;
  }

  .error-text {
    color: #e74c3c;
    font-size: 0.85rem;
    margin-top: 0.3rem;
    text-align: left;
  }
</style>
