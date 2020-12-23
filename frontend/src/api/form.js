import request from '@/utils/request'

export function form_user(data) {
  return request({
    url: '/form/user',
    method: 'get'
  })
}
