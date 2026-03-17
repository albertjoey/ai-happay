<template>
  <div class="material-list">
    <a-card title="物料管理">
      <template #extra>
        <a-space>
          <a-select v-model:value="filterType" style="width: 120px" placeholder="物料类型" allowClear @change="handleFilter">
            <a-select-option value="">全部类型</a-select-option>
            <a-select-option value="image_text">图文</a-select-option>
            <a-select-option value="novel">小说</a-select-option>
            <a-select-option value="video">视频</a-select-option>
            <a-select-option value="banner">Banner</a-select-option>
            <a-select-option value="comic">漫剧</a-select-option>
            <a-select-option value="short_drama">短剧</a-select-option>
          </a-select>
          <a-input-search
            v-model:value="searchKeyword"
            placeholder="搜索标题/作者"
            style="width: 200px"
            @search="handleFilter"
          />
          <a-button type="primary" @click="handleAdd">
            <template #icon><PlusOutlined /></template>
            添加物料
          </a-button>
        </a-space>
      </template>

      <a-table
        :columns="columns"
        :data-source="tableData"
        :loading="loading"
        :pagination="{ pageSize: 10 }"
        row-key="id"
      >
        <template #bodyCell="{ column, record }">
        
        
        
        
        
        
        
        
        
        
      </template>
      </a-table>

      <div class="pagination">
        <a-pagination
          v-model:current="pagination.current"
          v-model:pageSize="pagination.pageSize"
          :total="pagination.total"
          show-size-changer
          @change="handlePageChange"
        />
      </div>
    </a-card>

    <!-- 添加/编辑物料对话框 -->
    <a-modal
      v-model:open="modalVisible"
      :title="modalTitle"
      @ok="handleModalOk"
      @cancel="handleModalCancel"
      width="800px"
    >
      <a-form
        ref="formRef"
        :model="formState"
        :rules="formRules"
        :label-col="{ span: 4 }"
        :wrapper-col="{ span: 18 }"
      >
        <a-form-item label="标题" name="title">
          <a-input v-model:value="formState.title" placeholder="请输入标题" :disabled="isView" />
        </a-form-item>
        <a-form-item label="副标题" name="subtitle">
          <a-input v-model:value="formState.subtitle" placeholder="请输入副标题" :disabled="isView" />
        </a-form-item>
        <a-form-item label="类型" name="type">
          <a-select v-model:value="formState.type" :disabled="isView">
            <a-select-option value="image_text">图文</a-select-option>
            <a-select-option value="novel">小说</a-select-option>
            <a-select-option value="video">视频</a-select-option>
            <a-select-option value="banner">Banner</a-select-option>
            <a-select-option value="comic">漫剧</a-select-option>
            <a-select-option value="short_drama">短剧</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="封面URL" name="cover_url">
          <a-input v-model:value="formState.cover_url" placeholder="请输入封面URL" :disabled="isView" />
        </a-form-item>
        <a-form-item label="内容URL" name="content_url">
          <a-input v-model:value="formState.content_url" placeholder="请输入内容URL(视频/音频等)" :disabled="isView" />
        </a-form-item>
        <a-form-item label="作者" name="author">
          <a-input v-model:value="formState.author" placeholder="请输入作者" :disabled="isView" />
        </a-form-item>
        <a-form-item label="描述" name="description">
          <a-textarea v-model:value="formState.description" placeholder="请输入描述" :rows="3" :disabled="isView" />
        </a-form-item>
        <a-form-item label="分类" name="category">
          <a-input v-model:value="formState.category" placeholder="请输入分类" :disabled="isView" />
        </a-form-item>
        <a-form-item v-if="formState.type === 'video' || formState.type === 'short_drama'" label="时长(秒)" name="duration">
          <a-input-number v-model:value="formState.duration" :min="0" :disabled="isView" />
        </a-form-item>
        <a-form-item v-if="formState.type === 'novel'" label="字数" name="word_count">
          <a-input-number v-model:value="formState.word_count" :min="0" :disabled="isView" />
        </a-form-item>
        <a-form-item v-if="formState.type === 'novel' || formState.type === 'comic'" label="章节数" name="chapter_count">
          <a-input-number v-model:value="formState.chapter_count" :min="0" :disabled="isView" />
        </a-form-item>
        <a-form-item label="状态" name="status">
          <a-radio-group v-model:value="formState.status" :disabled="isView">
            <a-radio :value="0">草稿</a-radio>
            <a-radio :value="1">已发布</a-radio>
            <a-radio :value="2">已下架</a-radio>
          </a-radio-group>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { message } from 'ant-design-vue';
import { PlusOutlined } from '@ant-design/icons-vue';
import type { FormInstance } from 'ant-design-vue';
import { getMaterialList, createMaterial, updateMaterial, deleteMaterial, type Material } from '@/api/material';

const router = useRouter();
const loading = ref(false);
const tableData = ref<Material[]>([]);
const filterType = ref('');
const searchKeyword = ref('');

const pagination = reactive({
  current: 1,
  pageSize: 20,
  total: 0,
});

const modalVisible = ref(false);
const modalTitle = ref('添加物料');
const isEdit = ref(false);
const isView = ref(false);
const currentId = ref<number>(0);
const formRef = ref<FormInstance>();

// 表格列定义
const columns = [
  { title: 'ID', dataIndex: 'id', key: 'id', width: 80 },
  { title: '标题', dataIndex: 'title', key: 'title', width: 200 },
  { title: '类型', dataIndex: 'type', key: 'type', width: 100 },
  { title: '状态', dataIndex: 'status', key: 'status', width: 100 },
  { title: '创建时间', dataIndex: 'created_at', key: 'created_at', width: 180 },
  { title: '操作', key: 'action', width: 200, fixed: 'right' as const },
];

const formState = reactive({
  title: '',
  subtitle: '',
  type: 'image_text',
  cover_url: '',
  content_url: '',
  description: '',
  author: '',
  category: '',
  duration: 0,
  word_count: 0,
  chapter_count: 0,
  status: 0,
});

const formRules = {
  title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
  type: [{ required: true, message: '请选择类型', trigger: 'change' }],
};

// 获取类型名称
const getTypeName = (type: string) => {
  const typeMap: Record<string, string> = {
    image_text: '图文',
    novel: '小说',
    video: '视频',
    banner: 'Banner',
    comic: '漫剧',
    short_drama: '短剧',
  };
  return typeMap[type] || type;
};

// 获取类型颜色
const getTypeColor = (type: string) => {
  const colorMap: Record<string, string> = {
    image_text: 'blue',
    novel: 'green',
    video: 'orange',
    banner: 'purple',
    comic: 'cyan',
    short_drama: 'magenta',
  };
  return colorMap[type] || 'default';
};

// 格式化数字
const formatNumber = (num: number) => {
  if (num >= 10000) {
    return (num / 10000).toFixed(1) + 'w';
  }
  return num.toString();
};

// 加载数据
const loadData = async () => {
  loading.value = true;
  try {
    const res = await getMaterialList({
      page: pagination.current,
      page_size: pagination.pageSize,
      type: filterType.value,
      keyword: searchKeyword.value,
    });
    tableData.value = res.list;
    pagination.total = res.total;
  } catch (error) {
    message.error('加载数据失败');
  } finally {
    loading.value = false;
  }
};

// 筛选
const handleFilter = () => {
  pagination.current = 1;
  loadData();
};

// 分页
const handlePageChange = (page: number, pageSize: number) => {
  pagination.current = page;
  pagination.pageSize = pageSize;
  loadData();
};

// 添加
const handleAdd = () => {
  modalTitle.value = '添加物料';
  isEdit.value = false;
  isView.value = false;
  currentId.value = 0;
  Object.assign(formState, {
    title: '',
    subtitle: '',
    type: 'image_text',
    cover_url: '',
    content_url: '',
    description: '',
    author: '',
    category: '',
    duration: 0,
    word_count: 0,
    chapter_count: 0,
    status: 0,
  });
  modalVisible.value = true;
};

// 编辑
const handleEdit = (row: Material) => {
  modalTitle.value = '编辑物料';
  isEdit.value = true;
  isView.value = false;
  currentId.value = row.id;
  Object.assign(formState, {
    title: row.title,
    subtitle: row.subtitle,
    type: row.type,
    cover_url: row.cover_url,
    content_url: row.content_url,
    description: row.description,
    author: row.author,
    category: row.category,
    duration: row.duration,
    word_count: row.word_count,
    chapter_count: row.chapter_count,
    status: row.status,
  });
  modalVisible.value = true;
};

// 查看
const handleView = (row: Material) => {
  modalTitle.value = '查看物料';
  isEdit.value = false;
  isView.value = true;
  currentId.value = row.id;
  Object.assign(formState, {
    title: row.title,
    subtitle: row.subtitle,
    type: row.type,
    cover_url: row.cover_url,
    content_url: row.content_url,
    description: row.description,
    author: row.author,
    category: row.category,
    duration: row.duration,
    word_count: row.word_count,
    chapter_count: row.chapter_count,
    status: row.status,
  });
  modalVisible.value = true;
};

// 删除
const handleDelete = async (row: Material) => {
  try {
    await deleteMaterial(row.id);
    message.success('删除成功');
    loadData();
  } catch (error) {
    message.error('删除失败');
  }
};

// 管理章节
const handleChapter = (row: Material) => {
  router.push({
    path: '/material/chapter',
    query: { material_id: row.id, title: row.title, type: row.type },
  });
};

// 提交
const handleModalOk = async () => {
  if (isView.value) {
    modalVisible.value = false;
    return;
  }

  try {
    await formRef.value?.validate();
    if (isEdit.value) {
      await updateMaterial(currentId.value, formState);
      message.success('更新成功');
    } else {
      await createMaterial(formState);
      message.success('创建成功');
    }
    modalVisible.value = false;
    loadData();
  } catch (error) {
    message.error('操作失败');
  }
};

// 取消
const handleModalCancel = () => {
  modalVisible.value = false;
  formRef.value?.resetFields();
};

onMounted(() => {
  loadData();
});
</script>

<style scoped>
.material-list {
  padding: 0;
}

.pagination {
  margin-top: 16px;
  text-align: right;
}
</style>
