<template>
  <LoggedNavbar/>
  <h2>Tests results in %</h2>
  <Knob v-model="result" readonly="true" :size="200" />

</template>



<script>

import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import { useRoute } from "vue-router/dist/vue-router";
import Knob from 'primevue/knob';

export default {
  name: "StudentTestResultView",
  components: {
    LoggedNavbar,
    Knob

  },
  setup() {
    let result = ref(null)
    const route = useRoute()

    onMounted(() => {
      getTestResult()
    })

    function getTestResult(){
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      const params = new URLSearchParams({
        testId: route.params.testId
      }).toString()

      console.log(params)

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'results?' + params, config).then(resp => {
        result.value = resp.data
      }).then(resp=>{
        result.value = parseFloat(result.value[0].result)*100
      }).catch(err => {
        console.log(err)
      })
    }

    return {
      result
    }
  }
};


</script>

<style lang="scss"></style>