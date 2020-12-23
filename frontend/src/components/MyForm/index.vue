<template>
  <!--https://stackoverflow.com/questions/50687897/computed-function-being-called-twice-in-vuejs-app/50688053-->
<el-dialog :visible="visible" :title="form.title" v-if="form.fields">
  <el-form
    :model="formData"
    :label-position="form.labelPosition"
    :size="form.size"
    :label-width="labelWidth">
    <el-form-item v-for="field in form.fields" :key="field.id" :label="field.label">
      <el-input v-if="field.tag === 'el-input'" v-bind:required="field.required" v-bind:disabled="field.disabled"></el-input>

      <el-input-number v-else-if="field.tag === 'el-input-number'"></el-input-number>

      <el-select v-else-if="field.tag === 'el-select'">
        <el-option
          v-for="option in field.options"
          :key="option.value"
          :label="option.label"
          :value="option.value"></el-option>
      </el-select>

      <el-checkbox-group v-else-if="field.tag === 'el-checkbox-group'">
        <el-checkbox
          v-for="option in field.options"
          :key="option.value"
          :label="option.label"
          :value="option.value"></el-checkbox>
      </el-checkbox-group>

      <el-radio-group v-model="list" v-else-if="field.tag === 'el-radio-group'">
        <el-radio
          v-for="option in field.options"
          :key="option.value"
          :label="option.label"
          :value="option.value"></el-radio>
      </el-radio-group>

      <el-slider v-else-if="field.tag === 'el-slider'"></el-slider>

      <el-switch v-else-if="field.tag === 'el-switch'"></el-switch>

      <el-upload v-else-if="field.tag === 'el-upload'" action="https://jsonplaceholder.typicode.com/posts/"></el-upload>
    </el-form-item>
  </el-form>
  <div slot="footer">
    <el-button @click="close">取消</el-button>
    <el-button type="primary" @click="handleConfirm">确定</el-button>
  </div>
</el-dialog>
</template>

<script>
export default {
  name: 'MyForm',
  props: [
    'form'
  ],
  data: function() {
    return {
      visible: true,
      list: [],
      ss: 'ss'
    }
  },
  computed: {
    labelWidth: function() {
      return this.form.labelWidth + 'px'
    },
    formData: function() {
      let mm = {}
      for (let i = 0; i < this.form.fields.length; i++) {
        const field = this.form.fields[i]
        console.log(field)
        mm[field.vModel] = field.defaultValue
      }
      return mm
    }
  },
  methods: {
    close() {
      this.$emit('hidden')
    },
    handleConfirm() {
      this.close()
    }
  }
}
</script>

<style scoped>
</style>
