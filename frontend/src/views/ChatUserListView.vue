<template>
  <LoggedNavbar/>
  <h2>All users list:</h2>
  <div>
    <DataTable :value="users" :paginator="true" :rows="10"
               paginatorTemplate="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown"
               :rowsPerPageOptions="[10,20,50]" responsiveLayout="scroll"
               currentPageReportTemplate="Showing {first} to {last} of {totalRecords}">
      <Column field="Email" header="Email">      </Column>
      <Column field="Username" header="Username">
        <template #body="slotProps">
          <a :href="'chatUser/' + slotProps.data.Username" v-text="slotProps.data.Username" />
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
  name: "ChatUserListView",
  components: {
    LoggedNavbar,
    DataTable,
    Column
  },
  setup() {
    let users = ref(null)

    onMounted(() => {
      getUsers()
    })

    function getUsers(){
      const params = new URLSearchParams({
        group: 'all'
      }).toString()

      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'users?' + params, config).then(resp => {
        users.value = resp.data
        console.log(resp)
      }).catch(err => {
        console.log(err)
      })
    }

    return {
      users,
      getUsers
    }
  }
};


</script>

<style lang="scss"></style>