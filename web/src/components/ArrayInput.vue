<template>
    <div>
        <div v-if="valid">
            <!-- <ul v-if="isArrayInitd" class="p-10px pl-0"> -->
            <ul class="p-10px pl-0">
                <li v-for="(_, index) in array" :key="index" class="flex items-center">
                    <el-input v-model="array[index]" type="text"
                        class="bg-transparent flex-grow color-black bg-[#070a10] placeholder-[#1d2634]"
                        :placeholder="placeholder">
                        <template #suffix>
                            <svg class="i-lucide-x h-16px w-16px" @click="remove(index)"></svg>
                        </template>
                        <!-- /> -->
                    </el-input> 
                        
                </li>
            </ul>

            <el-button class="mt-3" @click="addField" type="primary">
                {{ $t("addListItem", [ displayName ]) }}
            </el-button>
        </div>
        <div v-else>
            {{ $t("LongSyntaxNotSupported") }}
        </div>
    </div>
</template>

<script setup lang="ts">
const props = defineProps({
    name: {
        type: String,
        required: true,
    },
    placeholder: {
        type: String,
        default: "",
    },
    displayName: {
        type: String,
        required: true,
    },
    objectType: {
        type: String,
        default: "service",
    },
})

const array = computed(() => {
    if (!service.value[props.name]) {
        return [];
    }
    return service.value[props.name];
})



// const isArrayInitd = computed(() => {
//     return service.value[props.name] !== undefined;
// })

const valid = computed(() => {

    // Check if the array is actually an array
    if (!Array.isArray(array.value)) {
        return false;
    }

    // Check if the array contains non-object only.
    for (let item of array.value) {
        if (typeof item === "object") {
            return false;
        }
    }
    return true;
})

const service = computed(() => {
    return {
        "ports": ["8080/80"],
        "networks": [],
        "depends_on": [],
        "environment": [],
        "volumes": [],
    }
    if (props.objectType === "service") {
        // Used in Container.vue
        // return this.$parent.$parent.service;
    } else if (props.objectType === "x-dockge") {

        // if (!this.$parent.$parent.jsonConfig["x-dockge"]) {
        //     return {};
        // }

        // Used in Compose.vue
        // return this.$parent.$parent.jsonConfig["x-dockge"];
    } else {
        return {};
    }
})

const addField = () => {
    if (props.objectType === "x-dockge") {
        // if (!this.$parent.$parent.jsonConfig["x-dockge"]) {
        //     this.$parent.$parent.jsonConfig["x-dockge"] = {};
        // }
    }

    // Create the array if not exists.
    if (!service.value[props.name]) {
        service.value[props.name] = []
    }

    array.value.push("")
    console.log(array.value, 1231414312)
}
const remove = (index: number) => {
    array.value.splice(index, 1);
}

</script>
