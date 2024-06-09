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
{
  "username": "tester4",
  "password": "admin12"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» username|body|string| 是 |none|
|» passwor|body|string| 是 |none|

> 返回示例

> 成功

```json
{
  "errCode": 1000,
  "data": {
    "role": {
      "id": 0,
      "createdAt": "10112-01-01 00:00:00",
      "updatedAt": "10112-01-01 00:00:00",
      "name": "",
      "desc": "",
      "code": "",
      "users": null,
      "permission_id": 0,
      "permission": null
    },
    "token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3RlcjQiLCJ1c2VyX3JvbGUiOiIiLCJ0b2tlbl9idWZmZXJfdGltZSI6MCwiZXhwIjoxNzE4NDk3OTQyLCJpYXQiOjE3MTc4OTMxNDIsImlzcyI6InBsYXRmb3JtIiwic3ViIjoidXNlciB0b2tlbiJ9.4bKCX6ZkXSnYJP4KAJKXmn0AAnp6TxXQWg7FDDiyWxlDhEyCuGLtremTBC2bnyfwReXAyKgMjt3nKMSX1Yxwpw",
    "username": "tester4"
  },
  "errMsg": "login success"
}
```

```json
{
  "errCode": 1003,
  "data": {},
  "errMsg": "user name or password error"
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
|»» role|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» createdAt|string|true|none||none|
|»»» updatedAt|string|true|none||none|
|»»» name|string|true|none||none|
|»»» desc|string|true|none||none|
|»»» code|string|true|none||none|
|»»» users|null|true|none||none|
|»»» permission_id|integer|true|none||none|
|»»» permission|null|true|none||none|
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

> 成功

```json
"# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.\n# TYPE go_gc_duration_seconds summary\ngo_gc_duration_seconds{quantile=\"0\"} 6.5926e-05\ngo_gc_duration_seconds{quantile=\"0.25\"} 0.000159127\ngo_gc_duration_seconds{quantile=\"0.5\"} 0.000947659\ngo_gc_duration_seconds{quantile=\"0.75\"} 0.00154217\ngo_gc_duration_seconds{quantile=\"1\"} 0.00154217\ngo_gc_duration_seconds_sum 0.002714882\ngo_gc_duration_seconds_count 4\n# HELP go_goroutines Number of goroutines that currently exist.\n# TYPE go_goroutines gauge\ngo_goroutines 13\n# HELP go_info Information about the Go environment.\n# TYPE go_info gauge\ngo_info{version=\"go1.22.3\"} 1\n# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.\n# TYPE go_memstats_alloc_bytes gauge\ngo_memstats_alloc_bytes 3.463128e+06\n# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.\n# TYPE go_memstats_alloc_bytes_total counter\ngo_memstats_alloc_bytes_total 6.549096e+06\n# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.\n# TYPE go_memstats_buck_hash_sys_bytes gauge\ngo_memstats_buck_hash_sys_bytes 9974\n# HELP go_memstats_frees_total Total number of frees.\n# TYPE go_memstats_frees_total counter\ngo_memstats_frees_total 33403\n# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.\n# TYPE go_memstats_gc_sys_bytes gauge\ngo_memstats_gc_sys_bytes 3.097824e+06\n# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.\n# TYPE go_memstats_heap_alloc_bytes gauge\ngo_memstats_heap_alloc_bytes 3.463128e+06\n# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.\n# TYPE go_memstats_heap_idle_bytes gauge\ngo_memstats_heap_idle_bytes 2.187264e+06\n# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.\n# TYPE go_memstats_heap_inuse_bytes gauge\ngo_memstats_heap_inuse_bytes 5.644288e+06\n# HELP go_memstats_heap_objects Number of allocated objects.\n# TYPE go_memstats_heap_objects gauge\ngo_memstats_heap_objects 18774\n# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.\n# TYPE go_memstats_heap_released_bytes gauge\ngo_memstats_heap_released_bytes 2.00704e+06\n# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.\n# TYPE go_memstats_heap_sys_bytes gauge\ngo_memstats_heap_sys_bytes 7.831552e+06\n# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.\n# TYPE go_memstats_last_gc_time_seconds gauge\ngo_memstats_last_gc_time_seconds 1.717893039520691e+09\n# HELP go_memstats_lookups_total Total number of pointer lookups.\n# TYPE go_memstats_lookups_total counter\ngo_memstats_lookups_total 0\n# HELP go_memstats_mallocs_total Total number of mallocs.\n# TYPE go_memstats_mallocs_total counter\ngo_memstats_mallocs_total 52177\n# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.\n# TYPE go_memstats_mcache_inuse_bytes gauge\ngo_memstats_mcache_inuse_bytes 4800\n# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.\n# TYPE go_memstats_mcache_sys_bytes gauge\ngo_memstats_mcache_sys_bytes 15600\n# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.\n# TYPE go_memstats_mspan_inuse_bytes gauge\ngo_memstats_mspan_inuse_bytes 106080\n# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.\n# TYPE go_memstats_mspan_sys_bytes gauge\ngo_memstats_mspan_sys_bytes 114240\n# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.\n# TYPE go_memstats_next_gc_bytes gauge\ngo_memstats_next_gc_bytes 7.530584e+06\n# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.\n# TYPE go_memstats_other_sys_bytes gauge\ngo_memstats_other_sys_bytes 914698\n# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.\n# TYPE go_memstats_stack_inuse_bytes gauge\ngo_memstats_stack_inuse_bytes 557056\n# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.\n# TYPE go_memstats_stack_sys_bytes gauge\ngo_memstats_stack_sys_bytes 557056\n# HELP go_memstats_sys_bytes Number of bytes obtained from system.\n# TYPE go_memstats_sys_bytes gauge\ngo_memstats_sys_bytes 1.2540944e+07\n# HELP go_threads Number of OS threads created.\n# TYPE go_threads gauge\ngo_threads 10\n# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.\n# TYPE promhttp_metric_handler_requests_in_flight gauge\npromhttp_metric_handler_requests_in_flight 1\n# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.\n# TYPE promhttp_metric_handler_requests_total counter\npromhttp_metric_handler_requests_total{code=\"200\"} 0\npromhttp_metric_handler_requests_total{code=\"500\"} 0\npromhttp_metric_handler_requests_total{code=\"503\"} 0\n"
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

> 成功

```json
{
  "errCode": 0,
  "data": {
    "id": 0,
    "createdAt": "10112-01-01 00:00:00",
    "updatedAt": "10112-01-01 00:00:00",
    "uid": "",
    "username": "",
    "password": "",
    "Status": null,
    "role_id": 0,
    "role": {
      "id": 0,
      "createdAt": "10112-01-01 00:00:00",
      "updatedAt": "10112-01-01 00:00:00",
      "name": "",
      "desc": "",
      "code": "",
      "users": null,
      "permission_id": 0,
      "permission": null
    }
  },
  "errMsg": "register success"
}
```

```json
{
  "errCode": 1004,
  "data": {},
  "errMsg": "user tester4 already exists"
}
```

```json
{
  "errCode": 1004,
  "data": {},
  "errMsg": "user name or password is empty"
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
|»» id|integer|true|none||none|
|»» createdAt|string|true|none||none|
|»» updatedAt|string|true|none||none|
|»» uid|string|true|none||none|
|»» username|string|true|none||none|
|»» password|string|true|none||none|
|»» Status|null|true|none||none|
|»» role_id|integer|true|none||none|
|»» role|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» createdAt|string|true|none||none|
|»»» updatedAt|string|true|none||none|
|»»» name|string|true|none||none|
|»»» desc|string|true|none||none|
|»»» code|string|true|none||none|
|»»» users|null|true|none||none|
|»»» permission_id|integer|true|none||none|
|»»» permission|null|true|none||none|
|» errMsg|string|true|none||none|

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

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer|true|none||'Self-increment numbering'|
|created_at|string(date-time)|false|none||none|
|updated_at|string(date-time)|false|none||none|
|deleted_at|string(date-time)|false|none||none|
|uid|string|false|none||'user uid'|
|username|string|false|none||'user name'|
|password|string|false|none||'user password'|
|status|integer|false|none||'user status(enable/disable)'|
|role_id|integer|false|none||'role foreign id'|

<h2 id="tocS_Role">Role</h2>

<a id="schemarole"></a>
<a id="schema_Role"></a>
<a id="tocSrole"></a>
<a id="tocsrole"></a>

```json
{
  "id": 1,
  "created_at": "2019-08-24T14:15:22Z",
  "updated_at": "2019-08-24T14:15:22Z",
  "deleted_at": "2019-08-24T14:15:22Z",
  "name": "string",
  "desc": "string",
  "permission_id": -2147483648
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer|true|none||'Self-increment numbering'|
|created_at|string(date-time)|false|none||none|
|updated_at|string(date-time)|false|none||none|
|deleted_at|string(date-time)|false|none||none|
|name|string|false|none||'role name'|
|desc|string|false|none||'role description'|
|permission_id|integer|false|none||'Permission id Foreign key'|

