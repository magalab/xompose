<script setup lang="ts">
import { loginAPI } from '@/api/user'
// import Editor from '@/components/YamlEditor.vue'
interface Form {
  username: string
  password: string
}

const form = ref<Form>({} as Form)
const isLogin = ref(true)
const onLogin = async () => {
  await loginAPI({
    username: form.value.username,
    password: form.value.password,
  })
  isLogin.value = true
}


// import Prism from 'prismjs'

// const defaultTemplate = `
// services:
//   nginx:
//     image: nginx:latest
//     restart: unless-stopped
//     ports:
//       - "8080:80"
// `

// const html =  Prism.highlight(
//   defaultTemplate,
//   Prism.languages.yaml,
//   "yaml",
// )



</script>

<template>
  <el-container>
    <el-header class='bg-gray flex items-center justify-between'>
      <!-- 左侧 -->
      <router-link to='/' class="flex flex-row items-center justify-center no-underline" >
        <img src="@/assets/meow.png" class="w-50px h-50px rounded-50%">
        <p class="text-white text-xl ml-5px"> {{ $t("Meow") }}</p>
      </router-link>
      <!-- 右侧 -->
      <div class="flex">
        <el-dropdown trigger="click" oncommand="null">
          <span class="text-white! text-xl">admin</span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>
                <router-link to="/settings/about">{{ $t("Settings") }}</router-link>
              </el-dropdown-item>
              <el-dropdown-item>
                {{ $t("Logout") }}
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </el-header>
    <!-- 登录 -->
    <el-main>
      <router-view v-if="isLogin"></router-view>
      <el-container v-else>
        <div class="h-50% mx-auto mt-250px w-18% flex flex-col justify-center items-center bg-violet rounded-2xl">
          <img src="@/assets/meow.png" class="w-50px h-50px rounded-50% mb-1">
          <p class="mb-5"> {{ $t("Meow") }}</p>
          <el-form :model="form">
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
            <el-row justify="end">
              <el-form-item>
                <el-button type="primary" @click="onLogin">
                  {{ $t("Login") }}
                </el-button>
              </el-form-item>
            </el-row>
          </el-form>
        </div>
      </el-container>
    </el-main>
  </el-container>
</template>
