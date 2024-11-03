<!-- AppFooter -->
<template>
  <el-row gutter="15" style="margin-top: 5vh">
    <el-col :span="8">
      <a href="" target="_blank">
        <div class="custom-card">
          <img src="../assets/images/R-C.gif" alt="校徽" class="logo" />
        </div>
      </a>
    </el-col>
    <el-col :span="8">
      <a href="" target="_blank">
        <div class="custom-card">
          <img src="../assets/images/R-C.gif" alt="学院" class="logo" />
        </div>
      </a>
    </el-col>
    <el-col :span="8">
      <a href="" target="_blank">
        <div class="custom-card">
          <img src="../assets/images/R-C.gif" alt="北京交通大学研究生院" class="logo" />
        </div>
      </a>
    </el-col>
  </el-row>
  <div style="
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-top: 20px;
      margin-bottom: 5px;
    ">
    <a lang="en" href="/login" style="text-decoration: none;">Members Only</a>
    <div style="display: flex; align-items: center">
      <img src="../assets/images/ICP.png" style="margin-right: 10px;width: 20px;height: auto;" />
      <a href="https://beian.mps.gov.cn/#/query/webSearch" rel="noreferrer" target="_blank" style="text-decoration: none;margin-right: 10px">京公网安备xxxxxxxxxxx</a>
      <a href="https://beian.miit.gov.cn/" target="_blank" style="text-decoration: none;">京ICP备xxxxxxx号</a>
    </div>
  </div>
</template>

<script>
export default {
  name: "AppFooter",
};
</script>

<style scoped>
.custom-card {
  width: 100%;
  height: 14vh;
  margin: 0;
  padding: 0;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: white;
}

.logo {
  width: 100%;
  max-height: 100%;
  object-fit: cover;
  /* 或者 'cover' 根据需要 */
}
</style>
