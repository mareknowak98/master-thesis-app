<template>
  <div class="home">
    <div class="grid">
      <div class="col-6 col-offset-3">
        <div class="surface-card p-4 shadow-2 border-round w-full">
          <NavBar/>
          <div class="grid grid-nogutter surface-section text-800">
            <div class="col-12 md:col-6 p-0 text-center md:text-left flex align-items-center ">
              <section>
                <span class="block text-6xl font-bold mb-1">Create the lessons and manage your class</span>
                <div class="text-6xl text-primary font-bold mb-3">without any problems</div>
                <p class="mt-0 mb-4 text-700 line-height-3">A portal built entirely on AWS cloud services. (not only EC2) </p>

                <Button label="Project repository" onclick="location.href='https://github.com/mareknowak98/master-thesis-app';" type="button" class="mr-3 p-button-raised"></Button>
<!--                <Button label="XYZ" type="button" class="p-button-outlined"></Button>-->
              </section>
            </div>
            <div class="col-12 md:col-6  overflow-hidden">
              <img src="../assets/mylearn_logo.png" alt="Image" class="md:ml-auto block md:h-full" style="clip-path: polygon(8% 0, 100% 0%, 100% 100%, 0 100%)">
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue";
import TabMenu from 'primevue/tabmenu';

export default {
  name: "HomeView",
  components: {
    NavBar,
    TabMenu,
  },
};
</script>
