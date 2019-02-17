# 学习golang的一些笔记  
本仓库是学习用代码
## 关于环境变量  
请确保正确设置GOROOT和GOPATH，两者不能相同  
GOROOT：就是go lang SDK安装的目录，默认是/usr/local/go  
GOPATH：就是project的根目录，目录有bin，pkg，src组成。GOPATH上可以添加多个project，每个project的根目录用冒号分割就可以。  
但是通过go get下载下来的代码默认存在于第一个GOPATH的src里。  
## 关于go命令  
由于go get对有些包是无法下载，所以用gopm get -g来替代  
地址：github.com/gpmgo/gopm
先用gopm get -g从github上下载代码之后，在通过go install来安装，下载下来的源代码保存在GOPATH对应的目录里

## 关于hot deploy
默认情况下是不支持hot deploy，所以我们需要安装插件  
github.com/pilu/fresh  
在go install之后，不用使用go run，可以直接在project的src／myproject里输入fresh就可  
