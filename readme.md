# cdnode
The layer 1 Cosmos blockchain for the Crypto Dollar



# Info  
### [testnet]  

- Staking denom: `ubir`  
- Denom exponent: `6`  
- Wallet prefix: `birth`  
- Average block speed: `5s`  
- Public RPC: [https://chaos-rpc.cryptodollar.dev](https://chaos-rpc.cryptodollar.dev)  
- Public LCD: [https://chaos-lcd.cryptodollar.dev](https://chaos-lcd.cryptodollar.dev)  
- LCD Docs: [https://chaos-lcd.cryptodollar.dev/swagger/](https://chaos-lcd.cryptodollar.dev/swagger/)  
- Block explorer: [https://explorer.cryptodollar.dev](https://explorer.cryptodollar.dev)

# Chain Specs  

- Governance disabled  
- Automatic minting disabled
- CosmWasm 1.0.0 enabled  
- IBC enabled  
- Interchain Accounts enabled  
- Staking denom `ubir` total supply: `1B`


# Install

- [Install Golang v1.18+](https://go.dev/doc/install)
- Set your [$GOBIN and $GOPATH](https://pkg.go.dev/cmd/go#hdr-GOPATH_environment_variable) env
- `git clone https://github.com/cdbo/cdnode.git`  
- `cd cdnode`  
- `make install`  

At this point, you have a `cdnoded` binary installed, used to run a node/validator and execute/query the chain.

# Running a local development environment  
Read through [init_local.sh](init_local.sh) for more details.  

### 1. `./init_local.sh`  
### 2. `cdnoded start`

# Running a node  

### 1. `cdnoded init <moniker> --chain-id oasis-1`  
Moniker is your node name.  

### 2. Adjust some configuration parameters  

<a name="config"></a>
> The default location for configuration files is `$HOME/.cdnode/config`  

`app.toml`
```bash
minimum-gas-prices = "0.0025ubir"
```

`config.toml`
```bash
moniker = "moniker_entered_at_step_1"
persistent_peers = "d9f121783c3e80c0e2c98da9f9c33cf5838a49c1@167.99.177.244:26656"  
```

`client.toml`

```bash
chain-id = "oasis-1"  
```



### 3. Download the [genesis.json](https://raw.githubusercontent.com/cdbo/cdnode/master/genesis.json) file to your [config](#config) folder  

### 4. `cdnoded start`  

At this point, your node will start synchronizing with the existing network and catch up on blocks. This might take a while. You can verify the state of your node with the following command: `cdnoded status`, look for the `catching_up` property; once `false`, that means you are in sync with the rest of the chain.  

It is recommended to run the this binary as a daemon like systemd. Here is an example of a `/etc/systemd/system/cdnode.service`:  
_replace $USER with your username and $GOBIN with the path where `cdnoded` is installed._

```bash
[Unit]  
Description=CDNode Daemon  
After=network.target  
  
[Service]  
Type=simple  
User=$USER  
ExecStart=$GOBIN/cdnoded start    
Restart=on-abort  
  
[Install]  
WantedBy=multi-user.target  
  
[Service]  
LimitNOFILE=65535    
```


# Running a validator  

Once you have a fully sync'd node, you can start signing blocks by becoming a validator.  

### 1. Make you have a wallet configured  
with `cdnoded keys list`. If you don't, add one: `cdnoded keys add <wallet name>`. If you need funds, hit the faucet on discord with `/request <wallet address>`.  

### 2. Execute the **create-validator** transaction:  
```bash
cdnoded tx staking create-validator \\  
--amount="1000000000ubir" \\  
--pubkey=$(cdnoded tendermint show-validator) \\  
--moniker="My Node" \\  
--chain-id="oasis-1" \\  
--commission-rate="0.05" \\ 
--commission-max="0.20" \\ 
--commission-max-change-rate="0.01" \\ 
--min-self-delegation="1000000ubir" \\ 
--gas="auto" \\ 
--gas-prices="0.0025ubir" \\ 
--gas-adjustment="1.75" \\ 
--from="myWalletName" \\ 
```

* [CosmosHub example](https://hub.cosmos.network/main/validators/validator-setup.html#create-your-validator)


<br />
<br />
<br />
<br />
 
> Todo:
> - ~~add build script~~
> - ~~store releases~~
> - ~~store genesis~~
> - links to snapshots
> - ~~How-To guide~~
