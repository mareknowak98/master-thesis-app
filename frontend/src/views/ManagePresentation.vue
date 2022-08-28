<template>
  <LoggedNavbar/>
  <h2>Edit lesson: {{presentationId}}</h2>
  <div class="row">
    <div class="column">
      <h2>Edit existing presentation</h2>
      <DataTable :value="lessonsSlides" :paginator="true" :rows="10"
                 paginatorTemplate="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown"
                 :rowsPerPageOptions="[10,20,50]" responsiveLayout="scroll"
                 currentPageReportTemplate="Showing {first} to {last} of {totalRecords}">
        <Column field="slideId" header="Slide Id">
          <template #body="slotProps">
            <a :href="'slide/' + this.$route.params.presentationId + '/' + slotProps.data.slideId" v-text="slotProps.data.slideId" />
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
      <h2>Add new slide</h2>
      <router-link :to="'addNewSlide/' + presentationId" >
      <Button label="Add new slide" />
      </router-link>
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
import { useRoute } from "vue-router";

export default {
  name: "ManagePresentation",
  components: {
    LoggedNavbar,
    DataTable,
    Column,
  },
  setup() {
    let lessonsSlides = ref(null)
    let presentationId = ref('')
    const route = useRoute()

    onMounted(() => {
      presentationId.value = route.params.presentationId

      getSlides()
    })

    function getSlides() {
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      const params = new URLSearchParams({
        lesson: presentationId.value
      }).toString()

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'slides?' +params , config).then(resp => {
        lessonsSlides.value = resp.data
      }).catch(err => {
        console.log(err)
      })
    }

    return {
      lessonsSlides,
      presentationId
    }
  }
};


</script>

