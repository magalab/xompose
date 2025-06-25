<template>
    <div>
        <h1 class="mb-3"> {{ t("Settings") }}</h1>
        <div class="p-20px min-h-[calc(100vh-155px)]">
            <div class="flex flex-row">
                <div class="lg:w-3/12 md:w-5/12">
                    <router-link v-for="(item, key) in subMenus" :key="key" :to="`/settings/${key}`"
                        class="no-underline">
                        <div
                            class="rounded-[10px] m-2 p-[0.7em_1em] cursor-pointer border-l-0 
                            transition-all duration-100 ease-in-out 
                            hover:bg-[#e7faec] 
                            active:rounded-l-0 active:border-b-1 active:border-solid active:border-[#74c2ff] active:bg-[#e7faec]">
                            {{ item.title }}
                        </div>
                    </router-link>
                </div>
                <div class="lg:w-9/12 md:w-7/12">
                    <div v-if="currentPage" class="w-[calc(100%+20px)] 
                        border-b-1 border-b-solid border-b-[#dee2e6] rounded-r-[10px]
                        mt-[-20px] mr-[-20px] py-12.5px px-16px text-26px">
                        {{ subMenus[currentPage].title }}
                    </div>
                    <div class="mx-3">
                        <router-view v-slot="{ Component }">
                            <transition name="slide-fade" appear>
                                <component :is="Component" />
                            </transition>
                        </router-view>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
const { t } = useI18n()

const currentPage = computed(() => {
    let pathSplit = useRoute().path.split("/");
    let pathEnd = pathSplit[pathSplit.length - 1];
    if (!pathEnd || pathEnd === "settings") {
        return null;
    }
    return pathEnd;
})

const subMenus = ref({
    general: {
        title: t("General"),
    },
    // appearance: {
    //     title: t("Appearance"),
    // },
    security: {
        title: t("Security"),
    },
    about: {
        title: t("About"),
    },
})
</script>
