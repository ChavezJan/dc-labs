# Architecture Document

[Architecture File Link](https://drive.google.com/file/d/1jDTwMCjdXHRmfFqGWcl72SUwoDWDTgkd/view?usp=sharing)

# Distributed and Parallel Image Processing

_User Guide_

In order to be able to run the API, it is necessary to install Gin. It is a HTTP web framework, which helps to build applications and microservices in Go (Golang). To install Gin run in a terminal the following command:

    $ go get -u github.com/gin-gonic/gin

Once Gin is installed, the server needs to be available. For running the server use the next command:

    $ go run ./main.go

After running the server in a different terminal client requests may be sent. One of them is access providing username, password and port to be use:

    $ curl -u username:password http://localhost:8080/login

Receiving a similar JSON message with a access token:

> "message": "Hi username, welcome to the DPIP System",

> "token" "OjIE89GzFw"

Other client request is status adding the access token given previously:

    $ curl -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/status

Receiving JSON message with exact time and date:

> "message": "Hi username, the DPIP System is Up and Running"

> "time": "2015-03-07 11:06:39"

The third possible client request is upload, which allows the client to open a local file using the file path and the access token:

    $ curl -F 'data=@/path/to/local/image.png' -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/upload

Receiving JSON message with the file name and size:

> "message": "An image has been successfully uploaded",

> "filename": "image.png",

> "size": "500kb"

Finally the last cliente request is logout using the access token:

    $ curl -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/logout

Receiving JSON message with the following:

> "message": "Bye username, your token has been revoked"
