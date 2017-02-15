# EnergyChain

A blockchain platform to tokenize and track energy-grid-attached appliance production, usage, availablity,
share ownership, dividend and profit-sharing payouts, and renewable resource credits generation.

<b>SOLES EnergyChain Project</b> is a collaborative cross-industry effort to stand up a platform for the
automated tracking of energy production, usage, availablity, share ownership, dividend and profit-sharing
payouts from grid-attached devices and appliances, including tracking generation and retirement of renewable
resource credits (RRCs) and device audit/registrations. EnergyChain is meant as a platform for energy projects
from the "microgrid" to the municipal and regional level, solar / renewable installation and auditing organizations,
and government organizations to all collaborate to use blockchains to achieve positive environmental impact and
efficient energy distribution. SOLES is also building real-time markets for true pricing of energy and renewable
resource credits tokenized into blockchains in the EnergyChain Project.

# node installation instructions

### 1) install golang and set $GOPATH (compiler & runtime environment setup)

install latest golang from source is the best way... (instructions to follow)

$GOPATH can be set by the following

export $GOPATH='~/.go'

export $PATH="$PATH:$GOPATH"

set these lines in your ~/.bashrc or ~/.bash_profile file

### 2) install tendermint & basecoin (dependency setup)

go get -u github.com/tendermint/tendermint/cmd/tendermint

see : https://tendermint.com/docs , https://tendermint.com/intro   (for background on Tendermint)

note: currently basecoin currently requires the develop branch of tendermint

cd $GOPATH//src/github.com/tendermint/tendermint

git checkout develop

make all

make install

go get -d github.com/tendermint/basecoin/cmd/basecoin

see : https://github.com/tendermint/basecoin (for basecoin background & tutorials, see dev branch)

cd $GOPATH/src/github.com/tendermint/basecoin

make get_vendor_deps

make install

you should have a <b>basecoin</b> binary exectuable in $GOPATH/bin when this completes

### 3) install EnergyChain

go get https://github.com/soles-io/energychain

### 4) running the node, creating a recieve address, sending a transaction

in one command console window, run energychain:

> energychain

in another command console window, run tendermint:

> tendermint

finally, in a third window, you can issue commands to energychain such as:

> energychain --createaddress

> energychain --sendtoaddress [amount] [address]










