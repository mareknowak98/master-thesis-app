<template>
  <div class="home">
    <div class="grid">
      <div class="col-6 col-offset-3">
        <div class="surface-card p-4 shadow-2 border-round w-full">

        <NavBar/>
        <div class="login">
          <div class="surface-card border-round w-full">
            <div class="text-center mb-5">
              <img src="../assets/mylearn_logo.png" alt="Image" height="150" class="mb-3">
              <div class="text-900 text-3xl font-medium mb-3">Welcome Back</div>
              <span class="text-600 font-medium line-height-3">Don't have an account?</span>
              <a class="font-medium no-underline ml-2 text-blue-500 cursor-pointer">Create today!</a>
            </div>

            <div>
              <label for="email1" class="block text-900 font-medium mb-2">Email</label>
              <InputText id="email1" type="text" class="w-full mb-3" />

              <label for="password1" class="block text-900 font-medium mb-2">Password</label>
              <InputText id="password1" type="password" class="w-full mb-3" />

              <div class="flex align-items-center justify-content-between mb-6">
                <a class="font-medium no-underline ml-2 text-blue-500 text-right cursor-pointer">Forgot password?</a>
              </div>

              <Button label="Sign In" icon="pi pi-user" class="w-full"></Button>
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

export default {
  name: "Login",
  components: {
    Checkbox,
    NavBar,
  },
  props: {
    msg: String,
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="less"></style>
