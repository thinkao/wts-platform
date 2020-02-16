import axios from 'axios'

/*
* 用户登录
*/
export const login = (params) => {
  return axios.post(`http://127.0.0.1:8080/api/LoginAPI`, params).then(res => res)
}

/*
* 用户注册
*/
export const register = (params) => {
  return axios.post(`http://127.0.0.1:8080/RegistAPI`, params).then(res => res)
}
