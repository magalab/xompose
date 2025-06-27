<template>
    <transition name="slide-fade" appear>
        <div>
            <h1 v-if="isAdd" class="mb-3 text-xl fw-700"> Compose </h1>
            <h1 v-else class="mb-3">
                <Uptime :stack="stack" :pill="true" /> {{ stack.stackName }}
            </h1>
            <div v-if="stack.isManaged" class="mb-3 flex items-center p-4px">
                <el-button-group>
                    <el-button type="primary" size="large" v-if="isEditMode" :disabled="processing"
                        @click="deployStack">
                        <svg class="i-lucide-rocket"></svg>
                        {{ $t("deployStack") }}
                    </el-button>
                    <el-button type="primary" size="large" v-if="isEditMode" :disabled="processing" @click="saveStack">
                        <svg class="i-lucide-save"></svg>
                        {{ $t("saveStackDraft") }}
                    </el-button>
                    <el-button size="large" v-if="!isEditMode" :disabled="processing" @click="enableEditMode">
                        <svg class="i-lucide-edit"></svg>
                        {{ $t("editStack") }}
                    </el-button>
                    <el-button size="large" v-if="!isEditMode && !active" class="btn btn-primary" :disabled="processing"
                        @click="startStack">
                        <svg class="i-lucide-play"></svg>
                        {{ $t("startStack") }}
                    </el-button>
                    <el-button size="large" v-if="!isEditMode && active" class="btn btn-normal " :disabled="processing"
                        @click="restartStack">
                        <svg class="i-lucide-refresh-ccw"></svg>
                        {{ $t("restartStack") }}
                    </el-button>
                    <el-button size="large" v-if="!isEditMode" class="btn btn-normal" :disabled="processing"
                        @click="updateStack">
                        <svg class="i-lucide-refresh-ccw"></svg>
                        {{ $t("updateStack") }}
                    </el-button>

                    <el-button size="large" v-if="!isEditMode && active" :disabled="processing" @click="stopStack">
                        <svg class="i-lucide-circle-stop"></svg>
                        {{ $t("stopStack") }}
                    </el-button>
                    <el-button size="large" v-if="!isEditMode && active" :disabled="processing" @click="downStack">
                        <svg class="i-lucide-circle-stop"></svg>
                        {{ $t("downStack") }}
                    </el-button>
                </el-button-group>
                <el-button-group class="ml-1px " round>
                    <el-button size="large" round v-if="isEditMode && !isAdd" :disabled="processing"
                        @click="discardStack">
                        {{ $t("discardStack") }}
                    </el-button>
                    <el-button size="large" round type="danger" v-if="!isEditMode" :disabled="processing"
                        @click="showDeleteDialog = !showDeleteDialog">
                        <svg class="i-lucide-trash"></svg>
                        {{ $t("deleteStack") }}
                    </el-button>
                </el-button-group>
            </div>

            <!-- URLs -->
            <!-- <div v-if="urls.length > 0" class="mb-3">
                <a v-for="(url, index) in urls" :key="index" target="_blank" :href="url.url">
                    <span class="badge bg-secondary me-2">{{ url.display }}</span>
                </a>
            </div> -->

            <!-- Progress Terminal -->
            <!-- <transition name="slide-fade" appear>
                <Terminal v-show="showProgressTerminal" ref="progressTerminal" class="mb-3 h-200px" :name="terminalName"
                    :endpoint="endpoint" :rows="progressTerminalRows">
                </Terminal>
            </transition> -->

            <div v-if="stack.isManaged" class="flex flex-row">
                <div class="lg:w-1/2">
                    <!-- General -->
                    <div v-if="isAdd">
                        <h4 class="mb-3">{{ $t("General") }}</h4>
                        <div class="shadow-[0_15px_70px_rgba(0,0,0,0.1)] rounded-lg p-4 mb-3 ">
                            <!-- Stack Name -->
                            <div>
                                <label for="name">{{ $t("stackName") }}</label>
                                <el-input v-model="stack.stackName" type="text" @blur="stackNameToLowercase"
                                    class="[&_.el-input\_\_wrapper]:rounded-5">
                                </el-input>
                                <div class="text-sm color-gray">
                                    {{ $t("Lowercase only") }}
                                </div>
                            </div>
                        </div>
                    </div>

                    <!-- Containers -->
                    <h4 class="mb-3">{{ $t("container", 2) }}</h4>
                    <div v-if="isEditMode" class="flex mb-3">
                        <el-input v-model="newContainerName" :placeholder="$t(`New Container Name...`)"
                            class="[&_.el-input\_\_wrapper]:rounded-l-5 [&_.el-input\_\_wrapper]:rounded-r-0" />
                        <el-button type="primary" @click="addContainer" class="rounded-l-0! rounded-r-5!">
                            {{ $t("addContainer") }}
                        </el-button>
                    </div>
                    <div ref="containerList">
                         <!-- :service="envSubstJSONConfig.services[name]" -->
                        <Container v-for="(_, name) in jsonConfig.services" :key="name" :name="`${name}`"
                            :is-edit-mode="isEditMode" :first="name === Object.keys(jsonConfig.services)[0]"
                            :status="serviceStatusList[name]"
                            :stackName="stack.stackName">
                        </Container>
                    </div>

                    <!-- General -->
                    <!-- 额外 -->
                    <!-- <div v-if="isEditMode">
                        <h4 class="mb-3">{{ $t("extra") }}</h4>
                        <div class="p-4 mb-3">
                            <div class="mb-4">
                                <label>{{ $t("url") }}</label>
                                <ArrayInput name="urls" :display-name="$t('url')" placeholder="https://"
                                    object-type="x-dockge" >
                                </ArrayInput>
                            </div>
                        </div>
                    </div> -->
                </div>
                <div class="lg:w-1/2">
                    <h4 class="mb-3">{{ stack.yamlPath.split("/")[-1] }}</h4>
                    <!-- YAML editor -->
                    <div class="mb-3" :class="{ 'edit-mode': isEditMode }">
                        <YamlEditor v-model="stack.yamlContent" v-model:editable="isEditMode" class="max-h-400px">
                        </YamlEditor>
                    </div>
                    <div v-if="isEditMode" class="mb-3">
                        {{ yamlError }}
                    </div>

                    <!-- ENV editor -->
                    <div v-if="isEditMode">
                        <h4 class="mb-3">.env</h4>
                        <div class="mb-3" :class="{ 'edit-mode': isEditMode }">
                            <YamlEditor v-model="stack.envContent"></YamlEditor>
                        </div>
                    </div>

                    <div v-if="isEditMode">
                        <!-- Volumes -->
                        <!-- <div v-if="false">
                            <h4 class="mb-3">{{ $tc("volume", 2) }}</h4>
                            <div class="shadow-box big-padding mb-3">
                            </div>
                        </div> -->

                        <!-- Networks -->
                        <h4 class="mb-3">{{ $t("Network", 2) }}</h4>
                        <div class="shadow-box big-padding mb-3">
                            <NetworkInput />
                        </div>
                    </div>

                    <!-- <div class="shadow-box big-padding mb-3">
                        <div class="mb-3">
                            <label for="name" class="form-label"> Search Templates</label>
                            <input id="name" v-model="name" type="text" class="form-control" placeholder="Search..." required>
                        </div>

                        <prism-editor v-if="false" v-model="yamlConfig" class="yaml-editor" :highlight="highlighter" line-numbers @input="yamlCodeChange"></prism-editor>
                    </div>-->
                </div>
            </div>

            <div v-if="!stack.isManaged && !processing">
                {{ $t("stackNotManaged") }}
            </div>

            <!-- Delete Dialog -->
            <el-dialog v-model="showDeleteDialog">
                <span> {{ $t("deleteStackMsg") }} </span>
                <template #footer>
                    <el-button type="primary" @click="showDeleteDialog = false"> {{ $t("Cancel") }}</el-button>
                    <el-button type="danger" @click="deleteStack">{{ $t("Confirm") }}</el-button>
                </template>
            </el-dialog>
            <!-- <template> -->
            <div ref="iTerminalRef" class="w-full h-full bg-red-7"></div>
            <!-- </template> -->
        </div>
    </transition>
</template>

<script setup lang="ts">

import "@xterm/xterm/css/xterm.css"

// terminal 测试
import { Terminal } from "@xterm/xterm";
import { AttachAddon } from '@xterm/addon-attach'
const iTerminalRef = ref<HTMLElement>()

const terminal = ref<Terminal>()

// const terminalFitAddOn = ref()

onMounted(() =>{
    terminal.value = new Terminal({
        fontSize: 14,
        cursorBlink: true,
        cols: 105,
        rows: 10,
    })

    const ws = new WebSocket("ws://127.0.0.1:9527/ws/terminal?container=xxx-nginx-3-1")
    ws.onclose = () => {
        console.log("服务器跑了")
    }
    const attachAddon = new AttachAddon(ws)
    terminal.value.loadAddon(attachAddon)
    terminal.value.open(iTerminalRef.value!)

})




import {
    COMBINED_TERMINAL_COLS,
    COMBINED_TERMINAL_ROWS,
    getCombinedTerminalName,
    getComposeTerminalName,
    PROGRESS_TERMINAL_ROWS,
} from "@/utils/terminal"
import NetworkInput from "../components/NetworkInput.vue"
import dotenv from "dotenv"
import { envSubstYAML, yamlToJSON, copyYAMLComments } from "@/utils/yaml"
import { type ListItem, type StackStatusItem } from "@/types/stack"
import YamlEditor from "@/components/YamlEditor.vue"
import { deleteStackAPI, deployStackAPI, downStackAPI, getStackAPI, getStackStatusAPI, reStartStackAPI, saveStackAPI, startStackAPI, stopStackAPI, updateStackAPI, } from "@/api/stack"

import { Document } from "yaml"

// onBeforeRouteUpdate((to, from, next) => {
//     exitConfirm(next)
// })

// onBeforeRouteLeave((to, from, next) => {
//     exitConfirm(next)
// })

const defaultTemplate = `services:
  nginx:
    image: nginx:latest
    restart: unless-stopped
    ports:
      - "8080:80"
    environment:
      - DEBUG=1
      - PROD
      - A=\${a}
      - C=\${c}
`
const defaultEnv = "# VARIABLE=value #comment\na=b\nc=d\ne=false"

// let yamlErrorTimeout: ReturnType<typeof setTimeout>

let serviceStatusTimeout: ReturnType<typeof setTimeout>

const route = useRoute()
const router = useRouter()

const isEditMode = ref(false)




// const exitConfirm = (next) => {
//     if (isEditMode.value) {
//         if (confirm("退出")) {
//             exitAction()
//             next()
//         } else {
//             next(false)
//         }
//     } else {
//         exitAction()
//         next()
//     }
// }

const stack = ref<ListItem>({} as ListItem)
const stopServiceStatusTimeout = ref(false)

const exitAction = () => {
    stopServiceStatusTimeout.value = true
    clearTimeout(serviceStatusTimeout)

    // Leave Combined Terminal
    console.debug("leaveCombinedTerminal", endpoint.value, stack.value.stackName)
    // this.$root.emitAgent(endpoint.value, "leaveCombinedTerminal", stack.value.stackName, () => { })
}

const startServiceStatusTimeout = () => {
    clearTimeout(serviceStatusTimeout)
    serviceStatusTimeout = setTimeout(async () => {
        requestServiceStatus()
    }, 5000)
}

const submitted = ref(false)

// 是否新增
const isAdd = computed(() => {
    return route.path === "/compose" && !submitted.value
})

const processing = ref(true)

// status of containers in stack
// const serviceStatusList = ref({} as Record<string, string>)
const serviceStatusList = ref({} as Record<string, StackStatusItem>)

// const readOnly = ref(false)


onMounted(() => {
    if (isAdd.value) {
        processing.value = false
        isEditMode.value = true

        let yamlContent = defaultTemplate
        let composeENV = defaultEnv

        // if (this.$root.composeTemplate) {
        //     composeYAML = this.$root.composeTemplate
        //     this.$root.composeTemplate = ""
        // } else {
        //     composeYAML = defaultTemplate
        // }
        // if (this.$root.envTemplate) {
        //     composeENV = this.$root.envTemplate
        //     this.$root.envTemplate = ""
        // } else {
        //     composeENV = defaultEnv
        // }

        // Default Values
        stack.value = {
            id: "",
            stackName: "",
            stackStatus: "running(1)",
            yamlPath: "",
            yamlContent: "",
            envContent: "",
            isManaged: true,
            // endpoint: "",
        }
        yamlCodeChange()

    } else {
        // console.log(route.params)
        stack.value.stackName = route.params.stackName[0]
        loadStack()
    }

    requestServiceStatus()
})

const requestServiceStatus = async () => {
    if (isAdd.value) {
        return
    }
    const stackName = route.params.stackName as string

    const res = await getStackStatusAPI({ stackName })
    serviceStatusList.value = res.items.reduce((object, item) => {
        object[item.service] = item
        return object
    }, {} as Record<string, StackStatusItem>)


    // serviceStatusList.value =  res.items.reduce((object, item) => {
    //     object[item.service] =  item.state
    //     return object
    // }, {} as Record<string, string>)
    // console.log(serviceStatusList, res.items, "list")
    // TODO 请求 status
    // this.$root.emitAgent(this.endpoint, "serviceStatusList", this.stack.name, (res) => {
    //     if (res.ok) {
    //         this.serviceStatusList = res.serviceStatusList
    //     }
    //     if (!this.stopServiceStatusTimeout) {
    //         this.startServiceStatusTimeout()
    //     }
    // })
}

const enableEditMode = () => {
    isEditMode.value = true
}

// stack 操作
const loadStack = async () => {
    processing.value = true
    const stackName = route.params.stackName as string
    const res = await getStackAPI({ stackName })
    stack.value = res
    yamlCodeChange()
    processing.value = false
}

const deployStack = async () => {
    processing.value = true
    if (!jsonConfig.value.services) {
        ElMessage.error("No services found in compose.yaml")
        processing.value = false
        return
    }
    // Check if services is object
    if (typeof jsonConfig.value.services !== "object") {
        ElMessage.error("Services must be an object")
        processing.value = false
        return
    }
    let serviceNameList = Object.keys(jsonConfig.value.services)
    // Set the stack name if empty, use the first container name
    if (stack.value.stackName && serviceNameList.length > 0) {
        let serviceName = serviceNameList[0]
        let service = jsonConfig.value.services[serviceName]

        if (service && service.container_name) {
            stack.value.stackName = service.container_name
        } else {
            stack.value.stackName = serviceName
        }
    }

    // TODO
    bindTerminal()

    await deployStackAPI({
        stackName: stack.value.stackName,
        yamlContent: stack.value.yamlContent,
        envContent: stack.value.envContent
    })
    processing.value = true
    ElMessage.success("部署成功")
    isEditMode.value = false
    // TODO ?
    router.push(route.path)
}

const downStack = async () => {
    processing.value = true
    await downStackAPI({ stackName: stack.value.stackName })
    processing.value = false
    ElMessage.success("删除成功")
}

const startStack = async () => {
    processing.value = true
    await startStackAPI({ stackName: stack.value.stackName })
    processing.value = false
    ElMessage.success("启动成功")
}

const stopStack = async () => {
    processing.value = true
    await stopStackAPI({ stackName: stack.value.stackName })
    processing.value = false
    ElMessage.success("停止成功")
}

const saveStack = async () => {
    processing.value = true
    await saveStackAPI({
        stackName: stack.value.stackName,
        yamlContent: stack.value.yamlContent,
        envContent: stack.value.envContent,
        yamlPath: stack.value.yamlPath,
    })
    processing.value = false
    ElMessage.success("保存成功")
}

const restartStack = async () => {
    processing.value = true
    await reStartStackAPI({ stackName: stack.value.stackName })
    processing.value = false
    ElMessage.success("重启成功")
}

const updateStack = async () => {
    processing.value = true
    await updateStackAPI({ stackName: stack.value.stackName })
    processing.value = false
    ElMessage.success("更新成功")
}

const deleteStack = async () => {
    processing.value = true
    await deleteStackAPI({ stackName: stack.value.stackName })
    processing.value = false
    ElMessage.success("删除成功")
    setTimeout(() => {
        router.push("/")
    }, 300)
}

const discardStack = () => {
    loadStack()
    isEditMode.value = false
}

const yamlDoc = ref({} as Document)
const yamlError = ref("")

interface JsonConfig {
    version: string
    services: {
        [key: string]: {
            container_name?: string
            image?: string
            restart?: string
            ports?: string[]
            volumes?: string[]
            environment?: string[]
            networks?: string[]
        }
    }
    networks: {
        [key: string]: {
            driver: string
        }
    }
}

const jsonConfig = ref<JsonConfig>({} as JsonConfig)
const envSubstJSONConfig = ref<JsonConfig>({} as JsonConfig)
// 编辑器失焦点
const editorFocus = ref(false)

const yamlCodeChange = () => {
    try {
        let { config, doc } = yamlToJSON(stack.value.yamlContent)
        yamlDoc.value = doc
        jsonConfig.value = config
        let env = dotenv.parse(stack.value.envContent)
        // 使用 env 替换 compose 文件中, 类似 A=${A} 的环境声明
        let envYAML = envSubstYAML(stack.value.yamlContent, env)
        envSubstJSONConfig.value = yamlToJSON(envYAML).config
        console.log(envSubstJSONConfig.value.services["nginx-1"].ports, "envSubstJSONConfig.value")
        yamlError.value = ""
    } catch (e: any) {
        yamlError.value = e.message
    }
}

// 监听主要的几个内容变化
watch(
    () => stack.value.yamlContent, () => {
        console.debug("yaml code changed")
        yamlCodeChange()
    }
)

watch(
    () => stack.value.envContent, () => {
        console.debug("env code changed")
        yamlCodeChange()
    }
)

watch(
    jsonConfig, () => {
        console.debug("jsonConfig changed")
        if (editorFocus.value) {
            return
        }
        let doc = new Document(jsonConfig.value)
        if (yamlDoc.value) {
            copyYAMLComments(doc, yamlDoc.value)
        }
        stack.value.yamlContent = doc.toString()
        yamlDoc.value = doc
    }, { deep: true }
)


const stackNameToLowercase = () => {
    stack.value.stackName = stack.value.stackName.toLowerCase()
}

const newContainerName = ref("")

const addContainer = () => {
    if (jsonConfig.value.services[newContainerName.value]) {
        ElMessage.error("Container name already exists")
        return
    }

    if (!newContainerName.value) {
        ElMessage.error("Container name cannot be empty")
        return
    }

    jsonConfig.value.services[newContainerName.value] = {
        restart: "unless-stopped",
    }
    newContainerName.value = ""
    // let element = this.$refs.containerList.lastElementChild
    // element.scrollIntoView({
    //     block: "start",
    //     behavior: "smooth"
    // })
}

const bindTerminal = () => {
    // this.$refs.progressTerminal?.bind(this.endpoint, this.terminalName)
}

const status = computed(() => {
    return stack.value.stackStatus as string
})

const active = computed(() => {
    return status.value.includes("running")
})

const terminalName = computed(() => {
    if (!stack.value.stackName) {
        return ""
    }
    return getComposeTerminalName(endpoint.value, stack.value.stackName)
})

const combinedTerminalName = computed(() => {
    if (!stack.value.stackName) {
        return ""
    }
    return getCombinedTerminalName(endpoint.value, stack.value.stackName)
})

const networks = computed(() => {
    return jsonConfig.value.networks
})

const endpoint = computed(() => {
    return stack.value.endpoint || route.params.endpoint || ""
})

// const url = computed(() => {
//     if (stack.value.endpoint) {
//         return `/compose/${stack.value.stackName}/${stack.value.endpoint}`
//     } else {
//         return `/compose/${stack.value.stackName}`
//     }
// })

// const urls = computed(() => {

//     if (!envSubstJSONConfig.value["x-dockge"] || !envSubstJSONConfig.value["x-dockge"].urls || !Array.isArray(envSubstJSONConfig.value["x-dockge"].urls)) {
//         return []
//     }

//     let urls = []
//     for (const url of envSubstJSONConfig.value["x-dockge"].urls) {
//         let display
//         try {
//             let obj = new URL(url)
//             let pathname = obj.pathname
//             if (pathname === "/") {
//                 pathname = ""
//             }
//             display = obj.host + pathname + obj.search
//         } catch (e) {
//             display = url
//         }

//         urls.push({ display, url })
//     }
//     return urls
// })

const showProgressTerminal = ref(true)

const progressTerminalRows = PROGRESS_TERMINAL_ROWS
const combinedTerminalRows = COMBINED_TERMINAL_ROWS
const combinedTerminalCols = COMBINED_TERMINAL_COLS

const showDeleteDialog = ref(false)


</script>
