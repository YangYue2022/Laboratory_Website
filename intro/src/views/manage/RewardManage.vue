<template>
  <div>
    <div class="management-tools">
      <div class="search-section">
        <el-select v-model="searchType" placeholder="选择类型" @change="handleSearch">
          <el-option label="所有" value=""></el-option>
          <el-option label="竞赛" value="竞赛"></el-option>
          <el-option label="其他" value="其他"></el-option>
        </el-select>
        <el-input placeholder="按竞赛名称搜索" v-model="searchName" class="search-input" @input="handleSearch"></el-input>
        <el-button type="primary" @click="handleSearch">搜索</el-button>
      </div>
      <div class="action-section">
        <el-button type="primary" @click="handleAdd">新增</el-button>
        <el-button type="success" @click="handleImport">批量导入</el-button>
      </div>
    </div>

    <el-table :data="rewards" style="width: 100%" class="table">
      <el-table-column prop="year" label="年度" width="100"></el-table-column>
      <el-table-column prop="type" label="类型" width="150"></el-table-column>
      <el-table-column prop="name" label="竞赛名称" width="250"></el-table-column>
      <el-table-column prop="project" label="项目名称" fit="fill"></el-table-column>
      <el-table-column prop="winner" label="获奖人" width="350"></el-table-column>
      <el-table-column fixed="right" label="操作" width="150">
        <template #default="scope">
          <el-button type="text" @click="handleEdit(scope.row)">编辑</el-button>
          <el-button type="text" style="color: red" @click="handleDelete(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination @current-change="handleCurrentChange" :current-page="currentPage" :page-size="10"
      layout="total, prev, pager, next, jumper" :total="totalItems">
    </el-pagination>

    <!-- 编辑奖励信息对话框 -->
    <el-dialog v-model="dialogVisible" :title="currentReward.id ? '编辑奖励信息' : '新增奖励信息'" width="30%">
      <el-form :model="currentReward" :rules="rules" ref="rewardForm" label-width="80px">
        <el-form-item label="年度">
          <el-input v-model="currentReward.year" @input="handleYearChange"></el-input>
        </el-form-item>
        <el-form-item label="类型">
          <el-select v-model="currentReward.type" placeholder="请选择">
            <el-option label="竞赛" value="竞赛"></el-option>
            <el-option label="其他" value="其他"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="竞赛名称">
          <el-input v-model="currentReward.name"></el-input>
        </el-form-item>
        <el-form-item label="项目名">
          <el-input v-model="currentReward.project"></el-input>
        </el-form-item>
        <el-form-item label="获奖人">
          <el-input v-model="currentReward.winner"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="updateReward">保存</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import axios from '@/utils/axios';

export default {
  name: "RewardManage",
  data() {
    return {
      rewards: [],
      searchName: "",
      searchType: "",
      currentPage: 1,
      dialogVisible: false,
      currentReward: {},
      totalItems: 0,
      rules: {
        year: [
          { required: true, message: "请输入年度", trigger: "blur" },
          { min: 4, max: 4, message: "年度必须为4位数字", trigger: "blur" },
          { pattern: /^\d{4}$/, message: "请输入合法的年份", trigger: "blur" },
        ],
        type: [{ required: true, message: "请选择类型", trigger: "change" }],
        name: [{ required: true, message: "请输入竞赛名称", trigger: "blur" }],
        project: [{ required: true, message: "请输入项目名", trigger: "blur" }],
        winner: [
          { required: true, message: "请输入获奖人名称", trigger: "blur" },
        ],
      },
    };
  },
  mounted() {
    this.getRewards();
  },
  methods: {
    handleYearChange(value) {
      this.currentReward.year = parseInt(value, 10) || 0; // 转换为整数，非数字时默认为0
    },
    handleCurrentChange(newPage) {
      this.currentPage = newPage;
      this.getRewards(); // 重新获取当前页的数据
    },
    handleSearch() {
      // 设置当前页为第一页
      this.currentPage = 1;
      this.getRewards();
    },
    handleImport() {
      this.$message({
        type: "info",
        message: "待实现",
      });
    },
    getRewards() {
      axios
        .get("/api/rewards", {
          params: {
            name: this.searchName,
            type: this.searchType,
            page: this.currentPage,
          },
        })
        .then((response) => {
          if (response.data.code === 200) {
            this.rewards = response.data.data.lists;
            this.totalItems = response.data.data.total;
          } else {
            console.error("Failed to fetch reward list:", response.data.msg);
          }
        })
        .catch((error) => {
          console.error("Error fetching reward list:", error);
        });
    },
    handleAdd() {
      this.currentReward = {
        year: "",
        type: "",
        name: "",
        project: "",
        winner: "",
      };
      this.dialogVisible = true;
    },
    updateReward() {
      if (!Number.isInteger(this.currentReward.year)) {
        this.currentReward.year = parseInt(this.currentReward.year, 10) || 0; // 再次确保是整数
      }
      const url = this.currentReward.id
        ? `/api/admin/reward/${this.currentReward.id}`
        : "/api/admin/reward";
      const method = this.currentReward.id ? "put" : "post";

      axios({
        method: method,
        url: url,
        headers: {
          "Content-Type": "application/json",
        },
        data: this.currentReward,
      })
        .then((response) => {
          if (response.data.code === 200) {
            this.getRewards();
            this.dialogVisible = false;
            this.$message({
              type: "success",
              message: "奖励已保存!",
            });
          } else {
            console.error("Save failed:", response.data.msg);
            this.$message({
              type: "error",
              message: "保存失败!",
            });
          }
        })
        .catch((error) => {
          console.error("Error saving reward:", error);
        });
    },
    handleEdit(reward) {
      this.currentReward = { ...reward };
      this.dialogVisible = true;
    },
    handleDelete(reward) {
      this.$confirm("确认删除这条记录吗?", "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(() => {
        axios
          .delete(`/api/admin/reward/${reward.id}`)
          .then((response) => {
            if (response.data.code === 200) {
              this.getRewards();
              this.$message({
                type: "success",
                message: "删除成功!",
              });
            } else {
              console.error("Delete failed:", response.data.msg);
            }
          })
          .catch((error) => {
            console.error("Error deleting reward:", error);
          });
      }).catch(() => {
        this.$message({
          type: "info",
          message: "已取消删除",
        });
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
  gap: 10px;
}

.search-input {
  width: 500px;
}

.search-section {
  width: 500px;
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
