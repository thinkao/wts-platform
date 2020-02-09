#! /bin/bash

echo "--------------create volumes directory------------"
mkdir -p ./mysql/db
mkdir -p ./mysql/conf

echo "--------------load docker images------------"
docker load -i wts_server.tar.gz

echo "-------------docker-compose up--------------"
docker-compose up -d
echo "------------docker-compose done------------"

# figlet "WTSplatform"
cat <<-'EOF'

__        _______ ____        _       _    __
\ \      / /_   _/ ___| _ __ | | __ _| |_ / _| ___  _ __ _ __ ___
 \ \ /\ / /  | | \___ \| '_ \| |/ _` | __| |_ / _ \| '__| '_ ` _ \
  \ V  V /   | |  ___) | |_) | | (_| | |_|  _| (_) | |  | | | | | |
   \_/\_/    |_| |____/| .__/|_|\__,_|\__|_|  \___/|_|  |_| |_| |_|
                       |_|

Server running...
EOF