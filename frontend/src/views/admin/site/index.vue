<template>
  <div class="app-container" />
</template>

<script>
import { loadComponent } from '@/utils/vue-loader'
import { formList, formDelete } from '@/api/form'
import { metaGet } from '@/api/meta'

export default {
  created() {
    metaGet('form').then(response => {
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
            return formList(req)
          },
          createItem() {
            // 创建数据
          },
          deleteItem(row) {
            // 删除数据
            return formDelete(row.id)
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
