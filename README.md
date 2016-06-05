# Go-APNS

Interface for Apple Push Notification System written in Go using their HTTP2 API

# Installation

After setting up Go and the GOPATH variable, you download and install the dependencies by executing the following commands

```
go get -u golang.org/x/net/http2
go get -u golang.org/x/crypto/pkcs12
```

And then install Go-APNS using this line

```
go get -u https://github.com/Tantalum73/Go-APNS
```

# Usage

**First step**: creating a `Connection`.

```go
conn, err := goapns.NewConnection("<File path to your certificate in p12 format>", "<password of your certificate>")
if err != nil {
  //Do something responsible with the error, like printing it.
  return
}
```

Optionally, you can specify a development or production environment by calling `conn.Development()`. Development is the default environment.

Now you are ready for the next step.

**Second step**: build your notification.

According to Apples documentation, a notification consists of a header and a payload containing meta-information and the actual alert. In Go-APNS, I condensed it into a `Message`.

You only operate with the `Message` struct. It provides a method for every property that you can set. Let's jump right in by looking at an example.

```go
message := goapns.NewMessage().Title("Title").Body("A Test notification :)").Sound("Default").Badge(42)
m.Custom("key", "val")
```

- You create a new `Message` by calling `goapns.NewMessage()`.
- Specifying the fields is done by calling a method on the message object.
- You can chain it together or call them individually.

# License

```The MIT License (MIT)

Copyright (c) 2016 Andreas Neusuess. (@Klaarname)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE. ```
