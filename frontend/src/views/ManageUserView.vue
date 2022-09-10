<template>
  <LoggedNavbar/>
  <h1>Manage user {{this.$route.params.userId}}</h1>

  <h5></h5>
  <Dropdown v-model="selectedGroup" :options="groups" optionLabel="name" optionValue="code" placeholder="Select group" />
  <Button @click="changeUserGroup" label="Change group" />

</template>



<script>

import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted, onUpdated } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import { useRoute } from "vue-router";
import DataTable from 'primevue/datatable';
import Column from 'primevue/column'
import Dropdown from "primevue/dropdown";

export default {
  name: "ManageUserView",
  components: {
    LoggedNavbar,
    DataTable,
    Column,
    Dropdown
  },
  setup() {
    const groups = ref([
      {name: 'Student group', code: 'student-group'},
      {name: 'Teacher group', code: 'teacher-group'},
      {name: 'Parent group', code: 'parent-group'},
      {name: 'Admin group', code: 'admin-group'},
    ]);
    let selectedGroup = ref('')
    const route = useRoute()

    function changeUserGroup(){
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }

      let m = {
        operation: "addToGroup",
        username: route.params.userId,
        groupName: selectedGroup.value,
      }

      axios.post(process.env.VUE_APP_BACKEND_RESP_API + 'manageGroups', m, config).then(resp => {
        alert("User group has been changed")
        selectedGroup.value = ''
      }).catch(err => {
        alert("Error occurred changing user group.")
      })

    }

    return {
      changeUserGroup,
      groups,
      selectedGroup
    }
  }
};


</script>

<style lang="scss">



</style>