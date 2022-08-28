<template>
  <h2>Question:</h2>
  <p>{{this.slide.slideContent}}</p>
  <h2>Choose answers:</h2>

  <DataTable :value="answerList" v-model:selection="selectedAnswerList" selectionMode="multiple" dataKey="id" responsiveLayout="scroll" >
    <Column field="text" header="Select correct answers(multiple select with Alt)"></Column>
  </DataTable>

</template>



<script>


import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import { useRoute } from "vue-router";
import DataTable from 'primevue/datatable';
import Column from 'primevue/column'
import router from "@/router";

export default {
  name: "QASlide",
  components: {
    DataTable,
    Column
  },
  setup() {
    let presentationId = ref('')
    let slideId = ref('')
    let slide = ref('')
    let answerList = ref(null)
    let correctAnswerList = ref(null)
    let selectedAnswerList = ref(null)

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
      }).then(resp => {
        slide.value.questionAnswers = JSON.parse(slide.value.questionAnswers)
      }).then(resp => {
        answerList.value = slide.value.questionAnswers.answers
        correctAnswerList.value = slide.value.questionAnswers.correctAnswers
      }).catch(err => {
        console.log(err)
      })
    }

    return {
      presentationId,
      slideId,
      slide,
      answerList,
      correctAnswerList,
      selectedAnswerList
    }
  }
};


</script>

