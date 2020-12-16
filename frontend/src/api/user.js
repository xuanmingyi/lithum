import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

export function logout() {
  return request({
    url: '/logout',
    method: 'post'
  })
}

export function userList(params) {
  return request({
    url: '/user',
    method: 'get',
    params
  })
}

export function me(token) {
  return request({
    url: '/me',
    method: 'get'
  })
}

