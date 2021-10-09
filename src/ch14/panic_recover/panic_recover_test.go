package panic_recover

// panic 退出先会调用defer 会打印调用栈的信息 ,os.exit不会 且不会打印堆栈

// 僵尸服务进程
// 错误恢复 recover是 好的 吗
// Let it Crash 一旦检测到 ，守护进程就把服务进程 重启
