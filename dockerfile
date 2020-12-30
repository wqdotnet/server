#基于的镜像不是golang(733M),而是alpine(4.14M)
FROM alpine:latest
 
#ENV 设置环境变量
#ENV PATH /usr/local/nginx/sbin:$PATH

#WORKDIR 相当于cd
WORKDIR $GOPATH/src/login

#ADD  文件放在当前目录下，拷过去会自动解压
ADD config　 $GOPATH/src/config


#RUN 执行以下命令 
#RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

#EXPOSE 指定容器要打开的端口 格式为 EXPOSE <port> [<port>...]
EXPOSE 8080 3344

#CMD 运行以下命令
#CMD ["nginx"]

#ENTRYPOINT ["executable", "param1", "param2"]  
#ENTRYPOINT command param1 param2
ENTRYPOINT ["./server"]