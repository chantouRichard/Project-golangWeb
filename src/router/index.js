import { createRouter,createWebHistory } from "vue-router";

//导入组件
import LoginVue from '@/views/Login.vue'
import Layout from "@/views/Layout.vue";

//定义路由关系
const routes = [
    {path: '/api/register', component: LoginVue},
    {path: '/' , component: Layout}
]

//创建路由对象
const router = createRouter({
    history:createWebHistory(),
    routes:routes
})

//导出路由对象
export default router;