<template>
 <div>
   <el-menu :default-active="activeIndex" class="el-menu-demo"
            mode="horizontal" @select="handleSelect" background-color="#252737" text-color="#fff">
     <el-menu-item><img src="../../static/img/logo_green.png" style="width: 50px;height: 50px;"></el-menu-item>
     <el-menu-item style="font-size: 25px;color: #5dd5c8">摘星</el-menu-item>
     <el-menu-item index="1"><router-link to="/main" style="text-decoration-line: none">首页</router-link></el-menu-item>
     <el-menu-item index="2"><router-link to="/fight" style="text-decoration-line: none">对战</router-link></el-menu-item>
     <el-submenu index="3">
       <template slot="title">题库</template>
       <el-menu-item index="3-1"><router-link to="/practice" style="color: white;text-decoration-line: none">专项练习</router-link></el-menu-item>
       <el-menu-item index="3-2"><router-link to="/interview" style="color: white;text-decoration-line: none">面试必备</router-link></el-menu-item>
     </el-submenu>
     <el-menu-item index="3"><router-link to="/circle" style="text-decoration-line: none">圈子</router-link></el-menu-item>
     <el-menu-item index="4" v-if="this.UserType === 'admin'"><router-link to="/manage" style="text-decoration-line: none">管理</router-link></el-menu-item>
     <el-submenu index="5" style="float: right">
       <template slot="title">{{UserName}}</template>
       <router-link to="/account/Info"><el-menu-item index="5-1" style="text-align: center">我的信息</el-menu-item></router-link>
       <router-link to="/"><el-menu-item index="5-2" style="text-align: center">返回首页</el-menu-item></router-link>
       <el-menu-item index="5-3" style="text-align: center" @click="logout">退出登录</el-menu-item>
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
      ruleForm:{
        UserName:"",
        Avatar:"",
        UserType:""
      }
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    },
    search (){

    },
    logout (){
      this.$router.push('/login')
    }
  },
  mounted () {
    this.search()
  }
}
</script>
<style>
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
  *{
    margin: 0;
    padding: 0;
  }

</style>
