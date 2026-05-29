import apiClient from './client'

export const authAPI = {
  register(email, password) {
    return apiClient.post('/auth/register', { email, password })
  },
   verify(email, code) {
    return apiClient.post('/auth/verify', { email, code })
  },

  login(email, password) {
    return apiClient.post('/auth/login', { email, password })
  },

  logout(refreshToken) {
    return apiClient.post('/auth/logout', { refresh_token: refreshToken })
  },

   forgotPassword(email) {
    return apiClient.post('/auth/forgot-password', { email })
  },

  // Проверка кода восстановления
  verifyResetCode(email, code) {
    return apiClient.post('/auth/verify-reset-code', { email, code })
  },

  // Установка нового пароля
  setNewPassword(email, newPassword) {
    return apiClient.post('/auth/set-password', { email, new_password: newPassword })
  },

  refresh(refreshToken) {
    return apiClient.post('/auth/refresh', { refresh_token: refreshToken })
  }
}