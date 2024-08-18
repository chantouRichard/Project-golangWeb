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
    

    <el-empty
      v-else-if="!loading"
      description="No Room Info Available"
    ></el-empty>
    <el-spinner v-else></el-spinner>
    <label class="screen"></label>
    <label class="discuss"></label>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { ElMessage } from "element-plus";

const room = ref(null);
const loading = ref(false);

import { movieRoomInfoService } from "@/api/movie";
const getRoomInfo = async () => {
  loading.value = true;
  try {
    let result = await movieRoomInfoService();
    if (result && result.room) {
      room.value = result.room;
    } else {
      ElMessage.error("Failed to get room information");
    }
  } catch (error) {
    ElMessage.error("An error occurred while fetching room information");
  } finally {
    loading.value = false;
  }
};

// 退出房间
import { exitRoomService } from "@/api/movie";
const leaveRoom = async () => {
  loading.value = true;
  try {
    let result = await exitRoomService();
    ElMessage.success("Successfully left room");
    room.value = null;
  } catch (error) {
    ElMessage.error("An error occurred while leaving room");
  } finally {
    loading.value = false;
  }
  
};

getRoomInfo();
</script>

<style scoped>
.el-menu-vertical {
  width: 300px;
  background-color: #f5f5f5;
  border-radius: 8px;
  padding: 10px;
}

.screen {
  background-color: black;
  margin: 10px;
  width: 1000px;
  height: 800px;
}

.discuss {
  margin: 10px;
}
</style>
