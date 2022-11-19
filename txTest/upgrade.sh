#!/bin/bash

HEIGHT=$1
UPGRADE_NAME="v0.7.3"
BINARY_DIR=".cdnode"
CHAIN_ID="oasis-2"
DENOM="unoria"
GAS_PRICE="0.0025"
export DAEMON_NAME="cdnoded"
export DAEMON_HOME="$HOME/$BINARY_DIR"

if ! command -v cosmovisor &>/dev/null; then
  echo "\n\ncosmovisor could not be found"
  exit
fi



# submit upgrade proposal
$DAEMON_NAME tx gov submit-proposal software-upgrade $UPGRADE_NAME --title "Upgrade to $UPGRADE_NAME" --description "Upgrade to $UPGRADE_NAME" --upgrade-info='{"binaries":{"linux/amd64":"https://github.com/cdbo/cdnode/releases/download/v0.7.3/cdnode_linux_amd64.tar.gz?checksum=sha256:af61e03eb0c3c8b2af43a8dbf61558b2520b2de576f4238ee80a874b21893b71"}}' --deposit 10000000$DENOM --upgrade-height $HEIGHT --from me --chain-id $CHAIN_ID --keyring-backend test --home $DAEMON_HOME --node tcp://localhost:26657 --yes --gas-prices $GAS_PRICE$DENOM --gas auto --gas-adjustment 1.5 --broadcast-mode block

# vote on the proposal
$DAEMON_NAME tx gov vote 1 yes --from me --chain-id $CHAIN_ID --keyring-backend test --home $DAEMON_HOME --node tcp://localhost:26657 --yes --gas-prices $GAS_PRICE$DENOM --gas auto --gas-adjustment 1.5 --broadcast-mode block


# make install
# mkdir -p $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME/bin/
# cp `which $DAEMON_NAME` $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME/bin/