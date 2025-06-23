<template>
    <div class="shadow-box mb-3 bg-gray" :style="boxStyle">
        <div class="list-header">
            <div class="header-top">
                <div class="placeholder"></div>
                <div class="flex items-center justify-center">
                    <i-ep-search class="p-1px text-white" />
                    <div></div>
                    <el-input v-model="searchText" clearable class="w-150px! ml-2px" />
                </div>
            </div>
        </div>
        <div ref="stackList" class="stack-list" :class="{ scrollbar: false }" :style="stackListStyle">
            <div v-if="sortedStackList.length === 0" class="text-center mt-3">
                <router-link to="/compose"> 组合你的第一个堆栈！</router-link>
            </div>

            <StackListItem v-for="(item, index) in sortedStackList" :key="index" :stack="item"
                :isSelectMode="selectMode" :isSelected="isSelected" :select="select" :deselect="deselect" />
        </div>
    </div>

    <!-- <Confirm ref="confirmPause" :yes-text="$t('Yes')" :no-text="$t('No')" @yes="pauseSelected">
        {{ $t("pauseStackMsg") }}
    </Confirm> -->
    <el-popconfirm ref="confirmPause" trigger="click" @confirm="pauseSelected">

    </el-popconfirm>
</template>

<script setup lang="ts">
// import Confirm from "../components/Confirm.vue";
import StackListItem from "../components/StackListItem.vue";
import { CREATED_FILE, CREATED_STACK, EXITED, RUNNING, UNKNOWN } from "@/types/stack";
import type { StackList } from "@/types/stack";

const confirmPause = ref<HTMLBodyElement>()

const stackList = ref<StackList>({} as StackList)

const searchText = ref("")
let selectMode = false

// let selectAll = false
// let disableSelectAllWatcher = false
let selectedStacks = {}
let windowTop = 0
// let filterState: {
//     status: null,
//     active: null,
//     tags: null,
// }
// props: {
//     /** Should the scrollbar be shown */
//     scrollbar: {
//         type: Boolean,
//     },
// },

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
            name: "ahahah",
            status: 3,
            tags: [{
                name: "",
                value: "",
            }],
        },
        {
            id: 2,
            isManagedByDockge: true,
            active: "active",
            name: "aha12313hah",
            status: 4,
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
            searchTextMatch = stack.name.toLowerCase().includes(loweredSearchText)
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

        // sort by status
        if (m1.status !== m2.status) {
            if (m2.status === RUNNING) {
                return 1;
            } else if (m1.status === RUNNING) {
                return -1;
            } else if (m2.status === EXITED) {
                return 1;
            } else if (m1.status === EXITED) {
                return -1;
            } else if (m2.status === CREATED_STACK) {
                return 1;
            } else if (m1.status === CREATED_STACK) {
                return -1;
            } else if (m2.status === CREATED_FILE) {
                return 1;
            } else if (m1.status === CREATED_FILE) {
                return -1;
            } else if (m2.status === UNKNOWN) {
                return 1;
            } else if (m1.status === UNKNOWN) {
                return -1;
            }
        }
        return m1.name.localeCompare(m2.name);
    });

    return result;
})

const stackListStyle = computed(() => {
    let listHeaderHeight = 60;
    if (selectMode) {
        listHeaderHeight += 42;
    }

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

const pauseDialog = () => {
    confirmPause.value!.click();
}

const resumeSelected = () => {
    Object.keys(selectedStacks)
        .filter(id => !stackList[id].active)
    // .forEach(id => this.$root.getSocket().emit("resumeStack", id, () => { }));

    cancelSelectMode()
}

const cancelSelectMode = () => {
    selectMode = false;
    selectedStacks = {};
}

const pauseSelected = () => {
    Object.keys(selectedStacks)
        .filter(id => stackList[id].active)
    // .forEach(id => this.$root.getSocket().emit("pauseStack", id, () => { }));

    cancelSelectMode()
}

const isSelected = (id: string) => {
    return id in selectedStacks;
}

const select = (id: string) => {
    selectedStacks[id] = true;
}

const deselect = (id: string) => {
    delete selectedStacks[id];
}

// const filtersActive = (() => {
//     return filterState.status != null || filterState.active != null || filterState.tags != null || searchText !== "";
// }
// )
const updateFilter = (newFilter) => {
    filterState = newFilter;
}

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

<style lang="scss" scoped>
@import "../styles/vars.scss";

.shadow-box {
    height: calc(100vh - 600px);
    position: sticky;
    top: 100px;
}

.small-padding {
    padding-left: 5px !important;
    padding-right: 5px !important;
}

.list-header {
    border-bottom: 1px solid #dee2e6;
    border-radius: 10px 10px 0 0;
    margin: -10px;
    margin-bottom: 10px;
    padding: 10px;

    .dark & {
        background-color: $dark-header-bg;
        border-bottom: 0;
    }
}

.header-top {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.header-filter {
    display: flex;
    align-items: center;
}

@media (max-width: 770px) {
    .list-header {
        margin: -20px;
        margin-bottom: 10px;
        padding: 5px;
    }
}

.search-wrapper {
    display: flex;
    align-items: center;
}

.search-icon {
    padding: 10px;
    color: #c0c0c0;

    // Clear filter button (X)
    svg[data-icon="times"] {
        cursor: pointer;
        transition: all ease-in-out 0.1s;

        &:hover {
            opacity: 0.5;
        }
    }
}

.search-input {
    max-width: 15em;
}

.stack-item {
    width: 100%;
}

.tags {
    margin-top: 4px;
    padding-left: 67px;
    display: flex;
    flex-wrap: wrap;
    gap: 0;
}

.bottom-style {
    padding-left: 67px;
    margin-top: 5px;
}

.selection-controls {
    margin-top: 5px;
    display: flex;
    align-items: center;
    gap: 10px;
}
</style>
