import MainLayout from '../components/MainLayout.vue';
import HomePage from '../views/HomePage.vue';
import MemberPage from '../views/MemberPage.vue';
import WorkPage from '@/views/WorkPage.vue';
import RecordPage from '@/views/RecordPage.vue';
import RewardPage from '@/views/RewardPage.vue';
import AboutPage from '@/views/AboutPage.vue';
import LoginPage from '@/views/LoginPage.vue';
import AdminPage from '@/views/manage/AdminPage.vue';
import MenuLayout from '@/components/MenuLayout.vue';
import MemberManage from '@/views/manage/MemberManage.vue';
import WorkManage from '@/views/manage/WorkManage.vue';
import RecordManage from '@/views/manage/RecordManage.vue';
import RewardManage from '@/views/manage/RewardManage.vue';
import IssueManage from '@/views/manage/IssueManage.vue';
import { createRouter, createWebHistory } from 'vue-router';



const routes = [
  {
    path: '/',
    component: MainLayout,
    children: [
      {
        path: '', // 默认子路由
        name: 'Home',
        component: HomePage
      },
      {
        path: 'members',
        name: 'Member',
        component: MemberPage
      },
      {
        path: 'work',
        name: 'Work',
        component: WorkPage
      },
      {
        path: 'record',
        name: 'Record',
        component: RecordPage
      },
      {
        path: 'reward',
        name: 'Reward',
        component: RewardPage
      },
      {
        path: 'about',
        name: 'About',
        component: AboutPage
      }
      // 更多页面...
    ]
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginPage
  },
  {
    path: '/admin',
    component: MenuLayout,
    meta: { requiresAuth: true, breadcrumb: '主界面' },
    children: [
      { path: '', name: 'Admin', component: AdminPage },
      { path: 'member', name: 'MemberManage', component: MemberManage, meta: { breadcrumb: '成员管理' } },
      { path: 'work', name: 'WorkManage', component: WorkManage, meta: { breadcrumb: '成果管理' } },
      { path: 'record', name: 'RecordManage', component: RecordManage, meta: { breadcrumb: '记录管理' } },
      { path: 'reward', name: 'RewardManage', component: RewardManage, meta: { breadcrumb: '奖励管理' } },
      { path: 'issues', name: 'IssueManage', component: IssueManage, meta: { breadcrumb: '问题管理' } }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token');

  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!token || isTokenExpired(token)) {
      next({ name: 'Login' });
    } else {
      next();
    }
  } else {
    next();
  }
});

function parseJwt(token) {
  try {
      const base64Url = token.split('.')[1];
      const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
      const jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
          return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
      }).join(''));

      return JSON.parse(jsonPayload);
  } catch (error) {
      return null;
  }
}

function isTokenExpired(token) {
  const payload = parseJwt(token);
  if (!payload) return true;  // 如果 token 解析失败，视为已过期

  const currentTime = Date.now() / 1000;  // 获取当前时间的 UNIX 时间戳
  return payload.exp < currentTime;  // 比较过期时间
}

export default router;
