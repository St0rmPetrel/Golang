#!/bin/bash

if [[ "$1" == "" ]] || [[ $1 -lt 1 ]]; then
	echo "Bad input"
	exit 1
fi

i=1;
while [ $i -le $1 ]
do
	# Print num [-32767, 32767]
	echo $(($RANDOM - $RANDOM))
	i=$((i + 1));
done
