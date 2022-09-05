<template>

  <div class="surface-section px-4 py-8 md:px-6 lg:px-8 text-center">
    <div class="text-center">
      <img src="../assets/mylearn_logo.png" alt="Image" height="150" class="mb-3">
    </div>
    <div class="text-700 text-sm mb-6">You are on teacher view</div>
    <div class="grid">
      <div class="col-12 md:col-4 mb-4 px-5">
        <router-link to="/gradeBook">
            <span class="p-3 shadow-2 mb-3 inline-block surface-card" style="border-radius: 10px">
                <i class="pi pi-desktop text-4xl text-blue-500"></i>
            </span>
        <div class="text-900 mb-3 font-medium">Online gradebook</div>
        <span class="text-700 text-sm line-height-3"></span>
        </router-link>

      </div>
      <div class="col-12 md:col-4 mb-4 px-5">
        <router-link to="/managePresentations">
            <span class="p-3 shadow-2 mb-3 inline-block surface-card" style="border-radius: 10px">
                <i class="pi pi-lock text-4xl text-blue-500"></i>
            </span>
        </router-link>
        <div class="text-900 mb-3 font-medium">Create/Edit Lesson</div>
        <span class="text-700 text-sm line-height-3">Create or edit your interactive presentations</span>
      </div>
      <div class="col-12 md:col-4 mb-4 px-5">
        <router-link to="/startLesson">
            <span class="p-3 shadow-2 mb-3 inline-block surface-card" style="border-radius: 10px">
                <i class="pi pi-check-circle text-4xl text-blue-500"></i>
            </span>
        <div class="text-900 mb-3 font-medium">Conduct Lesson</div>
        <span class="text-700 text-sm line-height-3">Run an interactive lesson</span>
        </router-link>
      </div>
      <div class="col-12 md:col-4 mb-4 px-5">
        <router-link to="/chatList">
            <span class="p-3 shadow-2 mb-3 inline-block surface-card" style="border-radius: 10px">
                <i class="pi pi-globe text-4xl text-blue-500"></i>
            </span>
        <div class="text-900 mb-3 font-medium">User Chat</div>
        <span class="text-700 text-sm line-height-3">Exchange messages with other users</span>
        </router-link>
      </div>
      <div class="col-12 md:col-4 mb-4 px-5">
        <router-link to="/manageClasses">
            <span class="p-3 shadow-2 mb-3 inline-block surface-card" style="border-radius: 10px">
                <i class="pi pi-github text-4xl text-blue-500"></i>
            </span>
        </router-link>
        <div class="text-900 mb-3 font-medium">Classes management</div>
        <span class="text-700 text-sm line-height-3">Create classes and attach users</span>
      </div>
      <div class="col-12 md:col-4 md:mb-4 mb-0 px-3">
        <router-link to="/classContent">
            <span class="p-3 shadow-2 mb-3 inline-block surface-card" style="border-radius: 10px">
                <i class="pi pi-shield text-4xl text-blue-500"></i>
            </span>
        </router-link>

        <div class="text-900 mb-3 font-medium">Homeworks and learning materials</div>
        <span class="text-700 text-sm line-height-3">Share files with students and see uploaded materials</span>
      </div>
    </div>
  </div>
</template>



<script>

export default {
  name: "TeacherMenu",
  components: {
  },
  setup() {
    return {
    }
  }
};


</script>

<style scoped lang="less"></style>