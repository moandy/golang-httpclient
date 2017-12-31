## Http-Client

#### 测试

```golang
➜  goroutine ./goroutine -h

Usage of ./goroutine:

  -n    Whether to use native method

使用native方法：

➜  goroutine time ./goroutine -n

./goroutine -n  0.00s user 0.00s system 0% cpu 5.378 total

不使用native方法：

➜  goroutine time ./goroutine   

./goroutine  0.00s user 0.00s system 0% cpu 0.728 total
```



#### 知识框架

![](http://o7d2h0gjo.bkt.clouddn.com/2017-12-25-15142095283938.png)

![](http://o7d2h0gjo.bkt.clouddn.com/2017-12-25-15142096539053.png)

![](http://o7d2h0gjo.bkt.clouddn.com/2017-12-25-15142096602737.png)

#### 结论

通过测试，我们很容易发现异步方法比同步方法更加有效。