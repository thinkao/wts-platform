<template >
  <div class="body">
    <div class="content">
      <div class="carousel">
        <el-carousel :interval="4000" type="card" height="260px">
          <el-carousel-item v-for="item in picture" :key="item">
            <img :src="item.img" style="width: 580px;height: 260px">
          </el-carousel-item>
        </el-carousel>
      </div>
      <div class="news">
        <div class="news-first">
            <img src="../../static/img/propa1.jpg" style="width: 100%;height: 120px">
            <p>wocacacacacac</p>
        </div>
        <div class="news-second">
            <div>
              <img src="../../static/img/propa2.jpg" style="width: 100%;height: 120px">
              <p>wocacacacacac</p>
            </div>
        </div>
        <div class="news-third">
            <div>
              <img src="../../static/img/propa5.jpeg" style="width: 100%;height: 120px">
              <p>wocacacacacac</p>
            </div>
        </div>
        <div class="news-third">
            <div>
              <img src="../../static/img/propa4.jpg" style="width: 100%;height: 120px">
              <p>wocacacacacac</p>
            </div>
          </div>
      </div>
      <div class="dynamic">
        <el-card>
          <el-tabs v-model="dynamic" type="card">
            <el-tab-pane label="当前动态" name="dynamicAll">
              <div style="margin: 10px">
                <ul>
                  <li v-for="item in list" :key="item">
                    <div>
                      <!--<img :src="item.ImgPath" style="float: left">-->
                      <el-avatar style="float: left"> user </el-avatar>
                      <p style="line-height: 40px;margin-left: 50px">{{item.Username}}</p>
                    </div>
                    <div class="dynamicContent">
                      <p>{{item.Content}}</p>
                      <img :src="item.ImgPath" style="margin-top: 10px">
                    </div>
                    <p style="float: right">{{item.CreateAt}}</p>
                  </li>
                </ul>
              </div>
            </el-tab-pane>
            <!--action="/api/DynamicAPI"-->
            <el-tab-pane label="发表动态" name="dynamicPush" ref="ruleForm" :model="ruleForm">
                <el-upload
                        multiple
                        name="file"
                        action=""
                        list-type="picture-card"
                        :auto-upload=false
                        :limit=1
                        :on-change="onchange"
                        :on-preview="handlePictureCardPreview"
                        accept="image/jpeg,image/gif,image/png,image/bmp"
                        :file-list="fileList2"
                        :on-remove="handleRemove">
                  <i class="el-icon-plus"></i>
                </el-upload>
                <el-dialog :visible.sync="dialogVisible">
                  <el-image width="100%" :src="dialogImageUrl" alt="" v-model="ruleForm.imgPath"></el-image>
                </el-dialog>
                <div style="margin: 20px 0;"></div>
                <el-input
                    type="textarea"
                    :autosize="{ minRows: 8, maxRows: 16}"
                    placeholder="请输入内容"
                    v-model="ruleForm.content"
                    maxlength="500"
                    show-word-limit
                >
                </el-input>
              <div style="margin-top: 20px">
                <el-button type="success" @click="publishDynamic">确认发表</el-button>
                <el-button type="info" @click="unPublishDynamic">取消发表</el-button>
              </div>
            </el-tab-pane>
          </el-tabs>
        </el-card>
      </div>
    </div>
    <div class="person-info">
      <el-card class="box-card1">
        <div>
          <i class="el-icon-money"></i>
          <span>这是一个新消息</span>
        </div>
      </el-card>
      <el-card class="box-card2">
        <div>
          <i class="el-icon-money"></i>
          <span>这是二个新消息</span>
        </div>
      </el-card>
      <el-card class="box-card3">
        <div>
          <i class="el-icon-s-data"></i>
          <span>今日推荐</span>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import global from "../global";
export default {
  name: 'Main',
  data () {
    return {
      isPush:true,
      list: [
        {
          name:"apple"
        },
        {
          name:"banana"
        }
      ],
      params2:{},
      fileList2:[],
      dialogVisible:false,
      dialogImageUrl: '',
      textarea: '',
      allNum: '66',
      ruleForm: {
        content: '',
        imgPath: ''
      },
      dynamic: '',
      tableData: [],
      picture: [
        {
          img: 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1555932101740&di=9ee42bcea75b9a6b91f15b6da964ac37&imgtype=0&src=http%3A%2F%2Fpic1.win4000.com%2Fwallpaper%2F9%2F5879c03369db1.jpg'
        },
        {
          img: 'https://th.bing.com/th/id/OIP.c70wvKqS4QjmOCEa7EzR6wHaEo?pid=Api&rs=1'
        },
        {
          img: 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1556526887&di=bfbad078380fffd1ec63abaf7d7dd163&imgtype=jpg&er=1&src=http%3A%2F%2Fpic1.win4000.com%2Fwallpaper%2Fd%2F573534849e26b.jpg'
        }
      ],
    }
  },
  methods: {
    search () {
        this.$account.request("/api/DynamicAPI","","GET").then(resp => {
          const jsonObj = JSON.parse(JSON.stringify(resp.data.data));
          for(let i = 0; i<jsonObj.length; i++){
            jsonObj[i].index = jsonObj[i].id;
          }
          this.list = jsonObj;
      }).catch(function (error){
        console.log(error)
      })
    },
    onchange(file, fileList){
      this.param2 = new FormData();
      this.param2.append('file', file.raw, file.name);
      console.log(fileList)
    },
    handleRemove (file, fileList) {
      console.log(file, fileList)
      this.param2 = null
    },
    handlePictureCardPreview (file) {
      this.dialogImageUrl = file.url
      this.dialogVisible = true
    },

    publishDynamic () {
      if(this.param2 == null&&this.ruleForm.content == ""){
        this.$message({message: '内容不能为空', type: 'error'})
      }else {
        let headers = {
          headers:{'Content-Type':'multipart/form-data', 'X-Csrftoken':global.CsrfToken}
        };

        axios.post('/api/DynamicAPI',this.param2,headers).then(resp=>{
          if(resp.data.data!=null){
            this.ruleForm.imgPath = resp.data.data.ImgPath
          }
          if (resp.data.err == null) {
            var pushParams = {Content: this.ruleForm.content,ImgPath:this.ruleForm.imgPath}
            this.$account.request('/api/DynamicAPI',pushParams,"POST").then(resp=>{
              if (resp.data.err != null) {
                this.isPush=false
              }else {
                this.$message({message: '发表成功', type: 'success'})
                this.search()
                this.dynamic = "dynamicAll"
              }
            })
          }
        })
      }
    },
    unPublishDynamic (){
      this.param2 = null
      this.ruleForm.content = ""
    }
  },
  mounted () {
    this.dynamic = "dynamicAll"
    this.search()
  }
}
</script>
<style scoped>
  .body{
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    background-color:#f6f6f8;
  }
  .content{
    /*border: 1px solid black;*/
    width: 70%;
  }
  .person-info{
    width: 20%;
    /*border: 1px solid black;*/
  }

  .person-info .box-card1{
    height: 200px;
  }
  .person-info .box-card2{
    height: 300px;
    margin-top: 20px;
  }
  .person-info .box-card3{
    height: 500px;
    margin-top: 20px;
  }

  .content .carousel{

  }
  .el-carousel__item h3 {
    color: #475669;
    font-size: 14px;
    opacity: 0.75;
    line-height: 200px;
    margin: 0;
  }

  .content .news{
    margin-top: 30px;
    display: flex;
    flex-direction: row;
    justify-content: space-around;
  }
  .content .dynamic{
    margin-top: 30px;
  }

  .content .news .news-first,.content .news .news-second,.content .news .news-third{
    flex-grow: 1;
    position: relative;
    margin-left: 3px;
  }

  

  .content .news .small-box-card{
    width: 330px;
    height:110px;
  }

  .dynamicContent{
    margin-top: 20px;
    margin-left: 50px;
  }
</style>
