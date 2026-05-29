<script setup>
  import { ref } from 'vue'
  import { useRouter } from 'vue-router'
  import Input from '@/components/EmailInput.vue'
  import PasswordInput from '@/components/PasswordInput.vue'
  import BackButton from '@/components/BackButton.vue'
  import { authAPI } from '@/api/auth'

  const router = useRouter()
  const email = ref('')
  const password = ref('')
  const confirmPassword = ref('')
  const errorMessage = ref('')

  const handleRegister = async () => {
    errorMessage.value = ''
    if (!email.value || !password.value || !confirmPassword.value) {
      errorMessage.value = 'Заполните все поля'
      return
    }
    if (password.value !== confirmPassword.value) {
      errorMessage.value = 'Пароли не совпадают'
      return
    }
    try {
      const response = await authAPI.register(email.value, password.value)
      // При успешной регистрации сервер возвращает 202 и { message, email }
      console.log('Регистрация успешна:', response.data)
      router.push({ name: 'verify-email', query: { email: email.value } })
    } catch (error) {
      if (error.response) {
        // ошибка с ответом от сервера (400, 409, 500)
        errorMessage.value = error.response.data?.error || error.response.data || 'Ошибка сервера'
      } else {
        errorMessage.value = 'Сетевая ошибка. Проверьте подключение.'
      }
    }
  }
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
          <div
            id="iw3pdlkas_0"
            class="container container--u-iw3pdlkas container--s1-iw3pdlkas container--s2-iw3pdlkas">
            <BackButton to="/" label="← На главную" />
          </div>
        </div>

        <div
          id="iryfzmc8c_0"
          class="container container--u-iryfzmc8c container--s1-iryfzmc8c container--s2-iryfzmc8c">
          <div class="form-content">
            <h1 class="logo">Регистрация</h1>
            <p class="subtitle">Создайте новый аккаунт</p>

            <form class="login-form" @submit.prevent="handleRegister">
              <Input
                id="email"
                v-model="email"
                label="Email"
                type="email"
                placeholder="your@email.com" />

              <PasswordInput
                id="password"
                v-model="password"
                label="Пароль"
                placeholder="••••••••" />

              <PasswordInput
                id="confirm-password"
                v-model="confirmPassword"
                label="Подтвердите пароль"
                placeholder="••••••••" />

              <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>

              <button type="submit" class="submit-button">Зарегистрироваться</button>
            </form>

            <p class="register-link">
              Уже есть аккаунт?
              <router-link to="/login" class="link-green">Войти</router-link>
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
  /* === Общие стили (как в HomeView и LoginView) === */

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

  /* Форма */
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

  .login-form {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    width: 100%;
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
      box-shadow 0.2s;
    box-shadow:
      0px 10px 20px 0px rgba(88, 204, 2, 0.04),
      0px 5px 10px 0px rgba(0, 0, 0, 0.06);

    margin-top: 0.5rem;
  }

  .submit-button:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(88, 204, 2, 0.3);
    box-shadow: 0 6px 16px rgba(88, 204, 2, 0.4);
    background: linear-gradient(135deg, #58cc02, #4cae01);
    color: white;
  }

  .error-message {
    color: #e74c3c;
    font-size: 0.9rem;
    margin: -0.5rem 0 0;
    text-align: left;
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
</style>
