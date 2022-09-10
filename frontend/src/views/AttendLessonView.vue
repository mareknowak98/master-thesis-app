<template>
  <LoggedNavbar/>
    <h2>Lesson: {{presentationId}}</h2>

  <div v-if="this.currentSlideType == 'TEXT'">
    <Editor v-model="currentSlide.slideContent" editorStyle="height: 320px" readonly="true">
      <template #toolbar>
		<span class="ql-formats">
		</span>
      </template>
    </Editor>
  </div>

  <div v-if="this.currentSlideType == 'QA'">
    <h4>Question:</h4>
    <h2>{{this.currentSlide.slideContent}}</h2>
    <DataTable :value="answerList" v-model:selection="selectedAnswerList" selectionMode="multiple" dataKey="id" responsiveLayout="scroll" >
      <Column field="text" header="Select correct answers(multiple select with Alt)"></Column>
    </DataTable>
    <Button @click="checkAnswers" label="Check your answer" />
  </div>


  <div v-if="this.isCorrectAnswer === true">
    <Message severity="success">Correct</Message>
  </div>
  <div v-if="this.isCorrectAnswer === false">
    <Message severity="error">Try again</Message>
  </div>


</template>
<script>



import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import { useRoute } from "vue-router";
import Editor from "primevue/editor";
import Column from "primevue/column";
import DataTable from "primevue/datatable";
import Message from 'primevue/message';

export default {
  name: "AttendLessonView",
  components: {
    LoggedNavbar,
    Editor,
    Column,
    DataTable,
    Message
  },
  setup() {
    let lessonsSlides = ref(null)
    let presentationId = ref('')
    let currentSlide = ref(null)
    let currentSlideType = ref(null)

    let socket = ref(null)
    let msg = ref(null)
    const route = useRoute()

    let group = ref("")
    let decodedToken = ref(null)

    let answerList = ref(null)
    let correctAnswerList = ref(null)
    let selectedAnswerList = ref(null)

    let isCorrectAnswer = ref(null)

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

        socket.value.onmessage = (m) => {
          currentSlide.value = JSON.parse(m.data)
          currentSlideType.value = JSON.parse(m.data).slideType
          answerList.value = JSON.parse(JSON.parse(m.data).questionAnswers).answers
          correctAnswerList.value = JSON.parse(JSON.parse(m.data).questionAnswers).correctAnswers
        }

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
      }).then(resp => {
        if(currentSlide.value == null){
          currentSlide.value = lessonsSlides.value[0]
        }
      }).then(resp => {
        currentSlideType.value = currentSlide.value.slideType
      }).then(resp =>{
        currentSlide.value.questionAnswers = JSON.parse(currentSlide.value.questionAnswers)
      }).then(resp =>{
        answerList.value = currentSlide.value.questionAnswers.answers
        correctAnswerList.value = currentSlide.value.questionAnswers.correctAnswers
      })
        .catch(err => {
        console.log(err)
      })
    }

    const objectsEqual = (o1, o2) =>
      Object.keys(o1).length === Object.keys(o2).length
      && Object.keys(o1).every(p => o1[p] === o2[p]);

    const arraysEqual = (a1, a2) =>
      a1.length === a2.length && a1.every((o, idx) => objectsEqual(o, a2[idx]));

    function checkAnswers(){
      isCorrectAnswer.value = arraysEqual(correctAnswerList.value, selectedAnswerList.value)
    }

    return {
      lessonsSlides,
      presentationId,
      msg,
      group,
      currentSlide,
      currentSlideType,
      answerList,
      correctAnswerList,
      selectedAnswerList,
      checkAnswers,
      isCorrectAnswer
    }
  }
};


</script>

