<template>
  <div class="home container mx-auto">
    <div class="tableFilters">
      <h2>Users List</h2>
      <div class="filters">
        <div class="filter">
          <label for="status">Filter</label>
          <select
            name="status"
            id="status"
            v-model="this.$store.state.filters.status"
            @change="filterUsers"
          >
            <option value="-1" selected>All</option>
            <option value="0" selected>Disabled</option>
            <option value="1" selected>Enabled</option>
          </select>
        </div>
        <div class="filter">
          <label for="order">Order</label>
          <select
            name="order"
            id="order"
            v-model="this.$store.state.filters.order"
            @change="filterUsers"
          >
            <option
              :value="item.value"
              v-for="(item, index) in orderItems"
              :key="index"
              :selected="
                item.value == this.$store.state.filters.order ? true : false
              "
            >
              {{ item.name }}
            </option>
          </select>
        </div>
      </div>
    </div>

    <div class="tableHeaders">
      <div><strong>ID</strong></div>
      <div><strong>Full Name</strong></div>
      <div><strong>Username</strong></div>
      <div><strong>Email</strong></div>
      <div class="sm-col"><strong>Status</strong></div>
      <div class="sm-col"><strong>Actions</strong></div>
    </div>
    <div class="userTable" v-for="user in users" :key="user.id">
      <div>{{ user.id }}</div>
      <div>{{ user.first_name }} {{ user.last_name }}</div>
      <div>{{ user.username }}</div>
      <div>{{ user.email }}</div>
      <div class="sm-col">{{ getStatusLabel(user.status) }}</div>
      <div class="sm-col">
        <button class="icon" @click="editUser(user.id)">
          <i class="bi bi-pencil-square"></i>
        </button>
        <button class="icon" @click="deleteItem(user.id)">
          <i class="bi bi-trash"></i>
        </button>
      </div>
    </div>
    <div class="pagination">
      <div class="prev" @click="prevPage()">Prev</div>
      <div
        class="page"
        @click="goToPage(index)"
        v-bind:class="this.$store.state.filters.page == index ? 'current' : ''"
        v-for="index in pages"
        :key="index"
      >
        {{ index }}
      </div>
      <div class="next" @click="nextPage()">Next</div>
    </div>
    <Modal :value="dialog" @cancel="cancel" @confirm="deleteUser" />
  </div>
</template>

<script>
import Modal from "@/components/ModalView";

export default {
  name: "HomeView",
  components: { Modal },
  data() {
    return {
      dialog: false,
      deleteCandidate: null,
      orderItems: [
        { name: "default", value: "" },
        { name: "username (ASC)", value: "username ASC" },
        { name: "username (DESC)", value: "username DESC" },
        { name: "email (ASC)", value: "email ASC" },
        { name: "email (DESC)", value: "email DESC" },
        { name: "name (ASC)", value: "first_name ASC" },
        { name: "name (DESC)", value: "first_name DESC" },
      ],
    };
  },
  computed: {
    users() {
      return this.$store.state.users;
    },
    pages() {
      return Math.ceil(
        this.$store.state.users_count / this.$store.state.filters.limit
      );
    },
  },
  methods: {
    getUsers() {
      this.$store.dispatch("getUsers", this.$store.state.filters);
    },
    deleteItem(userID) {
      console.log("detele action", userID);
      this.dialog = true;
      this.deleteCandidate = userID;
    },
    editUser(userID) {
      this.$router.push({ name: "edit", params: { id: userID } });
    },
    cancel() {
      this.dialog = false;
      this.deleteCandidate = null;
    },
    deleteUser() {
      this.$store.dispatch("deleteUser", this.deleteCandidate);
      this.dialog = false;
    },
    getStatusLabel(statusID) {
      if (statusID == 0) return "Disabled";

      return "Enabled";
    },
    nextPage() {
      this.$store.state.filters.page += 1;
      this.$store.dispatch("getUsers", this.$store.state.filters);
    },
    prevPage() {
      if (this.$store.state.filters.page != 1)
        this.$store.state.filters.page -= 1;
      this.$store.dispatch("getUsers", this.$store.state.filters);
    },
    goToPage(index) {
      this.$store.state.filters.page = index;
      this.$store.dispatch("getUsers", this.$store.state.filters);
    },
    filterUsers() {
      this.$store.dispatch("filterUsers");
    },
  },
  created() {
    this.getUsers();
  },
};
</script>

<style scoped>
.home {
  margin-top: 50px;
}

.tableHeaders {
  display: flex;
}

.userTable {
  display: flex;
}

.tableHeaders div {
  width: 20%;
  background: #5bb7d6;
  padding: 5px;
  border: 1px solid #fff;
}

.userTable div {
  width: 20%;
  padding: 5px;
}

div.sm-col {
  width: 10%;
}

.pagination {
  margin-top: 50px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.prev,
.next {
  color: #5bb7d6;
  cursor: pointer;
}

.pagination .page {
  color: #5bb7d6;
  margin: 0px 10px;
  cursor: pointer;
  transition: all 0.25s ease-in-out;
}

.pagination .page:hover {
  background: #5bb7d638;
}

.pagination .current {
  background: #5bb7d6;
  padding: 2px 7px;
  color: #000;
}

.tableFilters {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
  background: #5bb7d6;
  padding: 10px;
}

.tableFilters .filters {
  display: flex;
}

.filter {
  display: flex;
  flex-direction: column;
  margin-left: 20px;
}

.filter label {
  margin-bottom: 5px;
}

.filter select {
  padding: 5px;
  border: 1px solid #5bb7d6;
  border-radius: 5px;
}

button {
  cursor: pointer;
}
</style>
