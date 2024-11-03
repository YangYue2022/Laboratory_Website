import axios from 'axios';

// 创建 axios 实例
const service = axios.create({
  baseURL: '', // 配置基础 API 路径
  timeout: 5000, // 请求超时时间
});

// 请求拦截器
service.interceptors.request.use(
  config => {
    // 从 localStorage 中获取 token
    const token = localStorage.getItem('token');
    if (token) {
      // 如果 token 存在，统一在请求头中添加 token
      config.headers['Authorization'] = 'Bearer ' + token;
    }
    return config;
  },
  error => {
    // 请求错误处理
    console.error('Request Error: ', error);
    return Promise.reject(error);
  }
);

export default service;
