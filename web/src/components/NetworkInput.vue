<template>
    <div>
        <h5> {{ $t("Internal Networks") }} </h5>
        <ul>
            <li v-for="(networkRow, index) in networkList" :key="index">
                <el-input v-model="networkRow.key" type="text" :placeholder="$t(`Network name...`)">
                    <template #suffix>
                        <svg class="i-lucide-x h-16px w-16px" @click="remove(index)"></svg>
                    </template>
                </el-input>
            </li>
        </ul>

        <el-button class="mt-3 me-2" type="primary" @click="addField">{{ $t("addInternalNetwork") }}</el-button>

        <h5 class="mt-3"> {{ $t("External Networks") }} </h5>

        <div v-if="externalNetworkList.length === 0">
            {{ $t("No External Networks") }}
        </div>
        <div v-for="(item, index) in externalNetworkList" :key="index" class="my-3">
            <el-switch v-model="selectedExternalList[item.name]" @change="() => onChange(item)">
            </el-switch>
            <span>{{ item.name }}</span>

            <!-- <span v-if="false" class="text-danger ms-2 delete">Delete</span> -->
        </div>

        <!-- <div v-if="false" class="input-group mb-3">
            <input placeholder="New external network name..." class="form-control"
                @keyup.enter="createExternelNetwork" />
            <button class="btn btn-normal btn-sm  me-2" type="button">
                {{ $t("createExternalNetwork") }}
            </button>
        </div>

        <div v-if="false">
            <button class="btn btn-primary btn-sm mt-3 me-2" @click="applyToYAML">{{ $t("applyToYAML") }}</button>
        </div> -->
    </div>
</template>

<script setup lang="ts">
import { networkListAPI } from '@/api/network'
import type { NetworkItem } from '@/types/network'

const networkList = ref([])
const externalList = ref([])
const selectedExternalList = ref<Object>({})
const externalNetworkList = ref<NetworkItem[]>([])

const jsonConfig = computed(() => {
    // return this.$parent.$parent.jsonConfig

    return {
        networks: {
            "a": {},
        },

    }
})

const stack = computed(() => {
    // return this.$parent.$parent.stack
})

const editorFocus = computed(() => {
    // return this.$parent.$parent.editorFocus
})

const endpoint = computed(() => {
    // return this.$parent.$parent.endpoint
})

watch(networkList, (newList, oldList) => {
    // console.log(oldList, 234, newList, 123)
}, { deep: true })

watch(selectedExternalList, (newList, oldList) => {
    // console.log(oldList, 561, newList, 123)
}, { deep: true })

watchEffect(() => {
    // console.log(selectedExternalList, 55555)
    
})

const onChange = (item) => {
    console.log(selectedExternalList.value, 861)
}

onMounted(() => {
    loadNetworkList()
    loadExternalNetworkList()
})

const loadNetworkList = () => {
    // let networkList = []
    // let externalList = {}

    for (const key in jsonConfig.value.networks) {
        let obj = {
            key: key,
            value: jsonConfig.value.networks[key],
        }

        if (obj.value && obj.value.external) {
            externalList.value[key] = Object.assign({}, obj.value)
        } else {
            networkList.value.push(obj)
        }
    }

    // Restore selectedExternalList

    for (const networkName in externalList.value) {
        selectedExternalList.value[networkName] = true
    }
}

const loadExternalNetworkList = async () => {
    const res = await networkListAPI()
    externalNetworkList.value = res.items
    console.log(externalNetworkList.value)
}

const addField = () => {
    networkList.value.push({
        key: "",
        value: {},
    })
}

const remove = (index: number) => {
    networkList.value.splice(index, 1)
    applyToYAML()
}

const applyToYAML = () => {
    if (editorFocus.value) {
        return
    }

    jsonConfig.value.networks = {}

    // Internal networks
    for (const networkRow of networkList.value) {
        jsonConfig.value.networks[networkRow.key] = networkRow.value
    }

    // External networks
    for (const networkName in externalList.value) {
        jsonConfig.value.networks[networkName] = externalList.value[networkName]
    }

    console.debug("applyToYAML", jsonConfig.value.networks)
}

</script>
