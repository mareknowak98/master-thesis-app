<template>
  <LoggedNavbar/>
  <div class="row">
    <div class="column">
      <h2>Files list</h2>
      <DataTable :value="files" :paginator="true" :rows="10"
                 paginatorTemplate="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown"
                 :rowsPerPageOptions="[10,20,50]" responsiveLayout="scroll"
                 currentPageReportTemplate="Showing {first} to {last} of {totalRecords}">
        <Column field="fileName" header="Class Id">
          <template #body="slotProps">
            <a v-text="slotProps.data.fileName"/>
          </template>
        </Column>
        <Column field="fileName" header="Download file">
          <template #body="slotProps">
            <i @click="downloadFile(slotProps.data.fileName)" class="pi pi-download"></i>
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
    <h2>Upload new file</h2>
<!--      <FileUpload ref="newFile" name="demo[]" @upload="onUpload" @before-upload="uploadFile2(this.files)" :multiple="true" accept="image/*" :maxFileSize="1000000">-->
<!--        <template #empty>-->
<!--          <p>Drag and drop files to here to upload.</p>-->
<!--        </template>-->
<!--      </FileUpload>-->

<!--      <div class="mb-3">-->
<!--        <label for="formFile" class="form-label">Default file input example</label>-->
<!--        <input ref="newFile" class="form-control" type="file" id="newFile">-->
<!--      </div>-->
      <input ref="newFile" type="file" id="newFile">
      <Button @click="uploadFile()" label="Add new file" />

<!--      <label>File-->
<!--        <input type="file" @change="handleFileUpload( $event )"/>-->
<!--      </label>-->
<!--      <Button @click="submitFile()" label="Add new file" />-->

    </div>
  </div>
</template>



<script>

import DataTable  from 'primevue/datatable';
import Column from 'primevue/column';
import Divider from 'primevue/divider';
import FileUpload from 'primevue/fileupload';

import LoggedNavbar from "@/components/LoggedNavbar";
import { ref, onMounted } from "vue";
import axios from "axios";
import { TokenService } from "@/store/tokenService";
import { useRoute } from "vue-router/dist/vue-router";

import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap"

export default {
  name: "S3ClassMaterialsView",
  components: {
    LoggedNavbar,
    DataTable,
    Column,
    Divider,
    FileUpload
  },
  setup() {
    // let files = ref(null)
    let filesFormatted = ref([])
    const route = useRoute()
    //temp
    let files = ref([
    {
      fileName: "1AB/file_new.jpg"
    },   {
      fileName: "1AB/new_file2.jpg"
    }])
    let newFile = ref(null)


    onMounted(() => {
      // getFiles()

    })

    function getFiles() {
     console.log("getfiles")
      let config = {
        headers: {
          Authorization: TokenService.getToken(),
        }
      }

      const params = new URLSearchParams({
        folder: route.params.classId + '/'
      }).toString()

      axios.get(process.env.VUE_APP_BACKEND_RESP_API + 'files?' + params, config).then(resp => {
        files.value = resp.data
      }).catch(err => {
        console.log(err)
      })

    }

    function uploadFile2(sth) {
      console.log("---------uplkoad")
      console.log(sth[0])


    }
    function uploadFile(sth) {
      // newFile.value = this.$refs.file.files[0];

      let config = {
        headers: {
          Authorization: TokenService.getToken(),
          'Content-Type': 'multipart/form-data'
        }
      }
      console.log(config)

      console.log("---------upload2")
      console.log("selected file",newFile.value.files)
      let filename = newFile.value.files[0].name
      console.log("selected filename",filename)
      const formData = new FormData();
      formData.append('src', newFile.value.files[0]);
      console.log("---------filepath")

      let filepath = process.env.VUE_APP_S3_BUCKET_NAME + "/" + route.params.classId + '/' + filename
      console.log(filepath)

      let path = process.env.VUE_APP_S3_API + 's3?key=' + filepath
      console.log(path)

      axios.put(path, formData, config).then(resp => {
        console.log(resp)
      }).catch(err => {
        console.log(err)
      })
    }

    function handleFileUpload(){
      newFile.value = this.$refs.file.files[0];
    }

    function submitFile(){

    }

    function test(){
      console.log(newFile.value)
    }

    const onUpload = () => {
      toast.add({severity: 'info', summary: 'Success', detail: 'File Uploaded', life: 3000});
    }

    function downloadFile(filename){
      let config = {
        headers: {
          Authorization: TokenService.getIdToken(),
        }
      }

      let filepath = process.env.VUE_APP_S3_BUCKET_NAME + "/" + filename

      console.log(filename)
      console.log(filepath)

      let path = process.env.VUE_APP_S3_API + 's3?key=' + filepath
      console.log(path)

      axios({
        url: path,
        method: 'GET',
        responseType: 'blob',
        headers: {
          Authorization: TokenService.getToken()
        }
      }).then((response) => {
        var fileDownload = require('js-file-download');
        fileDownload(response.data, filename);
      });
    }


    return {
      files,
      filesFormatted,
      downloadFile,
      newFile,
      uploadFile,
      onUpload,
      test,
      uploadFile2,
      submitFile,
      handleFileUpload
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


</style>