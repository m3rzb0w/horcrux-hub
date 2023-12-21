<script setup lang="ts">
import { ref, reactive, computed, onMounted, onUnmounted } from 'vue'

interface Message {
    username: string,
    message: string,
}

const userName = ref<string>("")
const userMessage = ref<string>("")
const input = ref<HTMLElement>()
const allMessages = reactive<Array<Message>>([]);
const socket = new WebSocket(import.meta.env.VITE_WS_CHAT)

socket.onmessage = (event) => {
    allMessages.push(JSON.parse(event.data))
}

const reversedMessages = computed(() => allMessages.slice().reverse());


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

const focusOnInput = () => {
    input.value?.focus()
}

const handleKeyboard = (event: any) => {
    if (event.key === "Enter") {
        sendMessage()
    }
}
onMounted(() => {
    document.addEventListener("keyup", handleKeyboard)
})
onUnmounted(() => {
    document.removeEventListener("keyup", handleKeyboard)
})

onMounted(() =>{
    focusOnInput()
})


</script>

<template>
    <div id="messages">
        <div class="" v-for="(message, index) in reversedMessages" :key="index">
            <strong>{{ message.username }}</strong>: {{ message.message }}
        </div>
    </div>
        <div class="msg-inputs">
            <input class="input" type="text" id="userName" v-model="userName" placeholder="Nickname" />
            <input class="input is-primary is-normal" type="text" placeholder="Enter your message" v-model="userMessage" ref="input">
            <button class="button is-primary is-outlined" @click="sendMessage">Send</button>
        </div>
</template>

<style scoped>

#userName {
    width: 10%;
}

#messages {
    margin: auto;
    /* margin-top: 5vh; */
    background-color: rgb(190, 199, 199);
    display: flex;
    flex-direction: column-reverse;
    height: 90%;
    width: 100%;
    overflow-y: scroll;
    padding-left: 5px;
    word-wrap: break-word;
    position: fixed;
}

.msg-inputs{
   position: fixed;
   bottom: 0;
   display: flex;
   width: 100%;
}
        

  
</style>