<script setup>
  import { ref, watch } from 'vue'

  const props = defineProps({
    label: { type: String, default: '' },
    placeholder: { type: String, default: '' },
    modelValue: { type: String, default: '' },
  })

  const emit = defineEmits(['update:modelValue'])

  const showPassword = ref(false)
  const errorMessage = ref('')

  const togglePassword = (event) => {
    showPassword.value = !showPassword.value
    const input = event.currentTarget.parentElement.querySelector('input')
    if (input) input.focus()
  }

  const validatePassword = (value) => {
    if (!value) {
      errorMessage.value = ''
      return
    }

    if (/[^\x20-\x7E]/.test(value)) {
      errorMessage.value =
        'Пароль должен содержать только латинские буквы, цифры и специальные символы'
      return
    }
    // Разрешены любые символы кроме пробелов
    if (/\s/.test(value)) {
      errorMessage.value = 'Пароль не должен содержать пробелов'
      return
    }
    if (value.length < 6) {
      errorMessage.value = 'Пароль должен содержать минимум 6 символов'
      return
    }
    if (value.length > 128) {
      errorMessage.value = 'Пароль слишком длинный (максимум 128 символов)'
      return
    }
    // Минимум одна буква и одна цифра (опционально, можно закомментировать)
    if (!/[a-zA-Z]/.test(value) || !/[0-9]/.test(value)) {
      errorMessage.value = 'Пароль должен содержать хотя бы одну букву и одну цифру'
      return
    }
    //Хотя бы один спец символ
    if (!/[^a-zA-Z0-9]/.test(value)) {
      errorMessage.value = 'Пароль должен содержать хотя бы один специальный символ'
      return
    }

    errorMessage.value = ''
  }

  const onBlur = (event) => {
    validatePassword(event.target.value)
  }

  // Автоматическая перепроверка при изменении значения, если уже была ошибка
  watch(
    () => props.modelValue,
    (newVal) => {
      if (errorMessage.value) {
        validatePassword(newVal)
      }
    },
  )
</script>

<template>
  <div class="form-group">
    <label v-if="label" :for="$attrs.id || label">{{ label }}</label>
    <div class="input-wrapper">
      <input
        :type="showPassword ? 'text' : 'password'"
        :value="modelValue"
        :placeholder="placeholder"
        class="form-input"
        :class="{ 'input--error': errorMessage }"
        v-bind="$attrs"
        @input="emit('update:modelValue', $event.target.value)"
        @blur="onBlur" />
      <button type="button" class="eye-button" @click="togglePassword" tabindex="-1">
        <!-- иконка открытого глаза -->
        <svg
          v-if="showPassword"
          xmlns="http://www.w3.org/2000/svg"
          width="22"
          height="22"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round">
          <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z" />
          <circle cx="12" cy="12" r="3" />
        </svg>
        <!-- иконка закрытого глаза (перечёркнутый) -->
        <svg
          v-else
          xmlns="http://www.w3.org/2000/svg"
          width="22"
          height="22"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round">
          <path
            d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94" />
          <path d="M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19" />
          <path d="M1 1l22 22" />
          <path d="M14.12 14.12a3 3 0 1 1-4.24-4.24" />
        </svg>
      </button>
    </div>
    <p v-if="errorMessage" class="error-text">{{ errorMessage }}</p>
  </div>
</template>

<style scoped>
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

  .input-wrapper {
    position: relative;
    display: flex;
    align-items: center;
  }

  .form-input {
    width: 100%;
    padding: 0.8rem 3rem 0.8rem 1rem; /* правый отступ под кнопку */
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

  .input--error {
    border-color: #e74c3c !important;
    box-shadow: 0 0 0 3px rgba(231, 76, 60, 0.1) !important;
  }

  .eye-button {
    position: absolute;
    right: 12px;
    top: 50%;
    transform: translateY(-50%);
    background: transparent;
    border: none;
    cursor: pointer;
    color: #777;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2px;
    border-radius: 4px;
    transition: color 0.2s;
  }

  .eye-button:hover {
    color: #333;
  }

  .error-text {
    color: #e74c3c;
    font-size: 0.85rem;
    margin-top: 0.3rem;
    text-align: left;
  }
</style>
