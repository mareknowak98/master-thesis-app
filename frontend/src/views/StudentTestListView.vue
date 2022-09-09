<template>
  <LoggedNavbar/>
  <h2>Tests list</h2>


</template>



<script>

import DataTable  from 'primevue/datatable';
import Column from 'primevue/column'
import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";

export default {
  name: "StudentTestListView",
  components: {
    LoggedNavbar,
    DataTable,
    Column
  },
  setup() {
    let tests = ref(null)

    onMounted(() => {
      getTests()
    })

    function getTests(){

    }

    // function getClasses() {
    //   let config = {
    //     headers: {
    //       Authorization: TokenService.getToken(),
    //     }
    //   }
    //
    //   axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'classes', config).then(resp => {
    //     classes.value = resp.data
    //   }).then( resp => {
    //     classes.value.forEach(name => {
    //       classesFormatted.value.push({
    //         name
    //       })
    //     })}).catch(err => {
    //     console.log(err)
    //   })
    // }

    return {
      tests,
    }
  }
};


</script>

<style lang="scss"></style>