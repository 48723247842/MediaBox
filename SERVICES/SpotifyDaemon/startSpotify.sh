#!/bin/bash
# https://unix.stackexchange.com/a/416031
DISPLAY=""
# Guess the active DISPLAY
while read session; do
    # Explode to needed variables
    set -- $session; tty=$2; display=$3
    # If there is an X session in thet TTY
    if [ "$display" != "-" ]; then
        # 1st non root display is used if TTY is not matched with the active
        [ "$DISPLAY" == "" ] && DISPLAY="$display"
        # If it is the active TTY we can stop, this is the active X!
        [ "$tty" == "$(cat /sys/class/tty/tty0/active)" ] && DISPLAY="$display";
    fi
done <<< "$(w -hs)"
t="1"
if [[ "$t" == "1" ]]; then
	set -x
	set -e
	#export DISPLAY=:0;
	export "DISPLAY=$DISPLAY"
	#nohup /usr/bin/spotify
	/usr/bin/spotify
	#/usr/bin/spotify
fi