<template>
  <div>

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
      class="table"
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
            size="mini"
            @click="handleEdit(scope.$index, scope.row)"
          >编辑</el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)"
          >删除</el-button>
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
import { loadComponent } from '@/utils/vue-loader'

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
    this.fetch()
  },
  methods: {
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
    create() {
    },
    delete() {
    },
    handlePageSizeChange() {},
    handlePageCurrentChange() {},
    handleTableAction(action) {
      import('@/components/LDialog').then(cmp => {
        loadComponent.call(this, cmp, {}, '.table')
      })
    },
    getColumnLabel(name) {
      for (let i = 0; i < this.meta.model.attributes.length; i++) {
        if (this.meta.model.attributes[i].name === name) {
          return this.meta.model.attributes[i].display
        }
      }
    },
    getColumnValue(row, name) {
      const func_name = 'getColumnValue' + name
      if (func_name in this) {
        return this[func_name](row)
      } else {
        return row[name]
      }
    },
    handleRowAction(action) {
      console.log(action)
    },
    handleEdit(index, row) {
      import('@/components/LDialog').then(cmp => {
        loadComponent.call(this, cmp, {}, document.querySelector('.table'))
      })
    },
    handleDelete(index, row) {
      this.$confirm('是否继续删除操作?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteItem(row).then(response => {
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
