# Laboratory_Website

## 网站介绍

基于gin+vue框架的前后端分离的**实验室网站**，包含主页、成员、研究成果、奖项、记录、关于、管理等模块。

- 主页：介绍实验室的方向；
- 成员：实验室的老师、博士生、研究生、本科生、毕业生；
- 研究成果：以年划分的包含期刊论文、会议论文等分类的研究成果；
- 奖项：包含竞赛和其他两个分类；
- 记录：参加活动的照片；
- 关于：地址和QA。

## 在线预览

:clap:顺便宣传一下我们实验室

[UIS实验室](http://www.bjtuuis.cn/)

管理界面：

![image-20241103135324297](https://serena-typora-img.oss-cn-beijing.aliyuncs.com/202411031353453.png)

![image-20241103144551812](https://serena-typora-img.oss-cn-beijing.aliyuncs.com/202411031445033.png)

## 怎么启动

### mysql配置

**data.sql**

管理员账号没有创建接口，需要自己在数据库创建

### 前端

``` 
npm run serve
```

> 使用高德地图api需要申请密钥

### 后端

1. 修改配置文件

2. 运行

   ```
   go build .
   ./UISwebsite_backend
   ```

   

## TODO

- [ ] 网站访问量统计
- [ ] 批量导入导出
