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
        <iframe style="border: 0; width: 100%; height: 42px;" src="https://bandcamp.com/EmbeddedPlayer/album=3942234626/size=small/bgcol=333333/linkcol=ffffff/artwork=none/track=4003073828/transparent=true/" seamless><a href="https://holedweller.bandcamp.com/album/flies-the-coop">flies the coop by Hole Dweller</a></iframe>
      </div>
    </div>
  </div>
</template>

<style scoped>
h1.countdown {
  font-size: 3.2em;
  line-height: 1.1;
  color: #57e389
}

div.body {
  margin: 0;
  display: flex;
  place-items: center;
  min-width: 320px;
  min-height: 80vh;
}
</style>
