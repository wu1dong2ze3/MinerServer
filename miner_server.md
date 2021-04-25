# 矿机控制服务接口

# 版本1.0.2 
####修定时间 2021-04-22
####修改内容:
#####1.0.2版本：2021-04-22
修复了修改用户名密码也返回token，否则用户校验的时候，用的是旧的token，不符合逻辑
#####1.0.1版本：2021-04-20
在矿机算力图接口中，修改返回参数start_collection_time,end_collection_time
用于前端网页方便计算点数据


# 错误码表

|code|message|
|:-----  |:-------|
|200       |(成功） 服务器已成功处理了请求。|
|401       |(未授权） 请求要求身份验证。|
|421       |用户名密码校验失败|
|424       |用户名或者密码尝试次数超过5次，请五分钟之后再次登陆|
|423       |挖矿地址校验失败，地址格式不合法|
|424       |网络信息校验失败，参数不合法|
|425       |工作模式设置参数不合法|
|426       |升级文件的md5校验失败|
|427       |矿机重启失败|
|428       |恢复出厂失败|
|500       |服务器内部错误|
|501       |服务器无法读取到请求数据|

# 接口基本参数

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-lang  | -否|- string    | -用户标识 默认zh-cn，可选 en-us|

#### 成功

```json
{
  "code": 200,
  //错误码
  "message": "成功",
  //描述
  "data": {
    ......
  }
}
``` 

#### 失败

```json
{
  "code": 500,
  //错误码
  "message": "服务器内部错误"
  //描述
}
``` 

# 一 用户

### 1.1 登陆

###### 接口功能

> 登陆

###### URL

> http://localhost:8080/user/login

###### 支持格式

> JSON

###### HTTP请求方式

> POST

###### header参数

无

###### body参数

|参数|必选|类型|说明| 
|:-----  |:-------|:-----|:----- | 
|-user       | -是|- string    | -用户名,大小写+数字 最小5最大12位|
|-pwd   | -是|- string    | -密码 大小写+数字 最小5最大12位|
|-device   | -否|- int    | -设备类型 0=浏览器 1=android 默认为0|

###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- token      | -string |- 登陆状态标识,通过用户名，密码，当前时间，设备类型算出的唯一字符串|

###### 接口示例

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiJ9.eyJ1c2VyTm8iOiIyMDIwMDMyMzA5MjcwNDc0IiwiaXNzIjoiUk9OQ09PIiwidHlwZSI6Img1IiwiZXhwIjoxNjAxNzExMjM0fQ.rAyn1iOHsed5td6HrlSkeSfPFkTqmr8hkUdMuvQQTOQ"
    //token
  }
}
``` 

### 1.2 修改密码

###### 接口功能

> 修改用户名密码

###### URL

> http://localhost:8080/user/update

###### 支持格式

> JSON

###### HTTP请求方式

> POST

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

|参数|必选|类型|说明| 
|:-----  |:-------|:-----|:----- | 
|-user       | -是|- string    | -用户名 密码 大小写+数字 最小5最大12位|
|-pwd   | -是|- string    | -密码 密码 大小写+数字 最小5最大12位|
|-old_pwd   | -是|- string    | -旧密码 密码 大小写+数字 最小5最大12位|

###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- token      | -string |-因为用户名密码修改了，需要用新的 token 访问|

###### 接口示例

```json
{
  "code": 421,
  "message": "用户名密码校验失败"
}
```
成功
```json
{
  "code": 200,
  "message": "",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiJ9.eyJ1c2VyTm8iOiIyMDIwMDMyMzA5MjcwNDc0IiwiaXNzIjoiUk9OQ09PIiwidHlwZSI6Img1IiwiZXhwIjoxNjAxNzExMjM0fQ.rAyn1iOHsed5td6HrlSkeSfPFkTqmr8hkUdMuvQQTOQ"
    //token
  }
}
```

# 二 矿机信息

### 2.1 矿机运行状态信息

###### 接口功能

> 获取当前矿机运行状态，对应用户面板，左4显示区

###### URL

> http://localhost:8080/miner/summary

###### 支持格式

> JSON

###### HTTP请求方式

> GET

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

无

###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- hashrate      | -string |- 实时算力|
|- duration      | -int |- 运行时长 |
|- reject_rate      | -float |- 拒绝率|
|- temp      | -string |- 温度

###### 接口示例

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "summary": {
      "hashrate": "473.45 MH/s", //实时算力，带单位
      "duration": 1341242434, //运行时常
      "reject_rate": 0.12, //拒绝率 百分比
      "temp": "58C"//温度，暂时统一成摄氏度
    }
  }
}
``` 

### 2.2 矿机算力图

###### 接口功能

> 获取矿机算力图接口

###### URL

> http://localhost:8080/miner/hashrate

###### 支持格式

> JSON

###### HTTP请求方式

> POST

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

|参数|必选|类型|说明| 
|:-----  |:-------|:-----|:----- | 
|-last_collection_time       | -是|- int    | -最后一次采集时间。注意，此参数的时间，特指服务端的时钟。一般情况下，需要输入两种类型值，如果为0，则返回当前服务器时间的12小时范围内，所有点的信息。如果不为0，那么服务器会取这次时间对比当前服务器的时钟差值范围内的所有点的信息。|
|-type       | -是|- int    | -采集类型 0=每15分钟为单位统计平均算力，1=每1小时为单位统计算力|


###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- index      | -int |- 算力板对应的index,如果index=-1,则是总算力|
|- start_collection_time      | -int |- 开始采集时间|
|- end_collection_time      | -int |- 结束采集时间|
|- type      | -int |- 采集类型|
|- count      | -int |-采集点数量|
|- points[]      | -array<flaot> |- 点数组|

###### 接口示例

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "hashrate": {
      "start_collection_time": 123123432,
      "end_collection_time": 126178431,
      "type": 0,
      "count": 5,
      "points": [
        {
          "index:": 0,
          "points": [
            12.00,
            13.00,
            0,
            0,
            1.0
          ]
        },
        {
          "index:": 1,
          "points": [
            12.00,
            13.00,
            0,
            0,
            1.0
          ]
        },
        {
          "index:": 2,
          "points": [
            12.00,
            13.00,
            0,
            0,
            1.0
          ]
        },
        {
          "index:": 3,
          "points": [
            12.00,
            13.00,
            0,
            0,
            1.0
          ]
        },
        {
          "index:": -1,
          "points": [
            12.00,
            13.00,
            0,
            0,
            1.0
          ]
        }
      ]
    }
  }
}
```
说明：</br>
此接口所有时间节点（start_collection_time,end_collection_time，last_collection_time）都以服务端为主，除测试情况外，不要把客户端时间直接填入，避免时钟不同步问题。</br>
1 首先参数last_collection_time=0 进行第一次采集。<br>
2 返回的点中将包含这次采集的服务器的开始采集时间和结束采集时间例如：start_collection_time=123123432 end_collection_time=126178431<br>
3 如果需要动态图，只需要把下次采集累加值，仅需要把这个end_collection_time，传入last_collection_time，服务器就会自动计算出两个时间间隔范围内的点的信息。<br>
4 所以，需要动图的客户端可靠的做法是开个定时器，每间隔一段时间（一般是大于一个（end_collection_time-start_collection_time）/count）秒，采集一次信息。循环即可。<br>


### 2.3 矿池状态

###### 接口功能

> 获取当前矿池的基本运行状态

###### URL

> http://localhost:8080/miner/pool/info

###### 支持格式

> JSON

###### HTTP请求方式

> GET

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

无

###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- index      | -int |- 序号|
|- address      | -string |- 挖矿地址 |
|- status      | -int |- 在线状态 （0 在线 ，1 离线）|
|- name      | -string |- 矿工名|
|- diff      | -string |- 难度|
|- receive_num      | -int |- 接收数|
|- reject_num      | -int |- 拒绝数|
|- ls_time      | int |- 上次登陆时间|

###### 接口示例

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "pools": [
      {
        "index": 0,
        "address": "stratum+tcp://btc.ss.poolin.com:443",
        "status": 1,
        "name": "zhiyuan.5x36",
        "diff": "262K",
        "receive_num": 2344,
        "reject_num": 1,
        "ls_time": 1663445224
      },
      {
        "index": 1,
        "address": "stratum+tcp://btc.ss.poolin.com:443",
        "status": 1,
        "name": "zhiyuan.5x37",
        "diff": "262K",
        "receive_num": 2344,
        "reject_num": 1,
        "ls_time": 1663445224
      },
      {
        "index": 2,
        "address": "stratum+tcp://btc.ss.poolin.com:443",
        "status": 1,
        "name": "zhiyuan.5x38",
        "diff": "262K",
        "receive_num": 2344,
        "reject_num": 1,
        "ls_time": 1663445224
      }
    ]
  }
}
```

### 2.4 算力板状态

###### 接口功能

> 获取当前矿机算力板信息

###### URL

> http://localhost:8080/miner/board/info

###### 支持格式

> JSON

###### HTTP请求方式

> GET

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

无

###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- index      | -int |- 序号|
|- hashrate_real_time      | -string |- 实时算力 |
|- hashrate_in_theory      | -string |- 理论算力|
|- hardware_errs      | -int |- 硬件错误|
|- receive_num      | -int |- 接收数|
|- reject_num      | -int |- 拒绝数|
|- temp      | string |- 温度|
|- status      | int |- 状态，0正常 1异常 2故障|

###### 接口示例

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "boards": [
      {
        "index": 0,
        "hashrate_real_time": "26.39 TH/s",
        "hashrate_in_theory": "27.5 TH/s",
        "hardware_errs": 425,
        "receive_num": 1722,
        "reject_num": 12,
        "temp": "65.1C",
        "status": 0
      },
      {
        "index": 1,
        "hashrate_real_time": "26.39 TH/s",
        "hashrate_in_theory": "27.5 TH/s",
        "hardware_errs": 425,
        "receive_num": 1722,
        "reject_num": 12,
        "temp": "65.1C",
        "status": 0
      },
      {
        "index": 2,
        "hashrate_real_time": "26.39 TH/s",
        "hashrate_in_theory": "27.5 TH/s",
        "hardware_errs": 425,
        "receive_num": 1722,
        "reject_num": 12,
        "temp": "65.1C",
        "status": 0
      }
    ]
  }
}
```

### 2.5 风扇状态

###### 接口功能

> 获取当前矿机风扇状态

###### URL

> http://localhost:8080/miner/fan/info

###### 支持格式

> JSON

###### HTTP请求方式

> GET

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

无

###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- index      | -int |- 序号|
|- status      | -int |- 风扇状态 0 启用 1 未启用（水冷默认此类型） |
|- fan_rpm      | -int |- 风扇的每分钟转数|

###### 接口示例

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "fans": [
      {
        "index": 0,
        "status": 0,
        "fan_rpm": 7233
      },
      {
        "index": 1,
        "status": 0,
        "fan_rpm": 6233
      },
      {
        "index": 2,
        "status": 1,
        "fan_rpm": 0
      },
      {
        "index": 3,
        "status": 0,
        "fan_rpm": 5333
      }
    ]
  }
}
```


# 三 设置和维护

### 3.1 矿工信息获取

###### 接口功能

> 获取当前矿机矿工的账户的基本信息

###### URL

> http://localhost:8080/miner/user/info

###### 支持格式

> JSON

###### HTTP请求方式

> GET

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

无

###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- index      | -int |- 账户序号|
|- address      | -string |- 挖矿地址 |
|- name      | -string |- 矿工名|
|- pwd      | -string |- 密码 选填|

###### 接口示例

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "miners": [
      {
        "index": 0,
        "address": "stratum+tcp://btc.ss.poolin.com:443",
        "name": "zhiyuan.5x36",
        "pwd": ""
      },
      {
        "index": 1,
        "address": "stratum+tcp://btc.ss.poolin.com:443",
        "name": "zhiyuan.5x37",
        "pwd": ""
      },
      {
        "index": 2,
        "address": "stratum+tcp://btc.ss.poolin.com:443",
        "name": "zhiyuan.5x38",
        "pwd": ""
      }
    ]
  }
}
``` 

### 3.2 修改矿工信息

###### 接口功能

> 修改当前矿机矿工的账户的基本信息

###### URL

> http://localhost:8080/miner/user/update

###### 支持格式

> JSON

###### HTTP请求方式

> POST

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|miner    |  是  |  array   | 矿工地址数组 |
|miner[0][address]    |  是  |  string   | 地址 |
|miner[0][name]    |  是  |  string   | 矿工名 |
|miner[0][pwd]    |  是  |  string   | 密码 |

###### 返回字段

无

###### 接口示例

```json
{
  "code": 200,
  "message": "成功"
}
``` 

### 3.3 获取网络信息

###### 接口功能

> 获取当前矿机的网络信息

###### URL

> http://localhost:8080/system/net/info

###### 支持格式

> JSON

###### HTTP请求方式

> GET

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

无

###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- mac      | -string |- mac地址|
|- router_type      | -int |- 路由方式 0=static 1=dhcp |
|- ip      | -string |- ip地址|
|- subnet_mask      | -string |- 子网掩码|
|- gateway      | -string |- 网关|
|- dns1      | -string |- dns1|
|- dns2      | -string |- dns2|

###### 接口示例

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "net_info": {
      "mac": "38-AD-45-7E-6D-3E",
      "router_type": 1,
      "ip": "10.0.0.2",
      "subnet_mask": "255.255.255.0",
      "gateway": "10.0.0.1",
      "dns1": "114.114.114.114",
      "dns2": "202.106.0.20"
    }
  }
}
```

### 3.4 修改网络信息

###### 接口功能

> 修改当前矿机网络的基本信息

###### URL

> http://localhost:8080/system/net/update

###### 支持格式

> JSON

###### HTTP请求方式

> POST

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-router_type    |- 是  | - int   |- 路由方式 0=static 1=dhcp </br> 如果是dhcp模式，设置地址 网关等 </br> 等信息不再生效 |
|-ip    | - 否  | - string   |- 地址 |
|-subnet_mask   | - 否  |- string   |- 子网掩码 |
|-gateway    | - 否  | - string   |- 网关 |
|-dns1    | - 否  |- string   |- dns1 |
|-dns2    | - 否  |- string   |- dns2 |

###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- mac      | -string |- mac地址|
|- router_type      | -int |- 路由方式 0=static 1=dhcp |
|- ip      | -string |- ip地址|
|- subnet_mask      | -string |- 子网掩码|
|- gateway      | -string |- 网关|
|- dns1      | -string |- dns1|
|- dns2      | -string |- dns2|

###### 接口示例

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "net_info": {
      "mac": "38-AD-45-7E-6D-3E",
      "router_type": 1,
      "ip": "10.0.0.2",
      "subnet_mask": "255.255.255.0",
      "gateway": "10.0.0.1",
      "dns1": "114.114.114.114",
      "dns2": "202.106.0.20"
    }
  }
}
```

### 3.5 工作模式获取

###### 接口功能

> 获取当前矿机工作模式

###### URL

> http://localhost:8080/miner/mode/info

###### 支持格式

> JSON

###### HTTP请求方式

> GET

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

无

###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- frequency      | -int |- 工作频率|
|- voltage      | -float |- 电压 |

###### 接口示例

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "mode": {
      "frequency": 800,
      "voltage": 12.34
    }
  }
}
```

### 3.6 修改工作模式

###### 接口功能

> 修改当前矿机的工作模式参数

###### URL

> http://localhost:8080/miner/mode/update

###### 支持格式

> JSON

###### HTTP请求方式

> POST

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-frequency    |- 是  | - int   |- 频率|
|-voltage    | - 否  | - float   |- 电压 |

###### 返回字段

无

###### 接口示例

```json
{
  "code": 200,
  "message": "成功"
}
```

### 3.7 获取固件信息

###### 接口功能

> 获取固件信息，用于固件升级页面

###### URL

> http://localhost:8080/system/ota/info

###### 支持格式

> JSON

###### HTTP请求方式

> GET

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

无

###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- version      | -string |- 版本号|
|- time      | -string |- 描述 |
|- ota_name      | -string |- 上次升级文件名称 |

###### 接口示例

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "ota": {
      "version": "ota-aa-bb-cc-1",
      "time": "2021.4.5 22:22",
      "ota_name": "/user/file/ota_01_01.rar"
    }
  }
}
```

### 3.8 上传固件

###### 接口功能

> 固件升级接口

###### URL

> http://localhost:8080/system/ota/upgrade

###### 支持格式

> JSON

###### HTTP请求方式

> POST

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|
|-multipart  | -是|- string    |=form-data  -上传文件数据类型|

###### body参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-filename    |- 是  | - string   |- 上传文件标识 </br> 举例子：filename=@C:\Users\user\ota_2021_01_01.rar|
|-md5   | - 否  | - string   |- md5校验码 |

###### 返回字段

无

###### 接口示例

```json
{
  "code": 200,
  "message": "成功"
}
```

### 3.9 重启

###### 接口功能

> 重启矿机

###### URL

> http://localhost:8080/system/reboot

###### 支持格式

> JSON

###### HTTP请求方式

> GET

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

无

###### 返回字段

无

###### 接口示例

```json
{
  "code": 200,
  "message": "成功"
}
```

### 3.10 恢复出厂

###### 接口功能

> 恢复出厂

###### URL

> http://localhost:8080/system/reset

###### 支持格式

> JSON

###### HTTP请求方式

> GET

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

无

###### 返回字段

无

###### 接口示例

```json
{
  "code": 200,
  "message": "成功"
}
```

### 3.11 获取日志

###### 接口功能

> 获取固件信息，用于固件升级页面

###### URL

> http://localhost:8080/system/log

###### 支持格式

> JSON

###### HTTP请求方式

> POST

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-start_line  | -是|- int    | -从多少行开始读|

###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- start_line      | -int |- 启始行 |
|- lines      | -int |- 此次日志本体中包含多少行 |
|- log      | -string |-日志本体  |

###### 接口示例

post start_line=0

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "log": {
      //共100行
      "start_line": 0,
      //从0行开始读取
      "lines": 2,
      //2
      "log": "Mem: 170420K used, 63324K free, 0K shrd, 643184K buff, 643228K cached\n CPU:   0% usr  20% sys   0% nic  80% idle   0% io   0% irq   0% sirq"
    }
  }
}
```

post start_line=2

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "log": {
      "start_line": 2,
      //从2行开始读取 (包含2)
      "lines": 0,
      //没有文件返回0
      "log": ""
      //没有返回""
    }
  }
}
```

注意： 客户端可以用这样的逻辑读日志，</br>
1 首先日志以行为单位读取和发送，每次发送的行数由服务端指定，不定长。</br> 2 第一此展开页面，先发送 start_line=0的请求，服务端如果没有获取到日志（start_line=0
lines=0），则间隔几秒继续刷日志start_line=0的请求</br>
如果客户端获取到日志（lines>0），进行步骤3 </br>
3 先把获取到的日志显示再浏览器上，间隔几秒可以继续发送请求 start_line=start_line+lines, </br>
4 如果返回的lines大于0，证明有新的日志，否则暂无新日志，只要客户不关闭浏览器，重复3步骤。</br>
5 如果服务端15秒内没有接受到任何日志读取请求，服务端自动关闭日志管道。并释放当前日志文件fd，此时，意味着当前系统日志，正常滚动刷新，用户会丢失被刷新掉的部分。</br>
6 在一次页面交互中，不保证每一次start_line=0获取到的初始日志是相同的，因为每次调用日志，日志文件都可能在滚动刷新，日志文件只记录第一次打开日志文件的快照，以及滚动刷新的部分。</br>

# 四 系统状态

### 4.1 页面标题兰信息

###### 接口功能

> 获取标题栏页面的信息

###### URL

> http://localhost:8080/system/ui/title_bar

###### 支持格式

> JSON

###### HTTP请求方式

> GET

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

无

###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- ip      | -string |- ip|
|- mac      | -string |- 字符串 |
|- version      | -string  |- 固件版本|

###### 接口示例

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "title_bar": {
      "ip": "192.168.1.1",
      "mac": "C4:E5:E7:3F:E0:F4",
      "version": "a10s_20210108_053449"
    }
  }
}
``` 

### 4.2 硬件版本

###### 接口功能

> 获取硬件版本相关信息

###### URL

> http://localhost:8080/system/hardware/version

###### 支持格式

> JSON

###### HTTP请求方式

> GET

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

无

###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- model      | -string |- 型号|
|- cpu      | -string |- cpu |
|- firmware_version      | -string  |- 固件版本|
|- firmware_date      | -string  |- 固件发布时间|

###### 接口示例

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "version": {
      "model": "A10s",
      "cpu": "g19",
      "firmware_version": "a10s_20210108_053449",
      "firmware_date": "2020-11-21 15:21"
    }
  }
}
``` 

### 4.3 获取网络状态

参考3.3接口文档

### 4.4 硬件状态

###### 接口功能

> 获取硬件版本相关信息

###### URL

> http://localhost:8080/system/hardware/status

###### 支持格式

> JSON

###### HTTP请求方式

> GET

###### header参数

|参数|必选|类型|说明|
|:-----  |:-------|:-----|:----- | 
|-token  | -是|- string    | -用户标识|

###### body参数

无

###### 返回字段

|返回字段|字段类型|说明 |
|:-----  |:-------|:-----|
|- time      | -string |- 运行时间|
|- cpu      | -string |- cpu |
|- ram      | -string  |- 固件版本|
|- rom      | -string  |- 固件发布时间|

###### 接口示例

```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "status": {
      "time": "2021 03-29 09:39:51",
      //开机时间 
      "cpu": "50",
      //cpu占用 （百分比）
      "ram": "63108/240384",
      //内存占用
      "rom": "34368/223144"
      //存储占用
    }
  }
}
``` 

