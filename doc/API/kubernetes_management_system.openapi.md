---
title: cloudManagementSystem
language_tabs:
- shell: Shell
- http: HTTP
- javascript: JavaScript
- ruby: Ruby
- python: Python
- php: PHP
- java: Java
- go: Go
  toc_footers: []
  includes: []
  search: true
  code_clipboard: true
  highlight_theme: darkula
  headingLevel: 2
  generator: "@tarslib/widdershins v4.0.23"

---

# cloudManagementSystem

Base URLs:

# Authentication

# Default

## GET ping接口

GET /api/v1/platform/ping

> 返回示例

> 成功

```json
"pong"
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取client的IP地址和端口号

GET /api/v1/platform/address

> 返回示例

> 成功

```json
"IP:port //192.168.31.87:56203"
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 登录接口

POST /api/v1/user/login

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{
  "errCode": 0,
  "data": {
    "token": "string",
    "username": "string"
  },
  "msg": "string",
  "errMsg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» errCode|integer|true|none||none|
|» data|object|true|none||none|
|»» token|string|true|none||none|
|»» username|string|true|none||none|
|» msg|string|true|none||none|
|» errMsg|string|true|none||none|

## GET 获取用户信息接口

GET /api/v1/user/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|query|string| 否 |none|
|token|header|string| 否 |none|

> 返回示例

> 成功

```json
{
  "data": {
    "user": {
      "id": 3,
      "CreatedAt": "2023-02-24 15:40:40",
      "UpdatedAt": "2023-02-24 15:40:40",
      "uid": "123",
      "userName": "gaochuang22",
      "password": "$2a$04$EEVgxdCVoD2bFQ8t5H4U6OEF1fh.ev2Uoa4Qp7TdLU0e7yvo/BQdS",
      "Status": true
    }
  },
  "errcode": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» user|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» CreatedAt|string|true|none||none|
|»»» UpdatedAt|string|true|none||none|
|»»» uid|string|true|none||none|
|»»» userName|string|true|none||none|
|»»» password|string|true|none||none|
|»»» Status|boolean|true|none||none|
|» errcode|integer|true|none||none|

## GET 服务metrics

GET /metrics

metrics

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 用户注册

POST /api/v1/user/register

> Body 请求参数

```json
{
  "username": "tester3",
  "password": "admin123"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» username|body|string| 是 |none|
|» password|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

<h2 id="tocS_Users">Users</h2>

<a id="schemausers"></a>
<a id="schema_Users"></a>
<a id="tocSusers"></a>
<a id="tocsusers"></a>

```json
{
  "id": 1,
  "created_at": "2019-08-24T14:15:22Z",
  "updated_at": "2019-08-24T14:15:22Z",
  "deleted_at": "2019-08-24T14:15:22Z",
  "uid": "string",
  "username": "string",
  "password": "string",
  "status": "1",
  "role_id": 4294967296
}

```
