<template>
    <router-link :to="url"
        class="no-underline flex items-center min-h-52px rounded-10px transition-all duration-150 ease-in-out w-full px-8px py-[5px] disabled:opacity-30 hover:bg-[#e7faec] active:bg-[#cdf8f4]">
        <div class="mr-4px"> {{ props.stack.stackStatus }}</div>
        <div>{{ props.stack.stackName }}</div>
    </router-link>
</template>

<script setup lang="ts">
// import Uptime from "./Uptime.vue";

const props = defineProps({
    stack: {
        type: Object,
        default: null,
    },
    // isSelectMode: {
    //     type: Boolean,
    //     default: false,
    // },
    // isSelected: {
    //     type: Boolean,
    //     default: false,
    // },
    // depth: {
    //     type: Number,
    //     default: 0,
    // },
    // select: {
    //     type: Function,
    //     default: () => { }
    // },
    // deselect: {
    //     type: Function,
    //     default: () => { }
    // },
})

// const endpointDisplay = computed(() => {
//     // return this.$root.endpointDisplayFunction(props.stack.endpoint);
// })

const url = computed((): string => {
    if (props.stack.endpoint) {
        return `/compose/${props.stack.stackName}/${props.stack.endpoint}`;
    } else {
        return `/compose/${props.stack.stackName}`;
    }
})

// const stackName = computed((): string => {
//     return props.stack.stackName;
// })

// const changeCollapsed = (() => {
//     isCollapsed.value = !isCollapsed.value
//     let storage = window.localStorage.getItem("stackCollapsed");
//     let storageObject = {};
//     if (storage !== null) {
//         storageObject = JSON.parse(storage);
//     }
//     storageObject[`stack_${props.stack.id}`] = isCollapsed.value;

//     window.localStorage.setItem("stackCollapsed", JSON.stringify(storageObject));
// })

// const toggleSelection = () => {
//     if (props.isSelected(props.stack.id)) {
//         props.deselect(props.stack.id);
//     } else {
//         props.select(props.stack.id);
//     }
// }


</script>

<style lang="scss" scoped>
.dim {
    opacity: 0.5;
}
</style>
