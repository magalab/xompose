export type StackListItem = StackItem & {
    isManaged: Boolean
}

export type StackListResp = {
    items: Array<StackListItem>
}


export type StackItem = {
    stackName: String
    stackStatus: String
    yamlContent: String
}