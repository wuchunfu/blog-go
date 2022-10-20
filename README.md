## 博客介绍

<p align=center>
  <a href="https://www.spxzx.xyz">
    <img src="https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/user/admin.jpg" width="200" hight="200" alt="水平线之下的个人博客" style="border-radius: 50%">
  </a>
</p>

<p align=center>
   基于Gin + Vue 开发的前后端分离博客
</p>

<p align="center">
   <a target="_blank" href="https://github.com/spxzx/blog">
      <img src="https://img.shields.io/badge/Go-1.19-blue"/>
      <img src="https://img.shields.io/badge/Gin-v1.8.1-blue"/>
      <img src="https://img.shields.io/badge/Casbin-v2.56.0-blue"/>
      <img src="https://img.shields.io/badge/mysql-5.7-blue"/>
      <img src="https://img.shields.io/badge/GORM-v1.24.0-blue"/>
      <img src="https://img.shields.io/badge/redis-5.0.7-red"/>
      <img src="https://img.shields.io/badge/vue-v2.X-green"/>
    </a>
</p>

​             [在线地址](#在线地址) | [目录结构](#目录结构) | [项目介绍](#项目介绍) | [技术介绍](#技术介绍) | [运行环境](#运行环境) | [开发环境](#开发环境) | [项目截图](#项目截图) | [项目总结](#项目总结) 

** 本项目一开始写的代码比较混乱，后面有时间会重新将逻辑重写一遍 **

## 在线地址

**项目链接：** [www.spxzx.xyz](https://www.spxzx.xyz)

**后台链接：** [admin.spxzx.xyz](https://admin.spxzx.xyz)

测试账号：test@qq.com，密码：1，可登入后台查看。

**Github地址：** [https://github.com/spxzx/blog](https://github.com/spxzx/blog)



**需要更多信息请移步至 [X1192176811/blog](https://github.com/X1192176811/blog)**

**需要更多信息请移步至 [X1192176811/blog](https://github.com/X1192176811/blog)**

**需要更多信息请移步至 [X1192176811/blog](https://github.com/X1192176811/blog)**

## 目录结构

前端代码位于web下，blog为前台，admin为后台。

SQL文件位于根目录下的blog.sql，需要mysql5以上应该都兼容。

可直接导入该项目于本地，修改后端配置文件中的数据库等连接信息，项目中使用到的关于腾讯云对象存储功能和一些其他功能需要自己开通。

**先运行后端项目，再启动前端项目，前端项目配置由后端动态加载。** 

```
blog
├── api/v1	      	--  API
├── config        	--  配置模块
├── dao           	--  数据库模块
├── interal         --  没什么东西就了一个Set数据结构
├── model         	--  entity、dto、vo都放这里了
├── router        	--  路由模块
│   └── middleware  --  中间件
├── service       	--  服务模块
├── util        	--  工具模块
└── web            	--  前端资源
    ├── backend     --  打包好的后台资源文件
    ├── blog        --  打包好的前台资源文件
    └── vue         --  vue代码
```

## 项目介绍

原来写的博客项目由于前端技术比较渣，写的很简陋，主要的模块就一个文章模块，所以后来找到了风丶宇的博客代码，带了前端的代码（主要是能够读懂），而且前端Butterfly主题设计我很喜欢，所以就打算使用go完成后端代码，为了更方便使用go编写所以我对前端 js 代码进行了一些小修改。

这个项目原Java版本地址为[X1192176811/blog](https://github.com/X1192176811/blog)，个人在其前端代码基础上相当于重构了后端（大部分功能基本都已完成）。



​    **前端的介绍请移步至[X1192176811/blog](https://github.com/X1192176811/blog)查看详情**

- 后端使用Gin框架完成，ORM框架用的是GORM
- 认证和权限管理采用了 jwt + casbin 完成
- 支持权限的动态修改，采用的是RBAC模型
- 使用Redis实现点赞、统计用户等功能
- 前后端分离部署，现在暂时用的是Nginx
- 后面会继续完善一些功能

## 技术介绍

**前端：** vue + vuex + vue-router + axios + vuetify + element + echarts

**后端：** Go + Gin + GORM + Mysql + Redis + nginx

**其他：** 腾讯云人机验证、腾讯云对象存储

## 运行环境

**服务器：** 腾讯云2核2G Ubuntu20.04

**对象存储：** 腾讯云OSS

## 开发环境

| 开发工具                      | 说明                                               |
| ----------------------------- | -------------------------------------------------- |
| GoLand                        | Go开发工具IDE                                      |
| VSCode                        | 轻量级开发工具IDE(最近体验了Fleet，一点都不轻量级) |
| Another Redis Desktop Manager | Redis远程连接工具                                  |
| X-shell                       | Linux远程连接工具                                  |
| WinSCP                        | Linux文件上传工具                                  |

| 开发环境      | 版本  |
| ------------- | ----- |
| GO            | 1.19  |
| MySQL         | 5.7   |
| Redis         | 5.0.7 |

## 项目截图

![QQ截图1.png](https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/U~90%28FZAXS7%5BNQY9BLX8BKI.png)

![QQ截图2.jpg](https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/HY2%24%60M2T10%7DR76%7DQ6G0A~T2.png)

![QQ截图3.png](https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/8O72WZRH%29%7D9T%24I%5BXU7E3%28RQ.png)

## 项目总结

首先十分感谢 风丶宇 这个开源的博客。

个人感觉这个博客作为入门项目是很好的，在这里面使用到的技术覆盖也挺多，后面有时间的话还会加入Elasticsearch、mq等。

个人认为在Go中主要难点在于权限管理、监听和改变用户状态这块，毕竟Go的生态并没有Java那么强。
