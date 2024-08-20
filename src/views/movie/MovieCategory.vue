<script setup>
import axios from "axios";
import { ref } from "vue";

const categorys = ref([]);
const currentPage = ref(1); // 当前页
const pageSize = ref(3); // 每页显示的条数
const total = ref(0); // 总记录数

const selectedMovie = ref(null);
const rooms = ref([]);
const showMovieDetails = ref(false);
const isDialogVisible = ref(false);

// 通过 rowStyle 函数为每一行设置自定义样式
const rowStyle = () => {
  return {
    height: "200px",
  };
};

const rowClassName = () => {
  return "custom-row";
};

const filter = ref({
  genre: "",
  minRating: "",
});

const movies = ref([]);

import { movieFilterListService } from "@/api/movie";
const fetchMovies = async () => {
  const params = {
    genre: filter.value.genre,
    min_rating: filter.value.minRating,
    page: currentPage.value, // 当前页
    page_size: pageSize.value, // 每页显示的条数
  };

  let response = await movieFilterListService(params);

  if (response.status === "success") {
    categorys.value = response.movies;
    total.value = response.total; // 从后端获取总记录数
  } else {
    categorys.value = [];
    total.value = 0;
  }
};

//声明一个异步的函数
import { movieCategoryListService } from "@/api/movie.js";
import { ElMessage } from "element-plus";
const movieCategoryList = async () => {
  const params = {
    page: currentPage.value, // 当前页
    page_size: pageSize.value, // 每页显示的条数
  };

  let result = await movieCategoryListService(params);

  if (result.status === "success") {
    ElMessage.success("movies:" + result.movies.length);
    categorys.value = result.movies;
    total.value = result.total; // 假设返回的数据中有 total 字段，表示总记录数
  } else {
    categorys.value = [];
    total.value = 0;
  }
};

//调用函数
movieCategoryList();

// 当页码或页大小改变时，重新获取数据
const handlePageSizeChange = (newPageSize) => {
  ElMessage.success("当前页码：" + newPageSize);
  pageSize.value = newPageSize;
  fetchMovies(); // 重新获取数据
};

const handlePageChange = (newPage) => {
  ElMessage.success("当前页码：" + newPage);
  currentPage.value = newPage;
  fetchMovies(); // 重新获取数据
};

const clearFetch = () => {
  filter.value = { genre: "", minRating: "" };
  currentPage.value = 1; // 重置到第一页
  fetchMovies();
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
import { addHistoryService } from "@/api/movie";
const openRoom = async (roomId) => {
  let response = await enterRoomService(roomId);
  if (response.status === "success") {
    ElMessage.success("进入房间成功");
  } else {
    ElMessage.error("进入房间失败");
  }
};

// 添加历史记录
const addHistory = async (movieId) => {
  let response = await addHistoryService(movieId);
  if (response.status === "success") {
    ElMessage.success("添加历史记录成功");
  }
};
const openMovieDrawer = async (row) => {
  console.log("row:", row);
  ElMessage.success(row.ID);
  await fetchMovieDetails(row.ID);
  await fetchRoomDetails(row.ID);
  showMovieDetails.value = true;
};

const formatDate = (dateString) => {
  const date = new Date(dateString);
  return date.toLocaleString();
};

//创建房间
const roomName = ref("");
// 打开弹窗
const showCreateRoomDialog = () => {
  isDialogVisible.value = true;
  ElMessage.success("打开弹窗:" + isDialogVisible.value);
};
// 提交房间创建请求
import { createRoomService } from "@/api/movie";
const submitCreateRoom = async () => {
  if (!roomName.value.trim()) {
    ElMessage.error("房间名称不能为空");
    return;
  }

  const movieID = selectedMovie.value.ID;

  // 构造请求数据
  const requestData = {
    movie_id: movieID,
    room_name: roomName.value.trim(),
  };

  let response = await createRoomService(requestData);

  if (response.status === "success") {
    ElMessage.success("创建房间成功");
    isDialogVisible.value = false; // 关闭弹窗
    await fetchRoomDetails(movieID); // 获取房间详情
  } else {
    ElMessage.error("创建房间失败");
  }
};

//获取电影封面图片
const getImageUrl = async(movieID) => {
  // let result = await movieDetailService(movieID);
  // ElMessage.success(movieID);
  // console.log("result:", result);
  // console.log("result.movie.url:", result.movie.thumbnail_url);
  // return result.movie.thumbnail_url;
  return `http://localhost:8080/api/movies/picture/${movieID}`;
};
</script>

<template>
  <div class="movie-filter">
    <el-form :model="filter" label-width="100px" class="filter-form">
      <el-form-item label="Genre">
        <el-select v-model="filter.genre" placeholder="Select Genre">
          <el-option label="Action" value="action"></el-option>
          <el-option label="Comedy" value="Comedy"></el-option>
          <el-option label="Drama" value="Drama"></el-option>
          <el-option label="Sci-Fi" value="Sci-Fi"></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="Min Rating">
        <el-input
          v-model="filter.minRating"
          placeholder="Enter minimum rating"
        ></el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="fetchMovies">搜索</el-button>
        <el-button type="primary" @click="clearFetch">清空</el-button>
      </el-form-item>
    </el-form>

    <el-table :data="movies" style="width: 100%" v-if="movies.length">
      <el-table-column prop="title" label="Title" width="180"></el-table-column>
      <el-table-column prop="genre" label="Genre" width="180"></el-table-column>
      <el-table-column
        prop="rating"
        label="Rating"
        width="100"
      ></el-table-column>
    </el-table>

    <el-table
      :data="categorys"
      style="width: 100%; height: inherit"
      :row-style="rowStyle"
      :row-class-name="rowClassName"
    >
      <el-table-column label="序号" width="100" type="index"> </el-table-column>
      <el-table-column label="电影名称" prop="title"></el-table-column>
      <!-- 电影图片列 -->
      <el-table-column label="电影图片">
        <template #default="{ row }">
          <img
            :src="getImageUrl(row.ID)"
            alt="电影图片"
            style="width: auto; height: 150px"
          />
        </template>
      </el-table-column>
      <el-table-column label="分类名称" prop="genre"></el-table-column>
      <el-table-column label="评分" prop="rating"></el-table-column>
      <el-table-column label="电影详情" width="100%">
        <template #default="{ row }">
          <el-button
            plain
            type="warning"
            @click="openMovieDrawer(row)"
          ></el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 添加分页控件 -->
    <el-pagination
      background
      layout="total, prev, pager, next, sizes, jumper"
      :total="total"
      :page-size="pageSize"
      :current-page="currentPage"
      @size-change="handlePageSizeChange"
      @current-change="handlePageChange"
      :page-sizes="[3, 5, 10, 20]"
      style="text-align: right; margin-top: 20px"
    ></el-pagination>
    <!-- 抽屉 -->
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
            <el-button
              type="primary"
              @click="
                openRoom(room.id);
                addHistory(room.movie_id);
              "
              >进入房间</el-button
            >
          </li>
        </ul>
        <el-empty v-else description="暂无相关房间"></el-empty>
        <el-button type="primary" @click="showCreateRoomDialog"
          >创建我的房间</el-button
        >
      </div>
      <el-empty v-else description="暂无电影详情"></el-empty>
    </el-drawer>

    <!-- 弹窗 -->
    <el-dialog title="创建房间" v-model="isDialogVisible" width="30%">
      <el-input v-model="roomName" placeholder="请输入房间名称"></el-input>

      <span class="dialog-footer">
        <el-button @click="isDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitCreateRoom">创建</el-button>
      </span>
    </el-dialog>
  </div>
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

.custom-row {
  background-color: #000000; /* 更改背景色 */
  height: 60px; /* 更改行高 */
  font-size: 32px; /* 更改字体大小 */
}

.movie-filter {
  padding: 20px;
  background-color: #1e1e1e;
  color: #fff;
}

.filter-form {
  margin-bottom: 20px;
  background-color: #2b2b2b;
  padding: 20px;
  border-radius: 8px;
}

.el-input,
.el-select {
  width: 100%;
}

.no-movies {
  text-align: center;
  margin-top: 20px;
  font-size: 18px;
  color: #f5f5f5;
}

.el-table {
  background-color: #2b2b2b;
}

.el-table th,
.el-table td {
  color: #fff;
  background-color: #2b2b2b;
}

.el-table__header-wrapper {
  background-color: #2b2b2b;
}
</style>
