# comchain

A simple chain which restricts blockchain addresses ending with letter `s`

## Approch 1

Branch: [approach-1-sdk-fork](https://github.com/arnabghose997/compchain/tree/approach-1-sdk-fork)

In this approach, I essentially forked the `v0.47.3` tag of `cosmos-sdk` (tagged it `v0.47.3-bank`), and added validation check in [Send](https://github.com/arnabghose997/cosmos-sdk/blob/0d9b367a6a688e27e81b7ceab4fca8192f488e10/x/bank/keeper/msg_server.go#L44) and [MultiSend](https://github.com/arnabghose997/cosmos-sdk/blob/0d9b367a6a688e27e81b7ceab4fca8192f488e10/x/bank/keeper/msg_server.go#L81) keepers which prevented address ending with letter `s` from recieving funds.

## Approach 2
Branch: [approach-2-ante-handlers](https://github.com/arnabghose997/compchain/tree/approach-2-ante-handlers)

In this approach, I created a module `x/address` which defines an [Ante Handler](https://github.com/arnabghose997/compchain/blob/approach-2-ante-handlers/x/address/ante/decorator.go) that looks for `MsgSend` and `MsgMultiSend` messages in a transaction. If present, an error is returned if the reciever's address ends with the letter "s".

