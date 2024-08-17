<template>
  <el-container class="layout">
    <!-- Header -->
    <el-header class="header">
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

    <!-- Sidebar & Main Content -->
    <el-container>
      <el-aside width="200px" class="sidebar">
        <el-menu
          class="el-menu-vertical-demo"
        >
          <el-menu-item index="1">首页</el-menu-item>
          <el-menu-item index="2">热门电影</el-menu-item>
          <el-menu-item index="3">Favorites</el-menu-item>
          <el-menu-item index="4">Settings</el-menu-item>
          <el-sub-menu index="5" title="个人">
            <template #title>
                <span style="color: #000000;">个人中心</span>
            </template>
            <el-menu-item index="6">
                个人资料
            </el-menu-item>
            <el-menu-item index="7">
                我的收藏
            </el-menu-item>
            <el-menu-item index="8">
                观看记录
            </el-menu-item>
        </el-sub-menu>
        </el-menu>
      </el-aside>

      <el-main class="main-content">
        <div class="content-header">Welcome to MovieStream</div>
        <div class="content-body">
          <p>
            Here you can stream the latest movies and series. Enjoy the best
            experience with our platform!
          </p>
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
