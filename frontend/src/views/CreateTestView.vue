<template>
  <LoggedNavbar/>
  <h2>Enter question:</h2>
  <div class="row">
    <div class="column">

    <InputText type="text" v-model="questionText" />

    <h2>Enter possible answers:</h2>
    <InputText type="text" v-model="answerText" />
    <Button @click="addAnswer" label="Add answer" />
    <p> </p>

    <DataTable :value="answerList" v-model:selection="answerCorrect" selectionMode="single" dataKey="id" responsiveLayout="scroll" >
      <Column field="text" header="Select correct answer"></Column>
    </DataTable>

    <Button @click="addQuestion" label="Add question" />


    </div>
  <div class="columnslim">
    <Divider layout="vertical" />
  </div>
  <div class="column">
    <li v-for="question in questions">
      {{ question }}
    </li>
    <Button @click="saveTest" label="Save test" />

  </div>
  </div>
</template>



<script>

import DataTable  from 'primevue/datatable';
import Column from 'primevue/column'
import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import Divider from 'primevue/divider'
import { useRoute } from "vue-router/dist/vue-router";
import router from "@/router";

export default {
  name: "CreateTestView",
  components: {
    LoggedNavbar,
    DataTable,
    Column,
    Divider
  },
  setup() {
    let questionText = ref('')
    let answerText = ref('')
    let answerList= ref([])
    let answerCorrect = ref([])
    let testId = ref('')
    const route = useRoute()

    let questions = ref([])
    let questionId = ref(0)

    let tmpId = ref(1)

    onMounted(() => {
      testId.value = Math.floor(Math.random() * 10000000)
    })

    function addAnswer() {
      let obj = {
        id: Math.floor(Math.random() * 100000),
        text: answerText.value
      }
      tmpId.value++
      answerList.value.push(obj)
      answerText.value = ''
    }

    function addQuestion() {
      questionId.value++
      let tmp = route.params.classId + ':' + questionId.value
      let tmp2 = answerList.value.map(function(item) {
        return item['text'];
      });


      let obj = {
        testId : testId.value.toString(),
        combinedKey : tmp,
        question : questionText.value,
        answers : tmp2,
        correctAnswer : answerCorrect.value.text
      }

      questions.value.push(obj)
      answerText.value = ''
      answerCorrect.value = []
      questionText.value = ''
      answerList.value = []
    }

    function saveTest(){
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      let promises = [];
      for (let elem of questions.value) {
        promises.push(
          axios.post(process.env.VUE_APP_BACKEND_RESP_API + 'tests', elem, config)
        )
      }

      Promise.all(promises)
        .then(responses => {
          console.log(responses)
          router.push({name: "createTests"})
        })
        .catch(responses => console.log(responses));
    }

      return {
      questionText,
      answerText,
      answerCorrect,
      answerList,
      addAnswer,
      addQuestion,
      questionId,
      questions,
      saveTest
    }
  }
};


</script>

<style lang="scss"></style>