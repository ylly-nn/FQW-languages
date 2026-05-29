import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  // Главная страница
  {
    path: '/',
    name: 'home',
    component: () => import('@/views/HomeView.vue')
  },

  // Вход и регистрация
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/LoginView.vue')
  },
  {
  path: '/password-recovery',
  name: 'password-recovery',
  component: () => import('@/views/PasswordRecView.vue')
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('@/views/RegisterView.vue')
  },
  {
    path: '/verify-email',
    name: 'verify-email',
    component: () => import('@/views/CodeVerifyView.vue')
  },

  // Выбор языка
  {
    path: '/language',
    name: 'language',
    component: () => import('@/views/LangSelectView.vue'),
    meta: { requiresAuth: true }
  },

  // Главная страница с заданиями
  {
    path: '/dashboard',
    name: 'dashboard',
    component: () => import('@/views/DashboardView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/dictionary',
    name: 'dictionary',
    component: () => import('@/views/DictionaryView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/progress',
    name: 'progress',
    component: () => import('@/views/ProgressView.vue'),
    meta: { requiresAuth: true }
  },

  // Блок аудирования
  {
    path: '/audition/topics',
    name: 'audition-topics',
    component: () => import('@/views/audition/LvlTopicsView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/audition/topics/:topicId',
    name: 'audition-topic-lessons',
    component: () => import('@/views/audition/TopicView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/audition/lesson/:id',
    name: 'audition-lesson',
    component: () => import('@/views/audition/LessonView.vue'),
    meta: { requiresAuth: true }
  },

  // Блок слов
  {
    path: '/vocabulary/topics',
    name: 'VocabularyTopics',
    component: () => import('@/views/vocabulary/LvlTopicsVue.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/vocabulary/topics/:topicId',
    name: 'VocabularyTopicDetail',
    component: () => import('@/views/vocabulary/TopicVue.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/vocabulary/lesson/:lessonId/:wordIndex?',
    name: 'VocabularyLesson',
    component: () => import('@/views/vocabulary/LessonView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/vocabulary/lesson/:lessonId/:wordIndex/text',
    name: 'VocabularyTextLesson',
    component: () => import('@/views/vocabulary/TextLessonView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/vocabulary/lesson/:lessonId/cards',
    name: 'VocabularyCardText',
    component: () => import('@/views/vocabulary/CardTextView.vue'),
    meta: { requiresAuth: true }
  },

  // Грамматика
  {
    path: '/grammar/topics',
    name: 'grammar-topics',
    component: () => import('@/views/grammar/LvlTopicsView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/grammar/topics/:topicId',
    name: 'grammar-topic-theory',
    component: () => import('@/views/grammar/TopicView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/grammar/topics/:topicId/exercises',
    name: 'grammar-topic-exercises',
    component: () => import('@/views/grammar/ExercisesView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/grammar/exercise/:topicId/:type',
    name: 'grammar-exercise',
    component: () => import('@/views/grammar/ExerciseView.vue'),
    meta: { requiresAuth: true }
  },

  // 404 — Страница не найдена
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/ErrView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

// Navigation guard
router.beforeEach((to, from, next) => {
  const accessToken = localStorage.getItem('access_token')
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth)

  if (requiresAuth && !accessToken) {
    next({ name: 'login' })
  } else {
    next()
  }
})

export default router