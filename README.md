# ap1_Resns_backend
Resnsのバックエンド

#### api1
アカウント作成、認証api

#### api2
ユーザー情報の登録更新api

#### api3
記事のデータ送信のapi

#### api4
記事に対するコメント送信、更新のapi

#### api5
記事に対するいいね数送信、更新のapi

#### api6
リストに記事の追加、削除のapi

<<<<<<< HEAD
#### DB
性別 0:男 1:女  

### データベースの接続情報を設定する
環境変数にデータベースの接続情報を設定します。
ターミナルのセッション毎に設定したり、.bash_profileで設定を行います。

- Macの場合
```cassandraql
$ export MYSQL_USER=root \
    MYSQL_PASSWORD=rootpassword \
    MYSQL_HOST=127.0.0.1 \
    MYSQL_PORT=3306 \
    MYSQL_DATABASE=resns_app
```

- Windowsの場合
```cassandraql
$ SET MYSQL_USER=root
$ SET MYSQL_PASSWORD=rootpassword
$ SET MYSQL_HOST=192.168.99.100
$ SET MYSQL_PORT=3306
$ SET MYSQL_DATABASE=resns_app
```

#### DBの補足
ユーザーのプロフィール  
性別 未設定:0 男:1 女:2

記事のジャンルは数字で管理  
例)  ドラマ＝1 スポーツ=2 音楽=3 グルメ=4 アニメ・漫画=5  
ゲーム=6 エンタメ=7 ファッション=8 アニマル=9   
国際=10 政治=11 その他=12
