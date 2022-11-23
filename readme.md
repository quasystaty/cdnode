# cdnode

The layer 1 Cosmos blockchain for the Crypto Dollar

# Info

### [testnet]

- Current binary version: v0.7.3  
- Staking denom: `unoria`
- Denom exponent: `6`
- Wallet prefix: `noria`
- Average block speed: `5s`
- Public RPC: [https://chaos-rpc.cryptodollar.dev](https://chaos-rpc.cryptodollar.dev)
- Public LCD: [https://chaos-lcd.cryptodollar.dev](https://chaos-lcd.cryptodollar.dev)
- LCD Docs: [https://chaos-lcd.cryptodollar.dev/swagger/](https://chaos-lcd.cryptodollar.dev/swagger/)
- Block explorer: [https://explorer.cryptodollar.dev](https://explorer.cryptodollar.dev)

# Chain Specs

- Automatic minting disabled
- Manual minting/burning of custom denoms by whitelisted minters
- CosmWasm 1.1 enabled
- IBC enabled
- Staking denom `unoria` total supply: `1B`

# Install

- [Install Golang v1.18+](https://go.dev/doc/install)
- Set your [$GOBIN and $GOPATH](https://pkg.go.dev/cmd/go#hdr-GOPATH_environment_variable) env
- `git clone https://github.com/cdbo/cdnode.git`
- `cd cdnode`
- `make install`

At this point, you have a `cdnoded` binary installed, used to run a node/validator and execute/query the chain.

# Running a local development environment

Read through [scripts/init_local.sh](./scripts/init_local.sh) for more details.

### 1. `./scripts/init_local.sh`

### 2. `cdnoded start`

# Running a node

### 1. `cdnoded init <moniker> --chain-id oasis-2`

Moniker is your node name.

### 2. Adjust some configuration parameters

<a name="config"></a>

> The default location for configuration files is `$HOME/.cdnode/config`

`app.toml`

```bash
minimum-gas-prices = "0.0025unoria"
```

`config.toml`

```bash
moniker = "moniker_entered_at_step_1"
persistent_peers = "5dc187550a021298f88c3e5905c7dbdd790ff040@167.99.177.244:26656"
```

`client.toml`

```bash
chain-id = "oasis-2"
```

### 3. Download the [genesis.json](https://raw.githubusercontent.com/cdbo/cdnode/master/genesis.json) file to your [config](#config) folder

StateSync is enabled on the public RPC server and can dramatically speed up catch-up time to the latest block.  
It can be enabled by modifying some config parameters before starting the `cdnoded`.  
Here are the changes required to enable StateSync catch-up:

```bash
# config.toml

[statesync]

enable = true
rpc_servers = "https://chaos-rpc.cryptodollar.dev:443,http://167.99.177.244:26657"
trust_height = <insert previous block height which is a factor of 500>
trust_hash = <insert block hash of that block height>

```

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
--amount="1000000000unoria" \\
--pubkey=$(cdnoded tendermint show-validator) \\
--moniker="My Node" \\
--chain-id="oasis-2" \\
--commission-rate="0.05" \\
--commission-max-rate="0.20" \\
--commission-max-change-rate="0.01" \\
--min-self-delegation="1000000" \\
--gas="auto" \\
--gas-prices="0.0025unoria" \\
--gas-adjustment="1.75" \\
--from="myWalletName" \\
```

- [CosmosHub example](https://hub.cosmos.network/main/validators/validator-setup.html#create-your-validator)

# Minting / Burning CRD

The **Coinmaster** module allows the minting and burning of whitelisted native coins.

### Permission

The **Coinmaster** module has two guards in place:

- Whitelisted denominations. You can only mint/burn whitelisted denoms. The initial whitelist is set to an empty string, which allows minting and burning any denomination. This can be changed through governance with a `param-change` proposal such as:

```bash
{
  "title": "Update whitelisted denominations",
  "description": "Update whitelisted denominations",
  "changes": [
    {
      "subspace": "coinmaster",
      "key": "Denoms",
      "value": "ucrd,unoria,nwbtc"
    }
  ],
  "deposit": "1000000unoria"
}
```

- Whitelisted minters. Only Minters are allowed to mint/burn coins. This can be changed through governance with a `param-change` proposal such as:

```bash
{
  "title": "Update whitelisted minters",
  "description": "Update whitelisted minters",
  "changes": [
    {
      "subspace": "coinmaster",
      "key": "Minters",
      "value": "noria1qrh465lh5ygkaqu0nc2wdxfv5nkmwl3xlqf7jl"
    }
  ],
  "deposit": "1000000unoria"
}
```

In its current implementation, the module only supports 1 minter, but this can be changed in the future to allow more minters.  
The default value of the minter is an empty string, which allows anyone to mint/burn.

### Minting

```bash
cdnoded tx coinmaster mint \\
--amount="1000000ucrd" \\
--from="myWalletName" \\
--gas="auto" \\
--gas-prices="0.0025unoria" \\
--gas-adjustment="1.75" \\
--chain-id="oasis-2"
```

### Burning

```bash
cdnoded tx coinmaster burn \\
--amount="1000000ucrd" \\
--from="myWalletName" \\
--gas="auto" \\
--gas-prices="0.0025unoria" \\
--gas-adjustment="1.75" \\
--chain-id="oasis-2"
```

<br />
<br />
 
> Todo:
> - ~~add build script~~
> - ~~store releases~~
> - ~~store genesis~~
> - ~~How-To guide~~
> - links to snapshots
> - improve security on coinmaster
