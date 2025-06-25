<template>
    <transition ref="tableContainer" name="slide-fade" appear>
        <!-- hide the home page -->
        <div class="w-full" v-if="$route.name==='DashboardHome'">
            <h1 class="mb-3">{{ $t("Home") }}</h1>
            <div class="w-full">
                <div class="flex flex-row">
                    <div class="col">
                        <h3>{{ $t("Active") }}</h3>
                        <span class="num active">{{ activeNum }}</span>
                    </div>
                    <div class="col">
                        <h3>{{ $t("Exited") }}</h3>
                        <span class="num exited">{{ exitedNum }}</span>
                    </div>
                    <div class="col">
                        <h3>{{ $t("Inactive") }}</h3>
                        <span class="num inactive">{{ inactiveNum }}</span>
                    </div>
                </div>

                <!-- Docker Run -->
                <h2 class="mb-3"> {{ $t("Docker Run") }}</h2>
                <div class="mb-3">
                    <el-input v-model="dockerRunCommand" type="textarea" :autosize="{ minRows: 3 }"
                        placeholder="docker run ..." resize="none" class="border-none text-xl! max-w-60% min-w-250px">
                    </el-input>
                </div>
                <el-button type="primary" class="mb-4" @click="convertDockerRun">
                    {{ $t("Convert to Compose") }}
                </el-button>
            </div>
        </div>
    </transition>
    <router-view ref="child" />
</template>

<script setup lang="ts">

// const props = defineProps({
//     calculatedHeight: {
//         type: Number,
//         default: 0
//     }
// })

// const page = ref(1)
const perPage = ref(25)
const initialPerPage = ref(25)
// const paginationConfig = ref({
//     hideCount: true,
//     chunksNavigation: "scroll",
// })

// const importantHeartBeatListLength = ref(0)
// const displayedRecords = ref([])
const dockerRunCommand = ref("")

const activeNum = computed(() => {
    return getStatusNum("active");
})

const inactiveNum = computed(() => {
    return getStatusNum("inactive");
})

const exitedNum = computed(() => {
    return getStatusNum("exited");
})

// watch: {
//         perPage() {
//             this.$nextTick(() => {
//                 this.getImportantHeartbeatListPaged();
//             });
//         },

//         page() {
//             this.getImportantHeartbeatListPaged();
//         },
//     },

onMounted(() => {
    initialPerPage.value = perPage.value
    window.addEventListener("resize", updatePerPage);
    updatePerPage()
})

onBeforeUnmount(() => {
    window.removeEventListener("resize", updatePerPage);
})

const tableContainer = ref()

const updatePerPage = () => {
    // const tableContainer = this.$refs.tableContainer;
    const tableContainerHeight = tableContainer.value.offsetHeight;
    const availableHeight = window.innerHeight - tableContainerHeight;
    const additionalPerPage = Math.floor(availableHeight / 58);

    if (additionalPerPage > 0) {
        perPage.value = Math.max(initialPerPage.value, perPage.value + additionalPerPage);
    } else {
        perPage.value = initialPerPage.value;
    }
}

const getStatusNum = (statusName: string): number => {
    let num = 0;

    // for (let stackName in this.$root.completeStackList) {
    //     const stack = this.$root.completeStackList[stackName];
    //     if (statusNameShort(stack.status) === statusName) {
    //         num += 1;
    //     }
    // }
    return num;
}

const convertDockerRun = () => {
    if (dockerRunCommand.value.trim() === "docker run") {
        ElMessage.error("Please enter a docker run command");
    }

    // composerize is working in dev, but after "vite build", it is not working
    // So pass to backend to do the conversion
    // this.$root.getSocket().emit("composerize", this.dockerRunCommand, (res) => {
    //     if (res.ok) {
    //         this.$root.composeTemplate = res.composeTemplate;
    //         this.$router.push("/compose");
    //     } else {
    //         this.$root.toastRes(res);
    //     }
    // });
}

/**
* Updates the displayed records when a new important heartbeat arrives.
* @param {object} heartbeat - The heartbeat object received.
* @returns {void}
*/
// onNewImportantHeartbeat(heartbeat) {
//     if (this.page === 1) {
//         this.displayedRecords.unshift(heartbeat);
//         if (this.displayedRecords.length > this.perPage) {
//             this.displayedRecords.pop();
//         }
//         this.importantHeartBeatListLength += 1;
//     }
// }

// /**
//  * Retrieves the length of the important heartbeat list for all monitors.
//  * @returns {void}
//  */
// getImportantHeartbeatListLength() {
//     this.$root.getSocket().emit("monitorImportantHeartbeatListCount", null, (res) => {
//         if (res.ok) {
//             this.importantHeartBeatListLength = res.count;
//             this.getImportantHeartbeatListPaged();
//         }
//     });
// }

// /**
//  * Retrieves the important heartbeat list for the current page.
//  * @returns {void}
//  */
// getImportantHeartbeatListPaged() {
//     const offset = (this.page - 1) * this.perPage;
//     this.$root.getSocket().emit("monitorImportantHeartbeatListPaged", null, offset, this.perPage, (res) => {
//         if (res.ok) {
//             this.displayedRecords = res.data;
//         }
//     });
// }





</script>

<style lang="scss" scoped>
@import "@/styles/vars";

.num {
    font-size: 30px;

    font-weight: bold;
    display: block;

    &.active {
        color: $primary;
    }

    &.exited {
        color: $danger;
    }
}
</style>
