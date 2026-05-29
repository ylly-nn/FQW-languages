<script setup>
  import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import Input from '@/components/EmailInput.vue'
  import BackButton from '@/components/BackButton.vue'
  import { authAPI } from '@/api/auth'

  const route = useRoute()
  const router = useRouter()

  // Восстанавливаем email/пароль из query или localStorage
  const saved = JSON.parse(localStorage.getItem('pendingVerification') || '{}')
  const email = ref(route.query.email || saved.email || '')
  const password = ref(route.query.password || saved.password || '')

  const code = ref('')
  const errorMessage = ref('')
  const successMessage = ref('')

  // Если нет данных для верификации — возвращаем на регистрацию
  if (!email.value || !password.value) {
    router.replace('/register')
  }

  // Таймер повторной отправки
  const cooldownSeconds = 60
  const lastResendTime = ref(null)
  const canResend = computed(() => {
    if (!lastResendTime.value) return true
    const elapsed = (Date.now() - lastResendTime.value) / 1000
    return elapsed >= cooldownSeconds
  })
  const cooldownRemaining = ref(0)
  let countdownInterval = null

  function startCountdown(initial) {
    cooldownRemaining.value = initial
    clearInterval(countdownInterval)
    countdownInterval = setInterval(() => {
      cooldownRemaining.value--
      if (cooldownRemaining.value <= 0) {
        clearInterval(countdownInterval)
        cooldownRemaining.value = 0
      }
    }, 1000)
  }

  onMounted(() => {
    if (lastResendTime.value) {
      const elapsed = (Date.now() - lastResendTime.value) / 1000
      if (elapsed < cooldownSeconds) {
        startCountdown(cooldownSeconds - Math.floor(elapsed))
      }
    }
  })

  const resendCode = async () => {
    errorMessage.value = ''
    successMessage.value = ''

    if (!canResend.value) return

    try {
      await authAPI.register(email.value, password.value)
      lastResendTime.value = Date.now()
      startCountdown(cooldownSeconds)
      successMessage.value = 'Код повторно отправлен на ваш email.'
    } catch (error) {
      errorMessage.value = error.response?.data?.error || 'Ошибка при повторной отправке кода.'
    }
  }

  const handleVerify = async () => {
    errorMessage.value = ''
    if (!email.value || !code.value) {
      errorMessage.value = 'Введите email и код'
      return
    }
    try {
      await authAPI.verify(email.value, code.value)
      // Успешно — удаляем временные данные и перенаправляем
      localStorage.removeItem('pendingVerification')
      router.push('/login')
    } catch (error) {
      errorMessage.value = error.response?.data?.error || 'Неверный код'
    }
  }

  onBeforeUnmount(() => {
    clearInterval(countdownInterval)
  })
</script>

<template>
  <div class="mosaic-wrap">
    <div id="i8z5xo3zc_0" class="root root--u-i8z5xo3zc root--s1-i8z5xo3zc root--s2-i8z5xo3zc">
      <div
        id="ikc0yythz_0"
        class="container container--u-ikc0yythz container--s1-ikc0yythz container--s2-ikc0yythz">
        <div
          id="iw3pdlkas_0"
          class="container container--u-iw3pdlkas container--s1-iw3pdlkas container--s2-iw3pdlkas">
          <BackButton to="/register" label="← Назад" />
        </div>

        <div
          id="iryfzmc8c_0"
          class="container container--u-iryfzmc8c container--s1-iryfzmc8c container--s2-iryfzmc8c">
          <div class="form-content">
            <h1 class="logo">Подтверждение</h1>
            <p class="subtitle">На ваш email отправлен код. Введите его ниже.</p>

            <form class="verify-form" @submit.prevent="handleVerify">
              <Input
                id="email"
                v-model="email"
                label="Email"
                type="email"
                placeholder="your@email.com"
                disabled />
              <Input id="code" v-model="code" label="Код" type="text" placeholder="123456" />

              <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
              <p v-if="successMessage" class="success-message">{{ successMessage }}</p>

              <button type="submit" class="submit-button">Подтвердить</button>
            </form>

            <div class="resend-block">
              <button class="resend-button" :disabled="!canResend" @click="resendCode">
                <span v-if="canResend">Отправить код снова</span>
                <span v-else>Повторно через {{ cooldownRemaining }} сек</span>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
  /* === Базовые стили (как в HomeView и LoginView) === */

  .root--u-i8z5xo3zc {
    background-size: auto;
    background-image: linear-gradient(rgba(255, 253, 247, 1) 0%, rgba(255, 253, 247, 1) 100%);
    background-repeat: no-repeat;
    background-position: left 0px top 0px;
    background-attachment: scroll;
    display: flex;
    align-items: center;
    justify-content: center;
    min-height: 100vh;
  }

  .container--u-ikc0yythz {
    min-height: 100vh;
    width: 100%;
    max-width: 1200px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 2rem 0;
  }

  .container--u-iw3pdlkas {
    display: flex;
    justify-content: flex-start;
    align-items: center;
    width: 100%;
    max-width: 1000px;
    padding: 0 2rem;
    margin-bottom: 2rem;
  }

  /* Белый центральный блок */
  .container--u-iryfzmc8c {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 90%;
    max-width: 500px;
    min-height: auto;
    margin: 0 auto;
    padding: 3rem 2.5rem;
    border-radius: 50px;
    background-size: auto;
    background-image: linear-gradient(rgba(255, 255, 255, 1) 0%, rgba(255, 255, 255, 1) 100%);
    background-repeat: no-repeat;
    background-position: left 0px top 0px;
    background-attachment: scroll;
    border: 1px solid rgba(88, 204, 2, 0.6);
    box-shadow:
      0px 20px 40px 0px rgba(88, 204, 2, 0.06),
      0px 10px 20px 0px rgba(0, 0, 0, 0.04);
  }

  /* Стили формы */
  .form-content {
    width: 100%;
    text-align: center;
  }

  .logo {
    font-size: 2.5rem;
    color: #2c3e50;
    margin-bottom: 0.5rem;
    font-weight: 700;
  }

  .subtitle {
    font-size: 1.1rem;
    color: #777;
    margin-bottom: 2rem;
    font-weight: 300;
  }

  .verify-form {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .error-message {
    color: #e74c3c;
    font-size: 0.9rem;
    margin: 0;
    text-align: center;
  }

  .submit-button {
    width: 100%;
    padding: 0.9rem;
    font-size: 1.1rem;
    font-weight: 600;
    color: #2c3e50;
    background-image: linear-gradient(rgba(255, 255, 255, 1) 0%, rgba(255, 255, 255, 1) 100%);
    border: 1px solid rgba(88, 204, 2, 0.25);
    border-radius: 50px;
    cursor: pointer;
    transition:
      transform 0.2s,
      box-shadow 0.2s,
      background 0.2s,
      color 0.2s;
    box-shadow:
      0px 10px 20px 0px rgba(88, 204, 2, 0.04),
      0px 5px 10px 0px rgba(0, 0, 0, 0.06);
    margin-top: 0.5rem;
  }

  .submit-button:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 16px rgba(88, 204, 2, 0.4);
    background: linear-gradient(135deg, #58cc02, #4cae01);
    color: white;
  }

  .register-link {
    margin-top: 2rem;
    font-size: 0.95rem;
    color: #555;
  }

  .link-green {
    color: #58cc02;
    text-decoration: none;
    font-weight: 600;
    transition: color 0.2s;
  }

  .link-green:hover {
    color: #3d8b01;
  }

  .success-message {
    color: #58cc02;
    font-size: 0.9rem;
    margin: 0;
    text-align: center;
  }

  .resend-block {
    margin-top: 2rem;
    text-align: center;
  }

  .resend-button {
    background: transparent;
    border: none;
    color: #58cc02;
    font-weight: 600;
    cursor: pointer;
    font-size: 0.95rem;
    text-decoration: underline;
    transition: opacity 0.2s;
  }

  .resend-button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
    text-decoration: none;
  }

  .resend-button:hover:not(:disabled) {
    color: #3d8b01;
  }
</style>
