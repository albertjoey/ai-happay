<template>
  <div class="topic-list">
    <a-card title="话题管理">
      <template #extra>
        <a-button type="primary">
          <template #icon><PlusOutlined /></template>
          添加话题
        </a-button>
      </template>
      <a-empty description="话题管理功能开发中..." />
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { PlusOutlined } from '@ant-design/icons-vue';
</script>
