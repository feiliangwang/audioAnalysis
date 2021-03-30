# audioAnalysis
录音文件采集工具

本项目支持通过http上传文件，并在数据库中生成对应的记录，供后续查询分析使用。同时提供http的文件列表，支持查看和下载文件

项目地址https://github.com/feiliangwang/audioAnalysis.git



## 编译

本项目使用go编写，go版本>=1.14

```
go build -o audio
```



## 运行

```
./audio -c config.yaml
```

配置文件参考example.config.yaml,其中fileDir为文件存放根目录已经http文件列表根目录

命令行参数可运行-h进行查看

```
./audio -c -h
Usage of ./audio:
  -c string
        config file path,defalut is current dir config.yaml (default "config.yaml")
  -h    this help

```



## 测试

本项目已对接口做了基本的测试。

测试用例详见 

[audio.postman_collection.json]: audio.postman_collection.json



## API文档

详见

[api.md]: ami.md



