-- 删除已存在的表格
DROP TABLE IF EXISTS uis_work;
DROP TABLE IF EXISTS uis_reward;
DROP TABLE IF EXISTS uis_qa;
DROP TABLE IF EXISTS uis_member;
DROP TABLE IF EXISTS uis_activity;
DROP TABLE IF EXISTS uis_record;
DROP TABLE IF EXISTS uis_auth;


CREATE TABLE uis_work (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自增ID',  -- 主键，自增ID
    type ENUM('期刊论文', '会议论文', '专著', '软著/专利', '其他') NOT NULL COMMENT '作品类型，枚举值：1=期刊论文, 2=会议论文, 3=专著, 4=软著/专利, 5=其他',  -- 枚举类型，表示作品的类别
    year YEAR NOT NULL COMMENT '发布年份',  -- 年份，使用YEAR类型存储年份信息
    `desc` VARCHAR(500) NOT NULL COMMENT '作品描述'  -- 描述，变长字符串，适用于较短的文本

) COMMENT='存储作品信息，包括类型、年份和描述';

-- 为type字段添加索引
CREATE INDEX idx_type ON uis_work (type);

-- 为year字段添加索引
CREATE INDEX idx_year ON uis_work (year);

CREATE TABLE uis_reward (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自增ID',  -- 主键，自增ID
    type ENUM('竞赛', '其他') NOT NULL COMMENT '作品类型，枚举值：1=竞赛, 2=其他',  -- 枚举类型，表示奖励的类别
    year YEAR NOT NULL COMMENT '发布年份',  -- 年份，使用YEAR类型存储年份信息
    name VARCHAR(255) NOT NULL COMMENT '竞赛名称',  -- 描述，变长字符串
    project VARCHAR(255) NOT NULL COMMENT '项目名称',  -- 描述，变长字符串
    winner  VARCHAR(255) NOT NULL COMMENT '获奖人' -- 用,分隔开
) COMMENT='存储奖励信息，包括类型、年份、获奖信息等';

-- 为type字段添加索引
CREATE INDEX idx_type ON uis_reward (type);

-- 为year字段添加索引
CREATE INDEX idx_year ON uis_reward (year);

CREATE TABLE uis_qa (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自增ID',  -- 主键，自增ID
    que VARCHAR(255) NOT NULL COMMENT '问题',  
    ans  TEXT NOT NULL COMMENT '答案' 

) COMMENT='q&a信息，包括问题、答案';

CREATE TABLE uis_member(
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自增ID',  -- 主键，自增ID
    photo_url VARCHAR(255) DEFAULT '' COMMENT '图片地址',  -- 照片地址
    name VARCHAR(255) NOT NULL COMMENT '姓名',  -- 姓名
	identity ENUM('教职员', '学生') NOT NULL COMMENT '作品类型，枚举值：1=教职员 2=学生',  -- 枚举类型
    title ENUM('教授', '副教授','讲师','工程师','博士研究生','硕士研究生','本科生','毕业生') NOT NULL COMMENT '枚举类',  -- 职称或学历
    phone_number VARCHAR(20) DEFAULT '' COMMENT '手机号码',  -- 手机号码
    email VARCHAR(255) DEFAULT '' COMMENT '邮箱地址',  -- 邮箱地址
    description TEXT COMMENT '个人介绍'  -- 个人介绍
) COMMENT='成员信息表';

CREATE TABLE uis_activity(
	id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自增ID',  -- 主键，自增ID
    year YEAR NOT NULL COMMENT '年份', -- 年份，使用YEAR类型存储年份信息
    name VARCHAR(255) NOT NULL COMMENT '活动名称' -- 活动名称
)COMMENT='活动信息表';

CREATE TABLE uis_record(
	id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自增ID',  -- 主键，自增ID
    photo_url VARCHAR(255) DEFAULT '' COMMENT '图片地址',  -- 照片地址
    activity_id INT DEFAULT '0' COMMENT '活动ID',
    name VARCHAR(255) NOT NULL COMMENT '图片名称', -- 图片名称
    visibility BOOL DEFAULT '0' NOT NULL COMMENT '可见性'
)COMMENT='记录信息表';

CREATE TABLE uis_auth (
  id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自增ID',  -- 主键，自增ID
  username VARCHAR(255) NOT NULL COMMENT '用户名',
  password VARCHAR(255) NOT NULL COMMENT '密码'
) COMMENT='用户认证表';


