import { http } from "@/utils/request"
import type {  NetworkListResp } from "@/types/network"

export const networkListAPI = () => {
    return http.get<NetworkListResp>(
        '/networks',
    )
}