<!--
 * @Description: 全部商品页面组件(包括全部商品,商品分类,商品搜索)
 * @Author: hai-27
 * @Date: 2020-02-07 16:23:00
 * @LastEditors: hai-27
 * @LastEditTime: 2020-03-08 12:11:13
 -->
<template>
  <div class="goods" id="goods" name="goods">
    <!-- 面包屑 -->
    <div class="breadcrumb">
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item to="/">首页</el-breadcrumb-item>
        <el-breadcrumb-item>全部货品</el-breadcrumb-item>
        <el-breadcrumb-item >分类</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <!-- 面包屑END -->



    <!-- 主要内容区 -->
    <div class="main">

        <div>
        <el-table :data="tableData.slice((currentPage-1)*pagesize,currentPage*pagesize)" style="width: 100%" highlight-current-row=true >
      <el-table-column prop="goodtype" label="货品类型"       
      :filters="[{text: '钢材', value: '1'}, {text: '木方', value: '2'}, {text: '烧结砖', value: '3'}]"
      :filter-method="filterHandler">
      </el-table-column>
      <el-table-column prop="name" label="货品名称" >
      </el-table-column>
      <el-table-column prop="standards" label="规格" >
      </el-table-column>
      <el-table-column prop="material" label="材料">
      </el-table-column>
      <el-table-column prop="workmanship" label="制作工艺">
      </el-table-column>
      <el-table-column prop="supplier" label="供应商">
      </el-table-column>
      <el-table-column  >
        <template slot-scope="scope">
        <el-button @click="handleNewOrder(scope.row)">下单</el-button>
        </template>
      </el-table-column>
    </el-table>
          <el-pagination background layout="prev, pager, next" :page-size="pagesize" :total="this.total" @current-change="current_change" class="page">
		</el-pagination>
      </div>
    </div>
    <el-dialog title="订单详情"
      class="detail-dialog"
      :visible.sync="showDetailDialog"
      :close-on-click-modal="false"
      :close-on-press-escape="false">
      <el-form class="detail-form" :model="detailForm" ref="detailForm" :rules="rules">
        <el-form-item label="发起企业：" prop="buyer" label-width="100px">
        <span>{{detailForm.buyer}}</span>
        </el-form-item>
        <el-form-item label="供应商：" prop="seller" label-width="100px">
          <span>{{detailForm.seller}}</span>
        </el-form-item>
          <el-form-item label-width="100px" label="货品信息：">
        <el-col span="5">货品名称：{{goodsname}}</el-col>
        <el-col span="5">规格：{{standards}}</el-col>
        <el-col span="5">原材料：{{material}}</el-col>
        <el-col span="5">制作工艺：{{workmanship}}</el-col>
        </el-form-item>
      <el-form-item label="货品数量："  prop="goodsnum" label-width="100px">
        <el-input v-model="detailForm.goodsnum"></el-input>
      </el-form-item>
        <el-form-item label="订单金额：" prop="orderamount" label-width="100px">
        <el-input v-model="detailForm.orderamount"></el-input>
        </el-form-item>
          <el-form-item label="收货地址：" prop="address" label-width="100px">
        <el-input v-model="detailForm.address"></el-input>
        </el-form-item>
        <el-form-item label="发起时间：" prop="createtime" label-width="100px">
          <span>{{detailForm.createtime}}</span>
        </el-form-item>
        <el-form-item label="截止时间：" prop="deadline" label-width="100px">
        <el-date-picker type="date" placeholder="选择截止日期" v-model="detailForm.deadline" style="width: 100%;"></el-date-picker>
        </el-form-item>
        <el-form-item label="订单备注：" prop="remark" label-width="100px">
      <el-input v-model="detailForm.orderamount" type="textarea"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="">
        <el-button @click="placeOrder('detailForm')">确定下单</el-button>
        <el-button @click="closeForm()">取消</el-button>
      </div>
    </el-dialog>
    <!-- 主要内容区END -->
  </div>
</template>

<script>
export default {
  data() {
    return {
      tableData:[{"goodsID":'1',"name":'钢筋',"goodtype":'1',"amount":'100吨',"standards":'10*100',"material":'精钢',"workmanship":'冷拉',"supplier":'北京钢铁集团',"adddate":'2021-6-7',"tel":'123123123',"shippingarea":'北京市辖区',"remark":'无'},
      {"goodsID":'1',"name":'钢筋',"goodtype":'1',"amount":'100吨',"standards":'10*100',"material":'精钢',"workmanship":'冷拉',"supplier":'北京钢铁集团',"adddate":'2021-6-7',"tel":'123123123',"shippingarea":'北京市辖区',"remark":'无'},
      {"goodsID":'1',"name":'钢筋',"goodtype":'1',"amount":'100吨',"standards":'10*100',"material":'精钢',"workmanship":'冷拉',"supplier":'北京钢铁集团',"adddate":'2021-6-7',"tel":'123123123',"shippingarea":'北京市辖区',"remark":'无'}],
      goodsname:'',
      standards:'',
      material:'',
      workmanship:'',
      showDetailDialog: false,
      detailForm: {
        ID:'',
        buyer:'',
        seller:'',
        buyerID:'',
        sellerID:'',
        address:'',
        goodsID:'',
        goodsname:'',
        goodsnum:'',
        orderamount:'',
        createtime:'',
        deadline:'',
        remark:'',
        transferstatus:'',
        transferID:'',
        status:'0'
      },
      rules:{
        goodsnum:[{required:true}],orderamount:[{required:true}],deadline:[{required:true}],address:[{required:true}]
      },
      total:0,
      currentPage:1,
      pagesize:15
    };
  },
  created() {
    this.getData()
  },
  activated() {
  },
  watch: {},
  methods: {
    handleNewOrder(row) {
        console.log("下单啦")
        this.detailForm.buyer=''
        this.detailForm.seller = row.supplier
        //this.detailForm.buyerID=userID
        this.detailForm.sellerID=row.supplierID
        this.detailForm.address=''
        this.detailForm.goodsID=row.goodsID
        this.detailForm.goodsname=row.name
        this.detailForm.goodsnum=''
        this.detailForm.orderamount=''
        this.detailForm.deadline=''
        this.detailForm.reamrk=''
        this.detailForm.transferstatus='0'
        this.status = '0'
        this.showDetailDialog = true
        this.goodsname = row.name
        this.material = row.material
        this.workmanship = row.workmanship
        this.standards = row.standards
        let date = new Date();
       let year = date.getFullYear(); // 年
       year = year +  "年"
       let month = date.getMonth() + 1; // 月
       month = month + "月"
        let day = date.getDate(); // 日
        day = day + "日"
        let hour = date.getHours(); // 时
        hour = hour < 10 ? "0" + hour : hour; // 如果只有一位，则前面补零
        let minute = date.getMinutes(); // 分
        minute = minute < 10 ? "0" + minute : minute; // 如果只有一位，则前面补零
        let time = hour + ":" + minute
        this.detailForm.createtime = year + month + day + time
    },
     filterHandler(value, row, column) {
        const property = column['property'];
        return row[property] === value;
    },
    closeForm() {
      this.showDetailDialog = false
    },
    placeOrder(formName) {
      this.$refs[formName].validate((valid) =>{
        if(valid) {
          this.showDetailDialog = false;
          console.log(this.detailForm);
          this.$message({
          message: '订单提交成功',
          type: 'success'
        });
        }else{
        this.$message({
          message: '请补全订单信息',
          type: 'error'
        });}
      })
    },
      current_change(currentPage){
      this.currentPage = currentPage;
      },

    getData() {
      this.total = this.tableData.length
    },


  }
};
</script>

<style scoped>
.goods {
  background-color: #f5f5f5;
}
/* 面包屑CSS */
.el-tabs--card .el-tabs__header {
  border-bottom: none;
}
.goods .breadcrumb {
  height: 50px;
  background-color: white;
}
.goods .breadcrumb .el-breadcrumb {
  width: 1225px;
  line-height: 30px;
  font-size: 16px;
  margin: 0 auto;
}
/* 面包屑CSS END */


/* 主要内容区CSS */
.goods .main {
  margin: 0 auto;
  max-width: 1225px;
}
.goods .main .list {
  min-height: 650px;
  padding-top: 14.5px;
  margin-left: -13.7px;
  overflow: auto;
}

.goods .main .none-product {
  color: #333;
  margin-left: 13.7px;
}

    .page {
    margin-top: 20px;
    text-align: center
  }
/* 主要内容区CSS END */
</style>