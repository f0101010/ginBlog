<template>
  <router-view :key="key" />
</template>

<script>
export default {
  name: 'App',
  computed: {
    key() {
      return this.$route.name ? this.$route.name + +new Date() : this.$route + +new Date()
    }
  }
}
</script>