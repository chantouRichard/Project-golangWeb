//导入request.js请求工具
import request from '@/utils/request'

//提供调用注册接口的函数
export const userRegisterService = (registerData) => {
    const params = new URLSearchParams()
    for(let key in registerData) {
        params.append(key, registerData[key])
    }
    return request.post('/api/register', params);
}

//提供调用登录接口的函数
export const userLoginService = (loginData) => {
    const params = new URLSearchParams()
    for(let key in loginData) {
        params.append(key, loginData[key])
    }
    return request.post('/api/login', params);
}


//提供调用获取用户信息的函数
export const userInfoService = () => {
    return request.get('/api/user');
}

//提供调用修改用户信息的函数
export const updateUserInfoService = (userData) => {
    const params = new URLSearchParams()
    for(let key in userData) {
        params.append(key, userData[key])
    }
    return request.put('/api/user', params);
}

//提供调用修改密码的函数
export const updatePasswordService = (passwordData) => {
    return request.post('/api/user/updatepwd', passwordData);
}

//我的收藏
export const fetchCollectionsService = () => {
    return request.get('/api/user/collections');
}

//观看记录
export const fetchHistoryService = () => {
    return request.get('/api/user/history');
}

//删除收藏
export const deleteCollectionService = (id) => {
    return request.delete(`/api/user/collections/${id}`);
}

//删除观看记录
export const deleteHistoryService = () => {
    return request.delete('/api/user/history');
}
