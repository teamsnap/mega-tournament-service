# TeamSnap Mega Tournament Service

Greetings! This is the Mega Tournament Service, based loosely on the ever popular https://github.com/teamsnap/teamsnap-mega-tournament

## Your Task

This Service is barely a skeleton that does _nothing_ useful. Your task is to make it do _anything_ useful. You may find inspiration in the [openapi spec](spec.yml), but that's not a strict requirement by any means.
Feel free to add or modify APIs as you deem necessary for the objective of making the service useful.    
We are a curious group, so we're probably going to ask about what you did, and why you did it! This shouldn't be a stressful exercise. We also recognize that everyone has their own stuff going on in life. We don't expect you to spend huge amounts of time on this, so put into it what you feel comfortable with. When you are done, zip it up and send it back to us!


## Bootstrapping

This service is written in Go. To get it running locally, you'll need to head over over to https://go.dev/ and get the Go binary. We make some assumptions that this is running on a Linux machine. There are some mostly functional docker files included, but you don't have to use those!

### Starting the app

`go run main.go` will start the service.

`make lint` will run the linter.

`make test` will run the tests.


### Navigating the code

The handlers directory is a great place to start! 