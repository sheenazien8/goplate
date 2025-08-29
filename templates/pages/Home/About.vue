<template>
  <div class="container">
    <div class="card">
      <h1 class="text-xl mb-4">{{ title }}</h1>
      <p class="text-gray-600 mb-4">{{ description }}</p>

      <div class="mt-4">
        <a href="/" class="btn">Back to Home</a>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HomeAbout',
  props: {
    title: String,
    description: String
  }
}
</script>
