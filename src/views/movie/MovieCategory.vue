<script setup>
import axios from "axios";
import { ref } from "vue";

const categorys = ref([]);
const currentPage = ref(1); // 当前页
const pageSize = ref(3); // 每页显示的条数
const total = ref(0); // 总记录数

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
      <el-table-column label="电影图片" prop="video_url"></el-table-column>
      <el-table-column label="分类名称" prop="genre"></el-table-column>
      <el-table-column label="评分" prop="rating"></el-table-column>
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
  </div>
</template>

<style scoped>
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
