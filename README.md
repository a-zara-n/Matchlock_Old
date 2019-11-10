# Matchlock
```
   __  ___     __      __   __         __     ====
  /  |/  /__ _/ /_____/ /  / /__  ____/ /__  ====   Matchlock
 / /|_/ / _ '/ __/ __/ _ \/ / _ \/ __/  '_/ ====	〔 https://github.com/a-zara-n/Matchlock 〕
/_/  /_/\_,_/\__/\__/_//_/_/\___/\__/_/\_\ ====
=============================================
```
## About
このツールはWebアプリケーションへの動的なセキュリティテストツールです。

**このような用途に対応したい**
- 社内のWebアプリケーション開発時のセキュリティテスト

**このツールに備わっている機能**
- プロキシ機能
- 簡単なXSS検知機能
- HTTP改変機能
- historyの保存
- クローリング機能

## Description
このツールを使う際は、お好きなChromeかFirefox(こっちがオススメ)とSQLiteをインストールしておいてください。
## Install & Start UP
**Install**
```sh
# When putting under GOPATH
mkdir $GOPATH/src/github.com/a-zara-n/
cd $GOPATH/src/github.com/a-zara-n/
#  If you put in that directory from here
git clone https://github.com/a-zara-n/Matchlock.git
go mod download
```
**Start UP by Server Run**
```sh
go run .
# or
go build .
./Matchlock
   __  ___     __      __   __         __     ====
  /  |/  /__ _/ /_____/ /  / /__  ____/ /__  ====   Matchlock
 / /|_/ / _ '/ __/ __/ _ \/ / _ \/ __/  '_/ ====	〔 https://github.com/a-zara-n/Matchlock 〕
/_/  /_/\_,_/\__/\__/_//_/_/\___/\__/_/\_\ ====
=============================================

⇨ http server started on [::]:8888
```
**Start UP by Web UI**

次のようなURLを利用してサイトにアクセスしてください
```
http(s)?://<hostname>:8888
example : http://localhost:8888
```
