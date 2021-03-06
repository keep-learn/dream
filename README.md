[toc]

## 项目简介
* 动态读取用户输入并生成指定格式的Excel
* 程序是一个命令行小工具，安全、可靠、高性能

## 技术栈
* 日志使用 [zap](https://github.com/uber-go/zap)
* 日志分割使用 [lumberjack](https://github.com/natefinch/lumberjack)
* 配置使用 [viper](https://github.com/spf13/viper)
* Excel使用 [excelize](https://github.com/qax-os/excelize)
* 命令行使用 [cobra](https://github.com/spf13/cobra")
* 命令行输出 [color](https://github.com/fatih/color)

## 安装
```shell
# 日志和excel的路径（可以在配置文件修改）
go mod tidy
mkdir -p /tmp/log
mkdir -p /tmp/excel
make build
bin/dream export "X 信息、Y 信息；甲类、乙类"
```

## 演示
输入：
```shell
bin/dream export "X 信息、Y 信息；甲类、乙类"
```
1. 控制台结果
![demo](./images/demo02.png)
   
2. 生成的excel截图：
![demo](./images/img_1.png)


## 项目结构

```api
.
├── Makefile
├── README.md
├── business  //业务逻辑目录
│   ├── dto //结构体目录
│   │   ├── export.go
│   │   └── input.go
│   └── service //service服务
│       ├── export_excel.go
│       ├── export_excel_test.go
│       ├── user_input.go
│       └── user_input_test.go
├── config //配置文件目录
│   └── default.toml
├── go.mod
├── go.sum
├── images //图片、演示视频目录
│   ├── Kapture 2022-03-22 at 23.37.26.gif
│   ├── img.png
│   └── img_1.png
├── main.go
└── pkg //内部包
    ├── conf  //配置文件包
    │   └── conf.go
    └── log //日志包
        └── log.go
```

[//]: # (### 演示视频)

[//]: # (![demo]&#40;./images/Kapture%202022-03-22%20at%2023.37.26.gif&#41;)


### TODO
- [x] 使用viper，项目读取config配置
- [x] 使用lumberjack进行日志拆分
- [x] 完善 README 文件
- [x] 输入、输出新增友好提示、错误纠正功能
- [x] 日志和Excel文件分开存储
- [x] 日志新增时间、代码行数、日志级别字段
- [x] 日志新增trace_id跟踪
- [ ] 应用Validator进行参数校验

