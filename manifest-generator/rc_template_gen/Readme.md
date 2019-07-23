Run the script with the following args-

`go run rc_temp.go -APP_NAME=app1 -ENV=dev -TYPE=type1 -VOLUMES=volume -APP_GROUP=group1 -CONTAINER_PORT=23 -BUILD_NUMBER=4 -REPLICA_COUNT=4 -VOLUME_MOUNTS=vol_mount -ARGS=arg -MEMORY_REQUESTS=520g -CPU_REQUESTS=40g -MEMORY_LIMITS=23g -CPU_LIMITS=22g`


File generated in the current directory-

`RC-{app_name}-{build_number}.yml`