export type Response<T = any> = {
    code: number,
    message: string,
    data: T,
}