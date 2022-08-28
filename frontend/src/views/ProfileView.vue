<template>
  <LoggedNavbar/>
  <div class="space"></div>

  <Card>
    <template #header>
      <div class="img">
      <img src="../assets/mylearn_logo.png" alt="user header">
      </div>
    </template>
    <template #title>
      Profile
    </template>
    <template #content>
      <Divider type="dashed"></Divider>
      <h3>Username: {{this.username}}</h3>
      <h3>User group: {{this.userGroup}}</h3>
      <h3>User class: {{this.userClass}}</h3>

    </template>
<!--    TODO add profile edit-->
<!--    <template #footer>-->
<!--      <Button icon="pi pi-check" label="Save" />-->
<!--      <Button icon="pi pi-times" label="Cancel" class="p-button-secondary" style="margin-left: .5em" />-->
<!--    </template>-->
  </Card>
</template>



<script>

import "primevue/resources/themes/saga-blue/theme.css";
import "primevue/resources/primevue.min.css";
import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import Card from "primevue/card";

import Divider from "primevue/divider";

export default {
  name: "ProfileView",
  components: {
    LoggedNavbar,
    Card,
    Divider
  },
  setup() {
    let me = ref(null)
    let decodedToken = ref(null)
    let username = ref('')
    let userGroup = ref('')
    let userClass = ref('')

    onMounted(() => {
      decodeToken()
      username.value = decodedToken.value.username
      userGroup.value = decodedToken.value['cognito:groups'][0]
      getUserClass()
    })

    function decodeToken(){
      decodedToken.value = TokenService.decodeToken(TokenService.getToken())
    }

    function getUserClass(){
      let params = new URLSearchParams({
        info: "class"
      }).toString()

      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'me?' + params, config).then(resp => {
        userClass.value = resp.data
      }).catch(err => {
        console.log(err)
      })
    }

    return {
      me,
      decodedToken,
      username,
      userGroup,
      userClass
    }
  }
};


</script>

<style>
.img{
  max-width: 10%;
  height: auto;
  display: block;
  margin-left: auto;
  margin-right: auto;
}

.space{
  margin-bottom: 2%
}

.p-divider-solid.p-divider-horizontal:before {
  border-top-style: solid !important;
}
.p-divider-dashed.p-divider-horizontal:before {
  border-top-style: dashed !important;
}
.p-divider-solid.p-divider-vertical:before {
  border-left-style: solid !important;
}

</style>