<template>
 <div>
   <el-menu :default-active="activeIndex" class="el-menu-demo"
            mode="horizontal" @select="handleSelect" background-color="#252737" text-color="#fff">
     <el-menu-item><img src="../../static/img/logo1_green.png" style="width: 50px;height: 50px;"></el-menu-item>
     <el-menu-item style="font-size: 25px;color: #5dd5c8">摘星</el-menu-item>
     <el-menu-item index="1"><router-link to="/main">首页</router-link></el-menu-item>
     <el-menu-item index="2">处理中心</el-menu-item>
     <el-submenu index="3">
       <template slot="title">我的工作台</template>
       <el-menu-item index="3-1">选项1</el-menu-item>
       <el-menu-item index="3-2">选项2</el-menu-item>
       <el-menu-item index="3-3">选项3</el-menu-item>
     </el-submenu>
     <el-menu-item index="3">消息中心</el-menu-item>
     <el-menu-item index="4"><a href="https://www.ele.me" target="_blank">订单管理</a></el-menu-item>
     <el-submenu index="9" style="float: right">
       <template slot="title">{{NiName}}</template>
       <router-link to="/account/Info"><el-menu-item index="9-2" style="text-align: center">我的信息</el-menu-item></router-link>
       <router-link to="/"><el-menu-item index="9-2" style="text-align: center">返回首页</el-menu-item></router-link>
       <router-link to="/Login"><el-menu-item index="9-2" style="text-align: center">退出登录</el-menu-item></router-link>
     </el-submenu>
     <div style="float: right;margin-top: 10px">
       <el-avatar> user </el-avatar>
     </div>
   </el-menu>
   <el-container>
     <router-view></router-view>
     <el-footer>Footer</el-footer>
   </el-container>
 </div>
</template>
<script>
export default {
  data () {
    return {
      activeIndex: '1',
      NiName: '张三',
      NiName1: sessionStorage.getItem('username')
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    }
  },
  mounted () {
    this.search()
  }
}
</script>
<style scoped>
  @import url("//unpkg.com/element-ui@2.13.0/lib/theme-chalk/index.css");
  .el-header, .el-footer {
    background-color: #252737;
    color: #333;
    text-align: center;
    line-height: 60px;
  }
  body > .el-container {
    margin-bottom: 40px;
  }

</style>
