import Vue from 'vue'
import Router from 'vue-router'

import Login from '@/views/login'
import Layout from '@/layout/layout'

Vue.use(Router)

/**
 * 关于 route 的配置属性说明：
 *
 * alwaysShow: true       // if set true, will always show the root menu, whatever its child routes length
 *                        // if not set alwaysShow, only more than ont route under the children
 *                        // it will becomes nested mode, otherwise not show the root menu
 *                        // 如果设置为true,它将总是出现在根目录。如果不设置的话，当它只有1个子路由的时候，会把
 *                        // 它的唯一子路由放到跟目录上来，而不显示它自己本身。
 *
 * hidden: true           // if set ture, 将不会出现在左侧菜单栏中
 */

/**
 * 基础路由： 任何角色都包含的路由
 * @type {Array}
 */
export const constantRouterMap = [
  {
    path: '/',
    redirect:'/login',
    hidden:true
  },
  {
    path: '/login',
    name: 'login',
    hidden: true,
    component: Login,
    meta: {
      title: '登录'
    }
  },
  {
    path: '/404',
    component: () => import('@/views/error-page/404'),
    hidden: true
  },
]

export default new Router({
  // mode: 'history',  require service support
  // scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap
})

/**
 * 动态路由： 根据用户角色
 * @type {Array}
 */
export const asyncRouterMap = [
  { path: '*', redirect: '/404', hidden: true },
  {
    path: '/company',
    component: () => import('@/views/company/company'),
    children:[
      {
        path: '/',
        name: 'Goods',
        component: () => import('../views/company/Home.vue')
      },
      {
        path: '/goods',
        name: 'Goods',
        component: () => import('../views/company/Goods.vue')
      },
      {
        path: '/orders',
        name: 'Orders',
        component: () => import('../views/company/Orders.vue')
      },
      {
        path: '/home',
        name: 'Orders',
        component: () => import('../views/company/Home.vue')
      }

    ]
  },
  {
    path: '/home',
    // hidden: true,
    component: Layout,
    redirect:'/home/home',
    children: [
      {
        path: 'home',
        name: 'home',
        component: () => import('@/views/homepage'),
        meta: {icon: 's-home', title: '首页'}
      }
    ]
  },
  {
    path: '/user',
    component: Layout,
    hidden: true,
    meta: {
      icon: 'tickets',
      title: '首页'
    },
    children: [
      {
        path: 'profile',
        name: 'Profile',
        component: () => import('@/views/homepage'),
        meta: {icon: 'warning', title: '个人中心'}
      }
    ]
  },
  {
    path: '/orderlist',
    component: Layout,
    redirect: '',
    alwaysShow: true,
    meta: {
      title: '订单管理',
      icon: 's-order'
    },
    children: [
      {
        path: 'allorder',
        name: '已确认订单',
        component: () => import('@/views/orderlist/allorder'),
        meta: {icon: 'document', title: '已确认订单'}
      },
      {
        path: 'undo',
        name: '待确认订单',
        component: () => import('@/views/orderlist/undo'),
        meta: {icon: 'document', title: '待确认订单'}
      }
    ]
  },
  {
    path: '/goods',
    component: Layout,
    redirect: '/goods/index',
    alwaysShow: true,
    meta: {
      title: '商品管理',
      icon: 's-grid'
    },
    children: [
      {
        path: 'allogoods',
        name: '全部商品',
        component: () => import('@/views/goods/allgoods'),
        meta: {icon: 'svg-layers', title: '全部商品'}
      },
      {
        path: 'uploadgoods',
        name: '上架商品',
        component: () => import('@/views/goods/uploadgoods'),
        meta: {icon: 'svg-layers', title: '上架商品'}
      }
    ]
  },
  {
    path: '/finance',
    component: Layout,
    redirect: '',
    alwaysShow: true,
    meta: {
      title: '融资管理',
      icon: 's-grid'
    },
    children: [
      {
        path: 'all',
        name: '全部融资',
        component: () => import('@/views/finance/allfinance'),
        meta: {icon: 'svg-layers', title: '全部融资'}
      },
    ]
  },
  // {
  //   path: '/bankpage',
  //   component: Layout,
  //   redirect: '/bankpage/undo',
  //   alwaysShow: true,
  //   meta: {
  //     title: '融资管理',
  //     icon: 's-grid'
  //   },
  //   children: [
  //     {
  //       path: 'undo',
  //       name: '未审批融资',
  //       component: () => import('@/views/bankpage/undo'),
  //       meta: {icon: 'svg-layers', title: '未审批融资'}
  //     },
  //     {
  //       path: 'done',
  //       name: '已审批融资',
  //       component: () => import('@/views/bankpage/'),
  //       meta: {icon: 'svg-layers', title: '已审批融资'}
  //     },
  //   ]
  // },
  {
    path: '/bankpage',
    component: Layout,
    redirect:'/bankpage/undo',
    meta: {
      icon: 'document',
      title: 'theme'
    },
    children: [
      {
        path: 'undo',
        component: () => import('@/views/bankpage/undo'),
        name: '未审批融资',
        meta: {icon: 'svg-layers', title: '未审批融资'}
      }
    ]
  },
  {
    path: '/bankpage',
    component: Layout,
    meta: {
      icon: 'document',
      title: 'theme'
    },
    children: [
      {
        path: 'done',
        component: () => import('@/views/bankpage/index'),
        name: '已审批融资',
        meta: {icon: 'svg-layers', title: '已审批融资'}
      }
    ]
  },
  
]
