<template>
  <div>
    <el-card class="box-card">
      <i class="el-icon-caret-left">
        计算机基础
      </i>
      <el-card class="box-card-one">
        <div slot="header" class="clearfix">
          <span>算法</span>
          <el-checkbox style="float: right; padding: 3px 0">全选</el-checkbox>
        </div>
        <div v-for="o in 4" :key="o" class="text item">
          {{'列表内容 ' + o }}
        </div>
      </el-card>
      <el-card class="box-card-one">
        <div slot="header" class="clearfix">
          <span>数据结构</span>
          <el-checkbox style="float: right; padding: 3px 0">全选</el-checkbox>
        </div>
        <div v-for="o in 4" :key="o" class="text item">
          {{'列表内容 ' + o }}
        </div>
      </el-card>
      <el-card class="box-card-one">
        <div slot="header" class="clearfix">
          <span>计算机组成原理</span>
          <el-checkbox style="float: right; padding: 3px 0">全选</el-checkbox>
        </div>
        <div v-for="o in 4" :key="o" class="text item">
          {{'列表内容 ' + o }}
        </div>
      </el-card>
      <el-card class="box-card-one">
        <div slot="header" class="clearfix">
          <span>操作系统</span>
          <el-checkbox style="float: right; padding: 3px 0">全选</el-checkbox>
        </div>
        <div v-for="o in 4" :key="o" class="text item">
          {{'列表内容 ' + o }}
        </div>
      </el-card>
      <el-card class="box-card-one">
        <div slot="header" class="clearfix">
          <span>计算机网络</span>
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
        name: "ProblemBasis",
        data() {
            return {
                star: null,
            }
        }
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
  /*.clearfix:before,
  .clearfix:after {
    display: table;
    content: "";
  }
  .clearfix:after {
    clear: both
  }*/



</style>
