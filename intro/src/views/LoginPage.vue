<template>
    <div class="login-page">
        <div class="login-container">
            <p class="login-title">登录</p>
            <form @submit.prevent="handleSubmit">
                <el-row style="margin-bottom: 20px">
                    <input type="text" id="username" v-model="username" required class="form-input" placeholder="请输入用户名" />
                </el-row>
                <el-row style="margin-bottom: 20px">
                    <input type="password" id="password" v-model="password" required class="form-input" placeholder="请输入密码" />
                </el-row>
                <button type="submit" class="submit-button">登录</button>
            </form>
            <div class="social-login">
                <p class="social-login-title">使用社交账号登录</p>
                <div class="icons">
                    <img :src="feishuIcon" alt="Feishu Login" @click="handleFeishuLogin" />
                    <img :src="wechatIcon" alt="WeChat Login" @click="handleWeChatLogin" />
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
export default {
    name: "LoginPage",
    data() {
        return {
            username: "",
            password: "",
            feishuIcon: require("@/assets/images/feishu-icon.png"),
            wechatIcon: require("@/assets/images/wechat-icon.png"),
        };
    },
    methods: {
        handleSubmit() {
            axios.get('/api/auth', {
                params: {
                    username: this.username,
                    password: this.password
                }
            })
            .then(response => {
                if (response.data.code === 200) {
                    const token = response.data.data.token;
                    localStorage.setItem('token', token);
                    this.$message.success('登录成功！');
                    this.$router.push({ name: 'Admin' });
                } else {
                    this.$message.error('登录失败，请检查用户名和密码');
                }
            })
            .catch(error => {
                console.error('Error logging in:', error);
                this.$message.error('登录失败，请稍后再试');
            });
        },
        handleFeishuLogin() {
            console.log("Logging in with Feishu");
        },
        handleWeChatLogin() {
            console.log("Logging in with WeChat");
        },
    },
};
</script>

<style scoped>
.login-page {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
}

.login-container {
    background-color: #fff;
    padding: 30px;
    border-radius: 15px;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
    width: 360px;
}

.login-title {
    font-size: 24px;
    color: #333;
    text-align: center;
    margin-bottom: 20px;
}

.form-input {
    width: 100%;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    font-size: 16px;
}

.submit-button {
    display: block;
    width: 100%;
    padding: 12px;
    background-color: #007bff;
    border: none;
    border-radius: 5px;
    color: white;
    font-size: 18px;
    cursor: pointer;
    margin-top: 20px;
}

.submit-button:hover {
    background-color: #0056b3;
}

.social-login {
    margin-top: 20px;
    text-align: center;
}

.social-login-title {
    font-size: 16px;
    color: #666;
}

.icons {
    margin-top: 10px;
    display: flex;
    justify-content: center;
    gap: 20px;
}

.icons img {
    width: 40px;
    height: 40px;
    cursor: pointer;
}

.icons img:hover {
    transform: scale(1.1);
}
</style>
