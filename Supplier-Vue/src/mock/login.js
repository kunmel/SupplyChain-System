// import Mock from 'mockjs'
import { userMap } from './user'

function login(userInfo) {
  let { username } = JSON.parse(userInfo.body)
  return userMap[username]
  
}

function logout() {
  return 'success'
}

function register(userInfo) {
  return userMap['supplier']
}

export { login, logout ,register}
