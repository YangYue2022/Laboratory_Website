import { createApp } from "vue";
import App from "./App.vue";
import router from "@/router";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";

import { VueMasonryPlugin } from 'vue-masonry';


// 创建Vue应用
const app = createApp(App);

// 使用ElementPlus和VueMasonryPlugin插件
app.use(ElementPlus);
app.use(VueMasonryPlugin);


// 使用router
app.use(router);

// 挂载Vue应用
app.mount("#app");
