<template>
  <div id="textinput">
    <div class="row">
      <div class="column">
      <h2>Enter question:</h2>
      <InputText type="text" v-model="slideText" />

      <h2>Enter possible answers:</h2>
      <InputText type="text" v-model="answerText" />
      <Button @click="addAnswer" label="Add answer" />
      <p> </p>
      <Button @click="addQASlide" label="Add QA slide" />
      </div>

      <div class="column">
        <DataTable :value="answerList" v-model:selection="answerCorrect" selectionMode="multiple" dataKey="id" responsiveLayout="scroll" >
          <Column field="text" header="Select correct answers(multiple select with Alt)"></Column>
        </DataTable>
      </div>
    </div>

    </div>
</template>



<script>


import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import { useRoute } from "vue-router";
import Editor from 'primevue/editor';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column'
import Divider from 'primevue/divider'

import router from "@/router";

export default {
  name: "AddQASlide",
  components: {
    Editor,
    DataTable,
    Column,
    Divider
  },
  setup() {
    let presentationId = ref('')
    let slideText = ref('')
    let answerText = ref('')
    let answerList = ref([])
    let answerCorrect= ref([])

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


    function addQASlide() {

      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }

      let maxSlide = Math.max.apply(Math, currentPresentationSlides.value.map(function(o) { return parseInt(o.slideId); }))
      maxSlide++
      console.log(maxSlide)

      let qa = {
        answers: answerList.value,
        correctAnswers: answerCorrect.value
      }
      console.log(qa)
      console.log("ss")

      let m = {
        lessonId: presentationId.value,
        slideId: maxSlide.toString(),
        slideType: "QA",
        slideContent: slideText.value,
        questionAnswers: JSON.stringify(qa)
      }

      console.log(m)


      axios.post(process.env.VUE_APP_BACKEND_RESP_API + 'slides', m, config).then(resp => {
        router.push({name: "managePresentation"})
      }).catch(err => {
        console.log(err)
      })
    }

    function addAnswer() {
      let obj = {
        id: Math.floor(Math.random() * 100000),
        text: answerText.value
      }
      answerList.value.push(obj)
      answerText.value = ''
    }

      return {
      presentationId,
      slideText,
      addQASlide,
      currentPresentationSlides,
      addAnswer,
      answerText,
      answerList,
      answerCorrect
    }
  }
};


</script>



<style lang="scss">
.column {
  float: left;
  width: 45%;
}

.columnslim {
  float: left;
  width: 10%;
}

/* Clear floats after the columns */
.row:after {
  content: "";
  display: table;
  clear: both;
}

</style>
