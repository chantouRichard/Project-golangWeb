### **API 接口文档**

------

#### **1. 用户管理**

##### **1.1. 用户注册**

- **HTTP 方法**: `POST`

- **URL**: `/api/register`

- 请求参数:

  - `username` (string) - 用户名
  - `email` (string) - 邮箱
  - `password` (string) - 密码

- 响应示例:

  ```
  {
    "status": "success",
    "message": "登录成功"
  }
  ```

##### **1.2. 用户登录**

- **HTTP 方法**: `POST`

- **URL**: `/api/login`

- 请求参数:

  - `email` (string) - 邮箱
  - `password` (string) - 密码

- 响应示例:

  ```
  {
    "status": "success",
    "token": "jwt-token-here",
    "user": {
      "id": 1,
      "username": "user1",
      "email": "user1@example.com"
    }
  }
  ```

##### **1.3. 获取用户信息**

- **HTTP 方法**: `GET`

- **URL**: `/api/user`

- 请求头:

  - `Authorization: Bearer jwt-token-here`

- 响应示例:

  ```
  {
    "status": "success",
    "user": {
      "id": 1,
      "username": "user1",
      "email": "user1@example.com"
    }
  }
  ```

##### **1.4. 更新用户信息**

- **HTTP 方法**: `PUT`

- **URL**: `/api/user`

- 请求头:

  - `Authorization: Bearer jwt-token-here`

- 请求参数:

  - `username` (string) - 新的用户名
  - `email` (string) - 新的邮箱

- 响应示例:

  ```
  {
    "status": "success",
    "message": "更新成功"
  }
  ```

------

#### **2. 电影管理**

##### **2.1. 获取电影列表**

- **HTTP 方法**: `GET`

- **URL**: `/api/movies`

- 响应示例:

  ```
  {
    "status": "success",
    "movies": [
      {
        "id": 1,
        "title": "Movie 1",
        "description": "This is a movie description",
        "thumbnail_url": "https://example.com/thumbnail1.jpg",
        "video_url": "https://example.com/video1.m3u8"
      },
      {
        "id": 2,
        "title": "Movie 2",
        "description": "This is another movie description",
        "thumbnail_url": "https://example.com/thumbnail2.jpg",
        "video_url": "https://example.com/video2.m3u8"
      }
    ]
  }
  ```

##### **2.2. 获取电影详情**

- **HTTP 方法**: `GET`

- **URL**: `/api/movies/{id}`

- 响应示例:

  ```
  {
    "status": "success",
    "movie": {
      "id": 1,
      "title": "Movie 1",
      "description": "This is a movie description",
      "thumbnail_url": "https://example.com/thumbnail1.jpg",
      "video_url": "https://example.com/video1.m3u8"
    }
  }
  ```

------

#### **3. 评论系统**

##### **3.1. 获取电影评论**

- **HTTP 方法**: `GET`

- **URL**: `/api/movies/{id}/comments`

- 响应示例:

  ```
  {
    "status": "success",
    "comments": [
      {
        "id": 1,
        "user_id": 1,
        "username": "user1",
        "content": "Great movie!",
        "created_at": "2024-08-15T12:34:56Z"
      },
      {
        "id": 2,
        "user_id": 2,
        "username": "user2",
        "content": "Not bad",
        "created_at": "2024-08-15T13:00:00Z"
      }
    ]
  }
  ```

##### **3.2. 发表评论**

- **HTTP 方法**: `POST`

- **URL**: `/api/movies/{id}/comments`

- 请求头:

  - `Authorization: Bearer jwt-token-here`

- 请求参数:

  - `content` (string) - 评论内容

- 响应示例:

  ```
  {
    "status": "success",
    "message": "评论成功"
  }
  ```

------

#### **4. 实时聊天**

##### **4.1. 加入聊天房间**

- **WebSocket 连接**: `ws://example.com/ws/chat`

- 消息格式:

  ```
  {
    "type": "join",
    "room_id": "movie-1",
    "username": "user1"
  }
  ```

##### **4.2. 发送聊天消息**

- WebSocket 消息:

  ```
  {
    "type": "message",
    "room_id": "movie-1",
    "username": "user1",
    "content": "This is an awesome scene!"
  }
  ```

##### **4.3. 接收聊天消息**

- WebSocket 消息:

  ```
  {
    "type": "message",
    "room_id": "movie-1",
    "username": "user2",
    "content": "Totally agree!"
  }
  ```

------

#### **5. 同步播放**

##### **5.1. 同步播放进度**

- **WebSocket 连接**: `ws://example.com/ws/sync`

- 消息格式:

  ```
  {
    "type": "sync",
    "room_id": "movie-1",
    "username": "user1",
    "current_time": 1234.56  // 当前播放时间（秒）
  }
  ```

##### **5.2. 接收同步进度**

- WebSocket 消息:

  ```
  {
    "type": "sync",
    "room_id": "movie-1",
    "username": "user2",
    "current_time": 1234.56
  }
  ```