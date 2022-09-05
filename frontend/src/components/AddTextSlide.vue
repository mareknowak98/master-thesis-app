<template>
  <div id="textinput">
    <Editor v-model="slideText" editorStyle="height: 320px"/>
    <Button @click="addTextSlide" label="Add text slide" />
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
  name: "AddTextSlide",
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


    function addTextSlide() {
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }

      let maxSlide = Math.max.apply(Math, currentPresentationSlides.value.map(function(o) { return parseInt(o.slideId); }))
      maxSlide++

      if (maxSlide < 1 || maxSlide > 1000){
        maxSlide = 1
      }

      let m = {
        lessonId: presentationId.value,
        slideId: maxSlide.toString(),
        slideType: "TEXT",
        slideContent: slideText.value,
        questionAnswers: ""
      }


      axios.post(process.env.VUE_APP_BACKEND_RESP_API + 'slides', m, config).then(resp => {
        router.push({name: "managePresentation"})
      }).catch(err => {
        console.log(err)
      })
    }


    return {
      presentationId,
      slideText,
      addTextSlide,
      currentPresentationSlides
    }
  }
};


</script>

