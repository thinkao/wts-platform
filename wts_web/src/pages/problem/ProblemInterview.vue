<template>
  <div class="body">
    <div class="picture">
      <img src="../../../static/img/temp3.jpg" style="height:200px;width: 100%">
    </div>
    <div>
      <el-card class="box-card">
        <div class="search interval">
          <span style="font-family: 'Adobe 楷体 Std R'">关键词搜索</span>
          <el-input
            style="width: 50%"
            v-model="search"
            placeholder="输入关键字搜索">
          </el-input>
        </div>
        <div class="content interval">
          <el-tabs type="border-card">
            <el-tab-pane>
              <span slot="label"><i class="el-icon-date"></i> 公司面经</span>
              公司面经
            </el-tab-pane>
          </el-tabs>
          <el-tabs type="border-card">
            <el-tab-pane>
              <span slot="label"><i class="el-icon-date"></i> 面试宝典</span>
              面试宝典
            </el-tab-pane>
          </el-tabs>
          <el-tabs type="border-card">
            <el-tab-pane>
              <span slot="label"><i class="el-icon-date"></i> 公众号必备</span>
              公众号必备
            </el-tab-pane>
          </el-tabs>
          <el-tabs type="border-card">
            <el-tab-pane>
              <span slot="label"><i class="el-icon-date"></i>热门文章</span>
              热门文章
            </el-tab-pane>
          </el-tabs>
        </div>
      </el-card>
    </div>

  </div>
</template>

<script>
    export default {
        name: "ProblemInterview"
    }
</script>

<style scoped>
  .body{
    background-color:#f6f6f8;
  }
  .body .picture{
    margin: auto;
    width: 100%;
  }

  .body .box-card {
    width: 80%;
    margin: auto;
  }

  .body .search{
    text-align: center;
  }

  .body .content{
    width: 90%;
    margin: auto;
  }

  .body .interval{
    margin-top: 30px;
  }


</style>
