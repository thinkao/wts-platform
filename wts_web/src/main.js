// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import ElementUI from 'element-ui'
import api from './api/index'
import router from './router/router'
import './filters/index'
import axios from 'axios'

Vue.prototype.$api = api
Vue.prototype.$account = api.account
Vue.use(ElementUI)
Vue.config.productionTip = false

axios.defaults.withCredentials = true

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  render: h => h(App),
}).$mount('#app')