<template>
  <div>
    <div class="management-tools">
      <div class="search-section">
        <el-select
          v-model="searchType"
          placeholder="选择类型"
          @change="handleSearch"
        >
          <el-option label="所有" value=""></el-option>
          <el-option label="期刊论文" value="期刊论文"></el-option>
          <el-option label="会议论文" value="会议论文"></el-option>
          <el-option label="科研项目" value="科研项目"></el-option>
          <el-option label="专著" value="专著"></el-option>
          <el-option label="软著/专利" value="软著/专利"></el-option>
          <el-option label="其他" value="其他"></el-option>
        </el-select>
        <el-input
          placeholder="按名称搜索"
          v-model="searchName"
          class="search-input"
          @input="handleSearch"
        ></el-input>
        <el-button type="primary" @click="handleSearch">搜索</el-button>
      </div>
      <div class="action-section">
        <el-button type="primary" @click="handleAdd">新增</el-button>
        <el-button type="success" @click="handleImport">批量导入</el-button>
      </div>
    </div>

    <el-table :data="works" style="width: 100%" class="table">
      <el-table-column prop="type" label="类型" width="150"></el-table-column>
      <el-table-column prop="year" label="年份" width="100"></el-table-column>
      <el-table-column prop="desc" label="描述" fit="fill"></el-table-column>
      <el-table-column fixed="right" label="操作" width="150">
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

    <el-dialog
      v-model="dialogVisible"
      :title="currentWork.id ? '编辑成果信息' : '新增成果信息'"
      width="50%"
    >
      <el-form :model="currentWork" >
        <el-form-item label="类型" prop="type">
          <el-select v-model="currentWork.type" placeholder="请选择">
            <el-option label="期刊论文" value="期刊论文"></el-option>
            <el-option label="会议论文" value="会议论文"></el-option>
            <el-option label="科研项目" value="科研项目"></el-option>
            <el-option label="专著" value="专著"></el-option>
            <el-option label="软著/专利" value="软著/专利"></el-option>
            <el-option label="其他" value="其他"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="年份" prop="year">
          <el-input
            v-model="currentWork.year"
            placeholder="请输入年份"
          ></el-input>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="currentWork.desc" type="textarea" rows="5" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="updateWork">保存</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import axios from '@/utils/axios';

export default {
  name: "WorkManage",
  data() {
    return {
      works: [],
      searchName: "",
      searchType: "",
      currentPage: 1,
      dialogVisible: false,
      currentWork: {},
      totalItems: 0,
    };
  },
  mounted() {
    this.getWorks();
  },
  methods: {
    getWorks() {
      axios
        .get("/api/works", {
          params: {
            desc: this.searchName,
            type: this.searchType,
            page: this.currentPage,
          },
        })
        .then((response) => {
          if (response.data.code === 200) {
            this.works = response.data.data.lists;
            this.totalItems = response.data.data.total;
          } else {
            console.error("Failed to fetch work list:", response.data.msg);
          }
        })
        .catch((error) => {
          console.error("Error fetching work list:", error);
        });
    },

    handleAdd() {
      this.currentWork = { type: "", year: "", desc: "" };
      this.dialogVisible = true;
    },
    updateWork() {
      if (!Number.isInteger(this.currentWork.year)) {
        this.currentWork.year = parseInt(this.currentWork.year, 10) || 0; // 再次确保是整数
      }
      const url = this.currentWork.id
        ? `/api/admin/work/${this.currentWork.id}`
        : "/api/admin/work";
      const method = this.currentWork.id ? "put" : "post";

      axios
        .request({
          url,
          method,
          data: this.currentWork,
          headers: {
            "Content-Type": "application/json",
          },
        })
        .then((response) => {
          if (response.data.code === 200) {
            this.dialogVisible = false;
            this.getWorks();
            this.$message({
              type: "success",
              message: "成果已保存!",
            });
          } else {
            console.error("Failed to update work:", response.data.msg);
            this.$message({
              type: "error",
              message: "保存失败!",
            });
          }
        });
    },
    handleSearch() {
      // 设置当前页为第一页
      this.currentPage = 1;
      this.getWorks();
    },
    handleEdit(member) {
      this.currentWork = { ...member };
      this.dialogVisible = true;
    },
    handleDelete(currentWork) {
      this.$confirm("确认删除这条记录吗?", "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          axios
            .delete(`/api/admin/work/${currentWork.id}`)
            .then((response) => {
              if (response.data.code === 200) {
                this.getWorks();
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
            })
            .catch((error) => {
                console.error("Error deleting work:", error);
            });
        })
        .catch(() => {
          this.$message({ type: "info", message: "删除取消" });
        });
    },
    handleCurrentChange(newPage) {
      this.currentPage = newPage;
      this.getWorks();
    },
    handleImport() {
      this.$message({
        type: "info",
        message: "待实现",
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

.search-section {
  display: flex;
  align-items: center;
  gap: 10px;
}

.action-section {
  display: flex;
  align-items: center;
  gap: 10px;
}

.search-input {
  width: 400px;
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

.dialog-footer {
  text-align: right;
}
</style>
