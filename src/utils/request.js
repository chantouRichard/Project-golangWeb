//定制请求的实例

//导入axios  npm install axios
import axios from 'axios';

import { ElMessage } from 'element-plus';
//定义一个变量,记录公共的前缀  ,  baseURL
// const baseURL = 'http://localhost:8080';
const baseURL = '/sub';
const instance = axios.create({baseURL})

//添加请求拦截器
import { useTokenStore } from '@/stores/token.js';
instance.interceptors.request.use(
    (config)=>{
        //请求前的回调
        //添加token
        const tokenStore = useTokenStore();
        //判断token是否存在
        if(tokenStore.token)
        {
            config.headers.Authorization = tokenStore.token;
        }
        return config;
    },
    (err)=>{
        //请求错误的回调
        Promise.reject(err);
    }
)

import router from '@/router';
//添加响应拦截器
instance.interceptors.response.use(
    result=>{
        //判断业务状态码
        if(result.data.status==="success"){
            return result.data;
        }

        //业务状态码不等于0,则提示错误信息
        // alert(result.data.msg?result.data.msg:"服务异常");
        ElMessage.error(result.data.msg ?result.data.msg:"服务异常");
        //异步操作的状态转化成失败的状态
        return Promise.reject(result.data);
    },
    err=>{
        //判断响应状态码，如果响应状态码为401,则跳转到登录页
        if(err.response&&err.response.status===401){
            //没有权限,跳转到登录页
            ElMessage.error('没有权限,请登录');
            router.push('/api/register')
        }
        else ElMessage.error('服务异常');
        return Promise.reject(err);//异步的状态转化成失败的状态
    }
)

export default instance;