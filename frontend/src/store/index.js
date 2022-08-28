import { createStore } from "vuex";

export default createStore({
  state: {
    slide: null
  },
  getters: {
    returnSlide(state) {
      return state.slide
    }
  },
  mutations: {
    setSlide(state, newSlide) {
      state.slide = newSlide
    }
  },
  actions: {},
  modules: {},
});
