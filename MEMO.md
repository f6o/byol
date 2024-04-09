 # Build your own lisp/micro parser combinator

* https://www.buildyourownlisp.com/
* https://github.com/orangeduck/mpc

## Chapter 6 Parsing

mpcライブラリをつかって、逆ポーランド記法の式をパースする。

### mpcライブラリの使い方

1. mpc_newでパーサを定義する
1. mpca_langで文法を定義する
1. mpc_parseで入力をパースする
1. mpc_ast_printで構文木を出力する
1. mpc_ast_deleteでメモリを解放する
1. mpc_cleanup

see [mpc/exmaple/foobar.c](mpc/example/foobar.c)

### 文法を定義する

```
/* Define them with the following Language */
mpca_lang(MPCA_LANG_DEFAULT,
  "                                                     \
    number   : /-?[0-9]+/ ;                             \
    operator : '+' | '-' | '*' | '/' ;                  \
    expr     : <number> | '(' <operator> <expr>+ ')' ;  \
    lispy    : /^/ <operator> <expr>+ /$/ ;             \
  ",
  Number, Operator, Expr, Lispy);
```

see [mpc/exmaple/lispy.c](mpc/example/lispy.c)

## Chapter 7 Evaluation

この章では逆ポーランド記法の式を評価する。

* number そのまま数値をかえす
* 2つの数値をつかって四則演算をする

### mpc.hでのASTの定義

```
mpc_ast_t {
    tag string
    contents string
    state mpc_state_t
    children_num int
    children mpc_ast_t**
}
```

* tag expr,number,regexなどのノードの型みたいなもの。ルールの定義につかう。
* contents ノードの中身。空のときもある。
* state このノードをパースしたときのパーサの状態。行数と列とか。
* children 子ノード。ポインタで木の構造にする

### ASTの操作

* new
* build
* add_root
* add_child
* add_tag
* add_route_tag
* tag
* state
* delete
* print
* get_index
* get_child

### ASTの走査順

* pre
* post

`mpc_parse`の結果としてASTがある

## Chapter 8 Error Handling

### Lisp Value

* エラーは「式の結果のひとつ」
* Lisp valueとして定義
  * type int (enum)
    * lisp value, lval としてエラー、数字、シンボル、S式
  * num long
  * err int (enum)
* read/eval関数
  * read: AST -> lval
  * eval_op: lval -> lval
  * eval : AST -> lval
* plumbling: つなげてつくるというのもプログラミングの性質。
中身についてくわしく知ることなく、ソフトウェアを作成できる。

## Chapter 9 S式

* ポインタの話: structをコピーするとつらいから、アドレスでやりとりしよう
* スタックとヒープ
* S式の文法
  * number
  * symbol
  * sexpr
  * expr
* lvalをS式むけに変更
  * count int
  * cell lval**
* 式の読み込み: tag/contentsをつかってASTからS式にする
* S式の印字は相互再帰でおこなう
* S式の評価: 子の評価から

## Chapter 10 Q-Expressions

ここで触れたいのは追加機能の考え方。
今回はインタプリタへの追加なので、Syntax, Representation, Parsing, Semanticsにわかれる

* Common Lispだと`'(1 2 3 4)`だけど、ここでは `{1 2 3 4}`という形式
* evalはさわらなくていい
* 組み込み関数としてQ式を扱うものを定義: list, head, tail, join, eval


## Chapter 11 Variables

* Immutableな変数
* name -> value という対応を環境 envrionment として定義
* 変数シンボルの定義
* lvalにlbuiltinという関数ポインタを追加
  シンボルに紐づく関数を定義する
* lvalでlbuiltinを使っている一方で、lbuiltinの定義でもlvalを使っている
* lvalにLVAL_FUNを追加

### Environment

* 環境は lval* のリストと char* のリストを作り、要素の位置で対応付けする
* `def {a b} 3 4` で `a` と `b` に値が対応付けられる

## Chapter 12 Functions

* 関数とは
  * 手続き
  * 入力すると、値を返す
  * partial computations/lambda calculus
* `\\`を関数のシンボルとしてあつかう
* 関数もlvalとしてあつかう

```
typedef lval *(lbuiltin)(lenv *,lval *)

struct lval {
  // ...
  
  builtin lbuiltin
  env  lenv*
  formals  lval*
  body lval*
  
  // ...
}
```

```
(* define *)
\ {x y} {+ x y}

(* apply *)
(\ {x y} {+ x y}) 10 20

(* bind *)
def {f} (\ {x y} {+ x y})
f 10 20
```

* lambda式の導入
* 環境の拡張
  * ユーザ定義関数でもビルトインの関数を使いたい 
  * lenvにparentという環境を定義する
    * defはグローバルな環境、putや=はローカルな環境に変数を定義できることとする

```
struct lenv {
    parent lenv*
    count  int
    syms   []string
    vals   []string
}
```

### Function calling

1. ビルトインなら、そのまま呼び出す
1. そうでないなら、ローカルの環境に、仮引数の名前と値をひもづける
1. parentのEnvを設定する
1. bodyを評価する

これだと不完全？

「引数が少ないときは、部分適用された関数を返す」ようにする

* `{x & xs}` という表現で可変長の引数をとれるようにする
* 関数定義のための手続きはlispy内で作れる
* カリー化

## Chapter 13 Conditionals

* 組み込みとしてconditionを定義: `<, >, <=, >=`
* 等号 `==, !=`

再帰関数が定義できるようになる！

## Chapter 14 Strings

* lvalとしてstring型を追加: LVAL_STRで実態はchar *
* printのときはエスケープをする
* 文法
* 入力となるソースコード上はエスケープ表現なので、readするときはunescapeをする
* コメントの実装: `;` から行末まで
  readするときにはなにもしないが、パーサでは読み込む(ASTには入っている！)

### ファイルロード load

* mpc_parse_contentsを使ってファイルからASTにする
* ファイルは１つの式として扱わず、1つずつ式を評価する
* 1つの式にロードエラーがあったらエラーのlvalにし、エラーを表示して、ロードを続ける

