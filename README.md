# game server

💡❌💙 💔💜 💚💬⭐️⚠️💃🏻📄📚🛠 😎 🔧 🐭🐮🐯🐇🐉🐍🐎🐑🐒🐔🐶🐷

## 🔨 command
Available Commands:
-  clean       &emsp;&emsp;&emsp;清理数据
-  completion  &emsp;生成补全脚本
-  debug       &emsp;&emsp;&emsp;控制台
-  pb  [int] [obj]         &emsp;生成protobuf 
-  reloadcfg   &emsp;&emsp;&emsp;重新加载配置
-  start       &emsp;&emsp;&emsp;启动服务
-  state       &emsp;&emsp;&emsp;获取服务器运行状态
-  stop        &emsp;&emsp;&emsp;关闭服务器

 


### ✅ 安全退出
    ctrl + | 
  

### 📄 协议
|  2Byte (包长)  | 2Byte  |  2Byte | message|
|  ----  | ----  |----  |----  |
| 4Byte+ len(消息体)  | 模块ID | 方法ID | 消息体|


  
## 🛠 构建镜像 
```
sudo docker build -t gamedocker .
```

## 🏃 运行容器  
```
sudo docker run -t -i -d -v /mnt/e/dockerconfig:/home/server/config -p 3344:3344 -p 8080:8080 -p 8081:8081 --name gameserver  gamedocker:latest
```

## 📝 进入容器 
```
sudo docker exec -it gameserver /bin/sh
```

## 📥 保存镜像
```
sudo docker save gamedocker:latest  -o  /home/wq/gamedocker:latest.tar
```
## 💡  加载镜像
```
sudo docker load -i gamedocker:latest.tar
```
