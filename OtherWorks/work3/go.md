# GO的优缺点

## go的优点
### 高并发
原生支持多核、并发，运行效率高。
### 轻量级
语法简单，易学、易读。开发效率高，编译速度快。
### 云原生
对容器化快速弹性的业务支持好，易于部署。
不用每个容器，微服务都装jar包，体积小

## go的缺点
### 生态
新语言难以避免的生态不如java、Python丰富。
### 性能
虽然被称为新世纪的C语言，但性能并不能达到C语言的水平，和java差不多。
### 生的晚
越来越多的新项目在使用go，但是互联网行业又有多少新项目呢？
对于大多数已有web项目的并发量，go并不比java有太多优势。
即使有一定优势，重构的动机也可能不足。

## 为什么大厂在用go重构
### 替代PHP
go与PHP开发效率近似，但性能远高于PHP。
### 替代C++
go开发效率远高于C++，内存管理强于C++，
虽然运行效率还略逊一筹，但追求极致性能为什么不用C呢？
### 替代java
go比java更加轻量级，开发效率更高，并发性能更强。
### 云原生
go是目前最适合云原生的语言。