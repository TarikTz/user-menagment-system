<template>
  <div class="modal" v-if="value">
    <div class="modal-content">
      <div class="modal-content-title">Delete</div>
      <div class="modal-content-text">
        Are you sure you wan't to delete this user?
      </div>
      <div class="modal-content-actions">
        <button class="text error" @click="closeModal">Cancel</button>
        <button class="text success" @click="confirmDelete">Confirm</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ["value"],
  methods: {
    closeModal() {
      this.$emit("cancel");
    },
    confirmDelete() {
      this.$emit("confirm");
    },
  },
};
</script>

<style>
.modal {
  width: 100vw;
  height: 100vh;
  position: absolute;
  background: rgba(0, 0, 0, 0.3);
  top: 0;
  left: 0;
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  max-width: 300px;
  background: #fff;
  padding: 20px;
  border-radius: 3px;
}

.modal-content-title {
  font-weight: bold;
  margin-bottom: 10px;
}

.modal-content-text {
  margin-bottom: 10px;
}

.modal-content-actions {
  display: flex;
  justify-content: end;
}

button.text {
  border: 0;
  background: transparent;
  padding: 0px 10px;
  cursor: pointer;
}

.error {
  color: red;
}

.success {
  color: green;
}
</style>
