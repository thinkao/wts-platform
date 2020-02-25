<template>
  <div>
    <el-card class="box-card">
      <el-tabs v-model="focusList" @tab-click="handleClick" style="width: 100%;align-content: center" stretch="true">
        <el-tab-pane label="我的关注" name="first">
          <div>
            <el-table
              :data="tableData"
              stripe
              style="width: 100%">
              <el-table-column
                prop="name"
                label="用户名"
                width="180">
                <template slot-scope="scope">
                  <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
                  <span style="margin-left: 10px">{{ scope.row.name }}</span>
                </template>
              </el-table-column>
              <el-table-column
                prop="declaration"
                label="宣言"
                width="280">
              </el-table-column>
              <el-table-column
                align="right">
                <template slot="header" slot-scope="scope">
                  <el-input
                    v-model="search"
                    size="mini"
                    placeholder="输入关键字搜索"/>
                </template>
                <template slot-scope="scope">
                  <el-button
                    size="mini"
                    type="danger"
                    @click="handleDelete(scope.$index, scope.row)">关注</el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-tab-pane>
        <el-tab-pane label="我的粉丝" name="second">
          <div>
            <el-table
              :data="tableData1"
              stripe
              style="width: 100%">
              <el-table-column
                prop="name"
                label="用户名"
                width="180">
                <template slot-scope="scope">
                  <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
                  <span style="margin-left: 10px">{{ scope.row.name }}</span>
                </template>
              </el-table-column>
              <el-table-column
                prop="declaration"
                label="宣言"
                width="280">
              </el-table-column>
              <el-table-column
                align="right">
                <template slot="header" slot-scope="scope">
                  <el-input
                    v-model="search"
                    size="mini"
                    placeholder="输入关键字搜索"/>
                </template>
                <template slot-scope="scope">
                  <el-button
                    size="mini"
                    type="danger"
                    @click="handleDelete(scope.$index, scope.row)">关注</el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-tab-pane>
      </el-tabs>
    </el-card>

  </div>
</template>

<script>
export default {
  name: 'InfoFocusPeo',
  data () {
    return {
      focusList: 'first',
      tableData: [{
        name: '王小虎',
        declaration: '我没个性，所以不宣言！！！'
      }, {
        name: '王小虎',
        declaration: '我没个性，所以不宣言！！！'
      }, {
        name: '王小虎',
        declaration: '我没个性，所以不宣言！！！'
      }, {
        name: '王小虎',
        declaration: '我没个性，所以不宣言！！！'
      }, {
        name: '王小虎',
        declaration: '我没个性，所以不宣言！！！'
      }, {
        name: '王小虎',
        declaration: '我没个性，所以不宣言！！！'
      }, {
        name: '王小虎',
        declaration: '我没个性，所以不宣言！！！'
      }, {
        name: '王小虎',
        declaration: '我没个性，所以不宣言！！！'
      }, {
        name: '王小虎',
        declaration: '我没个性，所以不宣言！！！'
      }, {
        name: '王小虎',
        declaration: '我没个性，所以不宣言！！！'
      }],
      tableData1: [{
        name: '李小龙',
        declaration: '我没个性，所以不宣言！！！'
      }, {
        name: '李小龙',
        declaration: '我没个性，所以不宣言！！！'
      }, {
        name: '李小龙',
        declaration: '我没个性，所以不宣言！！！'
      }]

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
