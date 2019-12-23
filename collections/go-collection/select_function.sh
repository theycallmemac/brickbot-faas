#!/bin/bash

# positional args
  args=()

  # named args
  while [ "$1" != "" ]; do
      case "$1" in
          -f | --function )   function="$2";     shift;;
          -a | --argument )   argument="$2";      shift;;
      esac
      shift # move to next kv pair
  done

if [ -z /go/${function} ];then
    exit 1
else
    if [ ! -f /go/${function} ];then
        echo "Function does not exist : ${function}. Exiting"
        exit 1
    fi
fi

/go/${function} "${argument}"

exit $?
