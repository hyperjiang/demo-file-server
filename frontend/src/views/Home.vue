<template>
  <div>
    <el-upload
      action="http://localhost:9090/upload"
      list-type="picture"
      :on-preview="handlePreview"
      :on-remove="handleRemove"
      :before-remove="beforeRemove"
      multiple
      :file-list="fileList"
    >
      <el-button size="small" type="primary">Upload</el-button>
    </el-upload>
    <el-dialog :visible.sync="dialogVisible">
      <img width="100%" :src="dialogImageUrl" alt />
    </el-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      dialogImageUrl: '',
      dialogVisible: false,
      fileList: [],
    };
  },
  created() {
    const vm = this;
    this.$axios.get('http://localhost:9090/list').then((response) => {
      if (response.data != null && response.data.length > 0) {
        vm.fileList = response.data;
      }
    });
  },
  methods: {
    handleRemove(file) {
      this.$axios.delete(file.url);
    },
    handlePreview(file) {
      this.dialogImageUrl = file.url;
      this.dialogVisible = true;
    },
    beforeRemove(file) {
      return this.$confirm(`Are you sure to delete ${file.name}？`);
    },
  },
};
</script>
