
export type StackList = {
    id: string
    name: string
    tags: string[]
}

export type StackListResp = {
    items: Array<ListItem>
}

export type ListItem = StackItem & {
    isManaged: Boolean
}

export type StackItem = {
    id?: string
    stackName: string
    stackStatus?: string
    yamlContent: string
    yamlPath: string
    envContent: string
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

export type StackStatusListResp = {
    items: StackStatusItem[]
}

export type StackStatusItem = {
    service: string
    image: string
    state: string
    ports: Port[]
    networks: string[]
    volumes: string[]
    dependsOn: string[]
}

export type Port = {
    publishPort: number
    targetPort : number
    protocol: string
}