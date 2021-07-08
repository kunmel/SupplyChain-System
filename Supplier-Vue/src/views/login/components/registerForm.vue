<template>
	<el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="sign-up-form">
   <el-form-item label="用户名" prop="username">
    <el-input v-model="ruleForm.username"></el-input>
  </el-form-item>
    <el-form-item label="企业名称" prop="usermail">
    <el-input v-model="ruleForm.usermail"></el-input>
  </el-form-item>
  <el-form-item label="密码" prop="pass">
    <el-input type="password" v-model="ruleForm.pass" autocomplete="off"></el-input>
  </el-form-item>
  <el-form-item label="确认密码" prop="checkPass">
    <el-input type="password" v-model="ruleForm.checkPass" autocomplete="off"></el-input>
  </el-form-item>
  <el-form-item label="手机号" prop="phone">
    <el-input v-model.number="ruleForm.phone"></el-input>
  </el-form-item>
  <el-form-item label="用户类型" prop="usertype">
    <el-radio-group v-model="ruleForm.usertype">
      <el-radio label="企业"></el-radio>
      <el-radio label="供应商"></el-radio>
      <el-radio label="银行"></el-radio>
    </el-radio-group>
  </el-form-item>
  <el-form-item>
    <el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>
    <el-button @click="resetForm('ruleForm')">重置</el-button>
  </el-form-item>
</el-form>
</template>

<script>
//login css
// import '../css/global.css'
// import '../css/reset.css'
import { StoreUser} from '@/api/account'
import { saveToLocal, loadFromLocal } from '@/common/local-storage'
import { mapActions } from 'vuex'

  export default {
    data() {
	  var validateUsername = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入用户名'));
        } else {
          callback();
        }
      };
	  	  var validateUsermail = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入企业名称'));
        } else {
		  // if(!(value.includes('@')&&value.includes('.'))){
			//   callback(new Error('邮箱格式错误'));
		  // }else{
			callback();
		  //}
        }
      };
      var checkPhone = (rule, value, callback) => {
        if (!value) {
          return callback(new Error('手机号不能为空'));
        }
        setTimeout(() => {
          if (!Number.isInteger(value)) {
            callback(new Error('请输入数字值'));
          } else {
            if (value.toString().length != 11) {
				
              callback(new Error("电话格式错误！"));
            } else {
              callback();
            }
          }
        }, 1000);
      };
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
      var validatePass2 = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请再次输入密码'));
        } else if (value !== this.ruleForm.pass) {
          callback(new Error('两次输入密码不一致!'));
        } else {
          callback();
        }
      };
	  var validateUsertype = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请选择用户类型'));
        } else {
          callback();
        }
      };
      return {
        ruleForm: {
          pass: '',
          checkPass: '',
          phone: '',
		  username: '',
		  usertype: '',
		  usermail: '',
        },
        rules: {
          pass: [
            { validator: validatePass, trigger: 'blur' }
          ],
          checkPass: [
            { validator: validatePass2, trigger: 'blur' }
          ],
          phone: [
            { validator: checkPhone, trigger: 'blur' }
          ],
		  username: [
            { validator: validateUsername, trigger: 'blur' }
          ],
		  usertype: [
            { validator: validateUsertype, trigger: 'blur' }
          ],
		  usermail: [
            { validator: validateUsermail, trigger: 'blur' }
          ]

        }
      };
    },
    methods: {
      ...mapActions([
      'login',
      'register'
      ]),
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
              this.loading = true
              this.register(this.ruleForm).then(() => {
                // 保存账号
                if (this.remember) {
                  saveToLocal('username', this.ruleForm.username)
                  saveToLocal('password', this.ruleForm.userpass)
                  saveToLocal('remember', true)
                } else {
                  saveToLocal('username', '')
                  saveToLocal('password', '')
                  saveToLocal('remember', false)
                }
                var identity = '0'
                if(this.ruleForm.usertype == '企业'){
                  identity = '1'
                }else if(this.ruleForm.usertype == '供应商'){
                  identity = '2'
                }else{
                  identity = '3'
                }
                  const data = {
                  Account:this.ruleForm.username,
                  Password:this.ruleForm.pass,
                  Name:this.ruleForm.usermail,
                  Identity:identity
                }

                StoreUser(data).then(response => {
                  console.log("StoreUser")
                  console.log(response)
                })
                this.$message({
                  message: '注册成功，跳转至登陆界面',
                  type: 'success'
                });
                //this.$parent.changemode()
                this.$router.go(0);
              }).catch(() => {
                this.loading = false
              })
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      }
    }
  }
</script>

<style scoped>
</style>