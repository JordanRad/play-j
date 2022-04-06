import { defineStore } from 'pinia'

export const useCounterStore = defineStore({
  id: 'user',
  state: () => ({
    activeUser: null
  }),
  getters: {
    email: (state) => state.counter * 2
  },
  actions: {
    increment() {
      this.counter++
    }
  }
})
