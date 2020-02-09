<template>
  <div class="body">
    <div>
      <el-card class="box-card">
        <div style="float: left;">
          <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
        </div>
        <div style="float: left;margin-left: 15px">
          I<div style="float: left">
          <span>张三</span>
          <span>ID:53665522</span>
          </div>
          <div>
            <span>这里是我的个人宣言......</span>
          </div>
        </div>
        <div style="float: right">
          <el-button type="success" @click="updateInfo">编辑</el-button>
        </div>
      </el-card>
    </div>
    <div class="list">
      <el-row class="tac">
        <el-col :span="6">
          <el-menu
            default-active="1"
            class="el-menu-vertical-demo"
            @open="handleOpen"
            @close="handleClose">
            <el-menu-item index="1">
              <i class="el-icon-menu"></i>
              <router-link to="/account/main" class="router-link-active"><span>主页</span></router-link>
            </el-menu-item>
            <el-menu-item index="2">
              <i class="el-icon-menu"></i>
              <router-link to="/account/infoFocusPeo" class="router-link-active"><span>关注的人</span></router-link>
            </el-menu-item>
            <el-menu-item index="3">
              <i class="el-icon-document"></i>
              <router-link to="/account/infoCollection" class="router-link-active"><span>我的收藏</span></router-link>
            </el-menu-item>
            <el-menu-item index="4">
              <i class="el-icon-setting"></i>
              <router-link to="/account/infoDynamic" class="router-link-active"><span>我的动态</span></router-link>
            </el-menu-item>
            <el-menu-item index="5">
              <i class="el-icon-setting"></i>
              <router-link to="/account/infoLevel" class="router-link-active"><span>等级信息</span></router-link>
            </el-menu-item>
            <el-menu-item index="6">
              <i class="el-icon-setting"></i>
              <router-link to="/account/infoRecord" class="router-link-active"><span>我的战绩</span></router-link>
            </el-menu-item>
            <el-menu-item index="7">
              <i class="el-icon-setting"></i>
              <router-link to="/account/infoNotes" class="router-link-active"><span>个人笔记</span></router-link>
            </el-menu-item>
            <el-menu-item index="8">
              <i class="el-icon-setting"></i>
              <router-link to="/account/infoFoot" class="router-link-active"><span>我的足迹</span></router-link>
            </el-menu-item>
            <el-menu-item index="9">
              <i class="el-icon-setting"></i>
              <router-link to="/account/infoFeedback" class="router-link-active"><span>我要反馈</span></router-link>
            </el-menu-item>
            <el-menu-item index="10">
              <i class="el-icon-setting"></i>
              <router-link to="/account/infoSetting" class="router-link-active"><span>账号设置</span></router-link>
            </el-menu-item>
          </el-menu>
        </el-col>
        <el-col :span="18">
          <router-view></router-view>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Info',
  data () {
    return {
      tabPosition: 'left'
    }
  },
  methods: {
    updateInfo () {
      this.$router.push('/account/infoUpdate')
    },
    main () {
      this.$router.push('/account/main')
    }

  }
}
</script>

<style scoped>
  @import url("//unpkg.com/element-ui@2.13.0/lib/theme-chalk/index.css");
  .box-card {
    width: 65%;
    height:100px;
    margin: auto;
    margin-top: 30px;
  }
  .body{
    background-color:#f6f6f8;
  }
  .list{
    /*margin-left: 200px;
    margin-top: 20px;*/
    width: 65%;
    margin: auto;
    margin-top: 20px;
  }
  .router-link-active {
    text-decoration: none;
    color: #000000;
  }

</style>
