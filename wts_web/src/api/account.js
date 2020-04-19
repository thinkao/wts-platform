import axios from 'axios'
import global from "../global";

export const request = (api,data,method) => {
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
    return  axios({
        method: method,
        url: api,
        data: data,
        headers:{
            'X-Csrftoken':global.CsrfToken
        }
    });

}



