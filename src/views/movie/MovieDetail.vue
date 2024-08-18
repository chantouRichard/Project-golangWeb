<template>
  <el-dialog
    v-model:visible="isVisible"
    title="Movie Details"
    width="50%"
    @close="closeDetail"
  >
    <div v-if="movie">
      <h3>{{ movie.title }}</h3>
      <p>{{ movie.description }}</p>
      <p><strong>Director:</strong> {{ movie.director }}</p>
      <p><strong>Release Date:</strong> {{ movie.releaseDate }}</p>
      <!-- 显示更多电影详细信息 -->
    </div>
    <div v-else>
      <p>Loading...</p>
    </div>
  </el-dialog>
</template>

<script setup>
import { ref, watch } from "vue";
import { movieDetailService } from "@/services/movieService";

const props = defineProps({
  id: String,
  visible: Boolean,
});

const emit = defineEmits(["update:visible"]);

const isVisible = ref(props.visible);
const movie = ref(null);

watch(
  () => props.visible,
  async (newVal) => {
    isVisible.value = newVal;
    if (newVal && props.id) {
      // 请求电影详情数据
      try {
        const response = await movieDetailService({ id: props.id });
        movie.value = response.data.movie;
      } catch (error) {
        console.error("Failed to fetch movie details:", error);
      }
    }
  }
);

const closeDetail = () => {
  emit("update:visible", false);
};
</script>

<style scoped>
/* 添加你的样式 */
</style>
