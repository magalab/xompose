import type { Response } from "@/types/base"
import type { LoginRes } from "@/types/user"
import { http } from "@/utils/request"

export const loginAPI = (data: { username: string, password: string }) => {
    return http.post<Response<LoginRes>>(
        '/login/username',
        data,
    )
}

export const changePassword = (data: {password: string, password2: string}) => {
    return http.patch<Response>(
        '/user/password',
        data,
    )
}