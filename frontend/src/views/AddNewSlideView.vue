<template>
  <LoggedNavbar/>
  <h3>Select slide type</h3>
  <Dropdown v-model="selectedSlide" :options="slideTypes" optionLabel="name" optionValue="code" placeholder="Select slide type" />

  <div v-if="this.selectedSlide == 'TEXT'">
    <AddTextSlide/>
  </div>
  <div v-if="this.selectedSlide == 'QA'">
    <AddQASlide/>
  </div>
</template>



<script>


import Dropdown from 'primevue/dropdown'

import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import { useRoute } from "vue-router";
import AddTextSlide from "@/components/AddTextSlide";
import AddQASlide from "@/components/AddQASlide";

import router from "@/router";

export default {
  name: "AddNewSlideView",
  components: {
    LoggedNavbar,
    Dropdown,
    AddTextSlide,
    AddQASlide
  },
  setup() {
    let slide = ref(null)
    let presentationId = ref('')
    const slideTypes = ref([
      {name: 'Text slide', code: 'TEXT'},
      {name: 'Questions', code: 'QA'},
      {name: 'Drag and drop', code: 'DRAG'},
    ]);
    let selectedSlide = ref('')

    const route = useRoute()

    onMounted(() => {
      presentationId.value = route.params.presentationId

    })




    return {
      slide,
      slideTypes,
      selectedSlide
    }
  }
};


</script>

