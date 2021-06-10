<template>
  <div class="home" id="home" name="home" style="  background-color:#f5f5f5  ">
    <!-- 轮播图 -->
    <div class="block"> 
      <el-carousel height="460px">
        <el-carousel-item >
          <img style="height:460px;width:1225px;" src="../assets/1.jpg"/>
        </el-carousel-item>
        <el-carousel-item >
          <img style="height:460px;width:1225px;" src="../assets/2.jpg"/>
        </el-carousel-item>
        <el-carousel-item >
          <img style="height:460px;width:1225px;" src="../assets/3.jpg"/>
        </el-carousel-item>
      </el-carousel>
    </div>
    <!-- 轮播图END -->
    <div class="main-box">
      <div class="main">
      <h2 style="color:grey;margin-top:20px">畅销产品</h2>
        <el-row gutter="20" style="margin-left:120px; margin-top:10px;">
          <el-col v-for="(item,i) in dataList.slice(0,4)" :key="i" span="5" style="height:200px" >
            <el-card style="height:200px">
              <p class="el-icon-goods" style="  margin-left:10px;margin-top:15px;color:grey">货品名称：{{item.name}}</p>
              <p class="el-icon-s-home" style="  margin-left:10px;margin-top:8px;color:grey">库存：{{item.amount}}</p>
              <p class="el-icon-s-shop" style="  margin-left:10px;margin-top:8px;color:grey">供应商：{{item.supplier}}</p>
              <p class="el-icon-phone" style="  margin-left:10px;margin-top:8px;color:grey">联系方式:</p>
               <p  style="  margin-left:30px;margin-top:3px;color:grey">{{item.tel}}</p>
            </el-card>
          </el-col>
        </el-row>
        <el-row gutter="20" style="margin-left:120px; margin-top:30px;">
          <el-col v-for="(item,i) in dataList.slice(4,7)" :key="i" span="5" style="height:200px">
            <el-card style="height:200px">
              <p class="el-icon-goods" style="  margin-left:10px;margin-top:15px;color:grey">货品名称：{{item.name}}</p>
              <p class="el-icon-s-home" style="  margin-left:10px;margin-top:8px;color:grey">库存：{{item.amount}}</p>
              <p class="el-icon-s-shop" style="  margin-left:10px;margin-top:8px;color:grey">供应商：{{item.supplier}}</p>
              <p class="el-icon-phone" style="  margin-left:10px;margin-top:8px;color:grey">联系方式:</p>
               <p  style="  margin-left:30px;margin-top:3px;color:grey">{{item.tel}}</p>
            </el-card>
          </el-col>
          <el-col span="5">
            <el-card  style="height:200px" >
              <p style="margin-top:60px;margin-left:20px;font-size:25px;color:grey" @click="gotoGoods()">>>更多好货</p>
            </el-card>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>
<script>
import router from '../router'
export default {
  data() {
    return {
      dataList:[{name:'钢材',amount:'10吨',supplier:'北京建材厂',tel:'010-1234567'},{name:'钢材',amount:'10吨',supplier:'北京建材厂',tel:'010-1234567'},{name:'钢材',amount:'10吨',supplier:'北京建材厂',tel:'010-1234567'},{name:'钢材',amount:'10吨',supplier:'北京建材厂',tel:'010-1234567'},
      {name:'钢材',amount:'10吨',supplier:'北京建材厂',tel:'010-1234567'},{name:'钢材',amount:'10吨',supplier:'北京建材厂',tel:'010-1234567'},{name:'钢材',amount:'10吨',supplier:'北京建材厂',tel:'010-1234567'},{name:'钢材',amount:'10吨',supplier:'北京建材厂',tel:'010-1234567'}]
    };
  },

  created() {
    this.getData()
  },
  methods: {
    gotoGoods(){
      router.push({ path: "/goods" });
    },
    getData(){
      //向后端请求数据接口
    }
  }
};
</script>
<style scoped>
@import "../assets/css/index.css";
.cardText{
  margin-left:10px;
  margin-top:3px;
}
</style>