# はじめに
Go言語を勉強する中で簡易的なcatコマンドを作成したのでまとめます。
Go言語に慣れるための練習として書いているため、エラー処理などは思いついたまま書いています。よりGo言語らしい(？)書き方があると思いますが、今回は気にせず書いています。
簡易版ということで、オプションは`[-n]`と`[-b]`の２つだけを実装しています。

# コード
今回はオプション機能に関するコードを main.go とは別のファイルに書いています。
よって、goのファイルは `main.go` と `cat_option.go` の２つになっています。

# 簡単にコードの説明
今回はコマンド引数を扱うために flag パッケージを使用しています。

オプションとして `-n` と `-b` を実装しており、それぞれ以下のような機能となっています。
`-n：行番号を付けて表示`
`-b：空行以外に行番号を付けて表示`

例として`-n`が引数として与えられた時、 変数 n は true となって下の if 文によって関数 `catN` が呼び出されます。また、オプションとして何も指定されず、引数にファイル名(パス名)だけが渡された場合はファイルの中身をそのまま表示します。
```
	if *n {
		catN(scanner)
	} else if *b {
		catB(scanner)
	} else {
		cat(scanner)
	}
```

# 実行してみる
以下のサンプルを使って実行してみます。
```sample.text
apple
banana

orange, orange

mango
```

## ビルド
`＄go build -o $GOBIN/cat_command`
※ビルドに関する説明は省略します。

## 実行
実行する際には以下の２通りで実行します。
（今回は対象のtxtファイルなどがカレントディレクトリにある場合を想定しています。パス名を指定することで他のディレクトリにあるファイルも表示することができます。）

オプションを指定しない場合：`＄cat_command [ファイル名]`
オプションを指定する場合：`＄cat_command [オプション] [ファイル名]`

それでは、いくつかの方法で実行してみます。
`＄cat_command sample.txt`
```
$cat_command sample.txt
apple
banana

orange, orange

mango
```

`＄cat_command -b sample.txt`
オプションを指定し、空行以外に行番号をつけて表示します。
```
$cat_command -b sample.txt
    1: apple
    2: banana

    3: orange, orange

    4: mango
```
`＄cat_command -a sample.txt`
存在しないオプションを指定した場合、実行に失敗しヘルプが表示されます。
```
$cat_command -a sample.txt
flag provided but not defined: -a
Usage of cat_command:
  -b    空行以外に行番号を付けて表示する
  -n    行番号を付けて表示する
```
