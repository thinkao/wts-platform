<template>
  <el-row class="tac">
    <el-col :span="7" style="margin-left: 20px" >
    <el-card class="small-box-card" ></el-card>
    <el-card class="small-box-card"></el-card>
  </el-col>
    <el-col :span="7" style="float: left" >
      <el-card class="small-box-card" ></el-card>
      <el-card class="small-box-card"></el-card>
    </el-col>
    <el-col :span="7" style="float: left" >
      <el-card class="small-box-card" ></el-card>
      <el-card class="small-box-card"></el-card>
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
    width: 200px;
    height:180px;
  }
</style>
