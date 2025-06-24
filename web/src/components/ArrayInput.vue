<template>
    <div>
        <div v-if="valid">
            <ul v-if="isArrayInitd" class="bg-[#070a10] p-10px pl-0">
                <li v-for="(_, index) in array" :key="index" class="flex items-center">
                    <input v-model="array[index]" type="text"
                        class="bg-transparent flex-grow border-none outline-none color-black bg-[#070a10] placeholder-[#1d2634]"
                        :placeholder="placeholder" />
                    <i-lucide-x @click="remove(index)" class="h-16px w-16px" />
                </li>
            </ul>

            <button class="mt-3" @click="addField">
                "addListItem", {{ displayName }}
            </button>
        </div>
        <div v-else>
            LongSyntaxNotSupported
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



const isArrayInitd = computed(() => {
    return service.value[props.name] !== undefined;
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
        service.value[props.name] = [];
    }

    array.value.push("");
}
const remove = (index: number) => {
    array.value.splice(index, 1);
}

</script>
