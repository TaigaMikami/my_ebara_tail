# mytail
-  Gopher道場 申し込み課題

次の仕様を満たすtailコマンドに似たコマンドをGoで作成してください。
- 引数で渡された1つのファイルの末尾の最大N行をそのまま出力
- Nはデフォルトで10
- オプション-nでNを指定できる

## 追加機能
1. 複数ファイル対応
2. headコマンドになるオプション追加(-o=f, 最終行からではなく最初の行からn行出力する)
3. ファイルの行をランダムにn行出力するオプション追加(-o=r)
4. エバラおかずランキングtail機能
エバラの素を使用したおかずのレシピの人気順表示をスクレイピングし、おかずのランキングを100位から表示するオプション
https://www.ebarafoods.com/recipe/cla_menu/49/?&limit=100

## 実行例

```
❱ ./mytail -n=2 txts/main_dish.txt txts/soup.txt txts/staple_food.txt        
==> main_dish.txt <==
さばのみそに
すぶた

==> soup.txt <==
わかめスープ
たまごスープ

==> staple_food.txt <==
パスタ
ラーメン

```

## オプション
-h : ヘルプ
```
❱ ./mytail -h

Usage of ./mytail:
   ./mytail [OPTIONS] ARGS...
Options\n  -n int
        How many lines to read from the end (default 10)
  -o string
        Read in reverse option, b: back(default), f: front, r: random shuffle (default "b")
  -okazu
        I'm hungry! Yes, let's decide the side dish!
```

- n : 行数指定
```
❱ ./mytail -n=1 -o=b  txts/hoge.txt
==> hoge.txt <==
z


❱ ./mytail -n=3 hoge.txt                              
==> hoge.txt <==
x
y
z
```

-o : b: 最終行から, f: 最初行から, r: ランダム
```
❱ ./mytail -n=1 -o=b  txts/hoge.txt
==> hoge.txt <==
z


❱ ./mytail -n=1 -o=f  txts/hoge.txt
==> hoge.txt <==
x

❱ ./mytail -n=2 -o=r  txts/hoge.txt
==> hoge.txt <==
v
z

```

-okazu : true or false (default: false)
今日のおかずに悩んだ際におかずをtailコマンドで表示します。
前述した -o=f や -o=r コマンドを用いることで楽しくカスタマイズできます。

tailコマンドなので最下位から出力されます。
```
❱ ./mytail -okazu=true -n=3     
98位:魚介たっぷりプルコギ風: https://www.ebarafoods.com/recipe/detail/recipe3057.php
99位:牛肉の長芋磯辺巻き: https://www.ebarafoods.com/recipe/detail/recipe3054.php
100位:牛肉のたたき: https://www.ebarafoods.com/recipe/detail/recipe3053.php
```

-o=f をつけることで、人気ランキングになります！
```
❱ ./mytail -okazu=true -o=f -n=3
1位:肉巻きレタス: https://www.ebarafoods.com/recipe/detail/recipe2801.php
2位:プチッと豚キムチ炒め: https://www.ebarafoods.com/recipe/detail/recipe1559.php
3位:なすのボリューム味噌炒め: https://www.ebarafoods.com/recipe/detail/recipe1192.php
```

## 各ソースコード説明
- mytail.go
tailコマンド実装コード

- mytail_test.go
mytail.goのテストコード

- ebrafood_scraping.go
エバラの素を使用したおかずのレシピの人気順表示をスクレイピングするコード
