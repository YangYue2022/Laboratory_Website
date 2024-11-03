<template>
  <div>
    <div class="management-tools">
      <div class="search-section">
        <el-input
          v-model="searchQuery"
          placeholder="搜索活动名称"
          class="search-input"
          @input="handleSearch"
        ></el-input>
        <el-button type="primary" @click="handleSearch">搜索</el-button>
      </div>
      <div class="action-section">
        <el-button type="primary" @click="handleAdd">新增活动</el-button>
      </div>
    </div>

    <el-table :data="records" style="width: 100%" class="table">
      <el-table-column prop="photo" label="图片" fit="fill">
        <template #default="scope">
          <el-image
            style="width: auto; height: 100px"
            :src="scope.row.photo_url"
          ></el-image>
        </template>
      </el-table-column>
      <el-table-column
        prop="activity.name"
        label="活动名称"
        width="200"
      ></el-table-column>
      <el-table-column
        prop="activity.year"
        label="年份"
        width="200"
      ></el-table-column>
      <el-table-column prop="name" label="图片名称" width="200">
        <template #default="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>
      <el-table-column prop="visibility" label="可见" width="200">
        <template #default="scope">
          <el-switch
            v-model="scope.row.visibility"
            @change="updateVisibility(scope.row)"
          ></el-switch>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template #default="scope">
          <el-button type="text" @click="handleEdit(scope.row)">编辑</el-button>
          <el-button
            type="text"
            style="color: red"
            @click="handleDelete(scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      @current-change="handleCurrentChange"
      :current-page="currentPage"
      :page-size="10"
      layout="total, prev, pager, next, jumper"
      :total="totalItems"
    >
    </el-pagination>

    <el-dialog v-model="dialogVisible" title="新增活动" width="50%">
      <el-form ref="form" :model="currentActivity" label-width="80px">
        <el-form-item label="活动名称">
          <el-input v-model="currentActivity.name"></el-input>
        </el-form-item>
        <el-form-item label="年份" prop="year">
          <el-input
            v-model="currentActivity.year"
            placeholder="请输入年份"
          ></el-input>
        </el-form-item>
        <el-form-item label="上传图片">
          <el-upload
            class="upload-demo"
            action="/api/upload"
            multiple
            :on-success="handleUploadSuccess"
            :before-upload="beforeAvatarUpload"
            list-type="picture-card"
            :file-list="currentActivity.images"
            name="image"
          >
            <i class="el-icon-plus"></i>
          </el-upload>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="addActivity">保存</el-button>
        </div>
      </template>
    </el-dialog>
    <el-dialog v-model="editDialogVisible" title="编辑图片名称" width="30%">
      <el-form :model="editingRecord">
        <el-form-item label="图片名称" label-width="80px">
          <el-input v-model="editingRecord.name"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="editDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="updateRecord">保存</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import axios from "@/utils/axios";

export default {
  name: "RecordManage",
  data() {
    return {
      records: [],
      searchQuery: "",
      currentPage: 1,
      dialogVisible: false,
      currentActivity: {
        name: "",
        year: "",
        images: [],
      },
      totalItems: 0,
      editingRecord: {
        id: null,
        name: "",
      },
      editDialogVisible: false,
    };
  },
  mounted() {
    this.getRecords();
  },
  methods: {
    handleUploadSuccess(response, file) {
      if (response.code === 200) {
        const fileNameWithoutExtension = file.name.replace(/\.[^/.]+$/, "");
        this.currentActivity.images.push({
          name: fileNameWithoutExtension,
          url: response.data.image_url,
        });
        this.$message.success("图片上传成功！");
      } else {
        this.$message.error("图片上传失败！");
      }
    },
    beforeAvatarUpload(file) {
      const isJPEG = file.type === "image/jpeg";
      const isJPG = file.type === "image/jpg";
      const isPNG = file.type === "image/png";

      const isLt5M = file.size / 1024 / 1024 < 5;

      if (!isJPEG && !isPNG && !isJPG) {
        this.$message.error("上传头像图片只能是 JPG 或 PNG 或 JPEG 格式!");
      }
      if (!isLt5M) {
        this.$message.error("上传头像图片大小不能超过 5MB!");
      }
      return (isJPG || isPNG || isJPEG) && isLt5M;
    },
    getRecords() {
      axios
        .get("/api/records", {
          params: {
            page: this.currentPage,
            activity_name: this.searchQuery,
          },
        })
        .then((response) => {
          if (response.data.code === 200) {
            this.records = response.data.data.lists;
            this.totalItems = response.data.data.total;
          } else {
            console.error("failed to get record list", response.data.msg);
          }
        });
    },
    handleSearch() {
      this.currentPage = 1; // 搜索时重置到第一页
      this.getRecords();
    },
    handleAdd() {
      this.currentActivity = { name: "", year: "", images: [] };
      this.dialogVisible = true;
      this.$message.info("先传高度大的图片");
    },
    addActivity() {
      // 检查年份是否为四位数字
      if (!/^\d{4}$/.test(this.currentActivity.year)) {
        this.$message.error("年份必须是四位数字");
        return; // 提前返回，不执行后续代码
      }
      if (this.currentActivity.images.length === 0) {
        this.$message.error("图片不能为空");
        return; // 提前返回，不执行后续代码
      }
      const newActivity = {
        name: this.currentActivity.name,
        year: this.currentActivity.year,
        images: this.currentActivity.images,
      };
      // 发送 POST 请求到后端以保存新活动
      axios
        .post("/api/admin/record", newActivity)
        .then((response) => {
          if (response.data.code === 200) {
            this.$message.success("活动添加成功！");
            this.dialogVisible = false;
            this.getRecords(); // 刷新活动列表
          } else {
            this.$message.error("活动添加失败！");
          }
        })
        .catch((error) => {
          console.error("Error adding activity:", error);
          this.$message.error("活动添加失败！");
        });
    },
    handleDelete(activity) {
      this.$confirm("确认删除这条记录吗?", "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          axios.delete(`/api/admin/record/${activity.id}`).then((response) => {
            if (response.data.code === 200) {
              this.getRecords();
              this.$message({
                type: "success",
                message: "删除成功!",
              });
            } else {
              this.$message({
                type: "error",
                message: "删除失败!",
              });
            }
          });
        })
        .catch(() => {
          this.$message({ type: "info", message: "删除取消" });
        });
    },
    handleCurrentChange(newPage) {
      this.currentPage = newPage;
      this.getRecords();
    },
    async updateVisibility(activity) {
      try {
        await axios.put("/api/admin/record/visibility", {
          id: activity.id,
          visibility: activity.visibility,
        });
        console.log("Visibility updated successfully");
      } catch (error) {
        console.error("Error updating visibility", error);
      }
    },
    handleEdit(record) {
      this.editingRecord.id = record.id;
      this.editingRecord.name = record.name;
      this.editDialogVisible = true;
    },
    updateRecord() {
      axios
        .put(`/api/admin/record/${this.editingRecord.id}`, {
          name: this.editingRecord.name,
        })
        .then((response) => {
          if (response.data.code === 200) {
            this.$message.success("更新成功!");
            this.editDialogVisible = false;
            this.getRecords(); // 刷新记录列表
          } else {
            this.$message.error("更新失败!");
          }
        })
        .catch((error) => {
          console.error("更新失败:", error);
          this.$message.error("更新失败!");
        });
    },
  },
};
</script>

<style scoped>
.management-tools {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.search-section,
.action-section {
  display: flex;
  align-items: center;
}

.search-input {
  width: 300px;
  margin-right: 10px;
}

.dialog-footer {
  text-align: right;
}

.table {
  height: 600px;
  overflow-y: auto;
}

.el-pagination {
  display: flex;
  justify-content: flex-end;
  margin-top: 20px;
}
</style>
