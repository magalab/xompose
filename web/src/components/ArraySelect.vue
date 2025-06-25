<template>
    <div>
        <div v-if="valid">
            <ul v-if="isArrayInitd" class="bg-[#070a10] p-10px pl-0">
                <li v-for="(_, index) in array" :key="index">
                    <select v-model="array[index]"
                        class="bg-transparent flex-grow border-none outline-none color-black bg-[#070a10] placeholder-[#1d2634]">
                        <option value="">{{ $t(`Select a network...`) }}</option>
                        <option v-for="option in options" :key="option" :value="option">{{ option }}</option>
                    </select>
                    <svg class="i-lucide-x h-16px w-16px" @click="remove(index)" ></svg>
                </li>
            </ul>

            <el-button class="mt-3" type="primary" @click="addField">
               {{ $t("addListItem", [ displayName ]) }}
            </el-button>
        </div>
        <div v-else>
            Long syntax is not supported here. Please use the YAML editor.
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
    options: {
        type: Array,
        required: true,
    },
})

const array = computed(() => {
    if (!service.value[props.name]) {
        return []
    }
    return service.value[props.name]
})

const isArrayInitd = computed(() => {
    return service.value[props.name] !== undefined
})

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
    // return this.$parent.$parent.service
     return {
        "ports": [],
        "networks": [],
        "depends_on": [],
        "environment": [],
        "volumes": [],
    }
})

const addField = () => {
    // Create the array if not exists.
    if (!service.value[this.name]) {
        service.value[this.name] = []
    }
    array.value.push("")
}

const remove = (index: number) => {
    array.value.splice(index, 1);
}

</script>
