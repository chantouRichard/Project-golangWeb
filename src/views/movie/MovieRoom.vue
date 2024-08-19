<template>
  <div>
    <el-menu v-if="room" :default-active="room.RoomID" class="el-menu-vertical">
      <el-menu-item>
        <h3><strong>Room ID:</strong> {{ room.RoomID }}</h3>
      </el-menu-item>
      <el-menu-item>
        <h3><strong>Room Name:</strong> {{ room.Room.RoomName }}</h3>
      </el-menu-item>

      <!-- 在此处添加更多房间信息 -->
      <el-button type="primary" @click="leaveRoom">退出房间</el-button>
    </el-menu>

    <!-- 视频播放区域 -->
    <div v-if="room && videoUrl" class="video-container">
      <video controls autoplay :src="videoUrl" class="video-player"></video>
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

// 连接WebSocket
const connectWebSocket = (roomID) => {
  if (socket.value) {
    socket.value.close(); // 关闭现有连接
  }

  socket.value = new WebSocket(`ws://localhost:8080/api/rooms/${roomID}/chat`);

  socket.value.onopen = () => {
    console.log("Connected to chat room", roomID);
  };

  socket.value.onmessage = (event) => {
    messages.value.push(event.data);
    scrollToBottom();
  };

  socket.value.onclose = () => {
    console.log("Disconnected from chat room", roomID);
    socket.value = null; // 清除连接引用
  };

  socket.value.onerror = (error) => {
    console.error("WebSocket error:", error);
  };
};

// 发送消息
const sendMessage = () => {
  if (
    newMessage.value.trim() !== "" &&
    socket.value &&
    socket.value.readyState === WebSocket.OPEN
  ) {
    // 发送消息到 WebSocket
    socket.value.send(userName.value+":"+newMessage.value);
    ElMessage.success("Sent message to room:" + newMessage.value);
    // 将消息推送到 messages 数组，以便显示在聊天框中
    // messages.value.push(newMessage.value);

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
