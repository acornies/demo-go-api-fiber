
const demoApp = {
  data: () => ({
    todos: null,
    loading: false,
    timer: null
  }),
  methods: {
    fetchTodos() {
      this.loading = true;
      fetch(`/demo/todos`)
        .then(response => response.json())
        .then(response => {
          this.loading = false;
          this.todos = response.Todos
        })
        .catch(err => {
          console.log(err.message || err);
          this.loading = false
        })
    }
  },
  mounted: function () {
    this.fetchTodos();
    this.timer = setInterval(this.fetchTodos, 5000);
  }
};

Vue.createApp(demoApp).mount('#name-list')
