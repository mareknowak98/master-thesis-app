<template>
  <div class="home">
    <div class="grid">
      <div class="col-6 col-offset-3">
        <HelloWorld/>
      </div>
    </div>

  </div>
</template>

<script>
import HelloWorld from "@/components/HelloWorld.vue";

export default {
  name: "HomeView",
  components: {
    HelloWorld,
  },
};
</script>
