import Countdown from '../components/Countdown.vue'
import Home from '../components/Home.vue'
import Chat from '../components/Chat.vue'
import { createRouter, createWebHashHistory } from 'vue-router'


const routes = [
    {
        path: '/',
        name: 'home',
        component: Home
    },
    {
        path: '/countdown',
        name: 'countdown',
        component: Countdown
    },
    {
        path: '/chat',
        name: 'chat',
        component: Chat
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router