import { http } from "@/utils/request"

export const loginAPI = (data: { username: string, password: string }) => {
    return http.post(
        '/login/username',
        data,
    )
}

export const changePasswordAPI = (data: {currentPassword: string, password: string, password2: string}) => {
    return http.patch(
        '/user/password',
        data,
    )
}