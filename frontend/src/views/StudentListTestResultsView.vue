<template>
  <LoggedNavbar/>
  <h2>Tests list</h2>
  <DataTable :value="tests" :paginator="true" :rows="10"
             paginatorTemplate="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown"
             :rowsPerPageOptions="[10,20,50]" responsiveLayout="scroll"
             currentPageReportTemplate="Showing {first} to {last} of {totalRecords}">
    <Column field="testId" header="Test Id">
      <template #body="slotProps">
        <a :href="'testResultList/' + slotProps.data.testId" v-text="slotProps.data.testId" />
      </template>
    </Column>
    <template #paginatorstart>
      <Button type="button" icon="pi pi-refresh" class="p-button-text" />
    </template>
    <template #paginatorend>
      <Button type="button" icon="pi pi-cloud" class="p-button-text" />
    </template>
  </DataTable>

</template>



<script>

import DataTable  from 'primevue/datatable';
import Column from 'primevue/column'
import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import { useRoute } from "vue-router/dist/vue-router";

export default {
  name: "StudentListTestResultView",
  components: {
    LoggedNavbar,
    DataTable,
    Column
  },
  setup() {
    let tests = ref(null)
    const route = useRoute()
    let myClass = ref(null)

    onMounted(() => {
      getClass()
      getTests()
    })

    function getClass(){
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      const params = new URLSearchParams({
        info: "class"
      }).toString()

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'me?' + params, config).then(resp => {
        myClass.value = resp.data
      }).catch(err => {
        console.log(err)
      })
    }

    function getTests(){
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'tests', config).then(resp => {
        tests.value = resp.data
      }).then(resp =>{
        tests.value = tests.value.filter(function (el)
        {
          return el.combinedKey.startsWith(myClass.value)
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

    return {
      tests,
      myClass
    }
  }
};


</script>

<style lang="scss"></style>