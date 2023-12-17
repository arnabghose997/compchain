# compchain

A simple chain with a unique restriction

## Approch 1

Branch: `main`

In this approach, I essentially forked the `v0.47.3` tag of `cosmos-sdk` (tagged it `v0.47.3-bank`), and added a check [here](https://github.com/arnabghose997/cosmos-sdk/blob/b81512b37907449b40175978d514dfc8c02c4308/x/bank/keeper/msg_server.go#L44) which prevented address ending with letter `s` from recieving funds

## Approach 2

Branch: `approach-2-ante-handlers`

In this approach, I created a module `x/address` which defines an Ante Handler which looks for `MsgSend` and `MsgMultiSend` in a transaction. If present, an error is returned if the reciever's address ends with the letter `s`