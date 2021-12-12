# go_practice
go进阶训练营作业

第二周作业

问： 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答： 应该。因为sql语句报错后，外部调用者想知道的是哪个操作导致了sql.ErrNoRows,此时返回方可以使用wrap包装一层返回出去。同时在真正处理错误的地方，可以打印出调用堆栈信息，方便更快地定位问题。

作业地址：https://github.com/Wei-Song/go_practice/tree/master/src/week2/work
