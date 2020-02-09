<template>
<div>
  <el-card class="box-card">
    <div>
      <span>我的反馈</span>
      <span style="float: right">共{{allNum}}条</span>
      <el-upload
        style="margin-top: 20px"
        action="https://jsonplaceholder.typicode.com/posts/"
        list-type="picture-card"
        :on-preview="handlePictureCardPreview"
        :on-remove="handleRemove">
        <i class="el-icon-plus"></i>
      </el-upload>
      <el-dialog :visible.sync="dialogVisible">
        <img width="100%" :src="dialogImageUrl" alt="">
      </el-dialog>
      <div style="margin: 20px 0;"></div>
      <el-input
        type="textarea"
        :autosize="{ minRows: 4, maxRows: 8}"
        placeholder="请输入内容"
        v-model="textarea"
        maxlength="200"
        show-word-limit
      >
      </el-input>
      <el-button>确认反馈</el-button>
      <!--<el-button>立即反馈</el-button>-->
      <el-button>反馈记录</el-button>

    </div>
  </el-card>

</div>
</template>

<script>
export default {
  name: 'InfoFeedback',
  data () {
    return {
      dialogImageUrl: '',
      dialogVisible: false,
      text: '',
      textarea: '',
      allNum: '66',
      ruleForm: {
        username: '',
        declaration: ''
      }
    }
  },
  methods: {
    handleRemove (file, fileList) {
      console.log(file, fileList)
    },
    handlePictureCardPreview (file) {
      this.dialogImageUrl = file.url
      this.dialogVisible = true
    },
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
