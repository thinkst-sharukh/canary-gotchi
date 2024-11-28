import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { IGotchi } from '@/types'

export const useGotchiStore = defineStore('gotchi', () => {
  const gotchi = ref<IGotchi>()
  function setGotchi(g: IGotchi) {
    gotchi.value = g
  }

  return { gotchi,setGotchi }
})
