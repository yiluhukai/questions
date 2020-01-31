### 使用 golang 实现一个问答系统

1. session 模块

利用 session 和 cookie 来实现。

-   session 需要实现 Session 接口，使用接口来设置、删除、获取、保存 session 中的 key
-   SessionManager 用来创建 session 和通过 sessionId 获取 session
-   SessionManager 和 Session 有两种实现，一种位于内存中，一种存在于 redis 中
-   session 需要调用 Init 函数初始化

2.gen_id 模块

-   gen_id 模块用来生成数据库中用户的唯一 id(userid),为啥不用数据库中自增的 id 作为主键呢？ 当分库分表时自增的 id 会造成 id 重复
-   gen_id 使用前需要初始化

3.middlewar 模块

-   middware 模块时一些相关的中间件
-   account 中间件用来处理登录相关的逻辑
