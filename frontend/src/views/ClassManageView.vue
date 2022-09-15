<template>
  <LoggedNavbar/>
  <div class="row">
    <div class="column">
      <h2>Class student list</h2>
      <DataTable :value="classStudentsFormatted" :paginator="true" :rows="10"
                 paginatorTemplate="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown"
                 :rowsPerPageOptions="[10,20,50]" responsiveLayout="scroll"
                 currentPageReportTemplate="Showing {first} to {last} of {totalRecords}">
        <Column field="username" header="Class Id">
          <template #body="slotProps">
            <a v-text="slotProps.data.username" />
          </template>
        </Column>
        <Column field="name" header="Delete from class">
          <template #body="slotProps">
            <i @click="removeUserFromClass(slotProps.data.username)" class="pi pi-trash"></i>
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
      <h2>All student list</h2>
      <DataTable :value="allStudents" :paginator="true" :rows="10"
                 paginatorTemplate="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown"
                 :rowsPerPageOptions="[10,20,50]" responsiveLayout="scroll"
                 currentPageReportTemplate="Showing {first} to {last} of {totalRecords}">
        <Column field="name" header="Class Id">
          <template #body="slotProps">
            <a v-text="slotProps.data.Username" />
          </template>
        </Column>
        <Column field="name" header="Add to class">
          <template #body="slotProps">
            <i @click="addUserToClass(slotProps.data.Username)" class="pi pi-plus"></i>
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
import { useRoute } from "vue-router/dist/vue-router";

export default {
  name: "ClassManageView",
  components: {
    LoggedNavbar,
    DataTable,
    Column,
    Divider
  },
  setup() {
    let classStudents = ref(null)
    let classStudentsFormatted = ref([])
    let allStudents = ref(null)
    let className = ref('')
    const route = useRoute()

    onMounted(() => {
      getAllStudents()
      className.value = route.params.classId
      getClassStudents()
    })

    function getAllStudents() {
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      const params = new URLSearchParams({
        group: 'student-group'
      }).toString()

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'users?' + params, config).then(resp => {
        allStudents.value = resp.data
      }).catch(err => {
        console.log(err)
      })
    }

    function getClassStudents() {
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      const params = new URLSearchParams({
        class: className.value
      }).toString()

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'classUsers?' + params, config).then(resp => {
        classStudents.value = resp.data.userList
        classStudentsFormatted.value = []
      }).then( resp => {
        classStudents.value.forEach(username => {
          classStudentsFormatted.value.push({
            username
          })
        })}).catch(err => {
        console.log(err)
      })
    }


    function addUserToClass(username) {
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      let m = {
        UserClass: className.value,
        Username: username
      }
      axios.post(process.env.VUE_APP_BACKEND_RESP_API + 'classUsers', m, config).then(resp => {
        getClassStudents()
      }).catch(err => {
        console.log(err)
      })
    }

    function removeUserFromClass(username) {
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        },
        data: {
          UserClass: className.value,
          Username: username,
        }
      }

      console.log(config)

      axios.delete(process.env.VUE_APP_BACKEND_RESP_API + 'classUsers', config).then(resp => {
        getClassStudents()
      }).catch(err => {
        console.log(err)
      })
    }

      return {
      classStudents,
      allStudents,
      className,
      classStudentsFormatted,
      addUserToClass,
      removeUserFromClass
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