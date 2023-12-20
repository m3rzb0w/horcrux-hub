<script setup lang="ts">
import { useRouter } from 'vue-router'
import { computed } from 'vue'
import Icons from '../assets/Icons.vue'

const router = useRouter()
const currentRouter= computed(() => router.options.routes)
const routeFiltered = currentRouter.value.filter(el => el.name !== 'home')

</script>

<template>
    <div class="body">
        <div class="content">
            <img src="../../src/assets/horcruxhub_logo.png" alt="logo">
            <hr>
            <ul class="horizontal-list">
                <li v-for="route in routeFiltered">
                    <RouterLink :to="route.path">
                        <Icons :iconName="(route.name as string)" />
                    </RouterLink>
                </li>
            </ul>
        </div>
    </div>
</template>

<style scoped>
ul.horizontal-list {
    list-style: none;
}

ul.horizontal-list li {
    display: inline-block;
}

ul.horizontal-list li {
    padding: 0 0.5rem;
}

ul.horizontal-list li {
    min-width: 12rem;
    text-align: center;
}


</style>