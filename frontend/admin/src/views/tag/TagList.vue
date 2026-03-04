<template>
  <div class="tag-list">
    <a-card title="标签管理">
      <template #extra>
        <a-button type="primary">
          <template #icon><PlusOutlined /></template>
          添加标签
        </a-button>
      </template>
      <a-empty description="标签管理功能开发中..." />
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { PlusOutlined } from '@ant-design/icons-vue';
</script>
