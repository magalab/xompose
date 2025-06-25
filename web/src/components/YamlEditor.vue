<script setup lang="ts">
// 引入所需的 CodeMirror 模块
import { basicSetup } from 'codemirror'
import { EditorView } from '@codemirror/view'
import { Compartment, EditorState } from '@codemirror/state'
import { yaml } from '@codemirror/lang-yaml'

// 初始化编辑器内容
const defaultTemplate = `services:
  nginx:
    image: nginx:latest
    restart: unless-stopped
      ports:
      - "8080:80"
    `
const code = defineModel(
    {
        type: String,
        default: defaultTemplate,
    }
)

const view = ref()
const viewRef = ref<Element>()

// 初始化编辑器实例
const init = () => {
    if (view.value) return

    const language = new Compartment()
    const tabSize = new Compartment()

    // 创建 CodeMirror 编辑器状态
    const state = EditorState.create({
        doc: code.value,
        extensions: [
            basicSetup,
            language.of(yaml()), 
            tabSize.of(EditorState.tabSize.of(2)),
            EditorView.updateListener.of(v => {
                code.value = v.state.doc.toString()
            }),
            EditorView.theme({
                '&': { maxHeight: `400px`, overflow: 'auto' },
                '.cm-scroller': { maxHeight: `400px` },
            }),
        ],
    })

    view.value = new EditorView({
        state,
        parent: viewRef.value,
    })
}

// 在组件挂载时初始化编辑器
onMounted(() => {
    init()
})
</script>

<template>
    <div ref="viewRef" ></div>
</template>

<style lang="scss" scoped>
/* 样式可以根据需要调整 */
</style>
