<!-- @format -->

# Push Notification
The notification feature will be wrapped as a rpc service, called by passing a general parameter, here, called `Carrier`.

### Executor
The specific executor that pushing the message asynchronously in the composition pattern.

- Each of your third-party executor should be instantiated somewhere somehow.
- It's initialized in the pattern of functional option.
