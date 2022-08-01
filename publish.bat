set GOARCH=amd64
set GOOS=linux
go build -o httptestgo.exe
scp -P 29558 -i ~\bandwagon.pem  -r .\httptestgo.exe root@las.zeddal.com:~/
