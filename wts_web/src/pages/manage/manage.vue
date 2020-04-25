<template>
    <div>
        <el-row class="tac">
            <el-col :span="3">
                <el-menu
                        default-active="2"
                        class="el-menu-vertical-demo menu"
                        @open="handleOpen"
                        @close="handleClose"
                        background-color="#545c64"
                        text-color="#fff"
                        active-text-color="#ffd04b">
                    <el-submenu>
                        <template slot="title">
                            <h3>摘星管理控制台</h3>
                        </template>
                    </el-submenu>
                    <el-submenu index="1">
                        <template slot="title">
                            <i class="el-icon-menu"></i>
                            <span>用户管理</span>
                        </template>
                        <el-menu-item-group>
                            <el-menu-item index="1-1">
                                <i class="el-icon-location"></i>
                                <span slot="title"><router-link to="/manage/user" style="color: white;text-decoration-line: none">基本信息</router-link></span>
                            </el-menu-item>
                        </el-menu-item-group>
                    </el-submenu>
                    <el-submenu index="2">
                        <template slot="title">
                            <i class="el-icon-menu"></i>
                            <span>题库管理</span>
                        </template>
                        <el-menu-item-group>
                            <el-menu-item index="2-1">
                                <i class="el-icon-star-on"></i>
                                <span slot="title"><router-link to="/manage/problem" style="color: white;text-decoration-line: none">题库管理</router-link></span>
                            </el-menu-item>
                        </el-menu-item-group>
                    </el-submenu>
                </el-menu>
            </el-col>
            <el-col :span="20">
                <router-view></router-view>
            </el-col>
        </el-row>
    </div>
    
</template>

<script>
    export default {
        name: "manage"
    }
</script>

<style scoped>

    .menu{
        height: 760px;
    }

</style>