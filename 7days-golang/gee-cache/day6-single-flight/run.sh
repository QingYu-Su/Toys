#!/bin/bash
# 设置脚本使用的 Shell 解释器为 Bash

# 定义一个 trap 命令，用于在脚本退出时执行清理操作。
# 当脚本接收到 EXIT 信号时，执行以下操作：
# 1. 删除生成的可执行文件 "server"。
# 2. 使用 "kill 0" 终止当前脚本启动的所有子进程。
trap "rm server; kill 0" EXIT

# 使用 Go 编译当前目录下的程序，生成一个名为 "server" 的可执行文件。
go build -o server

# 启动三个缓存服务实例，分别监听不同的端口：
# - 第一个实例监听端口 8001。
# - 第二个实例监听端口 8002。
# - 第三个实例监听端口 8003，并启用 API 服务（监听端口 9999）。
./server -port=8001 &
./server -port=8002 &
./server -port=8003 -api=1 &

# 等待 2 秒，确保所有服务实例已经启动并准备好。
sleep 2

# 输出测试开始的提示信息。
echo ">>> start test"

# 启动多个后台 curl 请求，访问 API 服务并查询键 "Tom" 的值。
# 每个请求都会通过 API 服务转发到缓存服务实例，验证缓存服务的分布式功能。
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &

# 等待所有后台进程结束。
# 这包括启动的缓存服务实例和 curl 请求。
wait