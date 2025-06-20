<template>
    <div>aaaaaa {{ stackList.length }} 121</div>
    <div v-if="stackList.length > 0">
        <div v-for="item in stackList">
            {{ item.stackName }}: {{ item.stackStatus }} : {{ item.isManaged }}
        </div>

    </div>
</template>

<script setup lang="ts">
import { listStackAPI } from '@/api/stack';
import { type StackListItem } from '@/types/stack'
import { onMounted, ref } from 'vue';

const stackList = ref<Array<StackListItem>>([])

const getStackList = async () => {
    const res = await listStackAPI()
    stackList.value = res.data.data.items
}

onMounted(async () => {
    await getStackList()
})
</script>

<style lang="scss" scoped></style>