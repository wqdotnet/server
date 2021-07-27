# game server

💡❌💙 💔💜 💚💬⭐️⚠️💃🏻📄📚🛠 😎 🔧 🐭🐮🐯🐇🐉🐍🐎🐑🐒🐔🐶🐷

## 🔨 command
- server start    &emsp;&emsp;&emsp;    start game server   
- server protobuf &emsp;    protobuf 协议生成  
- server clean&emsp;&emsp;&emsp;  清理数据



### ✅ 安全退出
    ctrl + | 
  

  
  
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
