<template>
  <div class="app-container">
    <div class="table-header">
      <div class="table-search">
        <el-input v-model="search" placeholder="搜索内容" @keyup.enter.native="fetchData" />
      </div>
      <div class="table-actions">
        <el-button type="primary" @click="openCreateDialog">创建</el-button>
      </div>
    </div>

    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="ID" width="95">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column label="标题" align="center">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column label="尺寸" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.size }}</span>
        </template>
      </el-table-column>
      <el-table-column label="标签对齐" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.labelPosition }}</span>
        </template>
      </el-table-column>
      <el-table-column label="标签宽度" align="center">
        <template slot-scope="scope">
          {{ scope.row.labelWidth }}
        </template>
      </el-table-column>
      <el-table-column label="Span" align="center">
        <template slot-scope="scope">
          {{ scope.row.span }}
        </template>
      </el-table-column>
      <el-table-column label="gutter" align="center">
        <template slot-scope="scope">
          {{ scope.row.gutter }}
        </template>
      </el-table-column>
      <el-table-column label="formBtns" align="center">
        <template slot-scope="scope">
          {{ scope.row.formBtns }}
        </template>
      </el-table-column>
      <el-table-column label="disabled" align="center">
        <template slot-scope="scope">
          {{ scope.row.disabled }}
        </template>
      </el-table-column>
      <el-table-column label="field数量" align="center">
        <template slot-scope="scope">
          11
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="primary"
            @click="handleEdit(scope.$index, scope.row)"
          >编辑</el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)"
          >删除</el-button>
          <router-link :to="`/admin/form/${scope.row.id}`">
            <el-button
              size="mini"
              @click="handleView(scope.$index, scope.row)"
            >查看</el-button>
          </router-link>
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
    <Dialog ref="dialog" @fetch="fetchData()" />
  </div>
</template>

<script>
import { formList, formDelete } from '@/api/form'
import { default as Dialog } from './dialog'

export default {
  components: {
    'Dialog': Dialog
  },
  filters: {
    timeFilter(time) {
      var delta = (Date.parse(new Date()) - Date.parse(time)) / 1000
      if (delta < 60) {
        return delta + ' 秒前'
      }
      if (delta < 3600) {
        return parseInt(delta / 60) + ' 分钟前'
      }
      if (delta < 86400) {
        return parseInt(delta / 3600) + ' 小时前'
      }
      return parseInt(delta / 86400) + ' 天前'
    }
  },
  data() {
    return {
      search: '',
      pagination: {
        total: 0,
        current: 1,
        size: 10
      },
      list: null,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    openCreateDialog() {
      this.$refs['dialog'].open()
    },
    fetchData() {
      this.listLoading = true
      const req = {
        search: this.search,
        page: this.pagination.current,
        limit: this.pagination.size
      }
      formList(req).then(response => {
        this.pagination.total = response.data.count
        this.list = response.data.items
        this.listLoading = false
      }).catch(reason => {
        console.log(reason)
      })
    },
    handlePageSizeChange(val) {
      this.pagination.size = val
      this.pagination.current = 1
      this.fetchData()
    },
    handlePageCurrentChange(val) {
      this.pagination.current = val
      this.fetchData()
    },
    handleEdit(index, row) {
      this.updateDialog.username = row.username
      this.updateDialog.nickname = row.nickname
      this.updateDialog.visible = true
    },
    handleView(index, row) {
      console.log(row)
    },
    handleDelete(index, row) {
      this.$confirm('此操作将删除 ' + row.title + ' , 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        formDelete(row.id).then(response => {
          this.$message({
            type: 'success',
            message: '删除成功!'
          })
          this.fetchData()
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
