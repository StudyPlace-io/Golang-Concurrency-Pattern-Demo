### 生产者-消费者模式介绍：共享一个固定大小的缓冲区(chan)，A goroutine生产数据放入缓冲区，B goroutine从缓冲区中取出数据进行消费。
![](https://github.com/StudyPlace-io/Golang-Concurrency-Pattern-Demo/blob/main/image/producer-consumer.jpg?raw=true)
- 一句话概括：
- 实现方法：
    1. 一对一：
    2. 多对一：
    3. 多对多：
- 适用场景：把生产数据和消费数据双方分开，将生产数据与消费数据的过程解耦。
  1. Excutor任务执行框架
  2. 任务的处理时间较长的情况下
  
