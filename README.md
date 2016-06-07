# Go-APNS

Interface for Apple Push Notification System written in Go using their HTTP2 API

## Installation

After setting up Go and the GOPATH variable, you download and install the dependencies by executing the following commands:

```bash
go get -u golang.org/x/net/http2
go get -u golang.org/x/crypto/pkcs12
```

And then install Go-APNS using this line

```bash
go get -u https://github.com/Tantalum73/Go-APNS
```

## Usage

**First step**: creating a `Connection`.

```go
conn, err := goapns.NewConnection("<File path to your certificate in p12 format>", "<password of your certificate>")
if err != nil {
  //Do something more responsible in production
  log.Fatal(err)
}
```

Optionally, you can specify a development or production environment by calling `conn.Development()`. Development is the default environment. Now you are ready for the next step.

**Second step**: build your notification.

According to Apples documentation, a notification consists of a header and a payload, that contains meta-information and the actual alert. In Go-APNS, I condensed it to `Message`.

You only operate with the `Message` struct. It provides a method for every property that you can set. Let's jump right in by looking at an example.

```go
message := goapns.NewMessage().Title("Title").Body("A Test notification :)").Sound("Default").Badge(42)
message.Custom("customKey", "customValue")
```

- You create a new `Message` by calling `goapns.NewMessage()`.
- Specifying the fields is done by calling a method on the message object.
- You can chain it together or call them individually.

**Third Step** push your notification to a device token. Once you have you connection ready and configured the message according to your gusto, you can send the notification to a device token. Often, you gather the tokens in a database and you know best how to get them off there. So let's assume, they are contained in an array or statically typed, like in my case.

The magic happens when you call `Push()` on a `Connection`. The provided `message` is sent to Apples servers asynchronously. Therefore, you get the result delivered in a `chan`. When I say 'response', I mean a `Response` object.

```go
tokens := []string{"<token1>",
  "<token2>"}
responseChannel := make(chan goapns.Response, len(tokens))
conn.Push(message, tokens, responseChannel)

for response := range responseChannel {
  if !response.Sent() {
    //handle the error in response.Error
  } else {
    //the notification was delivered successfully
  }
}
```

_In case, you want to know, what JSON string exactly is pushed to Apple, you can call_ `fmt.Println(message.JSONstring())`_._

Now it is up to you how to handle the error case.

For example, if the device you tried to push to has removed the app you get an `Unregistered` Error (`response.Error == ErrorUnregistered`). In this case, Apple provides the timestamp on which the device started to become unavailable. You can store this status update and the timestamp for the case that the device re-registeres itself. Then, you can compare the received timestamp and decide which token to keep and if you keep pushing to it.

## Values you can set

As mentioned above, you only interact with a `Message`object. There are plenty of methods and I will list them here. You can chain those methods like this

```go
message.Title("Title").Body("A Test notification :)").Sound("Default").Badge(42)
```

_This method will change the Alert_

- Title(string)
- Body(string)
- Sound(string)
- Badge(int) (if left empty, the badge will remain unchanged)
- NoBadgeChange() _if you set the Badge to an int and want to unset it so it stays unchained on the app_

### Tests

_Where are all your tests?_ I am working on it.

## License

```The MIT License (MIT)

Copyright (c) 2016 Andreas Neusuess. (@Klaarname)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.```
