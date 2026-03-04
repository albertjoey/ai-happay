# Happy API 接口文档

## 基础信息

- **Base URL**: `http://localhost:8001`
- **认证方式**: JWT Bearer Token
- **数据格式**: JSON
- **字符编码**: UTF-8

## 通用响应格式

### 成功响应
```json
{
  "code": 0,
  "message": "成功",
  "data": {}
}
```

### 错误响应
```json
{
  "code": 10001,
  "message": "参数错误",
  "data": null
}
```

## 用户服务 API

### 1. 用户注册

**接口**: `POST /api/v1/user/register`

**请求参数**:
```json
{
  "username": "test",
  "email": "test@example.com",
  "phone": "13800138000",
  "password": "password123",
  "nickname": "测试用户"
}
```

**响应示例**:
```json
{
  "code": 0,
  "message": "成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "test",
      "email": "test@example.com",
      "phone": "13800138000",
      "nickname": "测试用户",
      "avatar": "",
      "gender": 0,
      "bio": "",
      "role": 0,
      "follow_count": 0,
      "fans_count": 0,
      "like_count": 0
    }
  }
}
```

### 2. 用户登录

**接口**: `POST /api/v1/user/login`

**请求参数**:
```json
{
  "username": "test",
  "password": "password123"
}
```

**响应示例**:
```json
{
  "code": 0,
  "message": "成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "test",
      "email": "test@example.com",
      "nickname": "测试用户",
      "avatar": "",
      "role": 0
    }
  }
}
```

### 3. 获取用户信息

**接口**: `GET /api/v1/user/info`

**请求头**:
```
Authorization: Bearer {token}
```

**响应示例**:
```json
{
  "code": 0,
  "message": "成功",
  "data": {
    "id": 1,
    "username": "test",
    "email": "test@example.com",
    "nickname": "测试用户",
    "avatar": "https://example.com/avatar.jpg",
    "gender": 1,
    "bio": "这是我的个人简介",
    "role": 0,
    "follow_count": 10,
    "fans_count": 100,
    "like_count": 500
  }
}
```

### 4. 更新用户信息

**接口**: `PUT /api/v1/user/info`

**请求头**:
```
Authorization: Bearer {token}
```

**请求参数**:
```json
{
  "nickname": "新昵称",
  "avatar": "https://example.com/new-avatar.jpg",
  "gender": 1,
  "bio": "更新后的个人简介"
}
```

**响应示例**:
```json
{
  "code": 0,
  "message": "成功",
  "data": null
}
```

### 5. 关注用户

**接口**: `POST /api/v1/user/follow`

**请求头**:
```
Authorization: Bearer {token}
```

**请求参数**:
```json
{
  "user_id": 2
}
```

**响应示例**:
```json
{
  "code": 0,
  "message": "成功",
  "data": null
}
```

### 6. 获取用户列表

**接口**: `GET /api/v1/user/list`

**请求参数**:
- `page`: 页码，默认1
- `page_size`: 每页数量，默认20
- `keyword`: 搜索关键词
- `role`: 用户角色

**响应示例**:
```json
{
  "code": 0,
  "message": "成功",
  "data": {
    "total": 100,
    "list": [
      {
        "id": 1,
        "username": "user001",
        "nickname": "用户001",
        "avatar": "https://example.com/avatar1.jpg",
        "role": 0,
        "follow_count": 10,
        "fans_count": 100
      }
    ]
  }
}
```

## 内容服务 API

### 1. 创建内容

**接口**: `POST /api/v1/content`

**请求头**:
```
Authorization: Bearer {token}
```

**请求参数**:
```json
{
  "title": "精彩短视频分享",
  "description": "这是一个很棒的视频",
  "cover": "https://example.com/cover.jpg",
  "type": 2,
  "media": [
    {
      "type": 2,
      "url": "https://example.com/video.mp4",
      "thumbnail": "https://example.com/thumb.jpg",
      "duration": 60,
      "width": 1920,
      "height": 1080,
      "size": 10485760
    }
  ],
  "topics": [1, 2],
  "tags": [1, 2, 3]
}
```

**内容类型说明**:
- 1: 长视频
- 2: 短视频
- 3: 短剧
- 4: 漫剧
- 5: 小说
- 6: 图文

**响应示例**:
```json
{
  "code": 0,
  "message": "成功",
  "data": {
    "id": 1
  }
}
```

### 2. 获取内容列表

**接口**: `GET /api/v1/content/list`

**请求参数**:
- `page`: 页码，默认1
- `page_size`: 每页数量，默认20
- `type`: 内容类型
- `channel`: 频道ID
- `user_id`: 用户ID
- `status`: 状态

**响应示例**:
```json
{
  "code": 0,
  "message": "成功",
  "data": {
    "total": 100,
    "list": [
      {
        "id": 1,
        "title": "精彩短视频分享",
        "description": "这是一个很棒的视频",
        "cover": "https://example.com/cover.jpg",
        "type": 2,
        "status": 1,
        "view_count": 1234,
        "like_count": 567,
        "comment_count": 89,
        "collect_count": 45,
        "share_count": 23,
        "media": [
          {
            "type": 2,
            "url": "https://example.com/video.mp4",
            "thumbnail": "https://example.com/thumb.jpg",
            "duration": 60
          }
        ],
        "topics": [
          {"id": 1, "name": "热门话题"}
        ],
        "tags": [
          {"id": 1, "name": "搞笑"}
        ],
        "author": {
          "id": 1,
          "nickname": "创作者A",
          "avatar": "https://example.com/avatar.jpg"
        },
        "created_at": "2024-01-01 10:00:00"
      }
    ]
  }
}
```

### 3. 获取内容详情

**接口**: `GET /api/v1/content/{id}`

**响应示例**:
```json
{
  "code": 0,
  "message": "成功",
  "data": {
    "id": 1,
    "title": "精彩短视频分享",
    "description": "这是一个很棒的视频",
    "cover": "https://example.com/cover.jpg",
    "type": 2,
    "status": 1,
    "view_count": 1234,
    "like_count": 567,
    "comment_count": 89,
    "collect_count": 45,
    "share_count": 23,
    "media": [
      {
        "type": 2,
        "url": "https://example.com/video.mp4",
        "thumbnail": "https://example.com/thumb.jpg",
        "duration": 60,
        "width": 1920,
        "height": 1080,
        "size": 10485760
      }
    ],
    "topics": [
      {"id": 1, "name": "热门话题"}
    ],
    "tags": [
      {"id": 1, "name": "搞笑"}
    ],
    "author": {
      "id": 1,
      "nickname": "创作者A",
      "avatar": "https://example.com/avatar.jpg"
    },
    "created_at": "2024-01-01 10:00:00"
  }
}
```

### 4. 更新内容

**接口**: `PUT /api/v1/content/{id}`

**请求头**:
```
Authorization: Bearer {token}
```

**请求参数**:
```json
{
  "title": "更新后的标题",
  "description": "更新后的描述",
  "cover": "https://example.com/new-cover.jpg"
}
```

**响应示例**:
```json
{
  "code": 0,
  "message": "成功",
  "data": null
}
```

### 5. 删除内容

**接口**: `DELETE /api/v1/content/{id}`

**请求头**:
```
Authorization: Bearer {token}
```

**响应示例**:
```json
{
  "code": 0,
  "message": "成功",
  "data": null
}
```

## 频道服务 API

### 1. 获取频道列表

**接口**: `GET /api/v1/channel/list`

**响应示例**:
```json
{
  "code": 0,
  "message": "成功",
  "data": {
    "list": [
      {
        "id": 1,
        "name": "推荐",
        "code": "recommend",
        "description": "推荐内容",
        "icon": "https://example.com/icon.png"
      }
    ]
  }
}
```

### 2. 获取Banner列表

**接口**: `GET /api/v1/banner/list`

**请求参数**:
- `channel_id`: 频道ID

**响应示例**:
```json
{
  "code": 0,
  "message": "成功",
  "data": {
    "list": [
      {
        "id": 1,
        "title": "活动Banner",
        "image": "https://example.com/banner.jpg",
        "link": "https://example.com/activity"
      }
    ]
  }
}
```

## 错误码说明

| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 10000 | 服务器内部错误 |
| 10001 | 参数错误 |
| 10002 | 未授权 |
| 10003 | 无权限 |
| 10004 | 资源不存在 |
| 10005 | 请求错误 |
| 20001 | 用户不存在 |
| 20002 | 用户已存在 |
| 20003 | 密码错误 |
| 20004 | 用户已被禁用 |
| 30001 | 内容不存在 |
| 30002 | 内容已存在 |
| 30003 | 内容状态错误 |
| 30004 | 无权限操作此内容 |

## 注意事项

1. 所有需要认证的接口都需要在请求头中携带 `Authorization: Bearer {token}`
2. Token有效期为24小时，过期后需要重新登录
3. 所有时间字段格式为 `YYYY-MM-DD HH:mm:ss`
4. 分页参数从1开始
5. 文件上传接口需要使用 `multipart/form-data` 格式
