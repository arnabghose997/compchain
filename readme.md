# comchain

A simple chain with a unique restriction

## Approch 1

Branch: [approach-1-sdk-fork](https://github.com/arnabghose997/compchain/tree/approach-1-sdk-fork)

In this approach, I essentially forked the `v0.47.3` tag of `cosmos-sdk` (tagged it `v0.47.3-bank`), and added validation check in [Send](https://github.com/arnabghose997/cosmos-sdk/blob/0d9b367a6a688e27e81b7ceab4fca8192f488e10/x/bank/keeper/msg_server.go#L44) and [MultiSend](https://github.com/arnabghose997/cosmos-sdk/blob/0d9b367a6a688e27e81b7ceab4fca8192f488e10/x/bank/keeper/msg_server.go#L81) keepers which prevented address ending with letter `s` from recieving funds.

