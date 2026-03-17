<template>
  <div class="chapter-list">
    <a-card :title="`${materialTitle} - ${chapterTypeName}管理`">
      <template #extra>
        <a-space>
          <a-button @click="handleBack">返回列表</a-button>
          <a-button type="primary" @click="handleAdd">
            <template #icon><PlusOutlined /></template>
            添加{{ chapterTypeName }}
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
    </a-card>

    <!-- 添加/编辑对话框 -->
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
          <a-input v-model:value="formState.title" placeholder="请输入标题" />
        </a-form-item>

        <!-- 小说：文本内容 -->
        <template v-if="materialType === 'novel'">
          <a-form-item label="章节内容" name="content">
            <a-textarea v-model:value="formState.content" placeholder="请输入章节内容" :rows="10" />
          </a-form-item>
        </template>

        <!-- 漫剧：图片列表 -->
        <template v-if="materialType === 'drama'">
          <a-form-item label="图片列表" name="images">
            <div v-for="(img, index) in formState.images" :key="index" style="margin-bottom: 8px;">
              <a-input-group compact>
                <a-input v-model:value="formState.images[index]" style="width: calc(100% - 80px)" placeholder="图片URL" />
                <a-button danger @click="removeImage(index)">删除</a-button>
              </a-input-group>
            </div>
            <a-button type="dashed" block @click="addImage">添加图片</a-button>
          </a-form-item>
        </template>

        <!-- 短剧/长视频：视频URL -->
        <template v-if="materialType === 'short_drama' || materialType === 'long_video'">
          <a-form-item label="视频URL" name="video_url">
            <a-input v-model:value="formState.video_url" placeholder="请输入视频URL" />
          </a-form-item>
          <a-form-item label="时长(秒)" name="duration">
            <a-input-number v-model:value="formState.duration" :min="0" style="width: 200px" />
          </a-form-item>
        </template>

        <a-form-item label="排序" name="sort">
          <a-input-number v-model:value="formState.sort" :min="0" placeholder="留空自动递增" />
        </a-form-item>
        <a-form-item label="是否免费" name="is_free">
          <a-radio-group v-model:value="formState.is_free">
            <a-radio :value="1">免费</a-radio>
            <a-radio :value="0">付费</a-radio>
          </a-radio-group>
        </a-form-item>
        <a-form-item v-if="formState.is_free === 0" label="价格" name="price">
          <a-input-number v-model:value="formState.price" :min="0" addon-after="币" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { message } from 'ant-design-vue';
import { PlusOutlined } from '@ant-design/icons-vue';
import type { FormInstance } from 'ant-design-vue';
import {
  getChapterList,
  createChapter,
  updateChapter,
  deleteChapter,
  type Chapter,
} from '@/api/chapter';

const route = useRoute();
const router = useRouter();
const materialId = Number(route.query.material_id);
const materialType = ref(route.query.type || 'novel');
const materialTitle = ref(route.query.title || '物料');

const loading = ref(false);
const tableData = ref<Chapter[]>([]);

const modalVisible = ref(false);
const modalTitle = ref('');
const currentId = ref<number>(0);
const formRef = ref<FormInstance>();

// 根据物料类型显示不同名称
const chapterTypeName = computed(() => {
  const typeNames: Record<string, string> = {
    novel: '章节',
    drama: '话',
    short_drama: '集',
    long_video: '集',
  };
  return typeNames[materialType.value] || '章节';
});

// 表格列定义
const columns = [
  { title: 'ID', dataIndex: 'id', key: 'id', width: 80 },
  { title: '章节名称', dataIndex: 'name', key: 'name', width: 200 },
  { title: '排序', dataIndex: 'sort', key: 'sort', width: 80 },
  { title: '状态', dataIndex: 'status', key: 'status', width: 100 },
  { title: '操作', key: 'action', width: 150, fixed: 'right' as const },
];

const formState = reactive({
  title: '',
  content: '',
  images: [] as string[],
  video_url: '',
  duration: 0,
  sort: 0,
  is_free: 1,
  price: 0,
});

const formRules = {
  title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
};

// 格式化数字
const formatNumber = (num: number) => {
  if (num >= 10000) {
    return (num / 10000).toFixed(1) + 'w';
  }
  return num.toString();
};

// 格式化时长
const formatDuration = (seconds: number) => {
  if (!seconds) return '-';
  const min = Math.floor(seconds / 60);
  const sec = seconds % 60;
  return min > 0 ? `${min}分${sec}秒` : `${sec}秒`;
};

// 添加图片
const addImage = () => {
  formState.images.push('');
};

// 删除图片
const removeImage = (index: number) => {
  formState.images.splice(index, 1);
};

// 加载数据
const loadData = async () => {
  loading.value = true;
  try {
    const res = await getChapterList(materialId);
    tableData.value = res.list || [];
  } catch (error) {
    message.error('加载数据失败');
  } finally {
    loading.value = false;
  }
};

// 返回
const handleBack = () => {
  router.push('/material');
};

// 添加
const handleAdd = () => {
  modalTitle.value = `添加${chapterTypeName.value}`;
  currentId.value = 0;
  Object.assign(formState, {
    title: '',
    content: '',
    images: [],
    video_url: '',
    duration: 0,
    sort: 0,
    is_free: 1,
    price: 0,
  });
  modalVisible.value = true;
};

// 编辑
const handleEdit = (row: Chapter) => {
  modalTitle.value = `编辑${chapterTypeName.value}`;
  currentId.value = row.id;
  Object.assign(formState, {
    title: row.title,
    content: row.content || '',
    images: row.images || [],
    video_url: row.video_url || '',
    duration: row.duration || 0,
    sort: row.sort,
    is_free: row.is_free,
    price: row.price,
  });
  modalVisible.value = true;
};

// 删除
const handleDelete = async (row: Chapter) => {
  try {
    await deleteChapter(row.id);
    message.success('删除成功');
    loadData();
  } catch (error) {
    message.error('删除失败');
  }
};

// 提交
const handleModalOk = async () => {
  try {
    await formRef.value?.validate();
    
    const chapterType = materialType.value === 'novel' ? 'text' : 
                        materialType.value === 'drama' ? 'image' : 'video';
    
    if (currentId.value) {
      await updateChapter(currentId.value, {
        chapter_type: chapterType,
        ...formState,
      });
      message.success('更新成功');
    } else {
      await createChapter({
        material_id: materialId,
        chapter_type: chapterType,
        ...formState,
      });
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
.chapter-list {
  padding: 0;
}
</style>
