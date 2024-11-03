<template>
    <el-row gutter="15" style="margin-bottom: 1vh;margin-top: 20px;">
        <el-col :span="12">
            <a  target="_blank">
                <div class="custom-card">
                    <img src="../assets/images/R-C.gif" alt="智慧交通" class="card-image">
                    
                </div>
            </a>
        </el-col>
        <el-col :span="12">
            <a target="_blank">
                <div class="custom-card">
                    <img src="../assets/images/R-C.gif" alt="虚拟现实" class="card-image">
                </div>
            </a>
        </el-col>
    </el-row>
    <el-row gutter="15">
        <el-col :span="12">
            <a  target="_blank">
                <div class="custom-card">
                    
                    <img src="../assets/images/R-C.gif" alt="无线通信" class="card-image">
                
                </div>
            </a>
        </el-col>
        <el-col :span="12">
            <a  target="_blank">
                <div class="custom-card">
                    <img src="../assets/images/R-C.gif" alt="人机交互" class="card-image">
                </div>
            </a>
        </el-col>
    </el-row>
    <el-row style="margin-top: 35px;">
        <div style="display: flex; align-items: baseline;">
            <h1 lang="en" style="margin-right: 8px;">xxx</h1>
            <h1>实验室</h1>
        </div>
    </el-row>
    <el-row style="margin-top: 0;">
        <div class="text-container">
            <p>待填充</p>
            <p>我们感兴趣的研究领域包括下列主题：</p>
            <ul>
                <li>xxxx</li>
                <li>xxxx</li>
                <li>xxxx</li>
            </ul>
        </div>
    </el-row>


</template>

<style scoped>
.custom-card {
    width: 100%;
    height: 38vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    background-color: white;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
    overflow: hidden;
}

.card-image {
    width: 100%;
    height: 100%;
}

.card-title {
    font-size: 18px;
    color: #333;
    padding: 5px;
}

.text-container {
    font-size: 18px;
    line-height: 1.5;
}

h1 {
    margin-top: 0;    /* 减少上边距 */
    margin-bottom: 0; /* 减少下边距 */
}
</style>
