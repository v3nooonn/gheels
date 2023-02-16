<!-- @format -->

# Push Notification
The notification feature will be wrapped as a rpc service, called by passing a general parameter, here, called `Carrier`.

Treat the `notification` package as a service, the `carrier` as the handler in it.

### Carrier
Carrier in the notification service calling, which will generate/mapping the content in the executor.

And the carrier layer is just like a rpc client, receiving outer calling. Later, I will make it a rpc service.

### Executor
The specific executor that pushing the message asynchronously.

- Each of your third-party executor should be instantiated somewhere somehow.
- It's in the pattern of functional option.
