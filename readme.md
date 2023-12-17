# compchain

A simple chain with a unique restriction

## Approach 2
Branch: [approach-2-ante-handlers](https://github.com/arnabghose997/compchain/tree/approach-2-ante-handlers)

In this approach, I created a module `x/address` which defines an [Ante Handler](https://github.com/arnabghose997/compchain/blob/approach-2-ante-handlers/x/address/ante/decorator.go) that looks for `MsgSend` and `MsgMultiSend` messages in a transaction. If present, an error is returned if the reciever's address ends with the letter "s".
