<template>
  <div>
    <div>
      <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
        <el-menu-item><el-input v-model="input" placeholder="请输入搜索内容" style="width: 260px;float: left;display: inline"></el-input></el-menu-item>
        <el-menu-item index="2">处理中心</el-menu-item>
        <el-submenu index="3">
          <template slot="title">关于我们</template>
          <el-menu-item index="3-1">选项1</el-menu-item>
          <el-menu-item index="3-2">选项2</el-menu-item>
          <el-menu-item index="3-3">选项3</el-menu-item>
        </el-submenu>
        <el-menu-item index="4" @click="loginTo">登录</el-menu-item>
        <el-menu-item index="5">注册</el-menu-item>
      </el-menu>
    </div>
    <img src="../../static/img/hengtiao.png" style="width: 100%">
    <div>
      <el-row :gutter="60" style="margin: 50px">
        <el-col :span="8">
          <el-card shadow="hover">
            <img src="../../static/img/duizhan.jpeg" style="width: 100%;;height: 230px">
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="hover">
            <img src="../../static/img/lianxi.jpeg" style="width: 100%;height: 230px">
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="hover">
            <img src="../../static/img/shequ.jpeg" style="width: 100%;;height: 230px">
          </el-card>
        </el-col>
      </el-row>
      <el-row :gutter="60" style="margin: 50px">
        <el-col :span="8">
          <span>Passion versus mode, let you immersed in the joy of the confrontation experience,Passion versus mode, let you immersed in the joy of the confrontation experience</span>
        </el-col>
        <el-col :span="8">
          <span>Special practice mode, mother in also do not have to worry about my study,Special practice mode, mother in also do not have to worry about my study</span>
        </el-col>
        <el-col :span="8">
          <span>Community sharing mode, where you can communicate with people in the community indefinitely,Community sharing mode, where you can communicate with people in the community indefinitely</span>
        </el-col>
      </el-row>
    </div>
    <el-container>
      <span style="margin: auto">
        Reach the star platform original @ WS
      </span>
    </el-container>
  </div>
</template>

<script>
export default {
  name: 'Welcome',
  data () {
    return {
      input: '',
      activeIndex: '1',
      activeIndex2: '1'
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    },
    loginTo () {
      this.$router.push('/login')
    }
  },
}
</script>

<style scoped>
  @import url("//unpkg.com/element-ui@2.13.0/lib/theme-chalk/index.css");
  .el-container {
    height: 200px;
    background: #f6f6f8;
  }
</style>
