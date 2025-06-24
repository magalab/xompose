<template>
    <div class="shadow-box big-padding mb-3">
        <div class="row">
            <div class="col-7">
                <h4>{{ name }}</h4>
                <div class="image mb-2">
                    <span class="me-1">{{ imageName }}:</span><span class="tag">{{ imageTag }}</span>
                </div>
                <div v-if="!isEditMode">
                    <span class="badge me-1" :class="bgStyle">{{ status }}</span>

                    <a v-for="port in envSubstService.ports" :key="port" :href="parsePort(port).url" target="_blank">
                        <span class="badge me-1 bg-secondary">{{ parsePort(port).display }}</span>
                    </a>
                </div>
            </div>
            <div class="col-5">
                <div class="function">
                    <router-link v-if="!isEditMode" class="btn btn-normal" :to="terminalRouteLink" disabled="">
                        <font-awesome-icon icon="terminal" />
                        Bash
                    </router-link>
                </div>
            </div>
        </div>

        <div v-if="isEditMode" class="mt-2">
            <button class="btn btn-normal me-2" @click="showConfig = !showConfig">
                <i-ep-edit />
                <!-- {{ $t("Edit") }} -->
                编辑
            </button>
            <!-- <button v-if="false" class="btn btn-normal me-2">Rename</button> -->
            <button class="btn btn-danger me-2" @click="remove">
                <i-ep-delete />
                删除容器
            </button>
        </div>

        <transition name="slide-fade" appear>
            <div v-if="isEditMode && showConfig" class="mt-3">
                <!-- Image -->
                <div class="mb-4">
                    <label class="form-label">
                        镜像
                    </label>
                    <div class="input-group mb-3">
                        <input v-model="service.image" class="form-control" list="image-datalist" />
                    </div>

                    <!-- TODO: Search online: https://hub.docker.com/api/content/v1/products/search?q=louislam%2Fuptime&source=community&page=1&page_size=4 -->
                    <datalist id="image-datalist">
                        <option value="louislam/uptime-kuma:1" />
                    </datalist>
                    <div class="form-text"></div>
                </div>

                <!-- Ports -->
                <div class="mb-4">
                    <label class="form-label">
                        <!-- {{ $tc("port", 2) }} -->
                        ports
                    </label>
                    <!-- <ArrayInput name="ports" :display-name="$t('port')" placeholder="HOST:CONTAINER" /> -->
                    <ArrayInput name="ports" display-name="port" placeholder="HOST:CONTAINER" />
                </div>

                <!-- Volumes -->
                <div class="mb-4">
                    <label class="form-label">
                        <!-- {{ $tc("volume", 2) }} -->
                        volumn
                    </label>
                    <!-- <ArrayInput name="volumes" :display-name="$t('volume')" placeholder="HOST:CONTAINER" /> -->
                    <ArrayInput name="volumes" display-name="volume" placeholder="HOST:CONTAINER" />
                </div>

                <!-- Restart Policy -->
                <div class="mb-4">
                    <label class="form-label">
                        <!-- {{ $t("restartPolicy") }} -->
                        restartPolicy
                    </label>
                    <select v-model="service.restart" class="form-select">
                        <option value="always">restartPolicyAlways </option>
                        <option value="unless-stopped">restartPolicyUnlessStopped</option>
                        <option value="on-failure">restartPolicyOnFailure</option>
                        <option value="no">restartPolicyNo</option>
                    </select>
                </div>

                <!-- Environment Variables -->
                <div class="mb-4">
                    <label class="form-label">
                        <!-- {{ $tc("environmentVariable", 2) }} -->
                        environmentVariable
                    </label>
                    <!-- <ArrayInput name="environment" :display-name="$t('environmentVariable')" placeholder="KEY=VALUE" /> -->
                    <ArrayInput name="environment" display-name="environmentVariable" placeholder="KEY=VALUE" />
                </div>

                <!-- Container Name -->
                <div v-if="false" class="mb-4">
                    <label class="form-label">
                        <!-- {{ $t("containerName") }} -->
                        containerName
                    </label>
                    <div class="input-group mb-3">
                        <input v-model="service.container_name" class="form-control" />
                    </div>
                    <div class="form-text"></div>
                </div>

                <!-- Network -->
                <div class="mb-4">
                    <label class="form-label">
                        <!-- {{ $tc("network", 2) }} -->
                        network
                    </label>

                    <div v-if="networkList.length === 0 && service.networks && service.networks.length > 0"
                        class="text-warning mb-3">
                        <!-- {{ $t("NoNetworksAvailable") }} -->
                        NoNetworksAvailable
                    </div>

                    <!-- <ArraySelect name="networks" :display-name="$t('network')" placeholder="Network Name" -->
                    <ArraySelect name="networks" display-name="network" placeholder="Network Name"
                        :options="networkList" />
                </div>

                <!-- Depends on -->
                <div class="mb-4">
                    <label class="form-label">
                        <!-- {{ $t("dependsOn") }} -->
                        dependsOn
                    </label>
                    <!-- <ArrayInput name="depends_on" :display-name="$t('dependsOn')" :placeholder="$t(`containerName`)" /> -->
                    <ArrayInput name="depends_on" display-name="dependsOn" placeholder="containerName" />
                </div>
            </div>
        </transition>
    </div>
</template>

<script setup lang="ts">

import { parseDockerPort } from "@/utils/container"

const showConfig = ref(false)

const props = defineProps({
    name: {
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
        type: String,
        default: "N/A",
    }
})

const networkList = computed(() => {
    let list = []
    for (const networkName in jsonObject.networks) {
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
                endpoint: this.endpoint,
                stackName: this.stackName,
                serviceName: this.name,
                type: "bash",
            },
        }
    } else {
        return {
            name: "containerTerminal",
            params: {
                stackName: this.stackName,
                serviceName: this.name,
                type: "bash",
            },
        }
    }
})

const envSubstJSONConfig = computed(() => {
    // return this.$parent.$parent.envsubstJSONConfig;
})

const envSubstService = computed(() => {
    if (!envSubstJSONConfig.services[props.name]) {
        return {}
    }
    return envSubstJSONConfig.services[props.name]
})

const imageName = computed(() => {
    if (envSubstService.image) {
        return envSubstService.image.split(":")[0]
    } else {
        return ""
    }
})
const imageTag = (() => {
    if (envSubstService.image) {
        let tag = envSubstService.image.split(":")[1]

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
})

const stackName = computed(() => {
    // return this.$parent.$parent.stack.name
})

const service = computed(() => {
    if (!jsonObject.services[props.name]) {
        return {}
    }
    return jsonObject.services[props.name]
})

const jsonObject = computed(() => {
    // return this.$parent.$parent.jsonConfig
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
