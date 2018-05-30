# slackbot

* capabilities

    * search music from AppleMusic by keyword using ```$ hime /music <keyword>``` format
    
* how to build

``` 
$ env GOOS=linux GOARCH=amd64 go build -v github.com/nerocrux/slackbot
$ zip deployment.zip slackbot
```
