import axios, { type AxiosError, type AxiosInstance, type AxiosRequestConfig, type AxiosResponse, type InternalAxiosRequestConfig } from "axios";


type Response<T = any> = {
    code: number,
    message: string,
    data: T,
}

const request: AxiosInstance = axios.create({
    // baseURL: import.meta.env.VITE_APP_BASE_URL,
    baseURL: "/api/xompose/v1",
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
    (error: AxiosError) => {
        return Promise.reject(error)
    }
)

// 响应拦截器
request.interceptors.response.use(
    (response: AxiosResponse<Response>) => {
        const { code, data, message } = response.data
        if (code === 0) {
            return data
        }
        ElMessage.error(message || '系统出错了')
        return Promise.reject(message)
    },
    async (error) => {
        const response = error.response
        if (response) {
            const { message } = response.data
            ElMessage.error(message || '系统出错啦')
        }

        return Promise.reject(error)
    }
)

export const http = {
    get<T = any>(url: string, config?: AxiosRequestConfig): Promise<T> {
        return request.get(url, config)
    },

    post<T = any>(url: string, data?: object, config?: AxiosRequestConfig): Promise<T> {
        return request.post(url, data, config)
    },

    put<T = any>(url: string, data?: object, config?: AxiosRequestConfig): Promise<T> {
        return request.put(url, data, config)
    },

    patch<T = any>(url: string, data?: object, config?: AxiosRequestConfig): Promise<T> {
        return request.put(url, data, config)
    },

    delete<T = any>(url: string, config?: AxiosRequestConfig): Promise<T> {
        return request.delete(url, config)
    }
}