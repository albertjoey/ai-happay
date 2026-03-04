import { createRouter, createWebHistory } from 'vue-router';
import type { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue'),
    meta: { title: '登录' },
  },
  {
    path: '/',
    component: () => import('@/views/Layout.vue'),
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/Dashboard.vue'),
        meta: { title: '仪表盘', icon: 'DashboardOutlined' },
      },
      {
        path: 'user',
        name: 'User',
        component: () => import('@/views/user/UserList.vue'),
        meta: { title: '用户管理', icon: 'UserOutlined' },
      },
      {
        path: 'material',
        name: 'Material',
        component: () => import('@/views/material/MaterialList.vue'),
        meta: { title: '物料管理', icon: 'PictureOutlined' },
      },
      {
        path: 'material/chapter',
        name: 'MaterialChapter',
        component: () => import('@/views/material/ChapterList.vue'),
        meta: { title: '章节管理' },
      },
      {
        path: 'topic',
        name: 'Topic',
        component: () => import('@/views/topic/TopicList.vue'),
        meta: { title: '话题管理', icon: 'TagOutlined' },
      },
      {
        path: 'tag',
        name: 'Tag',
        component: () => import('@/views/tag/TagList.vue'),
        meta: { title: '标签管理', icon: 'TagsOutlined' },
      },
      {
        path: 'channel',
        name: 'Channel',
        meta: { title: '频道管理', icon: 'AppstoreOutlined' },
        children: [
          {
            path: 'list',
            name: 'ChannelList',
            component: () => import('@/views/channel/ChannelList.vue'),
            meta: { title: '频道列表' },
          },
          {
            path: 'diamond',
            name: 'Diamond',
            component: () => import('@/views/channel/DiamondList.vue'),
            meta: { title: '金刚位管理' },
          },
          {
            path: 'recommend',
            name: 'Recommend',
            component: () => import('@/views/channel/RecommendList.vue'),
            meta: { title: '推荐位管理' },
          },
          {
            path: 'feed-config',
            name: 'FeedConfig',
            component: () => import('@/views/channel/FeedConfigList.vue'),
            meta: { title: 'Feed流配置' },
          },
          {
            path: 'ad-slot',
            name: 'AdSlot',
            component: () => import('@/views/channel/AdSlotList.vue'),
            meta: { title: '广告位管理' },
          },
          {
            path: 'config',
            name: 'ChannelConfig',
            component: () => import('@/views/channel/ChannelConfig.vue'),
            meta: { title: '频道配置' },
          },
        ],
      },
      {
        path: 'system',
        name: 'System',
        meta: { title: '系统管理', icon: 'SettingOutlined' },
        children: [
          {
            path: 'role',
            name: 'Role',
            component: () => import('@/views/system/RoleList.vue'),
            meta: { title: '角色管理' },
          },
          {
            path: 'permission',
            name: 'Permission',
            component: () => import('@/views/system/PermissionList.vue'),
            meta: { title: '权限管理' },
          },
          {
            path: 'admin-user',
            name: 'AdminUser',
            component: () => import('@/views/system/AdminUserList.vue'),
            meta: { title: '管理员管理' },
          },
        ],
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// 路由守卫
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('admin_token');
  if (to.path !== '/login' && !token) {
    next('/login');
  } else {
    next();
  }
});

export default router;
