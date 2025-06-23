<script setup>
// import { useUserStore } from '@/store/user';
// import { useRouter } from 'vue-router';

// const router = useRouter()
// const userStore = useUserStore()

// // 退出登录
// const confirm = () => {
//   userStore.clearUserInfo()
//   router.push('/login')
// }

</script>

<template>
  <el-container>
    <el-header class='bg-gray flex items-center justify-between'>
      <!-- 左侧 -->
      <div class="flex flex-row items-center justify-center">
        <img src="@/assets/meow.png" class="w-50px h-50px">
        <p  class="text-white text-xl ml-5px">Meow</p>
      </div>
      <!-- 右侧 -->
      <div class="flex ">
        <el-dropdown trigger="click" oncommand="null">
          <span class="text-white! text-xl">admin</span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>
                个人信息
              </el-dropdown-item>
              <el-dropdown-item>
                退出登录
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </el-header>
  </el-container>
</template>


<style scoped lang="scss"></style>