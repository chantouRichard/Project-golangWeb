<template>
  <div>
    <el-menu v-if="room" :default-active="room.RoomID" class="el-menu-vertical">
      <el-menu-item>
        <h3><strong>Room ID:</strong> {{ room.RoomID }}</h3>
      </el-menu-item>
      <el-menu-item>
        <h3><strong>Room Name:</strong> {{ room.Room.RoomName }}</h3>
      </el-menu-item>
      <el-menu-item v-if="isCreator === true">
        <h3><strong>身份：</strong>房主</h3>
      </el-menu-item>
      <el-menu-item v-else>
        <h3><strong>身份：</strong>观众</h3>
      </el-menu-item>

      <!-- 在此处添加更多房间信息 -->
      <el-button type="primary" v-if="!isCreator" @click="leaveRoom"
        >退出房间</el-button
      >
      <el-button type="primary" v-if="isCreator" @click="closeRoom"
        >关闭房间</el-button
      >
    </el-menu>

    <!-- 视频播放区域 -->
    <div v-if="room && videoUrl" class="video-container">
      <video
        ref="videoElementRef"
        controls
        autoplay
        :src="videoUrl"
        class="video-player"
      ></video>
    </div>
    <!-- 聊天框部分 -->
    <div v-if="room" class="chat-container">
      <div class="chat-messages" ref="chatMessages">
        <p v-for="(message, index) in messages" :key="index">
          {{ message }}
        </p>
      </div>
      <el-input
        v-model="newMessage"
        placeholder="输入消息..."
        @keyup.enter="sendMessage"
        class="chat-input"
      />
      <el-button type="primary" @click="sendMessage">发送</el-button>
    </div>

    <el-empty
      v-else-if="!loading && !room"
      description="No Room Info Available"
    ></el-empty>
    <el-spinner v-else></el-spinner>
  </div>
</template>

<script setup>
import { onMounted, onBeforeUnmount } from "vue";

const userName = ref("");
const room = ref(null);
const loading = ref(false);
const videoUrl = ref("");
const socket = ref(null);
const messages = ref([]);
const newMessage = ref("");
const isCreator = ref(false);
const videoElementRef = ref(null);

// 获取房间信息
import { movieRoomInfoService } from "@/api/movie";
const getRoomInfo = async () => {
  loading.value = true;
  try {
    let result = await movieRoomInfoService();
    room.value = result.room;
    if (room.value && room.value.RoomID) {
      // 设置视频源URL
      videoUrl.value = `http://localhost:8080/api/rooms/${room.value.RoomID}/stream`;

      if (room.value.Room.CreatorID === room.value.User.id)
        isCreator.value = true;
      // 连接WebSocket
      connectWebSocket(room.value.RoomID);
      ElMessage.success("Connected to room");
    }
  } catch (error) {
    ElMessage.error("An error occurred while fetching room information");
  } finally {
    loading.value = false;
  }
};
//同步播放进度
let progressInterval = null;

const sendVideoProgress = () => {
  if (socket.value && socket.value.readyState === WebSocket.OPEN) {
    const progressData = {
      type: "video-progress",
      progress: getVideoProgress(), // 获取当前视频播放进度
      timestamp: new Date().toISOString(),
    };

    // 发送视频进度信息给房间中的其他用户
    socket.value.send(JSON.stringify(progressData));
  }
};

const startSendingProgress = () => {
  // 如果是房主，定期发送视频进度
  progressInterval = setInterval(sendVideoProgress, 10000); // 每秒发送一次
  ElMessage.success("Sending video progress:" + getVideoProgress());
};

const stopSendingProgress = () => {
  if (progressInterval) {
    clearInterval(progressInterval);
    progressInterval = null;
  }
};

// 连接WebSocket
const connectWebSocket = (roomID) => {
  if (socket.value) {
    socket.value.close(); // 关闭现有连接
  }

  socket.value = new WebSocket(`ws://localhost:8080/api/rooms/${roomID}/chat`);

  socket.value.onopen = () => {
    console.log("Connected to chat room", roomID);
    if (isCreator.value) startSendingProgress(); // 如果是房主，开始发送视频进度
  };

  socket.value.onmessage = (event) => {
    // 假设接收到的消息是 JSON 格式
    const data = JSON.parse(event.data);
    // ElMessage.success(data.type)

    // 检查消息类型并根据类型处理
    switch (data.type) {
      case "chat":
        messages.value.push(data.user + ":" + data.message);
        scrollToBottom();
        break;
      case "video-progress": {
        // ElMessage.success("isCreator::"+isCreator.value);
        if (!isCreator.value) {
          // 只有非房主同步视频进度
          // ElMessage.success("2");
          syncVideoProgress(data.progress);
          ElMessage.success("Received video progress:" + data.progress);
          const progress = JSON.parse(data.progress);
          const videoElement = videoElementRef.value;
          videoElement.currentTime = progress;

          // if (videoElement.paused) {
          //   videoElement.play();
          // }
        }
        break;
      }
      case "notification":
        showNotification(data.message);
        break;
      default:
        console.warn("Unknown message type:", data.type);
    }
  };

  socket.value.onclose = () => {
    console.log("Disconnected from chat room", roomID);
    stopSendingProgress(); // 停止发送视频进度
    socket.value = null; // 清除连接引用
  };

  socket.value.onerror = (error) => {
    console.error("WebSocket error:", error);
  };
  if (isCreator.value) {
    const videoElement = videoElementRef.value;
    videoElement.play(); // 房主自动播放视频

    setInterval(() => {
      sendVideoProgress(); // 每隔一段时间发送视频进度
      // const currentTime = Math.floor(videoElement.currentTime);
      // socket.value.send(JSON.stringify(currentTime));
    }, 10000); // 每秒发送一次视频进度
  }
};
//辅助函数
const getVideoProgress = () => {
  // 返回当前视频的播放进度，例如：
  // 如果视频播放长度为100秒，当前播放到50秒，则返回0.5
  const videoElement = videoElementRef.value;
  if (videoElement && videoElement.duration > 0)
    return videoElement.currentTime;
  return 0;
};

const syncVideoProgress = (progress) => {
  // 同步视频进度，例如：
  const videoElement = videoElementRef.value;
  if (videoElement && videoElement.duration > 0 && !isCreator)
    videoElement.currentTime = progress * videoElement.duration;
};

// 发送消息
const sendMessage = () => {
  if (
    newMessage.value.trim() !== "" &&
    socket.value &&
    socket.value.readyState === WebSocket.OPEN
  ) {
    // 构造消息对象
    const messageData = {
      type: "chat", // 消息类型
      user: userName.value, // 发送消息的用户名
      message: newMessage.value, // 消息内容
      timestamp: new Date().toISOString(), // 时间戳
    };

    // 将消息对象转换为 JSON 字符串
    const jsonMessage = JSON.stringify(messageData);

    // 发送 JSON 格式的消息到 WebSocket
    socket.value.send(jsonMessage);

    ElMessage.success("Sent message to room: " + newMessage.value);

    // 清空输入框
    newMessage.value = "";

    // 滚动到聊天框底部
    scrollToBottom();
  } else {
    ElMessage.error("Cannot send message. WebSocket is not connected.");
  }
};

// 滚动到聊天框底部
const scrollToBottom = () => {
  const chatMessages = document.querySelector(".chat-messages");
  chatMessages.scrollTop = chatMessages.scrollHeight;
};

// 组件挂载时获取历史观看记录
// onMounted(() => {
//   getRoomInfo();
// });

// 组件卸载时关闭WebSocket连接
onBeforeUnmount(() => {
  if (socket.value) {
    socket.value.close(); // 正确关闭 WebSocket 连接
  }
});
// Watch for changes in room and videoUrl
import { watch } from "vue";
watch([room, videoUrl], ([newRoom, newVideoUrl]) => {
  if (newRoom && newVideoUrl) {
    const videoElement = videoElementRef.value;
    if (videoElement) {
      console.log("Video URL updated:", newVideoUrl);
    }
  }
});

onMounted(() => {
  // Initial setup after component is mounted
  const videoElement = videoElementRef.value;
  if (videoElement) {
    console.log("Video element is ready:", videoElement);
  }
});
import { ref } from "vue";
import { ElMessage } from "element-plus";

// 退出房间
import { exitRoomService } from "@/api/movie";
const leaveRoom = async () => {
  loading.value = true;
  try {
    await exitRoomService();
    ElMessage.success("Successfully left room");
    room.value = null;
    videoUrl.value = ""; // 清空视频URL
  } catch (error) {
    ElMessage.error("An error occurred while leaving room");
  } finally {
    loading.value = false;
  }
};

//获取用户信息
import { userInfoService } from "@/api/user";
const getUserInfo = async () => {
  loading.value = true;
  try {
    let result = await userInfoService();
    userName.value = result.user.username;
  } catch (error) {
    ElMessage.error("An error occurred while fetching user information");
  }
};
//关闭房间
import { closeRoomService } from "@/api/movie";
import { ElMessageBox } from "element-plus";
const closeRoom = async () => {
  loading.value = true;
  //提示是否删除

  ElMessageBox.confirm("是否删除该分类？", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      //调用删除接口函数
      let result = await closeRoomService(room.value.RoomID);
      room.value = null;
      videoUrl.value = ""; // 清空视频URL
      getRoomInfo();
      ElMessage({
        type: "success",
        message: "关闭房间成功！",
      });
      //调用获取所有文章分类的函数
      articleCategoryList();
    })
    .catch(() => {
      ElMessage({
        type: "info",
        message: "已取消关闭房间！",
      });
    });
};
getRoomInfo();
getUserInfo();
</script>

<style scoped>
.el-menu-vertical {
  width: 300px;
  background-color: #f5f5f5;
  border-radius: 8px;
  padding: 10px;
}

.video-container {
  margin-top: 20px;
}

.video-player {
  width: 100%;
  max-width: 1000px;
  height: auto;
}

.screen {
  background-color: black;
  margin: 10px;
}

.discuss {
  margin: 10px;
}

.chat-container {
  margin-top: 20px;
  background-color: #000000;
  padding: 10px;
  border-radius: 8px;
  width: 100%;
  max-width: 500px;
}

.chat-messages {
  height: 200px;
  overflow-y: auto;
  border: 1px solid #000000;
  padding: 10px;
  margin-bottom: 10px;
  font-size: large;
  color: black;
  background-color: #e6bd57;
}

.chat-input {
  width: calc(100% - 100px);
  margin-right: 10px;
}
</style>
