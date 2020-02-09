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
    <el-container>
      <el-footer>Footer</el-footer>
    </el-container>

  </div>
</template>

<script>
export default {
  name: 'WelcomePage',
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
  }
}
</script>

<style scoped>
  @import url("//unpkg.com/element-ui@2.13.0/lib/theme-chalk/index.css");
  .el-footer {
    text-align: center;
    line-height: 200px;
  }
  .el-container {
    height: 200px;
    background: #f6f6f8;
  }
</style>
