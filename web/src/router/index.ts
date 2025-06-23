import { createRouter, createWebHashHistory, type RouteRecordRaw } from "vue-router";

const routes: RouteRecordRaw[] = [
    {
        path: "/empty",
        component: () => import("@/layouts/Layout.vue"),
        children: [
            {
                path: "",
                component: () => import("@/views/Dashboard.vue"),
                children: [
                    {
                        name: "DashboardHome",
                        path: "/",
                        component:  () => import("@/views/DashboardHome.vue"),
                        children: [
                            {
                                path: "/compose",
                                component:   () =>import("@/views/Compose.vue"),
                            },
                            {
                                path: "/compose/:stackName/:endpoint",
                                component: () => import("@/views/Compose.vue"),
                            },
                            {
                                path: "/compose/:stackName",
                                component:  () =>import("@/views/Compose.vue"),
                            },
                            {
                                path: "/terminal/:stackName/:serviceName/:type",
                                component:  () =>import("@/views/ContainerTerminal.vue"),
                                name: "containerTerminal",
                            },
                            {
                                path: "/terminal/:stackName/:serviceName/:type/:endpoint",
                                component:  () => import("@/views/ContainerTerminal.vue"),
                                name: "containerTerminalEndpoint",
                            },
                        ]
                    },
                    {
                        path: "/console",
                        component: () => import("@/views/Console.vue"),
                    },
                    {
                        path: "/console/:endpoint",
                        component: () => import("@/views/Console.vue"),
                    },
                    {
                        path: "/settings",
                        component:  () =>import("@/views/Settings.vue"),
                        children: [
                            {
                                path: "general",
                                component: () => import("@/components/settings/General.vue"),
                            },
                            {
                                path: "appearance",
                                component: () => import("@/components/settings/Appearance.vue"),
                            },
                            // {
                            //     path: "security",
                            //     component: import("@/components/settings/Security.vue"),
                            // },
                            {
                                path: "about",
                                component: () => import("@/components/settings/About.vue"),
                            },
                        ]
                    },
                ]
            },
        ]
    },
    {
        path: "/setup",
        component: () => import("@/views/Setup.vue"),
    },
    {
        path: "/:catchAll(.*)",
        component: () => import("@/views/404.vue"),
    },

]


const router = createRouter({
    linkActiveClass: "active",
    routes,
    history: createWebHashHistory(import.meta.env.BASE_URL),
    scrollBehavior: () => {
        return {
            top: 0,
            behavior: "smooth",
        }
    }
})

export default router