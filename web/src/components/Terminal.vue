<template>
    <div class="shadow-[0_15px_70px_rgba(0,0,0,0.1)] p-2.5 rounded-lg">
        <div ref="terminalRef" class="h-full bg-red-3"></div>
    </div>
</template>

<script setup lang="ts">
import '@xterm/xterm/css/xterm.css'
import { Terminal } from "@xterm/xterm"
import { FitAddon } from "@xterm/addon-fit"
import { AttachAddon } from '@xterm/addon-attach'
import { WebLinksAddon } from '@xterm/addon-web-links'
import { TERMINAL_COLS, TERMINAL_ROWS } from "@/utils/terminal"
const terminalRef = ref<HTMLElement>()

const terminalFitAddOn = ref()

// const first = ref(true)
const terminalInputBuffer = ref("")
const cursorPosition = ref(0)

let ws: WebSocket
const props = defineProps({
    name: {
        type: String,
        require: true,
    },

    endpoint: {
        type: String,
        require: true,
    },

    // Require if mode is interactive
    stackName: {
        type: String,
    },

    // Require if mode is interactive
    serviceName: {
        type: String,
    },

    // Require if mode is interactive
    type: {
        type: String,
        default: "bash",
    },

    rows: {
        type: Number,
        default: TERMINAL_ROWS,
    },

    cols: {
        type: Number,
        default: TERMINAL_COLS,
    },

    // Mode
    // displayOnly: Only display terminal output
    // mainTerminal: Allow input limited commands and output
    // interactive: Free input and output
    mode: {
        type: String,
        default: "displayOnly",
    }
})

const terminal = ref<Terminal>()

const emit = defineEmits(["has-data"])

onMounted(() => {
    let cursorBlink = true
    // if (props.mode === "displayOnly") {
    //     cursorBlink = false
    // }
    ws = new WebSocket("ws://127.0.0.1:9527/ws/terminal?container=xxx-nginx-3-1")
    ws.onclose = () => {
        console.log("服务器跑了")
    }
    const attachAddon = new AttachAddon(ws)
    terminal.value = new Terminal({
        fontSize: 14,
        cursorBlink,
        cols: props.cols,
        rows: props.rows,
    })
    terminalFitAddOn.value = new FitAddon()
    terminal.value.loadAddon(terminalFitAddOn.value)
    terminal.value.loadAddon(new WebLinksAddon())
    terminal.value.loadAddon(attachAddon)
    // if (props.mode === "mainTerminal") {
    //     mainTerminalConfig()
    // } else if (props.mode === "interactive") {
    //     interactiveTerminalConfig()
    // }
    terminal.value.open(terminalRef.value!)
    terminalFitAddOn.value.fit()
    // terminal.value.write('stty rows 20 columns 80\r\n')
    terminal.value.focus()

    // Notify parent component when data is received
    // terminal.value.onCursorMove(() => {
    //     console.debug("onData triggered")
    //     if (first.value) {
    //         // this.$emit("has-data")
    //         emit('has-data')
    //         first.value = false
    //     }
    // })

    // // Create a new Terminal
    // if (props.mode === "mainTerminal") {
    //     // this.$root.emitAgent(props.endpoint, "mainTerminal", props.name, (res) => {
    //     //     if (!res.ok) {
    //     //         this.$root.toastRes(res)
    //     //     }
    //     // })
    // } else if (props.mode === "interactive") {
    //     console.debug("Create Interactive terminal:", props.name)
    //     // this.$root.emitAgent(props.endpoint, "interactiveTerminal", props.stackName, props.serviceName, props.shell, (res) => {
    //     //     if (!res.ok) {
    //     //         this.$root.toastRes(res)
    //     //     }
    //     // })
    // }
    // updateTerminalSize()

})

onUnmounted(() => {
    // window.removeEventListener("resize", onResizeEvent)
    // this.$root.unbindTerminal(this.name)
    if(ws){
        ws.close()
    }
    terminal.value!.dispose()
})

const mainTerminalConfig = () => {
    terminal.value!.onKey(e => {
        const code = e.key.charCodeAt(0)
        console.debug("Encode: " + JSON.stringify(e.key))

        if (e.key === "\r") {
            // Return if no input
            if (terminalInputBuffer.value.length === 0) {
                return
            }

            // const buffer = terminalInputBuffer.value

            // Remove the input from the terminal
            removeInput()

            // TODO
            // this.$root.emitAgent(this.endpoint, "terminalInput", this.name, buffer + e.key, (err) => {
            //     this.$root.toastError(err.msg)
            // })

        } else if (code === 127) { // Backspace
            if (cursorPosition.value > 0) {
                terminal.value!.write("\b \b")
                cursorPosition.value--
                terminalInputBuffer.value = terminalInputBuffer.value.slice(0, -1)
            }
        } else if (e.key === "\u001B\u005B\u0041" || e.key === "\u001B\u005B\u0042") {      // UP OR DOWN
            // Do nothing

        } else if (e.key === "\u001B\u005B\u0043") {      // RIGHT
            // TODO
        } else if (e.key === "\u001B\u005B\u0044") {      // LEFT
            // TODO
        } else if (e.key === "\u0003") {      // Ctrl + C
            console.debug("Ctrl + C")
            // TODO
            // this.$root.emitAgent(this.endpoint, "terminalInput", this.name, e.key)
            removeInput()
        } else {
            cursorPosition.value++
            terminalInputBuffer.value += e.key
            terminal.value!.write(e.key)
        }
    })
}
const interactiveTerminalConfig = () => {
    terminal.value!.onKey((e: any) => {
        // TODO 发送终端交互
        // this.$root.emitAgent(this.endpoint, "terminalInput", this.name, e.key, (res) => {
        //     if (!res.ok) {
        //         this.$root.toastRes(res)
        //     }
        // })
    })
}

const removeInput = () => {
    const backspaceCount = terminalInputBuffer.value.length
    const backspaces = "\b \b".repeat(backspaceCount)
    cursorPosition.value = 0
    terminal.value!.write(backspaces)
    terminalInputBuffer.value = ""
}

// const updateTerminalSize = () => {
//     if (!terminalFitAddOn.value) {
//         terminalFitAddOn.value = new FitAddon()
//         terminal.value!.loadAddon(terminalFitAddOn.value)
//         window.addEventListener("resize", onResizeEvent)
//     }
//     terminalFitAddOn.value.fit()
// }
// const onResizeEvent = () => {
//     terminalFitAddOn.value.fit()
//     let rows = terminal.value!.rows
//     let cols = terminal.value!.cols
//     if (ws) {
//         ws.send(JSON.stringify({
//             "action": "resize",
//             "rows": rows,
//             "cols": cols,
//         }))
//     }
// }

</script>