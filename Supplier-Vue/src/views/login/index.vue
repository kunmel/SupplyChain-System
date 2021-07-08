<template>
  <div class="container" :class="{ 'sign-up-mode': signUpMode }">
		<div class="forms-container source" >
			<div class="signin-signup" id="login-form">
				<login-form></login-form>
      </div>
      <div class="signin-signup" id="signup-form">
				<register-form :registerForm="registerForm"></register-form>
			</div>
		</div>
		<div class="panels-container">
			<div class="panel left-panel">
				<div class="content">
					<h3>振兴东北老工业基地，钢材区块链助力发展!</h3>
					
					<p>BIT 星之链钢材供应区块链</p>
					<button @click="switchMode()" class="btn transparent">
						去注册
					</button>
				</div>
				<!-- <img src="../../assets/img/log.svg" class="image" alt="" /> -->
        <img src="../../assets/img/cyborg-107.png" class="image" alt="" />
			</div>
			<div class="panel right-panel">
				<div class="content">
					<h3>振兴东北老工业基地，钢材区块链助力发展!</h3>
					<p>BIT 星之链钢材供应区块链</p>
					<button @click="switchMode()" class="btn transparent">
						去登录
					</button>
				</div>
				<img src="../../assets/img/cyborg-sign-up.png" class="image" alt="" />
			</div>
		</div>
	</div>
</template>
<script>
import { queryAccountList } from '@/api/account'
import { isValidUsername } from '@/utils/validate'
import { saveToLocal, loadFromLocal } from '@/common/local-storage'
import { mapActions } from 'vuex'

import LoginForm from "./components/LoginForm.vue"
import RegisterForm from "./components/registerForm.vue"
/* eslint-disable*/
import particles from 'particles.js'
export default {
  data() {
    return {
      signUpMode : false,
      // 粒子开关
      toggleParticles: false,
      loginForm: {
        username: '',
        pwd: ''
      },
      remember: false,
      loading: false,
    }
  },
  created() {
    // 初始化时读取localStorage用户信息
    if (loadFromLocal('remember', false)) {
      this.loginForm.username = loadFromLocal('username', '')
      this.loginForm.pwd = loadFromLocal('password', '')
    } else {
      this.loginForm.username = ''
      this.loginForm.pwd = ''
    }

      console.log("进入函数")
    	// 	queryAccountList().then(response => {
      //   this.accountList = response
      //   console.log("AccountList")
      //   console.log(response)
      // })
  },
  components: { LoginForm, RegisterForm },
  methods: {
    ...mapActions([
      'login'
    ]),
    // 登陆注册切换效果
    switchMode() {
      this.signUpMode = !this.signUpMode
      var signin = document.getElementById("login-form");
      if (getComputedStyle(signin, null)["display"] != "none") {
      signin.style.opacity = 0;
      setTimeout(function () {
        signin.style.display = "none";
      }, 1200);//这里要等待过渡的1s，然后才消失
    } else {
      setTimeout(function() {
        signin.style.opacity = 1;
      }, 0);
      signin.style.display = "block";
      signin.style.opacity = 0;
    }

      var signup = document.getElementById("signup-form").firstChild;
      if (getComputedStyle(signup, null)["display"] != "none") {
      signup.style.opacity = 0;
      setTimeout(function () {
        signup.style.display = "none";
      }, 1200);//这里要等待过渡的1s，然后才消失
    } else {
      setTimeout(function() {
        signup.style.opacity = 1;
      }, 0);
      signup.style.display = "block";
      signup.style.opacity = 0;
    }
    },
    changemode() {
      this.signUpMode = !this.signUpMode
    },
    // 用户名输入框回车后切换到密码输入框
    goToPwdInput() {
      this.$refs.pwd.$el.getElementsByTagName('input')[0].focus()
    },
    // 登录操作
    onLogin() {
      this.$refs.pwd.$el.getElementsByTagName('input')[0].blur()
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.login(this.loginForm).then(() => {
            // 保存账号
            if (this.remember) {
              // saveToLocal('username', this.loginForm.username)
              // saveToLocal('password', this.loginForm.pwd)
              // saveToLocal('remember', true)
            } else {
              // saveToLocal('username', '')
              // saveToLocal('password', '')
              // saveToLocal('remember', false)
            }
            this.$router.push({ path: '/home' })
          }).catch(() => {
            this.loading = false
          })
        } else {
          return false
        }
      })
    },
    accountTip() {
      this.$notify({
        title: '测试供应商账号：gongyingshang',
        dangerouslyUseHTMLString: true,
        message: '<strong>密码：<i>123</i></strong>',
        type: 'success',
        position: 'bottom-left'
      }),
            this.$notify({
        title: '测试银行账号：yinhang',
        dangerouslyUseHTMLString: true,
        message: '<strong>密码：<i>123</i></strong>',
        type: 'success',
        position: 'bottom-left',
        offset: 90
      }),
            this.$notify({
        title: '测试企业账号：qiye',
        dangerouslyUseHTMLString: true,
        message: '<strong>密码：<i>123</i></strong>',
        type: 'success',
        position: 'bottom-left',
        offset: 160
      })
    }
  },
  watch: {
  },
  mounted() {
    this.accountTip()
  }
}
</script>
<style scoped lang="stylus">
@import '../../css/reset.css'
</style>
