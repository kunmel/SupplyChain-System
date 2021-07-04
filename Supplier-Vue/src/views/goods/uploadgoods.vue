<template>
  <div class="components-container">
    <div class="container">
      <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
        <!-- 货品信息 -->
        <el-tab-pane label="货品信息" name="first">
          <h3 class="title">货品信息</h3>
          <!-- 基本信息-表单内容 -->
          <el-form class="basic-form" ref="basicForm" :model="basicForm" :rules="rules" label-width="100px">
            <el-form-item label="货品名称" prop="Name">
              <el-input v-model="basicForm.Name"></el-input>
            </el-form-item>
            <el-form-item label="货品类型" :label-width="formLabelWidth">
              <el-select v-model="basicForm.Goodtype" placeholder="请选择货品类型">
                <el-option label="钢材" value="1"></el-option>
                <el-option label="木方" value="2"></el-option>
                <el-option label="烧板砖" value="3"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="货品库存" prop="Amount">
              <el-input v-model="basicForm.Amount"></el-input>
            </el-form-item>
            <el-form-item label="货品规格" prop="Standards">
              <el-input v-model="basicForm.Standards"></el-input>
            </el-form-item>
              <el-form-item label="货品单价" prop="Price">
              <el-input v-model="basicForm.Price"></el-input>
            </el-form-item>
            <el-form-item label="原材料" prop="Material">
              <el-input v-model="basicForm.Material"></el-input>
            </el-form-item>
            <el-form-item label="制作工艺" prop="Workmanship">
              <el-input v-model="basicForm.Workmanship"></el-input>
            </el-form-item>
            <!-- <el-form-item label="供应商" prop="Supplier">
              <el-input v-model="basicForm.Supplier"></el-input>
            </el-form-item> -->
            <el-form-item label="联系方式" prop="Tel">
              <el-input v-model="basicForm.Tel"></el-input>
            </el-form-item>
            <el-form-item label="所在地域" required>
              <el-col :span="11">
                <el-form-item prop="Region1">
                  <el-select v-model="basicForm.Region1" placeholder="请选择" @change="chooseProvince">
                    <el-option v-for="item of basicForm.Region1List" :label="item.label" :key="item.index" :value="item.value"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col class="line" :span="1">&nbsp;-</el-col>
              <el-col :span="11">
                <el-form-item prop="Region2">
                  <el-select v-model="basicForm.Region2" placeholder="请选择" @change="chooseCity">
                    <el-option v-for="item of basicForm.Region2List" :label="item.label" :key="item.index" :value="item.value"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
            </el-form-item>
            <el-form-item label="备注" prop="Remark">
              <el-input type="textarea" v-model="basicForm.Remark"></el-input>
            </el-form-item>            
            <el-form-item>
              <el-button type="primary" @click="submitbasicForm('basicForm')">保存</el-button>
              <el-button @click="cancel">取消</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        
      </el-tabs>
    </div>
  </div>
</template>

<script>
import { provinceAndCityData,CodeToText } from 'element-china-area-data'
import { StoreProduct} from '@/api/excel'
import { getToken,getUserID,getUserName } from '@/common/auth'
export default {
  data() {
    return {
      activeName: 'first',
      address:'',
      basicForm: {
        Name: '',
        Goodtype:'',
        Amount:'',
        Standards: '',
        Material: '',
        Price:'',
        Workmanship: '',
        Supplier: '',
        Region1: '',
        Region1List: [],
        Region2: '',
        Region2List: [],
        Adddate: '',
        Tel: '',
        Remark: '',
        Shippingarea:''
      },
      content: '',
      // 基本信息验证规则
      rules: {
        Name: [
          { required: true, message: '请输入货品名称', trigger: 'blur' },
          { min: 1, max: 20, message: '长度在 1 到 20 个字符', trigger: 'blur' }
        ],
        Standards: [
          { required: true, message: '请输入货品规格', trigger: 'blur' }
        ],
        Amount: [
          { required: true, message: '请输入货品库存', trigger: 'blur' }
        ],
        Price: [
          {  required: true, message: '请输入货品单价', trigger: 'blur' }
        ],
        Goodtype: [
          { required: true, message: '请选择货品类型', trigger: 'blur' }
        ],
        Material: [
          { required: true, message: '请输入货品原材料', trigger: 'blur' }
        ],
        Workmanship: [
          { required: true, message: '请输入制作工艺', trigger: 'blur' }
        ],
        Region1: [
          { required: true, message: '请选择地区', trigger: 'blur' }
        ],
        Region2: [
          { required: true, message: '请选择城市', trigger: 'blur' }
        ],
        // Supplier: [
        //   { required: true, message: '请输入供货商名称', trigger: 'blur' }
        // ],
        Tel: [
          { required: true, message: '请输入联系电话', trigger: 'blur' }
        ],
        Remarks: [
          { max: 200, message: '200 字以内，非必填', trigger: 'blur' }
        ]
      }
    }
  },
  // computed: {
  //   sysMark() {
  //     return this.basicForm.sysMark
  //   }
  // },
  watch: {
    'basicForm.sysMark': function(newVal, oldVal) {
      if(this.basicForm.sysMarkType === '2') {
        this.basicForm.accountMark = newVal
      }
    }
  },
  methods: {
    // 系统标识切换
    handleSysMarkType(val) {
      // 系统标识-标准
      if(val === '1') {
        this.basicForm.sysMark = '.tencent.com'
      } else if(val === '2') { // 系统标识-个性
        this.basicForm.sysMark = ''
      }
    },
    // 省-选择
    chooseProvince(val) {
      this.basicForm.Region1 = val
      this.address = CodeToText[val]
      this.basicForm.Region2 = ''
      let _provinceList = this.basicForm.Region1List
      for(let provinceItem of _provinceList) {
        if(val === provinceItem.value) {
  
          this.basicForm.Region2List = provinceItem.children
          break
        }
      }
    },
    // 城市-选择
    chooseCity(val) {
      this.basicForm.Region2 = val
      this.address = this.address + CodeToText[val]
      this.basicForm.Shippingarea = this.basicForm.Region1 + this.basicForm.Region2
    },
    // 基本信息提交
    submitbasicForm(formName) {
      var date = new Date();
      var year = date.getFullYear();
      var month = date.getMonth() + 1;
      var day = date.getDate();
      if (month < 10) {
      month = '0' + month;
      }
      if (day < 10) {
      day = '0' + day;
      }
      var nowDate = year + '-'+ month +'-' + day;
      this.$refs[formName].validate((valid) => {
        if(valid) {
          
          let data = {
            Goodtype:Number(this.basicForm.Goodtype),
            Amount:Number(this.basicForm.Amount),
            Price:Number(this.basicForm.Price),
            Name: this.basicForm.Name,
            Standards: this.basicForm.Standards,
            Material: this.basicForm.Material,
            Workmanship: this.basicForm.Workmanship,
            Supplier: getToken(),
            Adddate: nowDate,
            Tel: this.basicForm.Tel,
            Shippingarea: this.address,
            Remark: this.basicForm.Remark,
          }
        StoreProduct(data).then(res => {
        console.log("StoreProduct",res)
          this.$message({
          message: '上架商品成功',
          type: 'success'
        });
        this.$router.push({ path: "/goods/allogoods" });
         })
          //console.log(params)
          // addOrg(_data).then( res => {
          //   console.log(res)
          // })
        } else {
          return false
        }
      })
    },
    // 取消提交
    cancel() {
      this.$refs['basicForm'].resetFields()
    },
    // tab切换点击事件
    handleClick(tab, e) {
      if(tab.name === 'second') {
        // nothing
      }
    }
  },
  created() {
    this.basicForm.Region1List = provinceAndCityData
  }

}
</script>
<style scoped lang="stylus">
.basic-form
  width 500px
.sys-mark-wrapper
  .input-two
    position absolute
    top 0
    right -110px
    width 100px
</style>
