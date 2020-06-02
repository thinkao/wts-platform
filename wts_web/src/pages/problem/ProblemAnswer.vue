<template>
    <div class="answerBody">
        <el-card class="answerBoxCard">
            <i class="el-icon-time">29:39</i>
            <div>
                <div style="text-align: center;margin-top: 40px">
                    <el-button type="success" style="width: 40px" circle>1</el-button>
                    <el-button style="margin-left: 25px;width: 40px" type="info" circle>2</el-button>
                    <el-button style="margin-left: 25px;width: 40px" type="info" circle>3</el-button>
                    <el-button style="margin-left: 25px;width: 40px" type="info" circle>4</el-button>
                    <el-button style="margin-left: 25px;width: 40px" type="info" circle>5</el-button>
                    <el-button style="margin-left: 25px;width: 40px" type="info" circle>6</el-button>
                    <el-button style="margin-left: 25px;width: 40px" type="info" circle>7</el-button>
                    <el-button style="margin-left: 25px;width: 40px" type="info" circle>8</el-button>
                    <el-button style="margin-left: 25px;width: 40px" type="info" circle>9</el-button>
                    <el-button style="margin-left: 25px;width: 40px" type="info" circle>10</el-button>
                </div>
                <div class="fighting">
                    <span style="font-size: 20px;">1.下列有关进程与线程描述正确的是？</span>
                    <br/>
                    <el-radio label="1" style="margin-top: 20px">进程是资源调度的基本单位</el-radio>
                    <br/>
                    <el-radio label="2" style="margin-top: 20px">线程是资源分配的基本单位</el-radio>
                    <br/>
                    <el-radio label="3" style="margin-top: 20px">一个进程最多可有一个线程</el-radio>
                    <br/>
                    <el-radio label="4" style="margin-top: 20px">一个进程可以有 n多个线程</el-radio>
                    <el-checkbox-group v-model="checkedOptions" @change="handleCheckedOptionChange">
                        <el-checkbox v-for="option in this.ruleForm.Option" :label="option" :key="option">{{option}}</el-checkbox>
                    </el-checkbox-group>
                    <el-button type="success" style="margin-top: 65px;width: 200px" @click="postProblem(this.ruleForm.ID,this.myAnswer)">提交</el-button>
                </div>
            </div>
        </el-card>
    </div>
</template>

<script>
    export default {
        name: "ProblemAnswer",
        methods:{
            handleCheckedOptionChange(value){
                this.myAnswer = value
                let checkedCount = value.length;
                this.isIndeterminateOptions = checkedCount > 0 && checkedCount < (this.ruleForm.Option).length;
            },
            /*提交*/
            postProblem(id,myAnswer) {
                const Params = {ID:id,Answer: myAnswer}
                this.$account.request("/api/FightAPI",Params,"POST").then(resp => {
                    if(resp.data.err == null){
                        if (resp.data.data === true){
                            /*回答正确*/
                            this.problemSuccess()
                        }else {
                            /*回答错误*/
                            this.problemFailed()
                        }
                    }else {
                        this.$message({message: '请联系管理员', type: 'fail'})
                    }
                })
            },
            /*获取题目信息*/
            nextProblem(id) {
                const Params = {ID: id};
                this.$account.request("/api/FightAPI",Params,"GET").then(resp => {
                    if(resp.data.err == null){
                        this.ruleForm.ID = resp.data.data.ID
                        this.ruleForm.Content = resp.data.data.Content
                        this.ruleForm.Option = resp.data.data.Option
                        console.log("====="+this.ruleForm.Content)
                        console.log("====="+this.ruleForm.Option)
                    }else {
                        this.$message({message: '请联系管理员', type: 'fail'})
                    }
                })
            },
            /*答题成功*/
            problemSuccess() {
                this.$alert('回答正确', {
                    confirmButtonText: '下一题',
                    /*callback: action => {
                        this.nextProblem()
                    }*/
                });
            },
            /*答题失败*/
            problemFailed() {
                this.$alert('回答错误', {
                    /*confirmButtonText: '重新作答',
                    confirmButtonText: '下一题',
                    callback: action => {
                        this.nextProblem()
                    }*/
                });
            }

        },
        data() {
            return {
                problemIds:[],
                checkedOptions:[],
                isIndeterminateOptions: true,
                Answers:['A:2','B:3','C:4'],
                test:'test',
                myAnswer:'',
                ruleForm:{
                    ID: '',
                    Content: '',
                    Option: [],
                    Answer: '',
                    Type: '',
                    Difficult: ''
                }
            }
        },
        created() {
            let proId = this.$route.query.ProblemIds;
            for(let i=0;i<proId.length;i++){
                this.problemIds.push(proId[i].id)
            }
            //this.nextProblem(this.problemIds[0])
        }
    }
</script>

<style scoped>

    .answerBoxCard{
        width: 70%;
        margin: auto;
        height: 655px;
        margin-top: 50px;
        margin-bottom: 50px;
    }

    .fighting{
        margin-top: 120px;
        text-align: center;
    }

    .answerBody{
        background-color:#f6f6f8;
    }
</style>