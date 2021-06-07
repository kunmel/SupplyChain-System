<template>
  <div class="components-container">
    <div class="container">
      <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
        <!-- 货品信息 -->
        <el-tab-pane label="货品信息" name="first">
          <h3 class="title">货品信息</h3>
          <!-- 基本信息-表单内容 -->
          <el-form class="basic-form" ref="basicForm" :model="basicForm" :rules="rules" label-width="100px">
            <el-form-item label="货品名称" prop="name">
              <el-input v-model="basicForm.name"></el-input>
            </el-form-item>
            <el-form-item label="货品规格" prop="standards">
              <el-input v-model="basicForm.standards"></el-input>
            </el-form-item>
            <el-form-item label="原材料" prop="material">
              <el-input v-model="basicForm.material"></el-input>
            </el-form-item>
            <el-form-item label="制作工艺" prop="workmanship">
              <el-input v-model="basicForm.workmanship"></el-input>
            </el-form-item>
            <el-form-item label="供应商" prop="supplier">
              <el-input v-model="basicForm.supplier"></el-input>
            </el-form-item>
            <el-form-item label="联系方式" prop="tel">
              <el-input v-model="basicForm.tel"></el-input>
            </el-form-item>
            <el-form-item label="所在地域" required>
              <el-col :span="11">
                <el-form-item prop="region1">
                  <el-select v-model="basicForm.region1" placeholder="请选择" @change="chooseProvince">
                    <el-option v-for="item of basicForm.region1List" :label="item.label" :key="item.index" :value="item.value"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col class="line" :span="1">&nbsp;-</el-col>
              <el-col :span="11">
                <el-form-item prop="region2">
                  <el-select v-model="basicForm.region2" placeholder="请选择" @change="chooseCity">
                    <el-option v-for="item of basicForm.region2List" :label="item.label" :key="item.index" :value="item.value"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
            </el-form-item>
            <el-form-item label="备注" prop="remark">
              <el-input type="textarea" v-model="basicForm.remark"></el-input>
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
import { provinceAndCityData } from 'element-china-area-data'
export default {
  data() {
    return {
      activeName: 'first',
      basicForm: {
        name: '',
        standards: '',
        material: '',
        workmanship: '',
        supplier: '',
        region1: '',
        region1List: [],
        region2: '',
        region2List: [],
        adddate: '',
        tel: '',
        remark: '',
        shippingarea:''
      },
      content: '',
      // 基本信息验证规则
      rules: {
        name: [
          { required: true, message: '请输入货品名称', trigger: 'blur' },
          { min: 1, max: 20, message: '长度在 1 到 20 个字符', trigger: 'blur' }
        ],
        standards: [
          { required: true, message: '请输入货品规格', trigger: 'blur' }
        ],
        material: [
          { required: true, message: '请输入货品原材料', trigger: 'blur' }
        ],
        workmanship: [
          { required: true, message: '请输入制作工艺', trigger: 'blur' }
        ],
        region1: [
          { required: true, message: '请选择地区', trigger: 'blur' }
        ],
        region2: [
          { required: true, message: '请选择城市', trigger: 'blur' }
        ],
        supplier: [
          { required: true, message: '请输入供货商名称', trigger: 'blur' }
        ],
        tel: [
          { required: true, message: '请输入联系电话', trigger: 'blur' }
        ],
        remarks: [
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
      this.basicForm.region1 = val
      this.basicForm.region2 = ''
      let _provinceList = this.basicForm.region1List
      for(let provinceItem of _provinceList) {
        if(val === provinceItem.value) {
          this.basicForm.region2List = provinceItem.children
          break
        }
      }
    },
    // 城市-选择
    chooseCity(val) {
      this.basicForm.region2 = val
      this.basicForm.shippingarea = this.basicForm.region1 + this.basicForm.region2
    },
    // 基本信息提交
    submitbasicForm(formName) {
      this.$refs[formName].validate((valid) => {
        if(valid) {
          let params = {
            name: this.basicForm.name,
            standards: this.basicForm.standards,
            material: this.basicForm.material,
            workmanship: this.basicForm.workmanship,
            supplier: this.basicForm.supplier,
            adddate: this.basicForm.adddate,
            tel: this.basicForm.tel,
            shippingarea: this.basicForm.shippingarea,
            remark: this.basicForm.remark,
          }
          console.log(params)
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
    this.basicForm.region1List = provinceAndCityData
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
