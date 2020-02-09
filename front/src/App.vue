<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
import Navigation from './pages/Navigation'
export default {
  name: 'App',
  components: {
    'Navigation': Navigation
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
