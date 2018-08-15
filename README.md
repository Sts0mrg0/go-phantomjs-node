
Команда для запуска узла: <br>
1.  ```go build .```<br>
2.  ```docker build ./ -t node```<br>
3.  ```docker run -ti --rm -p 6677:6677 -e "token=YOURTOKEN" -e "server=HUBIP" node``` переменная __server__ должна быть вида - __192.168.0.1__ <br>

Хаб - https://github.com/arkadybag/go-hub
