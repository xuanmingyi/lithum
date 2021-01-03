<template>
  <!--https://stackoverflow.com/questions/50687897/computed-function-being-called-twice-in-vuejs-app/50688053-->
  <el-dialog id="dialog-1" :visible="true" :title="action.dialog.title">
    <el-form
      :model="form"
      size="medium"
      label-position="left"
      label-width="80px"
    >
      <el-form-item v-for="field in action.dialog.fields" :key="field.name" :label="field.display">
        <el-input v-if="field.tag === 'el-input'" v-model="form[field.name]" :required="field.required" :disabled="field.disabled" />

        <el-input v-if="field.tag === 'el-area'" v-model="form[field.name]" type="textarea" :required="field.required" :disabled="field.disabled" />

        <el-input-number v-else-if="field.tag === 'el-input-number'" />

        <el-select v-else-if="field.tag === 'el-select'" v-model="form[field.name]">
          <el-option
            v-for="option in field.options"
            :key="option.value"
            :label="option.label"
            :value="option.value"
          />
        </el-select>

        <el-checkbox-group v-else-if="field.tag === 'el-checkbox-group'" v-model="form[field.name]">
          <el-checkbox
            v-for="option in field.options"
            :key="option.value"
            :label="option.label"
          />
        </el-checkbox-group>

        <el-radio-group v-else-if="field.tag === 'el-radio-group'" v-model="form[field.name]">
          <el-radio
            v-for="option in field.options"
            :key="option.value"
            :label="option.label"
            :value="option.value"
          />
        </el-radio-group>

        <el-slider v-else-if="field.tag === 'el-slider'" v-model="form[field.name]" />

        <el-switch v-else-if="field.tag === 'el-switch'" v-model="form[field.name]" />

        <el-upload v-else-if="field.tag === 'el-upload'" action="https://jsonplaceholder.typicode.com/posts/" />
      </el-form-item>
    </el-form>
    <div slot="footer">
      <el-button @click="close">取消</el-button>
      <el-button type="primary" @click="handleConfirm">确定</el-button>
    </div>
  </el-dialog>
</template>

<script>export default {
  name: 'LDialog',
  data: function() {
    return {
      visible: true,
      form: {}
    }
  },
  created() {
    for (let i = 0; i < this.action.dialog.fields.length; i++) {
      const field = this.action.dialog.fields[i]
      this.form[field.name] = field.default
    }
    console.log(this.form)
  },
  methods: {
    close() {
      document.querySelector('#dialog-1').remove()
      document.querySelector('.v-modal').remove()
    },
    handleConfirm() {
      console.log(this.form)
      this.$bus.emit('create', this.form)
    }
  }
}
</script>

<style scoped>

</style>
