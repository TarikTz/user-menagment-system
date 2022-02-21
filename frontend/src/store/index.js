import { createStore } from "vuex";
import axios from "axios";

export default createStore({
  state: {
    api: "http://localhost:3010",
    users: [],
    users_count: 0,
    user: {
      first_name: "",
      last_name: "",
      username: "",
      email: "",
      password: "",
      status: 0,
    },
    filters: {
      page: 1,
      limit: 10,
      status: -1,
      order: "",
    },
    notification: {
      status: false,
      message: "",
    },
  },
  getters: {},
  mutations: {
    setUsers(state, payload) {
      if (payload) return (state.users = payload);

      return (state.users = []);
    },
    setNotification(state, payload) {
      return (
        (state.notification.status = true),
        (state.notification.message = payload)
      );
    },
    setUser(state, payload) {
      if (payload) return (state.user = payload);

      return {};
    },
    setUsersCount(state, payload) {
      if (payload) return (state.users_count = payload);

      return (state.users_count = 0);
    },
    updateFilters(state) {
      state.filters.page = 1;
    },
  },
  actions: {
    getUsers(context, payload) {
      return new axios.get(this.state.api + "/users", {
        params: payload,
      })
        .then((response) => {
          context.commit("setUsers", response.data.users);
          context.commit("setUsersCount", response.data.users_count);
          console.log(response.data);
        })
        .catch((err) => {
          console.log(err);
        });
    },
    getUser(context, payload) {
      return new axios.get(this.state.api + "/user/" + payload)
        .then((response) => {
          context.commit("setUser", response.data);
        })
        .catch((err) => {
          console.log(err);
        });
    },
    createUser(context, payload) {
      return new axios.post(this.state.api + "/user", payload)
        .then(() => {
          context.commit("setNotification", "User created successfully");
        })
        .catch((err) => {
          console.log(err);
        });
    },
    deleteUser(context, payload) {
      return new axios.delete(this.state.api + "/user/" + payload)
        .then((response) => {
          context.commit("setNotification", response.data);
          this.dispatch("getUsers", this.state.filters);
        })
        .catch((err) => {
          console.log(err);
        });
    },
    filterUsers(context) {
      context.commit("updateFilters");
      this.dispatch("getUsers", this.state.filters);
    },
  },
  modules: {},
});
