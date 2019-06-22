# OpenAPI 本地浏览教程

## 简介

由于 `Device Cloud` 采用 `OpenAPI` 格式编写接口文档, 可以采用 `Swagger-UI` 工具浏览.

本教程目的是搭建本地的 `OpenAPI` 预览工具.

## 具体教程

1. 运行 `OpenAPI HTTP Server`

    $ cd $GOPATH/src/github.com/nayotta/metathings
    $ cat << EOF > httpd.py
    #!/usr/bin/env python3
    from http.server import HTTPServer, SimpleHTTPRequestHandler, test
    import sys
    class CORSRequestHandler (SimpleHTTPRequestHandler):
        def end_headers (self):
            self.send_header('Access-Control-Allow-Origin', '*')
            SimpleHTTPRequestHandler.end_headers(self)

    if __name__ == '__main__':
        test(CORSRequestHandler, HTTPServer, port=int(sys.argv[1]) if len(sys.argv) > 1 else 8000)
    EOF
    $ chmod +x httpd.py
    $ ./httpd.py

2. 运行 `Swagger-UI`

    $ docker run --rm -p 80:8080 swaggerapi/swagger-ui

3. 浏览 Device Cloud 接口

打开浏览器, 输入网址: http://localhost

打开网页之后, 输入: http://localhost:8000/openapi/device/openapi.yaml

最后点击 `Explore` 按钮即可
