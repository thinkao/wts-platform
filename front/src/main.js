// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import VueRouter from 'vue-router'
import ElementUI from 'element-ui'
import api from './api/index'
import routes from './router/router'
import './filters/index'

Vue.prototype.$api = api
Vue.prototype.$account = api.account
Vue.use(ElementUI)
Vue.config.productionTip = false

const router = new VueRouter({
  routes
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
