<template>
  <div class="app-container">
    <div class="table-header">
      <div class="table-search">
        <el-input v-model="search" placeholder="搜索内容" @keyup.enter.native="fetchData" />
      </div>
      <div class="table-actions">
        <el-button type="primary" @click="createDialog.visible = true">创建</el-button>
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
      <el-table-column label="用户名" align="center">
        <template slot-scope="scope">
          {{ scope.row.username }}
        </template>
      </el-table-column>
      <el-table-column label="昵称" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.nickname }}</span>
        </template>
      </el-table-column>
      <el-table-column label="介绍" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.introduction }}</span>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.created_at | timeFilter }}</span>
        </template>
      </el-table-column>
      <el-table-column label="更新时间" align="center">
        <template slot-scope="scope">
          {{ scope.row.updated_at | timeFilter }}
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center">
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
          <el-dropdown class="el-button el-button--mini more" @command="handleCommand">
            <span class="el-dropdown-link">
              更多操作<i class="el-icon-arrow-down el-icon--right" />
            </span>
            <el-dropdown-menu slot="dropdown" class="more-items">
              <el-dropdown-item command="a">黄金糕</el-dropdown-item>
              <el-dropdown-item command="b">狮子头</el-dropdown-item>
              <el-dropdown-item command="c">螺蛳粉</el-dropdown-item>
              <el-dropdown-item command="d" disabled>双皮奶</el-dropdown-item>
              <el-dropdown-item command="e" divided>蚵仔煎</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
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
    <el-dialog title="创建用户" :visible.sync="createDialog.visible">
      <el-form :model="createDialog">
        <el-form-item label="用户名" :label-width="createDialog.labelWidth">
          <el-input v-model="createDialog.username" autocomplete="off" />
        </el-form-item>
        <el-form-item label="昵称" :label-width="createDialog.labelWidth">
          <el-input v-model="createDialog.nickname" autocomplete="off" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="createDialog.visible = false">取 消</el-button>
        <el-button type="primary" @click="createDialog.visible = false">确 定</el-button>
      </div>
    </el-dialog>
    <el-dialog title="更新用户" :visible.sync="updateDialog.visible">
      <el-form :model="updateDialog">
        <el-form-item label="用户名" :label-width="updateDialog.labelWidth">
          <el-input v-model="updateDialog.username" autocomplete="off" disabled />
        </el-form-item>
        <el-form-item label="昵称" :label-width="updateDialog.labelWidth">
          <el-input v-model="updateDialog.nickname" autocomplete="off" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="updateDialog.visible = false">取 消</el-button>
        <el-button type="primary" @click="updateDialog.visible = false">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { userList } from '@/api/user'

export default {
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
      createDialog: {
        visible: false,
        username: '',
        nickname: '',
        labelWidth: '100'
      },
      updateDialog: {
        visible: false,
        username: '',
        nickname: '',
        labelWidth: '100'
      },
      list: null,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      console.log(this.search)
      this.listLoading = true
      const req = {
        username: this.search,
        page: this.pagination.current,
        limit: this.pagination.size
      }
      userList(req).then(response => {
        this.pagination.total = response.data.count
        this.list = response.data.items
        this.listLoading = false
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
    handleDelete(index, row) {
      this.$confirm('此操作将删除 ' + row.username + ' , 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$message({
          type: 'success',
          message: '删除成功!'
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    handleCommand() {
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

.more-items {
  width: 96px;
}
</style>
