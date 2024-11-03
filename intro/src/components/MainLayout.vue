<template>
    <el-row>
        <el-col :span="4">
            <div class="grid-content ep-bg-purple-dark" />
        </el-col>
        <el-col :span="16">
            <div>
                <AppHeader  />
                <router-view />
                <AppFooter />
            </div>
        </el-col>
        <el-col :span="4">
            <div class="grid-content ep-bg-purple-dark" />
        </el-col>
    </el-row>
</template>

<script>
import AppHeader from "./AppHeader.vue";
import AppFooter from "./AppFooter.vue";

export default {
    components: {
        AppHeader,
        AppFooter,
    },
};
</script>
