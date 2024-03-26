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

Common Lispだと`'(1 2 3 4)`だけど、ここでは
`{1 2 3 4}`という形式

Chapter9まででS式のREPLができる
Chapter11では環境を導入して、変数をつかえるようにした

## Chapter 12 Functions

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

```
struct lenv {
    parent lenv*
    count  int
    syms   []string
    vals   []string
}
```

Function callingの手前まで。


