<template>
  <LoggedNavbar/>
  <div>
    <div v-if="this.group == 'teacher-group'">
      <TeacherMenu/>
    </div>
    <div v-if="this.group == 'student-group'">
      <StudentMenu/>
    </div>
    <div v-if="this.group == 'admin-group'">
      <AdminView/>
    </div>
  </div>
</template>



<script>

import TeacherMenu from "@/components/TeacherMenu";
import StudentMenu from "@/components/StudentMenu";
import LoggedNavbar from "@/components/LoggedNavbar";
import AdminView from "@/components/AdminView";

import { ref, onMounted } from "vue";
import { TokenService } from "@/store/tokenService";
import TabMenu from 'primevue/tabmenu';
import router from "@/router";

export default {
  name: "MainView",
  components: {
    TeacherMenu,
    StudentMenu,
    LoggedNavbar,
    AdminView,
    TabMenu
  },
  data() {
    return {
      items: [
        {label: 'Profile', icon: 'pi pi-fw pi-home'},
        {label: 'Settings', icon: 'pi pi-fw pi-cog'},
        {label: 'Log out', icon: 'pi pi-fw pi-calendar', command:() => {this.logOut()}},
      ]
    }
  },

  setup() {
    let decodedToken = ref(null)
    let group = ref("")

    onMounted(() => {
      decodedToken.value = TokenService.decodeToken(TokenService.getToken())
      group.value = decodedToken.value['cognito:groups'][0]
    })

    function logOut(){
      TokenService.deleteToken()
      router.push({name: "home"})
    }

    return {
      decodedToken,
      group,
      logOut
    }
  }
};


</script>

<style lang="scss">

</style>