<template>
    <div class="w-full h-full">
        <div class="h-50% mx-auto mt-250px w-18% flex flex-col justify-center items-center bg-violet rounded-2xl">
            <img src="@/assets/meow.png" class="w-50px h-50px rounded-50% mb-1">
            <p class="mb-5">meow</p>
            <el-form :model="form">
                <el-form-item>
                    <el-input v-model="form.username" size="large" clearable>
                        <template #prefix>
                            <i-ep-user class="text-violet"></i-ep-user>
                        </template>
                    </el-input>
                </el-form-item>
                <el-form-item>
                    <el-input v-model="form.password" type="password" show-password size="large" clearable>
                        <template #prefix>
                            <i-ep-lock class="text-violet"></i-ep-lock>
                        </template>
                    </el-input>
                </el-form-item>
                <el-row justify="end">
                    <el-form-item>
                        <el-button type="primary">
                            登录
                        </el-button>
                    </el-form-item>
                </el-row>
            </el-form>
        </div>
    </div>
</template>

<script setup lang="ts">
interface Form {
    username: string
    password: string
}


const form = ref<Form>({} as Form)
</script>

<style scoped></style>