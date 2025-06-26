<template>
    <div class="shadow-[0_15px_70px_rgba(0,0,0,0.1)] rounded-lg p-4 mb-3">
        <div class="flex flex-row">
            <div class="flex flex-col w-7/12">
                <h4>{{ name }}</h4>
                <div class="mb-2">
                    {{ props.status.image }}
                </div>
                <div v-if="!isEditMode">
                    <span class="m-auto w-30% me-1" :class="bgStyle">{{ props.status.state }}</span>
                    <span v-for="(port, index) in props.status.ports" :key="index" class="py-1 px-2 bg-red-3 rounded-xl">
                        {{ port.publishPort }}:{{ port.targetPort }}/{{ port.protocol }}
                    </span>
                </div>
            </div>
            <div class="flex w-5/12">
                <div class="flex content-center w-full h-full items-center justify-end">
                    <router-link v-if="!isEditMode" :to="terminalRouteLink" class="no-underline">
                        <svg class="i-lucide-terminal"></svg>
                        Bash
                    </router-link>
                </div>
            </div>
        </div>
        <div v-if="isEditMode" class="mt-2">
            <el-button type="primary" @click="showConfig = !showConfig" class="rounded-10!">
                <svg class="i-lucide-edit"></svg>
                {{ $t("Edit") }}
            </el-button>
            <el-button type="danger" @click="remove" class="rounded-10!">
                <svg class="i-lucide-trash-2"></svg>
                {{ $t("Delete") }}
            </el-button>
        </div>

        <transition name="slide-fade" appear>
            <div v-if="isEditMode && showConfig" class="mt-3">
                <!-- Image -->
                <div class="mb-4">
                    <label>
                        {{ $t("Image") }}
                    </label>
                    <div class="mb-3">
                        <el-input v-model="service.image"></el-input>
                    </div>

                    <!-- TODO: Search online: https://hub.docker.com/api/content/v1/products/search?q=louislam%2Fuptime&source=community&page=1&page_size=4 -->
                    <!-- <datalist id="image-datalist">
                        <option value="louislam/uptime-kuma:1" />
                    </datalist> -->
                    <!-- <div class="form-text"></div> -->
                </div>

                <!-- Ports -->
                <div class="mb-4">
                    <label>
                        {{ $t("Port", 2) }}
                    </label>
                    <ArrayInput name="ports" :display-name="$t('Port')" placeholder="HOST:CONTAINER" />
                </div>

                <!-- Volumes -->
                <div class="mb-4">
                    <label class="form-label">
                        {{ $t("Volume", 2) }}
                    </label>
                    <ArrayInput name="volumes" :display-name="$t('Volume')" placeholder="HOST:CONTAINER" />
                </div>

                <!-- Restart Policy -->
                <div class="mb-4">
                    <label class="form-label">
                        {{ $t("RestartPolicy") }}
                    </label>
                    <el-select v-model="service.restart" clearable @change="onSelect">
                        <el-option v-for="item in networkOptions" :key="item.value" :label="item.label"
                            :value="item.value">
                        </el-option>
                    </el-select>
                </div>

                <!-- Environment Variables -->
                <div class="mb-4">
                    <label class="form-label">
                        {{ $t("EnvironmentVariable", 2) }}
                    </label>
                    <ArrayInput name="environment" :display-name="$t('EnvironmentVariable')" placeholder="KEY=VALUE" />
                </div>

                <!-- Container Name -->
                <div v-if="false" class="mb-4">
                    <label class="form-label">
                        {{ $t("ContainerName") }}
                    </label>
                    <div class="input-group mb-3">
                        <input v-model="service.container_name" />
                    </div>
                    <div></div>
                </div>

                <!-- Network -->
                <div class="mb-4">
                    <label class="form-label">
                        {{ $t("Network", 2) }}
                    </label>

                    <div v-if="networkList.length === 0 && service.networks && service.networks.length > 0"
                        class="text-warning mb-3">
                        {{ $t("NoNetworksAvailable") }}
                    </div>

                    <ArraySelect name="networks" :display-name="$t('Network')" placeholder="Network Name"
                        :options="networkList" />
                </div>

                <!-- Depends on -->
                <div class="mb-4">
                    <label class="form-label">
                        {{ $t("DependsOn") }}
                    </label>
                    <ArrayInput name="depends_on" :display-name="$t('DependsOn')" :placeholder="$t(`ContainerName`)" />
                </div>
            </div>
        </transition>
    </div>
</template>

<script setup lang="ts">

import { parseDockerPort } from "@/utils/container"

import { type StackStatusItem } from "@/types/stack"
import type { PropType } from "vue"

const { t } = useI18n()

const showConfig = ref(false)

const props = defineProps({
    name: {
        type: String,
        required: true,
    },
    stackName: {
        type: String,
        required: true,
    },
    isEditMode: {
        type: Boolean,
        default: false,
    },
    first: {
        type: Boolean,
        default: false,
    },
    status: {
        type: Object as PropType<StackStatusItem>,
        default: {} as StackStatusItem,
    },
    // service: {
    //     type: Object,
    //     default: {}
    // },
    endpoint: {
        type: String,
        default: "",
        required: false,
    },
})

onMounted(() => {
    console.log(props.status, 1231412312)
})

const onSelect = (e: any) => {
    console.log(e, 123)
    service.value.restart = e
    console.log(service.value, 421)
}

const networkOptions = [
    {
        label: t("RestartPolicyAlways"),
        value: "always",
    },
    {
        label: t("RestartPolicyUnlessStopped"),
        value: "unless-stopped",
    },
    {
        label: t("RestartPolicyOnFailure"),
        value: "on-failure",
    },
    {
        label: t("RestartPolicyNo"),
        value: "no",
    },
]

const networkList = computed(() => {
    let list = []
    for (const networkName in jsonObject.value.networks) {
        list.push(networkName)
    }
    return list
})

const bgStyle = computed(() => {
    if (status === "running" || status === "healthy") {
        return "bg-primary"
    } else if (status === "unhealthy") {
        return "bg-danger"
    } else {
        return "bg-secondary"
    }
})


const remove = () => {
    // delete this.jsonObject.services[this.name]
}

const terminalRouteLink = computed(() => {
    if (props.endpoint) {
        return {
            name: "containerTerminalEndpoint",
            params: {
                endpoint: props.endpoint,
                stackName: props.stackName,
                serviceName: props.name,
                type: "bash",
            },
        }
    } else {
        return {
            name: "containerTerminal",
            params: {
                stackName: props.stackName,
                serviceName: props.name,
                type: "bash",
            },
        }
    }
})

const envSubstJSONConfig = computed(() => {
    // return this.$parent.$parent.envsubstJSONConfig;
    return {
        "services": {}
    }
})

const envSubstService = computed(() => {
    if (!envSubstJSONConfig.value.services[props.name]) {
        return {}
    }
    return envSubstJSONConfig.value.services[props.name]
})

const imageName = computed(() => {
    if (envSubstService.value.image) {
        return envSubstService.value.image.split(":")[0]
    } else {
        return ""
    }
})
const imageTag = computed(() => {
    if (envSubstService.value.image) {
        let tag = envSubstService.value.image.split(":")[1]

        if (tag) {
            return tag
        } else {
            return "latest"
        }
    } else {
        return ""
    }
})


const endpoint = computed(() => {
    // return this.$parent.$parent.endpoint
})

const stack = computed(() => {
    // return this.$parent.$parent.stack
    return {}
})

const stackName = computed(() => {
    // return this.$parent.$parent.stack.name
})

const service = computed(() => {
    return { image: "", restart: 'no' }
    if (!jsonObject.value.services[props.name]) {
        return {}
    }
    return jsonObject.value.services[props.name]
})

const jsonObject = computed(() => {
    // return this.$parent.$parent.jsonConfig
    return {
        services: {
            "ports": {},
            "networks": {},
            "depends_on": [],
            "environment": [],
            "volumes": {},
        }
    }
})

const parsePort = (port) => {
    // if (props.stack.endpoint) {
    // return parseDockerPort(port, props.stack.primaryHostname);
    // } else {
    // let hostname = this.$root.info.primaryHostname || location.hostname;
    // return parseDockerPort(port, hostname);
    // }
}



</script>
