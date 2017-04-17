<template>
  <div class="register">
    <h2>Register</h2>

    <div class="form">
      <div class="form__group">
        <label for="username">Username:</label>
        <input v-model="username" v-bind:class="{ error: usernameError }">
      </div>

      <div class="form__group">
        <label for="email">Email:</label>
        <input v-model="email" type="email" v-bind:class="{ error: emailError }">
      </div>

      <div class="form__group">
        <label for="weight">Weight:</label>
        <input v-model="weight" type="number" v-bind:class="{ error: weightError }">
      </div>

      <div class="form__group">
        <label for="password">Password</label>
        <input v-model="password" type="password" v-bind:class="{ error: passwordError }">
      </div>

      <div class="form__group">
        <label for="confirmation">Confirmation</label>
        <input v-model="confirmation" type="password" v-bind:class="{ error: confirmationError }">
      </div>

      <div class="form__group-button">
        <button @click="register" class="form__button">Submit</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'register',
  data () {
    return {
      username: '',
      usernameError: false,
      email: '',
      emailError: false,
      password: '',
      passwordError: false,
      confirmation: '',
      confirmationError: false,
      weight: null,
      weightError: false
    }
  },
  methods: {
    register () {
      const valid = this.validate()
      const user = {
        username: this.username,
        email: this.email,
        password: this.password,
        weight: this.weight
      }

      if (valid) {
        this.$store.dispatch('register', user)
      }
    },
    validate () {
      const errors = ['username', 'email', 'password', 'confirmation', 'weight'].map(field => {
        const model = this[field]

        if (model === null || model === '') {
          this[`${field}Error`] = true
        } else {
          this[`${field}Error`] = false
        }

        return this[`${field}Error`]
      })

      if (errors.indexOf(true) !== -1) {
        return false
      } else {
        return true
      }
    }
  }
}
</script>

<style scoped>
  .error {
    border: 1px solid red;
  }
</style>
