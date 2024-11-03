<!-- AppHeader -->
<template>
    <el-row style="padding: 20px 0">
        <el-col :span="10">
            <div class="header-content">
                <div class="logo">
                    <img src="../assets/images/R-C.gif" class="logo" />
                </div>
                <div class="logo-title">
                    <h2>泛在智能软件实验室</h2>
                    <p lang="en">Ubiquitous Intelligent Software Laboratory</p>
                </div>
            </div>
        </el-col>
        <el-col :span="2"></el-col>
        <el-col :span="12">
            <el-menu :ellipsis="false" mode="horizontal" class="el-menu-demo" active-text-color="grey">
                <el-menu-item index="1">
                    <router-link to="/" class="no-underline">主页</router-link>
                </el-menu-item>
                <el-menu-item index="2">
                    <router-link to="/members" class="no-underline">成员</router-link>
                </el-menu-item>
                <el-menu-item index="3">
                    <router-link to="/work" class="no-underline">研究成果</router-link>
                </el-menu-item>
                <el-menu-item index="4">
                    <router-link to="/reward" class="no-underline">奖项</router-link>
                </el-menu-item>
                <el-menu-item index="5">
                    <router-link to="/record" class="no-underline">记录</router-link>
                </el-menu-item>
                <el-menu-item index="6">
                    <router-link to="/about" class="no-underline">关于</router-link>
                </el-menu-item>
            </el-menu>
        </el-col>
    </el-row>
</template>

<style scoped>
.header-content {
    display: flex;
    align-items: center;
}

.logo {
    height: 80px;
}

.logo-title {
    margin-left: 15px;
}

.logo-title h2 {
    font-size: 40px;
    margin: 0;
    margin-bottom: 10px;
    text-align: center;
}

.logo-title p {
    font-size: 18px;
    margin: 0;
    text-align: center;
}

.el-menu-demo {
    background-color: transparent;
    box-shadow: none;
    margin-top: 3vh;
}

.el-menu-demo .el-menu-item {
    font-size: 23px;
    text-align: center;
}

.no-underline {
    text-decoration: none;
    color: inherit; /* 保持原有文字颜色 */
}

.logo-title h2,
.logo-title p {
    white-space: nowrap;  /* 防止文本换行 */
    overflow: hidden;     /* 隐藏超出部分的文本 */
    text-overflow: ellipsis; /* 超出部分显示省略符号 */
}

</style>

