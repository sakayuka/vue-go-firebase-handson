<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <h2>Essential Links</h2>
    <button @click="apiPublic">public</button>
    <button @click="apiPrivate">private</button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import axios from 'axios'

export default defineComponent({
  name: 'HelloWorld',
  props: {
    msg: String,
  },
  setup(props, { emit }) {
    const apiPublic = async () => {
      const res = await axios.get('http://localhost:8000/public')
      emit('update:msg', res.data)
    }
    const apiPrivate = async () => {
      const res = await axios.get('http://localhost:8000/private')
      emit('update:msg', res.data)
    }
    return {
      apiPublic,
      apiPrivate,
    }
  },
})
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
h1, h2
  font-weight normal
ul
  list-style-type none
  padding 0
li
  display inline-block
  margin 0 10px
a
  color #42b983
button
  margin 10px
  padding 10px
</style>
