<template>
  <div class="container">
    <div class="card">
      <h1 class="text-xl mb-4">Welcome to GoPlate</h1>
      <p class="text-gray-600 mb-4">{{ message }}</p>

      <div class="card" v-if="user">
        <h2 class="text-xl mb-4">User Information</h2>
        <p><strong>Name:</strong> {{ user.name }}</p>
        <p><strong>Email:</strong> {{ user.email }}</p>
      </div>

      <div class="mt-4">
        <a href="/about" class="btn">About Page</a>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HomeIndex',
  props: {
    message: String,
    user: Object
  }
}
</script>
