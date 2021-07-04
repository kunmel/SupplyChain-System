import axios from 'axios'
import request from '@/utils/request'

export function QueryProduct() {
    // return axios({
    //     method: 'post',
    //     url:'/company/QueryProduct'
    // })
    return request({
      url: '/api/v1/QueryProduct',
      method: 'post'
    })
    // return axios({
    //     method: 'post',
    //     url:'QueryProduct'
    // })
}


// 获取登录界面角色选择列表
export function QueryProductTest() {
  return request({
    url: '/api/v1/QueryProduct',
    method: 'post'
  })
  // return request({
  //   url: '/accountlist.json',
  //   method: 'get'
  // })
}

export function StoreOrder(data) {
    // return axios({
    //     method:'post',
    //     url:'/company/StoreOrder',
    //     orderData
    // })
    return request({
      url: '/api/v1/StoreOrder',
      method: 'post',
      data
    })
    // return axios({
    //     method:'post',
    //     url:'/StoreOrder',
    //     orderData
    // })
}

export function QueryOrderByBuyer(data) {
  console.log("QueryOrderByBuyer canshu",data)
  return request({
    url: '/api/v1/QueryOrderByBuyer',
    method: 'post',
    data
    
  })
    // return axios({
    //     method: 'post',
    //     url: '/QueryOrderByBuyer',
    //      companyID
    //   })
  }

export function QueryOrder() {
    return request({
      url: '/api/v1/QueryOrder',
      method: 'post'
    })
    // return axios({
    //     method: 'post',
    //     url:'QueryProduct'
    // })
}