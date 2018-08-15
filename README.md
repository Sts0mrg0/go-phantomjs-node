
Команда для запуска узла: <br>
1.  ```go build .```<br>
2.  ```docker build ./ -t node```<br>
3.  ```docker run -ti --rm -p 6677:6677 -e "token=YOURTOKEN" -e "server=HUBIP" node``` переменная server должна быть вида - **192.168.0.1** <br>

Исполняемый узел для хаба (https://github.com/arkadybag/go-hub)
