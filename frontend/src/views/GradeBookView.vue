<template>
  <LoggedNavbar/>
  <h2>Student list:</h2>
  <div>
    <DataTable :value="students" :paginator="true" :rows="10"
               paginatorTemplate="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown"
               :rowsPerPageOptions="[10,20,50]" responsiveLayout="scroll"
               currentPageReportTemplate="Showing {first} to {last} of {totalRecords}">
      <Column field="Email" header="Email"></Column>
      <Column field="Username" header="Username"></Column>
      <template #paginatorstart>
        <Button type="button" icon="pi pi-refresh" class="p-button-text" />
      </template>
      <template #paginatorend>
        <Button type="button" icon="pi pi-cloud" class="p-button-text" />
      </template>
    </DataTable>
  </div>

</template>



<script>

import DataTable  from 'primevue/datatable';
import Column from 'primevue/column'
import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";

export default {
  name: "GradeBookView",
  components: {
    LoggedNavbar,
    DataTable,
    Column
  },
  setup() {
    let students = ref(null)

    onMounted(() => {
      getStudents()
    })

    function getStudents(){
      const params = new URLSearchParams({
        group: 'student-group'
      }).toString()

      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'users?' + params, config).then(resp => {
        students.value = resp.data
        console.log(resp)
      }).catch(err => {
        console.log(err)
      })
    }

    return {
      students,
      getStudents
    }
  }
};


</script>

<style lang="scss"></style>