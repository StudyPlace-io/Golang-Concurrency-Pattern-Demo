### 定时任务模式介绍：需要定期执行某些func时可以使用。

- 一句话概括：定时触发任务。
- 实现方法：
    1. 使用github.com/antlabs/timer包做二次封装。
    2. 使用github.com/robfig/cron/v3包做二次封装
    3. 使用定时器+select实现定时轮巡执行的方法(与k8s内部工具一致，cronfunc包中)

  
