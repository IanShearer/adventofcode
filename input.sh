#!/bin/zsh

./grab_input.py "$@"

DAY='01'
YEAR='2020'

if [ -z $1 ];then
   YEAR=$(date '+%Y');
   DAY=$(date '+%-d');
else
    echo "$1";
fi

cd "$YEAR/day$(printf %02d $DAY)";
$SHELL