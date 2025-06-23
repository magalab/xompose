
export type StackList = {
    id: string
    name: string
    tags: string[]
}


export const UNKNOWN = 0
export const CREATED_FILE = 1
export const CREATED_STACK = 2
export const RUNNING = 3
export const EXITED = 4


export const statusColor = (status: number): string => {
    switch (status) {
        case CREATED_FILE:
            return "dark";
        case CREATED_STACK:
            return "dark";
        case RUNNING:
            return "primary";
        case EXITED:
            return "danger";
        default:
            return "secondary";
    }
}

export const statusNameShort = (status: number): string => {
    switch (status) {
        case CREATED_FILE:
            return "inactive";
        case CREATED_STACK:
            return "inactive";
        case RUNNING:
            return "active";
        case EXITED:
            return "exited";
        default:
            return "?";
    }
}