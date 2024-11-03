<template>
  <div class="work-page" lang="en">
    <el-row gutter="15">
      <el-col :span="4">
        <el-menu
          :default-active="selectedYear.toString()"
          class="el-menu-vertical-demo"
          @select="handleYearSelect"
        >
          <el-menu-item
            v-for="year in years"
            :key="year"
            :index="year.toString()"
          >
            {{ year }}
          </el-menu-item>
        </el-menu>
      </el-col>
      <el-col :span="20">
        <div
          v-for="type in activeCategories" 
          :key="type"
        >
          <h3>{{ type }}</h3>
          <div
            v-for="(event, index) in filteredEventsByCategory[type]"
            :key="index"
          >
            <p>{{ event.desc }}</p>
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "WorkPage",
  data() {
    return {
      selectedYear: "",
      years: [],
      categories: [
        "期刊论文",
        "会议论文",
        "科研项目",
        "专著",
        "软著/专利",
        "其他",
      ],
      events: [],
      filteredEventsByCategory: {},
    };
  },
  computed: {
    // 计算属性来获取有数据的类别
    activeCategories() {
      return this.categories.filter(type => 
        this.filteredEventsByCategory[type] && 
        this.filteredEventsByCategory[type].length > 0
      );
    }
  },
  mounted() {
    this.getEventsData(); // 在组件加载时获取数据
  },
  methods: {
    handleYearSelect(year) {
      this.selectedYear = year;
      this.filterEvents();
    },
    getEventsData() {
      axios
        .get("/api/works")
        .then((response) => {
          if (response.data.code === 200) {
            this.events = response.data.data.lists;
            this.updateYears(this.events);
            this.filterEvents(); // 确保在数据加载后立即过滤事件
          } else {
            console.error("Failed to fetch data:", response.data.msg);
          }
        })
        .catch((error) => {
          console.error("Error fetching data:", error);
        });
    },
    updateYears(events) {
      const uniqueYears = new Set(events.map((event) => event.year));
      this.years = Array.from(uniqueYears).sort((a, b) => b - a);
      if (this.years.length > 0) {
        this.selectedYear = this.years[0].toString(); // 设置默认年份为数组的第一项或当前年份
        this.filterEvents(); // 一旦设置了 selectedYear，立即过滤事件
      } else {
        this.selectedYear = new Date().getFullYear().toString();
      }
    },
    filterEvents() {
      this.filteredEventsByCategory = this.categories.reduce(
        (acc, category) => {
          acc[category] = this.events.filter(
            (event) =>
              event.type === category &&
              event.year.toString() === this.selectedYear
          );
          return acc;
        },
        {}
      );
    },
  },
};
</script>

<style scoped>
.el-menu-vertical-demo {
  width: 100%;
  border-right: none;
  background-color: #f4f4f4;
  text-align: center;
  margin-top: 20px;
}

.el-menu-vertical-demo .el-menu-item {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 16px;
  height: 50px; /* Adjust as needed */
  width: 100%;
}

p {
  background-color: #f4f4f4;
  padding: 5px;
  margin: 0 0 10px 0;
}
</style>
