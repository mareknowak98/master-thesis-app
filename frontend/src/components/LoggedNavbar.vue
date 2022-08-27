<template>
<!--  <div v-if="this.decodedToken != null">-->
    <TabMenu class="myMenu" :model="items" />
<!--  </div>-->
<!--  <div v-else>-->
<!--    <h1>Log in to use website.</h1>-->
<!--  </div>-->
</template>



<script>

import { ref, onMounted } from "vue";
import { TokenService } from "@/store/tokenService";
import TabMenu from 'primevue/tabmenu';
import router from "@/router";

export default {
  name: "LoggedNavbar",
  components: {
    TabMenu,
  },
  data() {
    return {
      items: [
        {label: 'Menu', icon: 'pi pi-fw pi-home', command:() => {this.menu()}},
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

    function menu() {
      router.push({name: "mainView"})
    }

    function logOut(){
      TokenService.deleteToken()
      router.push({name: "home"})
    }

    return {
      decodedToken,
      group,
      logOut,
      menu
    }
  }
};


</script>

<style lang="scss">

.myMenu {

  .p-tabmenu-nav {
    padding-left: 40%;

  }
}


</style>