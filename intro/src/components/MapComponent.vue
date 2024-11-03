<template>
    <div id="container" style="width: 100%; height: 600px;"></div>
  </template>
  
  <script>
  /* global AMapLoader */
  export default {
    name: "MapComponent",
    mounted() {
      // 动态加载高德地图的安全配置脚本
      const securityScript = document.createElement("script");
      securityScript.type = "text/javascript";
      securityScript.innerHTML = `
        window._AMapSecurityConfig = {
          securityJsCode: "9657278e69186554bc9a289b99a3077f",
        };
      `;
      document.head.appendChild(securityScript);
  
      // 动态加载高德地图API的loader脚本
      const loaderScript = document.createElement("script");
      loaderScript.src = "https://webapi.amap.com/loader.js";
      loaderScript.onload = () => {
        AMapLoader.load({
          key: "737fd7cdfcdab3ae69f5c03870964858", //申请好的Web端开发者 Key，调用 load 时必填
          version: "2.0", //指定要加载的 JS API 的版本，缺省时默认为 1.4.15
        })
          .then((AMap) => {
            this.$emit("map-loaded", AMap);
          })
          .catch((e) => {
            console.error(e); //加载错误提示
          });
      };
      document.head.appendChild(loaderScript);
    },
  };
  </script>
  
  <style scoped>
  #container {
    width: 100%;
    height: 100%;
  }
  </style>
  