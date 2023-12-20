<script setup lang="ts">
import { useRouter } from 'vue-router'
import { computed, markRaw, ref, onMounted, onUnmounted } from 'vue'
import ModalAbout from './ModalAbout.vue'
import { useModal } from '../composables/useModal'
import { onClickOutside } from '@vueuse/core'


const modal = useModal()

const openConfirm = () => {
    console.log("toto")
    modal.component.value = markRaw(ModalAbout)
    modal.showModal()
}

const closeModal = () => {
    modal.hideModal()
}

//click outside to close modal
const modalCardRef = ref(null)

onClickOutside(modalCardRef, closeModal)

//escapte to close modal
const handleKeyboard = (event: any) => {
    if (event.key === "Escape") {
        closeModal()
    }
}
onMounted(() => {
    document.addEventListener("keyup", handleKeyboard)
})
onUnmounted(() => {
    document.removeEventListener("keyup", handleKeyboard)
})

const router = useRouter()
const currentRouter = computed(() => router.options.routes)
const routeFiltered = currentRouter.value.filter(el => el.name !== 'home')
</script>

<template>
    <nav class="navbar is-light" role="navigation" aria-label="main navigation">
        <div class="navbar-brand">
            <RouterLink class="navbar-item" to="/">
                <img src="../../src/assets/horcruxhub_logo.png">
            </RouterLink>
        </div>

        <div id="navbarBasicExample" class="navbar-menu">
            <div class="navbar-start">
                <div class="navbar-item has-dropdown is-hoverable">
                    <a class="navbar-link is-arrowless">
                        #projects
                    </a>
                    <div class="navbar-dropdown">
                        <div v-for="route in routeFiltered">
                            <RouterLink :to="route.path">
                                <a class="navbar-item">{{ route.name }}</a>
                            </RouterLink>
                        </div>
                    </div>
                </div>
                <a class="navbar-item" @click="openConfirm">
                    about
                </a>
            </div>
        </div>
    </nav>
    <Teleport to="#modal" ref="modalCardRef">
        <Transition>
            <component :is="modal.component.value" v-if="modal.show.value" @close="modal.hideModal" />
        </Transition>
    </Teleport>
</template>

<style scoped>
.v-enter-active,

.v-leave-active {
    transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
    opacity: 0;
}
</style>