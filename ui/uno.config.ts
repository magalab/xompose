// uno.config.ts
import { defineConfig, presetAttributify, presetWind3, transformerVariantGroup, transformerDirectives } from 'unocss'
import { presetIcons } from '@unocss/preset-icons'
import { presetRemToPx } from '@unocss/preset-rem-to-px'


export default defineConfig({
    presets: [
        presetWind3(),
        presetAttributify(),
        presetIcons(),
        presetRemToPx(),
    ],
    transformers: [
        transformerVariantGroup(),
        transformerDirectives(),
    ]
})