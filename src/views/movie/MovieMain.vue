<script setup>
import { ref } from "vue";
import { Axios } from "axios";

const images = ref(["/images/1.jpg", "/images/2.jpg", "/images/3.jpg"]);
const carousel = ref(null);
const recommends = ref([]);

const selectedMovie = ref(null);
const rooms = ref([]);
const showMovieDetails = ref(false);
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

const goToSlide = (index) => {
  carousel.value.setActiveItem(index + 1);
};

const handleChange = (index) => {
  // console.log("Current slide index:", index);
};

//电影推荐
import { movieRecommendService } from "@/api/movie";
import { ElMessage } from "element-plus"; // 获取推荐电影
const recommendMovie = async () => {
  let result = await movieRecommendService();
  console.log(result);
  recommends.value = result.movies;
  ElMessage.success("获取推荐电影成功"+result.movies);
  // recommends.value = result.data.movies;
};

recommendMovie();
</script>

<template>
  <div class="shell">
    <el-carousel
      :interval="3000"
      arrow="always"
      type="card"
      height="500px"
      @change="handleChange"
    >
      <el-carousel-item v-for="(image, index) in images" :key="index">
        <div
          :style="{
            backgroundImage: 'url(' + image + ')',
            height: '100%',
            backgroundSize: 'cover',
          }"
        ></div>
      </el-carousel-item>
    </el-carousel>
    <div class="min-images">
      <div
        v-for="(image, index) in images"
        :key="index"
        :style="{ backgroundImage: 'url(' + image + ')' }"
        class="min"
        @click="goToSlide(index)"
      ></div>
    </div>
  </div>
  <div class="homepage">
    <section class="hero">
      <div class="hero-content">
        <h1>Welcome to Cinema</h1>
        <p>Watch your favorite movies and shows in high quality</p>
        <button class="cta-button">Explore Now</button>
      </div>
    </section>

    <section class="movies-section">
      <h2>Featured Movies</h2>
      <div class="movies-grid">
        <div class="movie-card" v-for="movie in movies" :key="movie.id">
          <img :src="movie.image" :alt="movie.title" />
          <div class="movie-info">
            <h3>{{ movie.title }}</h3>
            <p>{{ movie.genre }}</p>
          </div>
        </div>
      </div>
      <el-main v-if="recommends.length > 0">
        <el-row :gutter="20">
          <el-col
            :span="8"
            v-for="(recommend, index) in recommends"
            :key="index"
          >
            <el-card
              :body-style="{ padding: '20px' }"
              @click="openMovieDrawer(recommend.ID)"
            >
              <div>
                <h3>{{ recommend.title }}</h3>
                <p>导演: {{ recommend.director }}</p>
                <p>上映时间: {{ recommend.CreatedAt }}</p>
                <p>{{ recommend.description }}</p>
              </div>
              
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
                  房间名: {{ room.room_name }} - 创建者:
                  {{ room.creator.name }} -
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
        <el-empty description="暂无电影推荐"></el-empty>
      </el-main>
    </section>
  </div>
</template>
<style scoped>
.homepage {
  font-family: "Arial", sans-serif;
  color: #fff;
  background: linear-gradient(135deg, #0f0f0f, #1a1a1a);
  padding: 0;
  margin: 0;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  background-color: #181818;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.5);
}

.logo {
  font-size: 1.8em;
  font-weight: bold;
  color: #fff;
}

.navbar ul {
  list-style: none;
  display: flex;
  gap: 20px;
  padding: 0;
  margin: 0;
}

.navbar ul li a {
  text-decoration: none;
  color: #aaa;
  transition: color 0.3s;
}

.navbar ul li a:hover {
  color: #fff;
}

.hero {
  height: 80vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: url("@/assets/1.jpg") no-repeat center center/cover;
  position: relative;
  text-align: center;
}

.hero::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
}

.hero-content {
  z-index: 1;
}

.hero h1 {
  font-size: 3em;
  margin: 0;
}

.hero p {
  font-size: 1.2em;
  margin: 10px 0 20px;
}

.cta-button {
  background-color: #ff4d4d;
  color: #fff;
  padding: 15px 30px;
  border: none;
  border-radius: 30px;
  cursor: pointer;
  font-size: 1em;
  transition: background-color 0.3s, transform 0.3s;
}

.cta-button:hover {
  background-color: #ff3333;
  transform: scale(1.05);
}

.movies-section {
  padding: 50px 20px;
}

.movies-section h2 {
  text-align: center;
  font-size: 2em;
  margin-bottom: 30px;
}

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
