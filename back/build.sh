echo "Building app $app"
cd src || exit 1
go get -d -v || exit 1
go build -o main main.go  || exit 1
echo "FINISHED Building app $app"