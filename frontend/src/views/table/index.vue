<template>
  <div class="app-container">
    <el-button v-on:click="openDialog()">创建</el-button>
    <MyForm @hidden="hideDialog" :form="form"></MyForm>
  </div>
</template>

<script>
import { default as MyForm } from '../../components/MyForm'
import { form_user } from '@/api/form'

export default {
  components: {
    'MyForm': MyForm
  },
  data() {
    return {
      form: {},
      list: null,
      listLoading: true
    }
  },
  created() {
    form_user().then(response => {
      this.form = JSON.parse(response.data)
    })
  },
  methods: {
    hideDialog() {
      this.show = false
    },
    openDialog() {
      this.show = true
    }
  }
}
</script>
