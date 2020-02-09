<template>
<div>
  <el-card class="box-card">
    <div style="text-align: center">
      <el-upload><el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" style="width: 60px;height: 60px;"></el-avatar></el-upload>
    </div>
    <div style="margin-top: 20px">
      <el-form label-width="100px" class="demo-ruleForm">
        <el-form-item
          label="昵称"
          prop="age"
          :rules="[
      { required: true, message: '昵称不能为空'},
      ]"
        >
          <el-input type="age" v-model="ruleForm.username" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item
          label="个人宣言"
          prop="declaration">
          <el-input type="age" v-model="ruleForm.declaration" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item style="float: left">
          <el-button type="info">取消修改</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="success">保存修改</el-button>
        </el-form-item>
      </el-form>
    </div>
  </el-card>

</div>
</template>

<script>
export default {
  name: 'InfoUpdate',
  data () {
    return {
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
    height:600px;
    margin: auto;
    float: right;
  }
</style>
