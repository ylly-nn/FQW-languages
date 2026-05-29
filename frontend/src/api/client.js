import axios from 'axios'

const apiClient = axios.create({
  baseURL: '/api',        // все запросы будут начинаться с /api
  headers: {
    'Content-Type': 'application/json',
  },
})

// Перехватчик для добавления токена, если он есть
apiClient.interceptors.request.use(config => {
  const token = localStorage.getItem('access_token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

export default apiClient