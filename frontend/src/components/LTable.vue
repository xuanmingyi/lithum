<template>
  <div
    class="table"
  >
    <div class="table-header">
      <div class="table-search">
        <el-input v-model="search" size="mini" placeholder="输入关键字搜索" @keyup.enter.native="fetch" />
      </div>
      <div class="table-actions">
        <el-button
          v-for="action in meta.table.actions"
          :key="action.name"
          type="primary"
          size="mini"
          @click="handleTableAction(action)"
        >{{ action.display }}</el-button>
      </div>
    </div>

    <el-table
      :data="list"
      style="width: 100%"
      border
    >
      <el-table-column
        v-for="column in meta.table.columns"
        :key="column.name"
        :label="getColumnLabel(column.name)"
        align="center"
      >
        <template slot-scope="scope">
          {{ getColumnValue(scope.row, column.name) }}
        </template>
      </el-table-column>

      <el-table-column
        align="center"
        label="操作"
      >
        <template slot-scope="scope">
          <el-button
            v-for="action in meta.table.row_actions"
            :key="action.name"
            size="mini"
            :type="computedType(action)"
            @click="handleRowAction(action, scope.$index, scope.row)"
          >{{ action.display }}</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      class="table-footer"
      :current-page="pagination.current"
      :page-sizes="[10, 20, 30, 50]"
      :page-size="10"
      :total="pagination.total"
      layout="total, sizes, prev, pager, next, jumper"
      @size-change="handlePageSizeChange"
      @current-change="handlePageCurrentChange"
    />
  </div>
</template>

<script>
import { loadComponent, initMeta } from '@/utils/vue-loader'

export default {
  name: 'LTable',
  data() {
    return {
      // 当前是否正在加载
      listLoading: false,

      // 分页
      pagination: {
        total: 0,
        current: 1,
        size: 10
      },

      // 列表数据
      list: [],

      // 搜索
      search: ''
    }
  },
  created() {
    // init meta
    initMeta(this.meta)

    this.$bus.on('create', this.create)
    this.$bus.on('delete', this.delete)
    this.$bus.on('update', this.update)

    this.fetch()
  },
  beforeDestroy() {
    this.$bus.off('create')
    this.$bus.off('delete')
    this.$bus.off('update')
  },
  methods: {
    create(obj) {
      // console.log('create hahahaha ')
      // console.log(this)
    },
    delete(id) {
    },
    update(obj) {
    },
    fetch() {
      this.listLoading = true
      this.fetchItems().then(response => {
        this.pagination.total = response.data.count
        this.list = response.data.items
        this.listLoading = false
      }).catch(reason => {
        console.log(reason)
      })
    },
    handlePageSizeChange() {
    },
    handlePageCurrentChange() {
    },
    getColumnLabel(name) {
      // 获取table的标题
      for (let i = 0; i < this.meta.model.attributes.length; i++) {
        if (this.meta.model.attributes[i].name === name) {
          return this.meta.model.attributes[i].display
        }
      }
    },
    getColumnValue(row, name) {
      // 获取table中的数据内容，相当于filter
      const func_name = 'getColumnValue' + name
      if (func_name in this) {
        return this[func_name](row)
      } else {
        return row[name]
      }
    },
    handleTableAction(action) {
      if (action.type === 'dialog') {
        const meta = this.meta
        const fetchItems = this.fetchItems
        const createItem = this.createItem
        const deleteItem = this.deleteItem
        const mixin = {
          data: function() {
            return {
              meta: meta,
              action: action,
              fetchItems: fetchItems,
              createItem: createItem,
              deleteItem: deleteItem
            }
          }
        }
        import('@/components/LDialog').then(cmp => {
          loadComponent.call(this, cmp, mixin, '.table')
        })
      }
    },
    handleRowAction(action, index, row) {
      if (action.type === 'dialog') {
        const mixin = {
          data: function() {
            return {
              meta: this.meta,
              action: action,
              index: index,
              row: row,
              fetchItems: this.fetchItems,
              createItem: this.createItem,
              deleteItem: this.deleteItem
            }
          }
        }
        import('@/components/LDialog').then(cmp => {
          loadComponent.call(this, cmp, mixin, '.table')
        })
      } else if (action.type === 'confirm') {
        this.$confirm(action.confirm.message, action.confirm.title, {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this[action.func](row).then(response => {
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
            this.fetch()
          }).catch(() => {
            this.$message({
              type: 'info',
              message: '删除失败!'
            })
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
      }
    },
    computedType(action) {
      if (action.name === 'delete') {
        return 'danger'
      } else {
        return ''
      }
    }
  }
}
</script>

<style scoped>
.table-header {
  margin: 10px 0px;
}

.table-search {
  display: inline-block;
  width: 40%;
}

.table-actions {
  float: right;
}

.table-footer {
  float: right;
  margin: 20px 20px;
}
</style>
