# アクセストークン

ここでは，`reflesh token`からアクセストークンを発行してみる．MS Graphでは，アクセストークンの`expire time`はデフォルトで60分らしい（間違ってるかもしれない）

ここまでで，`tenant id`, `client id`, `client secret`そして`reflesh token`を取得済とする．サンプルでは，それらは環境変数から読んで(`os.getEnv`)いる．

アクセストークンは`jwt`形式であり，複合するといろいろな情報が見れる．

[https://jwt.ms](https://jwt.ms/)

# エンドポイント

`https://login.microsoftonline.com/{{TenantID}}/oauth2/v2.0/token`