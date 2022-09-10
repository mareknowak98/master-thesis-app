<template>
  <LoggedNavbar/>
  <h2>Tests results</h2>
  <Dropdown v-model="selectedClass" :options="classesFormatted" optionLabel="name" optionValue="name" @change="loadTests"  placeholder="Select class" />

  <div v-if="this.selectedClass != null">
    <h2>Select test</h2>
    <Dropdown v-model="selectedTest" :options="tests" optionLabel="testId" optionValue="testId" @change="getTestResult"  placeholder="Select class" />
  </div>

  <div v-if="this.results != null">
    <DataTable :value="results" responsiveLayout="scroll">
      <Column field="userId" header="User"></Column>
      <Column field="result" header="Result"></Column>
    </DataTable>
  </div>

</template>



<script>

import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted, watch } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import Dropdown from 'primevue/dropdown';
import DataTable from "primevue/datatable";
import Column from "primevue/column";

export default {
  name: "TeacherTestResults",
  components: {
    LoggedNavbar,
    Dropdown,
    DataTable,
    Column
  },
  setup() {
    let results = ref(null)
    let classes = ref([])
    let classesFormatted = ref([])
    let selectedClass = ref(null)
    let tests = ref(null)
    let selectedTest = ref(null)

    onMounted(() => {
      getClasses()
    })

    function loadTests(){
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'tests', config).then(resp => {
        tests.value = resp.data
        console.log(resp.data)

      }).then(resp =>{
        let tmp = selectedClass.value
        console.log(tmp)
        tests.value = tests.value.filter(function (el)
        {
          return el.combinedKey.startsWith(tmp + ':')
        })

        tests.value = tests.value.filter((value, index, self) =>
            index === self.findIndex((t) => (
              t.testId === value.testId
            ))
        )
      }).catch(err => {
        console.log(err)
      })
    }

    function getClasses(){
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }

      const params = new URLSearchParams({
        folder: 'student-group'
      }).toString()

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'classes', config).then(resp => {
        classes.value = resp.data
      }).then( resp => {
        classes.value.forEach(name => {
          classesFormatted.value.push({
            name
          })
        })}).catch(err => {
        console.log(err)
      })
    }


    function getTestResult(){
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      const params = new URLSearchParams({
        testId: selectedTest.value
      }).toString()

      console.log(params)

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'results?' + params, config).then(resp => {
        results.value = resp.data
      }).then(resp =>{
        for (let result of results.value) {
          result.result = (parseFloat(result.result)*100).toString() + "%"
        }
      }).catch(err => {
        console.log(err)
      })
    }

    return {
      results,
      classes,
      classesFormatted,
      selectedClass,
      loadTests,
      tests,
      selectedTest,
      getTestResult
    }
  }
};


</script>

<style lang="scss"></style>