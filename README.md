# xdjprox-reader

## a reader for xdjprox log files

A small tool to create human readable files from xdjprox json logs.

## Installation

Compile with go for your desired platform and run the xdjprox-reader binary.

```cli
# running xdjprox-reader with defaults
./xdjprox-reader -buf 512 -in-file ./res_req.log 1> readable.txt
```

```cli
# show help
./xdjprox-reader -h

Usage of xdjprox-reader:
  -buf int
        buffer size in KB (default 64)
  -diff
        enable diff format - remove timestamps and request IDs
  -from string
        from 2006-02-01 15:04:05
  -in-file string
        input file name
  -to string
        from 2006-02-01 15:04:05
```

## Input

```json
{"level":"info","msg":"xdjprox started with config \u0026main.Config{TargetURL:\"http://127.0.0.1:7777\", EntryURL:\":8080\", TimeFormat:\"2006-02-01 15:04:05\", LogRequest:true, LogResponse:true, EnableWrite:true, LogFile:\"res_req.log\", LogAll:false}","time":"2019-12-02T20:49:01+01:00"}
{"level":"info","msg":"forwarded request /server/info","request_id":"5a9effca-07ba-4601-adf4-f67cd1be39fd","session":"","time":"2019-12-02T20:49:13+01:00","type":"forward"}
{"level":"info","msg":"GET /server/info HTTP/1.1\r\nHost: 127.0.0.1:7777\r\nAccept-Charset: utf-8\r\nAccept-Encoding: gzip, deflate\r\nConnection: keep-alive\r\nContent-Length: 0\r\nContent-Type: text/plain\r\nX-Forwarded-Host: \r\nX-Palo-Sv: 0\r\n\r\n","request_id":"5a9effca-07ba-4601-adf4-f67cd1be39fd","session":"","time":"2019-12-02T20:49:13+01:00","type":"request"}
{"level":"info","msg":"HTTP/1.1 200 OK\r\nContent-Length: 33\r\nContent-Type: text/plain;charset=utf-8\r\nServer: Palo\r\nX-Palo-Sv: 1639917714\r\n\r\n19;3;5;10339;0;0;1804290315;0;D;\n","request_id":"5a9effca-07ba-4601-adf4-f67cd1be39fd","session":"","time":"2019-12-02T20:49:13+01:00","type":"response"}
{"level":"info","msg":"forwarded request /server/login","request_id":"a61a1b0b-a832-4d34-99dd-060545b40650","session":"","time":"2019-12-02T20:49:13+01:00","type":"forward"}
{"level":"info","msg":"POST /server/login HTTP/1.1\r\nHost: 127.0.0.1:7777\r\nAccept-Charset: utf-8\r\nAccept-Encoding: gzip, deflate\r\nConnection: keep-alive\r\nContent-Length: 274\r\nContent-Type: text/plain\r\nX-Forwarded-Host: \r\nX-Palo-Sv: 0\r\n\r\nuser=admin\u0026extern_password=admin\u0026machine=127.0.0.1\u0026required=\u0026optional=3a00c07f02\u0026new_name={%22client%22:%22Excel%20Add-in%22,%22client_ver%22:%2216866%22,%22lib%22:%22libpalo_ng%22,%22lib_ver%22:%2219.3.2.5947%22,%22desc%22:%22user%20login%22}\u0026external_identifier=de_DE","request_id":"a61a1b0b-a832-4d34-99dd-060545b40650","session":"","time":"2019-12-02T20:49:13+01:00","type":"request"}
{"level":"info","msg":"HTTP/1.1 200 OK\r\nContent-Length: 46\r\nContent-Type: text/plain;charset=utf-8\r\nServer: Palo\r\nX-Palo-Sv: 1639917715\r\n\r\nB9i0QF5A5xhFQvvTFkRcRz1dp4TbQqMs;300;3b;0;\"\";\n","request_id":"a61a1b0b-a832-4d34-99dd-060545b40650","session":"","time":"2019-12-02T20:49:13+01:00","type":"response"}
{"level":"info","msg":"forwarded request /server/logout","request_id":"c011b5cb-ddc5-4821-aa6f-bfe13419a5c9","session":"B9i0QF5A5xhFQvvTFkRcRz1dp4TbQqMs","time":"2019-12-02T20:49:13+01:00","type":"forward"}
{"level":"info","msg":"GET /server/logout?sid=B9i0QF5A5xhFQvvTFkRcRz1dp4TbQqMs HTTP/1.1\r\nHost: 127.0.0.1:7777\r\nAccept-Charset: utf-8\r\nAccept-Encoding: gzip, deflate\r\nConnection: keep-alive\r\nContent-Length: 0\r\nContent-Type: text/plain\r\nX-Forwarded-Host: \r\nX-Palo-Sv: 1639917714\r\n\r\n","request_id":"c011b5cb-ddc5-4821-aa6f-bfe13419a5c9","session":"B9i0QF5A5xhFQvvTFkRcRz1dp4TbQqMs","time":"2019-12-02T20:49:13+01:00","type":"request"}
{"level":"info","msg":"HTTP/1.1 200 OK\r\nContent-Length: 3\r\nContent-Type: text/plain;charset=utf-8\r\nServer: Palo\r\nX-Palo-Sv: 1639917716\r\n\r\n1;\n","request_id":"c011b5cb-ddc5-4821-aa6f-bfe13419a5c9","session":"B9i0QF5A5xhFQvvTFkRcRz1dp4TbQqMs","time":"2019-12-02T20:49:13+01:00","type":"response"}
```

## Output
```txt
----------- 2019-12-02T20:49:01+01:00 - (/) -------------------
xdjprox started with config &main.Config{TargetURL:"http://127.0.0.1:7777", EntryURL:":8080", TimeFormat:"2006-02-01 15:04:05", LogRequest:true, LogResponse:true, EnableWrite:true, LogFile:"res_req.log", LogAll:false}

----------- 2019-12-02T20:49:13+01:00 - (request/5a9effca-07ba-4601-adf4-f67cd1be39fd) -------------------
GET /server/info HTTP/1.1
Host: 127.0.0.1:7777
Accept-Charset: utf-8
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 0
Content-Type: text/plain
X-Forwarded-Host: 
X-Palo-Sv: 0

----------- 2019-12-02T20:49:13+01:00 - (response/5a9effca-07ba-4601-adf4-f67cd1be39fd) -------------------
HTTP/1.1 200 OK
Content-Length: 33
Content-Type: text/plain;charset=utf-8
Server: Palo
X-Palo-Sv: 1639917714

19;3;5;10339;0;0;1804290315;0;D;

----------- 2019-12-02T20:49:13+01:00 - (request/a61a1b0b-a832-4d34-99dd-060545b40650) -------------------
POST /server/login HTTP/1.1
Host: 127.0.0.1:7777
Accept-Charset: utf-8
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 274
Content-Type: text/plain
X-Forwarded-Host: 
X-Palo-Sv: 0

user=admin&extern_password=admin&machine=127.0.0.1&required=&optional=3a00c07f02&new_name={"client":"Excel Add-in","client_ver":"16866","lib":"libpalo_ng","lib_ver":"19.3.2.5947","desc":"user login"}&external_identifier=de_DE

----------- 2019-12-02T20:49:13+01:00 - (response/a61a1b0b-a832-4d34-99dd-060545b40650) -------------------
HTTP/1.1 200 OK
Content-Length: 46
Content-Type: text/plain;charset=utf-8
Server: Palo
X-Palo-Sv: 1639917715

B9i0QF5A5xhFQvvTFkRcRz1dp4TbQqMs;300;3b;0;"";

----------- 2019-12-02T20:49:13+01:00 - (request/c011b5cb-ddc5-4821-aa6f-bfe13419a5c9) -------------------
GET /server/logout?sid=B9i0QF5A5xhFQvvTFkRcRz1dp4TbQqMs HTTP/1.1
Host: 127.0.0.1:7777
Accept-Charset: utf-8
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 0
Content-Type: text/plain
X-Forwarded-Host: 
X-Palo-Sv: 1639917714

----------- 2019-12-02T20:49:13+01:00 - (response/c011b5cb-ddc5-4821-aa6f-bfe13419a5c9) -------------------
HTTP/1.1 200 OK
Content-Length: 3
Content-Type: text/plain;charset=utf-8
Server: Palo
X-Palo-Sv: 1639917716

1;
```

## License

Licensed under the [MIT](./LICENSE) license.