#!/bin/bash

rm -rf ~/.certikd
ln -s ~/node0/certikd  ~/.certikd

### public keys
#echo 'export NODE0_KEY=$(certikcli keys show node0 -a --home /root/node0/certikd)' >> ~/.bashrc
#echo 'export NODE1_KEY=$(certikcli keys show node1 -a --home /root/node1/certikd)' >> ~/.bashrc
#echo 'export NODE2_KEY=$(certikcli keys show node2 -a --home /root/node2/certikd)' >> ~/.bashrc
#echo 'export NODE3_KEY=$(certikcli keys show node3 -a --home /root/node3/certikd)' >> ~/.bashrc

### private keys added to keychain
#echo -e "$(cat /root/node0/certikcli/key_seed.json | sed -r 's/^([^\"]+\"+){3}((\"*[^\"]+)).*/\2/')" "\n" | certikcli keys add --recover node0
#echo -e "$(cat /root/node1/certikcli/key_seed.json | sed -r 's/^([^\"]+\"+){3}((\"*[^\"]+)).*/\2/')" "\n" | certikcli keys add --recover node1
#echo -e "$(cat /root/node2/certikcli/key_seed.json | sed -r 's/^([^\"]+\"+){3}((\"*[^\"]+)).*/\2/')" "\n" | certikcli keys add --recover node2
#echo -e "$(cat /root/node3/certikcli/key_seed.json | sed -r 's/^([^\"]+\"+){3}((\"*[^\"]+)).*/\2/')" "\n" | certikcli keys add --recover node3