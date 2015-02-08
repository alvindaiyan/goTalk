# goTalk
A simple Go based IM system.

![Image of goTalk](https://github.com/alvindaiyan/goTalk/blob/master/config/goTalk.png)

# To Run
1. Install Golang to your computer,
2. Checkout the project,
3. Go to the project folder and run `go run imServer.go` to start the server.

The default port is 9000 (localhost:9000/send), can be changed in imServer.go.

The login services is not implemented yet.

Currently, the project has not connect to any database but will in the future.

# The Software Architecture

The user can has two role in this simple IM system. One is the sender, and one is the receiver. Each user has a channel ([see the detail](https://golang.org/doc/effective_go.html#concurrency)) to store message. 

## The Message Sending Process

This section will introduce the message sending process of the goTalk. The service name is `send`, requires FOUR parameters and is a POST method. The three params are the `Token`, the `id`(send user), the `username` (senduser)  the `sendTo`(user id of the receiver) and the message `content`. (After checking the validability of the token) When a message is sent, the server firstly receive the message and parse to a User object and a Message object. A receive message will be send back. Then the server find the channel regarding of the receive user's id and add to it. If the receiver Id is valid but there is not a channel for it, a new channel will be created and added to the appConfig. An example ajax call is: 

``` javascript
      $.post('/send', 
				{
					username: username,
					id: id,
					sendToId: sendToId,
					content: content
				}, function(data) {
				// data is in Json format
				...
				});
```

## The Message Receiving Process

There is a map to store all the message channel and indexed by the user id. When the server receive a request to sync the user's message status, the services name is `sync`, require two params: `Token`, `id`(user id). The message will be dequeue from the channel and pass back to the user as Json string.

> Problem: The message could be lost if the message is not successfully send at first time. need to be fixed!

## Security (Not Decide yet)

Currently, using AES256.
