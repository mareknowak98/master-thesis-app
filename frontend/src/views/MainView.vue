<template>
  <TabMenu class="myMenu" :model="items" />

  <TeacherMenu/>
</template>



<script>

import TeacherMenu from "@/components/TeacherMenu";
import { ref, onMounted } from "vue";
import { TokenService } from "@/store/tokenService";
import TabMenu from 'primevue/tabmenu';

export default {
  name: "MainView",
  components: {
    TeacherMenu,
    TabMenu
  },
  data() {
    return {
      items: [
        {label: 'Profile', icon: 'pi pi-fw pi-home'},
        {label: 'Settings', icon: 'pi pi-fw pi-cog'},
        {label: 'Log out', icon: 'pi pi-fw pi-calendar'},
      ]
    }
  },

  setup() {
    let decodedToken = ref(null)

    onMounted(() => {
      decodedToken.value = TokenService.decodeToken(TokenService.getToken())
    })

    return {
      decodedToken,
    }
  }
};


</script>

<style lang="scss">

.myMenu {

  .p-tabmenu-nav {
    position: relative;
    left: 50%;
    transform: translateX(-25%);

  }
}


</style>