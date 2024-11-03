<template>
  <div>
    <div class="management-tools">
      <div class="search-section">
        <el-select
          v-model="searchType"
          placeholder="选择身份"
          @change="handleSearch"
        >
          <el-option label="所有" value=""></el-option>
          <el-option label="教职员" value="教职员"></el-option>
          <el-option label="学生" value="学生"></el-option>
        </el-select>
        <el-input
          placeholder="按姓名搜索"
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

    <el-table :data="paginatedMembers" style="width: 100%" class="table">
      <el-table-column prop="photo" label="照片" width="100">
        <template #default="scope">
          <el-image
            style="width: 50px; height: 50px"
            :src="scope.row.photo_url"
            fit="cover"
          ></el-image>
        </template>
      </el-table-column>
      <el-table-column prop="name" label="姓名" width="100"></el-table-column>
      <el-table-column prop="identity" label="身份" width="100"></el-table-column>
      <el-table-column prop="title" label="职称/学历" width="120"></el-table-column>
      <el-table-column prop="phone_number" label="手机号" width="200"></el-table-column>
      <el-table-column prop="email" label="邮箱" width="200"></el-table-column>
      <el-table-column prop="description" label="个人介绍" fit="fill"></el-table-column>
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
      :page-size="pageSize"
      layout="total, prev, pager, next, jumper"
      :total="totalItems"
    ></el-pagination>

    <el-dialog
      v-model="dialogVisible"
      :title="currentMember.id ? '编辑成员信息' : '新增成员信息'"
      width="40%"
    >
      <div class="dialog-content">
        <el-form :model="currentMember" label-width="80px" class="form">
          <el-form-item label="姓名">
            <el-input v-model="currentMember.name"></el-input>
          </el-form-item>
          <el-form-item label="身份">
            <el-select v-model="currentMember.identity" placeholder="请选择">
              <el-option label="教职员" value="教职员"></el-option>
              <el-option label="学生" value="学生"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="职称/学历">
            <el-select v-model="currentMember.title" placeholder="请选择">
              <el-option label="教授" value="教授"></el-option>
              <el-option label="副教授" value="副教授"></el-option>
              <el-option label="讲师" value="讲师"></el-option>
              <el-option label="工程师" value="工程师"></el-option>
              <el-option label="博士研究生" value="博士研究生"></el-option>
              <el-option label="硕士研究生" value="硕士研究生"></el-option>
              <el-option label="本科生" value="本科生"></el-option>
              <el-option label="毕业生" value="毕业生"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="手机号">
            <el-input v-model="currentMember.phone_number"></el-input>
          </el-form-item>
          <el-form-item label="邮箱">
            <el-input v-model="currentMember.email"></el-input>
          </el-form-item>
          <el-form-item label="个人介绍">
            <el-input type="textarea" v-model="currentMember.description"></el-input>
          </el-form-item>
          <el-form-item label="照片">
            <el-upload
              class="avatar-uploader"
              action="/api/upload"
              :show-file-list="false"
              :on-success="handleUploadSuccess"
              :before-upload="beforeAvatarUpload"
              name="image"
            >
              <img
                v-if="currentMember.photo_url"
                :src="currentMember.photo_url"
                class="avatar"
              />
              <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
            </el-upload>
          </el-form-item>
        </el-form>
      </div>
      <div class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="updateMember">保存</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import axios from "@/utils/axios";

export default {
  name: "MemberManage",
  data() {
    return {
      dialogVisible: false,
      currentMember: {},
      currentPage: 1,
      pageSize: 10,
      searchName: "",
      searchType: "",
      totalItems: 0,
      fullMembersList: [], // 存储从API获取的所有成员
      members: [], // 过滤后显示在表格中的成员
    };
  },
  computed: {
    paginatedMembers() {
      const start = (this.currentPage - 1) * this.pageSize;
      const end = start + this.pageSize;
      return this.members.slice(start, end);
    },
  },
  mounted() {
    this.getMembers();
  },
  methods: {
    handleUploadSuccess(response) {
      if (response.code === 200) {
        this.currentMember.photo_url = response.data.image_url;
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
    handleSearch() {
      this.currentPage = 1;
      this.filterMembers();
    },
    getMembers() {
      axios
        .get("/api/members")
        .then((response) => {
          if (response.data.code === 200) {
            this.fullMembersList = response.data.data.lists;
            this.filterMembers();
            this.totalItems = this.fullMembersList.length;
          } else {
            console.error("Failed to get member list", response.data.msg);
          }
        })
        .catch((error) => {
          console.error("Error fetching member list:", error);
        });
    },
    filterMembers() {
      this.members = this.fullMembersList.filter((member) => {
        const matchesIdentity = this.searchType ? member.identity === this.searchType : true;
        const matchesName = this.searchName ? member.name.toLowerCase().includes(this.searchName.toLowerCase()) : true;
        return matchesIdentity && matchesName;
      });
      this.totalItems = this.members.length; // Update total count for pagination
    },
    handleAdd() {
      this.currentMember = {
        name: "",
        identity: "",
        title: "",
        phone_number: "",
        email: "",
        description: "",
        photo_url: "", // Ensure to initialize photo_url
      };
      this.dialogVisible = true;
    },
    handleImport() {
      this.$message({
        type: "info",
        message: "Feature not implemented yet",
      });
    },
    handleEdit(member) {
      this.currentMember = { ...member };
      this.dialogVisible = true;
    },
    updateMember() {
      if (this.currentMember.name === ""|| this.currentMember.identity === ""|| this.currentMember.title === "") {
        this.$message.error("姓名，身份，职称/学历不能为空");
        return; // 提前返回，不执行后续代码
      }
      const url = this.currentMember.id ? `/api/admin/member/${this.currentMember.id}` : "/api/admin/member";
      const method = this.currentMember.id ? "put" : "post";
      axios.request({
        url: url,
        method: method,
        data: this.currentMember,
        headers: { "Content-Type": "application/json" },
      })
      .then((response) => {
        if (response.data.code === 200) {
          this.$message.success("Member updated successfully");
          this.dialogVisible = false;
          this.getMembers(); // Refresh the list after update
        } else {
          this.$message.error("Failed to update member");
        }
      })
      .catch((error) => {
        this.$message.error("Error updating member: " + error.message);
      });
    },
    handleDelete(currentMember) {
      this.$confirm("Are you sure you want to delete this member?", "Warning", {
        confirmButtonText: "Yes",
        cancelButtonText: "No",
        type: "warning"
      })
      .then(() => {
        axios.delete(`/api/admin/member/${currentMember.id}`)
        .then((response) => {
          if (response.data.code === 200) {
            this.getMembers(); // Refresh the list after deletion
            this.$message.success("Member deleted successfully");
          } else {
            this.$message.error("Failed to delete member");
          }
        })
        .catch((error) => {
          this.$message.error("Error deleting member: " + error.message);
        });
      })
      .catch(() => {
        this.$message.info("Delete cancelled");
      });
    },
    handleCurrentChange(newPage) {
      this.currentPage = newPage;
      this.filterMembers();
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

.search-section, .action-section {
  display: flex;
  align-items: center;
  gap: 10px;
}

.search-input {
  width: 400px;
}

.table {
  overflow-y: auto;
  height: 600px;
}

.el-pagination {
  display:flex;
  justify-content: flex-end;
  margin-top: 20px;
}

.dialog-footer {
  text-align: right;
}

.avatar-uploader {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 178px;
  height: 178px;
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  overflow: hidden;
  position: relative;
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
}

.avatar {
  width: 100%;
  height: 100%;
  display: block;
}
</style>
