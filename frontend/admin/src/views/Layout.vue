<template>
  <a-layout style="min-height: 100vh">
    <a-layout-sider v-model:collapsed="collapsed" collapsible>
      <div class="logo">
        <h1 v-if="!collapsed">Happy Admin</h1>
        <h1 v-else>H</h1>
      </div>
      <a-menu
        v-model:selectedKeys="selectedKeys"
        v-model:openKeys="openKeys"
        theme="dark"
        mode="inline"
        @click="handleMenuClick"
      >
        <a-menu-item key="dashboard">
          <dashboard-outlined />
          <span>仪表盘</span>
        </a-menu-item>
        <a-menu-item key="user">
          <user-outlined />
          <span>用户管理</span>
        </a-menu-item>
        <a-menu-item key="material">
          <picture-outlined />
          <span>物料管理</span>
        </a-menu-item>
        <a-menu-item key="topic">
          <tag-outlined />
          <span>话题管理</span>
        </a-menu-item>
        <a-menu-item key="tag">
          <tags-outlined />
          <span>标签管理</span>
        </a-menu-item>
        <a-sub-menu key="discover">
          <template #icon>
            <compass-outlined />
          </template>
          <template #title>发现页管理</template>
          <a-menu-item key="discover-config">
            <span>模块配置</span>
          </a-menu-item>
          <a-menu-item key="discover-items">
            <span>内容管理</span>
          </a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="channel">
          <template #icon>
            <appstore-outlined />
          </template>
          <template #title>频道管理</template>
          <a-menu-item key="channel-list">
            <span>频道列表</span>
          </a-menu-item>
          <a-menu-item key="channel-diamond">
            <span>金刚位管理</span>
          </a-menu-item>
          <a-menu-item key="channel-recommend">
            <span>推荐位管理</span>
          </a-menu-item>
          <a-menu-item key="channel-feed-config">
            <span>Feed流配置</span>
          </a-menu-item>
          <a-menu-item key="channel-ad-slot">
            <span>广告位管理</span>
          </a-menu-item>
          <a-menu-item key="channel-config">
            <span>频道配置</span>
          </a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="system">
          <template #icon>
            <setting-outlined />
          </template>
          <template #title>系统管理</template>
          <a-menu-item key="system-role">
            <span>角色管理</span>
          </a-menu-item>
          <a-menu-item key="system-permission">
            <span>权限管理</span>
          </a-menu-item>
          <a-menu-item key="system-admin-user">
            <span>管理员管理</span>
          </a-menu-item>
        </a-sub-menu>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="background: #fff; padding: 0 16px">
        <div class="header-content">
          <a-breadcrumb>
            <a-breadcrumb-item>首页</a-breadcrumb-item>
            <a-breadcrumb-item>{{ currentRoute?.meta?.title }}</a-breadcrumb-item>
          </a-breadcrumb>
          <div class="header-right">
            <a-dropdown>
              <a class="ant-dropdown-link" @click.prevent>
                <a-avatar style="background-color: #1890ff">
                  <template #icon><UserOutlined /></template>
                </a-avatar>
                <span style="margin-left: 8px">管理员</span>
              </a>
              <template #overlay>
                <a-menu>
                  <a-menu-item key="profile">个人中心</a-menu-item>
                  <a-menu-item key="logout" @click="handleLogout">退出登录</a-menu-item>
                </a-menu>
              </template>
            </a-dropdown>
          </div>
        </div>
      </a-layout-header>
      <a-layout-content style="margin: 16px">
        <router-view :key="$route.name || $route.path" />
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import {
  DashboardOutlined,
  UserOutlined,
  PictureOutlined,
  FileOutlined,
  TagOutlined,
  TagsOutlined,
  AppstoreOutlined,
  SettingOutlined,
  CompassOutlined,
} from '@ant-design/icons-vue';

const router = useRouter();
const route = useRoute();
const collapsed = ref(false);
const selectedKeys = ref(['dashboard']);
const openKeys = ref(['channel', 'system', 'discover']);

const currentRoute = computed(() => route);

const handleMenuClick = ({ key }: { key: string }) => {
  if (key.startsWith('channel-')) {
    const path = key.replace('channel-', '');
    router.push(`/channel/${path}`);
  } else if (key.startsWith('system-')) {
    const path = key.replace('system-', '');
    router.push(`/system/${path}`);
  } else if (key.startsWith('discover-')) {
    const path = key.replace('discover-', '');
    router.push(`/discover/${path}`);
  } else {
    router.push(`/${key}`);
  }
};

const handleLogout = () => {
  localStorage.removeItem('admin_token');
  router.push('/login');
};
</script>

<style scoped>
.logo {
  height: 32px;
  margin: 16px;
  background: rgba(255, 255, 255, 0.3);
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.logo h1 {
  color: #fff;
  font-size: 18px;
  margin: 0;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-right {
  display: flex;
  align-items: center;
}
</style>
