# go工程实践

VERSIOVN MV0.1
git branch xxx      # 创建分支
git checkout        # 切换分支
git branch          # 查询当前分支

克隆V0版本：git clone -b V0 https://github.com/yhmain/go-project-example
提交一套流程的命令：参考网址：https://www.cnblogs.com/lxpblogs/p/15504450.html
进入当前项目目录下依次执行命令：
git add *
git commit -m "你自己的一些说明"
git push origin 分支名

## concurrence包
* 并发编程

## test包
* 单元测试 
* mock

## benchmark包
* 基准测试

---
---
## 社区话题页面需求描述
* 展示话题（标题，文字描述）和回帖列表
* 暂时不考虑前端页面实现，仅仅实现一个本地web服务
* 话题和回帖数据用文件存储
### data包
元数据文件
### repository包
数据层
### service包
服务层
### controller包
视图层
### go.mod文件
go module 依赖配置管理
### server.go
web服务main入口文件