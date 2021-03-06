# 预授权命令行工具

[预授权设计](../spec/txs/approve.md)
```
qoscli tx
tx subcommands

Usage:
  qoscli tx [command]

Available Commands:
  ...
  
  create-approve   Create approve
  increase-approve Increase approve
  decrease-approve Decrease approve
  use-approve      Use approve
  cancel-approve   Cancel approve
  
  ...
  
Flags:
  -h, --help   help for tx

Global Flags:
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
      --home string       directory for config and data (default "/home/imuge/.qoscli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors

Use "qoscli tx [command] --help" for more information about a command.
```

## create

```
qoscli tx create-approve --help
Create approve

Usage:
  qoscli tx create-approve [flags]

Flags:
      --async             broadcast transactions asynchronously
      --chain-id string   Chain ID of tendermint node
      --from string       Name of approve creator
  -h, --help              help for create-approve
      --max-gas int       gas limit to set per tx
      --node string       <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --qos int           Amount of QOS
      --qscs string       Names and amounts of QSCs
      --to string         Address of approve receiver
      --trust-node        Trust connected full node (don't verify proofs for responses) (default true)

Global Flags:
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
      --home string       directory for config and data (default "/home/imuge/.qoscli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors
```
主要参数：

- from  授权账户，keys中保存的name
- to    被授权账户地址
- qos   授权qos值
- qscs  授权qsc列表，[amount1][qsc1],[amount2][qsc2],...，以半角逗号相隔，eg: 100qsc1,100qsc2

Arya向Sansa授权100个qos，100个qstar
```
qoscli tx create-approve --from=Arya --to=address1t7eadnyl8g6ct9xyrasvz4rdztvkeqpc0hzujh --qos=100 --qscs=100qstar
Password to sign with 'Arya':
{"check_tx":{},"deliver_tx":{},"hash":"9917953D8CDE80F457CD072DBCE73A36449B7A7C","height":"333"}
```

## query

查询预授权
```
qoscli query approve --help
Query approve by from and to

Usage:
  qoscli query approve [flags]

Flags:
      --chain-id string   Chain ID of tendermint node
      --from string       Address of approve creator
      --height int        block height to query, omit to get most recent provable block
  -h, --help              help for approve
      --node string       <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --to string         Address of approve receiver
      --trust-node        Trust connected full node (don't verify proofs for responses)

Global Flags:
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
      --home string       directory for config and data (default "/home/imuge/.qoscli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors
```
主要参数：

- from  授权账户地址
- to    被授权账户地址
```
approve query approve --from=address1evmncf3z99a4uhq5n5yjwputfqmtjsuknv43fn --to=address1t7eadnyl8g6ct9xyrasvz4rdztvkeqpc0hzujh
{
  "from": "address1evmncf3z99a4uhq5n5yjwputfqmtjsuknv43fn",
  "to": "address1t7eadnyl8g6ct9xyrasvz4rdztvkeqpc0hzujh",
  "qos": "100",
  "qscs": [
    {
      "coin_name": "qstar",
      "amount": "100"
    }
  ]
}
```

## increase

```
qoscli tx increase-approve --help
Increase approve

Usage:
  qoscli tx increase-approve [flags]

Flags:
      --async             broadcast transactions asynchronously
      --chain-id string   Chain ID of tendermint node
      --from string       Name of approve creator
  -h, --help              help for increase-approve
      --max-gas int       gas limit to set per tx
      --node string       <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --qos int           Amount of QOS
      --qscs string       Names and amounts of QSCs
      --to string         Address of approve receiver
      --trust-node        Trust connected full node (don't verify proofs for responses) (default true)

Global Flags:
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
      --home string       directory for config and data (default "/home/imuge/.qoscli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors
```
主要参数：

- from  授权账户，keys中保存的name
- to    被授权账户地址
- qos   授权qos值
- qscs  授权qsc列表，[amount1][qsc1],[amount2][qsc2],...，以半角逗号相隔，eg: 100qsc1,100qsc2

Arya向Sansa增加授权100个qos，100个qstar
```
qoscli tx increase-approve --from=Arya --to=address1t7eadnyl8g6ct9xyrasvz4rdztvkeqpc0hzujh --qos=100 --qscs=100qstar
Password to sign with 'Arya':
{"check_tx":{},"deliver_tx":{},"hash":"3C06676C53A5439D39CB4D0FBA3213C44DC1BA8E","height":"406"}
```

## decrease

```
qoscli tx decrease-approve --help
Decrease approve

Usage:
  qoscli tx decrease-approve [flags]

Flags:
      --async             broadcast transactions asynchronously
      --chain-id string   Chain ID of tendermint node
      --from string       Name of approve creator
  -h, --help              help for decrease-approve
      --max-gas int       gas limit to set per tx
      --node string       <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --qos int           Amount of QOS
      --qscs string       Names and amounts of QSCs
      --to string         Address of approve receiver
      --trust-node        Trust connected full node (don't verify proofs for responses) (default true)

Global Flags:
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
      --home string       directory for config and data (default "/home/imuge/.qoscli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors
```
主要参数：

- from  授权账户，keys中保存的name
- to    被授权账户地址
- qos   授权qos值
- qscs  授权qsc列表，[amount1][qsc1],[amount2][qsc2],...，以半角逗号相隔，eg: 100qsc1,100qsc2

Arya向Sansa减少授权100个qos，100个qstar
```
qoscli tx decrease-approve --from=Arya --to=address1t7eadnyl8g6ct9xyrasvz4rdztvkeqpc0hzujh --qos=100 --qscs=100qstar
Password to sign with 'Arya':
{"check_tx":{},"deliver_tx":{},"hash":"9DC18AD3CB0B59FCD354C267D8C22A1CC75E5624","height":"414"}
```

## use

```
qoscli tx use-approve --help
Use approve

Usage:
  qoscli tx use-approve [flags]

Flags:
      --async             broadcast transactions asynchronously
      --chain-id string   Chain ID of tendermint node
      --from string       Address of approve creator
  -h, --help              help for use-approve
      --max-gas int       gas limit to set per tx
      --node string       <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --qos int           Amount of QOS
      --qscs string       Names and amounts of QSCs
      --to string         Name of approve receiver
      --trust-node        Trust connected full node (don't verify proofs for responses) (default true)

Global Flags:
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
      --home string       directory for config and data (default "/home/imuge/.qoscli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors
```
主要参数：

- from  授权账户地址
- to    被授权账户，keys中保存的name
- qos   授权qos值
- qscs  授权qsc列表，[amount1][qsc1],[amount2][qsc2],...，以半角逗号相隔，eg: 100qsc1,100qsc2

Sansa使用Arya向自己授权的10个qos，10个qstar
```
qoscli tx use-approve --from=address1evmncf3z99a4uhq5n5yjwputfqmtjsuknv43fn --to=Sansa --qos=10 --qscs=10qstar
Password to sign with 'Sansa':
{"check_tx":{},"deliver_tx":{},"hash":"0573760D6B316E6695FBB63A56F2A20C0635FCAE","height":"437"}
```

## cancel

```
qoscli tx cancel-approve --help
Cancel approve

Usage:
  qoscli tx cancel-approve [flags]

Flags:
      --async             broadcast transactions asynchronously
      --chain-id string   Chain ID of tendermint node
      --from string       Name of approve creator
  -h, --help              help for cancel-approve
      --max-gas int       gas limit to set per tx
      --node string       <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --to string         Address of approve receiver
      --trust-node        Trust connected full node (don't verify proofs for responses) (default true)

Global Flags:
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
      --home string       directory for config and data (default "/home/imuge/.qoscli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors
```
主要参数：

- from  授权账户，keys中保存的name
- to    被授权账户地址

Arya取消向Sansa授权任何资产
```
qoscli tx cancel-approve --from=Arya --to=address1t7eadnyl8g6ct9xyrasvz4rdztvkeqpc0hzujh
Password to sign with 'Arya':
{"check_tx":{},"deliver_tx":{},"hash":"BA45F8416780C76468C925E34372B05F5A7FEAAC","height":"484"}
```