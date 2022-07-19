<template>
  <div class="home">
    <div class="grid">
      <div class="col-6 col-offset-3">
        <div class="surface-card p-4 shadow-2 border-round w-full">
          <NavBar/>
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
