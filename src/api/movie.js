import request from '@/utils/request'
import {useTokenStore} from '@/stores/token.js'
import { da } from 'element-plus/es/locales.mjs'
//电影分类列表
export const movieCategoryListService = (params) =>{
    // const tokenStore = useTokenStore()
    
    // return request.get('/api/movies',{headers:{'Authorization':tokenStore.token}})
    return request.get('/api/movies',{params:params})
}

//筛选电影
export const movieFilterListService = (params) =>{
    return request.get('/api/movies/filter',{params:params})
}

// 电影详情服务
export const movieDetailService = (id) => {
    return request.get(`/api/movies/${id}`);
}
//电影评论列表
export const movieCommentListService = (params) =>{
    return request.get('/api/movies/comments',{params:params})
}

//电影房间信息
export const movieRoomInfoService = () =>{
    return request.get('/api/rooms')
}

//房间列表
export const movieRoomListService = (movieID) =>{
    return request.get(`/api/rooms/movie/${movieID}`)
}

//进入房间
export const enterRoomService = (roomID) => {
    const params = {
        room_id: roomID,  // 修改这里的键名
    };
    return request.post('/api/rooms/join', params);
};

//用户退出房间
export const exitRoomService = () => {
    return request.post('/api/rooms/leave');
}

//播放视频
export const playVideoService = (roomID) => {
    return request.get(`/api/rooms/${roomID}/stream`)
}
