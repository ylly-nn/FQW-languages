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
    component: () => import('@/views/LangSelectView.vue')
  },

  // Главная страница с заданиями
  {
    path: '/dashboard',
    name: 'dashboard',
    component: () => import('@/views/DashboardView.vue')
  },
  {
    path: '/dictionary',
    name: 'dictionary',
    component: () => import('@/views/DictionaryView.vue')
  },
  {
    path: '/progress',
    name: 'progress',
    component: () => import('@/views/ProgressView.vue')
  },

  // Блок аудирования
  {
    path: '/audition/topics',
    name: 'audition-topics',
    component: () => import('@/views/audition/LvlTopicsView.vue')
  },
  {
    path: '/audition/topics/:topicId',
    name: 'audition-topic-lessons',
    component: () => import('@/views/audition/TopicView.vue')
  },
  {
    path: '/audition/lesson/:id',
    name: 'audition-lesson',
    component: () => import('@/views/audition/LessonView.vue')
  },

  // Блок слов
  {
    path: '/vocabulary/topics',
    name: 'VocabularyTopics',
    component: () => import('@/views/vocabulary/LvlTopicsVue.vue')
  },
  {
    path: '/vocabulary/topics/:topicId',
    name: 'VocabularyTopicDetail',
    component: () => import('@/views/vocabulary/TopicVue.vue')
  },
  {
    path: '/vocabulary/lesson/:lessonId/:wordIndex?',
    name: 'VocabularyLesson',
    component: () => import('@/views/vocabulary/LessonView.vue')
  },
  {
    path: '/vocabulary/lesson/:lessonId/:wordIndex/text',
    name: 'VocabularyTextLesson',
    component: () => import('@/views/vocabulary/TextLessonView.vue')
  },
  {
    path: '/vocabulary/lesson/:lessonId/cards',
    name: 'VocabularyCardText',
    component: () => import('@/views/vocabulary/CardTextView.vue')
  },

  // Грамматика
  {
    path: '/grammar/topics',
    name: 'grammar-topics',
    component: () => import('@/views/grammar/LvlTopicsView.vue')
  },
  {
    path: '/grammar/topics/:topicId',
    name: 'grammar-topic-theory',
    component: () => import('@/views/grammar/TopicView.vue')
  },
  {
    path: '/grammar/topics/:topicId/exercises',
    name: 'grammar-topic-exercises',
    component: () => import('@/views/grammar/ExercisesView.vue')
  },
  {
    path: '/grammar/exercise/:topicId/:type',
    name: 'grammar-exercise',
    component: () => import('@/views/grammar/ExerciseView.vue')
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

export default router