import type { Response } from "@/types/base"
import type { LoginRes } from "@/types/user"
import { request } from "@/util/request"

export const loginAPI = (data: {username: string, password: string}) => {
    return request<Response<LoginRes>>({
        url: '/login/username',
        method: 'post',
        data: data
    })
}