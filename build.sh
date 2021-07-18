build_fn() {
  version="$(go version|awk '{print $3}'|sed 's/go//g')"
  if [ -z "$version" ]
    then
      echo "Please install go! Version need up than 1.14.1"
      exit 1
    else
      echo "start build..."
      rm -rf ./output/cmd
      rm -rf ./output/conf
      export GOPROXY="https://goproxy.cn,direct"
      go mod download
      mkdir -p output
      mkdir -p output/conf
      cp -r ./conf/* output/conf
      cd output || exit 1
      go build ../cmd
  fi
}

if [ "$1" = "build" ]
  then build_fn
elif [ "$1" = "run" ]
  then
    build_fn
    ./cmd &
else
  echo "command is not fit"
fi