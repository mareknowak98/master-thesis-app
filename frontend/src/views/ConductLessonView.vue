<template>
  <LoggedNavbar/>
  <div v-if="this.group == 'teacher-group'">

  <h2>Conduct lesson: {{presentationId}}</h2>
  <div class="row">
    <div class="column">
      <Button @click="previousSlide" label="< Previous slide" />
      <Button @click="nextSlide" label="Next slide >" />

    </div>
    <div class="columnslim">
      <Divider layout="vertical" />
    </div>

    <div class="column">
      <div v-if="this.currentSlideType == 'TEXT'">
        <Editor v-model="currentSlide.slideContent" editorStyle="height: 320px" readonly="true">
          <template #toolbar>
            <span class="ql-formats">
            </span>
          </template>
        </Editor>
      </div>

      <div v-if="this.currentSlideType == 'QA'">
        <h2>Question:</h2>
        <h3>{{this.currentSlide.slideContent}}</h3>

        <DataTable :value="answerList" disabled="true" selectionMode="multiple" dataKey="id" responsiveLayout="scroll" >
          <Column field="text" header="Select correct answers(multiple select with Alt)"></Column>
        </DataTable>

      </div>

    </div>
  </div>

  </div>

</template>



<script>

import DataTable  from 'primevue/datatable';
import Column from 'primevue/column'
import Divider from 'primevue/divider'
import Editor from "primevue/editor";

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
    Divider,
    Editor
  },
  setup() {
    let lessonsSlides = ref(null)
    let presentationId = ref('')
    let socket = ref(null)
    let msg = ref(null)
    const route = useRoute()
    let group = ref("")
    let decodedToken = ref(null)

    let currentSlide = ref(null)
    let currentSlideType = ref(null)
    let answerList = ref(null)

    let slidesIds = ref([])
    let slideI = ref(0)


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

    function nextSlide(){
      if (slideI.value < (slidesIds.value.length - 1)){
        slideI.value++
      }
      let m = {
        lessonId: presentationId.value,
        slideId: slidesIds.value[slideI.value]
      }
      // console.log(m)
      socket.value.send(JSON.stringify(m))

      currentSlide.value = lessonsSlides.value[slideI.value]
      currentSlideType.value = currentSlide.value.slideType
      currentSlide.value.questionAnswers = JSON.parse(currentSlide.value.questionAnswers)
      answerList.value = currentSlide.value.questionAnswers.answers
      slidesIds.value = lessonsSlides.value.map(function(item) {
        return item['slideId'];
      });
    }

    function previousSlide(){
      if (slideI.value >= 1){
        slideI.value--
      }
      let m = {
        lessonId: presentationId.value,
        slideId: slidesIds.value[slideI.value]
      }
      // console.log(m)
      socket.value.send(JSON.stringify(m))

      currentSlide.value = lessonsSlides.value[slideI.value]
      currentSlideType.value = currentSlide.value.slideType
      currentSlide.value.questionAnswers = JSON.parse(currentSlide.value.questionAnswers)
      answerList.value = currentSlide.value.questionAnswers.answers
      slidesIds.value = lessonsSlides.value.map(function(item) {
        return item['slideId'];
      });
    }

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
      }).then(resp => {
        if(currentSlide.value == null){
          currentSlide.value = lessonsSlides.value[0]
        }
      }).then(resp => {
        currentSlideType.value = currentSlide.value.slideType
      }).then(resp =>{
        if (currentSlideType.value === "QA") {
          currentSlide.value.questionAnswers = JSON.parse(currentSlide.value.questionAnswers)
        }
      }).then(resp =>{
        answerList.value = currentSlide.value.questionAnswers.answers
        slidesIds.value = lessonsSlides.value.map(function(item) {
          return item['slideId'];
        });
      })
        .catch(err => {
          console.log(err)
        })
    }

    return {
      lessonsSlides,
      presentationId,
      msg,
      group,
      currentSlide,
      currentSlideType,
      answerList,
      nextSlide,
      previousSlide,
      slidesIds,
      slideI
    }
  }
};


</script>

