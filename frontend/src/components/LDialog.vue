<template>
  <!--https://stackoverflow.com/questions/50687897/computed-function-being-called-twice-in-vuejs-app/50688053-->
  <el-dialog id="dialog-1" :visible="true" :title="action.dialog.title">
    <el-form
      size="medium"
      label-position="left"
      label-width="80px"
    >
      <el-form-item v-for="field in action.dialog.fields" :key="field.name" :label="field.display">
        <el-input v-if="field.tag === 'el-input'" v-model="field.value" :required="field.required" :disabled="field.disabled" />

        <el-input v-if="field.tag === 'el-area'" v-model="field.value" type="textarea" :required="field.required" :disabled="field.disabled" />

        <el-input-number v-else-if="field.tag === 'el-input-number'" />

        <el-select v-else-if="field.tag === 'el-select'" v-model="field.value">
          <el-option
            v-for="option in field.options"
            :key="option.value"
            :label="option.label"
            :value="option.value"
          />
        </el-select>

        <el-checkbox-group v-else-if="field.tag === 'el-checkbox-group'" v-model="checkboxvalue">
          <el-checkbox
            v-for="option in field.options"
            :key="option.value"
            :label="option.label"
          />
        </el-checkbox-group>

        <el-radio-group v-else-if="field.tag === 'el-radio-group'" v-model="field.value">
          <el-radio
            v-for="option in field.options"
            :key="option.value"
            :label="option.label"
            :value="option.value"
          />
        </el-radio-group>

        <el-slider v-else-if="field.tag === 'el-slider'" v-model="field.value" />

        <el-switch v-else-if="field.tag === 'el-switch'" v-model="field.value" />

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
      list: [],
      checkboxvalue: []
    }
  },
  computed: {
    formData: function() {
      const mm = {}
      console.log(this.action)
      for (let i = 0; i < this.action.dialog.fields.length; i++) {
        const field = this.action.dialog.fields[i]
        console.log(field)
        mm[field.name] = 1111
      }
      return mm
    }
  },
  methods: {
    close() {
      document.querySelector('#dialog-1').remove()
      document.querySelector('.v-modal').remove()
    },
    handleConfirm() {
      this.$emit('createItem')
    }
  }
}
</script>

<style scoped>

</style>
