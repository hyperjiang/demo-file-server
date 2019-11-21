<template>
  <div>
    <el-upload
      :action="uploadApi"
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
      uploadApi: '',
      fileList: [],
    };
  },
  created() {
    const vm = this;
    const listApi = `${process.env.VUE_APP_BACKEND_API}/list`;

    vm.uploadApi = `${process.env.VUE_APP_BACKEND_API}/upload`;

    this.$axios.get(listApi).then((response) => {
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
