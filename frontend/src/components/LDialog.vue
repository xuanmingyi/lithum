<template>
  <el-dialog title="收货地址" :visible.sync="dialogTableVisible">
  </el-dialog>
</template>

<script>
export default {
  name: 'LDialog',
  data: function() {
    return {
      dialogTableVisible: true
    }
  }
}
</script>

<style scoped>

</style>
