<template>
    <div>
        <div v-if="showList===1">
            <div class="conditions">
                <el-form :inline="true" v-model="queryForm" class="demo-form-inline">
                    <el-form-item label="题目ID">
                        <el-input v-model="queryForm.param.ID"></el-input>
                    </el-form-item>
                    <el-form-item label="题目名" class="condition">
                        <el-input v-model="queryForm.param.Content"></el-input>
                    </el-form-item>
                    <el-form-item label="类型" class="condition">
                        <el-input v-model="queryForm.param.Type"></el-input>
                    </el-form-item>
                    <el-form-item label="困难程度" class="condition">
                        <el-select v-model="queryForm.param.Difficult">
                            <el-option label="" value=""></el-option>
                            <el-option label="简单" value="difficultLevelOne"></el-option>
                            <el-option label="普通" value="difficultLevelTwo"></el-option>
                            <el-option label="困难" value="difficultLevelThree"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item style="float: right">
                        <el-button type="primary" @click="search" style="width: 100px">查询</el-button>
                    </el-form-item>
                </el-form>
            </div>
            <div>
                <el-button class="addProblem" type="primary" circle title="新增" @click="changeMap(2)">+</el-button>
            </div>
            <div class="conditions">
                <el-table
                        height="520"
                        :data="tableData"
                        stripe
                        style="width: 100%">
                    <el-table-column
                            prop="id"
                            label="题目ID"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            prop="content"
                            label="题目名"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            prop="option"
                            label="选项">
                    </el-table-column>
                    <el-table-column
                            prop="answer"
                            label="答案">
                    </el-table-column>
                    <el-table-column
                            prop="type"
                            label="类型">
                    </el-table-column>
                    <el-table-column
                            prop="difficult"
                            label="困难程度">
                    </el-table-column>
                    <el-table-column
                            label="操作">
                        <template slot-scope="scope">
                            <el-button @click="deleteRow(scope.row)" type="text" size="small"><i class="el-icon-delete" title="删除"></i></el-button>
                            <el-button @click="updateRow(scope.row)" type="text" size="small"><i class="el-icon-edit" title="编辑"></i></el-button>
                            <el-button @click="detailRow(scope.row)" type="text" size="small"><i class="el-icon-chat-dot-round" title="详细信息"></i></el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </div>
            <div class="block pageSelect">
                <el-pagination
                        background
                        @size-change="handleSizeChange"
                        @current-change="handleCurrentChange"
                        :current-page="1"
                        :page-sizes="[8, 15, 30]"
                        :page-size="2"
                        layout="total, sizes, prev, pager, next, jumper"
                        :total=sizeTotal>
                </el-pagination>
            </div>
        </div>
        <div class="layout-div" v-if="showList===2">

            <div class="layout-div form-data">
                <div slot="header" class="layout-div header-bar">
                    <div style="margin-top: 60px;text-align: right">
                        <el-button class="zdy-button" @click="saveUpdate">保存</el-button>
                        <el-button @click="changeMap(1)" class="zdy-button">取消</el-button>
                    </div>
                </div>
                <el-form :model="ruleForm" ref="ruleForm" label-width="100px" style="margin-top: 15px">
                    <el-form-item label="题目ID" prop="id">
                        <el-input :disabled="true" v-model="ruleForm.ID"></el-input>
                    </el-form-item>
                    <el-form-item label="题目名" prop="context">
                        <el-input v-model="ruleForm.Content"></el-input>
                    </el-form-item>
                    <el-form-item label="选项" prop="option">
                        <el-input v-model="ruleForm.Option"></el-input>
                    </el-form-item>
                    <el-form-item label="答案" prop="answer">
                        <el-input v-model="ruleForm.Answer"></el-input>
                    </el-form-item>
                    <el-form-item label="类型" prop="type">
                        <!--<el-input v-model="ruleForm.Type"></el-input>-->
                        <el-cascader style="width: 400px" :options="options" v-model="ruleForm.Type" clearable></el-cascader>
                    </el-form-item>
                    <el-form-item label="困难程度" prop="difficult">
                        <el-select v-model="ruleForm.Difficult">
                            <el-option label="简单" value="difficultLevelOne"></el-option>
                            <el-option label="普通" value="difficultLevelTwo"></el-option>
                            <el-option label="困难" value="difficultLevelThree"></el-option>
                        </el-select>
                    </el-form-item>
                </el-form>
            </div>
        </div>
        <el-dialog title="详细信息" :visible.sync="dialogVisible" width="40%">
            <el-form :data="this.openlaunch" label-width="90px" class="demo-ruleForm">
                <el-form-item label="题目ID:" prop="id">
                    <span>{{ruleForm.ID}}</span>
                </el-form-item>
                <el-form-item label="题目名:" prop="content">
                    <span>{{ruleForm.Content}}</span>
                </el-form-item>
                <el-form-item label="选项:" prop="option">
                    <span>{{ruleForm.Option}}</span>
                </el-form-item>
                <el-form-item label="答案:" prop="answer">
                    <span>{{ruleForm.Answer}}</span>
                </el-form-item>
                <el-form-item label="类型:" prop="type">
                    <span>{{ruleForm.Type}}</span>
                </el-form-item>
                <el-form-item label="困难程度:" prop="difficult">
                    <span>{{ruleForm.Difficult}}</span>
                </el-form-item>
            </el-form>
        </el-dialog>
    </div>
</template>

<script>
    export default {
        name: "ManageProblem",
        data() {
            return {
                showList: 1,
                tableData: [],
                openlaunch: [],
                fileList: [],
                dialogVisible: false,
                sizeTotal: 0,
                ruleForm: {
                    ID: '',
                    Content: '',
                    Option: '',
                    Answer: '',
                    Type: '',
                    Difficult: ''
                },
                queryForm: {
                    param:{
                        ID:'',
                        Content:'',
                        Type:'',
                        Difficult:''
                    },
                    currentPage: 1,
                    pageSize: 8
                },
                options: [{
                    value: '计算机基础',
                    label: '计算机基础',
                    children: [{
                        value: '算法',
                        label: '算法',
                        children: [{
                            value: '查找',
                            label: '查找'
                        }, {
                            value: '排序',
                            label: '排序'
                        }, {
                            value: '递归',
                            label: '递归'
                        }, {
                            value: '复杂度',
                            label: '复杂度'
                        }]
                    }, {
                        value: '数据结构',
                        label: '数据结构',
                        children: [{
                            value: '数组',
                            label: '数组'
                        }, {
                            value: '字符串',
                            label: '字符串'
                        },{
                            value: '链表',
                            label: '链表'
                        },{
                            value: '栈',
                            label: '栈'
                        },{
                            value: '队列',
                            label: '队列'
                        },{
                            value: '树',
                            label: '树'
                        },{
                            value: '哈希',
                            label: '哈希'
                        },{
                            value: '堆',
                            label: '堆'
                        },{
                            value: '图',
                            label: '图'
                        }]
                    },{
                        value: '计算机组成原理',
                        label: '计算机组成原理',
                        children: [{
                            value: '编程基础',
                            label: '编程基础'
                        }, {
                            value: '编译和原理',
                            label: '编译和原理'
                        }]
                    }, {
                        value: '操作系统',
                        label: '操作系统',
                        children: [{
                            value: '操作系统',
                            label: '操作系统'
                        }, {
                            value: 'Linux',
                            label: 'Linux'
                        }]
                    }, {
                        value: '计算机网络',
                        label: '计算机网络',
                        children: [{
                            value: '计算机网络',
                            label: '计算机网络'
                        }]
                    }]
                    }, {
                    value: '编程语言',
                    label: '编程语言',
                    children: [{
                        value: '热门语言',
                        label: '热门语言',
                        children: [{
                            value: 'Java',
                            label: 'Java'
                        }, {
                            value: 'Go',
                            label: 'Go'
                        }, {
                            value: 'Python',
                            label: 'Python'
                        }, {
                            value: 'C',
                            label: 'C'
                        }, {
                            value: 'Html',
                            label: 'Html'
                        }]
                    }, {
                        value: '非热门语言',
                        label: '非热门语言',
                        children: [{
                            value: 'C#',
                            label: 'C#'
                        }, {
                            value: 'PHP',
                            label: 'PHP'
                        }]
                    }]
                },{
                    value: '算法',
                    label: '算法',
                },{
                    value: 'SQL',
                    label: 'SQL',
                }]
            }
        },
        methods: {
            search () {
                this.$account.request("/api/ProblemCountAPI","","GET").then(resp => {
                    this.sizeTotal = resp.data.data
                }).catch(function (error){
                    console.log(error)
                })

                var selectConditionParams = {
                    ID: this.queryForm.param.ID,
                    Content: this.queryForm.param.Content,
                    Type: this.queryForm.param.Type,
                    Difficult: this.queryForm.param.Difficult,
                    CurrentPage: this.queryForm.currentPage,
                    PageSize: this.queryForm.pageSize,
                    //Offset: 0,
                }
                this.$account.request("/api/ProblemAPI",selectConditionParams,"GET").then(resp => {
                    this.tableData = resp.data.data
                }).catch(function (error){
                    console.log(error)
                })
            },
            saveUpdate(){
                var saveUpdateParams = {
                    ID: this.ruleForm.ID,
                    Content: this.ruleForm.Content,
                    Option:this.ruleForm.Option,
                    Answer: this.ruleForm.Answer,
                    Type:this.ruleForm.Type,
                    Difficult: this.ruleForm.Difficult,
                }
                if(!this.ruleForm.ID){
                    this.$account.request("/api/ProblemAPI",saveUpdateParams,"POST").then(resp => {
                        if(resp.data.err == null){
                            this.$message({message: '新增成功', type: 'success'})
                            this.search()
                            this.changeMap(1)
                        }else {
                            this.$message({message: '新增失败', type: 'fail'})
                        }
                    })
                }else{
                    console.log(saveUpdateParams)
                    this.$account.request("/api/ProblemAPI",saveUpdateParams,"PUT").then(resp => {
                        if(resp.data.err == null){
                            this.search()
                            this.$message({type: 'success', message: '修改成功!', customClass: 'zZindex'});
                            this.changeMap(1)
                        }else {
                            this.$message({type: 'fail', message: '修改失败!', customClass: 'zZindex'});
                        }
                    })
                }
            },
            handleSizeChange(val) {
                this.queryForm.pageSize = val
                this.search()
                console.log(`每页 ${val} 条`);
            },
            handleCurrentChange(val) {
                this.queryForm.currentPage = val
                this.search()
                console.log(`当前页: ${val}`);
            },
            changeMap (type) {
                if (this.showList === 2) {
                    this.resetForm()
                }
                this.showList = type
                if(type === 1){
                    this.ruleForm = {
                        ID: '',
                        Content: '',
                        Option: '',
                        Answer: '',
                        Type: '',
                        Difficult: ''
                    }
                }
            },
            resetForm () {
                this.$refs.ruleForm.resetFields()
            },
            select() {
                console.log('submit!');
            },
            deleteRow(row) {
                this.$confirm('是否确定删除?', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    var DeleteParams = {id: row.id}
                    this.$account.request("/api/ProblemAPI",DeleteParams,"DELETE").then(resp => {
                        if(resp.data.err == null){
                            this.search()
                            this.$message({type: 'success', message: '删除成功!'})
                        }else {
                            this.$message({type: 'fail', message: '删除失败!'})
                        }
                    })
                })
            },
            updateRow(row) {
                this.showList = 2
                var SelectByIdParams = {ID: row.id}
                this.$account.request("/api/ProblemAPI",SelectByIdParams,"GET").then(resp => {
                    this.resetForm()
                    if(resp.data.err == null){
                        this.ruleForm = resp.data.data
                    }
                }).catch(function (error){
                    console.log(error)
                })
            },

            detailRow(row){
                this.dialogVisible = true
                var SelectByIdParams = {ID: row.id}
                this.$account.request("/api/ProblemAPI",SelectByIdParams,"GET").then(resp => {
                    if(resp.data.err == null){
                        this.ruleForm = resp.data.data
                    }
                }).catch(function (error){
                    console.log(error)
                })
            },
        },
        mounted () {
            this.search()
        }
    }
</script>

<style scoped>

    .conditions{
        margin: 25px;
    }

    .conditions .condition{
        margin-left: 13px;
    }

    .form-data{
        margin-left: 180px;
        width: 60%;
    }

    .addProblem{
        margin-right: 70px;
        width: 50px;
        height: 50px;
        float: right;
        font-size: 25px;
    }

    .pageSelect{
        float: right;
        margin-right: 30px;
    }

</style>