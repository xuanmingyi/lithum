<template>
  <div class="app-container" />
</template>

<script>
import { connectionList, connectionCreate, connectionDelete, connectionUpdate } from '@/api/connection'
import { loadComponent } from '@/utils/vue-loader'
import { metaGet } from '@/api/meta'

export default {
  created() {
    metaGet('connection').then(response => {
      const mixin = {
        data: function() {
          return {
            meta: response.data
          }
        },
        methods: {
          createItem(fields) {
            console.log(fields)
            // 创建数据
            const req = {}
            return connectionCreate(req)
          },
          deleteItem(row) {
            // 删除数据
            const req = {}
            return connectionDelete(req)
          },
          updateItem(row) {
            // 修改数据
            const req = {}
            return connectionUpdate(req)
          },
          fetchItems() {
            // 获取数据
            const req = {
              search: this.search,
              page: this.page,
              limit: this.limit
            }
            return connectionList(req)
          }
        }
      }
      import('@/components/LTable').then(cmp => {
        loadComponent.call(this, cmp, mixin, '.app-container')
      })
    }).catch(reason => {
      console.log(reason)
    })
  }
}
</script>
