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
        <el-table-column property="Feedback" align="center" label="贷款反馈" >
        </el-table-column>
      <!-- <el-table-column align="center" width="220" label="操作"> -->
        <!-- <template slot-scope="scope">
          <el-button plain size="mini" @click="clickFileName(scope.row)">更新物流状态</el-button>
          <el-button plain size="mini" @click="delOne(scope.row)">拒绝</el-button>
        </template> -->
      <!-- </el-table-column> -->
      <!-- </el-table-column> -->
    </el-table>
    		<el-pagination background layout="prev, pager, next" :page-size="pagesize" :total="this.total" @current-change="current_change" class="page">
		</el-pagination>

  </div>
</template>

<script>
import { getMergeTable ,QueryFinanceBySupplier} from '@/api/excel'
import { getToken,getUserID,getUserName } from '@/common/auth'
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
      let data = {
        Supplier:getToken()
      } 
      QueryFinanceBySupplier(data).then(res => {
        console.log("QueryFinanceBySupplier",res)
        this.tableData = res
        this.total = res.length
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