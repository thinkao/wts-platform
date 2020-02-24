<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App',
  components: {
  }
}
</script>

<style scoped>
  #app {
    position: relative;
    top: 0px;
    bottom: 0px;
    width: 100%;
    overflow: hidden;
    background: #fff;
    min-height: 500px;
    min-width: 1366px;
  }
</style>
