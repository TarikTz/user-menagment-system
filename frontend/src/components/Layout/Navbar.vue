<template>
  <nav>
    <div class="container mx-auto">
      <div class="logo">UMS</div>
      <div class="navigation">
        <router-link to="/">Home</router-link>
        <router-link to="/about">About</router-link>
        <router-link to="/user/new" class="cta">Add New</router-link>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: "navBar",
};
</script>

<style></style>
