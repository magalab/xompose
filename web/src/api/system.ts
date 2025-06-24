import { http } from "@/utils/request"

export const setupAPI = (data: { username: string, password: string, password2: string }) => {
    return http.post(
        "/system/setup",
        data,
    )
}