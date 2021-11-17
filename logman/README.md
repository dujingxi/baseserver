# logman

#### 介绍
自己写的简单日志打印模块，可以在同一项目中定义多个 logger，分别将不同业务的Log记录在不同文件中。
可以设置以时间或是大小迭代保存

#### 软件架构
logman/*.go


#### 安装教程

做为模块移到项目中即可

#### 使用说明

1.  import logman

2. testlog := logman.NewLogMan("test.log")

3. testlog.SetSaveMode(logman.ByDay)

4. testlog.SetSaveVal(10)

5.  修改 joinFields 中 content 变量, 根据需要定义每条日志的字段、字段顺序、字段为空时以什么符号代替(like nginx log)。
    - Example: 时间 类型 信息
    ```
      var content = []map[string]string{
        map[string]string{"key": "time", "val": "-"},
	    map[string]string{"key": "type", "val": "-"},
	    map[string]string{"key": "message", "val": "-"},
    }
    ```

6.  日志打印
  testlog.Print(Fields{
    "level":  "INFO",
    "message": "This is a test log message.",
  })
  or .Infof()/.Debugf()/.Warnf()/.Errorf()/.Fatalf()
  testlog.Infof(Fields{
    "message": "This is another test log message.",
  })

* 默认mode = ByDay, 时长 10天

更多内容请直接看代码并根据自己的需要修改
  

