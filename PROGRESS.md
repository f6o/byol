# PROGRESS.md - 進捗状況

参考書籍: [Build Your Own Lisp](https://www.buildyourownlisp.com/)

## 章ごとの状況

| 章 | 内容 | 読了 | 実装 | 備考 |
|---|---|---|---|---|
| Ch 1-5 | 導入・環境構築・C基礎・REPL・文法概念 | o | o | Go プロジェクト + CGo REPL |
| Ch 6 | パーシング (mpc) | o | 部分 | Tokenizer 完成、Parser は stub |
| Ch 7 | 評価 (Evaluation) | o | 部分 | AST.Eval() で四則演算が動作 |
| Ch 8 | エラーハンドリング | o | 部分 | LVError 型あり (DIV_ZERO, ARGS_COUNT, TYPE_ASSERT) |
| Ch 9 | S式 | o | - | cell/count なし、S式の read/print/eval なし |
| Ch 10 | Q式 | o | - | list/head/tail/join/eval なし |
| Ch 11 | 変数・環境 | o | - | Environment 型なし |
| Ch 12 | 関数・ラムダ | o | - | クロージャ、部分適用なし |
| Ch 13 | 条件分岐 | o | - | 比較演算子なし |
| Ch 14 | 文字列・ファイルロード | o | - | string 型、load なし |
| Ch 15 | 標準ライブラリ | - | - | |
| Ch 16 | ボーナス | - | - | |

## 全体の進捗

- 読了: Ch 14 まで (16章中14章、約87%)
- 実装: Ch 6-8 の一部 (コア実装9章中 約2章分、約20-25%)

## 実装済みの機能

- **Tokenizer** (`lispy/tokenizer.go`): bufio.Reader ベースのストリーム型トークナイザ
  - TOKEN_NUMBER, TOKEN_OPERATOR, TOKEN_LPAREN, TOKEN_RPAREN, TOKEN_SYMBOL, TOKEN_EOF
  - イテレータパターン (NewTokenizer + NextToken)
  - 一括変換 (Tokenize, TokenizeReader)
- **AST** (`lispy/lispy.go`): Tag, Contents, Children による木構造
- **評価器** (`lispy/lispy.go`): 四則演算 (+, -, *, /) の再帰評価
- **エラー型** (`lispy/lispy.go`): LVError (ERR_DIV_ZERO, ERR_ARGS_COUNT, ERR_TYPE_ASSERT)
- **CGo REPL** (`main.go`): editline + mpc を使ったプロトタイプ

## 次のマイルストーン

1. **Parser の実装** (`lispy/parser.go`): トークン列から AST を構築し、Tokenizer → Parser → Eval のパイプラインを完成させる
2. **S式の導入** (Ch 9): LispValue に S式を追加し、read/print/eval を実装
3. **Q式と組み込み関数** (Ch 10): list, head, tail, join, eval
