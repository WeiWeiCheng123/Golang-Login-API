# Golang-Login-API
Use Gin and SQL to create a simple login API

## 執行方法
```sh
git clone https://github.com/WeiWeiCheng123/Golang-Login-system.git

cd Golang-Login-system

./start-project.sh
```

## API執行方法
這個專案有兩個API分別是
- Sign up
- Sign in

要注意本 API 有設定帳號長度要大於等於 6 以及至少要一個英文字

Sign up
Request
```sh
# username and password 可以自行輸入不同字串
curl -X POST -H "Content-Type:application/json" -d '{"username":"$username","passwd":"$password"}' http://localhost:8080/api/v1/signup
```
Response
```sh
# if the username does not exist and passes  the rule
User $username created

# if the username does not pass the rule
must have at least one character and six words

#if the username exists
the username is already used
```

Sign in
Request
```sh
# username and password 可以自行輸入不同字串
curl -X POST -H "Content-Type:application/json" -d '{"username":"$username","passwd":"$password"}' http://localhost:8080/api/v1/signin
```
Response
```sh
# if exists
Welcome, $username

#if not exists or wrong password
user not exist or wrong password
```


## Todo list
- add JWT
- Create a frontend (再看看XD)