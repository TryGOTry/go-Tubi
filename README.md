# Go-Tubi
Golang写的土司自动签到程序(Server酱推送,检测是否成功,没成功一小时后再执行)

## 说明
又是造轮子的一天.

根据大佬的操作，不用修改标准库也可以带cookie进行post发包了！！！！！！！！   果然还是自己太菜了。


## config.json说明
```
{
  "action":"login",  //不可修改
  "Username": "",    //登录账号
  "password": "",   // 登录密码，不加密
  "questionid": 0,  //安全问题id
  "answer": "",    //安全问题答案
  "serverkey": ""  //server酱的key,不要推送的话，不填写就好了
}
//问题id
//# 0 = 没有安全提问
//# 1 = 母亲的名字
//# 2 = 爷爷的名字
//# 3 = 父亲出生的城市
//# 4 = 您其中一位老师的名字
//# 5 = 您个人计算机的型号
//# 6 = 您最喜欢的餐馆名称
//# 7 = 驾驶执照的最后四位数字

```
---
### 如何编译
```
Linux:
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-s -w" -o Tubi_linux_x64 main.go

windows:
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags "-s -w" -o Tubi_x64.exe main.go
```
### 可以直接下载编译好的文件，修改config.json里的配置就可以了。
太菜了.

![运行截图](https://github.com/TRYblog/go-Tubi/blob/main/111.PNG "go-Tubi")

## 关于作者
一个菜鸟.
[个人博客](https://www.nctry.com)

## 时间
2021/02/22
