#!/bin/bash
set -e
trap "exit" INT

partition_name=''

while getopts i: flag
do
    case "${flag}" in
        i) partition_name=${OPTARG};;
    esac
done

if [[ -z $partition_name ]]; then
    echo "Please input partition_name with -i opt"
    exit 0
fi

is_lusk="$(sudo blkid -o value -s TYPE $partition_name)"

if [ $is_lusk == "crypto_LUKS" ]; then
    echo "encrypted"
else
    echo "unencrypted"
fi
