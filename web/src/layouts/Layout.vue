<script setup lang="ts">
import { loginAPI } from '@/api/user'
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

</script>

<template>
  <el-container>
    <el-header class='bg-gray flex items-center justify-between'>
      <!-- 左侧 -->
      <div class="flex flex-row items-center justify-center">
        <img src="@/assets/meow.png" class="w-50px h-50px">
        <p class="text-white text-xl ml-5px">Meow</p>
      </div>
      <!-- 右侧 -->
      <div class="flex ">
        <el-dropdown trigger="click" oncommand="null">
          <span class="text-white! text-xl">admin</span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>
                <router-link to="/settings/about">设置11111</router-link>
              </el-dropdown-item>
              <el-dropdown-item>
                退出登录
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
                <el-button type="primary" @click="onLogin">
                  登录
                </el-button>
              </el-form-item>
            </el-row>
          </el-form>
        </div>
      </el-container>
    </el-main>
  </el-container>
</template>


<style scoped lang="scss"></style>