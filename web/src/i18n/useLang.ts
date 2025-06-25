import { computed } from "vue";
import { useI18n } from "vue-i18n";

export default () => {
    const { locale } = useI18n();
    const langName = computed(() => {
        const langMap: Record<string, string> = {
            'zh': '简体中文',
            'en': 'English',
        }
        return langMap[locale.value] || '简体中文'
    });

    const langChange = (lang: string) => {
        locale.value = lang
        localStorage.setItem('lang', lang)
        window.location.reload()
    }

    return { langChange, langName }
}