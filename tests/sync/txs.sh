#!/bin/bash

set -e
set -x

checkConsensus() {
  if grep -q "CONSENSUS FAILURE" $DIR/node1/log.txt; then
    killall certikd
    echo "CONSENSUS FAILURE!"
    exit 1
  fi
}

# ------------------------------------------------------------------
# Transaction sequence
#
# use `$CERTIKCLI0` to send txs from jack or bob running old binary
# use `$CERTIKCLI1` to send txs from mary running latest binary
# ------------------------------------------------------------------

# Add tokens to mary
$CERTIKCLI0 tx send $jack $mary 100000000uctk --from $jack -y
sleep 6
$CERTIKCLI1 query account $jack
$CERTIKCLI1 query account $bob
$CERTIKCLI1 query account $mary

checkConsensus

# auth
$CERTIKCLI0 tx unlock $jack 500000uctk --from $bob -y
sleep 6
$CERTIKCLI1 query account $jack

checkConsensus

# bank
$CERTIKCLI1 tx locked-send $mary $jack 500000uctk --from $mary -y
sleep 6
$CERTIKCLI1 query account $jack
$CERTIKCLI1 query account $mary

checkConsensus

# cert
$CERTIKCLI1 query cert certifiers
$CERTIKCLI0 tx cert certify-validator certikvalconspub1zcjduepqff623akv26we89w9qz6nk7yq66ms5tlhmn5p7v8rqv4z2ur9puhqmxvkpk --from $bob -y
sleep 6
$CERTIKCLI1 query cert validators

$CERTIKCLI0 tx cert certify-platform certikvalconspub1zcjduepqff623akv26we89w9qz6nk7yq66ms5tlhmn5p7v8rqv4z2ur9puhqmxvkpk xxxx --from $bob -y
sleep 6
$CERTIKCLI1 query cert platform certikvalconspub1zcjduepqff623akv26we89w9qz6nk7yq66ms5tlhmn5p7v8rqv4z2ur9puhqmxvkpk

$CERTIKCLI0 tx cert issue-certificate AUDITING ADDRESS C --from $bob -y
sleep 6
$CERTIKCLI0 tx cert issue-certificate COMPILATION SOURCECODEHASH C --compiler A --bytecode-hash B --from $bob -y
sleep 6
$CERTIKCLI1 query cert certificates
id=$($CERTIKCLI1 query cert certificates | grep certificateid)
id=${id:17:60}

$CERTIKCLI0 tx cert decertify-validator certikvalconspub1zcjduepqff623akv26we89w9qz6nk7yq66ms5tlhmn5p7v8rqv4z2ur9puhqmxvkpk --from $bob -y
sleep 6
$CERTIKCLI1 query cert validators

$CERTIKCLI0 tx cert revoke-certificate $id --from $bob -y
sleep 6
$CERTIKCLI1 query cert certificates

checkConsensus

# cvm
txhash=$($CERTIKCLI1 tx cvm deploy $PROJ_ROOT/tests/simple.sol --from $mary -y | grep txhash)
txhash=${txhash:8}
sleep 6
addr=$($CERTIKCLI1 query tx $txhash | grep value | sed -n '2p')
addr=${addr:13}

$CERTIKCLI1 tx cvm call $addr set 123 --from $mary -y
sleep 6

$CERTIKCLI1 tx cvm call $addr get --from $mary -y
sleep 6

checkConsensus

# oracle
$CERTIKCLI1 tx oracle create-operator $mary 100000uctk --from $mary -y
sleep 6
$CERTIKCLI1 query oracle operators

$CERTIKCLI0 tx oracle create-task --contract A --function B --bounty 10000uctk --wait 4 --from $bob -y
sleep 6
$CERTIKCLI1 query oracle task --contract A --function B

$CERTIKCLI1 tx oracle deposit-collateral $mary 30000uctk --from $mary -y
sleep 6

$CERTIKCLI1 tx oracle withdraw-collateral $mary 10000uctk --from $mary -y
sleep 6
$CERTIKCLI1 query oracle operators

$CERTIKCLI1 tx oracle respond-to-task --contract A --function B --score 99 --from $mary -y
sleep 6
$CERTIKCLI1 query oracle response --contract A --function B --operator $mary
$CERTIKCLI1 query oracle operator $mary

$CERTIKCLI1 tx oracle claim-reward $mary --from $mary -y
sleep 6
$CERTIKCLI1 query oracle operator $mary

$CERTIKCLI0 tx oracle delete-task --contract A --function B --force=true --from $bob -y
sleep 6

$CERTIKCLI1 tx oracle remove-operator $mary --from $mary -y
sleep 6
$CERTIKCLI1 query oracle operators
$CERTIKCLI1 query oracle withdraws

checkConsensus

# shield
