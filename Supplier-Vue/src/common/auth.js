import Cookies from 'js-cookie'
// import Mock from 'mockjs'

const TokenKey = 'Vue-cms'

function getToken() {
  return Cookies.get(TokenKey)
}

function setToken(token) {
  return Cookies.set(TokenKey, token)
}

function removeToken() {
  return Cookies.remove(TokenKey)
}

const UserIdentity = 'UserIdentity'

function getUserIdentity() {
  return Cookies.get(UserIdentity)
}

function setUserIdentity(identity) {
  return Cookies.set(UserIdentity, identity)
}

function removeUserIdentity() {
  return Cookies.remove(UserIdentity)
}

const UserID = 'UserID'

function getUserID() {
  return Cookies.get(UserID)
}

function setUserID(userid) {
  return Cookies.set(UserID, userid)
}

function removeUserID() {
  return Cookies.remove(UserID)
}

const UserName = 'UserName'

function getUserName() {
  return Cookies.get(UserName)
}

function setUserName(username) {
  return Cookies.set(UserName, username)
}

function removeUserName() {
  return Cookies.remove(UserName)
}

export { getToken, setToken, removeToken,getUserID,setUserID,removeUserID,getUserIdentity,setUserIdentity,removeUserIdentit,getUserName,setUserName,removeUserName }
