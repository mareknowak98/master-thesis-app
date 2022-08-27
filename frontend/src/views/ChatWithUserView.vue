<template>
  <LoggedNavbar/>
  <h1>Chatting with user {{this.$route.params.userId}}</h1>
  <h2>-------------------</h2>

  <table class="table table-striped table-bordered">
    <thead>
    <tr>
      <th>UserFrom:UserTo</th>
      <th>Date</th>
      <th>Message</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="msg in rest_messages" :key="msg.id">
      <td>{{msg.userFromTo}}</td>
      <td>{{msg.timestamp}}</td>
      <td>{{msg.message}}</td>
    </tr>
    </tbody>
  </table>
  <h2>-------------------</h2>

  <InputText type="text" v-model="msg" />
  <Button @click="sendMessage" label="Send message" />

</template>



<script>

import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import { useRoute } from "vue-router";

export default {
  name: "ChatWithUserView",
  components: {
    LoggedNavbar,
  },
  setup() {
    let rest_messages = ref(null)
    let userid = ref(null)
    let msg = ref('')

    let websocket_msg = ref('')
    let socket = ref(null)
    let token = ref(null)


    const route = useRoute()

    onMounted(() => {
      userid.value = route.params.userId
      token.value = TokenService.getToken()
      getRestMessages()

      try {
        const params = new URLSearchParams({
          token: token.value
        }).toString()
        socket.value = new WebSocket(process.env.VUE_APP_WEBSOCKET_CHAT_API + '?' + params)
        socket.value.onmessage = (m) => {
          websocket_msg.value = JSON.parse(m.data)
          if (websocket_msg.value.payload.from === userid.value){
            rest_messages.value.push({
              userFromTo: websocket_msg.value.payload.from + ":" + websocket_msg.value.payload.to,
              timestamp: "",
              message: websocket_msg.value.payload.message
            })
          }
        }

      } catch (e) {
        console.log(e)
      }
    })

    function sendMessage() {
      var d = new Date();
      let m = {
          type: "message",
          payload: {
            from: TokenService.decodeToken(token.value).username,
            to: userid.value,
            message: msg.value
          }
      }
      socket.value.send(JSON.stringify(m))
      rest_messages.value.push({
        userFromTo: TokenService.decodeToken(token.value).username + ":" + userid.value,
        timestamp: d.toLocaleString(),
        message: msg.value
      })
    }


    function getRestMessages(){
      let params2 = new URLSearchParams({
        UserTo: userid.value
      }).toString()

      let config = {
        headers: {
          Authorization: token.value,
        }
      }

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'messages?' + params2, config).then(resp => {
        rest_messages.value = resp.data
      }).catch(err => {
        console.log(err)
      })
    }

    return {
      rest_messages,
      getRestMessages,
      websocket_msg,
      socket,
      msg,
      sendMessage
    }
  }
};


</script>

<style lang="scss"></style>