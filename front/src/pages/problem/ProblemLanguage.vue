<template>
  <div>
    <el-card class="box-card">
      <i class="el-icon-caret-left">
        编程语言
      </i>
      <el-card class="box-card-one">
        <div slot="header" class="clearfix">
          <span>编程语言</span>
          <el-checkbox style="float: right; padding: 3px 0">全选</el-checkbox>
        </div>
        <div v-for="o in 4" :key="o" class="text item">
          {{'列表内容 ' + o }}
        </div>
      </el-card>
      <el-card class="box-card-one">
        <div slot="header" class="clearfix">
          <span>脚本语言</span>
          <el-checkbox style="float: right; padding: 3px 0">全选</el-checkbox>
        </div>
        <div v-for="o in 4" :key="o" class="text item">
          {{'列表内容 ' + o }}
        </div>
      </el-card>
      <div class="box-card-one">
        <span class="demonstration" style="float: left;display: inline;margin-right: 30px">请选择难度等级</span>
        <el-rate v-model="star"></el-rate>
        <el-button type="success" style="float: right;display: inline;margin-bottom: 30px">确认</el-button>
      </div>
    </el-card>
  </div>
</template>

<script>
    export default {
        name: "ProblemLanguage"
    }
</script>

<style scoped>
  .box-card {
    margin-top: 30px;
    width: 100%;
  }
  .box-card .box-card-one{
    width: 95%;
    margin: auto;
    margin-top: 30px;
  }
</style>
