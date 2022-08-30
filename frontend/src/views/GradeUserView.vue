<template>
  <LoggedNavbar/>
  <h2>Student {{this.$route.params.userId}} grades list:</h2>
  <div class="card">
    <DataTable :value="grades" rowGroupMode="subheader" groupRowsBy="Subject"
               sortMode="single" sortField="Subject" :sortOrder="1" scrollable scrollHeight="400px">
      <Column field="Subject" ></Column>
<!--      <Column field="UserId" header="UserId" style="min-width:200px"></Column>-->
      <Column field="Date" header="Date" style="min-width:200px">
        <template #body="slotProps">
          <span class="image-text">{{slotProps.data.Date}}</span>
        </template>
      </Column>
      <Column field="Grade" header="Grade" style="min-width:200px"></Column>

      <template #groupheader="slotProps">
        <img width="32" style="vertical-align: middle" />
        <span class="image-text">{{slotProps.data.Subject}}</span>
      </template>
      <template #groupfooter="slotProps">
        <td style="min-width: 80%">
          <div style="text-align: right; width: 100%">Mean</div>
        </td>
        <td style="width: 20%">{{calculate(slotProps.data.Subject)}}</td>
      </template>
    </DataTable>
  </div>
  <h3>Add grade</h3>
  <h5>Subject</h5>
  <Dropdown v-model="selectedSubject" :options="subjects" optionLabel="name" optionValue="code" placeholder="Select a Subject" />
  <h5>Grade</h5>
  <Dropdown v-model="selectedGrade" :options="gradesList" optionLabel="name" optionValue="code" placeholder="Select a Subject" />
  <Button @click="addGrade" label="Add new grade" />

</template>



<script>

import DataTable  from 'primevue/datatable';
import Column from 'primevue/column'
import Dropdown from "primevue/dropdown";
import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import { useRoute } from "vue-router";

export default {
  name: "GradeUserView",
  components: {
    LoggedNavbar,
    DataTable,
    Column,
    Dropdown
  },
  setup() {
    let grades = ref(null)
    let parsedGrades = ref(null)

    let selectedSubject = ref(null)
    const subjects = ref([
      {name: 'Math', code: 'Math'},
      {name: 'Art', code: 'Art'},
      {name: 'English', code: 'English'},
      {name: 'Music', code: 'Music'},
      {name: 'History', code: 'History'},
      {name: 'Science', code: 'Science'},
      {name: 'Geography', code: 'Geography'},
      {name: 'Information technology', code: 'Information technology'},
      {name: 'Biology', code: 'Biology'}
    ]);

    let selectedGrade = ref(null)
    const gradesList = ref([
      {name: '1', code: '1'},
      {name: '2', code: '2'},
      {name: '3', code: '3'},
      {name: '4', code: '4'},
      {name: '5', code: '5'},
      {name: '6', code: '6'},
    ]);
    const route = useRoute()

    onMounted(() => {
      getGrades()
    })

    const calculate = (name) => {
      let num = 0;
      let sum = 0;

      if (grades.value) {
        for (let grade of grades.value) {
          if (grade.Subject === name) {
            sum += parseInt(grade.Grade)
            num++;
          }
        }
      }

      return sum/num;
    };

    function getGrades(){
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }
      const params = new URLSearchParams({
        UserId: route.params.userId
      }).toString()

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'grades?' + params, config).then(resp => {
        grades.value = resp.data
      }).then(

      ).catch(err => {
        console.log(err)
      })
    }

    function addGrade(){
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }

      let m = {
        UserId: route.params.userId,
        Grade: selectedGrade.value,
        Subject: selectedSubject.value,
      }

      axios.post(process.env.VUE_APP_BACKEND_RESP_API + 'grades', m, config).then(resp => {
        console.log(resp)
      }).catch(err => {
        console.log(err)
      })

    }

    return {
      grades,
      calculate,
      subjects,
      selectedSubject,
      gradesList,
      selectedGrade,
      addGrade
    }
  }
};


</script>

<style lang="scss" scoped>
.p-rowgroup-footer td {
  font-weight: 700;
}

::v-deep(.p-rowgroup-header) {
  span {
    font-weight: 700;
  }

  .p-row-toggler {
    vertical-align: middle;
    margin-right: .25rem;
  }
}
</style>