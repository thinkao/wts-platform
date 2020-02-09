import Vue from 'vue'
import Router from 'vue-router'

import Login from '../pages/Login.vue'
import Welcome from '../pages/Welcome'
import Navigation from '../pages/Navigation'
import Main from '../pages/Main'

import Info from '../pages/account/Info'
import InfoMain from '../pages/account/InfoMain'
import InfoUpdate from '../pages/account/InfoUpdate'
import InfoFocusPeo from '../pages/account/InfoFocusPeo'
import InfoDynamic from '../pages/account/InfoDynamic'
import InfoCollection from '../pages/account/InfoCollection'
import InfoNotes from '../pages/account/InfoNotes'
import InfoFeedback from '../pages/account/InfoFeedback'
import InfoSetting from '../pages/account/InfoSetting'
import InfoFoot from '../pages/account/InfoFoot'
import InfoLevel from '../pages/account/InfoLevel'
import InfoRecord from '../pages/account/InfoRecord'

Vue.use(Router)
let routes = [
  {path: '/', component: Welcome, name: '欢迎页'},
  {path: '/login', component: Login, name: '登录'},
  {path: '/navigation',
    redirect: '/account/info',
    component: Navigation,
    name: '导航',
    children: [
      {path: '/main', component: Main, name: '首页'},
      {path: '/account/info',
        redirect: '/account/main',
        component: Info,
        name: '用户信息',
        children: [
          {path: '/account/main', component: InfoMain, name: '主页'},
          {path: '/account/infoUpdate', component: InfoUpdate, name: '基本信息修改'},
          {path: '/account/infoFocusPeo', component: InfoFocusPeo, name: '关注的人'},
          {path: '/account/infoCollection', component: InfoCollection, name: '我的收藏'},
          {path: '/account/infoDynamic', component: InfoDynamic, name: '我的动态'},
          {path: '/account/infoNotes', component: InfoNotes, name: '个人笔记'},
          {path: '/account/infoFeedback', component: InfoFeedback, name: '我的动态'},
          {path: '/account/infoSetting', component: InfoSetting, name: '账号设置'},
          {path: '/account/infoFoot', component: InfoFoot, name: '我的足迹'},
          {path: '/account/infoLevel', component: InfoLevel, name: '等级信息'},
          {path: '/account/infoRecord', component: InfoRecord, name: '我的战绩'}
        ]}
    ]},
  {path: '*', redirect: '/404', name: '404'}
]

export default routes
