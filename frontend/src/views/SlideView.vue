<template>
  <LoggedNavbar/>
  <h2>Edit slide: {{this.$route.params.presentationId}} / {{this.$route.params.slideId}}</h2>

  <div v-if="this.slideType == 'TEXT'">
    <TextSlide/>
  </div>

  <div v-if="this.slideType == 'QA'">
    <QASlide/>
  </div>

  <h3>----------------</h3>
  <Button @click="deleteSlide" label="Delete slide" />

</template>



<script>

import DataTable  from 'primevue/datatable';
import Column from 'primevue/column'

import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import { useRoute } from "vue-router";
import router from "@/router";
import TextSlide from "@/components/TextSlide";
import QASlide from "@/components/QASlide";



export default {
  name: "SlideView",
  components: {
    LoggedNavbar,
    DataTable,
    Column,
    TextSlide,
    QASlide
  },
  setup() {
    let slide = ref(null)
    let presentationId = ref('')
    let slideId = ref('')
    let slideType = ref('')

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
        slide.value = resp.data
        slideType.value = resp.data[0].slideType
      }).catch(err => {
        console.log(err)
      })
    }

    function deleteSlide() {
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      const params = new URLSearchParams({
        lesson: presentationId.value,
        slide: slideId.value
      }).toString()

      axios.delete(process.env.VUE_APP_BACKEND_RESP_API + 'slides?' +params , config).then(resp => {
        router.push({name: "managePresentation"})
      }).catch(err => {
        console.log(err)
      })
    }

      return {
      slide,
      deleteSlide,
      slideType
    }
  }
};


</script>

