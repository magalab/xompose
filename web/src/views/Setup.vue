<template>
    <div class="flex items-center py-10" data-cy="setup-form">
        <!-- max-width: 330px;
    padding: 15px;
    margin: auto;
    text-align: center; -->
        <div class="w-full max-w-330px p-15px m-auto text-center">
            <!-- <form @submit.prevent="submit"> -->
            <!-- <div>
                    <object width="64" height="64" data="/icon.svg" />
                    <div style="font-size: 28px; font-weight: bold; margin-top: 5px;">
                        Meow
                    </div>
                </div>

                <p class="mt-3">
                    创建管理员账号
                </p> -->

            <!-- <div class="form-floating">
                    <select id="language" v-model="$root.language" class="form-select">
                        <option v-for="(lang, i) in $i18n.availableLocales" :key="`Lang${i}`" :value="lang">
                            {{ $i18n.messages[lang].languageName }}
                        </option>
                    </select>
                    <label for="language" class="form-label">{{ $t("Language") }}</label>
                </div> -->

            <!-- <div class="form-floating mt-3">
                    <input id="floatingInput" v-model="username" type="text" class="form-control" :placeholder="$t('Username')" required data-cy="username-input">
                    <label for="floatingInput">用户名</label>
                </div>

                <div class="form-floating mt-3">
                    <input id="floatingPassword" v-model="password" type="password" class="form-control" :placeholder="$t('Password')" required data-cy="password-input">
                    <label for="floatingPassword">密码</label>
                </div>

                <div class="form-floating mt-3">
                    <input id="repeat" v-model="repeatPassword" type="password" class="form-control" :placeholder="$t('Repeat Password')" required data-cy="password-repeat-input">
                    <label for="repeat">{{ $t("Repeat Password") }}</label>
                </div> -->
            <el-form :model="form">
                <div>
                    <!-- <object width="64" height="64" data="/logo.svg" /> -->
                     <img src="@/assets/meow.png" class="w-16 h-16">
                    <!-- <div style="font-size: 28px; font-weight: bold; margin-top: 5px;">
                        Meow
                    </div> -->
                </div>

                <p class="mt-3">
                    <!-- {{ $t("Create your admin account") }} -->
                    创建管理员账号
                </p>
                <el-form-item>
                    <el-input v-model="form.username" size="large" clearable>
                        <template #prefix>
                            <svg class="i-lucide-user text-violet"></svg>
                        </template>
                    </el-input>
                </el-form-item>
                <el-form-item>
                    <el-input v-model="form.password" type="password" show-password size="large" clearable>
                        <template #prefix>
                            <svg class="i-lucide-lock text-violet"></svg>
                        </template>
                    </el-input>
                </el-form-item>
                <el-form-item>
                    <el-input v-model="form.password2" type="password" show-password size="large" clearable>
                        <template #prefix>
                            <svg class="i-lucide-lock text-violet"></svg>
                        </template>
                    </el-input>
                </el-form-item>
                <el-row justify="end">
                    <el-form-item>
                        <el-button type="primary" @click="onCreate">
                            创建
                        </el-button>
                    </el-form-item>
                </el-row>
            </el-form>
        </div>
    </div>
</template>

<script setup lang="ts">
import { setupAPI } from '@/api/system'

const processing = ref(false)

interface Form {
    username: string
    password: string
    password2: string
}


const form = ref<Form>({} as Form)
const router = useRouter()

onMounted(() => {
    // TODO 是否能够创建新账号
    let need = false
    if (need) {
        router.push("/");
    }
})

const onCreate = async  () => {
    processing.value = true
    if (form.value.password !== form.value.password2) {
        ElMessage.error("密码不一致")
        processing.value = false
        return
    }
     await setupAPI({
        username: form.value.username,
        password: form.value.password,
        password2: form.value.password2,
    })
    ElMessage.success("创建成功")
    setTimeout(() => {
        router.push('/')
    }, 500)
}
</script>