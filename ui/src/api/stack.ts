import type { Response } from "@/types/base"
import type { StackItem, StackListResp } from "@/types/stack"
import { request } from "@/util/request"

export const listStackAPI = () => {
    return request<Response<StackListResp>>({
        url: '/stacks',
        method: 'get',
        // data: data
    })
}

export const saveStackAPI = (data: StackItem) => {
    return request<Response<any>>({
        url: '/stacks',
        method: 'get',
        data: data
    })
}