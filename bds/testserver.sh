#!/bin/bash

wget $(wget -U "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36" https://minecraft.net/en-us/download/server/bedrock/ -O - 2>/dev/null | grep -o 'https://minecraft.azureedge.net/bin-linux/[^"]*') -O update.zip
unzip update.zip ./bds/
echo server start
LD_LIBRARY_PATH=. ./bedrock_server