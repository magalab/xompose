import axios, { type AxiosInstance, type AxiosResponse, type AxiosError } from "axios";
export const request: AxiosInstance = axios.create({
    baseURL: "/api/xompose/v1",
    timeout: 1000,
    withCredentials: true,
    headers: {
        "Content-Type": "application/json",
    },
})

request.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('token');
        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`;
        }
        return config;
    },
    (error: AxiosError) => {
        console.error('Request Error:', error);
        return Promise.reject(error);
    }
)

request.interceptors.response.use(
    (response: AxiosResponse) => {
        return response;
    },
    (error: AxiosError) => {
        if (error.response) {
            console.error('Response Error:', error.response.status);
        } else if (error.request) {
            console.error('No Response:', error.request);
        } else {
            console.error('Error:', error.message);
        }
        return Promise.reject(error);
    }
)

