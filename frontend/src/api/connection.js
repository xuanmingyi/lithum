import request from '@/utils/request'

export function connectionCreate(data) {
  return request({
    url: '/connection',
    method: 'post',
    data
  })
}

export function connectionList(params) {
  return request({
    url: '/connection',
    method: 'get',
    params
  })
}

export function connectionDelete(id) {
  return request({
    url: '/connection/' + id,
    method: 'delete'
  })
}

export function connectionUpdate(data) {
  return request({
    url: '/connection',
    method: 'put'
  })
}
