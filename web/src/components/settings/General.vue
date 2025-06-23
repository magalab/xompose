<template>
    <div>
        <form class="my-4" autocomplete="off" @submit.prevent="saveGeneral">
            <!-- Client side Timezone -->
            <!-- <div v-if="false" class="mb-4">
                <label for="timezone" class="form-label">
                    {{ $t("Display Timezone") }}
                </label>
                <select id="timezone" v-model="$root.userTimezone" class="form-select">
                    <option value="auto">
                        {{ $t("Auto") }}: {{ guessTimezone }}
                    </option>
                    <option
                        v-for="(timezone, index) in timezoneList"
                        :key="index"
                        :value="timezone.value"
                    >
                        {{ timezone.name }}
                    </option>
                </select>
            </div> -->

            <!-- Server Timezone -->
            <!-- <div v-if="false" class="mb-4">
                <label for="timezone" class="form-label">
                    {{ $t("Server Timezone") }}
                </label>
                <select id="timezone" v-model="settings.serverTimezone" class="form-select">
                    <option value="UTC">UTC</option>
                    <option
                        v-for="(timezone, index) in timezoneList"
                        :key="index"
                        :value="timezone.value"
                    >
                        {{ timezone.name }}
                    </option>
                </select>
            </div> -->

            <!-- Primary Hostname -->
            <div class="mb-4">
                <label class="form-label" for="primaryBaseURL">
                    primaryHostname 
                </label>

                <div class="input-group mb-3">
                    <input
                        v-model="settings.primaryHostname"
                        class="form-control"
                        placeholder="未设置:沿用当前主机名"
                    />
                    <button class="btn btn-outline-primary" type="button" @click="autoGetPrimaryHostname">
                       自动获取
                    </button>
                </div>

                <div class="form-text"></div>
            </div>

            <!-- Save Button -->
            <div>
                <button class="btn btn-primary" type="submit">
                    保存
                </button>
            </div>
        </form>
    </div>
</template>

<script setup lang="ts">

// import dayjs from "dayjs";
// import { timezoneList } from "@/utils/time";

// const tList = ref({})

const settings = computed ( () => {
    // return this.$parent.$parent.$parent.settings;
    return {
        primaryHostname: "primaryHostname",
    }
})

const saveSettings = () => {
    // this.$parent.$parent.$parent.saveSettings();
}
const settingsLoaded = computed ( () => {
    // return this.$parent.$parent.$parent.settingsLoaded;
})


// const guessTimezone = computed ( () => {
//     return dayjs.tz.guess();
// })

const saveGeneral = () => {
    // localStorage.timezone = this.$root.userTimezone;
    // saveSettings();
}

const autoGetPrimaryHostname = () => {
    settings.primaryHostname = location.hostname;
}

</script>

