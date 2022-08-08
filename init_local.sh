#!/bin/bash

# This script wipe your config folder (~/.cdnode),
# creates a new wallet named "me"
# and prepares everything to be able to start running 
# a fresh chain from height 1.
# 
# This is not meant to be used when trying to sync to an existing chain,
# but rather to work in a local development environment.

rm -rf ~/.cdnode

# Set your keyring (the thing that saves your private keys) to the ~/.cdnode folder (not secure, only use for testing env)
cdnoded config keyring-backend test

# Set the default chain to use
cdnoded config chain-id oasis-2

# Create a new wallet named "me"
cdnoded keys add me

# Initialize a new genesis.json file
cdnoded init me --chain-id oasis-2 > /dev/null 2>&1 

# Add your freshly created account to the new chain genesis
cdnoded add-genesis-account me 1000000000unoria > /dev/null 2>&1 

# Generate the genesis transaction to create a new validator
cdnoded gentx me 100000000unoria --chain-id oasis-2 --commission-rate 0.1 --commission-max-rate 0.2 --commission-max-change-rate 0.01 > /dev/null 2>&1

# Add that gentx transaction to the genesis file
cdnoded collect-gentxs > /dev/null 2>&1

# Edit genesis to use a non-default demonimation
sed -i 's/stake/unoria/g' ~/.cdnode/config/genesis.json > /dev/null 2>&1

# Edit app.toml to set the minimum gas price
sed -i 's/minimum-gas-prices\ =\ \"\"/minimum-gas-prices\ =\ \"0.0025unoria\"/g' ~/.cdnode/config/app.toml > /dev/null 2>&1

# Edit app.toml to enable LCD REST server on port 1317 and REST documentation at http://localhost:1317/swagger/
sed -i 's/enable\ =\ false/enable\ =\ true/g' ~/.cdnode/config/app.toml > /dev/null 2>&1
sed -i 's/swagger\ =\ false/swagger\ =\ true/g' ~/.cdnode/config/app.toml > /dev/null 2>&1

echo ""
echo "You can now start your chain with 'cdnoded start'"
