import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";


const routes : Array<RouteRecordRaw> = [
   
    {
        path: '/login',
        component: () => import('@/views/Login/index.vue'),
    },
    {
        path: '/',
        component: () => import('@/views/Layout/index.vue'),
        children: [
            {
                path: '/stack',
                component: () => import('@/views/Index/index.vue'),
            }
        ]
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes : routes,
})

export {
    router,
}