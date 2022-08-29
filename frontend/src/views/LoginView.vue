<template>
  <div class="home">
    <div class="grid">
      <div class="col-6 col-offset-3">
        <div class="surface-card p-4 shadow-2 border-round w-full">

        <NavBar/>
        <div class="login">
          <div class="surface-card border-round w-full">
            <div class="text-center mb-3"></div>
            <div class="text-center mb-2">
              <div class="text-center">
                <img src="../assets/mylearn_logo.png" alt="Image" height="150" class="mb-3">
              </div>
              <div class="text-900 text-3xl font-medium mb-3">Welcome Back</div>
              <span class="text-600 font-medium line-height-3">Don't have an account?</span>
              <a class="font-medium no-underline ml-2 text-blue-500 cursor-pointer">Create today!</a>
            </div>

            <div>
              <label for="email1" class="block text-900 font-medium mb-2">Username</label>
              <InputText id="email1" v-model="email1" type="text" class="w-full mb-3" />

              <label for="password1" class="block text-900 font-medium mb-2">Password</label>
              <InputText id="password1" v-model="password1" type="password" class="w-full mb-3" />

              <div class="flex align-items-center justify-content-between mb-6">
                <a class="font-medium no-underline ml-2 text-blue-500 text-right cursor-pointer">Forgot password?</a>
              </div>

              <Button label="Sign In" icon="pi pi-lock" class="w-full" @click="onLogin" ></Button>
            </div>
          </div>
        </div>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
import Checkbox from 'primevue/checkbox';
import NavBar from "@/components/NavBar";
import {ref} from 'vue'
import axios from 'axios';
import { TokenService } from "@/store/tokenService";
import router from "@/router";

export default {
  name: "Login",
  components: {
    Checkbox,
    NavBar,
  },
  setup() {
    let email1 = ref('')
    let password1 = ref('')

    function onLogin() {
      axios.post(process.env.VUE_APP_BACKEND_RESP_API + 'login', {
        username: email1.value,
        password: password1.value,
      }).then(resp => {
        TokenService.setToken(resp.data.accessToken)
        TokenService.setIdToken(resp.data.idToken)
        email1.value = password1.value = ''
      }).then(resp => {
        let decodedToken = TokenService.decodeToken(TokenService.getToken())
        console.log(decodedToken)
        router.push({name: "mainView"})
      }).catch(err => {
        alert(err)
        email1.value = password1.value = ''
      })
    }

    return {
      email1,
      password1,
      onLogin
    }
  }
};


</script>

<style scoped lang="less"></style>
