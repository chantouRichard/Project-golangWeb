<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";
import { ElMessage } from "element-plus";

const collections = ref([]);
const selectedMovie = ref(null);
const rooms = ref([]);
const showMovieDetails = ref(false);

import { fetchCollectionsService } from "@/api/user";
const fetchCollections = async () => {
  let response = await fetchCollectionsService();

  if (response.status === "success") {
    collections.value = response.collections;
  } else {
    ElMessage.error("无法获取收藏数据");
  }
};

import { movieDetailService } from "@/api/movie";
const fetchMovieDetails = async (movieId) => {
  const response = await movieDetailService(movieId);
  if (response.status === "success") {
    selectedMovie.value = response.movie;
  } else {
    ElMessage.error("无法获取电影详情");
    selectedMovie.value = null;
  }
};

import { movieRoomListService } from "@/api/movie";
const fetchRoomDetails = async (movieId) => {
  try {
    let response = await movieRoomListService(movieId);
    if (response.status === "success") {
      rooms.value = response.rooms;
    } else {
      ElMessage.error("无法获取房间列表");
    }
  } catch (error) {
    ElMessage.error("无法获取房间列表");
  }
};

//进入房间
import { enterRoomService } from "@/api/movie";
const openRoom = async (roomId) => {
  let response = await enterRoomService(roomId);
  if (response.status === "success") {
    ElMessage.success("进入房间成功");
  } else {
    ElMessage.error("进入房间失败");
  }
};
const openMovieDrawer = async (movieId) => {
  await fetchMovieDetails(movieId);
  await fetchRoomDetails(movieId);
  showMovieDetails.value = true;
};

const formatDate = (dateString) => {
  const date = new Date(dateString);
  return date.toLocaleString();
};

//取消收藏
import { deleteCollectionService } from "@/api/user";
const deCollection = async (id) => {
  let response = await deleteCollectionService(id);
  if (response.status === "success") {
    ElMessage.success("取消收藏成功");
    fetchCollections();
  }
};

onMounted(() => {
  fetchCollections();
});
</script>

<template>
  <el-container>
    <el-header>
      <h2>我的收藏</h2>
    </el-header>

    <el-main v-if="collections.length > 0">
      <el-row :gutter="20">
        <el-col
          :span="8"
          v-for="(collection, index) in collections"
          :key="index"
        >
          <el-card
            :body-style="{ padding: '20px' }"
            @click="openMovieDrawer(collection.Movie.ID)"
          >
            <div>
              <h3>{{ collection.Movie.title }}</h3>
              <p>导演: {{ collection.Movie.director }}</p>
              <p>上映时间: {{ collection.Movie.CreatedAt }}</p>
              <p>{{ collection.Movie.description }}</p>
            </div>
            <el-button type="primary" @click="deCollection(collection.ID)">取消收藏</el-button>
          </el-card>
        </el-col>
      </el-row>
      <el-drawer
        v-model="showMovieDetails"
        title="电影详情"
        direction="rtl"
        size="50%"
      >
        <div v-if="selectedMovie">
          <h2>{{ selectedMovie.title }}</h2>
          <h3>导演: {{ selectedMovie.director }}</h3>
          <h3>上映时间: {{ formatDate(selectedMovie.updatedAt) }}</h3>
          <h3>描述: {{ selectedMovie.description }}</h3>

          <!-- Room List -->
          <h3>相关房间:</h3>
          <ul v-if="rooms">
            <li v-for="(room, index) in rooms" :key="index">
              <h3>
                房间名: {{ room.room_name }} - 创建者: {{ room.creator.name }} -
              </h3>
              <h3>创建时间: {{ formatDate(room.created_at) }}</h3>
              <el-button type="primary" @click="openRoom(room.id)"
                >进入房间</el-button
              >
            </li>
          </ul>
          <el-empty v-else description="暂无相关房间"></el-empty>
        </div>
        <el-empty v-else description="暂无电影详情"></el-empty>
      </el-drawer>
    </el-main>

    <el-main v-else>
      <el-empty description="您还没有收藏任何电影"></el-empty>
    </el-main>
  </el-container>
</template>

<style scoped>
h2 {
  margin: 0;
  padding: 10px 0;
  color: #409eff;
  font-weight: bold;
}

h3 {
  margin: 0 0 10px 0;
  font-size: 18px;
  color: #333;
}

.el-card {
  border-radius: 8px;
}

.el-empty {
  margin-top: 50px;
  text-align: center;
}
</style>
