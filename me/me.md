# 自分の情報

最も簡単に，自分の情報を取得する．`scope`は`User.read`のはず

`Bearer`形式の認証なので，ヘッダーにアクセストークンをつける．これは`MS Graph API`を使うときの基本になる（はず，たぶん）

なお，構造体は[JSON-to-Go](https://mholt.github.io/json-to-go/)を使って生成した．使用したサンプルのユーザー情報が適当だったのでinterface型になっているが，ちゃんとしていればちゃんとしたものが入る？

# エンドポイント

`https://graph.microsoft.com/v1.0/me`