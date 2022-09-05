<template>
  <LoggedNavbar/>
  <div class="row">
    <div v-if="this.currentLessons != null">

    <h2>Live lessons: </h2>
    <DataTable :value="currentLessons" :paginator="true" :rows="10"
                 paginatorTemplate="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown"
                 :rowsPerPageOptions="[10,20,50]" responsiveLayout="scroll"
                 currentPageReportTemplate="Showing {first} to {last} of {totalRecords}">
        <Column field="lessonId" header="Lesson Id">
          <template #body="slotProps">
            <a :href="'lesson/' + slotProps.data.lessonId" v-text="slotProps.data.lessonId" />
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

    <div v-else>
      <h2>There is no live lessons now</h2>
    </div>

  </div>
</template>



<script>

import DataTable  from 'primevue/datatable';
import Column from 'primevue/column'

import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";

export default {
  name: "CurrentLessonsView",
  components: {
    LoggedNavbar,
    DataTable,
    Column
  },
  setup() {
    let currentLessons = ref(null)

    onMounted(() => {
      getCurrentLessons()
    })

    function getCurrentLessons() {
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }

      const params = new URLSearchParams({
        ongoing: "true"
      }).toString()

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'lessons?' + params, config).then(resp => {
        currentLessons.value = resp.data
      }).catch(err => {
        console.log(err)
      })
    }


    return {
      currentLessons,
    }
  }
};


</script>

<style lang="scss">


</style>