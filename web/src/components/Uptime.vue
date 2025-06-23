<template>
    <span :class="className">{{ statusName }}</span>
</template>

<script setup lang="ts">
import { statusColor, statusNameShort } from '@/types/stack'

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

// const upTime = computed((): string => {
//     return "notAvailableShort"
// })

const color = computed(() => {
    statusColor(props.stack?.status)
})

const statusName = (() => {
    // return this.$t(statusNameShort(this.stack?.status));
    return statusNameShort(props.stack?.status);
})

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
