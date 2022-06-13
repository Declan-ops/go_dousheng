# 汇报文档



## 使用技术栈

- MySql
- gorm
- Gin
- OSS



## 成果展示

#### 视频地址

https://jxau7124.oss-cn-shenzhen.aliyuncs.com/%E5%B1%95%E7%A4%BA%E8%A7%86%E9%A2%91/202206131040.mp4


## 项目接口说明

### 视频模块

#### 视频流接口

##### 请求方式：GET

##### 请求路径：/douyin/feed

不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个

##### controller层

- 获取当前时间戳，获取参数
- 封装好后传给Service层
- 接收Service层的响应结果
- 在判断没有错误后返回浏览器Json数据

##### Service层

- 进行Token鉴权
- 把时间戳交给Dao层
- 接收Dao层封装的数据返回给Controller层

##### Dao层

- 根据时间戳去数据库中查发布时间小于当前时间戳的视频
- 封装好数据，检查错误后返回给Service层

##### 问题

数据库查询次数较多，响应处理结果较慢





#### 投稿接口

##### 请求方式：Post

##### 请求路径：/douyin/publish/action/

登录用户选择视频上传

##### controller层

- 接收客户端传过来的token与title，以及视频数据
- 封装到对应的结构体中传递给Service层
- 接收响应结果，返回对应的json数据



##### Service层

- 进行Token鉴权，读取用户信息
- 取出视频数据上传到阿里云OSS对象存储中
- 截至视频缩略图作为封面
- 把对象存储的地址、封面存储地址、用户数据封装传递给Dao层，保存到数据库中
- 返回响应结果到Controller层



##### Dao层

- 把封装好的数据插入到数据库中
- 检查错误后返回给Service层



##### 问题

没有做到异步上传视频，响应结果较慢



##### 亮点

- 使用Gin框架封装的api对编码方式为` multipart/form-data`的表单读取视频流数据

- 使用阿里云OSS对象存储服务存储视频

- 截取视频的缩略图作为视频的封面

  

#### 发布列表接口

##### 请求方式：Get

##### 请求路径：/douyin/publish/list/

用户的视频发布列表，直接列出用户所有投稿过的视频

##### controller层

- 接收客户端传过来的token与user_id
- 封装到对应的结构体中传递给Service层
- 接收响应结果，返回对应的json数据

##### Service层

- 进行Token鉴权，读取用户信息
- 把用户数据封装传递给Dao层，保存到数据库中
- 返回响应结果到Controller层

##### Dao层

- 根据用户id查询用户发布的所有视频
- 封装到对应的实体模型中
- 检查错误后返回给Service层

##### 问题

- 数据库查询次数较多，响应处理结果较慢





### 用户模块

#### 用户登录接口

##### 请求方式：Post

##### 请求路径：/douyin/user/login/

通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token

##### controller层

- 接收客户端传过来的username与password
- 封装到对应的结构体中传递给Service层
- 接收响应结果，返回对应的json数据



##### Service层

- 根据用户姓名从本地缓存map中查询用户，不存在则返回错误
- 将客户端传递过来的用户密码加密
- 比较缓存中用户与客户端传递过来的用户信息，不一致返回错误
- 检查token是否存在，是否过期，如果是则更新并把数据传递给Dao层，持久化到数据库中
- 更新本地map缓存



##### Dao层

- 把封装好的数据更新到数据库中
- 检查错误后返回给Service层



##### 亮点

- 使用缓存检查用户信息,t提升响应速度，减少数据库压力

- 使用MD5加密算法，对用户密码进行加密

  





#### 用户注册接口

##### 请求方式：Post

##### 请求路径：/douyin/user/register/

新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token

##### controller层

- 接收客户端传过来的username与password
- 封装到对应的结构体中传递给Service层
- 接收响应结果，返回对应的json数据



##### Service层

- 对密码进行MD5加密
- 把加密后的密码放入到用户数据中
- 根据用户信息下发Token
- 把用户数据封装传递给Dao层，保存到数据库中
- 返回响应结果到Controller层



##### Dao层

- 连接数据库

- 把封装好的数据插入到数据库中

- 把用户信息放入本地map缓存中

- 检查错误后返回给Service层

  

#### 用户信息接口

##### 请求方式：Get

##### 请求路径：/douyin/user/

获取用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数

##### controller层

- 接收客户端传过来的user_id与token
- 封装到对应的结构体中传递给Service层
- 接收响应结果，返回对应的json数据



##### Service层

- 对Token进行鉴权，读取用户信息
- 把用户数据封装传递给Dao层，保存到数据库中
- 返回响应结果到Controller层



##### Dao层

- 连接数据库

- 根据条件查找用户的粉丝数以及关注数
- 封装数据
- 检查错误后返回给Service层



##### 问题

多次查询数据库，响应结果较慢





## 技术说明

#### Token

##### 下发Token

根据用户信息下发token

> 借用JWT框架，把token信息分成三层。把用户数据json话，给定过期时间，保存到Claim中，调用api得到token字符串



##### Token鉴权

- 判断Token是否错误、Token是否过期
- 从Token中取出用户信息，并将json信息Unmarshal到用户结构体中
- 根据用户名字到数据库中取出对应的数据进行比较
- 信息一致返回用户信息



#### OSS对象存储

##### 总体流程

接收文件后，创建一个创建OSSClient实例，填写好相应的个人信息后，指定存储空间，就可以把本地的文件上传了。

##### 接收文件

这里需要前端上传文件的表单的`enctype`属性指定值为`multipart/form-data`

最早的HTTP是不支持文件上传的，后面为了开发，Content-Type的类型扩充了multipart/form-data用以支持向服务器发送二进制数据。因此发送post请求时候，form表单属性enctype共有二个值可选，这个属性管理的是表单的MIME编码：

①application/x-www-form-urlencoded(默认值) ②multipart/form-data 这里我们选择第二个编码方式。

**代码：**

```go
func UpLoadFile(c *gin.Context) {

    f, err := c.FormFile("data") // multipart/form-data
}
```

##### 创建OSSClient实例

```go
 	// 创建OSSClient实例。
	//  yourEndpoint填写Bucket对应的Endpoint
    AccessKeyID := "xxxxx"           //填写你的AccessKeyID
    AccessKeySecret := "xxxxx" 		//AccessKeySecret
    client, err := oss.New(Endpoint, AccessKeyID, AccessKeySecret)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }

```

##### 填写存储空间名称

```go
// 指定bucket
    bucket, err := client.Bucket("jxau7124") // 根据自己的填写
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)

    }
```

##### 指定路径

```go
    // 将文件流上传至指定目录下
    path := video.user.Name + "/" + video.title + file.Filename

    err = bucket.PutObject(path, src)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)

    }
```

然后就可以进行访问了

访问路径：

`https://桶名+oss-cn-服务器所在城市名+.aliyuncs.com/+存储路径`

###### 缩略图

一般我们上传完图片或者视频后想要得到一张缩略图，这一块oss也帮我们封装好了。

在图片的地址后面加上以下代码，可以生成缩略图

```go
?x-oss-process=image/resize,m_fill,w_200,quality,q_60
```

在视频的地址后面加上以下代码，可以获取视频的封面

```go
?x-oss-process=video/snapshot,t_7000,f_jpg,w_800,h_600,m_fast
```



#### MD5
```go
  // 进行md5加密
	newSig := md5.Sum([]byte(password)) //转成加密编码
	// 将编码转换为字符串
	newArr := fmt.Sprintf("%x", newSig)
	//输出字符串字母都是小写，转换为大写
	password = strings.ToTitle(newArr)
 ```




