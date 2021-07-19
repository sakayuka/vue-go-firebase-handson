<template>
  <div class="signin">
    <h2>Sign in</h2>
    <input type="text" placeholder="Email" v-model="email" />
    <input type="password" placeholder="Password" v-model="password" />
    <button>Signin</button>
    <p>
      You don't have an account?
      <router-link to="/signup">create account now!!</router-link>
    </p>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from 'vue'

export default defineComponent({
  name: 'SignUp',
  setup() {
    const state = reactive({
      email: '',
      password: '',
    })

    return {
      ...toRefs(state),
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
.signin
  margin-top 20px
  display flex
  flex-flow column nowrap
  justify-content center
  align-items center
</style>
