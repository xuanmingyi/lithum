import request from '@/utils/request'

export function metaGet(name) {
  return request({
    url: '/meta/' + name,
    method: 'get'
  })
}
