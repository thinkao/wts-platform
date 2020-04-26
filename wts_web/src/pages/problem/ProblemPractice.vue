<template>
  <div class="body">
    <div>
      <div class="picture">
        <img src="../../../static/img/temp1.jpg" style="margin: auto;width: 100%">
      </div>
      <div class="box" style="margin-top: 30px">
        <div class="box-one">
          <router-link to="/practice/basic"><el-button>计算机基础</el-button></router-link>
        </div>
        <div class="box-two">
          <router-link to="/practice/language"><el-button>编程语言</el-button></router-link>
        </div>
        <div class="box-three">
          <router-link to="/practice/algorithm"><el-button>算法</el-button></router-link>
        </div>
        <div class="box-four">
          <router-link to="/practice/database"><el-button>数据库</el-button></router-link>
        </div>
        <!--<div class="box-five">
          <router-link to="/practice/database"><el-button>逻辑</el-button></router-link>
        </div>-->
      </div>
    </div>
    <div class="content">
      <router-view></router-view>
    </div>
  </div>

</template>

<script>

export default {
  name: "ProblemPractice",
  data () {

  },
  methods: {

  },
  mounted() {
  }
}
</script>

<style scoped>
  @import url("//unpkg.com/element-ui@2.13.0/lib/theme-chalk/index.css");
  .body{
    background-color:#f6f6f8;
  }

  .body .content{
    width: 70%;
    margin: auto;
    margin-bottom: 30px;
  }

  .body .picture{
    margin: auto;
    width: 70%;
    margin-top: 30px;
  }

   .box{
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    width: 70%;
    margin: auto;
  }

  .box .box-one,.box .box-two,.box .box-three,.box .box-four,.box .box-five{
    flex-grow: 1;
    align-items: center;
  }


</style>
