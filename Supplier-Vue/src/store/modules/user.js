import { register, login, userInfo, logout } from '@/api/login'
import {QueryUser} from '@/api/account'
import { getToken, setToken, removeToken,getUserID,setUserID,removeUserID,getUserIdentity,setUserIdentity,removeUserIdentity,setUserName } from '@/common/auth'

const SET_ACCOUNT = 'SET_ACCOUNT'
const SET_TOKEN = 'SET_TOKEN'
const SET_NAME = 'SET_NAME'
const SET_IDENTITY = 'SET_IDENTITY'
const SET_AGE = 'SET_AGE'
const SET_SEX = 'SET_AEX'
const SET_AVATAR = 'SET_AVATAR'
const SET_PERMISSIONS = 'SET_PERMISSIONS'
const SET_TYPE = 'SET_TYPE'
const SET_DESC = 'SET_DESC'
const SET_ALL = 'SET_ALL'

const user = {
  state: {
    token: getToken(),
    identity: 0,
    account: '',
    name: '',
    age: 0,
    sex: '',
    avatar: '',
    permissions: '',
    type: [],
    desc: ''
  },
  mutations: {
    [SET_ACCOUNT](state, account) {
      state.account = account
    },
    [SET_IDENTITY](state, identity) {
      state.identity = identity
    },
    [SET_TOKEN](state, token) {
      state.token = token
    },
    [SET_NAME](state, name) {
      state.name = name
    },
    [SET_PERMISSIONS](state, permissions) {
      state.permissions = permissions
    },
    [SET_TYPE](state, type) {
      state.type = type
    },
    [SET_DESC](state, desc) {
      state.desc = desc
    },
    [SET_ALL](state, userInfo) {
      state = Object.assign(state, userInfo)
    }
  },
  actions: {
    // 用户登录
    login({ commit }, data) {
      return new Promise((resolve, reject) => {
        console.log('logindata',data)
        QueryUser(data).then(resp => {
          console.log("login",resp)
          let data = resp
          setToken(data.Account)
          setUserID(data.ID)
          setUserIdentity(data.Identity)
          setUserName(data.Account)
          commit(SET_TOKEN, data.Account)
          commit(SET_ACCOUNT, data.ID)
          commit(SET_IDENTITY, data.Identity)
          // commit(SET_NAME, data.name)
          // commit(SET_AGE, data.age)
          // commit(SET_AVATAR, data.avatar)
          // commit(SET_PERMISSIONS, data.permissions)
          return resolve(data)
        }).catch(err => {
          console.log("err",err)
          return reject(err)
        })
      })
    },  // 用户注册
    register({ commit }, userInfo) {
      return new Promise((resolve, reject) => {
        register(userInfo).then(resp => {
          //alert("用户注册成功")
          let data = resp.data
          setToken(data.token)
          commit(SET_TOKEN, data.token)
          // commit(SET_NAME, data.name)
          // commit(SET_AGE, data.age)
          // commit(SET_AVATAR, data.avatar)
          // commit(SET_PERMISSIONS, data.permissions)
          return resolve()
        }).catch(err => {
          return reject(err)
        })
      })
    },
    // 拉取用户信息
    pullUserInfo({ commit }) {
      return new Promise((resolve, reject) => {
        userInfo().then(resp => {
          let identity = getUserIdentity()
          if(identity == 1){
            var permissions = '/company'
            commit(SET_PERMISSIONS, permissions)
          }else if(identity == 2){
            var permissions =  '/home,/user,/orderlist,/goods,/finance'
            commit(SET_PERMISSIONS, permissions)
          }else if(identity == 3){
            var permissions = '/bankpage,/user'
            commit(SET_PERMISSIONS, permissions)
          }

          // commit(SET_ACCOUNT, data.account)
          // commit(SET_NAME, data.name)
          // commit(SET_PERMISSIONS, data.permissions)
          // commit(SET_TYPE, data.type)
          // commit(SET_DESC, data.desc)
          return resolve()
        }).catch(err => {
          return reject(err)
        })
      })
    },
    // 用户退出登录
    logout({ commit }) {
      return new Promise((resolve, reject) => {
        logout().then(resp => {
          removeToken()
          commit(SET_TOKEN, '')
          commit(SET_NAME, '')
          return resolve()
        }).catch(err => {
          return reject(err)
        })
      })
    },
    // 头像更新
    doUpdateAvatar({ commit }, imgFile) {
      return new Promise(resolve => {
        setTimeout(() => {
          commit(SET_AVATAR, imgFile)
          resolve()
        }, 1000)
      })
    },
    /**
     * 更新用户信息
     * userInfo: 用户信息表对象
     */
    doUpdateUser({ commit }, userInfo) {
      return new Promise(resolve => {
        commit(SET_ALL, userInfo)
        setTimeout(() => {
          resolve()
        }, 1000)
      })
    }
  },
  getters: {
    token: state => state.token,
    identity : state => state.identity,
    account : state => state.account,
    name: state => state.name,
    age: state => state.age,
    avatar: state => state.avatar,
    permissions: state => state.permissions,
    allInfo: state => state
  }
}

export default user
