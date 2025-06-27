<template>
    <ul class="h-300px list-none overflow-auto">
        <li v-for="(log , index) in props.logList" :key="index" class="flex items-center justify-start m-1px">
            <span v-html="formatLog(log)"></span>
        </li>
    </ul>
</template>

<script setup lang="ts">

const props = defineProps({
    logList: {
        type: [] as PropType<string[]>,
    }
})

const ansiColors: { [code: string]: string } = {
    '30': 'text-black',        // 黑色
    '31': 'text-red-500',      // 红色
    '32': 'text-green-500',    // 绿色
    '33': 'text-yellow-500',   // 黄色
    '34': 'text-blue-500',     // 蓝色
    '35': 'text-purple-500',   // 紫色
    '36': 'text-cyan-500',     // 青色
    '37': 'text-gray-300',     // 浅灰色
    '38': 'text-gray-400',     // 深灰色
    '90': 'text-gray-500',     // 深灰色
    '91': 'text-red-400',      // 浅红色
    '92': 'text-green-400',    // 浅绿色
    '93': 'text-yellow-400',   // 浅黄色
    '94': 'text-blue-400',     // 浅蓝色
    '95': 'text-purple-400',   // 浅紫色
    '96': 'text-cyan-400',     // 浅青色
    '97': 'text-white',        // 白色
}

const formatLog = (logLine: string) => {
    let html = logLine.replace("{", "")
    html = html.replace(/\x1b\[0m/g, '</span>')
    Object.keys(ansiColors).forEach(code => {
        const ansiPattern = new RegExp(`\\x1b\\[${code}m`, 'g')
        const htmlReplacement = `<span class="${ansiColors[code]}">`
        html = html.replace(ansiPattern, htmlReplacement)
    })

    return html
}

</script>

<style lang="scss" scoped></style>