<template>
    <div class="my-4">
        <div class="text-xl fw-bold mb-1">
            主机名
        </div>
        <div class="flex flex-row mb-3">
            <el-input v-model="settings.hostname" class="max-w-250px" placeholder="未设置:沿用当前主机名" clearable />
            <el-button type="primary" @click="autoGetPrimaryHostname" class="ml-5px">
                自动获取
            </el-button>
        </div>
    </div>
</template>

<script setup lang="ts">

const settings = ref({
    hostname: ""
})

const autoGetPrimaryHostname = () => {
    settings.value.hostname = location.hostname
}

</script>
