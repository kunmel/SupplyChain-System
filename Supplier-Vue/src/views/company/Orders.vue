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
        <el-breadcrumb-item>我的订单</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <!-- 面包屑END -->



    <!-- 主要内容区 -->
    <div class="main">
        <div>
        <el-table :data="tableData.slice((currentPage-1)*pagesize,currentPage*pagesize)" style="width: 100%" highlight-current-row=true >
      <el-table-column prop="ID" label="订单ID" >
      </el-table-column>
      <el-table-column prop="Seller" label="供货商" >
      </el-table-column>
      <el-table-column prop="Goodsname" label="货品名称">
      </el-table-column>
      <el-table-column prop="Deadline" label="截止时间">
      </el-table-column>
      <el-table-column prop="Status" label="订单状态">
          <template slot-scope="scope">
            {{ordermap.boolean[scope.row.Status]}}
          </template>
      </el-table-column>
      <el-table-column  >
        <template slot-scope="scope">
        <el-button  @click="handleNewOrder(scope.row)">详细信息</el-button>
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
      <el-form class="detail-form" :model="detailForm">
        <el-form-item label="订单ID：" prop="ID" label-width="100px">
          <span>{{detailForm.ID}}</span>
        </el-form-item>
        <el-form-item label="订单状态：" prop="Status" label-width="100px">
            {{ordermap.boolean[detailForm.Status]}}
          
        </el-form-item>
        <el-form-item label="物流状态：" prop="Transferstatus" label-width="100px">
          <span>{{transfermap.boolean[detailForm.Transferstatus]}}</span>
        </el-form-item>
        <el-form-item label="物流ID：" prop="TransferID" label-width="100px">
          <span>{{detailForm.TransferID}}</span>
        </el-form-item>
        <el-form-item label="发起企业：" prop="Buyer" label-width="100px">
          <span>{{detailForm.Buyer}}</span>
        </el-form-item>
        <el-form-item label="供货商：" prop="Seller" label-width="100px">
          <span>{{detailForm.Seller}}</span>
        </el-form-item>
        <el-form-item label="货品名称：" prop="Goodsname" label-width="100px">
          <span>{{detailForm.Goodsname}}</span>
        </el-form-item>
        <el-form-item label="货品数量：" prop="Goodsnum" label-width="100px">
          <span>{{detailForm.Goodsnum}}</span>
        </el-form-item>
        <el-form-item label="收货地址：" prop="Address" label-width="100px">
          <span>{{detailForm.Address}}</span>
        </el-form-item>
        <el-form-item label="订单总金额：" prop="Orderamount" label-width="100px">
          <span>{{detailForm.Orderamount }}</span>
        </el-form-item>
        <el-form-item label="创建时间：" prop="Createtime" label-width="100px">
          <span>{{detailForm.Createtime}}</span>
        </el-form-item>
        <el-form-item label="截止时间：" prop="Deadline" label-width="100px">
          <span>{{detailForm.Deadline}}</span>
        </el-form-item>
        <el-form-item label="备注：" prop="Remark" label-width="100px">
          <span>{{detailForm.Remark}}</span>
        </el-form-item>

      </el-form>
      <div slot="footer" class="">
        <el-button @click="closeForm()">关闭</el-button>
      </div>
    </el-dialog>
    <!-- 主要内容区END -->
  </div>
</template>

<script>
import {QueryOrderByBuyer,QueryOrder} from '@/api/company'
import { getToken,getUserID } from '@/common/auth'
export default {
  data() {
    return {
      ordermap : {boolean: {0:'待确认/融资',1:'已确认',2:'已完成',3:'供应商拒绝',4:'融资被拒绝'}},
      transfermap : {boolean: {0:'未发货',1:'已发货',2:'已送达'}},
      tableData:[],
      Goodsname:'',
      Standards:'',
      Material:'',
      Workmanship:'',
      showDetailDialog: false,
      detailForm: {},
      total:0,
      currentPage:1,
      pagesize:15
    };
  },
  created() {
    this.getData()
  },
  mounted(){
    this.getData()
  },
  activated() {
    this.getData()
  },
  watch: {},
  methods: {
    handleNewOrder(row) {
      this.showDetailDialog = true
      this.detailForm = row
      //   console.log("下单啦")
      //   this.detailForm.buyer=''
      //   this.detailForm.goodsnum=''
      //   this.detailForm.orderamount=''
      //   this.detailForm.deadline=''
      //   this.detailForm.reamrk=''
      //   this.showDetailDialog = true
      //   this.detailForm.goodsID = row.goodsID
      //   this.detailForm.seller = row.supplier
      //   this.goodsname = row.name
      //   this.material = row.material
      //   this.workmanship = row.workmanship
      //   this.standards = row.standards
      //   let date = new Date();
      //  let year = date.getFullYear(); // 年
      //  year = year +  "年"
      //  let month = date.getMonth() + 1; // 月
      //  month = month + "月"
      //   let day = date.getDate(); // 日
      //   day = day + "日"
      //   let hour = date.getHours(); // 时
      //   hour = hour < 10 ? "0" + hour : hour; // 如果只有一位，则前面补零
      //   let minute = date.getMinutes(); // 分
      //   minute = minute < 10 ? "0" + minute : minute; // 如果只有一位，则前面补零
      //   let time = hour + ":" + minute
      //   this.detailForm.createtime = year + month + day + time
    },
     filterHandler(value, row, column) {
        const property = column['property'];
        return row[property] === value;
    },
    closeForm() {
      this.showDetailDialog = false
    },
      current_change(currentPage){
      this.currentPage = currentPage;
      },

    getData() {
        const data ={
          Buyer :getToken()
        }
        console.log("QueryOrder参数",data)
        QueryOrderByBuyer(data).then(res => {
        //  QueryOrder().then(res=> {
        console.log("QueryOrder",res)
        this.tableData = res
      })
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