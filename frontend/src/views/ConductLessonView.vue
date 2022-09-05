<template>
  <LoggedNavbar/>
  <div v-if="this.group == 'teacher-group'">

  <h2>Conduct lesson: {{presentationId}}</h2>
  <div class="row">
    <div class="column">
      <h2>Lesson started</h2>
    </div>
    <div class="columnslim">
      <Divider layout="vertical" />
    </div>

    <div class="column">
      <h2>placehoder</h2>

    </div>
  </div>

  </div>

</template>



<script>

import DataTable  from 'primevue/datatable';
import Column from 'primevue/column'
import Divider from 'primevue/divider'

import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import { useRoute } from "vue-router";

export default {
  name: "ConductLessonView",
  components: {
    LoggedNavbar,
    DataTable,
    Column,
    Divider
  },
  setup() {
    let lessonsSlides = ref(null)
    let presentationId = ref('')
    let socket = ref(null)
    let msg = ref(null)
    const route = useRoute()
    let group = ref("")
    let decodedToken = ref(null)


    onMounted(() => {
      presentationId.value = route.params.presentationId
      getSlides()
      decodedToken.value = TokenService.decodeToken(TokenService.getToken())
      group.value = decodedToken.value['cognito:groups'][0]

      try {
        const params = new URLSearchParams({
          lessonId: presentationId.value
        }).toString()

        console.log(params)
        socket.value = new WebSocket(process.env.VUE_APP_WEBSOCKET_LESSON_API + '?' + params)

        // TODO only for students for receiving data
        // socket.value.onmessage = (m) => {
        //   console.log("eee")
        //   console.log(m.data)
        //   msg.value = JSON.parse(m.data)
        // }

      } catch (e) {
        console.log(e)
      }


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
        lessonsSlides.value = resp.data
      }).catch(err => {
        console.log(err)
      })
    }

    return {
      lessonsSlides,
      presentationId,
      msg,
      group
    }
  }
};


</script>

