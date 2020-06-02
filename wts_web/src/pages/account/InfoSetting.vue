<template>
<div>
  <el-card class="box-card">
    <div>
      <span>绑定手机号</span>
      <el-divider></el-divider>
      <el-form :label-position="labelPosition" label-width="80px" :model="ruleForm.account" class="form">
        <el-form-item label="原手机号">
          <el-input v-model="ruleForm.account.phone"></el-input>
        </el-form-item>
        <el-form-item label="新手机号">
          <el-input v-model="ruleForm.account.newPhone"></el-input>
        </el-form-item>
        <el-form-item label="验证码">
          <el-input v-model="ruleForm.account.code"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button style="width: 324px" type="success" plain>保存</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="margin-top">
      <span>绑定邮箱</span>
      <el-divider></el-divider>
      <el-form :label-position="labelPosition" label-width="80px" :model="ruleForm.email" class="form">
        <el-form-item label="新邮箱">
          <el-input v-model="ruleForm.email.email"></el-input>
        </el-form-item>
        <el-form-item label="验证码">
          <el-input v-model="ruleForm.email.code"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button style="width: 324px" type="success" plain>保存</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="margin-top">
      <span>设置密码</span>
      <el-divider></el-divider>
      <el-form :label-position="labelPosition" label-width="80px" :model="ruleForm.password" class="form">
        <el-form-item label="原密码">
          <el-input v-model="ruleForm.password.oldPass"></el-input>
        </el-form-item>
        <el-form-item label="新密码">
          <el-input v-model="ruleForm.password.newPass"></el-input>
        </el-form-item>
        <el-form-item label="确认密码">
          <el-input v-model="ruleForm.password.surePass"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button style="width: 324px" type="success" plain>保存</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="margin-top">
      <span>系统设置</span>
      <el-divider></el-divider>
          <span style="margin-left: 250px">是否接受通知</span>
          <el-checkbox v-model="checked">是</el-checkbox>
    </div>
  </el-card>
</div>
</template>

<script>
export default {
  name: 'InfoSetting',
  data () {
    return {
      checked: true,
      ruleForm: {
        account: {
          phone: '',
          newPhone: '',
          code: ''
        },
        email: {
          newEmail: '',
          code: ''
        },
        password: {
          oldPass: '',
          newPass: '',
          surePass: ''
        },
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
  .margin-top {
    margin-top: 40px;
  }
  .form {
    width: 60%;
    margin-left: 100px;
  }
</style>
