<template>
  <div class="login-main">
    <img height="100%" src="../../static/img/login.png" style="position: fixed;top:0;left: 0" width="100%">
    <div class="login-main-middle">
      <el-form :model="ruleForm" class="demo-ruleForm form-input-login" ref="ruleForm" status-icon>
        <el-tabs @tab-click="handleClick" stretch="true" style="width: 100%;align-content: center" v-model="LoginName">
          <el-tab-pane label="账号登录" name="first">
            <img src="../../static/img/logo_green.png" style="width: 90px;height: 90px;margin-top: 30px">
            <el-form-item prop="phone">
              <el-input auto-complete="off" placeholder="手机号或邮箱" style="width: 350px;margin-top: 10px"
                        v-model="ruleForm.phoneEmail"></el-input>
            </el-form-item>
            <el-form-item prop="password">
              <el-input auto-complete="off" placeholder="密码" style="width: 350px" type="password"
                        v-model="ruleForm.password"></el-input>
            </el-form-item>
            <el-form-item prop="login">
              <el-button :loading="loginLoading" @click="onLogin" @keyup.enter.native="onLogin"
                         style="background: #5dd5c8;border: #f8fcff;width: 350px" type="primary">登录
              </el-button>
            </el-form-item>
            <el-form-item>
              <el-button style="width: 100%;text-decoration: underline;color: #5dd5c8 " type="text">立即注册</el-button>
            </el-form-item>
          </el-tab-pane>
          <el-tab-pane label="短信登录" name="second">
            <img src="../../static/img/logo_green.png" style="width: 90px;height: 90px;margin-top: 30px">
            <el-form-item prop="phone">
              <el-input auto-complete="off" placeholder="手机号" style="width: 248px;margin-top: 10px"
                        v-model="ruleForm.phone"></el-input>
              <el-button>发送验证</el-button>
            </el-form-item>
            <el-form-item prop="password">
              <el-input auto-complete="off" placeholder="验证码" style="width: 350px" type="password"
                        v-model="ruleForm.password"></el-input>
            </el-form-item>
            <el-form-item prop="login">
              <el-button :loading="loginLoading" @click="onLogin" @keyup.enter.native="onLogin"
                         style="background: #5dd5c8;border: #f8fcff;width: 350px" type="primary">登录
              </el-button>
            </el-form-item>
            <el-form-item>
              <el-button style="width: 100%;text-decoration: underline;color: #5dd5c8 " type="text">立即注册</el-button>
            </el-form-item>
          </el-tab-pane>
        </el-tabs>
      </el-form>
      <br/>
      <el-form :model="addForm" :rules="rules" class="demo-ruleForm form-input-register" ref="addForm" status-icon
               style="margin-bottom: 130px">
        <el-tabs @tab-click="handleClick" stretch="true" style="width: 100%;align-content: center"
                 v-model="RegisterName">
          <el-tab-pane label="手机号注册" name="first">
            <el-form-item prop="phone">
              <el-input auto-complete="off" placeholder="请输入您的手机号" style="width: 248px;margin-top: 50px" type="text"
                        v-model="addForm.phone"></el-input>
              <el-button>发送验证</el-button>
            </el-form-item>
            <el-form-item prop="validate">
              <el-input auto-complete="off" placeholder="请输入验证码" style="width: 350px;" type="password"
                        v-model="addForm.validate"></el-input>
            </el-form-item>
            <el-form-item prop="pass">
              <el-input auto-complete="off" placeholder="请设置新的密码" style="width: 350px;" type="password"
                        v-model="addForm.pass"></el-input>
            </el-form-item>
            <el-form-item prop="checkPass">
              <el-input auto-complete="off" placeholder="请再次确认密码" style="width: 350px;" type="password"
                        v-model="addForm.checkPass"></el-input>
            </el-form-item>
            <el-form-item prop="username">
              <el-input auto-complete="off" placeholder="请设置昵称" style="width: 350px;" type="username"
                        v-model="addForm.username"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button style="width: 100%;text-decoration: underline;color: #5dd5c8 " type="text">已有账号，直接登录
              </el-button>
            </el-form-item>
            <el-form-item prop="login">
              <el-button :loading="loginLoading" @click="onRegister('addForm')"
                         @keyup.enter.native="onRegister('addForm')"
                         style="background: #5dd5c8;border: #f8fcff;width: 350px" type="primary">注册
              </el-button>
            </el-form-item>
          </el-tab-pane>
        </el-tabs>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
    data() {
        var validatePass = (rule, value, callback) => {
            if (value === '') {
                callback(new Error('请输入密码'));
            } else {
                if (this.addForm.checkPass !== '') {
                    this.$refs.addForm.validateField('checkPass');
                }
                callback();
            }
        };
        var validatePass2 = (rule, value, callback) => {
            if (value === '') {
                callback(new Error('请再次输入密码'));
            } else if (value !== this.addForm.pass) {
                callback(new Error('两次输入密码不一致!'));
            } else {
                callback();
            }
        };
        return {
            loginLoading: false,
            LoginName: 'first',
            RegisterName: 'first',
            ruleForm: {
                phoneEmail: '',
                password: ''
            },
            addForm: {
                phone: '',
                pass: '',
                checkPass: '',
                username: '',
                validate: ''
            },
            rules: {
                pass: [
                    {validator: validatePass, trigger: 'blur'}
                ],
                checkPass: [
                    {validator: validatePass2, trigger: 'blur'}
                ]
            }
        }
    },
    methods: {
        onRegister(formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    var registerParams = {phone: this.addForm.phone, password: this.addForm.pass,username: this.addForm.username}
                    this.$account.register(registerParams).then(resp => {
                        console.log(resp)
                        if (resp.data.err == null) {
                            this.$message({message: '注册成功', type: 'success'})
                            this.$router.push('/login')
                        }
                    })
                } else {
                    console.log('error submit!!')
                    return false
                }
            })
        },
        onLogin() {
            sessionStorage.clear()
            if (this.ruleForm.phoneEmail.trim() === '') {
                this.$message({message: '请输入账号', type: 'error'})
            } else if (this.ruleForm.password.trim() === '') {
                this.$message({message: '请输入密码', type: 'error'})
            } else {
                var loginParams = {phoneEmail: this.ruleForm.phoneEmail, password: this.ruleForm.password}
                this.$account.request("/api/LoginAPI",loginParams,"POST").then(resp => {
                    console.log(resp)
                    if (resp.data.err == null) {
                        this.$message({message: '登陆成功', type: 'success'})
                        this.$router.push('/main')
                    }else {
                        this.$message({message: '用户名或者密码错误', type: 'error'})
                    }
                })
            }
        },
        handleClick () {

        }
    },
    mounted() {
        this.$account.request("/api/CSRFTokenAPI","","GET").then(resp => {
            console.log(resp)
            //global.Csrftoken = resp.data.data.Csrftoken
        })
    }
}
</script>
<style scoped>
  @import url("//unpkg.com/element-ui@2.13.0/lib/theme-chalk/index.css");

  .login {
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

  .form-input-login {
    width: 650px;
    height: 470px;
    margin-left: 450px;
    background: #f8fcff;
    border-radius: 10px;
  }

  .form-input-register {
    width: 650px;
    height: 560px;
    margin-left: 450px;
    background: #f8fcff;
    border-radius: 10px;
  }

</style>
