#! /bin/bash
#! /bin/bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags '-w -s'
mkdir daily-windows
mv daily.exe daily-windows/
cp config.json daily-windows/config.json
cp README.md daily-windows/README.md
tar -zcvf daily-windows.tar.gz daily-windows
rm -rf daily-windows
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-w -s'
mkdir daily-mac
mv daily daily-mac/
cp config.json daily-mac/config.json
cp README.md daily-mac/README.md
tar -zcvf daily-mac.tar.gz daily-mac
rm -rf daily-mac
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s'
mkdir daily-linux
mv daily daily-linux/
cp config.json daily-linux/config.json
cp README.md daily-linux/README.md
tar -zcvf daily-linux.tar.gz daily-linux
rm -rf daily-linux
echo 'success'