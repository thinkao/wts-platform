import axios from 'axios'

/*X-Csrf请求*/
export const csrf = (params) => {
  return axios.get(`http://127.0.0.1:8080/api/CSRFTokenAPI`, params).then(res => res)
}

/*
* 用户登录
*/
export const login = (params) => {
  return axios.post(`http://127.0.0.1:8080/api/LoginAPI`, params).then(res => res)
}
/*

/*
* 用户注册
*/
export const register = (params) => {
  return axios.post(`http://127.0.0.1:8080/api/RegistAPI`, params).then(res => res)
}

/*获取全部动态*/
export const DynamicAll = () => {
  return axios.get(`http://127.0.0.1:8080/api/DynamicAPI`).then(res => res)
}

/*发表动态*/
export const PublishDynamic = (params) => {
  return axios.post(`http://127.0.0.1:8080/api/DynamicAPI`,params).then(res => res)
}

/*function request_data(method,api,params) {

  axios({
    method: method,
    header: {
      "X-Csrftoken": instance.headers
    },
    url: instance.host+api,
    data: params
  });
}*/


