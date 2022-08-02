GOARCH=amd64 GOOS=linux go build -o httptestgo.exe
scp -P 29558 -i ~\bandwagon.pem  -r .\httptestgo.exe root@las.zeddal.com:~/