# Quicker

![icon](https://user-images.githubusercontent.com/56575610/80867069-1a8f3a80-8ccd-11ea-9da8-d546d4d8c8f6.png)

短文送信サイト『[Quicker](https://quicker.netlify.app/)』のCUIアプリケーション

アカウント登録などを一切せずに簡単にCUI同士で文字列を送ることができます。

## インストール方法

```bash
$ go get github.com/moisutsu/quicker
```

`$GOPATH/bin`にパスが通すと、`quicker`で実行することができます。`$GOPATH`を特に設定していなければ、`$GOPATH=$HOME/go`になります。

## 使い方

文字列を送る、文字列を受け取るの2つを行うことができます。

### 文字列を送る

```bash
$ quicker -m <送りたい文字列>
```

こうすることでこの文字列と結びついた6桁の数字が表示されます。この数字は文字列を受け取る側で必要となります。

### 文字列を受け取る

```bash
$ quicker -i <6桁の数字>
```

文字列を送信したときに表示された6桁の数字を入力することで文字列を受け取ることができます。

## 最後に

このように6桁の数字を共有するだけで、URLなどの文字列を簡単に送受信することができます。
送信した文字列は10分たつと消去されます。

Webサイト『[Quicker](https://quicker.netlify.app/)』とサーバーは共通のため、CUIからWebサイトに文字列を送信することができます。

サーバーはGCPの無料枠なのでいじめないでください。
