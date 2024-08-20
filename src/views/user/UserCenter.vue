<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";

const user = ref({
  username: "",
  email: "",
  nickname: "",
  phone: "",
});

const passwordModel = ref({
  old_pwd: "",
  new_pwd: "",
  re_pwd: "",
});

const rules = {
  old_pwd: [
    { required: true, message: "请输入用户名", trigger: "blur" },
    { min: 5, max: 16, message: "长度在 5 到 16 个字符", trigger: "blur" },
  ],
  new_pwd: [
    { required: true, message: "请输入新密码", trigger: "blur" },
    { min: 5, max: 16, message: "长度在 5 到 16 个字符", trigger: "blur" },
  ],
  re_pwd: [
    { required: true, message: "请输入确认密码", trigger: "blur" },
    { min: 5, max: 16, message: "长度在 5 到 16 个字符", trigger: "blur" },
  ],
};

import { userInfoService } from "@/api/user";
import { ElMessage, ElMessageBox } from "element-plus";
const fetchUserInfo = async () => {
  // 从本地存储获取token
  const response = await userInfoService();
  if (response.status === "success") {
    ElMessage.success("获取用户信息成功");
    user.value = response.user;
  } else {
    ElMessage.error("获取用户信息失败");
    console.error(response.message);
  }
};

import { updateUserInfoService } from "@/api/user";
const updateUser = async () => {
  ElMessageBox.confirm("您确认要更新吗？", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      const response = await updateUserInfoService(user.value);
      if (response.status === "success") {
        ElMessage({
          type: "success",
          message: "更新成功！",
        });
      } else {
        console.error(response.data.message);
      }
    })
    .catch(() => {
      ElMessage({
        type: "info",
        message: "已取消更新信息！",
      });
    });
};

import { updatePasswordService } from "@/api/user";
const updatePassword = async () => {
  ElMessageBox.confirm("您确认要更新密码吗？", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      const response = await updatePasswordService(passwordModel.value);

      if (response.status === "success") {
        ElMessage({
          type: "success",
          message: "更新成功！",
        });
      }
    })
    .catch(() => {
      ElMessage({
        type: "info",
        message: "已取消修改密码！",
      });
    });
};

fetchUserInfo();
</script>

<template>
  <el-container>
    <!-- 个人资料部分 -->
    <el-card shadow="hover" style="margin: 20px;width: 500px;height: 500px;">
      <div class="header">
        <span>个人资料</span>
      </div>
      <el-form :model="user" ref="userForm" label-width="100px" style="margin-top: 50px;">
        <el-form-item label="用户名" style="margin-top: 50px;">
          <el-input v-model="user.username"></el-input>
        </el-form-item>
        <el-form-item label="邮箱" style="margin-top: 50px;">
          <el-input v-model="user.email"></el-input>
        </el-form-item>
        <el-form-item label="昵称" style="margin-top: 50px;">
          <el-input v-model="user.nickname"></el-input>
        </el-form-item>
        <el-form-item label="电话" style="margin-top: 50px;">
          <el-input v-model="user.phone"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="updateUser" style="margin-top: 0px;width: 200px;height: 50px;border-radius: 50px;">更新资料</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 重置密码部分 -->
    <el-card shadow="hover" style="margin: 20px;width: 500px;height: 500px;">
      <div class="header">
        <span>重置密码</span>
      </div>
      <el-form
        :model="passwordModel"
        :rules="rules"
        label-width="100px"
        size="large"
        style="margin-top: 50px;"
      >
        <el-form-item label="原密码" prop="old_pwd"  style="margin-top: 50px;">
          <el-input
            :prefix-icon="Lock"
            type="password"
            placeholder="请输入原密码"
            v-model="passwordModel.old_pwd"
          ></el-input>
        </el-form-item>
        <el-form-item label="新密码" prop="new_pwd"  style="margin-top: 50px;">
          <el-input
            :prefix-icon="Lock"
            type="password"
            placeholder="请输入新密码"
            v-model="passwordModel.new_pwd"
          ></el-input>
        </el-form-item>
        <el-form-item label="确认新密码" prop="re_pwd"  style="margin-top: 50px;">
          <el-input
            :prefix-icon="Lock"
            type="password"
            placeholder="请再次输入密码"
            v-model="passwordModel.re_pwd"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="updatePassword" style="margin-top: 50px;width: 200px;height: 50px;border-radius: 50px;">提交修改</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </el-container>
</template>

<style scoped>
.header {
  font-size: 20px;
  font-weight: bold;
  color: #409eff;
}

.el-card {
  border-radius: 8px;
  border: 1px solid #ebeef5;
}

.el-form-item {
  margin-bottom: 20px;
}
</style>
