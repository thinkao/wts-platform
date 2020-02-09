<template>
  <div class="login-main">
    <img src="/static/img/login.png" width="100%" height="100%" style="position: fixed;top:0;left: 0" >
    <div class="login-main-middle">
      <el-form :model="ruleForm" status-icon ref="ruleForm" :rules="rules" class="demo-ruleForm form-input">
        <el-tabs v-model="LoginName" @tab-click="handleClick" style="width: 100%;align-content: center" stretch="true">
          <el-tab-pane label="账号登录" name="first">
            <img src="/static/img/logo.png" style="width: 90px;height: 90px;margin-top: 30px">
            <el-form-item prop="phone"><el-input style="width: 350px;margin-top: 10px" v-model="ruleForm.phone" auto-complete="off"
                      placeholder="手机号或邮箱"></el-input>
            </el-form-item>
            <el-form-item prop="password"><el-input style="width: 350px" type="password" v-model="ruleForm.password" auto-complete="off"
                      placeholder="密码"></el-input></el-form-item>
            <el-form-item prop="login"><el-button type="primary" style="background: #5dd5c8;border: #f8fcff;width: 350px" @click="onLogin"
                       :loading="loginLoading" @keyup.enter.native="onLogin">登录</el-button></el-form-item>
            <el-form-item>
            <span style="color: #5dd5c8">立即注册</span>
            </el-form-item>
          </el-tab-pane>
          <el-tab-pane label="短信登录" name="second">
            <img src="/static/img/logo.png" style="width: 90px;height: 90px;margin-top: 30px">
            <el-form-item prop="phone"><el-input style="width: 350px;margin-top: 10px" v-model="ruleForm.phone" auto-complete="off"
                      placeholder="手机号或邮箱"></el-input>
            </el-form-item>
            <el-form-item prop="password"><el-input style="width: 350px" type="password" v-model="ruleForm.password" auto-complete="off"
                        placeholder="密码"></el-input>
            </el-form-item>
            <el-form-item prop="login"><el-button type="primary" style="background: #5dd5c8;border: #f8fcff;width: 350px" @click="onLogin"
                         :loading="loginLoading" @keyup.enter.native="onLogin">登录</el-button>
            </el-form-item>
            <el-form-item>
              <span style="color: #5dd5c8">立即注册</span>
            </el-form-item></el-tab-pane>
        </el-tabs>
      </el-form>
      <br/>
      <el-form :model="ruleForm" status-icon ref="ruleForm" class="demo-ruleForm form-input" style="margin-bottom: 130px">
        <el-tabs v-model="RegisterName" @tab-click="handleClick" style="width: 100%;align-content: center" stretch="true">
          <el-tab-pane label="手机号注册" name="first">
            <el-form-item prop="phone">
              <el-input style="width: 350px;margin-top: 50px" type="text"  v-model="addForm.phone" auto-complete="off"
                        placeholder="请输入您的手机号"></el-input>
            </el-form-item>
            <el-form-item prop="password">
              <el-input style="width: 350px;" type="password" v-model="addForm.password" auto-complete="off"
                        placeholder="请设置新的密码"></el-input>
            </el-form-item>
            <el-form-item prop="passwordS">
              <el-input style="width: 350px;" type="password" v-model="addForm.passwordS" auto-complete="off"
                        placeholder="请再次确认密码"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="text" style="width: 100%;text-decoration: underline;color: #5dd5c8 ">已有账号，直接登录</el-button>
            </el-form-item>
            <el-form-item prop="login">
              <el-button type="primary" style="background: #5dd5c8;border: #f8fcff;width: 350px" @click="onRegister"
                         :loading="loginLoading" @keyup.enter.native="onRegister">注册</el-button>
            </el-form-item>
          </el-tab-pane>
          <el-tab-pane label="邮箱注册" name="second">
            <el-form-item prop="phone">
              <el-input style="width: 350px;margin-top: 50px" type="text"  v-model="addForm.phone" auto-complete="off"
                        placeholder="请输入您的手机号"></el-input>
            </el-form-item>
            <el-form-item prop="password">
              <el-input style="width: 350px;" type="password" v-model="addForm.password" auto-complete="off"
                        placeholder="请设置新的密码"></el-input>
            </el-form-item>
            <el-form-item prop="passwordS">
              <el-input style="width: 350px;" type="password" v-model="addForm.passwordS" auto-complete="off"
                        placeholder="请再次确认密码"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="text" style="width: 100%;text-decoration: underline;color: #5dd5c8 ">已有账号，直接登录</el-button>
            </el-form-item>
            <el-form-item prop="login">
              <el-button type="primary" style="background: #5dd5c8;border: #f8fcff;width: 350px" @click="onRegister"
                         :loading="loginLoading" @keyup.enter.native="onRegister">注册</el-button>
            </el-form-item>
          </el-tab-pane>
        </el-tabs>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    var validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'))
      } else {
        if (this.addForm.passwordS !== '') {
          this.$refs.addForm.validateField('passwordS')
        }
        callback()
      }
    }
    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.addForm.password) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    }
    return {
      loginLoading: false,
      LoginName: 'first',
      RegisterName: 'first',
      ruleForm: {
        phone: '',
        password: ''
      },
      addForm: {
        phone: '',
        password: '',
        passwordS: ''
      },
      rules: {
        password: [{validator: validatePass, trigger: 'blur'}],
        passwordS: [{validator: validatePass2, trigger: 'blur'}]
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
    },
    onLogin () {
      this.$router.push('/main')
      sessionStorage.clear()
      if (this.ruleForm.phone.trim() === '') {
        this.$message({message: '请输入账号', type: 'error'})
      } else if (this.ruleForm.password.trim() === '') {
        this.$message({message: '请输入密码', type: 'error'})
      } else {
        var loginParams = {phone: this.ruleForm.phone, user_pass: this.ruleForm.password}
        this.$account.login(loginParams).then(resp => {
          console.log(resp)
        })
      }
    },
    onRegister () {
      if (this.ruleForm.phone.trim() === '') {
        this.$message({message: '请注册您的账号', type: 'error'})
      } else if (this.ruleForm.password.trim() === '') {
        this.$message({message: '请设置新的密码', type: 'error'})
      } else if (this.ruleForm.password !== this.ruleForm.passwordS) {
        this.$message({message: '两次密码不一致', type: 'error'})
      } else {
        var registerParams = {phone: this.ruleForm.phone, password: this.ruleForm.password}
        this.$account.register(registerParams).then(resp => {
          console.log(resp)
        })
      }
    }
  },
  mounted () {
  }
}
</script>
<style scoped>
  @import url("//unpkg.com/element-ui@2.13.0/lib/theme-chalk/index.css");
  .login{
    max-height: 750px;
    overflow: auto;
    position: relative;
  }
  .login-main {
    width: 100%;
    height: 100%;
    overflow: hidden;
  }
  .login-main-middle {
    position: relative;
    min-width: 360px;
    min-height: 320px;
    text-align: center;
    margin-top: 130px;
  }
  .form-input{
    width: 650px;
    height: 450px;
    margin-left: 450px;
    background: #f8fcff;
    border-radius: 10px;
  }

</style>
