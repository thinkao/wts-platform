<template>
  <div>
    <el-card class="box-card">
      <i class="el-icon-caret-left">
        编程语言
      </i>
      <el-card class="box-card-one">
        <div slot="header">
          <span>热门语言</span>
          <el-checkbox :indeterminate="isIndeterminateHotLanguages" v-model="checkAllHotLanguages" @change="handleCheckHotLanguageAllChange" style="float: right; padding: 3px 0">全选</el-checkbox>
        </div>
        <el-checkbox-group v-model="checkedHotLanguages" @change="handleCheckedHotLanguageChange">
          <el-checkbox v-for="hotLanguage in hotLanguages" :label="hotLanguage" :key="hotLanguage">{{hotLanguage}}</el-checkbox>
        </el-checkbox-group>
      </el-card>
      <el-card class="box-card-one">
        <div slot="header">
          <span>非热门语言</span>
          <el-checkbox :indeterminate="isIndeterminateCoolLanguages" v-model="checkAllCoolLanguages" @change="handleCheckCoolLanguagesAllChange" style="float: right; padding: 3px 0">全选</el-checkbox>
        </div>
        <el-checkbox-group v-model="checkedCoolLanguages" @change="handleCheckedCoolLanguagesChange">
          <el-checkbox v-for="coolLanguage in coolLanguages" :label="coolLanguage" :key="coolLanguage">{{coolLanguage}}</el-checkbox>
        </el-checkbox-group>
      </el-card>
      <div class="box-card-one">
        <span class="demonstration" style="float: left;display: inline;margin-right: 30px">请选择难度等级</span>
        <el-rate v-model="star"></el-rate>
        <el-button type="success" style="float: right;display: inline;margin-bottom: 30px" @click="selectOptions">确认</el-button>
      </div>
    </el-card>
  </div>
</template>

<script>
    const hotLanguageOptions = ['Java','Go','Python','C','Html'];
    const coolLanguageOptions =  ['C#','PHP'];
    export default {
      name: "ProblemLanguage",
      methods:{
        handleCheckHotLanguageAllChange(val) {
          console.log("val:"+val)
          this.checkedHotLanguages = val ? hotLanguageOptions : [];
          this.isIndeterminateHotLanguages = false;
        },
        handleCheckedHotLanguageChange(value) {
          console.log("value"+value)
          let checkedCount = value.length;
          this.checkAll = checkedCount === this.hotLanguages.length;
          this.isIndeterminateHotLanguages = checkedCount > 0 && checkedCount < this.hotLanguages.length;
        },
        handleCheckCoolLanguagesAllChange(val) {
          console.log("val:"+val)
          this.checkedCoolLanguages = val ? coolLanguageOptions : [];
          this.isIndeterminateCoolLanguages = false;
        },
        handleCheckedCoolLanguagesChange(value) {
          console.log("value"+value)
          let checkedCount = value.length;
          this.checkAll = checkedCount === this.coolLanguages.length;
          this.isIndeterminateCoolLanguages = checkedCount > 0 && checkedCount < this.coolLanguages.length;
        },
        selectOptions(){
          this.$router.push('/answer')
          if (this.star == 0){
            this.$message({message: '请选择难度等级', type: 'fail'})
            return false
          }
          /*if (this.star < 3) {
            this.Difficult = "difficultLevelOne"
          }else if(this.star <= 4 && this.star >= 3){
            this.queryForm.Difficult = "difficultLevelTwo"
          }else {
            this.queryForm.Difficult = "difficultLevelThree"
          }
          const selectProblemConditionParams = {
            Type: this.Type,
            Difficult: this.queryForm.Difficult,
            Size: 3,
          };
          this.$account.request("/api/ProblemAPI",selectProblemConditionParams,"GET").then(resp => {
            console.log("->"+resp)
            console.log(resp.data)
            console.log(resp.data.data)
            this.Answer = resp.data.data
          }).catch(function (error){
            console.log(error)
          })*/
        }

      },
      data() {
        return{
          checkAll: false,
          checkAllHotLanguages: false,
          checkAllCoolLanguages: false,
          checkedHotLanguages:[],
          checkedCoolLanguages:[],
          star: null,
          hotLanguages: hotLanguageOptions,
          coolLanguages: coolLanguageOptions,
          isIndeterminateHotLanguages: true,
          isIndeterminateCoolLanguages: true
        }
      }
    }
</script>

<style scoped>
  .box-card {
    margin-top: 30px;
    width: 100%;
  }
  .box-card .box-card-one{
    width: 95%;
    margin: auto;
    margin-top: 30px;
  }
</style>
