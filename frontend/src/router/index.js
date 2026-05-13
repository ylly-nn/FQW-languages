import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'
import LanguageView from '@/views/LangSelectView.vue'
import DashboardView from '@/views/DashboardView.vue'
import ListeningView from '@/views/ListeningView.vue'
import TopicsView from '@/views/TopicsView.vue'
import TopicLessonsView from '@/views/TopicLessonsView.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
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
    {
      path: '/language',
      name: 'language',
      component: LanguageView,
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: DashboardView,
    },
    {
      path: '/audition/topics',
      name: 'audition-topics',
      component: TopicsView,
    },
    {
      path: '/audition/topics/:topicId',
      name: 'audition-topic-lessons',
      component: TopicLessonsView,
    },
    {
      path: '/audition/lesson/:id',
      name: 'audition-lesson',
      component: ListeningView,
    },
  ],
})

export default router
