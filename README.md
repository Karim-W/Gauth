## Client Credentials Workflow


 ##### REST Request 

>POST @ /token endpoint

|param|value|
|------|------|
|headers Auth| base64 encoded (clientid:clientsecret)|
|body|grant_type=client_creds&scope=...|

#####  _Auth Server Response_

```
{
	accessToken:"w/e"
	tokenType:"who cares"
	expiresIn:"tmr"
	scope:"read_Users"
}
````

```mermaid 
  sequenceDiagram
	App->>AuthServer: authenticate via clientId and ClientSecret
	AuthServer->>AuthServer: Validate Client Id and secret
	AuthServer->>App: Send AccessToken
	App->>Api: Make Request using AccessToken
	Api->>App: send Response
	
```

## Auth Code Workflow
```mermaid
  sequenceDiagram
  User->>App: Click login link
  App->>AuthServer: Send Auth code request to /authorize
  AuthServer->>User: redirect to login/authorization promt
  User->>AuthServer: Authenticate and consent
  AuthServer->>App: send Authorixation code
  App->>AuthServer: AuthCode+ClientId+ClientSecret to /oauth/token
  AuthServer->>AuthServer: validate AuthCode+ClientId+ClientSecret
  AuthServer->>App: send tokenId and Access Token
  App->>API: Request userData using access token on behalf od user
  API->>App: send Response
```