<template>
  <div class="app-container">
    <div class="table-header">
      <div class="table-search">
        <el-input v-model="search" placeholder="搜索内容" @keyup.enter.native="fetchData" />
      </div>
      <div class="table-actions">
        <el-button type="primary" @click="createFormVisible = true">创建</el-button>
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

    <el-dialog v-bind="$attrs" title="创建表单" :visible="createFormVisible" v-on="$listeners" @open="onOpen" @close="onClose">
      <el-form
        ref="createForm"
        :model="createFormData"
        :rules="rules"
        size="medium"
        label-width="100px"
        label-position="left"
      >
        <el-form-item label="表单名" prop="Title">
          <el-input v-model="createFormData.Title" placeholder="请输入表单名" clearable :style="{width: '100%'}" />
        </el-form-item>
        <el-form-item label="表单尺寸" prop="Size">
          <el-radio-group v-model="createFormData.Size" size="medium">
            <el-radio-button
              v-for="(item, index) in SizeOptions"
              :key="index"
              :label="item.value"
              :disabled="item.disabled"
              border
            >{{ item.label }}</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="标签对齐" prop="LabelPosition">
          <el-radio-group v-model="createFormData.LabelPosition" size="medium">
            <el-radio-button
              v-for="(item, index) in LabelPositionOptions"
              :key="index"
              :label="item.value"
              :disabled="item.disabled"
              border
            >{{ item.label }}</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="标签宽度" prop="LabelWidth">
          <el-input-number v-model="createFormData.LabelWidth" placeholder="标签宽度" />
        </el-form-item>
        <el-form-item label="栅格间隔" prop="Span">
          <el-input-number v-model="createFormData.Span" placeholder="栅格间隔" />
        </el-form-item>
      </el-form>
      <div slot="footer">
        <el-button @click="close">取消</el-button>
        <el-button type="primary" @click="handelConfirm">确定</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
import { formList, formCreate } from '@/api/form'

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
      createFormVisible: false,
      createFormData: {
        Title: undefined,
        Size: 'medium',
        LabelPosition: 'left',
        LabelWidth: 100,
        Span: 15
      },
      rules: {
        Title: [{
          required: true,
          message: '请输入表单名',
          trigger: 'blur'
        }],
        Size: [{
          required: true,
          message: '表单尺寸不能为空',
          trigger: 'change'
        }],
        LabelPosition: [{
          required: true,
          message: '标签对齐不能为空',
          trigger: 'change'
        }],
        LabelWidth: [{
          required: true,
          message: '标签宽度',
          trigger: 'blur'
        }],
        Span: [{
          required: true,
          message: '栅格间隔',
          trigger: 'blur'
        }]
      },
      SizeOptions: [{
        'label': '中等',
        'value': 'medium'
      }, {
        'label': '较小',
        'value': 'small'
      }, {
        'label': '迷你',
        'value': 'mini'
      }],
      LabelPositionOptions: [{
        'label': '左对齐',
        'value': 'left'
      }, {
        'label': '右对齐',
        'value': 'right'
      }, {
        'label': '顶部对齐',
        'value': 'top'
      }],
      list: null,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    onOpen() {},
    onClose() {
      this.$refs['createForm'].resetFields()
    },
    close() {
      this.createFormVisible = false
    },
    handelConfirm() {
      this.$refs['createForm'].validate(valid => {
        if (!valid) return
        formCreate(this.createFormData).then(response => {
          this.fetchData()
        })
        this.close()
      })
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
