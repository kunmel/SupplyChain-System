import request from '@/utils/request'

// 获取登录界面角色选择列表
export function queryAccountList() {
  return request({
    url: '/api/v1/queryAccountList',
    method: 'post'
  })
  // return request({
  //   url: '/accountlist.json',
  //   method: 'get'
  // })
}


export function QueryUserTest(data) {
  return request({
    url: '/api/v1/QueryUserTest',
    method: 'post',
    data
  })

}

export function QueryUser(data) {
  console.log("queryuser canshu",data)
  return request({
    url: '/api/v1/QueryUser',
    method: 'post',
    data
  })

}

export function StoreUser(data) {
  return request({
    url: '/api/v1/StoreUser',
    method: 'post',
    data
  })

}

// 登录
export function login(data) {
  return request({
    url: '/api/v1/QueryUser',
    method: 'post',
    data
  })
}
