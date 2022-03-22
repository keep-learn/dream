## 项目概要
* 日志使用 [zap](https://github.com/uber-go/zap)
* 日志分割使用 [lumberjack](gopkg.in/natefinch/lumberjack.v2)
* 配置使用 [viper](https://github.com/spf13/viper)
* Excel使用 [excelize](https://github.com/qax-os/excelize)
* 命令行使用go自带的 flag

## Install
```shell
# 日志和excel的路径（可以在配置文件修改）
mkdir -p /tmp/test
```


## 输入、输出演示

输入：
```shell
go run main.go -export "X 信息、Y 信息；甲类、乙类"
```

生成的excel截图：
![demo](./images/img_1.png)


### TODO
- [x] 使用viper，项目读取config配置
- [x] 使用lumberjack进行日志拆分
- [ ] 完善细节...
