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
        <el-breadcrumb-item to="/company">首页</el-breadcrumb-item>
        <el-breadcrumb-item>全部货品</el-breadcrumb-item>
        <el-breadcrumb-item >分类</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <!-- 面包屑END -->



    <!-- 主要内容区 -->
    <div class="main">

        <div>
        <el-table :data="tableData.slice((currentPage-1)*pagesize,currentPage*pagesize)" style="width: 100%" highlight-current-row=true >
      <!-- <el-table-column prop="Goodtype" label="货品类型"       
      :filters="[{text: '钢材', value: 1}, {text: '木方', value: 2}, {text: '烧结砖', value: 3}]"
      :filter-method="filterHandler"> -->
      <el-table-column prop="Goodtype" label="货品类型">
          <template slot-scope="scope">
            {{typemap.boolean[scope.row.Goodtype]}}
          </template>
      </el-table-column>
      <el-table-column prop="Name" label="货品名称" >
      </el-table-column>
      <el-table-column prop="Standards" label="规格" >
      </el-table-column>
      <el-table-column prop="Material" label="材料">
      </el-table-column>
      <el-table-column prop="Workmanship" label="制作工艺">
      </el-table-column>
      <el-table-column prop="Supplier" label="供应商">
      </el-table-column>
      <el-table-column  >
        <template slot-scope="scope">
        <el-button @click="handleNewOrder(scope.row)">下单</el-button>
        </template>
      </el-table-column>
    </el-table>
          <el-pagination background layout="prev, pager, next" :page-size="pagesize" :total="this.tableData.length" @current-change="current_change" class="page">
		</el-pagination>
      </div>
    </div>
    <el-dialog title="订单详情"
      class="detail-dialog"
      :visible.sync="showDetailDialog"
      :close-on-click-modal="false"
      :close-on-press-escape="false">
      <el-form class="detail-form" :model="detailForm" ref="detailForm" :rules="rules">
        <el-form-item label="发起企业：" prop="Buyer" label-width="100px">
        <span>{{userName}}</span>
        </el-form-item>
        <el-form-item label="供应商：" prop="Seller" label-width="100px">
          <span>{{detailForm.Seller}}</span>
        </el-form-item>
          <el-form-item label-width="100px" label="货品信息：">
        <el-col span="5">货品名称：{{Goodsname}}</el-col>
        <el-col span="5">规格：{{Standards}}</el-col>
        <el-col span="5">原材料：{{Material}}</el-col>
        <el-col span="5">制作工艺：{{Workmanship}}</el-col>
        </el-form-item>
      <el-form-item label="货品数量："  prop="Goodsnum" label-width="100px">
        <el-input v-model="detailForm.Goodsnum"></el-input>
      </el-form-item>
        <el-form-item label="订单金额：" prop="Orderamount" label-width="100px">
        <el-input v-model="detailForm.Orderamount"></el-input>
        </el-form-item>
          <el-form-item label="收货地址：" prop="Address" label-width="100px">
        <el-input v-model="detailForm.Address"></el-input>
        </el-form-item>
        <el-form-item label="发起时间：" prop="Createtime" label-width="100px">
          <span>{{detailForm.Createtime}}</span>
        </el-form-item>
        <el-form-item label="截止时间：" prop="Deadline" label-width="100px">
        <el-date-picker type="date" placeholder="选择截止日期" v-model="detailForm.Deadline" style="width: 100%;" value-format="yyyy-MM-dd"></el-date-picker>
        </el-form-item>
        <el-form-item label="订单备注：" prop="Remark" label-width="100px">
      <el-input v-model="detailForm.Remark" type="textarea"></el-input>
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

import {QueryProduct,StoreOrder,QueryUserTest} from '@/api/company'
import { getToken,getUserID,getUserName } from '@/common/auth'
export default {
  data() {
    return {
      typemap:{boolean: {1:'钢材',2:'木方',3:'烧结砖'}},
      userName:'',
      tableData:[],
      Goodsname:'',
      Standards:'',
      Material:'',
      Workmanship:'',
      showDetailDialog: false,
      detailForm: {
        ID:'',
        Buyer:'',
        Seller:'',
        BuyerID:'',
        SellerID:'',
        Address:'',
        GoodsID:'',
        Goodsname:'',
        Goodsnum:'',
        Orderamount:'',
        Createtime:'',
        Deadline:'',
        Remark:'',
        Transferstatus:'',
        TransferID:'',
        Status:'0'
      },
      rules:{
        Goodsnum:[{required:true}],Orderamount:[{required:true}],Deadline:[{required:true}],Address:[{required:true}]
      },
      total:0,
      currentPage:1,
      pagesize:15
    };
  },
  created() {
    this.userName = getUserName()
    this.getData()
  },
  mounted() {
    this.userName = getUserName()
    this.getData()
  },
  activated() {
    this.userName = getUserName()
    this.getData()
  },
  watch: {},
  methods: {
    handleNewOrder(row) {
        console.log("下单啦")
        this.detailForm.Buyer=getToken()
        this.detailForm.Seller = row.Supplier
        this.detailForm.BuyerID= getUserID()
        this.detailForm.SellerID=row.SupplierID
        // this.detailForm.Address=''
        this.detailForm.GoodsID=row.GoodsID
        this.detailForm.Goodsname=row.Name
        // this.detailForm.Goodsnum=''
        // this.detailForm.Orderamount=''
        // this.detailForm.Deadline=''
        // this.detailForm.Reamrk=''
        // this.detailForm.ID = 'xxxx'
        this.detailForm.Deadline = row.Deadline
        this.detailForm.Transferstatus='0'
        this.Status = '0'
        this.showDetailDialog = true
        this.Goodsname = row.Name
        this.Material = row.Material
        this.Workmanship = row.Workmanship
        this.Standards = row.Standards
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
        this.detailForm.Createtime = year + month + day + time
    },
     filterHandler(value, row, column) {
        const property = column['property'];
        return row[property] === value;
    },
    closeForm() {
      this.showDetailDialog = false
    },
    placeOrder(formName) {
      console.log(this.detailForm)
      this.$refs[formName].validate((valid) =>{
        if(valid) {
          this.showDetailDialog = false;
        const data = {
          ID:this.detailForm.ID,
          Buyer:this.detailForm.Buyer,
          Seller:this.detailForm.Seller,
          BuyerID:this.detailForm.BuyerID,
          SellerID:this.detailForm.SellerID,
          Address:this.detailForm.Address,
          GoodsID:this.detailForm.GoodsID,
          Goodsname:this.detailForm.Goodsname,
          Goodsnum:this.detailForm.Goodsnum,
          Orderamount:this.detailForm.Orderamount,
          Createtime:this.detailForm.Createtime,
          //Deadline:'',
          Deadline: this.detailForm.Deadline,
          Remark:this.detailForm.Remark,
          Transferstatus:this.detailForm.Transferstatus,
          TransferID:this.detailForm.TransferID,
        }
        console.log
        StoreOrder(data).then(res => {
        console.log("StoreOrser 参数",data);
        console.log("res",res)
        this.$message({
          message: '订单提交成功',
          type: 'success'
        });
      })
          // StoreOrder(this.detailForm);
          // console.log("StoreOrser 参数",this.detailForm);
          
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
      QueryProduct().then(res => {
        console.log("QueryProduct",res)
        this.tableData = res
      })
    },

    getRealData(){
      QueryProductTest().then(resp => {
          console.log("QueryProduct",resp)
          let data = resp
          this.tableData = data
          // this.tableData.goodtype = resp.Goodtype
          // this.tableData.name = resp.Name
          // this.tableData.standards = resp.Standards
          // this.tableData.material = resp.Material
          // this.tableData.workmanship = resp.Workmanship
          // this.tableData.supplier = resp.Supplier
          return resolve(data)
        }).catch(err => {
          console.log("err",err)
          return reject(err)
        })
    }


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