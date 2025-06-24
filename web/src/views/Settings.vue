<template>
    <div>
        <h1 v-show="show" class="mb-3">
            设置
        </h1>

        <div class="shadow-box p-20px min-h-[calc(100vh-155px)]">
            <div class="flex flex-row">
                <div v-if="showSubMenu" class="lg:w-3/12 md:w-5/12">
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

const router = useRouter()

const show = ref(true)
const settings = ref({})

const settingsLoaded = ref(false)

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
        title: "General",
    },
    // appearance: {
    //     title: "Appearance",
    // },
    // security: {
    //     title: this.$t("Security"),
    // },
    about: {
        title: "About",
    },
})

const showSubMenu = computed(() => {
    // if (this.$root.isMobile) {
    //     return !this.currentPage;
    // } else {
    //     return true;
    // }
    return true
})

// watch: {
//     "$root.isMobile"() {
//         this.loadGeneralPage();
//     }
// },
onMounted(() => {
    loadSettings()
    loadGeneralPage()
})

const loadGeneralPage = () => {
    // mobile todo
    if (!currentPage) {
        router.push("/settings/appearance")
    }
}

const loadSettings = () => {
    // getSettings api
    // this.settings = res.data;
    // if (this.settings.checkUpdate === undefined) {
    //     this.settings.checkUpdate = true;
    // }
    // this.settingsLoaded = true;
}

const saveSettings = () => {
    let { success, msg } = validSettings()
    if (success) {
        // save api
        // this.$root.getSocket().emit("setSettings", this.settings, currentPassword, (res) => {
        //     this.$root.toastRes(res);
        //     this.loadSettings();

        //     if (callback) {
        //         callback();
        //     }
        // });
    } else {
        // Toast 警告
        alert(msg)
    }
}

const validSettings = () => {
    if (settings.keepDataPeriodDays < 0) {
        return {
            success: false,
            msg: "dataRetentionTimeError",
        };
    }
    return {
        success: true,
        msg: "",
    }
}
</script>
