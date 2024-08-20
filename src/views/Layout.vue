<script setup>
import { ElMessage, ElMessageBox } from "element-plus";
import { useTokenStore } from "@/stores/token.js";
import { useUserInfoStore } from "@/stores/userInfo.js";
import { useRouter } from "vue-router";
const userInfoStore = useUserInfoStore();
const tokenStore = useTokenStore();
const router = useRouter();
const exitLogin = async () => {
  useTokenStore.removeItem;
  ElMessageBox.confirm("您确认要退出登录吗？", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      //调用退出登录接口
      //1、清空pinia中的用户信息和token
      tokenStore.removeToken();
      userInfoStore.removeInfo();
      //2、跳转到登录页面
      router.push("/api/register");
      ElMessage({
        type: "success",
        message: "退出登录成功！",
      });
    })
    .catch(() => {
      ElMessage({
        type: "info",
        message: "已取消退出登录！",
      });
    });
};
</script>

<template>
  <el-container class="layout">
    <!-- Header -->
    <el-header class="header" style="position: relative">
      <div class="logo">MovieStream</div>
      <!-- <el-menu
        mode="horizontal"
        class="el-menu-demo"
      >
        <el-menu-item index="1">Home</el-menu-item>
        <el-menu-item index="2">Movies</el-menu-item>
        <el-menu-item index="3">Series</el-menu-item>
        <el-menu-item index="4">My Collection</el-menu-item>
        <el-menu-item index="5">My Collection</el-menu-item>
        <el-menu-item index="6">My Collection</el-menu-item>
      </el-menu> -->
    </el-header>

    <!-- 左侧菜单 -->
    <el-container>
      <el-aside width="200px" class="sidebar">
        <el-menu
          style="position: fixed; width: 200px"
          class="el-menu-vertical-demo"
          router
        >
          <el-menu-item index="/api/movies/main">首页</el-menu-item>
          <el-menu-item index="/api/movieHot">热门电影</el-menu-item>
          <el-menu-item index="/api/movies">电影分类</el-menu-item>
          <el-sub-menu index="4" title="个人">
            <template #title>
              <span style="color: #000000">个人中心</span>
            </template>
            <el-menu-item index="/api/userCenter"> 个人资料 </el-menu-item>
            <el-menu-item index="/api/userFavorite"> 我的收藏 </el-menu-item>
            <el-menu-item index="/api/userHistory"> 观看记录 </el-menu-item>
            <el-menu-item @click="exitLogin"> 退出登录 </el-menu-item>
          </el-sub-menu>
          <el-menu-item index="/api/movieRoom"> 我的房间 </el-menu-item>
          <el-menu-item index="/"> 联系我们 </el-menu-item>
        </el-menu>
      </el-aside>

      <!-- 中间区域 -->
      <el-main class="main-content">
        <!-- <div class="content-header">Welcome to MovieStream</div> -->
        <div class="content-body">
          <router-view></router-view>
          <!-- Add your dynamic content here -->
        </div>
      </el-main>
    </el-container>

    <!-- Footer -->
    <el-footer class="footer">
      © 2024 MovieStream | All Rights Reserved
    </el-footer>
  </el-container>
</template>

<style scoped>
/* General Layout Styles */
.layout {
  height: 100vh;
  display: flex;
  flex-direction: column;
}

/* .el-sub-menu {
  background-color: #1f2d3d;
  color: #ffffff;
}

.el-menu-item {
  background-color: #1f2d3d;
  color: #ffffff;
}

.el-menu-item:hover {
    background-color: #2b79d3;
}

.el-menu-item:focus, .el-menu-item:active {
    background-color: #2b79d3;
    color: #ffffff;
} */
/* Header Styles */
.header {
  background-color: #1f2d3d;
  display: flex;
  align-items: center;
  padding: 0 20px;
  color: #ffffff;
}

.logo {
  font-size: 1.5em;
  font-weight: bold;
  margin-right: 20px;
}

/* Sidebar Styles */
.sidebar {
  background-color: #304156;
  color: #ffffff;
  padding: 20px 0;
}

/* Main Content Styles */
.main-content {
  padding: 20px;
  background-color: #f5f7fa;
}

.content-header {
  font-size: 2em;
  margin-bottom: 20px;
}

.content-body {
  background: linear-gradient(135deg, #f6d365 0%, #fda085 100%);
  padding: 20px;
  border-radius: 10px;
  color: #ffffff;
  box-shadow: 0 4px 14px rgba(0, 0, 0, 0.1);
}

/* Footer Styles */
.footer {
  background-color: #1f2d3d;
  color: #ffffff;
  text-align: center;
  padding: 10px 0;
}
</style>
