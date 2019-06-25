# Matchlock
```
   __  ___     __      __   __         __     ====
  /  |/  /__ _/ /_____/ /  / /__  ____/ /__  ====   Matchlock
 / /|_/ / _ '/ __/ __/ _ \/ / _ \/ __/  '_/ ====	〔 https://github.com/a-zara-n/Matchlock 〕
/_/  /_/\_,_/\__/\__/_//_/_/\___/\__/_/\_\ ====
=============================================
```
## About
このツールはWebアプリケーションへのセキュリティテストツールです。
主にDevSecOpsでの脆弱性検査自動化（scaning)を目的にしており、使用する上で教育コストや組み込みコストが安価で利用者が扱いやすいツールを目指しています。

**このツールに備わっている機能**
- プロキシ機能
- 簡単なXSS検知機能
- HTTP改変機能
- historyの保存
- クローリング機能

**今後追加/更新する予定の機能**
- [ ] シナリオ機能
- [ ] SQLiなどの脆弱性検知機能と性能向上
- [ ] クローリンクデータを元にしたシナリオ自動生成機能
- [ ] CIやツールに組み込みやすい形でのAPI機能
- [ ] UIの更新
- [ ] クローリングと同時に画面記録とパラメタの取得
- [ ] 社内診断士向けチェックシートの吐き出し
- [ ] 脆弱性のレポーティング機能
- [ ] SPAへ対応するためのbrowser extension開発(衛星ツールの作成)
- [ ] 統計機能の追加
- [ ] 各種データの吐き出し
- [ ] openAPIを用いたAPIドキュメント作成
- [ ] パッケージ分割
- [ ] 僕が欲しいペネトレ用機能
- [ ] Webhook機能
- [ ] Google Driveを利用した機能の追加(保存はき出し)
- [ ] ペネトレーションツールとの連携機能(が個人的に欲しい)
- [ ] Docker化/k8s対応
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
