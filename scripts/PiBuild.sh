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
 

stop|reload|restart|build2pi)
#obviously need a way to enter the proper pi
  echo Building for the Pi
		env GOOS=linux GOARCH=arm GOARM=5 go build
		scp * jalspach@10.0.0.110:/home/jalspach/Public

	;;
  *)
	echo "Usage: $N {build|rebuild}" >&2
	exit 1
	;;
esac

exit 0