# DDNS
## Installation

```shell
go get -u github.com/izern/ddns-go
cd $GOPATH/src/github/izern/ddns-go
go build ddns.go
```

OR

```shell
go install github.com/izern/ddns-go@main
```

## quick start

save config file on your system,example `~/.ddns/config.yml`

```yml
logging:
  level:
    root: INFO    # default INFO
  encoding: json # json or console, default json. only encoding, console is plan text encoding
  encoder:
    TimeKey: time
    LevelKey: level
    NameKey: logger
    CallerKey: caller
    MessageKey: msg
    StacktraceKey: stacktrace
  output: # default is console
    file:
      path: /var/log/ddns.log
      async: true # async output,default false
yun:
  aliyun:
    accesskey: your_key               # 必填
    accessKeySecret: your_key_secret  # 必填
    endpoint: alidns.cn-hangzhou.aliyuncs.com # 选填，默认是杭州地区
dns:
  - type: IPV6  # IPV4 | IPV6
    rr: blog
    domain: izern.cn
    yun: aliyun   # select yun.aliyun
ip:
  parser: unixIpParser # ip解析器
  ext: # ip解析器 使用参数，使用了哪个ip解析器，对应的参数必填
    unixIpParser: # 解析ip使用unixIpExec时需要指定网卡名
      device: wlan0 # 网卡名，默认eth0
    osExecParser: # 执行命令
      ipv4Cmd: "ip -o -4 addr show wlan0 scope global | awk '{print $4}' | cut -d/ -f1"  # shell命令，或者Windows cmd命令
      ipv6Cmd: "ip -o -6 addr show wlan0 scope global | awk '{print $4}' | cut -d/ -f1"  # shell命令，或者Windows cmd命令
    public: # 公网解析
      ipv4: [ ]   # 解析ipv4的网址,地址请勿随意添加，除非你确定接口返回格式是兼容的
      ipv6: [ ]  # 解析ipv6的网址,地址请勿随意添加，除非你确定接口返回格式是兼容的
```

and then run `ddns-go --config ~/.ddns/config.yml`
