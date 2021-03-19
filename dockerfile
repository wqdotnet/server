#基于的镜像不是golang(733M),而是alpine(4.14M)   #scratch 空镜像
FROM alpine:latest

#ENV 设置环境变量
#ENV PATH /usr/local/nginx/sbin:$PATH

#[WORKDIR 进入docker内文件夹] 相当于cd $GOPATH/..
WORKDIR /home/slgserver

#[ADD 本地文件 docker文件地址]   文件放在当前目录下，拷过去会自动解压
ADD slgserver  /home/slgserver
#ADD cfg.yaml  /home/slgserver/cfg.yaml
#ADD config /home/slgserver/config

#RUN 执行以下命令 
#RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
#RUN go build 

#EXPOSE 指定容器要打开的端口 格式为 EXPOSE <port> [<port>...]
EXPOSE 8080 8081 3344

#CMD 运行以下命令
#CMD ["nginx"]

#ENTRYPOINT ["command", "param1", "param2"]
ENTRYPOINT ["./slgserver","start","--config=/home/slgserver/config/cfg.yaml"]

#=================================================================================================================================================
#构建镜像 [ -t name:tag ]
#sudo docker build -t slgdocker .

#运行容器 -v[本地配置地址 :docker内读取配置固定地址"/home/slgserver/config"]
#sudo docker run -t -i -d -v /mnt/e/dockerconfig:/home/slgserver/config -p 3344:3344 -p 8080:8080 -p 8081:8081 --name slggame  slgdocker:latest

#进入容器 
#sudo docker exec -it slggame /bin/sh

#保存 加载 镜像
#sudo docker save slgdocker:latest  -o  /home/wq/slgdocker:latest.tar
#sudo docker load -i slgdocker:latest.tar
#=================================================================================================================================================