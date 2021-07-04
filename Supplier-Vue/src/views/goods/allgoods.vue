<template>
  <div>
    <el-table
      :data="tableData"
      :max-height="tableData.slice((currentPage-1)*pagesize,currentPage*pagesize)"
      style="width: 100%">

      

      <el-table-column align="center" >
        <el-table-column type="expand">
      <template slot-scope="props">
        <el-form label-position="left" inline class="demo-table-expand">
          <el-form-item label="货品名称">
            <span>{{ props.row.Name}}</span>
          </el-form-item>
          <el-form-item label="规格">
            <span>{{ props.row.Standards}}</span>
          </el-form-item>
          <el-form-item label="原材料">
            <span>{{ props.row.Material}}</span>
          </el-form-item>
          <el-form-item label="制作工艺">
            <span>{{ props.row.Workmanship}}</span>
          </el-form-item>
          <el-form-item label="供应商">
            <span>{{ props.row.Supplier}}</span>
          </el-form-item>
          <el-form-item label="上架时间">
            <span>{{ props.row.Adddate}}</span>
          </el-form-item>
          <el-form-item label="联系方式">
            <span>{{ props.row.Tel}}</span>
          </el-form-item>
          <el-form-item label="发货地区">
            <span>{{ props.row.Shippingarea}}</span>
          </el-form-item>
          <el-form-item label="备注">
            <span>{{ props.row.Remark}}</span>
          </el-form-item>
        </el-form>
      </template>
      </el-table-column>
          <el-table-column type="index" label="ID"></el-table-column>
      </el-table-column>
      
      <el-table-column align="center" label="货品信息">
        <el-table-column property="Name" align="center" label="货品名称" show-overflow-tooltip>
        </el-table-column>
        <el-table-column property="Standards" align="center" label="规格" show-overflow-tooltip>
        </el-table-column>
        <!-- <el-table-column property="Material" align="center" label="材料" show-overflow-tooltip>
        </el-table-column> -->
         <el-table-column property="Workmanship" align="center" label="制作工艺" show-overflow-tooltip>
        </el-table-column>
        <el-table-column property="Amount" align="center" label="货品库存" show-overflow-tooltip>
        </el-table-column>
        <el-table-column property="Price" align="center" label="货品单价" show-overflow-tooltip>
        </el-table-column>
        <el-table-column property="Adddate" align="center" label="上架时间" >
        </el-table-column>
        <el-table-column property="Shippingarea" align="center" label="发货地区" >
        </el-table-column>
        <el-table-column align="center"  label="操作">
          <template slot-scope="scope">
            <el-button plain size="mini" @click="clickFileName(scope.row)">更新库存</el-button>
            <el-button plain size="mini" @click="delOne(scope.row)">下架</el-button>
          </template>
        </el-table-column>
      </el-table-column>
    </el-table>
    <el-pagination background layout="prev, pager, next" :page-size="pagesize" :total="this.total" @current-change="current_change" class="page">
		</el-pagination>

    <el-dialog title="更新库存" :visible.sync="dialogFormVisible">
  <el-form :model="form" ref="form" :rules="rules">
    
    <!-- <el-form-item label="货品名称" prop="name" :label-width="formLabelWidth">
      <el-input v-model="form.Name"></el-input>
    </el-form-item>
        <el-form-item label="规格" prop="standards" :label-width="formLabelWidth">
      <el-input v-model="form.Standards" autocomplete="off"></el-input>
    </el-form-item>
        <el-form-item label="材料" prop="material" :label-width="formLabelWidth">
      <el-input v-model="form.Material" autocomplete="off"></el-input>
    </el-form-item>
        <el-form-item label="制作工艺" prop="workmanship" :label-width="formLabelWidth">
      <el-input v-model="form.Workmanship" autocomplete="off"></el-input>
    </el-form-item> -->
            <el-form-item label="商品ID：" prop="GoodsID" label-width="100px">
          <span>{{form.GoodsID}}</span>
        </el-form-item>
        <el-form-item label="货品库存" prop="amount" :label-width="formLabelWidth">
      <el-input v-model="form.Amount" autocomplete="off"></el-input>
    </el-form-item>
        <!-- <el-form-item label="货品单价" prop="price" :label-width="formLabelWidth">
      <el-input v-model="form.Price" autocomplete="off"></el-input>
    </el-form-item>
    
    <el-form-item label="货品类型" prop="goodtype" :label-width="formLabelWidth">
      <el-select v-model="form.Goodtype" placeholder="请选择物流状态">
        <el-option label="钢材" value="1"></el-option>
        <el-option label="木方" value="2"></el-option>
        <el-option label="烧板砖" value="3"></el-option>
      </el-select>
    </el-form-item> -->
  </el-form>
  <div slot="footer" class="dialog-footer">
    <el-button @click="dialogFormVisible = false">取 消</el-button>
    <el-button type="primary" @click="UpdateProduct('form')">确 定</el-button>
  </div>
</el-dialog>

  </div>
</template>

<script>
import { DeleteProduct ,QueryProductBySupplier,UpdateProductByAmount} from '@/api/excel'
import { getToken,getUserID,getUserName } from '@/common/auth'
export default {
  data() {
    return {
      dialogFormVisible :false,
      tableData: [],
      total:0,
      form:{
        // Name: '',
        // Goodtype:'',
        GoodsID:'',
        Amount:'',
        // Standards: '',
        // Material: '',
        // Workmanship: '',
        // Supplier: '',
        // Region1: '',
        // Region1List: [],
        // Region2: '',
        // Region2List: [],
        // Adddate: '',
        // Tel: '',
        // Remark: '',
        // Shippingarea:''
      },
      currentPage:1,
      pagesize:20,
      rules: {
        // Name: [
        //   { required: true, message: '请输入货品名称', trigger: 'blur' },
        //   { min: 1, max: 20, message: '长度在 1 到 20 个字符', trigger: 'blur' }
        // ],
        // Standards: [
        //   { required: true, message: '请输入货品规格', trigger: 'blur' }
        // ],
        Amount: [
          { type: 'number',required: true, message: '请输入货品库存', trigger: 'blur' }
        ],
        // Price: [
        //   { required: true, message: '请输入价格', trigger: 'blur' }
        // ],
        // Goodtype: [
        //   { required: true, message: '请选择货品类型', trigger: 'blur' }
        // ],
        // Material: [
        //   { required: true, message: '请输入货品原材料', trigger: 'blur' }
        // ],
        // Workmanship: [
        //   { required: true, message: '请输入制作工艺', trigger: 'blur' }
        // ],
        // Supplier: [
        //   { required: true, message: '请输入供货商名称', trigger: 'blur' }
        // ],
      }
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
        console.log("row",row)
        let jsonStr = JSON.stringify(row)
        //this.form = JSON.parse(jsonStr);
        this.form.GoodsID = row.GoodsID
        this.dialogTableVisible = true
        this.dialogFormVisible = true
      }
    },
    delOne(row) {
        console.log("delOne",row.GoodsID)
        const data = {
          GoodsID:row.GoodsID
        }
        DeleteProduct(data).then(res => {
        console.log("DeleteProduct",res)
        this.$router.go(0);
         })
    },
    UpdateProduct(formName){
      
      this.$refs[formName].validate((valid) => {
        if(valid) {
          this.dialogFormVisible = false
          let data = {
            //Goodtype:this.form.Goodtype,
            Amount:Number(this.form.Amount),
            // Name: this.form.Name,
            // Standards: this.form.Standards,
            // Material: this.form.Material,
            // Workmanship: this.form.Workmanship,
            // Tel: this.form.Tel,
            // Shippingarea: this.form.Shippingarea,
            // Remark: this.form.Remark,
            GoodsID:this.form.GoodsID
          }
        UpdateProductByAmount(data).then(res => {
        console.log("UpdateProductByAmount",res)
        this.$router.go(0);
      })
  
          //alert("更新的goodsID为："+params.goodsID)
        } else {
          return false
        }
      })
    },
    getData() {
      // getGoodsTable().then(res => {
        const data = {
          Supplier:getToken()
        }
        QueryProductBySupplier(data).then(res => {
        console.log("QueryProductBySupplier",res)
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
    activated() {
    this.getData()
  },
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