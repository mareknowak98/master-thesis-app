<template>
  <Editor v-model="slide.slideContent" editorStyle="height: 320px" readonly="true">
    <template #toolbar>
		<span class="ql-formats">
		</span>
    </template>
  </Editor>
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
    let slideId = ref('')
    let slide = ref('')

    const route = useRoute()

    onMounted(() => {
      presentationId.value = route.params.presentationId
      slideId.value = route.params.slideId
      getSlide()
    })

    function getSlide() {
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      const params = new URLSearchParams({
        lesson: presentationId.value,
        slide: slideId.value
      }).toString()

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'slides?' +params , config).then(resp => {
        slide.value = resp.data[0]
      }).catch(err => {
        console.log(err)
      })
    }

    return {
      presentationId,
      slideId,
      slide
    }
  }
};


</script>

