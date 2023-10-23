# scheduler定时任务
- 定义了一个asynq.scheduler，定义了一个定时任务task := asynq.NewTask(jobtype.ScheduleSettleRecord, nil)
- 定义了一个定时任务task := asynq.NewTask(jobtype.ScheduleSettleRecord, nil)
- 注册了指定时间间隔的任务入队计划entryID, err := l.svcCtx.Scheduler.Register("*/1 * * * *", task)
- 总结：即每隔1分钟，会入队一个jobtype.ScheduleSettleRecord类型的任务，该任务payload为nil

# job(定时任务和延迟处理任务)
- 定义了用于处理异步任务的多路复用路由器ServeMux
- 对ServeMux注册了两种任务处理程序，分别是mux.Handle(jobtype.ScheduleSettleRecord,NewSettleRecordHandler(l.svcCtx))和mux.Handle(jobtype.DeferCloseHomestayOrder,NewCloseHomestayOrderHandler(l.svcCtx))


# 总结说明
- go-zero-looklook的scheduler定时任务仅仅用于展示功能用法示例
- 延迟处理任务用来启动一个worker(asynq.NewServer实例，app/mqueue/cmd/job/mqueue.go)以便异步处理超时未支付订单，避免使用定时任务轮训去关闭订单

