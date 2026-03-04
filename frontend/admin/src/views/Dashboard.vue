<template>
  <div class="dashboard">
    <a-row :gutter="16">
      <a-col :span="6">
        <a-card>
          <a-statistic
            title="总用户数"
            :value="statistics.totalUsers"
            :value-style="{ color: '#3f8600' }"
          >
            <template #prefix>
              <UserOutlined />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card>
          <a-statistic
            title="总内容数"
            :value="statistics.totalContents"
            :value-style="{ color: '#1890ff' }"
          >
            <template #prefix>
              <FileOutlined />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card>
          <a-statistic
            title="今日新增"
            :value="statistics.todayNew"
            :value-style="{ color: '#cf1322' }"
          >
            <template #prefix>
              <RiseOutlined />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card>
          <a-statistic
            title="活跃用户"
            :value="statistics.activeUsers"
            :value-style="{ color: '#722ed1' }"
          >
            <template #prefix>
              <TeamOutlined />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
    </a-row>

    <a-row :gutter="16" style="margin-top: 16px">
      <a-col :span="12">
        <a-card title="内容分布">
          <div id="content-chart" style="height: 300px"></div>
        </a-card>
      </a-col>
      <a-col :span="12">
        <a-card title="最近动态">
          <a-list :data-source="recentActivities">
            <template #renderItem="{ item }">
              <a-list-item>
                <a-list-item-meta :description="item.time">
                  <template #title>{{ item.content }}</template>
                </a-list-item-meta>
              </a-list-item>
            </template>
          </a-list>
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { UserOutlined, FileOutlined, RiseOutlined, TeamOutlined } from '@ant-design/icons-vue';

const statistics = ref({
  totalUsers: 12345,
  totalContents: 67890,
  todayNew: 234,
  activeUsers: 5678,
});

const recentActivities = ref([
  { content: '用户 张三 发布了新视频', time: '2分钟前' },
  { content: '用户 李四 创建了新话题', time: '5分钟前' },
  { content: '管理员 审核通过了3条内容', time: '10分钟前' },
  { content: '用户 王五 更新了个人资料', time: '15分钟前' },
  { content: '系统 完成了数据备份', time: '30分钟前' },
]);

onMounted(() => {
  // 这里可以初始化图表
  console.log('Dashboard mounted');
});
</script>

<style scoped>
.dashboard {
  padding: 0;
}
</style>
