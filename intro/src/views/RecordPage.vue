<template>
  <div class="record-page">
    <el-row gutter="15">
      <el-col :span="3">
        <el-menu
          :default-active="selectedYear"
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
      <el-col :span="21">
        <masonry :cols="{ default: 3, 1100: 2, 700: 1 }" gutter="20px">
          <div
            v-for="activity in filteredActivities"
            :key="activity.id"
            class="record-item"
          >
            <h2 style="margin: 5px 0">{{ activity.name }}</h2>
            <el-divider style="margin: 12px 0" />
            <div
              v-masonry
              transition-duration="0.3s"
              item-selector=".card"
              class="pubuliu"
            >
              <div
                v-masonry-tile
                v-for="image in activity.images"
                :key="image.id"
                class="card"
                @click="showImage(image.photo_url)"
              >
                <img
                  :src="image.photo_url"
                  alt="Activity Image"
                  class="activity-image"
                />
                <div class="image-name">{{ image.name }}</div>
              </div>
            </div>
          </div>
        </masonry>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "RecordPage",
  data() {
    return {
      selectedYear: "2024",
      years: [],
      activities: [],
    };
  },
  computed: {
    filteredActivities() {
      return this.activities
        .filter((activity) => activity.year === this.selectedYear)
        .sort((a, b) => b.id - a.id); // 添加这一行来根据 id 倒序排列活动
    },
  },
  created() {
    this.fetchActivities();
  },
  methods: {
    async fetchActivities() {
      try {
        const response = await axios.get("/api/records", {
          params: { visibility: true },
        });
        const data = response.data.data.lists;
        const activitiesMap = new Map();
        const yearsSet = new Set();

        data.forEach((item) => {
          yearsSet.add(item.activity.year);
          if (!activitiesMap.has(item.activity_id)) {
            activitiesMap.set(item.activity_id, {
              id: item.activity.id,
              name: item.activity.name,
              year: item.activity.year,
              images: [],
            });
          }
          activitiesMap.get(item.activity_id).images.push({
            id: item.id,
            photo_url: item.photo_url,
            name: item.name,
            visibility: item.visibility,
          });
        });

        this.activities = Array.from(activitiesMap.values());
        this.years = Array.from(yearsSet).sort((a, b) => b - a);
        if (this.years.length > 0) {
          this.selectedYear = this.years[0].toString();
        }
      } catch (error) {
        console.error("Error fetching activities:", error);
      }
    },
    handleYearSelect(year) {
      this.selectedYear = year;
    },
  },
};
</script>

<style scoped>
.record-page {
  padding: 20px;
}

.el-menu-vertical-demo {
  width: 100%;
  border-right: none;
  background-color: #f4f4f4;
  text-align: center;
}

.el-menu-vertical-demo .el-menu-item {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 16px;
  height: 50px;
  width: 100%;
}

.masonry {
  display: flex;
  margin-left: -20px;
  width: auto;
}

.masonry > div {
  padding-left: 20px;
  background-clip: padding-box;
}

.record-item {
  padding: 10px;
  box-sizing: border-box;
  width: 100%;
}

.carousel-item {
  position: relative;
  overflow: hidden;
  margin-bottom: 20px;
}

.activity-image {
  width: 100%;
  height: auto;
  display: block;
  border-radius: 5px;
  transition: transform 0.3s;
}

.activity-image:hover {
  transform: scale(1.05);
}

.card {
  width: 100%;
  max-width: 220px;
  margin: 0.25em;
  border-radius: 5px;
  box-shadow: 0px 2px 10px 1px rgba(0, 0, 0, 0.1);
  position: relative;
}

.image-name {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  color: white;
  text-align: center;
  padding: 5px 0;
  opacity: 0;
  transition: opacity 0.3s;
}

.card:hover .image-name {
  opacity: 1;
}
</style>
