<template >
  <div class="body">
    <div class="content">
      <div class="carousel">
        轮播图
      </div>
      <div class="news">
        <div class="news-first">
          news1
        </div>
        <div class="news-second">
          news2
        </div>
        <div class="news-third">
          news3
        </div>
      </div>
      <div class="dynamic">
        动态
      </div>
    </div>
    <div class="person-info">
      占位<br>占位<br>占位<br>占位<br>占位<br>占位<br>hahahah
    </div>
  </div>
</template>

<script>
export default {
  name: 'Main'
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
    border: 1px solid black;
    width: 70%;
  }
  .person-info{
    width: 20%;
    border: 1px solid black;
  }

  .content .carousel{

  }
  .content .news{
    display: flex;
    flex-direction: row;
    justify-content: space-around;
  }
  .content .dynamic{

  }

  .content .news .news-first,.content .news .news-second,.content .news .news-third{
    border: 1px solid black;
    flex-grow: 1;
  }
</style>
