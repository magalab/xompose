<template>
    <div class="min-w-400px w-50% mx-auto flex flex-col items-center">
        <svg class="h-20 w-20 color-purple mt-50 mb-5 i-lucide-home"></svg>
        <div class="flex flex-col">
            <div class="flex flex-wrap items-center border-b-1 border-b-solid border-b-[#d4d7de]">
                <svg class="text-4 h-35px w-35px leading-8 mr-2 i-lucide-user"></svg>
                <input id="username"
                    class="bg-white border-none inline-flex h-[35px] w-[200px] appearance-none items-center justify-center rounded-lg px-[10px] text-sm leading-none shadow-sm outline-none focus:shadow-[0_0_0_2px_black] selection:color-white selection:bg-blackA9"
                    :placeholder="$t('message.username')" type="text" v-model="username" v-bind="usernameAttrs">
            </div>
            <div class="h-5 text-3 color-[#f56c6c] ">{{ errors.username }}</div>

        </div>
        <div class="flex flex-col">
            <div class="flex flex-wrap items-center border-b-1 border-b-solid border-b-[#d4d7de]">
                <svg class="text-4 h-35px w-35px leading-8 mr-2 i-lucide-lock"></svg>
                <input id="password"
                    class="bg-white border-none inline-flex h-[35px] w-[200px] appearance-none items-center justify-center rounded-lg px-[10px] text-sm leading-none shadow-sm outline-none focus:shadow-[0_0_0_2px_black] selection:color-white selection:bg-blackA9"
                    :placeholder="$t('message.password')" type="password" v-model="password" v-bind="passwordAttrs">
            </div>
            <div class="h-5 text-3 color-[#f56c6c] ">{{ errors.password }}</div>

        </div>
        <button @click="login"
            class="w-25 mt-1px px-2 py-2 border-1 border-white cursor-pointer font-500 text-14px leading-14px bg-purple color-white">
            {{ $t("message.signin") }}
        </button>
    </div>
</template>

<script setup lang="ts">
import { toTypedSchema } from '@vee-validate/zod'
import { useForm } from 'vee-validate'
import { z } from 'zod'
import md5 from 'md5'
import { loginAPI } from '@/api/user'
import { onMounted } from 'vue'
import { listStackAPI } from '@/api/stack'

onMounted(async () => {
    const res =  await listStackAPI()
    console.log(res)
})


const { errors, defineField } = useForm({
    validationSchema: toTypedSchema(
        z.object({
            username: z.string().trim().min(1, '请输入用户名'),
            password: z.string().trim().min(1, '请输入密码'),
        })
    )
})
const [username, usernameAttrs] = defineField('username')
const [password, passwordAttrs] = defineField('password')


const login = async () => {
    //   try {
    //     await userStore.login(username.value, md5(password.value))
    //     if (redirectUrl) {
    //       router.replace(redirectUrl)
    //     } else {
    //       router.replace('/index')
    //     }

    //   } catch {

    //   }
    try {
        await loginAPI({ username: username.value as string, password: md5(password!.value as string) })
    } catch (err) {
        console.log(err);

    }
}


</script>