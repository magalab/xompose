<template>
    <transition name="slide-fade" appear>
        <div>
            <h1 class="mb-3">{{ $t("Terminal") }} - {{ serviceName }} ({{ stackName }})</h1>
            <div class="mb-3">
                <router-link :to="sh" >{{ $t("Switch to sh") }}</router-link>
            </div>

            <Terminal :rows="20" :cols="80" :name="terminalName" :stack-name="stackName" :service-name="serviceName" :type="type"
                :endpoint="endpoint">
            </Terminal>
        </div>
    </transition>
</template>

<script setup lang="ts">

const props = defineProps({
    stackName: String,
    endpoint: {
        type: String,
        default: "",
    },
    type: String,
    serviceName: String,

})

const terminalName = computed(() => {
    return `"container-exec-" + ${props.endpoint} + "-" + ${props.stackName} + "-" + ${props.serviceName}`

})

const sh = computed(() => {
    let data = {
        name: "containerTerminal",
        params: {
            endpoint: props.endpoint || "",
            stackName: props.stackName,
            serviceName: props.serviceName,
            type: props.type,
        }
    }

    return data
})

</script>

<style scoped lang="scss">
.terminal {
    height: 410px;
}
</style>
