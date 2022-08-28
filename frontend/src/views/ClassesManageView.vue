<template>
  <LoggedNavbar/>
  <div class="row">
    <div class="column">
      <h2>Classes list</h2>
      <DataTable :value="classesFormatted" :paginator="true" :rows="10"
                 paginatorTemplate="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown"
                 :rowsPerPageOptions="[10,20,50]" responsiveLayout="scroll"
                 currentPageReportTemplate="Showing {first} to {last} of {totalRecords}">
        <Column field="name" header="Class Id">
          <template #body="slotProps">
            <a :href="'manageClass/' + slotProps.data.name" v-text="slotProps.data.name" />
          </template>
        </Column>
        <Column field="name" header="Delete class">
          <template #body="slotProps">
            <i @click="deleteClass" class="pi pi-trash"></i>
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
    <div class="columnslim">
      <Divider layout="vertical" />
    </div>

    <div class="column">
      <h2>Add new class</h2>
      <InputText type="text" v-model="newClass" />
      <Button @click="addClass" label="Add new class" />
    </div>


  </div>
</template>



<script>

import DataTable  from 'primevue/datatable';
import Column from 'primevue/column'
import Divider from 'primevue/divider'

import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";

export default {
  name: "ClassesManage",
  components: {
    LoggedNavbar,
    DataTable,
    Column,
    Divider
  },
  setup() {
    let classes = ref(null)
    let classesFormatted = ref([])
    let newClass = ref(null)

    onMounted(() => {
      getLessons()
    })

    function getLessons() {
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

    function addClass() {
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      let m = {
        UserClass: newClass.value,
      }
      axios.post(process.env.VUE_APP_BACKEND_RESP_API + 'classes', m, config).then(resp => {
        newClass.value = ""
      }).catch(err => {
        console.log(err)
      })
    }

    function deleteClass() {
      console.log('delete')
    }

      return {
      classes,
      classesFormatted,
      addClass,
      newClass,
      deleteClass
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


</style>