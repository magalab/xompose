<template>
    <div class="w-98%">
        <div class="flex flex-row gap-4">
            <!-- <div v-if="!$root.isMobile" class="col-12 col-md-4 col-xl-3"> -->
            <div class="w-full md:w-1/3 xl:w-1/4 ">
                <div class="mb-20px">
                    <router-link to="/compose"  >
                        <el-button type="primary" class="rounded-50%">
                            <i-ep-plus class="mr-2px text-sm" />
                            <span>compose</span>
                        </el-button>
                    </router-link>
                </div>
                <StackList :scrollbar="true"  />
            </div>

            <div ref="container" class="w-full md:w-2/3 xl:w-3/4 mb-3">
                <!-- Add :key to disable vue router re-use the same component -->
                <router-view :key="$route.fullPath" :calculatedHeight="height" />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import StackList from "../components/StackList.vue";
const height = ref(0)
const container = ref < HTMLElement > ()

onMounted(() => {
    container.value!.offsetHeight
})

</script>
