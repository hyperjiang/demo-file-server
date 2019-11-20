<template>
  <div>
    <el-upload
      action="http://localhost:9090/upload"
      list-type="picture-card"
      :on-preview="handlePreview"
      :on-remove="handleRemove"
      :before-remove="beforeRemove"
      multiple
      :file-list="fileList"
    >
      <i class="el-icon-plus"></i>
    </el-upload>
    <el-dialog :visible.sync="dialogVisible" width="30%">
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
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePreview(file) {
      this.dialogImageUrl = file.url;
      this.dialogVisible = true;
    },
    beforeRemove(file) {
      return this.$confirm(`确定移除 ${file.name}？`);
    },
  },
};
</script>
