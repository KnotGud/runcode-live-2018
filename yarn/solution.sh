#!/bin/bash

a=$(strings $1 | grep -e "RCN{.*}")
echo $a