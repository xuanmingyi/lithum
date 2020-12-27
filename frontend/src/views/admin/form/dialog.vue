<template>
  <el-dialog v-bind="$attrs" title="创建表单" :visible="visible" v-on="$listeners" @open="onOpen" @close="onClose">
    <el-form
      ref="createForm"
      :model="form"
      :rules="rules"
      size="medium"
      label-width="100px"
      label-position="left"
    >
      <el-form-item label="表单名" prop="Title">
        <el-input v-model="form.Title" placeholder="请输入表单名" clearable :style="{width: '100%'}" />
      </el-form-item>
      <el-form-item label="表单尺寸" prop="Size">
        <el-radio-group v-model="form.Size" size="medium">
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
        <el-radio-group v-model="form.LabelPosition" size="medium">
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
        <el-input-number v-model="form.LabelWidth" placeholder="标签宽度" />
      </el-form-item>
      <el-form-item label="栅格间隔" prop="Span">
        <el-input-number v-model="form.Span" placeholder="栅格间隔" />
      </el-form-item>
      <el-form-item label="按钮" prop="">
        <el-switch v-model="form.FormBtns" />
      </el-form-item>
      <el-form-item label="禁用">
        <el-switch v-model="form.Disabled" />
      </el-form-item>
      <el-form-item label="列">
        <el-button>更新</el-button>
        <el-table
          :data="form.Fields"
          border
          style="width: 100%"
        >
          <el-table-column
            fixed
            prop="date"
            label="日期"
            width="150"
          />
          <el-table-column
            prop="name"
            label="姓名"
            width="120"
          />
          <el-table-column
            prop="province"
            label="省份"
            width="120"
          />
          <el-table-column
            prop="city"
            label="市区"
            width="120"
          />
          <el-table-column
            prop="address"
            label="地址"
            width="300"
          />
          <el-table-column
            prop="zip"
            label="邮编"
            width="120"
          />
          <el-table-column
            fixed="right"
            label="操作"
            width="100"
          >
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="handleClick(scope.row)">查看</el-button>
              <el-button type="text" size="small">编辑</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-form-item>
    </el-form>

    <div slot="footer">
      <el-button @click="handleCancel">取消</el-button>
      <el-button type="primary" @click="handelConfirm">确定</el-button>
    </div>
  </el-dialog>
</template>

<script>
import { formCreate } from '@/api/form'

export default {
  name: 'Dialog',
  data() {
    return {
      visible: false,
      form: {
        Title: 'abc',
        Size: 'medium',
        LabelPosition: 'left',
        LabelWidth: 100,
        Span: 15,
        FormBtns: true,
        Disabled: false,
        Fields: []
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
      }]
    }
  },
  methods: {
    onOpen() {
    },
    onClose() {
      this.$refs['createForm'].resetFields()
    },
    create() {
      console.log(this.$refs['createForm'])
      this.visible = true
    },
    update() {
      this.visible = true
    },
    view() {
      this.visible = true
    },
    close() {
      this.visible = false
    },
    handleCancel() {
      this.close()
    },
    handelConfirm() {
      this.$refs['createForm'].validate(valid => {
        if (!valid) return

        console.log(this.form)
        formCreate(this.form).then(response => {
          this.$emit('fetch')
        })
        this.close()
      })
    }
  }
}
</script>

<style scoped>

</style>
