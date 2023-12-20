<script setup lang="ts">
import { ref } from 'vue'

const countdown = ref<string>("")
const socket = new WebSocket(import.meta.env.VITE_WS_URL)
socket.onmessage = (event) => {
  countdown.value = event.data
}

</script>

<template>
  <div class="body">
    <div class="content">
  <h1 class="countdown">{{ countdown }}</h1>
  <img src="../../src/assets/barbu.gif" alt="barbu">
  <div style="padding-top: 20px;">
    <audio controls>
    <source src="../../src/assets/memories.mp3" type="audio/mpeg">
  </audio>
  </div>
</div>
</div>
</template>

<style scoped>

h1.countdown {
  font-size: 3.2em;
  line-height: 1.1;
  color:#57e389
}


</style>
