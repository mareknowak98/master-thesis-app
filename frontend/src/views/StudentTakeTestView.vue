<template>
  <LoggedNavbar/>
  <h2>Test: {{this.$route.params.testId}}</h2>

  <div v-if="this.isResolved == false">
    <div v-for="(answer, i) in answers">
      <Fieldset :legend="test[i].question">

        <div v-for="q of answer" :key="q.key" class="field-radiobutton">
          <RadioButton :inputId="q.key" name="category" :value="q.name" v-model="selectedAnswer[i]" :disabled="q.key === 'R'" />
          <label :for="q.key">{{q.name}}</label>
        </div>
      </Fieldset>
    </div>
    <Button @click="submitTest" label="Submit" />
  </div>
  <div v-else>
    <h2>You have already solved this test</h2>
  </div>

</template>



<script>

import DataTable  from 'primevue/datatable';
import Column from 'primevue/column'
import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import { useRoute } from "vue-router/dist/vue-router";
import Fieldset from 'primevue/fieldset';
import RadioButton from 'primevue/radiobutton';
import router from "@/router";

export default {
  name: "StudentTakeTestView",
  components: {
    LoggedNavbar,
    DataTable,
    Column,
    Fieldset,
    RadioButton
  },
  setup() {
    let test = ref(null)
    const route = useRoute()
    let answers = ref([])
    let selectedAnswer = ref([])
    let isResolved = ref(false)

    onMounted(() => {
      getTest()
      getUserTestResults()
    })

    function getTest(){
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      const params = new URLSearchParams({
        testId: route.params.testId
      }).toString()

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'tests?' + params, config).then(resp => {
        test.value = resp.data
      }).then(resp=>{
        for (const element of test.value) {
          let tmp = []
          for (const q of element.answers) {
            let obj = {
              name : q,
              key : q,
            }
            tmp.push(obj)
          }
          answers.value.push(tmp)
        }
      }).catch(err => {
        console.log(err)
      })
    }

    function submitTest(){
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      let questionAnswer = []
      for (let i = 0; i < answers.value.length; ++i){
        let tmp = {
          combinedKey: test.value[i].combinedKey,
          correctQuestionAnswer: selectedAnswer.value[i]
        }
        questionAnswer.push(tmp)
      }

      let submitTestPayload = {
        testId: route.params.testId,
        questionAnswer: questionAnswer
      }

      console.log(submitTestPayload)

      axios.post(process.env.VUE_APP_BACKEND_RESP_API + 'results', submitTestPayload, config).then(resp => {
        alert("Test submitted.")
      }).then(resp=>{
        router.push({name: "mainView"})
      }).catch(err => {
        console.log(err)
      })
    }

    function getUserTestResults(){
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      const params = new URLSearchParams({
        testId: route.params.testId
      }).toString()

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'results?' + params, config).then(resp => {
        isResolved.value = ((resp.data.length >= 1))
      }).catch(err => {
        console.log(err)
      })
    }

    return {
      test,
      selectedAnswer,
      answers,
      submitTest,
      isResolved
    }
  }
};


</script>

<style lang="scss"></style>