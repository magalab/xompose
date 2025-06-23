<template>
    <router-link :to="url" :class="{ 'dim': !stack.isManagedByDockge }" class="item">
        <Uptime :stack="stack" :fixed-width="true" class="me-2" />
        <div class="title">
            <span>{{ stackName }}</span>
            <!-- <div v-if="$root.agentCount > 1" class="endpoint">{{ endpointDisplay }}</div> -->
        </div>
    </router-link>
</template>

<script setup lang="ts">
import { type StackList } from "@/types/stack";
import Uptime from "./Uptime.vue";

const isCollapsed = ref(true)

const props = defineProps({
    stack: {
        type: Object,
        default: null,
    },
    isSelectMode: {
        type: Boolean,
        default: false,
    },
    isSelected: {
        type: Boolean,
        default: false,
    },
    depth: {
        type: Number,
        default: 0,
    },
    select: {
        type: Function,
        default: () => { }
    },
    deselect: {
        type: Function,
        default: () => { }
    },
})

const endpointDisplay = computed(() => {
    // return this.$root.endpointDisplayFunction(props.stack.endpoint);
})

const url = computed((): string => {
    if (props.stack.endpoint) {
        return `/compose/${props.stack.name}/${props.stack.endpoint}`;
    } else {
        return `/compose/${props.stack.name}`;
    }
})

const stackName = computed((): string => {
    return props.stack.name;
})

const changeCollapsed = (() => {
    isCollapsed.value = !isCollapsed.value
    let storage = window.localStorage.getItem("stackCollapsed");
    let storageObject = {};
    if (storage !== null) {
        storageObject = JSON.parse(storage);
    }
    storageObject[`stack_${props.stack.id}`] = isCollapsed.value;

    window.localStorage.setItem("stackCollapsed", JSON.stringify(storageObject));
})

const toggleSelection = () => {
    if (props.isSelected(props.stack.id)) {
        props.deselect(props.stack.id);
    } else {
        props.select(props.stack.id);
    }
}


</script>

<style lang="scss" scoped>
@import "../styles/vars.scss";

.small-padding {
    padding-left: 5px !important;
    padding-right: 5px !important;
}

.collapse-padding {
    padding-left: 8px !important;
    padding-right: 2px !important;
}

.item {
    text-decoration: none;
    display: flex;
    align-items: center;
    min-height: 52px;
    border-radius: 10px;
    transition: all ease-in-out 0.15s;
    width: 100%;
    padding: 5px 8px;

    &.disabled {
        opacity: 0.3;
    }

    &:hover {
        background-color: $highlight-white;
    }

    &.active {
        background-color: #cdf8f4;
    }

    .title {
        margin-top: -4px;
    }

    .endpoint {
        font-size: 12px;
        color: $dark-font-color3;
    }
}

.collapsed {
    transform: rotate(-90deg);
}

.animated {
    transition: all 0.2s $easing-in;
}

.select-input-wrapper {
    float: left;
    margin-top: 15px;
    margin-left: 3px;
    margin-right: 10px;
    padding-left: 4px;
    position: relative;
    z-index: 15;
}

.dim {
    opacity: 0.5;
}
</style>
