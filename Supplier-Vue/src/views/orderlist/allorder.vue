<template>
  <div>
    <el-table
      :data="tableData.slice((currentPage-1)*pagesize,currentPage*pagesize)"
      style="height:100%,width: 100%">
      <el-table-column align="center" >
        <el-table-column type="expand">
                <template slot-scope="props">
        <el-form label-position="left" inline class="demo-table-expand">
          <el-form-item label="订单ID">
            <span>{{ props.row.ID}}</span>
          </el-form-item>
          <el-form-item label="发起企业">
            <span>{{ props.row.Buyer}}</span>
          </el-form-item>
          <el-form-item label="供货商">
            <span>{{ props.row.Seller}}</span>
          </el-form-item>
          <el-form-item label="收货地址">
            <span>{{ props.row.Address}}</span>
          </el-form-item>
          <el-form-item label="货品名称">
            <span>{{ props.row.Goodsname}}</span>
          </el-form-item>
          <el-form-item label="货品数量">
            <span>{{ props.row.Goodsnum}}</span>
          </el-form-item>
                    <el-form-item label="总金额">
            <span>{{ props.row.Orderamount}}</span>
          </el-form-item>
                    <el-form-item label="创建时间">
            <span>{{ props.row.Createtime}}</span>
          </el-form-item>
                    <el-form-item label="截止日期">
            <span>{{ props.row.Deadline}}</span>
          </el-form-item>
                              <el-form-item label="物流状态">
            <span>{{transfermap.boolean[props.row.Transferstatus]}}</span>
          </el-form-item>
                              <el-form-item label="物流ID">
            <span>{{ props.row.TransferID}}</span>
          </el-form-item>
                              <el-form-item label="订单状态">
            <span>{{ordermap.boolean[props.row.Status]}}</span>
          </el-form-item>
          <el-form-item label="备注">
            <span>{{ props.row.Remark}}</span>
          </el-form-item>
        </el-form>
      </template>
      </el-table-column>
          <el-table-column type="index" label="ID"></el-table-column>
      </el-table-column>
      
      <el-table-column align="center" label="相关信息">
        <el-table-column property="Buyer" align="center" label="发起人(企业)" show-overflow-tooltip>
        </el-table-column>
         <el-table-column property="Goodsname" align="center" label="对应货品"  >
      </el-table-column>
      <!-- <el-table-column property="Standards" align="center" label="货品规格"  >
      </el-table-column> -->
      <el-table-column property="Goodsnum" align="center" label="商品总量" sortable="custom"  >
      </el-table-column>

      <el-table-column property="Orderamount" align="center" label="商品总金额" sortable="custom"  >
      </el-table-column>
        <el-table-column property="Deadline" align="center" label="截止日期" >
        </el-table-column>

        <el-table-column property="Status" align="center" label="订单状态" >
          <template slot-scope="scope">
            {{ordermap.boolean[scope.row.Status]}}
          </template>
        </el-table-column>

        <el-table-column property="Transferstatus" align="center" label="物流状态" >
          <template slot-scope="scope">
            {{transfermap.boolean[scope.row.Transferstatus]}}
          </template>
        </el-table-column>
        <el-table-column property="TransferID" align="center" label="物流单号" >
        </el-table-column>
        <el-table-column align="center" width="220" label="操作">
          <template slot-scope="scope">
            <el-button plain size="mini" @click="showtransferstatus(scope.row)">更新物流状态</el-button>
            <!-- <el-button plain size="mini" @click="delOne(scope.row)">拒绝</el-button> -->
          </template>
        </el-table-column>
      </el-table-column>
    </el-table>
    		<el-pagination background layout="prev, pager, next" :page-size="pagesize" :total="this.total" @current-change="current_change" class="page">
		</el-pagination>



<el-dialog title="更新物流状态" :visible.sync="dialogFormVisible">
  <el-form :model="form" ref="form" :rules="rules">
    
    <el-form-item label="收货地址" :label-width="formLabelWidth">
    <span>{{form.Address}}</span>
    </el-form-item>
    <el-form-item label="物流单号" :label-width="formLabelWidth" prop="TransferID">
      <el-input v-model="form.TransferID" autocomplete="off"></el-input>
    </el-form-item>
    
    <el-form-item label="物流状态" :label-width="formLabelWidth" prop="Transferstatus">
      <el-select v-model="form.Transferstatus" placeholder="请选择物流状态" >
        <el-option label="未发货" value="0"></el-option>
        <el-option label="已发货" value="1"></el-option>
        <el-option label="已送达" value="2"></el-option>
      </el-select>
    </el-form-item>
  </el-form>
  <div slot="footer" class="dialog-footer">
    <el-button @click="dialogFormVisible = false">取 消</el-button>
    <el-button type="primary" @click="UpdateTransferStatus('form')">确 定</el-button>
  </div>
</el-dialog>
  </div>
</template>

<script>
import { getMergeTable,QueryOrderBySeller,UpdateTransferStatus} from '@/api/excel'
import { getToken,getUserID,getUserName } from '@/common/auth'
export default {
  data() {
    return {
      transfermap : {boolean: {0:'未发货',1:'已发货',2:'已送达'}},
      ordermap : {boolean: {0:'待确认/融资',1:'已确认',2:'已完成',3:'供应商拒绝',4:'融资被拒绝'}},
      dialogFormVisible: false,
      form: {
      },
      formLabelWidth: '120px',
      tableData: [],
      total:0,
      currentPage:1,
      pagesize:20,
      rules: {
        Transferstatus: [
          { required: true, message: '请选择物流状态', trigger: 'blur' }
        ],
      }
    }
  },
  computed: {
    maxHeight() {
      return `${window.screen.availHeight - 220}px`
    }
  },
  methods: { 
    showtransferstatus(row){
        this.form = row
        this.dialogFormVisible = true
    },
    UpdateTransferStatus(formName){
       this.$refs[formName].validate((valid) => {
          this.dialogFormVisible = false
          const data = {
            ID:this.form.ID,
            TransferID:this.form.TransferID,
            Transferstatus:this.form.Transferstatus.toString()
          }

        UpdateTransferStatus(data).then(res => {
        console.log("UpdateTransferStatus",res)
        this.$message({
          message: '更新物流状态成功',
          type: 'success'
        });
      })

          //alert("更新ID为："+this.form.ID)
       })

    },
    getData() {
      let data = {
        Seller:getToken()
      }
      QueryOrderBySeller(data).then(res => {
        console.log("QueryOrderBySeller",res)
        let len=res.length;
        for(var j = 0; j < len; j++) {
          if(res[j].Status != 0){
            this.tableData.push(res[j])
          }
        }
        this.total = this.tableData.length
      })
    },
        current_change(currentPage){
        this.currentPage = currentPage;
      },
    // 合计
    getSummaries(param) {
      const { columns, data } = param
      const sums = []
      columns.forEach((column, index) => {
        if(index === 0) {
          sums[index] = ''
          return
        }
        if(index === 1) {
          sums[index] = '总计'
          return
        }
        if(index === 4 || index === 9) {
          const values = data.map(item => item[column.property])
          console.log(values)
          // 分子
          let s1 = values.reduce((prev, curr) => {
            const value = Number(curr.split('/')[0])
            if (!isNaN(value)) {
              return prev + value
            } else {
              return prev
            }
          }, 0)
          // 分母
          let s2 = values.reduce((prev, curr) => {
            const value = Number(curr.split('/')[1])
            if (!isNaN(value)) {
              return prev + value
            } else {
              return prev
            }
          }, 0)
          // 百分比
          let s3 = `${(s1 * 100 / s2).toFixed(2)}%`
          sums[index] = `${s1}/${s2} - (${s3})`
          sums[index] += ''
          return
        }
        const values = data.map(item => Number(item[column.property]))
        if (!values.every(value => isNaN(value))) {
          sums[index] = values.reduce((prev, curr) => {
            const value = Number(curr)
            if (!isNaN(value)) {
              return prev + curr
            } else {
              return prev
            }
          }, 0)
          sums[index] += ''
        } else {
          sums[index] = 'N/A'
        }
      })
      return sums
    }
  },
  created() {
    this.getData()
  },
  activated(){
    this.getData()
  }
}
</script>

<style scoped>
    .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
    .page {
    margin-top: 20px;
    text-align: center
  }
</style>