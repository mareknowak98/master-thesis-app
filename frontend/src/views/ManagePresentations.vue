<template>
  <LoggedNavbar/>
  <div class="row">
  <div class="column">
  <h2>Edit existing presentation</h2>
  <DataTable :value="lessons" :paginator="true" :rows="10"
             paginatorTemplate="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown"
             :rowsPerPageOptions="[10,20,50]" responsiveLayout="scroll"
             currentPageReportTemplate="Showing {first} to {last} of {totalRecords}">
    <Column field="lessonId" header="Lesson Id">
    <template #body="slotProps">
      <a :href="'managePresentation/' + slotProps.data.lessonId" v-text="slotProps.data.lessonId" />
    </template>
    </Column>
    <template #paginatorstart>
      <Button type="button" icon="pi pi-refresh" class="p-button-text" />
    </template>
    <template #paginatorend>
      <Button type="button" icon="pi pi-cloud" class="p-button-text" />
    </template>
  </DataTable>
  </div>
    <div class="columnslim">
      <Divider layout="vertical" />
    </div>

  <div class="column">
    <h2>Add new lesson</h2>
    <InputText type="text" v-model="newLesson" />
    <Button @click="addLesson" label="Add new lesson" />
  </div>


  </div>
</template>



<script>

import DataTable  from 'primevue/datatable';
import Column from 'primevue/column'
import Divider from 'primevue/divider'

import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";

export default {
  name: "ManagePresentations",
  components: {
    LoggedNavbar,
    DataTable,
    Column,
    Divider
  },
  setup() {
    let lessons = ref(null)
    let newLesson = ref('')

    onMounted(() => {
      getLessons()
    })

    function getLessons() {
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'lessons', config).then(resp => {
        lessons.value = resp.data
        console.log(resp)
      }).catch(err => {
        console.log(err)
      })
    }

    function addLesson() {
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      let m = {
        lessonId: newLesson.value,
        slideId: "1",
        slideType: "",
        slideContent: "",
        questionAnswers: ""
      }
      axios.post(process.env.VUE_APP_BACKEND_RESP_API + 'slides', m, config).then(resp => {
        console.log(resp)
        newLesson.value = ""
      }).catch(err => {
        console.log(err)
      })
    }

    return {
      lessons,
      newLesson,
      addLesson,
    }
  }
};


</script>

<style lang="scss">
.column {
  float: left;
  width: 45%;
}

.columnslim {
  float: left;
  width: 10%;
}

/* Clear floats after the columns */
.row:after {
  content: "";
  display: table;
  clear: both;
}

</style>