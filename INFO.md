gocrawl
=======

The purpose of this document is to make notes about how gocrawl should work under the hood. The idea is to wrap the relatively complex ssh connection between a Cisco iOS device in set of channels sort a command/response type architecture. Basic work flow would be something like

```
dev := NewDevice(hostname)
if err := dev.Connect(user, passwd); err == nil {
    // Could not connect
}

commands := []string{
    "show ver",
    "show cdp nei",
    "show int",
    "show route"
}

for _, command := range commands {
    if output, err := dev.Send(command); err != nil {
        fmt.Errorf("Command \"%s\" failed: %s", command, err)
    } else {
         // do someting with the output
    }
}
```

output should be free of extra information like the command prompt and just contain the response from the command. Additionally, the device should handle timeout issues, adding little things like "\r" to the end of the commands, establishing the prompt, etc.


Ideas
-----

**`connect` should be a separate procedure**

_pro_
- If the connection is dropped connect can attempt to reconnect
- Multiple connection attempts
- Can create devices and do work on them later
- Can attempt multiple authentication methods

_cons_
- Adds complexity
- Exposes connection methods

_decision_
Connnect should be a seperate procedure to allow for multiple connection attempts and to allow the calling program to handle cases where the connection fails

**Get output from timed out command**
There needs to be a way to get the output from the `Send` command when the call has timed out.

