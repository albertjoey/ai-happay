# 项目规范文档 (Project Specification)

> 本文档定义了项目的架构规范、技术栈约束和开发标准,确保AI助手能够一次性正确理解和实现需求。

---

## 1. 项目架构规范

### 1.1 后端架构 (Go-Zero)

#### 分层架构
```
Handler → Logic → Repository → Database
```

**严格规范**:
- ✅ **Handler层**: 只负责HTTP请求处理和响应
- ✅ **Logic层**: 业务逻辑处理,调用Repository
- ✅ **Repository层**: 数据库操作,所有SQL必须在这里
- ❌ **禁止**: Logic层直接写SQL

#### 文件结构
```
backend/app/channel/
├── internal/
│   ├── handler/          # HTTP处理器
│   │   └── discoverhandler.go
│   ├── logic/            # 业务逻辑
│   │   └── discoverlogic.go
│   ├── repository/       # 数据访问层
│   │   ├── discover.go        # 接口定义
│   │   └── discoverimpl.go    # 实现
│   ├── model/            # 数据模型
│   └── types/            # 类型定义
```

#### Repository规范
```go
// 1. 定义接口 (discover.go)
type DiscoverRepository interface {
    GetModuleList(ctx context.Context) ([]types.DiscoverConfig, error)
    CreateContent(ctx context.Context, req *types.DiscoverContentCreateRequest) error
}

// 2. 实现接口 (discoverimpl.go)
type discoverRepository struct {
    db *gorm.DB
}

func (r *discoverRepository) GetModuleList(ctx context.Context) ([]types.DiscoverConfig, error) {
    var modules []types.DiscoverConfig
    err := r.db.WithContext(ctx).Table("discover_module").Find(&modules).Error
    return modules, err
}

// 3. Logic层使用 (discoverlogic.go)
func (l *DiscoverLogic) Discover() (interface{}, error) {
    modules, err := l.repo.GetModuleList(l.ctx)
    if err != nil {
        return nil, err
    }
    // 业务逻辑处理
}
```

### 1.2 前端架构 (Vue 3 + Ant Design Vue)

#### 组件规范
- ✅ **使用 `<script setup>` 语法**
- ✅ **使用 Ant Design Vue 组件** (a-table, a-form, a-modal等)
- ❌ **禁止使用 vxe-table** (存在缓存问题)
- ✅ **路由切换必须使用 `:key="$route.name"`**

#### 页面结构模板
```vue
<template>
  <div class="page-container">
    <a-card title="页面标题">
      <template #extra>
        <a-space>
          <a-select v-model:value="selectedChannel" @change="loadData">
            <a-select-option v-for="ch in channels" :key="ch.id" :value="ch.id">
              {{ ch.name }}
            </a-select-option>
          </a-select>
          <a-button type="primary" @click="handleAdd">添加</a-button>
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
          <!-- 自定义列渲染 -->
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { message } from 'ant-design-vue';

// 状态定义
const loading = ref(false);
const tableData = ref([]);
const selectedChannel = ref(0);

// 表格列定义
const columns = [
  { title: 'ID', dataIndex: 'id', key: 'id', width: 80 },
  // ... 更多列
];

// 加载数据
const loadData = async () => {
  if (selectedChannel.value === 0) return;

  loading.value = true;
  try {
    const res = await getDataList({ channel_id: selectedChannel.value });
    tableData.value = res.list || [];
  } catch (error) {
    message.error('加载数据失败');
  } finally {
    loading.value = false;
  }
};

// 加载频道列表
const loadChannels = async () => {
  try {
    const res = await getChannelList({ page: 1, page_size: 100 });
    channels.value = res.list || [];
    if (channels.value.length > 0 && selectedChannel.value === 0) {
      selectedChannel.value = channels.value[0].id;
      loadData();
    }
  } catch (error) {
    message.error('加载频道列表失败');
  }
};

onMounted(() => {
  loadChannels();
});
</script>
```

#### 路由配置规范
```typescript
// router/index.ts
{
  path: 'ad-slot',
  name: 'AdSlot',  // 必须有唯一的name
  component: () => import('@/views/channel/AdSlotList.vue'),
  meta: { title: '广告位管理' },
}

// Layout.vue
<router-view :key="$route.name || $route.path" />
```

---

## 2. 数据库规范

### 2.1 字段命名规范

**重要**: 数据库字段与代码字段的映射关系

| 数据库字段 | 代码字段 | 说明 |
|----------|---------|------|
| `nickname` | `realname` | 管理员姓名 |
| `created_at` | `CreatedAt` | 创建时间 |
| `updated_at` | `UpdatedAt` | 更新时间 |

**Repository层必须处理字段映射**:
```go
// 错误示例
SELECT id, username, realname FROM admin_user

// 正确示例
SELECT id, username, nickname as realname FROM admin_user
```

### 2.2 字符集规范

- ✅ **数据库**: `utf8mb4_unicode_ci`
- ✅ **连接**: `charset=utf8mb4`
- ✅ **导入数据**: `SET NAMES utf8mb4`

### 2.3 空值处理

**前端必须处理null值**:
```vue
<!-- 错误示例 -->
<div>{{ record.insert_rule.fixed_position }}</div>

<!-- 正确示例 -->
<div v-if="record.insert_rule">
  {{ record.insert_rule.fixed_position }}
</div>
<div v-else>未配置</div>
```

---

## 3. API规范

### 3.1 RESTful API设计

```
GET    /api/v1/channel/list       # 获取列表
POST   /api/v1/channel            # 创建
PUT    /api/v1/channel/:id        # 更新
DELETE /api/v1/channel/:id        # 删除
```

### 3.2 响应格式

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [],
    "total": 100
  }
}
```

### 3.3 端口分配

| 服务 | 端口 | 说明 |
|-----|------|------|
| H5前端 | 4000 | Next.js |
| 管理后台 | 4002 | Vue 3 |
| 后端API | 4004 | Go-Zero |

---

## 4. 开发流程规范

### 4.1 新功能开发流程

1. **创建Repository层**
   - 定义接口 (xxx.go)
   - 实现接口 (xxximpl.go)

2. **创建Logic层**
   - 注入Repository依赖
   - 实现业务逻辑

3. **创建Handler层**
   - 注册路由
   - 处理HTTP请求

4. **前端开发**
   - 创建API接口
   - 创建页面组件
   - 注册路由

### 4.2 问题排查流程

1. **检查后端日志**
   ```bash
   tail -f /tmp/happy-channel.log
   ```

2. **测试API**
   ```bash
   curl -s "http://localhost:4004/api/v1/xxx" | python3 -m json.tool
   ```

3. **检查数据库**
   ```bash
   docker exec -it happy-mysql mysql -uroot -proot123456 happy
   ```

4. **检查浏览器控制台**
   - 打开开发者工具 (F12)
   - 查看Console和Network标签

---

## 5. 常见问题解决方案

### 5.1 路由缓存问题

**症状**: 切换路由后页面内容不更新

**解决方案**:
```vue
<!-- Layout.vue -->
<router-view :key="$route.name || $route.path" />
```

### 5.2 Loading状态卡住

**原因**:
1. API请求失败,loading未重置
2. 数据为null导致渲染错误

**解决方案**:
```typescript
const loadData = async () => {
  if (selectedChannel.value === 0) return;  // 检查前置条件

  loading.value = true;
  try {
    const res = await getDataList({ channel_id: selectedChannel.value });
    tableData.value = res.list || [];
  } catch (error) {
    message.error('加载数据失败');
  } finally {
    loading.value = false;  // 确保重置loading
  }
};
```

### 5.3 中文乱码

**解决方案**:
```sql
-- 导入数据前执行
SET NAMES utf8mb4;

-- 创建表时指定字符集
CREATE TABLE xxx (
  ...
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

---

## 6. 测试数据管理

### 6.1 导出数据

```bash
docker exec happy-mysql mysqldump -uroot -proot123456 \
  --default-character-set=utf8mb4 \
  --databases happy \
  --no-create-info \
  --complete-insert \
  --skip-add-locks \
  --skip-disable-keys \
  > backend/sql/init_data.sql
```

### 6.2 导入数据

```bash
docker exec -i happy-mysql mysql -uroot -proot123456 happy < backend/sql/init_data.sql
```

---

## 7. 如何向AI提问

### 7.1 提问模板

```markdown
## 需求描述
[清晰描述你要实现的功能]

## 技术要求
- 后端: 使用Repository模式,禁止Logic层直接写SQL
- 前端: 使用Ant Design Vue,禁止vxe-table
- 数据库: 字段映射关系 [具体说明]

## 相关文件
- 后端: [列出需要修改的文件]
- 前端: [列出需要修改的文件]

## 参考示例
[如果有类似功能,提供参考]
```

### 7.2 示例提问

**好的提问**:
```
需求: 完成发现页管理功能

技术要求:
1. 后端使用Repository模式:
   - 创建 discover.go (接口定义)
   - 创建 discoverimpl.go (实现)
   - 修改 discoverlogic.go (调用Repository)

2. 前端使用Ant Design Vue:
   - 创建 ConfigList.vue (模块配置)
   - 创建 ItemList.vue (内容管理)
   - 使用a-table,禁止vxe-table

3. 数据库字段映射:
   - discover_module表
   - discover_content表

参考: 可以参考 DiamondList.vue 的实现方式
```

**不好的提问**:
```
完成发现页管理
```

---

## 8. 项目目录结构

```
happytwo/
├── backend/
│   ├── app/
│   │   └── channel/
│   │       ├── internal/
│   │       │   ├── handler/
│   │       │   ├── logic/
│   │       │   ├── repository/
│   │       │   ├── model/
│   │       │   └── types/
│   │       └── channel.go
│   └── sql/
│       ├── init_data.sql
│       └── README.md
├── frontend/
│   ├── h5/              # Next.js H5前端
│   └── admin/           # Vue 3 管理后台
│       ├── src/
│       │   ├── api/
│       │   ├── views/
│       │   ├── router/
│       │   └── main.ts
│       └── package.json
├── docker-compose.yml
└── PROJECT_SPEC.md      # 本文档
```

---

## 9. 更新日志

| 日期 | 版本 | 更新内容 |
|-----|------|---------|
| 2026-03-16 | 1.0 | 初始版本,总结项目规范 |

---

## 10. 贡献者

- 项目负责人: [你的名字]
- AI助手: CodeArts代码智能体

---

**注意**: 本文档会随着项目发展持续更新,请定期同步最新版本。
