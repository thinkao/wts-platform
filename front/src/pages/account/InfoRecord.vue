<template>
<div>
  <el-card class="box-card">
    <span>战绩排名</span>
    <div class="rank">
      <el-form ref="selectForm" :model="selectForm" label-width="80px" style="float:left">
        <el-form-item label="排名选择">
          <el-select v-model="selectForm.rankType" placeholder="请选择排名类型">
            <el-option label="胜率排名" value="odds"></el-option>
            <el-option label="战数排名" value="warNum"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <span>我的当前排名为{{rank}}</span>
    </div>
    <div>
      <el-table
        :data="tableData"
        stripe
        style="width: 100%">
        <el-table-column
          prop="name"
          label="用户名"
          width="180">
        </el-table-column>
        <el-table-column
          prop="declaration"
          label="当前等级等基本信息"
          width="200">
        </el-table-column>
        <el-table-column
          prop="declaration"
          label="综合指标"
          width="280">
        </el-table-column>
      </el-table>
    </div>
  </el-card>
  <el-card class="box-card">
    <span>历史战绩</span>

  </el-card>
  <el-card class="box-card">
    <span>等级勋章</span>

  </el-card>
</div>
</template>
<script>
export default {
  name: 'InfoRecord',
  data () {
      return {
        tableData:[],
        rank:'',
        selectForm: {
            rankType:''
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
  },
}
</script>

<style scoped>
  .box-card {
    width: 97%;
    margin-bottom: 20px;
    float: right;
  }
  .rank{
    margin-top: 20px;
    text-align: center;
    line-height: 40px;
  }
</style>
