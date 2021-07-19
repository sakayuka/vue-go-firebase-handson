<template>
  <div class="hello">
    <h1>Hello {{ name }}!</h1>
    <h1>{{ msg }}</h1>
    <h2>Essential Links</h2>
    <button @click="signOut">Sign out</button>
    <button @click="apiPublic">public</button>
    <button @click="apiPrivate">private</button>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import firebase from 'firebase'

export default defineComponent({
  name: 'HelloWorld',
  props: {
    msg: String,
  },
  setup(props, { emit }) {
    const router = useRouter()
    const signOut = () => {
      firebase
        .auth()
        .signOut()
        .then(() => {
          localStorage.removeItem('jwt')
          router.push('/signin')
        })
    }
    const apiPublic = async () => {
      const res = await axios.get('http://localhost:8000/public')
      console.log(res)
      emit('update:msg', res.data)
    }
    const apiPrivate = async () => {
      const res = await axios.get('http://localhost:8000/private', {
        headers: { Authorization: `Bearer ${localStorage.getItem('jwt')}` },
      })
      console.log(res)
      emit('update:msg', res.data)
    }
    const name = computed(() => {
      const user = firebase.auth().currentUser
      if (user === null) {
        return ''
      }
      return user.email
    })
    return {
      name,
      signOut,
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
