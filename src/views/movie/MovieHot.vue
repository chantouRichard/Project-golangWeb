<template>
  <div class="movie-search">
    <el-input
      v-model="searchQuery"
      placeholder="Enter movie name"
      @keyup.enter="searchMovies"
      clearable
    ></el-input>

    <el-radio-group v-model="searchMode">
      <el-radio label="exact">Exact Search</el-radio>
      <el-radio label="fuzzy">Fuzzy Search</el-radio>
    </el-radio-group>

    <el-button type="primary" @click="searchMovies">Search</el-button>

    <el-table :data="movies" style="width: 100%" v-if="movies.length">
      <el-table-column
        prop="title"
        label="Movie Name"
        width="180"
      ></el-table-column>
      <el-table-column
        prop="release_date"
        label="Release Date"
        width="180"
      ></el-table-column>
      <el-table-column prop="genre" label="Genre"></el-table-column>
    </el-table>

    <el-empty v-else description="No movies found"></el-empty>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { ElMessage } from "element-plus";

const searchQuery = ref("");
const searchMode = ref("exact");
const movies = ref([]);

import { searchMovieService } from "@/api/movie";
const searchMovies = async () => {
  if (!searchQuery.value) {
    ElMessage.warning("Please enter a movie name to search");
    return;
  }

  try {
    let params = {
      query: searchQuery.value,
      mode: searchMode.value,
    };
    const response = await searchMovieService(params);

    movies.value = response.movies;
  } catch (error) {
    ElMessage.error("Failed to fetch movies");
    console.error(error);
  }
};
</script>

<style scoped>
.movie-search {
  padding: 20px;
}
</style>
