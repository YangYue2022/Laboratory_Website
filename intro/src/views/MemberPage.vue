<template>
    <div>
        <div class="members-page" v-if="teachers.length">
            <h2 style="margin: 2px 0">教职员</h2>
            <div class="members-container">
                <div class="member-card" v-for="member in teachers" :key="member.name">
                    <el-divider style="margin: 10px 0;"/>
                    <div class="member-details">
                        <img :src="member.photo_url" alt="member.name" class="member-image">
                        <div class="member-info">
                            <p class="member-name">{{ member.name }}</p>
                            <p class="member-title">{{ member.title }}</p>
                            <p v-html="renderMarkdown(member.description)"></p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="members-page" v-if="boshi.length">
            <h2 style="margin: 2px 0">博士研究生</h2>
            <div class="members-container">
                <div class="member-card" v-for="member in boshi" :key="member.name">
                    <el-divider style="margin: 10px 0;"/>
                    <div class="member-details">
                        <img :src="member.photo_url" alt="member.name" class="member-image">
                        <div class="member-info">
                            <p class="member-name">{{ member.name }}</p>
                            <p v-html="renderMarkdown(member.description)"></p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="members-page" v-if="shuoshi.length">
            <h2 style="margin: 2px 0">硕士研究生</h2>
            <div class="members-container">
                <div class="member-card" v-for="member in shuoshi" :key="member.name">
                    <el-divider style="margin: 10px 0;"/>
                    <div class="member-details">
                        <img :src="member.photo_url" alt="member.name" class="member-image">
                        <div class="member-info">
                            <p class="member-name">{{ member.name }}</p>
                            <p v-html="renderMarkdown(member.description)"></p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="members-page" v-if="benke.length">
            <h2 style="margin: 2px 0">本科生</h2>
            <div class="members-container">
                <div class="member-card" v-for="member in benke" :key="member.name">
                    <el-divider style="margin: 10px 0;"/>
                    <div class="member-details">
                        <img :src="member.photo_url" alt="member.name" class="member-image">
                        <div class="member-info">
                            <p class="member-name">{{ member.name }}</p>
                            <p v-html="renderMarkdown(member.description)"></p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="members-page" v-if="biye.length">
            <h2 style="margin: 2px 0">毕业生</h2>
            <div class="members-container">
                <div class="member-card" v-for="member in biye" :key="member.name">
                    <el-divider style="margin: 10px 0;"/>
                    <div class="member-details">
                        <img :src="member.photo_url" alt="member.name" class="member-image">
                        <div class="member-info">
                            <p class="member-name">{{ member.name }}</p>
                            <p v-html="renderMarkdown(member.description)"></p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import MarkdownIt from 'markdown-it';

export default {
    name: 'MemberPage',
    data() {
        return {
            teachers: [],
            boshi: [],
            shuoshi: [],
            benke: [],
            biye: []
        };
    },
    created() {
        this.fetchMembers();
    },
    methods: {
        fetchMembers() {
            axios.get('/api/members')
                .then(response => {
                    const members = response.data.data.lists;
                    this.categorizeMembers(members);
                })
                .catch(error => {
                    console.error('Error fetching members:', error);
                });
        },
        categorizeMembers(members) {
            members.forEach(member => {
                if (member.identity === '教职员') {
                    this.teachers.push(member);
                } else if (member.title === '博士研究生') {
                    this.boshi.push(member);
                } else if (member.title  === '硕士研究生') {
                    this.shuoshi.push(member);
                } else if (member.title === '本科生') {
                    this.benke.push(member);
                } else if (member.title  === '毕业生') {
                    this.biye.push(member);
                }
            });
        },
        renderMarkdown(md) {
            const mdParser = new MarkdownIt();
            return mdParser.render(md);
        }
    }
};
</script>

<style scoped>
.members-page {
    padding: 20px;
}

h1 {
    text-align: center;
    margin-bottom: 10px;
}

.members-container {
    display: flex;
    flex-wrap: wrap;
    gap: 20px;
}

.member-card {
    width: calc(50% - 20px);
    padding: 20px;
    box-sizing: border-box;
}

.member-details {
    display: flex;
}

.member-image {
    width: 120px;
    height: 180px;
    object-fit: cover;
    border-radius: 8px;
    margin-right: 20px;
    margin-top: 5px;
}

.member-info {
    flex: 1;
}

.member-info p {
    margin:  0 0;
    font-size: 14px
}

.member-name{
    font-size: 25px !important;
    margin-bottom: 5px; 
}

.member-title{
    font-size: 18px !important;
    margin: 0;
}
</style>
