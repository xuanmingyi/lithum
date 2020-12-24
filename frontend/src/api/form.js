import request from '@/utils/request'

export function formCreate(data) {
  return request({
    url: '/form',
    method: 'post',
    data
  })
}
