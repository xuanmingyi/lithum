<template>
  <div class="app-container" />
</template>

<script>
import { loadComponent } from '@/utils/vue-loader'
import { metaGet } from '@/api/meta'
import { connectionList, connectionCreate, connectionDelete } from '@/api/connection'

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
          fetchItems() {
            // 获取数据
            const req = {
              search: this.search,
              page: this.page,
              limit: this.limit
            }
            return connectionList(req)
          },
          createItem(fields) {
            // 创建数据
            console.log(fields)
            const req = {}
            return connectionCreate(req)
          },
          deleteItem(row) {
            const req = {}
            return connectionDelete(req)
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
