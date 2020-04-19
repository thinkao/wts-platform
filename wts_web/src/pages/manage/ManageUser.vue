<template>
    <div>
        <div v-if="showList===1">
            <div class="conditions">
                <el-form :inline="true" :model="ruleForm" class="demo-form-inline">
                    <el-form-item label="用户ID">
                        <el-input v-model="ruleForm.id"></el-input>
                    </el-form-item>
                    <el-form-item label="用户名" class="condition">
                        <el-input v-model="ruleForm.username"></el-input>
                    </el-form-item>
                    <el-form-item label="手机号" class="condition">
                        <el-input v-model="ruleForm.phone"></el-input>
                    </el-form-item>
                    <el-form-item label="邮箱" class="condition">
                        <el-input v-model="ruleForm.email"></el-input>
                    </el-form-item>
                    <el-form-item label="角色" class="condition">
                        <el-select v-model="ruleForm.userType">
                            <el-option label="管理员" value="admin"></el-option>
                            <el-option label="普通用户" value="normal"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item style="float: right">
                        <el-button type="primary" @click="select" style="width: 100px">查询</el-button>
                    </el-form-item>
                </el-form>
            </div>
            <div class="conditions">
                <el-table
                        :data="tableData"
                        stripe
                        style="width: 100%">
                    <el-table-column
                            prop="id"
                            label="用户ID"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            prop="username"
                            label="用户名"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            prop="phone"
                            label="手机号">
                    </el-table-column>
                    <el-table-column
                            prop="email"
                            label="邮箱">
                    </el-table-column>
                    <el-table-column
                            prop="userType"
                            label="角色">
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
        </div>
        <div class="layout-div" v-if="showList===2">

            <div class="layout-div form-data">
                <div slot="header" class="layout-div header-bar">
                    <div style="margin-top: 60px;text-align: right">
                        <el-button v-show="isUpdate" class="zdy-button bao_cun" @click="saveUpdate">保存</el-button>
                        <el-button @click="changeMap(1)" class="zdy-button">取消</el-button>
                    </div>
                </div>
                <el-form :model="ruleForm" ref="ruleForm" label-width="100px" style="margin-top: 15px">
                    <el-form-item label="用户ID" prop="id">
                        <el-input v-model="ruleForm.id"></el-input>
                    </el-form-item>
                    <el-form-item label="用户名" prop="username">
                        <el-input v-model="ruleForm.username"></el-input>
                    </el-form-item>
                    <el-form-item label="手机号" prop="phone">
                        <el-input v-model="ruleForm.phone" @blur="phoneValid($event)"></el-input>
                    </el-form-item>
                    <el-form-item label="邮箱" prop="email">
                        <el-input v-model="ruleForm.email"></el-input>
                    </el-form-item>
                    <el-form-item label="角色" prop="userType">
                        <el-select v-model="ruleForm.userType">
                            <el-option label="普通用户" value="1"></el-option>
                            <el-option label="管理员" value="2"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="宣言" prop="declaration">
                        <el-input v-model="ruleForm.declaration"></el-input>
                    </el-form-item>
                    <el-form-item label="等级" prop="level">
                        <el-input v-model="ruleForm.level"></el-input>
                    </el-form-item>
                    <el-form-item label="积分" prop="integral">
                        <el-input v-model="ruleForm.integral"></el-input>
                    </el-form-item>
                    <el-form-item label="上传头像" prop="avatar">
                        <el-upload
                                class="upload-demo"
                                v-model="ruleForm.avatar"
                                action="https://jsonplaceholder.typicode.com/posts/"
                                multiple
                                :file-list="fileList">
                            <el-button size="small" type="primary">点击上传</el-button>
                        </el-upload>
                    </el-form-item>
                </el-form>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "ManageUser",
    data() {
        return {
            isUpdate: true,
            showList: 1,
            tableData: [],
            fileList: [],
            sizeTotal: 0,
            ruleForm: {
                id: '',
                username: '',
                phone: '',
                email: '',
                userType: '',
                declaration: '',
                level: '',
                integral: '',
                avatar: ''
            }
        }
    },
    methods: {
        search () {
            this.$account.request("/api/UserAPI","","GET").then(resp => {
                this.tableData = resp.data.data
                this.sizeTotal = resp.data.data.total
            }).catch(function (error){
                console.log(error)
            })
        },
        saveUpdate(){
            if(!this.ruleForm.id){
                this.$account.request("/api/UserAPI","","POST").then(resp => {
                    if(resp.data.err == null){
                        this.$message({message: '新增成功', type: 'success'})
                        this.changeMap(1)
                    }
                    this.$message({message: '新增失败', type: 'fail'})
                })
            }else{
                this.$account.request("/api/UserAPI","","INPUT").then(resp => {
                    if(resp.data.err == null){
                        this.$message({type: 'success', message: '修改成功!', customClass: 'zZindex'});
                        this.changeMap(1)
                    }else {
                        this.$message({type: 'fail', message: '修改失败!', customClass: 'zZindex'});
                    }
                })
            }
        },
        changeMap (type) {
            if (this.showList === 2) {
                this.resetForm()
            }
            this.showList = type
            this.isUpdate = true
            if(type === 1){
                this.ruleForm = {
                    book_name: '',
                    book_editor: '',
                    book_price: '',
                    kind_name: '',
                    pub_name: '',
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
                var DeleteParams = {Id: row.id}
                this.$account.request("/api/UserAPI",DeleteParams,"DELETE").then(resp => {
                    console.log("----"+resp.data.err)
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
            var SelectByIdParams = {Id: row.id}
            this.$account.request("/api/UserAPI",SelectByIdParams,"GET").then(resp => {
                this.tableData = resp.data.data
            }).catch(function (error){
                console.log(error)
            })
        },

        detailRow(){

        },
        phoneValid(e){
            let boolean = new RegExp("^[1-9]\\d*.\\d*|0\\.\\d*[1-9]\\d*|[0-9]$").test(e.target.value)
            if(!boolean){
                this.$message.warning('请输入数字')
                e.target.value=''
            }
        },
    },
    mounted () {
        this.search()
    }
}
</script>

<style scoped>

    .conditions{
        margin: 40px;
    }

    .conditions .condition{
        margin-left: 13px;
    }

    .form-data{
        margin-left: 180px;
        width: 60%;
    }

</style>