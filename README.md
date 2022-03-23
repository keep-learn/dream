## 项目概要
* 日志使用 [zap](https://github.com/uber-go/zap)
* 日志分割使用 [lumberjack](gopkg.in/natefinch/lumberjack.v2)
* 配置使用 [viper](https://github.com/spf13/viper)
* Excel使用 [excelize](https://github.com/qax-os/excelize)
* 命令行使用go自带的 flag

## Install
```shell
# 日志和excel的路径（可以在配置文件修改）
go mod tidy
mkdir -p /tmp/test
go run main.go -export "X 信息、Y 信息；甲类、乙类"
```

## 输入、输出演示
输入：
```shell
go run main.go -export "X 信息、Y 信息；甲类、乙类"
```

生成的excel截图：
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
- [ ] 完善 README 文件
- [ ] 输入、输出新增友好提示、错误纠正功能
- [ ] 日志和Excel文件分开存储

