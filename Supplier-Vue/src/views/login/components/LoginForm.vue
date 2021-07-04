<template>
<el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="sign-in-form">
  <el-form-item label="用户名" prop="username">
    <el-input v-model="ruleForm.username"></el-input>
  </el-form-item>
   <el-form-item label="密码" prop="userpass">
    <el-input type="password" v-model="ruleForm.userpass" autocomplete="off"></el-input>
  </el-form-item>
    <el-form-item label="用户类型" prop="usertype">
    <el-radio-group v-model="ruleForm.usertype">
      <el-radio label="企业"></el-radio>
      <el-radio label="供应商"></el-radio>
      <el-radio label="银行"></el-radio>
    </el-radio-group>
  </el-form-item>
    <el-form-item>
    <el-button type="primary" @click="submitForm('ruleForm')">立即登陆</el-button>
    <el-button>取消</el-button>
  </el-form-item>

</el-form>	
</template>
<script>
import { queryAccountList,QueryUserTest,QueryUser,login} from '@/api/account'
import { saveToLocal, loadFromLocal } from '@/common/local-storage'
import { mapActions } from 'vuex'
import { mapGetters } from "vuex";


  export default {
    data() {
		var validatePass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入密码'));
        } else {
          if (this.ruleForm.checkPass !== '') {
            this.$refs.ruleForm.validateField('checkPass');
          }
          callback();
        }
      };
      return {
		loading: false,
		accountList:[],
    remember: false,
		redirect:'',
        ruleForm: {
          username: '',
          userpass:'',
		      usertype: '',
        },
        rules: {
          username: [
            { required: true, message: '请输入用户名', trigger: 'blur' },
            { min: 3, max: 15, message: 'Length should be 3 to 15', trigger: 'blur' }
          ],
		    userpass: [
			{ required: true,  trigger: 'blur' },
            { validator: validatePass, trigger: 'blur' }
          ],
		     usertype: [
            { required: true, message: '请选择用户类别', trigger: 'change' }
          ],
        }
      };
    },
	  //监听路由
  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
        console.log("动态监听路由：")
        console.log(route.query)
      },
      immediate: true
    }
  },
	created() {
    //区块链接口
    	//  queryAccountList().then(response => {
      //    this.accountList = response
      //    console.log("AccountList")
      //    console.log(response)
      //  })
      
       const data = {
         Account:'qiye',
         Password:'123'

       }

       QueryUser(data).then(response => {
         console.log("CreateQueryUser")
         console.log(response)
       })
         // 初始化时读取localStorage用户信息
    if (loadFromLocal('remember', false)) {
      // this.ruleForm.username = loadFromLocal('username', '')
      // this.ruleForm.pwd = loadFromLocal('password', '')
    } else {
      this.ruleForm.username = ''
      this.ruleForm.pwd = ''
    }
  },
    methods: {
    ...mapActions([
      'login'
    ]),
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
              
              this.loading = true

                const data = {
                  Account:this.ruleForm.username,
                  Password:this.ruleForm.userpass
                }
              console.log("huoqu data",data)
              this.login(data).then(response => {
                // 保存账号
                if (this.remember) {
  
                } else {

                }
              console.log('Identity:',response.Identity)
              // console.log('Permissions:',this.$store.getters.permissions)
                var Identity = response.Identity
                if (this.ruleForm.usertype  == "企业" && Identity == 1){
                  console.log("进入企业路由")
                  this.$router.push({ path: '/company' })
                } else if(this.ruleForm.usertype  == "供应商" && Identity == 2){     
                  this.$router.push({ path: '/home' })
                } else if(this.ruleForm.usertype  == "银行" && Identity == 3){
                  this.$router.push({ path: '/bankpage' })
                } else{
                  this.$notify.info({
                  title: '注意',
                  message: '用户身份错误'
                });
                }
        
              }).catch(() => {
                                 var Identity = this.$store.getters.identity
                if (this.ruleForm.usertype  == "企业" && Identity == 1){
                  this.$router.push({ path: '/company' })
                } else if(this.ruleForm.usertype  == "供应商" && Identity == 2){     
                  this.$router.push({ path: '/home' })
                } else if(this.ruleForm.usertype  == "银行" && Identity == 3){
                  this.$router.push({ path: '/bankpage' })
                } else{
                  this.$notify.info({
                  title: '注意',
                  message: '用户身份错误'
                });
                }
                // this.loading = false
                // this.$notify.info({
                //   title: '注意',
                //   message: '账户密码信息错误'
                // });
              })
          } else {
            return false;
          }
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
       
    },
}
</script>

<style scoped>
</style>
