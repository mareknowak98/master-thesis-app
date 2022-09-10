<template>
  <div class="surface-section px-4 py-8 md:px-6 lg:px-8 text-center">
    <div class="text-center">
      <img src="../assets/mylearn_logo.png" alt="Image" height="150" class="mb-3">
    </div>
    <div class="grid">
      <div class="col-12 md:col-4 mb-4 px-5">
        <router-link to="/manageUsers">
            <span class="p-3 shadow-2 mb-3 inline-block surface-card" style="border-radius: 10px">
                <i class="pi pi-desktop text-4xl text-blue-500"></i>
            </span>
          <div class="text-900 mb-3 font-medium">Manage users</div>
          <span class="text-700 text-sm line-height-3">Change users groups</span>
        </router-link>
      </div>

    </div>
  </div>


</template>



<script>

import { ref, onMounted } from "vue";
import { TokenService } from "@/store/tokenService";
import axios from "axios";

export default {
  name: "AdminMenuView",
  components: {
  },
  setup() {
    let myUsername = ref(null)
    let decodedToken = ref(null)

    onMounted(() => {
      decodedToken.value = TokenService.decodeToken(TokenService.getToken())
      myUsername.value = decodedToken.value.username

    })

    return {
      decodedToken,
    }
  }
};


</script>

<style scoped lang="less"></style>