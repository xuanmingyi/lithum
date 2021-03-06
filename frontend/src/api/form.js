import request from '@/utils/request'

export function formCreate(data) {
  return request({
    url: '/form',
    method: 'post',
    data
  })
}

export function formList(params) {
  return request({
    url: '/form',
    method: 'get',
    params
  })
}

export function formDelete(id) {
  return request({
    url: '/form/' + id,
    method: 'delete'
  })
}
