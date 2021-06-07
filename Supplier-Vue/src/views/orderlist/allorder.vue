<template>
  <div>
    <el-table
      :data="tableData"
      :max-height="maxHeight"
      style="width: 100%">

      

      <el-table-column align="center" >
        <el-table-column type="expand">
      <template slot-scope="props">
        <el-form label-position="left" inline class="demo-table-expand">
          <el-form-item label="发起企业">
            <span>{{ props.row.buyer}}</span>
          </el-form-item>
          <el-form-item label="对应货品">
            <span>{{ props.row.goodsname}}</span>
          </el-form-item>
          <el-form-item label="货品数量">
            <span>{{ props.row.goodsnum}}</span>
          </el-form-item>
          <el-form-item label="订单总金额">
            <span>{{ props.row.orderamount}}</span>
          </el-form-item>
          <el-form-item label="是否已接收">
            <span>{{ props.row.ifaccept}}</span>
          </el-form-item>
          <el-form-item label="发起时间">
            <span>{{ props.row.createtime}}</span>
          </el-form-item>
          <el-form-item label="截止时间">
            <span>{{ props.row.deadline}}</span>
          </el-form-item>
          <el-form-item label="订单备注">
            <span>{{ props.row.remark}}</span>
          </el-form-item>
        </el-form>
      </template>
      </el-table-column>
          <el-table-column type="index" label="ID"></el-table-column>
      </el-table-column>
      
      <el-table-column align="center" label="相关信息">
        <el-table-column property="buyer" align="center" label="发起人(企业)" show-overflow-tooltip>
        </el-table-column>
        <el-table-column property="goodsname" align="center" label="货品名称" show-overflow-tooltip>
        </el-table-column>
        <el-table-column property="deadline" align="center" label="截止日期" >
        </el-table-column>
        <el-table-column property="ifaccept" align="center" label="是否已接收" >
        </el-table-column>
      </el-table-column>
      

    </el-table>
  </div>
</template>

<script>
import { getMergeTable } from '@/api/excel'
export default {
  data() {
    return {
      tableData: []
    }
  },
  computed: {
    maxHeight() {
      return `${window.screen.availHeight - 220}px`
    }
  },
  methods: {
    getData() {
      getMergeTable().then(res => {
        console.log(res)
        this.tableData = res.data
      })
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
</style>