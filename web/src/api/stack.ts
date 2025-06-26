import type { ListItem, StackItem, StackListResp, StackStatusListResp } from "@/types/stack"
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

/**
 * 获取 stack 状态
 * @param data 
 * @returns 
 */
export const getStackStatusAPI = (data: {stackName: string}) => {
    return http.get<StackStatusListResp>(
        `/stack/status?stackName=${data.stackName}`,
    )
}


/**
 * 获取 stack 
 * @param data 
 * @returns 
 */
export const getStackAPI = (data: {stackName: string}) => {
    return http.get<ListItem>(
        `/stack?stackName=${data.stackName}`,
    )
}

/**
 * 停止并删除 stack 
 * @param data 
 * @returns 
 */
export const downStackAPI = (data: {stackName: string}) => {
    return http.post(
        `/stack/down`,
        data,
    )
}

/**
 * 部署 stack 
 * @param data 
 * @returns 
 */
export const deployStackAPI = (data: {stackName: string, yamlContent: string, envContent: string}) => {
    return http.post(
        `/stack/down`,
        data,
    )
}

/**
 * 启动 stack 
 * @param data 
 * @returns 
 */
export const startStackAPI = (data: {stackName: string}) => {
    return http.post(
        `/stack/start`,
        data,
    )
}

/**
 * 重启 stack 
 * @param data 
 * @returns 
 */
export const reStartStackAPI = (data: {stackName: string}) => {
    return http.post(
        `/stack/restart`,
        data,
    )
}

/**
 * 更新 stack 
 * @param data 
 * @returns 
 */
export const updateStackAPI = (data: {stackName: string}) => {
    return http.post(
        `/stack`,
        data,
    )
}

/**
 * 停止 stack 
 * @param data 
 * @returns 
 */
export const stopStackAPI = (data: {stackName: string}) => {
    return http.post(
        `/stack/stop`,
        data,
    )
}

/**
 * 删除 stack 
 * @param data 
 * @returns 
 */
export const deleteStackAPI = (data: {stackName: string}) => {
    return http.delete(
        `/stack?stackName=${data.stackName}`,
    )
}