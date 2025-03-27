# ユーザ主体認証による事業者認証
A社がユーザ認証システムで事業者認証をした後に、自身の事業者情報を取得する例を示します。

## 1. 事業者認証の実行
Action (A社): 下記の```curl```コマンドを実行し、認証情報を取得します。

前提条件として、プラットフォーム認定を受けた事業者に発行されるApiKeyを指定しAPIを実行します。
また運営事業者から各事業者はAccountIdとPasswordを事前に払い出されています。

```
curl --location --request POST 'http://localhost:8081/auth/login' \
--header 'Content-Type: application/json' \
--header 'apiKey: Sample-APIKey1' \
--data-raw '{
  "operatorAccountId": "oem_a@example.com",
  "accountPassword": "oemA&user_01"
}'
```

```json
{
    "accessToken":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJ2WVMyY3QwUVFqRGFWcEpWUktKTXRhRXFPWUFaQ0plNkN1RVZuQVJYUW44In0.eyJleHAiOjE3NDIxOTMyNjksImlhdCI6MTc0MjE5MjM2OSwianRpIjoiNTJmOWRhODktNzY0Ny00YTk3LWIwNDgtMDFhMGFmNGJiNWZhIiwiaXNzIjoiaHR0cDovL2hvc3QuZG9ja2VyLmludGVybmFsOjQwMDAvcmVhbG1zL21hc3RlciIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiI3ZTIyZjM3Ny1hMzExLTQ0ZjMtOTUzNC1hZjg5OWNlM2MzNDYiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJVc2VyQXV0aGVudGljYXRpb25TeXN0ZW0iLCJzaWQiOiI4NDA3ZWVhOS1kZDZiLTQ3MTktODBiOS02ZmQ1ZTRiY2NjZGMiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtbWFzdGVyIiwib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIGVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm9wZXJhdG9yX2lkIjoiYjM5ZTYyNDgtYzg4OC01NmNhLWQ5ZDAtODlkZTFiMWFkYzhlIiwicHJlZmVycmVkX3VzZXJuYW1lIjoib2VtX2FAZXhhbXBsZS5jb20ifQ.KR0v_CSWU54E1ZA8qPV2YHsCxsgpXSlMfhy1QTLSJNBRUkfOMJ5_Fspd58IPyZgW6_8zS1XmD2jpf31PuSJoLBFMqSnAJ_9Z0sVTKlO43kDsugsbThICgfU8En4kdD2b0pJrtIzC91ntisw-WPnZSEHFsh7ZSobXxJze_BoCOW78AxyaTzAOT7J-t-RmGjiluhyEaMVLks9JdKniakzh7nsdtgA2ypa5RRh1k8owCNBjBqT9xJWizP1905AvMdGd3mEXh95HKhyJrYcHI1B8rqIl1--RwcQYZiN-_1P1shr--x1ZDIkdIvDdzUW47LaDqKSRT5KHaNPF-Pd6fdQxWQ","refreshToken":"eyJhbGciOiJIUzUxMiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJhNDliMDRkMy1mYmEwLTQ0OTAtOWQzNS1mNTA0NjM4MzQ5NDkifQ.eyJleHAiOjE3NDIxOTQxNjksImlhdCI6MTc0MjE5MjM2OSwianRpIjoiZTMyNmNmMzItZjgxMy00NThmLTk3NTEtNzI2YzlhZDU2ZTBlIiwiaXNzIjoiaHR0cDovL2hvc3QuZG9ja2VyLmludGVybmFsOjQwMDAvcmVhbG1zL21hc3RlciIsImF1ZCI6Imh0dHA6Ly9ob3N0LmRvY2tlci5pbnRlcm5hbDo0MDAwL3JlYWxtcy9tYXN0ZXIiLCJzdWIiOiI3ZTIyZjM3Ny1hMzExLTQ0ZjMtOTUzNC1hZjg5OWNlM2MzNDYiLCJ0eXAiOiJSZWZyZXNoIiwiYXpwIjoiVXNlckF1dGhlbnRpY2F0aW9uU3lzdGVtIiwic2lkIjoiODQwN2VlYTktZGQ2Yi00NzE5LTgwYjktNmZkNWU0YmNjY2RjIiwic2NvcGUiOiJvcGVuaWQgcm9sZXMgZW1haWwgd2ViLW9yaWdpbnMgYWNyIHByb2ZpbGUgYmFzaWMifQ.uCWgBTmrQAuTAlh-hxnEY9WgGmZPGfFOMgS_idX1Qgc_kBNSLzkuwtpgyXhWZqH41Kxwxl1Mbsd0n9BUiJFQ-Q"
}
```

認証が成功すると返却値としてjson web tokenが払い出されます。
ユーザ認証システムにはheaderにApiKeyおよびTokenを必ず指定してAPIを実行します。

## 2. 事業者情報の取得
Action (A社): 下記のcurlコマンドを実行し、自社の事業者情報を取得します。
```
curl --location --request GET 'http://localhost:8081/api/v2/authInfo/operator' \
--header 'apiKey: Sample-APIKey1' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJ2WVMyY3QwUVFqRGFWcEpWUktKTXRhRXFPWUFaQ0plNkN1RVZuQVJYUW44In0.eyJleHAiOjE3NDIxOTMyNjksImlhdCI6MTc0MjE5MjM2OSwianRpIjoiNTJmOWRhODktNzY0Ny00YTk3LWIwNDgtMDFhMGFmNGJiNWZhIiwiaXNzIjoiaHR0cDovL2hvc3QuZG9ja2VyLmludGVybmFsOjQwMDAvcmVhbG1zL21hc3RlciIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiI3ZTIyZjM3Ny1hMzExLTQ0ZjMtOTUzNC1hZjg5OWNlM2MzNDYiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJVc2VyQXV0aGVudGljYXRpb25TeXN0ZW0iLCJzaWQiOiI4NDA3ZWVhOS1kZDZiLTQ3MTktODBiOS02ZmQ1ZTRiY2NjZGMiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtbWFzdGVyIiwib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIGVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm9wZXJhdG9yX2lkIjoiYjM5ZTYyNDgtYzg4OC01NmNhLWQ5ZDAtODlkZTFiMWFkYzhlIiwicHJlZmVycmVkX3VzZXJuYW1lIjoib2VtX2FAZXhhbXBsZS5jb20ifQ.KR0v_CSWU54E1ZA8qPV2YHsCxsgpXSlMfhy1QTLSJNBRUkfOMJ5_Fspd58IPyZgW6_8zS1XmD2jpf31PuSJoLBFMqSnAJ_9Z0sVTKlO43kDsugsbThICgfU8En4kdD2b0pJrtIzC91ntisw-WPnZSEHFsh7ZSobXxJze_BoCOW78AxyaTzAOT7J-t-RmGjiluhyEaMVLks9JdKniakzh7nsdtgA2ypa5RRh1k8owCNBjBqT9xJWizP1905AvMdGd3mEXh95HKhyJrYcHI1B8rqIl1--RwcQYZiN-_1P1shr--x1ZDIkdIvDdzUW47LaDqKSRT5KHaNPF-Pd6fdQxWQ'

```

A社の事業者情報が返却されます。
```json
[
    {
        "dataModelType":"test1",
        "attribute":
            {
                "operatorId":"b39e6248-c888-56ca-d9d0-89de1b1adc8e",
                "operatorName":"A社",
                "operatorAddress":"東京都渋谷区xx",
                "openOperatorId":"1234567890123",
                "operatorAttribute":
                {
                    "globalOperatorId":"1234ABCD5678EFGH0123"
                }
        }
    }
]
```

## 3. トークンイントロスペクションの実行
Action (A社): 下記のcurlコマンドを実行し、自社の事業者識別子(内部)を取得します。
```
curl --location --request POST 'http://localhost:8081/api/v2/systemAuth/token' \
--header 'Content-Type: application/json' \
--header 'apiKey: Sample-APIKey1' \
--data-raw '{
  "idToken": "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJ2WVMyY3QwUVFqRGFWcEpWUktKTXRhRXFPWUFaQ0plNkN1RVZuQVJYUW44In0.eyJleHAiOjE3NDIxOTMyNjksImlhdCI6MTc0MjE5MjM2OSwianRpIjoiNTJmOWRhODktNzY0Ny00YTk3LWIwNDgtMDFhMGFmNGJiNWZhIiwiaXNzIjoiaHR0cDovL2hvc3QuZG9ja2VyLmludGVybmFsOjQwMDAvcmVhbG1zL21hc3RlciIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiI3ZTIyZjM3Ny1hMzExLTQ0ZjMtOTUzNC1hZjg5OWNlM2MzNDYiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJVc2VyQXV0aGVudGljYXRpb25TeXN0ZW0iLCJzaWQiOiI4NDA3ZWVhOS1kZDZiLTQ3MTktODBiOS02ZmQ1ZTRiY2NjZGMiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtbWFzdGVyIiwib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIGVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm9wZXJhdG9yX2lkIjoiYjM5ZTYyNDgtYzg4OC01NmNhLWQ5ZDAtODlkZTFiMWFkYzhlIiwicHJlZmVycmVkX3VzZXJuYW1lIjoib2VtX2FAZXhhbXBsZS5jb20ifQ.KR0v_CSWU54E1ZA8qPV2YHsCxsgpXSlMfhy1QTLSJNBRUkfOMJ5_Fspd58IPyZgW6_8zS1XmD2jpf31PuSJoLBFMqSnAJ_9Z0sVTKlO43kDsugsbThICgfU8En4kdD2b0pJrtIzC91ntisw-WPnZSEHFsh7ZSobXxJze_BoCOW78AxyaTzAOT7J-t-RmGjiluhyEaMVLks9JdKniakzh7nsdtgA2ypa5RRh1k8owCNBjBqT9xJWizP1905AvMdGd3mEXh95HKhyJrYcHI1B8rqIl1--RwcQYZiN-_1P1shr--x1ZDIkdIvDdzUW47LaDqKSRT5KHaNPF-Pd6fdQxWQ"
}'
```

A社の事業者識別子(内部)が返却されます。
```json
{
    "operatorId": "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
    "active": true
}
```