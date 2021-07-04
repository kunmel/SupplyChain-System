import axios from 'axios'
import request from '@/utils/request'
// import jsonp from '@/common/jsonp'
import { isDev } from '@/utils'

// 获取当前运行环境
// const isDev = getEnviroument()

export function getTable() {
  if(isDev) {
    // 开发环境有代理
    return axios.post('/api/getmoviepiaofang')
  } else {
    // 使用 mock 拦截请求来造数据
    // setTimeout(function() {
    //   return axios.post('/api/getmoviepiaofang-mock')
    // }, 2000)
    return new Promise((resolve, reject) => {
      setTimeout(resolve, 2000, axios.post('/api/getmoviepiaofang-mock'))
    })
    // return axios.post('/api/getmoviepiaofang-mock')
  }
}
//获取借贷信息
export function getUserFiance() {
  return axios({
    method: 'post',
    url: '/excel/userFinance'
  })
}

// 合并&统计 - 表格内容
export function getMergeTable() {
  return axios({
    method: 'post',
    url: '/excel/getMergeTableData'
  })
}

export function QueryProductBySupplier(data) {
  console.log("QueryProductBySupplier canshu",data)
  return request({
    url: '/api/v1/QueryProductBySupplier',
    method: 'post',
    data
  })
}

export function UpdateProductByAmount(data) {
  console.log("UpdateProductByAmount canshu",data)
  return request({
    url: '/api/v1/UpdateProductByAmount',
    method: 'post',
    data
  })
}

export function StoreProduct(data) {
  console.log("StoreProduct canshu",data)
  return request({
    url: '/api/v1/StoreProduct',
    method: 'post',
    data
  })
}
    
export function DeleteProduct(data) {
  console.log("DeleteProduct canshu",data)
  return request({
    url: '/api/v1/DeleteProduct',
    method: 'post',
    data
  })
}

export function QueryOrderBySeller(data) {
  console.log("QueryOrderBySeller canshu",data)
  return request({
    url: '/api/v1/QueryOrderBySeller',
    method: 'post',
    data
  })
}

export function UpdateStatus(data) {
  console.log("UpdateStatus canshu",data)
  return request({
    url: '/api/v1/UpdateStatus',
    method: 'post',
    data
  })
}

export function UpdateTransferStatus(data) {
  console.log("UpdateTransferStatus canshu",data)
  return request({
    url: '/api/v1/UpdateTransferStatus',
    method: 'post',
    data
  })
}

export function StoreFinance(data) {
  console.log("StoreFinance canshu",data)
  return request({
    url: '/api/v1/StoreFinance',
    method: 'post',
    data
  })
}

export function QueryFinanceBySupplier(data) {
  console.log("QueryFinanceBySupplier canshu",data)
  return request({
    url: '/api/v1/QueryFinanceBySupplier',
    method: 'post',
    data
  })
}

export function QueryFinance() {
  console.log("QueryFinance canshu")
  return request({
    url: '/api/v1/QueryFinance',
    method: 'post'
  })
}



export function UpdateFinanceStatus(data) {
  console.log("UpdateFinanceStatus canshu",data)
  return request({
    url: '/api/v1/UpdateFinanceStatus',
    method: 'post',
    data
  })
}

export function getUndoTable() {
  return axios({
    method: 'post',
    url: '/excel/getUndoTable'
  })
}

export function getGoodsTable() {
  return axios({
    method: 'post',
    url: '/excel/getGoodsTable'
  })
}
/**
 * 自定义表格
 */
// 获取文件列表
export function getFiles() {
  return axios({
    method: 'post',
    url: '/excel/getFiles'
  })
}

// 删除选择的文件
export function delFiles(ids) {
  return axios({
    method: 'post',
    url: '/excel/delFiles'
  })
}
