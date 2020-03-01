import Vue from 'vue'
import dayjs from 'dayjs'

const filters = {
  dateTimeFormat (_timestamp, _format = 'YYYY-MM-DD HH:mm:ss') {
    if (!_timestamp) return '-'
    return dayjs(_timestamp).format(_format)
  },
  dateFormat (_timestamp, _format = 'YYYY-MM-DD') {
    if (!_timestamp) return '-'
    return dayjs(_timestamp).format(_format)
  }
}

// 注册全局通用过滤器
for (let filter in filters) {
  Vue.filter(filter, filters[filter])
}
