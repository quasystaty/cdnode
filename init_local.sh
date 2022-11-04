#!/bin/bash

# This script wipe your config folder (~/.cdnode),
# creates a new wallet named "me"
# and prepares everything to be able to start running 
# a fresh chain from height 1.
# 
# This is not meant to be used when trying to sync to an existing chain,
# but rather to work in a local development environment.

BINARY_HOME="$HOME/.cdnode"
CHAIN_ID="oasis-2"
DENOM="unoria"
GAS_PRICE="0.0025"

echo -e "\nRemoving previous config folder ($BINARY_HOME)"
rm -rf $BINARY_HOME

# Set your keyring (the thing that saves your private keys) to the ~/.cdnode folder (not secure, only use for testing env)
echo "Setting keyring to \"test\""
cdnoded config keyring-backend test

# Set the default chain to use
echo "Setting chain-id to \"$CHAIN_ID\""
cdnoded config chain-id $CHAIN_ID

# Create a new wallet named "me"
cdnoded keys add me

# Initialize a new genesis.json file
cdnoded init me --chain-id $CHAIN_ID > /dev/null 2>&1 

# Add your freshly created account to the new chain genesis
cdnoded add-genesis-account me 1000000000$DENOM > /dev/null 2>&1 

# Generate the genesis transaction to create a new validator
cdnoded gentx me 100000000$DENOM --chain-id oasis-2 --commission-rate 0.1 --commission-max-rate 0.2 --commission-max-change-rate 0.01 > /dev/null 2>&1

# Add that gentx transaction to the genesis file
cdnoded collect-gentxs > /dev/null 2>&1

# Edit genesis
sed -i "s/stake/$DENOM/g" ~/.cdnode/config/genesis.json > /dev/null 2>&1
sed -i 's/"inflation": "[^"]*"/"inflation": "0\.0"/g' ~/.cdnode/config/genesis.json > /dev/null 2>&1
sed -i 's/"inflation_rate_change": "[^"]*"/"inflation_rate_change": "0\.0"/g' ~/.cdnode/config/genesis.json > /dev/null 2>&1
sed -i 's/"inflation_min": "[^"]*"/"inflation_min": "0\.0"/g' ~/.cdnode/config/genesis.json > /dev/null 2>&1
sed -i 's/"voting_period": "[^"]*"/"voting_period": "60s"/g' ~/.cdnode/config/genesis.json > /dev/null 2>&1
sed -i 's/"quorum": "[^"]*"/"quorum": "0.000001"/g' ~/.cdnode/config/genesis.json > /dev/null 2>&1

# Edit config.toml to set the block speed to 1s
sed -i 's/^timeout_commit\ =\ .*/timeout_commit\ =\ \"1s\"/g' ~/.cdnode/config/config.toml > /dev/null 2>&1

# Edit app.toml to set the minimum gas price
sed -i "s/^minimum-gas-prices\ =\ .*/minimum-gas-prices\ =\ \"0.0025$DENOM\"/g" ~/.cdnode/config/app.toml > /dev/null 2>&1

# Edit app.toml to enable LCD REST server on port 1317 and REST documentation at http://localhost:1317/swagger/
sed -i 's/enable\ =\ false/enable\ =\ true/g' ~/.cdnode/config/app.toml > /dev/null 2>&1
sed -i 's/swagger\ =\ false/swagger\ =\ true/g' ~/.cdnode/config/app.toml > /dev/null 2>&1

. setup_cosmovisor.sh

echo -e "\n\nYou can now start your chain with 'cdnoded start' or 'cosmovisor start'\n"
