<template>
    <div class="answerBody">
        <el-card class="answerBoxCard">
            <span>这是答题页面</span>
            <span>{{problem.Content}}</span>
        </el-card>
    </div>
</template>

<script>
    export default {
        name: "ProblemAnswer",
        methods:{
            nextProblem(id) {
                const Params = {ID: id};
                this.$account.request("/api/FightAPI",Params,"GET").then(resp => {
                    if(resp.data.err == null){
                        this.problem = resp.data.data
                    }else {
                        this.$message({message: '请联系管理员', type: 'fail'})
                    }
                })
            }

        },
        data() {
            return {
                problemIds:[],
                problem:"",
                ruleForm: {
                    ID: '',
                    Content: '',
                    Option: '',
                    Answer: '',
                    Type: '',
                    Difficult: ''
                },
            }
        },
        created() {
            let proId = this.$route.query.ProblemIds;
            for(let i=0;i<proId.length;i++){
                this.problemIds.push(proId[i].id)
            }
            this.nextProblem(this.problemIds[0])
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

    .answerBody{
        background-color:#f6f6f8;
    }
</style>