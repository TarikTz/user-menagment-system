<template>
  <div class="snackbar" v-show="snackbar">
    {{ this.$store.state.notification.message }}
    <div class="close" @click="closeSnackBar">close</div>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
  computed: {
    snackbar() {
      return this.$store.state.snackbar;
    },
  },
  methods: {
    closeSnackBar() {
      this.$store.commit("setSnackBar", false);
    },
  },
  watch: {
    snackbar() {
      let that = this;
      setTimeout(function () {
        that.closeSnackBar();
      }, 3000);
    },
  },
};
</script>

<style scoped>
.snackbar {
  position: absolute;
  background: green;
  color: white;
  padding: 10px;
  border-radius: 3px;
  box-shadow: 0px 3px 6px rgba(0, 0, 0, 0.3);
  bottom: 30px;
  left: 0;
  right: 0;
  max-width: 300px;
  margin: 0px auto;
  display: flex;
  justify-content: space-between;
  transition: all 0.25s ease-in-out;
}

.success {
  background: green;
  color: white;
}

.error {
  background: red;
  color: white;
}

.close {
  cursor: pointer;
}
</style>
