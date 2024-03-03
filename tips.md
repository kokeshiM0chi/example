
## memo

以下全てのファイルをformat
```shell
go fmt ./...
```

## iden3

circom系は動作実績がある。polygonとか。

snarkjs: javascriptでpairingまでやってる
rapidsnark-go: pairingはcまかせ。proof作成にこれ呼び出してた。https://github.com/iden3/rapidsnark/blob/main/src/prover.h#L45
呼び出し先のcを見るとgoでラップしてるだけ。proof作成だけcliがある。cliは粗末。
rapidsnark: rapidsnarkは実行ファイルに引数渡して使う。rapidsnark-goと同じ。

### Q

trusted setupはなんかsnarkjsでやったんんだけど？？
置き換えろってrapidsnarkのreadmeに書いてあり？
todo: trusted setupをgo-rapidsnark/rapidsnarkでできないのかを調べる

## consensys

gnark
とにかく速いけど動作実績が乏しい。楕円曲線のアルゴリズムの最適化に強い人が一人いてその人がアルゴリズムを改良していってる。


##　benchmarkの項目

trusted setup
witness計算
proof生成
verify speed


