# game server

💡❌💙 💔💜 💚💬⭐️⚠️💃🏻📄📚🛠 😎 🔧 🐭🐮🐯🐇🐉🐍🐎🐑🐒🐔🐶🐷

### 基于  [halturin/ergo](https://github.com/halturin/ergo) 以erlang otp 模型方式组织创建的游戏服务器解决方案

##### 服务启动时会创建3个节点 

- gatewayNode 用户连接后创建网关连接
- serverNode  创建游戏公共服务 [cmdGenserver]
- dbNode      用做数据落地

serverNode 节点启动会创建一个 cmdGenserver 用于接收外部发送过来的命令，以
便于从内部 获取信息、更新配置、关闭服务

server运行时 执行 cmd [state|stop|debug|reloadcfg] 命令 

会在创建一个 debugNode 节点去接连服务器内部 serverNode 节点下的 cmdGenserver 发送命令消息


 



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
##### 使用 [spf13/cobra](https://github.com/spf13/cobra)  创建的服务器命令
 


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
