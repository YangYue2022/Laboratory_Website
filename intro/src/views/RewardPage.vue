<template>
  <div class="reward-page">
    <div class="reward-item">
      <h3 style="margin: 5px 0">竞赛</h3>
      <hr style="margin: 12px 0" />
      <table class="custom-table">
        <thead>
          <tr class="custom-header">
            <th>年度</th>
            <th>类型</th>
            <th>项目名</th>
            <th>获奖人</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(item, index) in competitionData"
            :key="index"
            :class="{ 'even-row': index % 2 === 0, 'odd-row': index % 2 !== 0 }"
          >
            <td>{{ item.year }}</td>
            <td>{{ item.name }}</td>
            <td>{{ item.project }}</td>
            <td>{{ item.winner }}</td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="reward-item">
      <h3 style="margin: 5px 0">其他</h3>
      <hr style="margin: 12px 0" />
      <table class="custom-table">
        <thead>
          <tr class="custom-header">
            <th>年度</th>
            <th>类型</th>
            <th>项目名</th>
            <th>获奖人</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(item, index) in otherData"
            :key="index"
            :class="{ 'even-row': index % 2 === 0, 'odd-row': index % 2 !== 0 }"
          >
            <td>{{ item.year }}</td>
            <td>{{ item.name }}</td>
            <td>{{ item.project }}</td>
            <td>{{ item.winner }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "RewardPage",
  data() {
    return {
      tableData: [], // 原始数据
      competitionData: [], // 竞赛数据
      otherData: [], // 其他数据
    };
  },
  mounted() {
    this.getTableData();
  },
  methods: {
    getTableData() {
      axios
        .get("/api/rewards")
        .then((response) => {
          if (response.data.code === 200) {
            this.tableData = response.data.data.lists;
            this.classifyData(); // 调用分类方法
          } else {
            console.error("Failed to fetch QA list:", response.data.msg);
          }
        })
        .catch((error) => {
          console.error("Error fetching reward data:", error);
        });
    },
    //根据 type 分类数据
    classifyData() {
        this.competitionData = this.tableData.filter(item => item.type === '竞赛').sort((a, b) => b.id - a.id).sort((a, b) => b.year - a.year);
        this.otherData = this.tableData.filter(item => item.type === '其他').sort((a, b) => b.id - a.id).sort((a, b) => b.year - a.year);
    }
  },
};
</script>

<style>
.custom-table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 20px;
}

.custom-table th,
.custom-table td {
  border: 2px solid #f4f4f4;
  text-align: center;
  padding: 8px;
}

.custom-table .custom-header {
  background-color: #fad7a0; /* Change this color to your desired header color */
}

.custom-table .even-row {
  background-color: #ececec;
}

.custom-table .odd-row {
  background-color: #f4f4f4;
}

.reward-page {
  margin-top: 20px;
}
</style>
