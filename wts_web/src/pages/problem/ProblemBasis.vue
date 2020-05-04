<template>
  <div>
    <el-card class="box-card">
      <i class="el-icon-caret-left">
        计算机基础
      </i>
      <el-card class="box-card-one">
        <div slot="header">
          <span>算法</span>
          <el-checkbox :indeterminate="isIndeterminateAlgorithms" v-model="checkAllAlgorithms" @change="handleCheckAlgorithmsAllChange" style="float: right; padding: 3px 0">全选</el-checkbox>
        </div>
        <el-checkbox-group v-model="checkedAlgorithms" @change="handleCheckedAlgorithmsChange">
          <el-checkbox v-for="algorithm in algorithms" :label="algorithm" :key="algorithm">{{algorithm}}</el-checkbox>
        </el-checkbox-group>
      </el-card>
      <el-card class="box-card-one">
        <div slot="header">
          <span>数据结构</span>
          <el-checkbox :indeterminate="isIndeterminateStructures" v-model="checkAllStructures" @change="handleCheckStructuresAllChange" style="float: right; padding: 3px 0">全选</el-checkbox>
        </div>
        <el-checkbox-group v-model="checkedStructures" @change="handleCheckedStructuresChange">
          <el-checkbox v-for="structure in structures" :label="structure" :key="structure">{{structure}}</el-checkbox>
        </el-checkbox-group>
      </el-card>
      <el-card class="box-card-one">
        <div slot="header">
          <span>计算机组成原理</span>
          <el-checkbox :indeterminate="isIndeterminateComponents" v-model="checkAllComponents" @change="handleCheckComponentsAllChange" style="float: right; padding: 3px 0">全选</el-checkbox>
        </div>
        <el-checkbox-group v-model="checkedComponents" @change="handleCheckedComponentsChange">
          <el-checkbox v-for="component in components" :label="component" :key="component">{{component}}</el-checkbox>
        </el-checkbox-group>
      </el-card>
      <el-card class="box-card-one">
        <div slot="header">
          <span>操作系统</span>
          <el-checkbox :indeterminate="isIndeterminateOperations" v-model="checkAllOperations" @change="handleCheckOperationsAllChange" style="float: right; padding: 3px 0">全选</el-checkbox>
        </div>
        <el-checkbox-group v-model="checkedOperations" @change="handleCheckedOperationsChange">
          <el-checkbox v-for="operation in operations" :label="operation" :key="operation">{{operation}}</el-checkbox>
        </el-checkbox-group>
      </el-card>
      <el-card class="box-card-one">
        <div slot="header">
          <span>计算机网络</span>
          <el-checkbox :indeterminate="isIndeterminateNetWorks" v-model="checkAllNetWorks" @change="handleCheckNetWorksAllChange" style="float: right; padding: 3px 0">全选</el-checkbox>
        </div>
        <el-checkbox-group v-model="checkedNetWorks" @change="handleCheckedNetWorksChange">
          <el-checkbox v-for="network in networks" :label="network" :key="network">{{network}}</el-checkbox>
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
  const algorithmsOptions = ['查找','排序','递归','复杂度'];
  const structureOptions =  ['数组','字符串','链表','栈','队列','树','哈希','堆','图'];
  const componentOptions =  ['编程基础','编译和原理'];
  const operationsOptions = ['操作系统','Linux']
  const NetWorksOptions = ['计算机网络']
    export default {
    name: "ProblemBasis",
    methods:{
      handleCheckAlgorithmsAllChange(val) {
        console.log("val:"+val)
        this.checkedAlgorithms = val ? algorithmsOptions : [];
        this.isIndeterminateAlgorithms = false;
      },
      handleCheckedAlgorithmsChange(value){
        this.AnswerAlgorithms = value
        let checkedCount = value.length;
        this.checkAll = checkedCount === this.algorithms.length;
        this.isIndeterminateAlgorithms = checkedCount > 0 && checkedCount < this.algorithms.length;
      },
      handleCheckStructuresAllChange(val) {
        console.log("val:"+val)
        this.checkedStructures = val ? structureOptions : [];
        this.isIndeterminateStructures = false;
      },
      handleCheckedStructuresChange(value) {
        this.AnswerStructures = value
        let checkedCount = value.length;
        this.checkAll = checkedCount === this.structures.length;
        this.isIndeterminateStructures = checkedCount > 0 && checkedCount < this.structures.length;
      },
      handleCheckComponentsAllChange(val) {
        console.log("val:"+val)
        this.checkedComponents = val ? componentOptions : [];
        this.isIndeterminateComponents = false;
      },
      handleCheckedComponentsChange(value) {
        this.AnswerComponents = value
        let checkedCount = value.length;
        this.checkAll = checkedCount === this.components.length;
        this.isIndeterminateComponents = checkedCount > 0 && checkedCount < this.components.length;
      },
      handleCheckOperationsAllChange(val) {
        console.log("val:"+val)
        this.checkedOperations = val ? operationsOptions : [];
        this.isIndeterminateOperations = false;
      },
      handleCheckedOperationsChange(value) {
        this.AnswerOperations = value
        let checkedCount = value.length;
        this.checkAll = checkedCount === this.operations.length;
        this.isIndeterminateOperations = checkedCount > 0 && checkedCount < this.operations.length;
      },
      handleCheckNetWorksAllChange(val) {
        console.log("val:"+val)
        this.checkedNetWorks = val ? NetWorksOptions : [];
        this.isIndeterminateNetWorks = false;
      },
      handleCheckedNetWorksChange(value) {
        this.AnswerNetWorks = value
        let checkedCount = value.length;
        this.checkAll = checkedCount === this.networks.length;
        this.isIndeterminateNetWorks = checkedCount > 0 && checkedCount < this.networks.length;
      },
      selectOptions(){
        this.$router.push('/answer')
        if (this.star == 0){
          this.$message({message: '请选择难度等级', type: 'fail'})
          return false
        }
        if (this.star < 3) {
          this.queryForm.Difficult = "difficultLevelOne"
        }else if(this.star <= 4 && this.star >= 3){
          this.queryForm.Difficult = "difficultLevelTwo"
        }else {
          this.queryForm.Difficult = "difficultLevelThree"
        }
        const selectProblemConditionParams = {
          Type : this.AnswerStructures.concat(this.AnswerNetWorks).concat(this.AnswerOperations).concat(this.checkedAlgorithms).concat(this.AnswerComponents),
          Difficult: this.queryForm.Difficult,
          Size: 10,
        };
        console.log("selectProblemConditionParams"+JSON.stringify(selectProblemConditionParams))
        this.$account.request("/api/ProblemAPI",selectProblemConditionParams,"GET").then(resp => {
          console.log("resp----->"+resp.data.data)
        }).catch(function (error){
          console.log(error)
        })
      }
    },
    data() {
        return {
          queryForm:{
            Type: [],
            Difficult: ""
          },
          Size: 0,
          AnswerAlgorithms:[],
          AnswerComponents:[],
          AnswerOperations:[],
          AnswerNetWorks:[],
          AnswerStructures:[],
          checkAll: false,
          checkAllAlgorithms: false,
          checkAllStructures: false,
          checkAllComponents: false,
          checkAllOperations: false,
          checkAllNetWorks: false,
          checkedAlgorithms:[],
          checkedStructures:[],
          checkedComponents:[],
          checkedOperations:[],
          checkedNetWorks:[],
          star: null,
          structures: structureOptions,
          algorithms: algorithmsOptions,
          components: componentOptions,
          operations: operationsOptions,
          networks: NetWorksOptions,
          isIndeterminateAlgorithms: true,
          isIndeterminateStructures: true,
          isIndeterminateComponents: true,
          isIndeterminateOperations: true,
          isIndeterminateNetWorks: true
        }
    },
    mounted() {
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
