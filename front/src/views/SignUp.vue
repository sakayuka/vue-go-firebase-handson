<template>
  <div class="signup">
    <h2>Sign up</h2>
    <input type="text" placeholder="Email" v-model="email" />
    <input type="password" placeholder="Password" v-model="password" />
    <button @click="signUp">Register</button>
    <p>
      Do you have an account?
      <router-link to="/signin">sign in now!!</router-link>
    </p>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from 'vue'
import firebase from 'firebase'

export default defineComponent({
  name: 'SignUp',
  setup() {
    const state = reactive({
      email: '',
      password: '',
    })
    const signUp = () => {
      firebase
        .auth()
        .createUserWithEmailAndPassword(state.email, state.password)
        .then((res) => {
          console.log('Create account: ', res)
        })
        .catch((e) => {
          console.log(e.message)
        })
    }
    return {
      ...toRefs(state),
      signUp,
    }
  },
})
</script>

<style lang="stylus" scoped>
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
input
  margin 10px 0
  padding 10px
.signup
  margin-top 20px
  display flex
  flex-flow column nowrap
  justify-content center
  align-items center
</style>
