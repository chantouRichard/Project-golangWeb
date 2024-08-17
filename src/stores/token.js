//定义store
import { defineStore } from "pinia";
import { ref } from "vue";

/*
    第一个参数：名字唯一性
    第二个参数：函数，返回一个对象

    返回值：函数
 */
export const useTokenStore = defineStore('token', ()=>{
    //定义状态的内容

    //1.响应式变量
    const token = ref('');

    //2.定义一个函数，用来修改token
    const setToken = (newToken)=>{
        token.value = newToken;
    }

    //3.函数，移除token的值
    const removeToken = ()=>{
        token.value = '';
    }

    //4.返回值
    return {
        token,
        setToken,
        removeToken
    }
},{
    persist: true//持久化存储，默认是关闭的
});