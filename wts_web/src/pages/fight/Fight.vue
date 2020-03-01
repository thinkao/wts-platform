<template>
  <div class="body">
    <el-card class="box-card">
      <div>
        <el-tabs v-model="fightType" type="card" @tab-click="handleClick">
          <el-tab-pane label="随机匹配" name="random">
            <el-form ref="form" :model="form" label-width="80px">
              <el-form-item label="题目类型">
                <el-select v-model="form.region" placeholder="请选择题目类型">
                  <el-option label="随机" value="random"></el-option>
                  <el-option label="选择题" value="choose"></el-option>
                  <el-option label="编程题" value="program"></el-option>
                </el-select>
              </el-form-item>
              <el-form-item label="难易程度">
                <el-rate style="display: inline" v-model="difficult"></el-rate>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="onSubmit">立即开始</el-button>
              </el-form-item>
          </el-form>
          </el-tab-pane>
          <el-tab-pane label="好友对战" name="friend">
            <el-form ref="form" :model="form" label-width="80px" style="width: 40%">
              <el-form-item label="选择好友">
                <el-input placeholder="请输入好友ID或昵称" v-model="form.name"></el-input>
              </el-form-item>
              <el-form-item label="题目类型">
                <el-select v-model="form.region" placeholder="请选择题目类型">
                  <el-option label="随机" value="random"></el-option>
                  <el-option label="选择题" value="choose"></el-option>
                  <el-option label="编程题" value="program"></el-option>
                </el-select>
              </el-form-item>
              <el-form-item label="难易程度">
                <el-rate style="display: inline" v-model="difficult"></el-rate>
              </el-form-item>
              <el-form-item label="设置时间">
                <el-input placeholder="请输入分钟" v-model="input2">
                  <template slot="append">分钟</template>
                </el-input>
              </el-form-item>

              <el-form-item>
                <el-button type="primary" @click="onSubmit">立即开始</el-button>
              </el-form-item>
            </el-form>
          </el-tab-pane>
        </el-tabs>
      </div>

    </el-card>

  </div>

</template>

<script>
    export default {
        name: "Fight",
        data () {
            return {
                difficult: null,
                fightType: 'random',
                form: {
                    name: '',
                    region: '',
                    date1: '',
                    date2: '',
                    delivery: false,
                    type: [],
                    resource: '',
                    desc: ''
                }
            };
        },
        methods: {
            handleClick(tab, event) {
                console.log(tab, event);
            }
        }
    }
</script>

<style scoped>
  .body{
    display: flex;
    background-color:#f6f6f8;
  }
  .body .box-card {
    margin: auto;
    margin-top: 30px;
    width: 60%;
    height: 600px;
  }

</style>
