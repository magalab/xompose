<template>
    <transition name="slide-fade" appear>
        <div>
            <h1 v-if="isAdd" class="mb-3 text-xl fw-700"> Compose </h1>
            <h1 v-else class="mb-3">
                <Uptime :stack="globalStack" :pill="true" /> {{ stack.name }}
            </h1>

            <div v-if="stack.isManaged" class="mb-3">
                <div class="btn-group me-2" role="group">
                    <button v-if="isEditMode" class="btn btn-primary" :disabled="processing" @click="deployStack">
                        <svg class="i-lucide-rocket"></svg>
                        {{ $t("deployStack") }}
                    </button>

                    <button v-if="isEditMode" class="btn btn-normal" :disabled="processing" @click="saveStack">
                        <svg class="i-lucide-save"></svg>
                        {{ $t("saveStackDraft") }}
                    </button>

                    <button v-if="!isEditMode" class="btn btn-secondary" :disabled="processing" @click="enableEditMode">
                        <svg class="i-lucide-edit"></svg>
                        {{ $t("editStack") }}
                    </button>

                    <button v-if="!isEditMode && !active" class="btn btn-primary" :disabled="processing"
                        @click="startStack">
                        <svg class="i-lucide-play"></svg>
                        {{ $t("startStack") }}
                    </button>

                    <button v-if="!isEditMode && active" class="btn btn-normal " :disabled="processing"
                        @click="restartStack">
                        <svg class="i-lucide-refresh-ccw"></svg>
                        {{ $t("restartStack") }}
                    </button>

                    <button v-if="!isEditMode" class="btn btn-normal" :disabled="processing" @click="updateStack">
                        <svg class="i-lucide-refresh-ccw"></svg>
                        {{ $t("updateStack") }}
                    </button>

                    <button v-if="!isEditMode && active" class="btn btn-normal" :disabled="processing"
                        @click="stopStack">
                        <svg class="i-lucide-circle-stop"></svg>
                        {{ $t("stopStack") }}
                    </button>

                    <!-- <BDropdown right text="" variant="normal">
                        <BDropdownItem @click="downStack">
                            <font-awesome-icon icon="stop" class="me-1" />
                            {{ $t("downStack") }}
                        </BDropdownItem>
                    </BDropdown> -->
                </div>

                <button v-if="isEditMode && !isAdd" class="btn btn-normal" :disabled="processing" @click="discardStack">
                    {{ $t("discardStack") }}
                </button>
                <button v-if="!isEditMode" class="btn btn-danger" :disabled="processing"
                    @click="showDeleteDialog = !showDeleteDialog">
                    <svg class="i-lucide-trash"></svg>
                    {{ $t("deleteStack") }}
                </button>
            </div>

            <!-- URLs -->
            <div v-if="urls.length > 0" class="mb-3">
                <a v-for="(url, index) in urls" :key="index" target="_blank" :href="url.url">
                    <span class="badge bg-secondary me-2">{{ url.display }}</span>
                </a>
            </div>

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
                    <!-- <h4 class="mb-3">{{ $tc("container", 2) }}</h4> -->
                    <h4 class="mb-3">{{ $t("container") }}</h4>
                    <div v-if="isEditMode" class="flex mb-3">
                        <el-input v-model="newContainerName" :placeholder="$t(`New Container Name...`)"
                            class="[&_.el-input\_\_wrapper]:rounded-l-5 [&_.el-input\_\_wrapper]:rounded-r-0" />
                        <el-button type="primary" @click="addContainer" class="rounded-l-0! rounded-r-5!">
                            {{ $t("addContainer") }}
                        </el-button>
                    </div>
                    <div ref="containerList">
                        <Container v-for="(_, name) in jsonConfig.services" :key="name" :name="name"
                            :is-edit-mode="isEditMode" :first="name === Object.keys(jsonConfig.services)[0]"
                            :status="serviceStatusList[name]">
                        </Container>
                    </div>

                    <!-- <el-button
                        v-if="false && isEditMode && jsonConfig.services && Object.keys(jsonConfig.services).length > 0"
                        class="btn btn-normal mb-3" @click="addContainer">
                        {{ $t("addContainer") }}
                    </el-button> -->

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

                    <!-- Combined Terminal Output -->
                    <!-- <div v-show="!isEditMode">
                        <h4 class="mb-3"> 终端 </h4>
                        <Terminal ref="combinedTerminal" class="mb-3 h-200px" :name="combinedTerminalName"
                            :endpoint="endpoint" :rows="combinedTerminalRows" :cols="combinedTerminalCols"
                            style="height: 315px">
                        </Terminal>
                    </div> -->
                </div>
                <div class="lg:w-1/2">
                    <h4 class="mb-3">{{ stack.composeFileName }}</h4>
                    <!-- YAML editor -->
                    <div class="mb-3" :class="{ 'edit-mode': isEditMode }">
                        <YamlEditor v-model="stack.composeYAML" class="max-h-400px"></YamlEditor>
                    </div>
                    <div v-if="isEditMode" class="mb-3">
                        {{ yamlError }}
                    </div>

                    <!-- ENV editor -->
                    <div v-if="isEditMode">
                        <h4 class="mb-3">.env</h4>
                        <div class="mb-3" :class="{ 'edit-mode': isEditMode }">
                            <YamlEditor v-model="stack.composeENV"></YamlEditor>
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
            <!-- <BModal v-model="showDeleteDialog" :cancelTitle="$t('cancel')" :okTitle="$t('deleteStack')" okVariant="danger" @ok="deleteDialog"> -->
            <!-- {{ $t("deleteStackMsg") }} -->
            <!-- </BModal> -->
        </div>
    </transition>
</template>

<script setup lang="ts">
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
import { RUNNING } from "@/types/stack"
import YamlEditor from "@/components/YamlEditor.vue"
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
`
const defaultEnv = "# VARIABLE=value #comment"

let yamlErrorTimeout = null

let serviceStatusTimeout = null

const route = useRoute()

const isEditMode = ref(false)



const exitConfirm = (next) => {
    console.log(next, isEditMode.value, 1231231)
    if (isEditMode.value) {
        if (confirm("退出")) {
            exitAction()
            next()
        } else {
            next(false)
        }
    } else {
        exitAction()
        next()
    }
}

interface Stack {
    stackName: string,
    composeFileName: string,
    stackStatus: string,
    composeYAML: string,
    isManaged: boolean,
    composeENV: string,
    endpoint: string,
}

const stack = ref<Stack>({} as Stack)

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


onMounted(() => {
    if (isAdd.value) {
        processing.value = false
        isEditMode.value = true

        let composeYAML = defaultTemplate
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
            stackName: "",
            stackStatus: "running(1)",
            composeFileName: "",
            composeYAML,
            composeENV,
            isManaged: true,
            endpoint: "",
        }

        yamlCodeChange()

    } else {
        // console.log(route.params)
        stack.value.stackName = route.params.stackName[0]
        loadStack()
    }

    requestServiceStatus()
})

const requestServiceStatus = () => {
    if (isAdd.value) {
        return
    }
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
const loadStack = () => {
    processing.value = true
    // TODO 获取 stack 信息
    // this.$root.emitAgent(this.endpoint, "getStack", this.stack.name, (res) => {
    //     if (res.ok) {
    //         this.stack = res.stack
    //         this.yamlCodeChange()
    //         this.processing = false
    //         this.bindTerminal()
    //     } else {
    //         this.$root.toastRes(res)
    //     }
    // })
}

const deployStack = () => {
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

    bindTerminal()

    // TODO 部署
    // this.$root.emitAgent(this.stack.endpoint, "deployStack", this.stack.name, this.stack.composeYAML, this.stack.composeENV, this.isAdd, (res) => {
    //     this.processing = false
    //     this.$root.toastRes(res)

    //     if (res.ok) {
    //         this.isEditMode = false
    //         this.$router.push(this.url)
    //     }
    // })
}

const downStack = () => {
    processing.value = true
    // TODO 停止
    // this.$root.emitAgent(this.endpoint, "downStack", this.stack.name, (res) => {
    //     this.processing = false
    //     this.$root.toastRes(res)
    // })
}

const startStack = () => {
    processing.value = true

    // TODO
    // this.$root.emitAgent(this.endpoint, "startStack", this.stack.name, (res) => {
    //     this.processing = false
    //     this.$root.toastRes(res)
    // })
}

const stopStack = () => {
    processing.value = true

    // TODO 
    // this.$root.emitAgent(this.endpoint, "stopStack", this.stack.name, (res) => {
    //     this.processing = false
    //     this.$root.toastRes(res)
    // })
}

const saveStack = () => {
    processing.value = true

    // TODO
    // this.$root.emitAgent(this.stack.endpoint, "saveStack", this.stack.name, this.stack.composeYAML, this.stack.composeENV, this.isAdd, (res) => {
    //     this.processing = false
    //     this.$root.toastRes(res)

    //     if (res.ok) {
    //         this.isEditMode = false
    //         this.$router.push(this.url)
    //     }
    // })
}

const restartStack = () => {
    processing.value = true

    // this.$root.emitAgent(this.endpoint, "restartStack", this.stack.name, (res) => {
    //     this.processing = false
    //     this.$root.toastRes(res)
    // })
}

const updateStack = () => {
    processing.value = true

    // this.$root.emitAgent(this.endpoint, "updateStack", this.stack.name, (res) => {
    //     this.processing = false
    //     this.$root.toastRes(res)
    // })
}

const deleteDialog = () => {
    // TODO
    // this.$root.emitAgent(this.endpoint, "deleteStack", this.stack.name, (res) => {
    //     this.$root.toastRes(res)
    //     if (res.ok) {
    //         this.$router.push("/")
    //     }
    // })
}

const discardStack = () => {
    loadStack()
    isEditMode.value = false
}

const yamlDoc = ref({})
const yamlError = ref("")

const envSubstJSONConfig = ref({})

const jsonConfig = ref({})
const editorFocus = ref(false)

const yamlCodeChange = () => {
    try {
        let { config, doc } = yamlToJSON(stack.value.composeYAML)

        yamlDoc.value = doc
        jsonConfig.value = config

        let env = dotenv.parse(stack.value.composeENV)
        let envYAML = envSubstYAML(stack.value.composeYAML, env)
        envSubstJSONConfig.value = yamlToJSON(envYAML).config

        clearTimeout(yamlErrorTimeout)
        yamlError.value = ""
    } catch (e) {
        clearTimeout(yamlErrorTimeout)

        if (yamlError.value) {
            yamlError.value = e.message

        } else {
            yamlErrorTimeout = setTimeout(() => {
                yamlError.value = e.message
            }, 3000)
        }
    }
}

// watch: {
//     "stack.composeYAML": {
//         handler() {
//             if (this.editorFocus) {
//                 console.debug("yaml code changed")
//                 this.yamlCodeChange()
//             }
//         },
//         deep: true,
//         },

//     "stack.composeENV": {
//         handler() {
//             if (this.editorFocus) {
//                 console.debug("env code changed")
//                 this.yamlCodeChange()
//             }
//         },
//         deep: true,
//         },

//     jsonConfig: {
//         handler() {
//             if (!this.editorFocus) {
//                 console.debug("jsonConfig changed")

//                 let doc = new Document(jsonConfig.value)

//                 // Stick back the yaml comments
//                 if (yamlDoc.value) {
//                     copyYAMLComments(doc, yamlDoc.value)
//                 }

//                 stack.value.composeYAML = doc.toString()
//                 yamlDoc.value = doc
//             }
//         },
//         deep: true,
//         },

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

const globalStack = computed(() => {
    // return this.$root.completeStackList[this.stack.name + "_" + this.endpoint]
    return stack
})

const status = computed(() => {
    return globalStack?.value.stackStatus
})

const active = computed(() => {
    return status === RUNNING
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

const url = computed(() => {
    if (stack.value.endpoint) {
        return `/compose/${stack.value.stackName}/${stack.value.endpoint}`
    } else {
        return `/compose/${stack.value.stackName}`
    }
})

const urls = computed(() => {

    if (!envSubstJSONConfig.value["x-dockge"] || !envSubstJSONConfig.value["x-dockge"].urls || !Array.isArray(envSubstJSONConfig.value["x-dockge"].urls)) {
        return []
    }

    let urls = []
    for (const url of envSubstJSONConfig.value["x-dockge"].urls) {
        let display
        try {
            let obj = new URL(url)
            let pathname = obj.pathname
            if (pathname === "/") {
                pathname = ""
            }
            display = obj.host + pathname + obj.search
        } catch (e) {
            display = url
        }

        urls.push({ display, url })
    }
    return urls
})

const showProgressTerminal = ref(false)

const progressTerminalRows = PROGRESS_TERMINAL_ROWS
const combinedTerminalRows = COMBINED_TERMINAL_ROWS
const combinedTerminalCols = COMBINED_TERMINAL_COLS

const showDeleteDialog = ref(false)

const serviceStatusList = ref({})
</script>


<style lang="scss" scoped>
//  :deep(.el-input__wrapper) {
//     border-top-left-radius: 20px;
//     border-bottom-left-radius: 20px;
//     border-top-right-radius: 0;
//     border-bottom-right-radius: 0;
//     // border: 0;
//     // box-shadow: 0 0 0 0px;
// }</style>