<template>
    <span :class="className">{{ props.stack.stackStatus }}</span>
</template>

<script setup lang="ts">
import { statusColor } from '@/types/stack'

const props = defineProps({
    stack: {
        type: Object,
        default: null,
    },
    fixedWidth: {
        type: Boolean,
        default: false,
    },
})

const color = computed(() => {
    statusColor(props.stack?.status)
})

// const statusName = computed (() => {
//     console.log(props.stack.stackStatus)
//     return props.stack.stackStatus
//     // return statusNameShort(props.stack?.status);
// })

const className = (() => {
    let className = `badge rounded-pill bg-${color}`;
    if (props.fixedWidth) {
        className += " fixed-width";
    }
    return className
})

</script>

<style scoped>
.badge {
    min-width: 62px;

}

.fixed-width {
    width: 62px;
    overflow: hidden;
    text-overflow: ellipsis;
}
</style>
