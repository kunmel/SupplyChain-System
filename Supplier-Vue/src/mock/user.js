import { getToken } from '@/common/auth'

/**
 * 用户列表list
 * sex: 1-男；2-女
 * @type {Object}
 */
const userMap = {
  // 键名和token保持一致
  // permissions 是用户的权限
  // 相比于用role来做权限，permissions这样可以定制每一个用户的权限
  // permissions不能为空，最少
  supplier: {
    ID:'id1',//用户ID
    account: 'MiaoZzz', //账户名
    identity:'',//用户身份

    token: 'supplier',
    name: 'supplier',
    permissions: '/home,/user,/orderlist,/goods,/finance',
    desc: 'supplier'
  },
  bank: {
    ID:'id2',
    account: 'bank',
    token: 'bank',
    name: 'bank',
    permissions: '/bankpage,/user',
    desc: 'bank'
  },
  company: {
    ID:'id3',
    account: 'company',
    token: 'company',
    name: 'company',

    permissions: '/company',
    desc: 'bank'
  }
}

function pullUserInfo() {
  // return userMap[getToken()]
  return 0
}

export { userMap, pullUserInfo }
