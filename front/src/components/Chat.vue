<script setup lang="ts">
import { ref, reactive } from 'vue'

interface Message {
    username: string,
    message: string
}

const userName = ref<string>("")
const userMessage = ref<string>("")
const allMessages = reactive<Array<Message>>([]);
const socket = new WebSocket(import.meta.env.VITE_WS_CHAT)

socket.onmessage = (event) => {
    allMessages.push(JSON.parse(event.data))
}

const userNameCheck = () => {
    if (userName.value.trim() == "") {
        const num = Math.floor(Math.random() * 10000)
        userName.value = "anon" + num
    }
}

const sendMessage = () => {
    if (userMessage.value.trim().length !== 0) {
        userNameCheck()
        const newMessage: Message = {
            username: userName.value,
            message: userMessage.value,
        }
        socket.send(JSON.stringify(newMessage))

        userMessage.value = ""
    }
}


</script>

<template>
    <div class="body">
        <div class="content">
            <h1 class="wip">
                WIP
            </h1>
            <div>

                <h2>All Messages</h2>
                <ul>
                    <li v-for="(message, index) in allMessages" :key="index">
                        <strong>{{ message.username }}</strong>: {{ message.message }}
                    </li>
                </ul>
            </div>
            <div>
                <label for="username">Username:</label>
                <input type="text" id="username" v-model="userName" />
            </div>
            <div>
                <label for="message">Message:</label>
                <input type="text" id="message" v-model="userMessage" />
            </div>
            <button @click="sendMessage">Add Message</button>
        </div>
    </div>
</template>

<style scoped>
h1.wip {
    font-size: 3.2em;
    line-height: 1.1;
    color: #57e389
}
</style>