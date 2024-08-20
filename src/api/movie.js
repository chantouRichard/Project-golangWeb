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

//创建房间
export const createRoomService = (params) => {
    return request.post('/api/rooms', params);
}

//删除房间
export const closeRoomService = (roomID) => {
    return request.delete(`/api/rooms/${roomID}`)
}

//加入收藏
export const addCollectionService = (movieID) => {
    return request.post('/api/movies/collections',{params:{movie_id: movieID}})
}

//添加历史记录
export const addHistoryService = (movieID) => {
    const params = {
        movie_id: movieID,  // 修改这里的键名
    };
    return request.post('/api/user/history',params);
}

//搜索电影
export const searchMovieService = (params) =>{
    return request.get('/api/movies/search',{params:params})
}

//电影推荐
export const movieRecommendService = () =>{
    return request.get('/api/movies/recommend')
}

// //获取电影封面图片
// export const movieCoverService = (movieID) =>{
//     return request.get(`/api/movies/picture/${movieID}`)
// }
