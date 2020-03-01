import axios from 'axios'
import qs from 'qs'

/*X-Csrf请求*/
export const csrf = () => fetchData('/api/CSRFTokenAPI',null,'GET')

/*用户登录*/
export const login = (data) => fetchData('/api/LoginAPI',data,'POST')

/*用户注册*/
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

export const request = (api,data,method) => {
    return fetchData(api,data,method)
}

const fetchData = (url = '', data = {}, method = 'GET') => {
    method = method.toUpperCase() // 将字符串转换为大写
    let reqdata = {}
    Object.keys(data).forEach(key => {
        if (!['', undefined, null, '全部'].includes(data[key])) {
            reqdata[key] = data[key]
        }
    })
    // 解析data = {} 里面的参数，转成JSONP格式
    if (method === 'GET') {
        let dataStr = ''
        Object.keys(reqdata).forEach(key => {
            if (!['', undefined, null].includes(reqdata[key])) {
                if (Array.isArray(reqdata[key])) {
                    reqdata[key].forEach(item => {
                        dataStr += `${key}=${item}&`
                    })
                } else {
                    dataStr += `${key}=${reqdata[key]}&` // 数据最终形式key1=data1&key2=data2&
                }
            }
        })
        if (dataStr !== '') {
            dataStr = dataStr.substr(0, dataStr.lastIndexOf('&'))
            url = `${url}?${dataStr}` // get方法下降url转化为url?key=data形式
        }
        // 尝试从远程获取数据
        try {
            // 异步操作
            return new Promise(resolve => {
                axios.get(encodeURI(url)).then(response => {
                    const responseData = response.data
                    resolve(responseData)
                })
            })
        } catch (error) {
            throw new Error(error)
        }
    } else {
        // 非get方法，尝试获取数据
        let requestData
        if (method === 'POST' || method === 'PUT') {
            requestData = qs.stringify(reqdata)
        } else {
            requestData = reqdata
        }
        try {
            // 异步操作
            return new Promise(resolve => {
                let reqParam = {
                    method: method,
                    url: url,
                    data: requestData,
                }
                if (url === '/api/UploadChunkAPI') {
                    reqParam.headers = { 'Content-Type': 'multipart/form-data' }
                    reqParam.data = data
                }
                axios(reqParam).then(response => {
                    const responseData = response.data
                    resolve(responseData)
                })
            })
        } catch (error) {
            throw new Error(error)
        }
    }
}


