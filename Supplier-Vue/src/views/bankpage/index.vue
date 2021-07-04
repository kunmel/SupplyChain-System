<template>
  <div>
    <el-table
      :data="tableData.slice((currentPage-1)*pagesize,currentPage*pagesize)"
      style="height:100%,width: 100%">

      <el-table-column align="center" >
          <el-table-column type="index" label="ID"></el-table-column>
      </el-table-column>
      
      <el-table-column align="center" label="相关信息">
        <el-table-column property="OrderID" align="center" label="订单号" show-overflow-tooltip>
        </el-table-column>
          <el-table-column property="Supplier" align="center" label="申请人" show-overflow-tooltip>
        </el-table-column>
        <el-table-column property="Amount" align="center" label="贷款总额" show-overflow-tooltip>
        </el-table-column>
        <el-table-column property="Period" align="center" label="贷款日期" >
        </el-table-column>
        <el-table-column property="Desc" align="center" label="贷款描述" >
        </el-table-column>
        <el-table-column property="Status" align="center" label="贷款状态" >
          <template slot-scope="scope">
                 {{map.boolean[scope.row.Status]}}
          </template>
        </el-table-column>
        </el-table-column>
        <!-- <el-table-column property="feedback" align="center" label="贷款反馈" >
        </el-table-column> -->
      <!-- <el-table-column align="center" width="220" label="操作"> -->
        <!-- <template slot-scope="scope">
          <el-button plain size="mini" @click="clickFileName(scope.row)">审批</el-button>
        </template> -->
      <!-- </el-table-column> --> 
      <!-- </el-table-column> -->
    </el-table>
    		<el-pagination background layout="prev, pager, next" :page-size="pagesize" :total="this.total" @current-change="current_change" class="page">
		</el-pagination>


    <!-- Form -->

<el-dialog title="更新融资状态" :visible.sync="dialogFormVisible">
  <el-form :model="form">
    
    <el-form-item label="融资人" :label-width="formLabelWidth">
    <span>{{form.Supplier}}</span>
    </el-form-item>
    <el-form-item label="贷款总额" :label-width="formLabelWidth">
    <span>{{form.Amount}}</span>
    </el-form-item>
    <el-form-item label="贷款日期" :label-width="formLabelWidth">
    <span>{{form.Period}}</span>

    </el-form-item>
    <el-form-item label="贷款描述" :label-width="formLabelWidth">
    <span>{{form.Desc}}</span>
    </el-form-item>
    <el-form-item label="贷款状态" :label-width="formLabelWidth">
    <span>{{map.boolean[form.Status]}}</span>
    </el-form-item>
    <el-form-item label="贷款反馈" :label-width="formLabelWidth">
      <el-input  type="textarea"  :autosize="{ minRows: 4, maxRows: 6}" v-model="form.Feedback" autocomplete="off"></el-input>
      
    </el-form-item>
    
  </el-form>
  <div slot="footer" class="dialog-footer">
    <el-button type="primary" @click="dialogFormVisible = false">通过申请</el-button>
    <el-button type="danger" @click="dialogFormVisible = false">拒绝申请</el-button>
  </div>
</el-dialog>
  </div>
</template>

<script>
import { getMergeTable ,getUserFiance,QueryFinance} from '@/api/excel'
export default {
  data() {
    return {

      map : {boolean: {0: '待审批', 1: '审批通过', 2: '审批拒绝'}},
      dialogTableVisible: false,
      dialogFormVisible: false,
      form: {
      },
      formLabelWidth: '120px',
      tableData: [],
      total:0,
      currentPage:1,
      pagesize:20
    }
  },
  computed: {
    maxHeight() {
      return `${window.screen.availHeight - 220}px`
    }
  },
  methods: { 
     // 点击文件名
    clickFileName(row) {
      // 如果是目录
      if(row.materialType === 0) {
        this.$emit('goFolder', row)
      } else {
        // 显示详情
        this.form = row
        this.dialogTableVisible = true
        this.dialogFormVisible = true
      }
    },
    getData() {
      QueryFinance().then(res => {
        let len=res.length;
        for(var j = 0; j < len; j++) {
          if(res[j].Status != 0){
            this.tableData.push(res[j])
          }
        }
        console.log("QueryFinance",res)
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
    activated() {
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