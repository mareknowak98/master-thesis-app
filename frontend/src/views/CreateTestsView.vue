<template>
  <LoggedNavbar/>
  <h2>Create test for class:</h2>
  <div class="column">
    <DataTable :value="classesFormatted" :paginator="true" :rows="10"
               paginatorTemplate="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown"
               :rowsPerPageOptions="[10,20,50]" responsiveLayout="scroll"
               currentPageReportTemplate="Showing {first} to {last} of {totalRecords}">
      <Column field="name" header="Class Id">
        <template #body="slotProps">
          <a :href="'createTest/' + slotProps.data.name" v-text="slotProps.data.name" />
        </template>
      </Column>
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
  name: "CreateTestsView",
  components: {
    LoggedNavbar,
    DataTable,
    Column
  },
  setup() {
    let classes = ref(null)
    let classesFormatted = ref([])

    onMounted(() => {
      getClasses()
    })

    function getClasses() {
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }

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

    return {
      classes,
      classesFormatted
    }
  }
};


</script>

<style lang="scss"></style>