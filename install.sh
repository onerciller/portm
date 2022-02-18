#!/bin/bash
set -e

version="1.0.1"

while [ -n "$1" ]
do
  case "$1" in
    -v)
        version=$2
        ;;
  esac
  shift
done

goos="linux"
goarch="amd64"

if [ `uname -s` == "Darwin" ];then
	goos="darwin"
fi

if [[ `arch` =~ "aarch64" ]];then
	goarch="arm64"
fi

filename="portm_"$version"_"$goos"_"$goarch".tar.gz"

rm -f $filename

curl -LJO "https://github.com/onerciller/portm/releases/download/v"$version"/"$filename""

tar -xvf $filename

chmod +x ./portm

mv ./portm /usr/local/bin

