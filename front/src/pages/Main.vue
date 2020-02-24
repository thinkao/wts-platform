<template >
  <div class="body">
    <div class="content">
      <div class="carousel">
        <el-carousel :interval="4000" type="card" height="260px">
          <el-carousel-item v-for="item in picture" :key="item">
            <img :src="item.img" style="width: 540px;height: 260px">
          </el-carousel-item>
        </el-carousel>
      </div>
      <div class="news">
        <div class="news-first">
          <el-card class="small-box-card">
            <div>
              <i class="el-icon-money"></i>
              <span>这是一个新消息</span>
            </div>
          </el-card>
        </div>
        <div class="news-second">
          <el-card class="small-box-card">
            <div>
              <i class="el-icon-money"></i>
              <span>这是一个新消息</span>
            </div>
          </el-card>
        </div>
        <div class="news-third">
          <el-card class="small-box-card">
            <div>
              <i class="el-icon-money"></i>
              <span>这是一个新消息</span>
            </div>
          </el-card>
        </div>
      </div>
      <div class="dynamic">
        动态
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
          <i class="el-icon-money"></i>
          <span>这是三个新消息</span>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Main',
    data(){
      return {
          tableData: [],
          picture: [
              {
                  img: 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1555932101740&di=9ee42bcea75b9a6b91f15b6da964ac37&imgtype=0&src=http%3A%2F%2Fpic1.win4000.com%2Fwallpaper%2F9%2F5879c03369db1.jpg'
              },
              {
                  img: 'http://img2.3lian.com/2014/f4/191/d/22.jpg'
              },
              {
                  img: 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1556526887&di=bfbad078380fffd1ec63abaf7d7dd163&imgtype=jpg&er=1&src=http%3A%2F%2Fpic1.win4000.com%2Fwallpaper%2Fd%2F573534849e26b.jpg'
              }
          ],
      }
    },
    methods:{
        search () {
            this.$account.DynamicAll().then(resp => {
                this.tableData = resp.data.data
                console.log("resppppp"+resp.data.data)
            }).catch(function (error){
                console.log(error)
            })
        },
    },
    mounted() {
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

  }

  .content .news .news-first,.content .news .news-second,.content .news .news-third{
    /*border: 1px solid black;*/
    flex-grow: 1;
  }

  .content .news .small-box-card{
    width: 340px;
    height:110px;
  }
</style>
