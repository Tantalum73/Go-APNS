# Go-APNS

Interface for Apples Push Notification System written in Go using their HTTP2 API ☄️

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

**First step**: Create a `Connection`.

```go
conn, err := goapns.NewConnection("<File path to your certificate in p12 format>", "<password of your certificate>")
if err != nil {
  //Do something more responsible in production
  log.Fatal(err)
}
```

Keep the `Connection` around as long as you can. Or as Apple puts it 'You should leave a connection open unless you know it will be idle for an extended period of time--for example, if you only send notifications to your users once a day it is ok to use a new connection each day.'

Optionally, you can specify a development or production environment by calling `conn.Development()`. Development is the default environment. Now you are ready for the next step.

--------------------------------------------------------------------------------

**Second step**: Build your notification.

According to Apples documentation, a notification consists of a header and a payload, that contains meta-information and the actual alert. In Go-APNS I condensed it to `Message`.

You operate only with the `Message` struct. It provides a method for every property that you can set. Let's jump right in by looking at an example.

```go
message := goapns.NewMessage().Title("Title").Body("A Test notification :)").Sound("Default").Badge(42)
message.Custom("customKey", "customValue")
```

- You create a new `Message` by calling `goapns.NewMessage()`.
- Specify values is done by calling a method on the message object.
- You can chain it together or call them individually.

--------------------------------------------------------------------------------

**Third Step** Push your notification to a device token. Once you have you connection ready and configured the message according to your gusto, you can send the notification to a device token. Often, you gather the tokens in a database and you know best how to get them off there. So let's assume, they are contained in an array or, like in my case statically typed.

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

**This method will change the Alert**

- `Title(string)`
- `Body(string)`
- `TitleLocKey(string)`
- `TitleLocArgs([] string)`
- `ActionLocKey(string)`
- `LocKey(string)`
- `LocArgs([] string)`
- `LaunchImage(string)`

**This method will change the Payload**

- `Badge(int)` _if left empty, the badge will remain unchanged_
- `NoBadgeChange()` _if you set the Badge to an int and want to unset it so it stays unchained on the app_
- `Sound(string)`
- `Category(string)`
- `ContentAvailable()` _sets ContentAvailable to 1 and the priority to low, according to Apples documentation_
- `ContentUnvailable()` _lets you reset the ContentAvailable flags you may have set earlier by accident_

**This method will change the Header**

- `APNSID(string)` _An UID you can set to identify the notification. If no ID is specified, Apples server will set one for you automatically_
- `Expiration(time.Time)`
- `PriorityHigh()` _Apple defines a value of 10 as high priority, if you do not specify the priority it will default to high_
- `PriorityLow()` _Apple defines a value of 5 as low priority_
- `Topic(string)` _typically the bundle ID for your app_

## Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/tantalum73/Go-APNS"
)

func main() {

    //creating the Connection by passing the path to a valid certificate and its passphrase
    conn, err := goapns.NewConnection("/../Push Test Push Cert.p12", "<passphrase>")
    if err != nil {
        log.Fatal(err)
    } else {
        conn.Development()
    }


    //composing the Message
    message := goapns.NewMessage().Title("Title").Body("A Test notification :)").Sound("Default").Badge(42)
    message.Custom("key", "val")


    //Tokens from a database or, in my case, statically typed
    tokens := []string{"a26f0000c052865e6631756b1b9d05b4a37ad512fabbe266dd21357b376f0e0e",
        "428dc1d681e576f69f3373d0065b1cdd8da9b76daab39203fa649c26187722c0"}

    //create a channel that gets the Response object passed in,
    //it expects as many responses as there are token to push to
    channel := make(chan goapns.Response, len(tokens))

    //Print the JSON as it is sent to Apples servers
    fmt.Println(message.JSONstring())

    //Perform the push asynchronosly
    conn.Push(message, tokens, channel)

    //iterate through the responses
    for response := range channel {

        if !response.Sent() {
          //handle the error in a way that fits you
            fmt.Printf("\nThere was an error sending to device %v : %v\n", response.Token, response.Error)

                  if response.Error == goapns.ErrorUnregistered {
                      //The device was removed from APNS so the token can't be used anymore.
                      //Update you database accordingly by using the Timestamp object that
                      //gives a hint since when the token coult not be reached anymore.
                      fmt.Printf("\nToken is not registered as valid token anymore since %v\n", response.Timestamp())
                  }

        } else {
            fmt.Printf("\nPush successful for token: %v\n", response.Token)
        }

    }
}
```

It will send this JSON to Apples servers:

```json
{
    "aps": {
        "alert": {
            "title": "Title",
            "body": "A Test notification :)"
        },
        "badge": 42,
        "sound": "Default"
    },
    "key": "val"
}
```

## Apples documentation

If you want to look something up, I also recommend Appled documentation of the APNS topic:

- [Introduction](https://developer.apple.com/library/ios/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/Chapters/Introduction.html#//apple_ref/doc/uid/TP40008194-CH1-SW1)

- [APNS Provider API](https://developer.apple.com/library/ios/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/Chapters/APNsProviderAPI.html#//apple_ref/doc/uid/TP40008194-CH101-SW1)

- [The Remote Notification Payload](https://developer.apple.com/library/ios/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/Chapters/TheNotificationPayload.html#//apple_ref/doc/uid/TP40008194-CH107-SW1)

## Related

I wrote some words about this project in my [blog](https://anerma.de/blog/open-source-project-go-apns). Any feedback is much appreciated. I am [@Klaarname on Twitter](https://twitter.com/klaarname) and would love to hear from you :)

## License

Go-APNS is published under MIT License.

```
Copyright (c) 2016 Andreas Neusuess (@Klaarname)
Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to use,
copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the
Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
