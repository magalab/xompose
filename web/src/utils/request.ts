import axios, { type AxiosInstance, type AxiosResponse, type InternalAxiosRequestConfig } from "axios";

export const request: AxiosInstance = axios.create({
    baseURL: import.meta.env.VITE_APP_BASE_URL,
    timeout: 10000,
    headers: {
        'Content-Type': 'application/json'
    }
})

// 请求拦截器
request.interceptors.request.use(
    (config: InternalAxiosRequestConfig) => {
        return config
    },
    error => {
        return Promise.reject(error)
    }
)

// 响应拦截器
request.interceptors.response.use(
    (response: AxiosResponse) => {
        const { code, data, msg } = response.data
        if (code === 200) {
            return data
        }
        ElMessage.error(msg || '系统出错')
    },
    async (error) => {
        const response = error.response
        if (response) {
            const { msg } = response.data
            ElMessage.error(msg || '系统出错')
        }

        return Promise.reject(error)
    }
)