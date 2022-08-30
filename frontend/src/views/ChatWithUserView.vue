<template>
  <LoggedNavbar/>
  <h1>Chatting with user {{this.$route.params.userId}}</h1>


  <div class="panel">
    <div class="messages" ref="messagesRef">
      <div class="inner">
        <div
          :key="index"
          v-for="(message, index) in rest_messages"
          class="message"
        >
          <div v-if="message.userFromTo === msgSent" class="user-self">
            You:&nbsp;
          </div>
          <div v-if="message.userFromTo === msgRec" class="user-them">
            {{this.userid}}:&nbsp;
          </div>
          <div class="text">{{ message.message }}</div>
        </div>
      </div>
    </div>

  </div>

  <InputText type="text" v-model="msg" />
  <Button @click="sendMessage" label="Send message" />

</template>



<script>

import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted, onUpdated } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import { useRoute } from "vue-router";
import { watch } from "vue";
import DataTable from 'primevue/datatable';
import Column from 'primevue/column'

export default {
  name: "ChatWithUserView",
  components: {
    LoggedNavbar,
    DataTable,
    Column
  },
  setup() {
    let rest_messages = ref(null)
    let userid = ref(null)
    let msg = ref('')

    let websocket_msg = ref('')
    let socket = ref(null)
    let token = ref(null)

    let msgSent = ref('')
    let msgRec = ref('')

    const route = useRoute()

    onMounted(() => {
      userid.value = route.params.userId
      token.value = TokenService.getToken()
      getRestMessages()

      msgSent.value = TokenService.decodeToken(token.value).username + ":" + userid.value
      msgRec.value = userid.value  + ":" +  TokenService.decodeToken(token.value).username

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

    watch(rest_messages.value, (oldValue, prevSelection) => {
      console.log("essa")
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
      msg.value = ""
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
      }).then(resp =>{
        rest_messages.value.sort(function(a,b){
          // Turn your strings into dates, and then subtract them
          // to get a value that is either negative, positive, or zero.
          return new Date(a.timestamp) - new Date(b.timestamp);
        });
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
      sendMessage,
      msgSent,
      msgRec,
      userid
    }
  }
};


</script>

<style lang="scss">
.panel {
  display: flex;
  flex-direction: column;
  padding: 20px;
  margin: 0 auto;
  //max-width: 300px;
  height: 400px;
  background: rgba(255, 255, 255, 0.7);
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
  backdrop-filter: blur(4px);
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.18);
  margin-bottom: 2%;
}
.messages {
  height: 100%;
  width: 100%;
  overflow-y: scroll;
  border-top-left-radius: 5px;
  border-top-right-radius: 5px;
  background-color: white;
}
.inner {
  padding: 10px;
}
.message {
  text-align: left;
  display: flex;
  margin-bottom: 6px;
}
.user-self {
  color: green;
}
.user-them {
  color: red;
}


</style>