import { createRouter,createWebHistory } from "vue-router";

//导入组件
import LoginVue from '@/views/Login.vue'
import Layout from "@/views/Layout.vue";

import MovieMain from "../views/movie/MovieMain.vue";
import MovieCategory from "../views/movie/MovieCategory.vue";
import MovieHot from "../views/movie/MovieHot.vue";
import MovieRoom from "../views/movie/MovieRoom.vue";
import UserCenter from "../views/user/UserCenter.vue";
import UserFavorite from "../views/user/UserFavorite.vue";
import UserHistory from "../views/user/UserHistory.vue";


//定义路由关系
const routes = [
    {path: '/api/register', component: LoginVue},
    {path: '/' , component: Layout,children:[
        {path: '/api/movies/main', component: MovieMain},
        {path: '/api/movies', component: MovieCategory},
        {path: '/api/movieHot', component: MovieHot},
        {path: '/api/movieRoom', component: MovieRoom},
        {path: '/api/userCenter', component: UserCenter},
        {path: '/api/userFavorite', component: UserFavorite},
        {path: '/api/userHistory', component: UserHistory}
    ]}
]

//创建路由对象
const router = createRouter({
    history:createWebHistory(),
    routes:routes
})

//导出路由对象
export default router;