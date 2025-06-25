import type { StackItem, StackListResp } from "@/types/stack"
import { http } from "@/utils/request"

export const listStackAPI = () => {
    return http.get<StackListResp>(
        '/stacks',
    )
}

export const saveStackAPI = (data: StackItem) => {
    return http.post(
        '/stack',
        data,
    )
}