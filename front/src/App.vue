<script setup lang="ts">
import Navbar from './components/Navbar.vue'
import { computed } from 'vue'
import { useRoute } from 'vue-router'
const route = useRoute()
const currentRouteName = computed(() => route.name)
</script>

<template>
  <template v-if="currentRouteName !== 'home'">
    <Navbar />  
  </template>
  <RouterView />
</template>

<style scoped>

</style>
