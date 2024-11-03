<template>
  <div>
    <div class="management-tools">
      <div class="search-section">
        <el-input
          placeholder="搜索问题..."
          v-model="searchQuery"
          class="search-input"
          @input="handleSearch"
        ></el-input>
        <el-button type="primary" @click="handleSearch">搜索</el-button>
      </div>
      <div class="action-section">
        <el-button type="success" @click="handleEdit">新增问题</el-button>
      </div>
    </div>

    <el-table :data="issues" style="width: 100%" class="table">
      <el-table-column prop="que" label="问题"></el-table-column>
      <el-table-column prop="ans" label="回答"></el-table-column>
      <el-table-column label="操作" width="180">
        <template #default="scope">
          <el-button size="text" @click="handleEdit(scope.row)">编辑</el-button>
          <el-button
            size="text"
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

    <!-- 编辑和新增问题的对话框 -->
    <el-dialog
      :title="currentIssue.id ? '编辑问题' : '新增问题'"
      v-model="dialogVisible"
    >
      <el-form :model="currentIssue">
        <el-form-item label="问题">
          <el-input v-model="currentIssue.que" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="回答">
          <el-input type="textarea" v-model="currentIssue.ans" placeholder="支持md语法"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveIssue">保存</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import axios from '@/utils/axios';

export default {
  name: "IssueManage",
  data() {
    return {
      searchQuery: "",
      issues: [],
      currentPage: 1,
      totalItems: 0,
      dialogVisible: false,
      currentIssue: {},
    };
  },
  methods: {
    handleCurrentChange(newPage) {
      this.currentPage = newPage;
      this.fetchQAs(); // 重新获取当前页的数据
    },
    handleSearch() {
      // 设置当前页为第一页
      this.currentPage = 1;

      // 调用 fetchQAs 方法来重新获取数据
      this.fetchQAs();
    },
    fetchQAs() {
      axios
        .get("/api/qas", {
          params: {
            que: this.searchQuery,
            page: this.currentPage,
          },
        })
        .then((response) => {
          if (response.data.code === 200) {
            this.issues = response.data.data.lists;
            this.totalItems = response.data.data.total;
          } else {
            console.error("Failed to fetch QA list:", response.data.msg);
          }
        })
        .catch((error) => {
          console.error("Error fetching QA list:", error);
        });
    },
    handleAdd() {
      this.currentIssue = { question: "", answer: "" };
      this.dialogVisible = true;
    },
    handleEdit(issue) {
      this.currentIssue = Object.assign({}, issue);
      this.dialogVisible = true;
    },
    handleDelete(issue) {
      this.$confirm("确认删除这个问题?", "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          axios
            .delete(`/api/admin/qa/${issue.id}`)
            .then((response) => {
              if (response.data.code === 200) {
                this.issues = this.issues.filter((i) => i.id !== issue.id);
                this.$message({
                  type: "success",
                  message: "删除成功!",
                });
              } else {
                console.error("Delete failed:", response.data.msg);
              }
            })
            .catch((error) => {
              console.error("Error deleting QA:", error);
            });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
    saveIssue() {
      const url = this.currentIssue.id
        ? `/api/admin/qa/${this.currentIssue.id}`
        : "/api/admin/qa";
      const method = this.currentIssue.id ? "put" : "post";

      axios({
        method: method,
        url: url,
        headers: {
          "Content-Type": "application/json",
        },
        data: this.currentIssue,
      })
        .then((response) => {
          if (response.data.code === 200) {
            const returnedIssue = response.data.data.item;
            if (this.currentIssue.id) {
              // 更新操作
              const index = this.issues.findIndex(
                (i) => i.id === this.currentIssue.id
              );
              if (index !== -1) {
                this.issues[index] = Object.assign({}, returnedIssue); // 使用后端返回的数据更新
              }
            } else {
              // 添加操作
              this.issues.push(returnedIssue); // 直接添加后端返回的新对象
              this.totalItems += 1; // 增加 totalItems 计数
              // 检查是否需要换页
              if (this.issues.length > this.pageSize) {
                this.currentPage += 1;
                this.issues = []; // 可以选择清空当前页的数据或重新加载当前页数据
                this.fetchQAs();
              }
            }
            this.dialogVisible = false;
            this.$message({
              type: "success",
              message: "问题已保存!",
            });
          } else {
            console.error("Failed to update work:", response.data.msg);
            this.$message({
              type: "error",
              message: "保存失败!",
            });
          }
        })
        .catch((error) => {
          console.error("Error saving QA:", error);
        });
    },
  },
  mounted() {
    this.fetchQAs();
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
  gap: 10px;
}

.search-input {
  width: 300px;
}

.table {
  overflow-y: auto;
  height: 600px;
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
