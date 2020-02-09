<template>
<div>
  <el-card class="box-card">
    <div>
      <span>我的笔记</span>
      <span style="float: right">共{{allNum}}条</span>
      <el-table
        :data="tableData"
        style="width: 100%">
        <el-table-column
          prop="date"
          width="180">
        </el-table-column>
        <el-table-column
          prop="name"
          width="180">
        </el-table-column>
        <el-table-column
          prop="address">
        </el-table-column>
        <el-table-column
          prop="date">
        </el-table-column>
      </el-table>

    </div>
  </el-card>

</div>
</template>

<script>
export default {
  name: 'InfoNotes',
  data () {
    return {
      allNum: '66',
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
