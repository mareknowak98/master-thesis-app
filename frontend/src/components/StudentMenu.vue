<template>
  <div class="surface-section px-4 py-8 md:px-6 lg:px-8 text-center">
    <div class="text-center">
      <img src="../assets/mylearn_logo.png" alt="Image" height="150" class="mb-3">
    </div>
    <div class="grid">
      <div class="col-12 md:col-4 mb-4 px-5">
            <span class="p-3 shadow-2 mb-3 inline-block surface-card" style="border-radius: 10px">
                <i class="pi pi-desktop text-4xl text-blue-500"></i>
            </span>
        <div class="text-900 mb-3 font-medium">Join lesson</div>
        <span class="text-700 text-sm line-height-3">Join virtual lesson</span>
      </div>
      <div class="col-12 md:col-4 mb-4 px-5">
        <router-link :to="'/gradeBook/'+ this.myUsername">
            <span class="p-3 shadow-2 mb-3 inline-block surface-card" style="border-radius: 10px">
                <i class="pi pi-lock text-4xl text-blue-500"></i>
            </span>
        <div class="text-900 mb-3 font-medium">Grades</div>
        <span class="text-700 text-sm line-height-3">See your grades</span>
        </router-link>
      </div>
      <div class="col-12 md:col-4 mb-4 px-5">
            <span class="p-3 shadow-2 mb-3 inline-block surface-card" style="border-radius: 10px">
                <i class="pi pi-check-circle text-4xl text-blue-500"></i>
            </span>
        <div class="text-900 mb-3 font-medium">XYZ</div>
        <span class="text-700 text-sm line-height-3">xyz</span>
      </div>
      <div class="col-12 md:col-4 mb-4 px-5">
        <router-link to="/chatList">
            <span class="p-3 shadow-2 mb-3 inline-block surface-card" style="border-radius: 10px">
                <i class="pi pi-globe text-4xl text-blue-500"></i>
            </span>
          <div class="text-900 mb-3 font-medium">User Chat</div>
          <span class="text-700 text-sm line-height-3">Exchange messages with other users</span>
        </router-link>
      </div>
      <div class="col-12 md:col-4 mb-4 px-5">
            <span class="p-3 shadow-2 mb-3 inline-block surface-card" style="border-radius: 10px">
                <i class="pi pi-github text-4xl text-blue-500"></i>
            </span>
        <div class="text-900 mb-3 font-medium">ZYX</div>
        <span class="text-700 text-sm line-height-3">XYZ</span>
      </div>
      <div class="col-12 md:col-4 md:mb-4 mb-0 px-3">
        <router-link :to="'/classContent/'+ this.myClass">
            <span class="p-3 shadow-2 mb-3 inline-block surface-card" style="border-radius: 10px">
                <i class="pi pi-shield text-4xl text-blue-500"></i>
            </span>
        <div class="text-900 mb-3 font-medium">Homeworks and learning materials</div>
        <span class="text-700 text-sm line-height-3"></span>
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
  name: "StudentMenu",
  components: {
  },
  setup() {
    let myUsername = ref(null)
    let myClass = ref(null)

    let decodedToken = ref(null)

    onMounted(() => {
      decodedToken.value = TokenService.decodeToken(TokenService.getToken())
      myUsername.value = decodedToken.value.username
      getClass()
    })

    function getClass(){
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      const params = new URLSearchParams({
        info: "class"
      }).toString()

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'me?' + params, config).then(resp => {
        myClass.value = resp.data
      }).catch(err => {
        console.log(err)
      })
    }

    return {
      myUsername,
      decodedToken,
      myClass
    }
  }
};


</script>

<style scoped lang="less"></style>