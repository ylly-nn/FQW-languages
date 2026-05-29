<script setup>
  import { ref } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import Input from '@/components/EmailInput.vue' // или простой input
  import BackButton from '@/components/BackButton.vue'
  import { authAPI } from '@/api/auth'

  const route = useRoute()
  const router = useRouter()
  const email = ref(route.query.email || '')
  const code = ref('')
  const errorMessage = ref('')

  const handleVerify = async () => {
    errorMessage.value = ''
    if (!email.value || !code.value) {
      errorMessage.value = 'Введите email и код'
      return
    }
    try {
      await authAPI.verify(email.value, code.value) // метод verify нужно добавить в auth.js
      // успех — перенаправляем на логин
      router.push('/login')
    } catch (error) {
      errorMessage.value = error.response?.data?.error || 'Неверный код'
    }
  }
</script>

<template>
  <div class="verify-page">
    <BackButton to="/register" label="← Назад" />
    <h1>Подтверждение email</h1>
    <p>На ваш email отправлен код подтверждения. Введите его ниже.</p>
    <form @submit.prevent="handleVerify">
      <Input id="email" v-model="email" label="Email" type="email" placeholder="your@email.com" />
      <Input id="code" v-model="code" label="Код" type="text" placeholder="123456" />
      <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
      <button type="submit">Подтвердить</button>
    </form>
  </div>
</template>

<style scoped>
  /* добавьте стили по аналогии с другими формами */
</style>
