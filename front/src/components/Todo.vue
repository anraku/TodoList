<template>
  <div>
    <el-input placeholder="タスクを追加" v-model="newTask" v-on:change="addTask()"></el-input>
    <ul class="list-group">
      <h2>残っているタスク</h2>
      <hr>
      <li v-for="todo in todos" :key="todo.id" class="list-group-item">
        <span v-if="!todo.done">
          <el-checkbox v-on:change="updateTask(todo)">{{ todo.text }}</el-checkbox>
        </span>
      </li>
    </ul>
    <ul class="list-group">
      <h2>終わったタスク</h2>
      <hr>
      <li v-for="todo in todos" :key="todo.id" class="list-group-item">
        <span v-if="todo.done">
          <el-checkbox v-on:change="updateTask(todo)"><s>{{ todo.text }}</s></el-checkbox>
        </span>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Todo',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App',
      newTask: '',
      todos: []
    }
  },
  created: function () {
    axios.get('http://localhost:1323/tasks')
      .then((response) => {
        console.log(response)
        this.todos = response.data.items || []
      })
      .catch((error) => {
        console.log(error)
      })
  },
  methods: {
    addTask: function (todo) {
      let params = new URLSearchParams()
      params.append('text', this.newTask)
      params.append('done', false)
      axios.post('http://localhost:1323/tasks', params)
        .then((response) => {
          this.todos.unshift(response.data)
          this.newTask = ''
        })
        .catch((error) => {
          console.log(error)
        })
    },
    updateTask: function (todo) {
      let params = new URLSearchParams()
      params.append('done', !todo.done)
      axios.put('http://localhost:1323/tasks/' + todo.id, params)
        .then((response) => {
          todo.done = !todo.done
          console.log(response)
        })
        .catch((error) => {
          console.log(error)
        })
    }
  }
}
</script>

<style>
.list-group {
text-align: left;
width: 50%;
padding-left: 30%;
}
.list-group-item {
list-style-type: none;
}
</style>
