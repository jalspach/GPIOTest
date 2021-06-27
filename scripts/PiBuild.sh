#!/bin/sh
### Script to build (or rebuild) for the pi
set -e

case "$1" in
  build)
	# Straight Build
    echo Building for the Pi
		env GOOS=linux GOARCH=arm GOARM=5 go build
	;;
  rebuild)
	#Remove old files
		echo Removing old stuff goes here
	# Straight Build
    echo Building for the pi
    env GOOS=linux GOARCH=arm GOARM=5 go build
;;
 

stop|reload|restart|force-reload)
	;;
  *)
	echo "Usage: $N {build|rebuild}" >&2
	exit 1
	;;
esac

exit 0