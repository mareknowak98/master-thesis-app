<template>
  <div class="navbar">
    <h1>Env variable: |{{VUE_APP_BACKEND_RESP_API}}|</h1>
    <h1>custom variable: |{{VUE_APP_BACKEND_ECS}}|</h1>

    <TabMenu :model="items" />
    <router-view/>
  </div>
</template>

<script>
import TabMenu from 'primevue/tabmenu';
import { ref } from 'vue';

export default {
  name: "Navbar",
  components: {
    TabMenu,
  },
  setup() {
    const items = ref([
      {
        label: 'Home',
        icon: 'pi pi-fw pi-home',
        to: '/'
      },
      {
        label: 'Login',
        icon: 'pi pi-fw pi-lock',
        to: '/login'
      },
      {
        label: 'Register',
        icon: 'pi pi-fw pi-user-plus',
        to: '/register'
      }
    ]);
    return { items }
  },

  data() {
    return {
      VUE_APP_BACKEND_RESP_API: process.env.VUE_APP_BACKEND_RESP_API,
      VUE_APP_BACKEND_ECS: process.env.VUE_APP_BACKEND_ECS
    }
  },
};
</script>
