import { createRouter, createWebHistory } from 'vue-router'

//! Страница не найдена
import ErrView from '@/views/ErrView.vue'

//# Главная страницa
import HomeView from '@/views/HomeView.vue'

//# Вход и регистрация
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'

//#Выбор языка
import LanguageView from '@/views/LangSelectView.vue'

//# Главная страница с заданиями
import DashboardView from '@/views/DashboardView.vue'
import DictionaryView from '@/views/DictionaryView.vue'
import ProgressView from '@/views/ProgressView.vue'

//# Блок аудирования
import AudLvlTopicsView from '@/views/audition/LvlTopicsView.vue'
import AudTopicView from '@/views/audition/TopicView.vue'
import AudLessonView from '@/views/audition/LessonView.vue'

//# Блок слов
import VocLvlTopicsView from '@/views/vocabulary/LvlTopicsVue.vue'
import VocTopicView from '@/views/vocabulary/TopicVue.vue'
import VocLessonView from '@/views/vocabulary/LessonView.vue'
import VocTextLessonView from '@/views/vocabulary/TextLessonView.vue'
import VocCardTextView from '@/views/vocabulary/CardTextView.vue'

//# Грамматика
import GrammarLvlTopicsView from '@/views/grammar/LvlTopicsView.vue'
import GrammarTopicView from '@/views/grammar/TopicView.vue'
import GrammarExercisesView from '@/views/grammar/ExercisesView.vue'
import GrammarExerciseView from '@/views/grammar/ExerciseView.vue'

//Do Блок грамматики

const routes = [
  //# Главная страницa
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },

  //# Вход и регистрация
  {
    path: '/login',
    name: 'login',
    component: LoginView,
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterView,
  },

  //#Выбор языка
  {
    path: '/language',
    name: 'language',
    component: LanguageView,
  },

  //# Главная страница с заданиями
  {
    path: '/dashboard',
    name: 'dashboard',
    component: DashboardView,
  },
  {
    path: '/dictionary',
    name: 'dictionary',
    component: DictionaryView,
  },
  {
    path: '/progress',
    name: 'progress',
    component: ProgressView,
  },

  //# Блок аудирования
  {
    path: '/audition/topics',
    name: 'audition-topics',
    component: AudLvlTopicsView,
  },
  {
    path: '/audition/topics/:topicId',
    name: 'audition-topic-lessons',
    component: AudTopicView,
  },
  {
    path: '/audition/lesson/:id',
    name: 'audition-lesson',
    component: AudLessonView,
  },

  //# Блок слов
  {
    path: '/vocabulary/topics',
    name: 'VocabularyTopics',
    component: VocLvlTopicsView,
  },
  {
    path: '/vocabulary/topics/:topicId',
    name: 'VocabularyTopicDetail',
    component: VocTopicView,
  },
  {
    path: '/vocabulary/lesson/:lessonId/:wordIndex?',
    name: 'VocabularyLesson',
    component: VocLessonView,
  },
  {
    path: '/vocabulary/lesson/:lessonId/:wordIndex/text',
    name: 'VocabularyTextLesson',
    component: VocTextLessonView,
  },
  {
    path: '/vocabulary/lesson/:lessonId/cards',
    name: 'VocabularyCardText',
    component: VocCardTextView,
  },

  //# Грамматика
  {
    path: '/grammar/topics',
    name: 'grammar-topics',
    component: GrammarLvlTopicsView,
  },
  {
    path: '/grammar/topics/:topicId',
    name: 'grammar-topic-theory',
    component: GrammarTopicView,
  },
  {
    path: '/grammar/topics/:topicId/exercises',
    name: 'grammar-topic-exercises',
    component: GrammarExercisesView,
  },
  {
    path: '/grammar/exercise/:topicId/:type',
    name: 'grammar-exercise',
    component: GrammarExerciseView,
  },

  //! Страница не найдена - все остальные
  {
    path: '/:pathMatch(.*)*', // ловит любой путь, который не совпал выше
    name: 'NotFound',
    component: ErrView,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
