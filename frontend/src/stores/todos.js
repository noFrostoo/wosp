import { defineStore } from 'pinia'

export const useTodos = defineStore('todos', {
  state: () => ({
    todos: [],
  }),
  getters: {
    finishedTodos(state) {
      return state.todos.filter((todo) => todo.done)
    },
    unfinishedTodos(state) {
      return state.todos.filter((todo) => !todo.done)
    },
    getAll(state){
      return state.todos
    }

  },
  actions: {
    addTodo(title, description) {
      // you can directly mutate the state
      this.todos.push({ title: title, description: description, done: false })
    },
    deleteTodo() {

    },
    retrieveAll() {

    },
  },
})