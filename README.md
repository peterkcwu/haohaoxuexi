学习资料




### 登录
执行以下命令按照提示操作:
```shell
jd_seckill login
```

### 自动获取eid,fp

> ⚠依赖谷歌浏览器，请安装谷歌浏览器，获取到的eid和fp请手动填入配置文件

执行以下命令按照提示操作:
```shell
jdTdudfp
```
> ⚠目前实验性阶段，请勿依赖该功能

### 预约
执行以下命令按照提示操作:
```shell
reserve
```

### 抢购
执行以下命令按照提示操作:
```shell
seckill
```

### 退出登录
```shell
 logout
```

### 获取版本号
```shell
jd_seckill version
```

> ⚠ 以上命令并不是每次都需要执行的，都是可选的，具体使用请参考提示。

### Linux下命令行方式显示二维码（以Ubuntu为例）

```bash
$ sudo apt-get install qrencode zbar-tools # 安装二维码解析和生成的工具，用于读取二维码并在命令行输出。
$ zbarimg qr_code.png > qrcode.txt && qrencode -r qrcode.txt -o - -t UTF8 # 解析二维码输出到命令行窗口。
```

## 使用教程

#### 1. 推荐Chrome浏览器
#### 2. 网页扫码登录，或者账号密码登录
#### 3. 填写config.ini配置信息
(1)`eid`和`fp`找个普通商品随便下单,然后抓包就能看到,这两个值可以填固定的
> 随便找一个商品下单，然后进入结算页面，打开浏览器的调试窗口，切换到控制台Tab页，在控制台中输入变量`_JdTdudfp`，即可从输出的Json中获取`eid`和`fp`。  
> 不会的话参考issue https://github.com/ztino/jd_seckill/issues/2

(2)`sku_id`,`default_user_agent`
> `sku_id`已经按照茅台的填好。
> `default_user_agent` 可以用默认的。谷歌浏览器也可以浏览器地址栏中输入about:version 查看`USER_AGENT`替换

(3)配置一下时间
> 现在不强制要求同步最新时间了，程序会自动同步京东时间
>> 但要是电脑时间快慢了好几个小时，最好还是同步一下吧

以上都是必须的.
> tips：
> 在程序开始运行后，会检测本地时间与京东服务器时间，输出的差值为本地时间-京东服务器时间，即-50为本地时间比京东服务器时间慢50ms。
> 本代码的执行的抢购时间以本地电脑/服务器时间为准

(4)修改抢购瓶数
> 可在配置文件中找到seckill_num进行修改

(5)其他配置
> 请自行参考使用

