<template>
  <div>
    <p class="title">地址</p>
    <el-divider style="margin: 5px 0" />
    <div id="container"></div>
    <div class="address-detail">
      <p>xxx实验室</p>
      <p>软件学院</p>
      <p>xxxxxxxxxxx</p>
    </div>
    <p class="title">Q & A</p>
    <el-divider style="margin: 5px 0" />
    <div v-for="(item, index) in qaList" :key="index">
      <div class="qa-item">
        <p class="question" v-html="renderMarkdown(item.que)"></p>
        <p class="answer" v-html="renderMarkdown(item.ans)"></p>
        <el-divider />
      </div>
    </div>
  </div>
</template>

<script>
import MarkdownIt from "markdown-it";
import axios from "axios";

/* global AMapLoader */
export default {
  name: "AboutPage",
  data() {
    return {
      qaList: [],
    };
  },
  mounted() {
    this.loadMapScript();
    this.fetchQAs();
  },
  methods: {
    renderMarkdown(md) {
      const mdParser = new MarkdownIt();
      return mdParser.render(md);
    },
    fetchQAs() {
      axios
        .get("/api/qas")
        .then((response) => {
          if (response.data.code === 200) {
            this.qaList = response.data.data.lists;
          } else {
            console.error("Failed to fetch QA list:", response.data.msg);
          }
        })
        .catch((error) => {
          console.error("Error fetching QA list:", error);
        });
    },
    loadMapScript() {
      // 检查是否已加载地图脚本
      if (
        !document.querySelector(
          'script[src="https://webapi.amap.com/loader.js"]'
        )
      ) {
        // 动态加载高德地图的安全配置脚本
        const securityScript = document.createElement("script");
        securityScript.type = "text/javascript";
        securityScript.innerHTML = `
            window._AMapSecurityConfig = {
              securityJsCode: "xxxxxxxxxxxxxxxxx",
            };
          `;
        document.head.appendChild(securityScript);

        // 动态加载高德地图API的loader脚本
        const loaderScript = document.createElement("script");
        loaderScript.src = "https://webapi.amap.com/loader.js";
        loaderScript.onload = () => {
          AMapLoader.load({
            key: "xxxxxxxxxxxxxxx", // 申请好的Web端开发者 Key，调用 load 时必填
            version: "2.0", // 指定要加载的 JS API 的版本，缺省时默认为 1.4.15
          })
            .then((AMap) => {
              this.initMap(AMap);
            })
            .catch((e) => {
              console.error(e); // 加载错误提示
            });
        };
        document.head.appendChild(loaderScript);
      } else {
        // 如果脚本已经加载，直接初始化地图
        this.initMap(window.AMap);
      }
    },
    initMap(AMap) {
      if (!AMap) {
        console.error("AMap is undefined");
        return;
      }

      // 初始化地图
      const map = new AMap.Map("container", {
        center: [116.34453, 39.9516], // 初始中心点，经纬度
        zoom: 17, // 初始缩放级别
      });

      // 添加标记
      const marker = new AMap.Marker({
        position: new AMap.LngLat(116.34453, 39.9516), // 标记位置
        title: "Beijing", // 标记标题
      });
      map.add(marker);
    },
  },
};
</script>

<style scoped>
.address-detail p{
  padding: 0;
  margin: 0;
  font-size:16px;
}

.address-detail{
  margin-bottom: 40px;
}

#container {
  width: 100%;
  height: 600px;
  margin-top: 10px;
  margin-bottom: 10px;
}

.title {
  margin-bottom: 10px;
  margin-top: 10px;
  font-size: 30px;
  font-weight: bold;
}

.question {
  font-size: 20px;
}

.answer {
  font-size: 16px;
}
</style>
