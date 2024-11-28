import { createRouter, createWebHistory } from 'vue-router'

export const names = {
  home: 'home',
  enrollment: 'enrollment',
  sequence: 'sequence',
  verifiedSequence: 'verified-sequence',
  leaderBoard: 'leader-board',
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: names.home,
      component: () => import('../views/HomeView.vue'),
    },
    {
      path: '/' + names.enrollment,
      name: names.enrollment,
      component: () => import('../views/EnrollmentView.vue'),
      meta: { title: 'Enrollment Page' },
    },
    {
      path: '/' + names.sequence,
      name: names.sequence,
      component: () => import('../views/SequenceView.vue'),
      meta: { title: 'Sequence Page' },
    },
    {
      path: '/' + names.sequence,
      name: names.sequence,
      component: () => import('../views/SequenceView.vue'),
      meta: { title: 'Sequence Page' },
    },
    {
      path: '/' + names.verifiedSequence,
      name: names.verifiedSequence,
      component: () => import('../views/VerifiedView.vue'),
      meta: { title: 'Verified Sequence' },
    },
    {
      path: '/' + names.leaderBoard,
      name: names.leaderBoard,
      component: () => import('../views/LeaderBoardView.vue'),
      meta: { title: 'Leader Board' },
    },
    // Catch-all route for 404
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: () => import('../views/HomeView.vue'),
      meta: { title: 'Page Not Found' },
    },
  ],
})

// Update document title on route change
router.afterEach((to) => {
  document.title = to.meta.title ? `${to.meta.title} | Canary Gotchi` : 'Canary Gotchi'
})

export default router
