<template>
    <el-container style="height: 100vh;">
        <!-- 头部区域 -->
        <el-header
            style="display: flex; justify-content: space-between; align-items: center; background-color: #1F2D3D; color: white;  box-shadow: 0 2px 8px rgba(0,0,0,0.2);">
            <!-- 左侧标题 -->
            <div style="margin-left: 20px;font-weight: bold;font-size: 20px; ">
                UIS网站管理平台
            </div>
            <!-- 右侧用户名 -->
            <div style="margin-right: 20px;font-size: 15px;">
                Admin
            </div>
        </el-header>


        <!-- 主体区域，包括侧边栏和主视图区 -->
        <el-container>
            <!-- 左侧菜单栏 -->
            <el-aside width="150px" style="background-color: #1F2D3D;">
                <el-menu :default-active="0" class="el-menu-vertical-demo" background-color="#1F2D3D"
                    text-color="#bfcbd9" active-text-color="#409EFF">
                    <el-menu-item index="1" @click="navigateTo('/admin/member')">
                        <template #title>
                            <el-icon>
                                <user />
                            </el-icon>
                            成员管理
                        </template>
                    </el-menu-item>
                    <el-menu-item index="2" @click="navigateTo('/admin/work')">
                        <el-icon><folder-opened /></el-icon>
                        成果管理
                    </el-menu-item>
                    <el-menu-item index="3" @click="navigateTo('/admin/record')">
                        <el-icon>
                            <tickets />
                        </el-icon>
                        记录管理
                    </el-menu-item>
                    <el-menu-item index="4" @click="navigateTo('/admin/reward')">
                        <el-icon>
                            <medal />
                        </el-icon>
                        奖励管理
                    </el-menu-item>
                    <el-menu-item index="5" @click="navigateTo('/admin/issues')">
                        <el-icon>
                            <help />
                        </el-icon>
                        问题管理
                    </el-menu-item>
                </el-menu>

            </el-aside>

            <!-- 主内容区 -->
            <el-main style="max-width: 1770px;">
                <!-- 面包屑导航 -->
                <el-breadcrumb separator="/" style="margin-bottom: 20px">
                    <el-breadcrumb-item :to="{ path: '/admin' }">主界面</el-breadcrumb-item>
                    <!-- 仅在有具体子路由时显示下一级面包屑 -->
                    <el-breadcrumb-item v-if="$route.meta.breadcrumb && $route.name !== 'Admin'">{{
                        $route.meta.breadcrumb }}</el-breadcrumb-item>
                </el-breadcrumb>
                <!-- 主内容将根据路由展示不同的组件 -->
                <router-view />
            </el-main>
        </el-container>
    </el-container>
</template>

<script>
import { User, FolderOpened, Tickets, Medal, Help } from '@element-plus/icons-vue';
export default {
    name: 'MenuLayout',
    components: {
        User,  // 导入的图标组件
        FolderOpened,
        Tickets,
        Medal,
        Help

    },
    methods: {
        navigateTo(route) {
            this.$router.push(route);
        },
        getDefaultActive() {
            // 根据路由决定默认激活的菜单项
            if (this.$route.name === 'Admin') {
                return '';  // 如果是Admin页面，没有菜单项应被激活
            }
            return this.$route.name;  // 否则，返回当前路由名
        }
    }
}


</script>

<style scoped>
.el-header,
.el-aside {
    background-color: #1F2D3D;
    color: white;
}

.el-aside {
    box-shadow: -2px 0 8px rgba(0, 0, 0, 0.2);
}

.el-menu-vertical-demo:not(.el-menu--collapse) {
    border-right: 0;
}
</style>