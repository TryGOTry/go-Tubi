# Go-Tubi
Golang写的土司自动签到程序(Server酱推送,检测是否成功,没成功一小时后再执行)

## 说明
又是造轮子的一天.

需要修改官方的net/http这个库才行。(ps:改了一下http.postform这个方法，支持带入cookie)
下载client.go这个文件替换官方的net\http\client.go文件
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
### 如果觉得麻烦，可以直接下载编译好的文件，修改config.json里的配置就可以了。
太菜了.

![运行截图](https://github.com/TRYblog/go-Tubi/blob/main/p.jpg "go-Tubi")

## 关于作者
一个菜鸟.
[个人博客](https://www.nctry.com)

## 时间
2021/02/22
