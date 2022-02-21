<template>
  <div class="container mx-auto add-new">
    <h2>{{ title }}</h2>
    <form action="" @submit.prevent="addNewUser">
      <div class="form-group w-half">
        <label for="first_name">First Name</label>
        <input
          type="text"
          name="first_name"
          id="first_name"
          v-model="user.first_name"
        />
      </div>
      <div class="form-group w-half">
        <label for="last_name">Last Name</label>
        <input
          type="text"
          name="last_name"
          id="last_name"
          v-model="user.last_name"
        />
      </div>
      <div class="form-group w-half">
        <label for="first_name">Email</label>
        <input type="text" name="email" id="email" v-model="user.email" />
      </div>
      <div class="form-group w-half">
        <label for="first_name">Username</label>
        <input
          type="text"
          name="username"
          id="username"
          v-model="user.username"
        />
      </div>
      <div class="form-group w-half">
        <label for="first_name">Password</label>
        <input
          type="password"
          name="password"
          id="password"
          v-model="user.password"
        />
      </div>
      <div class="form-group w-half">
        <label for="first_name">Status</label>
        <select name="status" id="status" v-model="user.status">
          <option value="0" selected>Disabled</option>
          <option value="1">Enabled</option>
        </select>
      </div>
      <div class="form-group w-half">
        <input type="submit" name="submit" value="Submit" class="submit" />
      </div>
    </form>

    <div class="messages" v-if="this.$store.state.notification.status">
      {{ this.$store.state.notification.message }}
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      submitText: "Submit",
      title: "Add New",
    };
  },
  computed: {
    user() {
      return this.$store.state.user;
    },
  },
  created() {
    console.log(this.$route.params.id);
    if (this.$route.params.id != "new") {
      this.$store.dispatch("getUser", this.$route.params.id);
      this.submitText = "Update";
      this.title = "Edit User";
    }
  },
  methods: {
    addNewUser() {
      console.log("validate here");
      this.user.status = parseInt(this.user.status);
      this.$store.dispatch("createUser", this.user);
    },
  },
};
</script>

<style scoped>
.add-new {
  margin-top: 50px;
}

.w-half {
  width: 50%;
  float: left;
}

form {
  display: flex;
  flex-wrap: wrap;
}

.form-group {
  display: flex;
  flex-direction: column;
  padding: 5px;
}

.form-group label {
  margin-bottom: 5px;
}

.form-group input,
select {
  padding: 5px;
  border: 1px solid #5bb7d6;
  border-radius: 5px;
}

.messages {
  display: block;
  margin-top: 30px;
  width: 100%;
  transition: all 0.25s ease-in-out;
}

input.submit {
  background: #5bb7d6;
  color: #000;
  padding: 5px;
  font-weight: bold;
}
</style>
