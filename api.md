# 服务接口文档

## 添加记录

**接口地址** `/api/audio/add`

**请求方式** `POST`

**consumes** `["application/x-www-form-urlencoded"]`

**produces** `["*/*"]`

**接口描述** 

**请求参数**

| 参数名称            | 参数说明       | 请求类型 | 是否必须 | 数据类型 |
| ------------------- | -------------- | -------- | -------- | -------- |
| TelephoneKind       | 手机类型       | body     | true     | string   |
| FromTelephoneNumber | 拨打手机号码   | body     | true     | string   |
| ToTelephoneNumber   | 接听手机号码   | body     | true     | string   |
| HappenTimestamp     | 录音时间戳(秒) | body     | true     | int      |
| TotalDuration       | 录音时长(秒)   | body     | true     | int      |
| File                | 录音文件       | body     | true     | file     |



​                                  

**响应状态**

| 状态码 | 说明         | schema        |
| ------ | ------------ | ------------- |
| 200    | OK           | 内容«boolean» |
| 201    | Created      |               |
| 401    | Unauthorized |               |
| 403    | Forbidden    |               |
| 404    | Not Found    |               |




**响应参数**

| 参数名称  | 参数说明                                | 类型   | schema |
| --------- | --------------------------------------- | ------ | ------ |
| error_no  | 结果码，0为正常                         | string |        |
| data      | 返回的内容, id                          | int    |        |
| error_msg | 返回的消息,仅做参考，服务端统一返回英文 | string |        |



**请求示例**

```json
"body": {
					"mode": "formdata",
					"formdata": [
						{
							"key": "TelephoneKind",
							"value": "vivo",
							"type": "text"
						},
						{
							"key": "FromTelephoneNumber",
							"value": "12345678901",
							"type": "text"
						},
						{
							"key": "ToTelephoneNumber",
							"value": "23456789012",
							"type": "text"
						},
						{
							"key": "HappenTimestamp",
							"value": "1617029034",
							"type": "text"
						},
						{
							"key": "TotalDuration",
							"value": "300",
							"type": "text"
						},
						{
							"key": "File",
							"type": "file",
							"src": ""
						}
					]
			}
}

```

**响应示例**

```json
{
    "error_no": 0,
    "error_msg": "",
    "data": 1
}
```

## 

## 删除记录

**接口地址** `/api/audio/delete`

**请求方式** `POST`

**consumes** `["application/json"]`

**produces** `["*/*"]`

**接口描述** 

**请求参数**

| 参数名称 | 参数说明   | 请求类型 | 是否必须 | 数据类型 |
| -------- | ---------- | -------- | -------- | -------- |
| id       | 数据条目ID | body     | true     | int      |



​                                  

**响应状态**

| 状态码 | 说明         | schema        |
| ------ | ------------ | ------------- |
| 200    | OK           | 内容«boolean» |
| 201    | Created      |               |
| 401    | Unauthorized |               |
| 403    | Forbidden    |               |
| 404    | Not Found    |               |




**响应参数**

| 参数名称  | 参数说明                                | 类型   | schema |
| --------- | --------------------------------------- | ------ | ------ |
| error_no  | 结果码，0为正常                         | string |        |
| data      | 返回的是否执行成功                      | bool   |        |
| error_msg | 返回的消息,仅做参考，服务端统一返回英文 | string |        |



**请求示例**

```json
{
	"id":1
}
```

**响应示例**

```json
{
    "error_no": 0,
    "error_msg": "",
    "data": true
}
```

## 

## 

## 修改记录信息

**接口地址** `/api/audio/update`

**请求方式** `POST`

**consumes** `["application/json"]`

**produces** `["*/*"]`

**接口描述** 

**请求参数**

| 参数名称            | 参数说明       | 请求类型 | 是否必须 | 数据类型 |
| ------------------- | -------------- | -------- | -------- | -------- |
| id                  | 数据条目ID     | body     | true     | int      |
| telephoneKind       | 手机类型       | body     | false    | string   |
| fromTelephoneNumber | 拨打手机号码   | body     | false    | string   |
| toTelephoneNumber   | 接听手机号码   | body     | false    | string   |
| happenTimestamp     | 录音时间戳(秒) | body     | false    | int      |
| totalDuration       | 录音时长(秒)   | body     | false    | int      |



​                                  

**响应状态**

| 状态码 | 说明         | schema        |
| ------ | ------------ | ------------- |
| 200    | OK           | 内容«boolean» |
| 201    | Created      |               |
| 401    | Unauthorized |               |
| 403    | Forbidden    |               |
| 404    | Not Found    |               |




**响应参数**

| 参数名称  | 参数说明                                | 类型   | schema |
| --------- | --------------------------------------- | ------ | ------ |
| error_no  | 结果码，0为正常                         | string |        |
| data      | 返回的是否执行成功                      | bool   |        |
| error_msg | 返回的消息,仅做参考，服务端统一返回英文 | string |        |



**请求示例**

```json
{
	"id":2,
	"telephoneKind":"iphone",
	"fromTelephoneNumber":"abc",
	"toTelephoneNumber":"efg",
	"happenTimestamp":123,
	"totalDuration":456
}
```

**响应示例**

```json
{
    "error_no": 0,
    "error_msg": "",
    "data": true
}
```

## 

## 

## 查询记录信息

**接口地址** `/api/audio/detail`

**请求方式** `POST`

**consumes** `["application/json"]`

**produces** `["*/*"]`

**接口描述** 

**请求参数**

| 参数名称 | 参数说明   | 请求类型 | 是否必须 | 数据类型 |
| -------- | ---------- | -------- | -------- | -------- |
| id       | 数据条目ID | body     | true     | int      |



​                                  

**响应状态**

| 状态码 | 说明         | schema        |
| ------ | ------------ | ------------- |
| 200    | OK           | 内容«boolean» |
| 201    | Created      |               |
| 401    | Unauthorized |               |
| 403    | Forbidden    |               |
| 404    | Not Found    |               |




**响应参数**

| 参数名称  | 参数说明                                | 类型   | schema       |
| --------- | --------------------------------------- | ------ | ------------ |
| error_no  | 结果码，0为正常                         | string |              |
| data      | 记录信息                                | object | 文件记录信息 |
| error_msg | 返回的消息,仅做参考，服务端统一返回英文 | string |              |



**schema属性说明**

**文件记录信息**

| 参数名称            | 参数说明           | 类型   | schema |
| ------------------- | ------------------ | ------ | ------ |
| id                  | 数据条目ID         | int    |        |
| telephoneKind       | 手机类型           | string |        |
| fromTelephoneNumber | 拨打手机号码       | string |        |
| toTelephoneNumber   | 接听手机号码       | string |        |
| happenTimestamp     | 录音时间戳(秒)     | string |        |
| totalDuration       | 录音时长(秒)       | int    |        |
| fileName            | 录音文件名         | int    |        |
| size                | 录音文件大小(Byte) | int    |        |
| fileType            | 录音文件类型       | string |        |
| md5                 | 录音文件MD5        | string |        |
| filePath            | 录音文件下载路径   | string |        |



**请求示例**

```json
{
	"id":2
}
```

**响应示例**

```json
{
    "error_no": 0,
    "error_msg": "",
    "data": {
        "id": 2,
        "telephoneKind": "iphone",
        "fromTelephoneNumber": "abc",
        "toTelephoneNumber": "efg",
        "happenTimestamp": 123,
        "totalDuration": 456,
        "fileName": "0.2.9.52-main-last1.tar.gz",
        "size": 13096841,
        "fileType": ".gz",
        "md5": "d41d8cd98f00b204e9800998ecf8427e",
        "filePath": "/12345678901/23456789012/20210329/0.2.9.52-main-last1.tar.gz"
    }
}
```

## 

## 

## 查询记录列表

**接口地址** `/api/audio/list`

**请求方式** `POST`

**consumes** `["application/json"]`

**produces** `["*/*"]`

**接口描述** 

**请求参数**

| 参数名称            | 参数说明                   | 请求类型 | 是否必须 | 数据类型 | schema                    |
| ------------------- | -------------------------- | -------- | -------- | -------- | ------------------------- |
| sortAsc             | 排序类型（0:DESC，1：ASC） | body     | false    | int      |                           |
| sortField           | 排序字段                   | body     | false    | string   |                           |
| pageNumber          | 页码                       | body     | false    | int      |                           |
| pageSize            | 页大小                     | body     | false    | int      |                           |
| telephoneKind       | 手机类型                   | body     | false    | string   | 模糊查询时使用%           |
| fromTelephoneNumber | 拨打手机号码               | body     | false    | string   | 模糊查询时使用%           |
| toTelephoneNumber   | 接听手机号码               | body     | false    | string   | 模糊查询时使用%           |
| happenTimestamp     | 录音时间戳(秒)             | body     | false    | int      | 范围查询时使用min,max对象 |
| totalDuration       | 录音时长(秒)               | body     | false    | int      | 范围查询时使用min,max对象 |
| fileName            | 录音文件名                 | body     | false    | int      | 范围查询时使用min,max对象 |
| size                | 录音文件大小(Byte)         | body     | false    | int      | 范围查询时使用min,max对象 |
| fileType            | 录音文件类型               | body     | false    | string   | 模糊查询时使用%           |
| md5                 | 录音文件MD5                | body     | false    | string   | 模糊查询时使用%           |
| filePath            | 录音文件下载路径           | body     | false    | string   | 模糊查询时使用%           |

​                                  

**响应状态**

| 状态码 | 说明         | schema        |
| ------ | ------------ | ------------- |
| 200    | OK           | 内容«boolean» |
| 201    | Created      |               |
| 401    | Unauthorized |               |
| 403    | Forbidden    |               |
| 404    | Not Found    |               |




**响应参数**

| 参数名称  | 参数说明                                | 类型   | schema       |
| --------- | --------------------------------------- | ------ | ------------ |
| error_no  | 结果码，0为正常                         | string |              |
| list      | 记录信息列表                            | object | 文件记录信息 |
| page      | 分页信息                                | object | 分页信息     |
| error_msg | 返回的消息,仅做参考，服务端统一返回英文 | string |              |



**schema属性说明**



**分页信息**

| 参数名称  | 参数说明            | 类型 | schema |
| --------- | ------------------- | ---- | ------ |
| page      | 当前页码（从1开始） | int  |        |
| size      | 当前页记录数        | int  |        |
| totalPage | 总页数              | int  |        |
| totalSize | 总记录数            | int  |        |

**文件记录信息**

| 参数名称            | 参数说明           | 类型   | schema |
| ------------------- | ------------------ | ------ | ------ |
| id                  | 数据条目ID         | int    |        |
| telephoneKind       | 手机类型           | string |        |
| fromTelephoneNumber | 拨打手机号码       | string |        |
| toTelephoneNumber   | 接听手机号码       | string |        |
| happenTimestamp     | 录音时间戳(秒)     | string |        |
| totalDuration       | 录音时长(秒)       | int    |        |
| fileName            | 录音文件名         | int    |        |
| size                | 录音文件大小(Byte) | int    |        |
| fileType            | 录音文件类型       | string |        |
| md5                 | 录音文件MD5        | string |        |
| filePath            | 录音文件下载路径   | string |        |



**请求示例**

```json
{
	"sortAsc":1,
	"sortField":"fromTelephoneNumber",
	"pageNumber":1,
	"pageSize":50,
	"telephoneKind":"iphone",
	"fromTelephoneNumber":"a%",
	"toTelephoneNumber":"%g",
	"happenTimestamp":{"min":100,"max":200},
	"totalDuration":{"min":100,"max":500},
	"fileName":"%main%",
	"size":{"max":23096841},
	"fileType":".gz",
	"md5":"%99%",
	"filePath":"%last%"
}
```

**响应示例**

```json
{
    "error_no": 0,
    "error_msg": "",
    "page": {
        "page": 1,
        "size": 1,
        "totalPage": 1,
        "totalSize": 1
    },
    "list": [
        {
            "id": 2,
            "telephoneKind": "iphone",
            "fromTelephoneNumber": "abc",
            "toTelephoneNumber": "efg",
            "happenTimestamp": 123,
            "totalDuration": 456,
            "fileName": "0.2.9.52-main-last1.tar.gz",
            "size": 13096841,
            "fileType": ".gz",
            "md5": "d41d8cd98f00b204e9800998ecf8427e",
            "filePath": "/12345678901/23456789012/20210329/0.2.9.52-main-last1.tar.gz"
        }
    ]
}
```

## 