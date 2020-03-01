<template>
<div>
  <el-card class="box-card">
    <el-tabs v-model="collectionList" @tab-click="handleClick" style="width: 100%;align-content: center" stretch="true">
      <el-tab-pane label="试题" name="first">
        <div>
          <i class="el-icon-tickets"></i>
          <span>收藏的试题:</span>
          <span>共{{allNum1}}条</span>
        </div>
      </el-tab-pane>
      <el-tab-pane label="文章/帖子" name="second">
        <div>
          <i class="el-icon-document"></i>
          <span>收藏的文章/帖子:</span>
          <span>共{{allNum2}}条</span>
        </div>
      </el-tab-pane>
    </el-tabs>
  </el-card>

</div>
</template>

<script>
export default {
  name: 'InfoCollection',
  data () {
    return {
      allNum1: '12',
      allNum2: '13',
      collectionList: 'first',
      ruleForm: {
        username: '',
        declaration: ''
      }
    }
  },
  methods: {
    submitForm (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          alert('submit!')
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style scoped>
  .box-card {
    width: 97%;
    margin: auto;
    float: right;
  }
</style>
