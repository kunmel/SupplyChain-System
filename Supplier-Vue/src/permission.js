/*
 * 全局权限检测
 * 包括1、路由的全局守卫
 */
import router from './router'
import store from './store'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'// progress bar style
import { getToken } from '@/common/auth'

NProgress.configure({ showSpinner: false })

// 路由全局前置守卫
const whiteList = ['/login','/404'] // 白名单
router.beforeEach((to, from, next) => {
  //NProgress.start() // start progress bar
  console.log("getToken",getToken())
  if (getToken()) {
    // alert("有Token"+getToken())
    // alert("有token")
    // 有token并且不是company，访问login页面，就跳到首页
    if (to.path == '/'||to.path == '/login') {
      store.dispatch('logout').then(() => {
        next()
      })
      //NProgress.done() // 这种情况不会触发router的后置钩子，所以这里需要单独处理
    } 
    else {
      //alert("跳到else")
    //   //NProgress.done()
    //   // 有token，没有permissions
      var permissions
      // store.dispatch('pullUserInfo').then(resp => {
      //   permissions = resp.permissions
      //   const token = resp.token
      //   // alert("permission： "+permissions+"token: "+token)
      //   // alert("store permission:"+store.getters.permissions)

      // })
      
      // permissions = store.getters.permissions
      // const token = store.getters.token

      console.log("permissions:",permissions)
      console.log("store.getters.permissions",store.getters.permissions)
      // if (store.getters.permissions.length === 0 || store.getters.permissions != permissions) {
        if (store.getters.permissions.length === 0) {
        // if  (permissions.indexOf(to.path) == -1){
        console.log("进入1")
        store.dispatch('pullUserInfo').then(resp => {
          const permissions = store.getters.permissions
          store.dispatch('GenerateRoutes', { permissions }).then(() => {
            // 动态添加可访问路由表
            router.addRoutes(store.getters.addRouters)
            console.log("可访问路由",store.getters.addRouters)
            // console.log({...to})
            // hack方法 确保addRoutes已完成，set the replace: true so the navigation will not leave a history record
            // 这样我们就可以简单的通过 `next(to)` 巧妙的避开之前的那个问题了。这行代码重新进入 `router.beforeEach` 这个钩子，这时候再通过 `next()` 来释放钩子，就能确保所有的路由都已经挂载完成了。
            next({
              ...to,
              replace: true
            })
          })
        }).catch(() => {
          alert("出大问题了")
          store.dispatch('logout').then(() => {
            next('/login')
          })
        })
      } else {
        console.log("进入2")
        next()
      }
    }
  } else {
    if (whiteList.includes(to.path)) { // 白名单，免密登录
      next()
    } else { // 否则就跳动登录页面
      next('/login')
     
    }
  }
})

router.afterEach(() => {
})
