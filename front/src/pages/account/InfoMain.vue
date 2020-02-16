<template>
  <el-row class="tac">
    <el-col :span="11" style="margin-left: 20px" >
    <el-card class="small-box-card">
      <div>
        <i class="el-icon-s-grid"></i>
        <span>信息完整度</span>
      </div>
      <el-progress type="circle" :percentage="25"></el-progress>
    </el-card>
    <el-card class="small-box-card">
      <div>
        <i class="el-icon-money"></i>
        <span>我的星币</span>
      </div>
    </el-card>
  </el-col>
    <el-col :span="11" style="float: left" >
      <el-card class="small-box-card" >
        <div>
          <i class="el-icon-star-off"></i>
          <span>我的星章</span>
        </div>
      </el-card>
      <el-card class="small-box-card">
        <div>
          <i class="el-icon-attract"></i>
          <span>个人技能</span>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: 'InfoMain'
}
</script>

<style scoped>
  .small-box-card{
    width: 330px;
    height:220px;
  }
</style>
