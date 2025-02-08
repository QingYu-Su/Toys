# QT_SqyMathTools
这是从MFC移植到QT中的项目，详情请看[MFC实现](../MFC/)

## 安装环境
---
- 操作系统：Windows10 1903
- QT版本：Qt 5.14.1
- 编译器：MinGW_32_bit-Debug
- IDE：Qt Creator 4.11.1

## 在本地运行
---

- 下载本项目，其中包含了几个小项目：Calculator（计算器）、FunctionIndicator（函数图像显示器）、SqyMathLibrary（相关数学功能实现库）。打开方式为进入各个项目文件夹，并使用Qt Creator打开项目中的.pro文件

- 请注意：其中SqyMathLibrary为静态库，被其他项目所引用，无法独立运行

- 用户可选择Calculator与FunctionIndicator为启动程序在本地运行，但是必须先构建SqyMathLibrary生成库文件才可以
