<template>
    <div class="mb-3" :style="boxStyle">
        <div class="b-b-1 b-b-solid b-b-[#dee2e6] rounded-t-10px m-[-10px] mb-10px p-10px dark:bg-[#161b22] dark:b-b-0">
            <div class="flex justify-between items-center">
                <div class="placeholder"></div>
                <div class="flex items-center justify-center">
                    <svg class="i-lucide-search p-1px text-black"></svg>
                    <div></div>
                    <el-input v-model="searchText" clearable class="w-150px! ml-2px" />
                </div>
            </div>
        </div>
        <div :class="{ scrollbar: false }" :style="stackListStyle">
            <div v-if="stackList.length === 0" class="text-center mt-3">
                <router-link to="/compose"> 组合你的第一个堆栈！</router-link>
            </div>
            <div v-else>
                <StackListItem v-for="(item, index) in stackList" :key="index" :stack="item" />
            </div>
        </div>
    </div>

    <!-- <el-popconfirm ref="confirmPause" trigger="click" @confirm="pauseSelected">
    </el-popconfirm> -->
</template>

<script setup lang="ts">
// import Confirm from "../components/Confirm.vue";
import { listStackAPI } from "@/api/stack";
import StackListItem from "../components/StackListItem.vue";
// import { CREATED_FILE, CREATED_STACK, EXITED, RUNNING, UNKNOWN } from "@/types/stack";
import type { ListItem } from "@/types/stack";
// import { statusConvert } from "@/utils/container";

// const confirmPause = ref<HTMLBodyElement>()
const stackList = ref<ListItem[]>([])
const searchText = ref("")

// let selectedStacks = {}
let windowTop = 0

const getStackList = async () => {
    const res = await listStackAPI()
    stackList.value = res.items
    // 将 list 中状态统一转换。 由 running(1) => 1
    // stackList.value.forEach(stack => {
    //     // stack.stackStatus = statusConvert(stack.stackStatus)
    // })
}

onMounted( async () => {
    await getStackList()
})


const boxStyle = computed(() => {
    // let heigth = 0
    if (window.innerWidth > 550) {
        return `calc(100vh - 160px + ${windowTop}px)`
    } else {
        return "calc(100vh - 160px)"
    }
})

const sortedStackList = computed(() => {
    // 
    // let result = Object.values([])
    let result = [
        {
            id: 1,
            isManagedByDockge: true,
            active: "active",
            stackName: "ahahah",
            stackStatus: "running(1)",
            tags: [{
                name: "",
                value: "",
            }],
        },
        {
            id: 2,
            isManagedByDockge: true,
            active: "inactive",
            stackName: "aha12313hah",
            stackStatus: "exited(1)",
            tags: [{
                name: "",
                value: "",
            }],
        }
    ]
    result = result.filter(stack => {
        let searchTextMatch = true;
        if (searchText.value !== "") {
            const loweredSearchText = searchText.value.toLowerCase();
            searchTextMatch = stack.stackName.toLowerCase().includes(loweredSearchText)
            // || stack.tags.find(tag => tag.name.toLowerCase().includes(loweredSearchText)
            //     || tag.value?.toLowerCase().includes(loweredSearchText))
        }

        return searchTextMatch //  && activeMatch && tagsMatch;
    })



    result.sort((m1, m2) => {
        // sort by managed by dockge
        if (m1.isManagedByDockge && !m2.isManagedByDockge) {
            return -1;
        } else if (!m1.isManagedByDockge && m2.isManagedByDockge) {
            return 1;
        }

        return m1.stackName.localeCompare(m2.stackName);
    });

    console.log(result, 1222223)

    return result;
})

const stackListStyle = computed(() => {
    let listHeaderHeight = 60;
    // if (selectMode) {
    //     listHeaderHeight += 42;
    // }

    return {
        "height": `calc(100% - ${listHeaderHeight}px)`
    };
})

// const selectedStackCount = computed(() => {
//     return Object.keys(selectedStacks).length;
// })


onMounted(() => {
    window.addEventListener("scroll", onScroll);
})

onBeforeUnmount(() => {
    window.addEventListener("scroll", onScroll);
})

const onScroll = () => {
    if (window.top!.scrollY <= 133) {
        windowTop = window.top!.scrollY;
    } else {
        windowTop = 133;
    }
}

// const clearSearchText = () => {
//     searchText = ""
// }

// const pauseDialog = () => {
//     confirmPause.value!.click();
// }

// const resumeSelected = () => {
//     Object.keys(selectedStacks)
//         .filter(id => !stackList[id].active)
//     // .forEach(id => this.$root.getSocket().emit("resumeStack", id, () => { }));

//     cancelSelectMode()
// }

// const cancelSelectMode = () => {
//     selectMode = false;
//     selectedStacks = {};
// }

// const pauseSelected = () => {
//     Object.keys(selectedStacks)
//         .filter(id => stackList[id].active)
//     // .forEach(id => this.$root.getSocket().emit("pauseStack", id, () => { }));

//     cancelSelectMode()
// }

// const isSelected = (id: string) => {
//     return id in selectedStacks;
// }

// const select = (id: string) => {
//     selectedStacks[id] = true;
// }

// const deselect = (id: string) => {
//     delete selectedStacks[id];
// }

// const filtersActive = (() => {
//     return filterState.status != null || filterState.active != null || filterState.tags != null || searchText !== "";
// }
// )
// const updateFilter = (newFilter) => {
//     filterState = newFilter;
// }

// watch: {
//     searchText() {
//         for (let stack of this.sortedStackList) {
//             if (!this.selectedStacks[stack.id]) {
//                 if (this.selectAll) {
//                     this.disableSelectAllWatcher = true;
//                     this.selectAll = false;
//                 }
//                 break;
//             }
//         }
//     },
//     selectAll() {
//         if (!this.disableSelectAllWatcher) {
//             this.selectedStacks = {};

//             if (this.selectAll) {
//                 this.sortedStackList.forEach((item) => {
//                     this.selectedStacks[item.id] = true;
//                 });
//             }
//         } else {
//             this.disableSelectAllWatcher = false;
//         }
//     },
//     selectMode() {
//         if (!this.selectMode) {
//             this.selectAll = false;
//             this.selectedStacks = {};
//         }
//     },
// },
</script>