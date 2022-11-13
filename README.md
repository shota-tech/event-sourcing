# イベントソーシング実装のデモ

## 概要
在庫管理のCLIアプリケーション。
イベントソーシングで実装。

## 使い方
```bash
# 1. ターミナルからアプリケーションを起動する。
go run ./cmd/

# 2. 実行したい操作を選択し、コマンドを入力する。
R: Receive Inventory
S: Ship Inventory
A: Inventory Adjustment
Q: Quantity On Hand
E: Events
Q: Quit
> r

# 3. 必要なパラメータを入力する。
> SKU: abc123
> Quantity: 30

# 4. 実行結果が表示される。
abc123 Received: 30

# 5. 2に戻る。
```
