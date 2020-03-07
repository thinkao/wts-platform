import axios from 'axios'
import global from "../global";


export const request = (api,data,method) => {
    return fetchData(api,data,method)
}

function fetchData(api,data,method){
    if(method=="GET"){
        let dataStr = ''
        Object.keys(data).forEach(key => {
            if (!['', undefined, null].includes(data[key])) {
                if (Array.isArray(data[key])) {
                    data[key].forEach(item => {
                        dataStr += `${key}=${item}&`
                    })
                } else {
                    dataStr += `${key}=${data[key]}&` // 数据最终形式key1=data1&key2=data2&
                }
            }
        })
        if (dataStr !== '') {
            dataStr = dataStr.substr(0, dataStr.lastIndexOf('&'))
            api = `${api}?${dataStr}` // get方法下降url转化为url?key=data形式
        }
    }
    axios({
        method: method,
        url: api,
        data: data,
        header:{
            'X-Csrftoken':global.CsrfToken
        }
    });
}

/*
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
    }
}
*/


