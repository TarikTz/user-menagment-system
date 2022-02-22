<template>
  <div class="wrapper">
    <NavBar />
    <router-view />
    <SnackBar />
  </div>
</template>

<script>
import NavBar from "@/components/Layout/Navbar";
import SnackBar from "@/components/Notification/SnackBar";

export default {
  components: {
    NavBar,
    SnackBar,
  },
};
</script>
<style></style>
