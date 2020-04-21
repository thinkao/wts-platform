<template>
    <div>
        <div v-if="showList===1">
            <div class="conditions">
                <el-form :inline="true" v-model="queryForm" class="demo-form-inline">
                    <el-form-item label="用户ID">
                        <el-input v-model="queryForm.param.ID"></el-input>
                    </el-form-item>
                    <el-form-item label="用户名" class="condition">
                        <el-input v-model="queryForm.param.Username"></el-input>
                    </el-form-item>
                    <el-form-item label="手机号" class="condition">
                        <el-input v-model="queryForm.param.Phone"></el-input>
                    </el-form-item>
                    <el-form-item label="邮箱" class="condition">
                        <el-input v-model="queryForm.param.Email"></el-input>
                    </el-form-item>
                    <el-form-item label="角色" class="condition">
                        <el-select v-model="queryForm.param.UserType">
                            <el-option label="管理员" value="admin"></el-option>
                            <el-option label="普通用户" value="normal"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item style="float: right">
                        <el-button type="primary" @click="search" style="width: 100px">查询</el-button>
                    </el-form-item>
                </el-form>
            </div>
            <div>
                <el-button class="addUser" type="primary" circle title="新增" @click="changeMap(2)">+</el-button>
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
                    <el-form-item label="用户ID" prop="id">
                        <el-input :disabled="true" v-model="ruleForm.ID"></el-input>
                    </el-form-item>
                    <el-form-item label="用户名" prop="username">
                        <el-input v-model="ruleForm.Username"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="password">
                        <el-input v-model="ruleForm.Password"></el-input>
                    </el-form-item>
                    <el-form-item label="手机号" prop="phone">
                        <el-input v-model="ruleForm.Phone" @blur="phoneValid($event)"></el-input>
                    </el-form-item>
                    <el-form-item label="邮箱" prop="email">
                        <el-input v-model="ruleForm.Email"></el-input>
                    </el-form-item>
                    <el-form-item label="角色" prop="userType">
                        <el-select v-model="ruleForm.UserType">
                            <el-option label="普通用户" value="normal"></el-option>
                            <el-option label="管理员" value="admin"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="宣言" prop="declaration">
                        <el-input v-model="ruleForm.Declaration"></el-input>
                    </el-form-item>
                    <el-form-item label="等级" prop="level">
                        <el-input v-model="ruleForm.Level"></el-input>
                    </el-form-item>
                    <el-form-item label="积分" prop="integral">
                        <el-input v-model="ruleForm.Integral"></el-input>
                    </el-form-item>
                    <!--<el-form-item label="上传头像" prop="avatar">
                        <el-upload
                                class="upload-demo"
                                v-model="ruleForm.Avatar"
                                action="https://jsonplaceholder.typicode.com/posts/"
                                multiple
                                :file-list="fileList">
                            <el-button size="small" type="primary">点击上传</el-button>
                        </el-upload>
                    </el-form-item>-->
                </el-form>
            </div>
        </div>
        <el-dialog title="详细信息" :visible.sync="dialogVisible" width="40%">
            <el-form :data="this.openlaunch" label-width="90px" class="demo-ruleForm">
                <el-form-item label="用户ID:" prop="id">
                    <span>{{ruleForm.ID}}</span>
                </el-form-item>
                <el-form-item label="用户名:" prop="username">
                    <span>{{ruleForm.Username}}</span>
                </el-form-item>
                <el-form-item label="密码:" prop="password">
                    <span>{{ruleForm.Password}}</span>
                </el-form-item>
                <el-form-item label="手机号:" prop="phone">
                    <span>{{ruleForm.Phone}}</span>
                </el-form-item>
                <el-form-item label="邮箱:" prop="email">
                    <span>{{ruleForm.Email}}</span>
                </el-form-item>
                <el-form-item label="角色:" prop="userType">
                    <span>{{ruleForm.UserType}}</span>
                </el-form-item>
                <el-form-item label="宣言:" prop="declaration">
                    <span>{{ruleForm.Declaration}}</span>
                </el-form-item>
                <el-form-item label="等级:" prop="level">
                    <span>{{ruleForm.Level}}</span>
                </el-form-item>
                <el-form-item label="积分:" prop="integral">
                    <span>{{ruleForm.Integral}}</span>
                </el-form-item>
            </el-form>
        </el-dialog>
    </div>
</template>

<script>
export default {
    name: "ManageUser",
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
                Username: '',
                Password: '',
                Phone: '',
                Email: '',
                UserType: '',
                Declaration: '',
                Level: '',
                Integral: '',
                Avatar: ''
            },
            queryForm: {
                param:{
                    ID:'',
                    Username:'',
                    Phone:'',
                    Email:'',
                    UserType:''
                },
                currentPage: 1,
                pageSize: 8
            }
        }
    },
    methods: {
        search () {
            var selectConditionParams = {
                ID: this.queryForm.param.ID,
                Username: this.queryForm.param.Username,
                Phone: this.queryForm.param.Phone,
                Email: this.queryForm.param.Email,
                UserType: this.queryForm.param.UserType,
                CurrentPage: this.queryForm.currentPage,
                PageSize: this.queryForm.pageSize,
                //Offset: 0,
            }
            console.log(JSON.stringify(selectConditionParams))
            this.$account.request("/api/UserAPI",selectConditionParams,"GET").then(resp => {
                this.tableData = resp.data.data
                this.sizeTotal = resp.data.data.total
            }).catch(function (error){
                console.log(error)
            })
        },
        saveUpdate(){
            var saveUpdateParams = {
                ID: this.ruleForm.ID,
                UserName: this.ruleForm.Username,
                Phone:this.ruleForm.Phone,
                Password: this.ruleForm.Password,
                Email:this.ruleForm.Email,
                UserType: this.ruleForm.UserType,
                Declaration:this.ruleForm.Declaration,
                Level: this.ruleForm.Level,
                Integral: parseInt(this.ruleForm.Integral),
                Avatar: this.ruleForm.Avatar
            }
            if(!this.ruleForm.ID){
                this.$account.request("/api/UserAPI",saveUpdateParams,"POST").then(resp => {
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
                this.$account.request("/api/UserAPI",saveUpdateParams,"PUT").then(resp => {
                    if(resp.data.err == null){
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
                this.$account.request("/api/UserAPI",DeleteParams,"DELETE").then(resp => {
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
            this.$account.request("/api/UserAPI",SelectByIdParams,"GET").then(resp => {
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
            this.$account.request("/api/UserAPI",SelectByIdParams,"GET").then(resp => {
                if(resp.data.err == null){
                    this.ruleForm = resp.data.data
                }
            }).catch(function (error){
                console.log(error)
            })
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
        margin: 25px;
    }

    .conditions .condition{
        margin-left: 13px;
    }

    .form-data{
        margin-left: 180px;
        width: 60%;
    }

    .addUser{
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