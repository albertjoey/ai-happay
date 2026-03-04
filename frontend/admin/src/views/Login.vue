<template>
  <div class="login-container">
    <a-card class="login-card" title="Happy 管理后台">
      <a-form
        :model="formState"
        :rules="rules"
        @finish="handleLogin"
      >
        <a-form-item name="username">
          <a-input
            v-model:value="formState.username"
            placeholder="用户名"
            size="large"
          >
            <template #prefix>
              <UserOutlined />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item name="password">
          <a-input-password
            v-model:value="formState.password"
            placeholder="密码"
            size="large"
          >
            <template #prefix>
              <LockOutlined />
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item>
          <a-button type="primary" html-type="submit" size="large" block :loading="loading">
            登录
          </a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import { message } from 'ant-design-vue';
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue';

const router = useRouter();
const loading = ref(false);

const formState = reactive({
  username: 'admin',
  password: 'admin123',
});

const rules = {
  username: [{ required: true, message: '请输入用户名' }],
  password: [{ required: true, message: '请输入密码' }],
};

const handleLogin = async () => {
  loading.value = true;
  try {
    // 模拟登录
    await new Promise(resolve => setTimeout(resolve, 1000));
    localStorage.setItem('admin_token', 'mock_token_12345');
    message.success('登录成功');
    router.push('/');
  } catch (error) {
    message.error('登录失败');
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.login-card {
  width: 400px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}
</style>
