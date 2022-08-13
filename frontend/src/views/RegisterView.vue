<template>
  <div class="home">
    <div class="grid">
      <div class="col-6 col-offset-3">
        <div class="surface-card p-4 shadow-2 border-round w-full">

          <NavBar/>
          <div class="login">
            <div class="surface-card border-round w-full">
              <div class="text-center mb-3"></div>
              <div class="text-center">
                <img src="../assets/mylearn_logo.png" alt="Image" height="150" class="mb-3">
              </div>
              <div>
                <label for="email" class="block text-900 font-medium">Email</label>
                <InputText id="email" v-model="email" type="email" class="w-full mb-3" />

                <label for="username" class="block text-900 font-medium">Username</label>
                <InputText id="username" v-model="username" type="text" class="w-full mb-3" />

                <label for="password1" class="block text-900 font-medium mb-2">Password</label>
                <InputText id="password1" v-model="password1" type="password" class="w-full mb-3" />

                <label for="password2" class="block text-900 font-medium mb-2">Repeat password</label>
                <InputText id="password2" v-model="password2" type="password" class="w-full mb-3" />

                <Button label="Register" icon="pi pi-user-plus" @click="onRegister" class="w-full"></Button>
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
import { ref } from "vue";
import axios from "axios";

export default {
  name: "Register",
  components: {
    Checkbox,
    NavBar,
  },
  setup() {
    let email = ref('')
    let username = ref('')
    let password1 = ref('')
    let password2 = ref('')

    function onRegister() {
      if (password1.value !== password2.value){
        alert("Passwords must be equal")
        return
      }

      axios.post(process.env.VUE_APP_BACKEND_RESP_API + 'register', {
        email: email.value,
        username: username.value,
        password1: password1.value,
        password2: password2.value,
      }).then(resp => {
        console.log(resp)
        alert("User successfully registered.")
      }).catch(err =>{
        console.log(err)
        alert("Error" + err)
      })
    }

    return {
      email,
      username,
      password1,
      password2,
      onRegister
    }

  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="less"></style>
