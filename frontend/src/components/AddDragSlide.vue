<template>
  <div id="textinput">
    <h2>TBD</h2>
  </div>
</template>



<script>


import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import { useRoute } from "vue-router";
import Editor from 'primevue/editor';

import router from "@/router";

export default {
  name: "AddDragSlide",
  components: {
    Editor
  },
  setup() {
    let presentationId = ref('')
    let slideText = ref('')
    let currentPresentationSlides = ref(null)

    const route = useRoute()

    onMounted(() => {
      presentationId.value = route.params.presentationId
      getSlides()
    })

    function getSlides() {
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      const params = new URLSearchParams({
        lesson: presentationId.value
      }).toString()

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'slides?' +params , config).then(resp => {
        currentPresentationSlides.value = resp.data
      }).catch(err => {
        console.log(err)
      })
    }


    return {
      presentationId,
      slideText
    }
  }
};


</script>

