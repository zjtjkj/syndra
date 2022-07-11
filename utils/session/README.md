# Session
Provide the best golang practice of session.

提供golang session的最佳实践。

## Creating sessions

session实现基本原理是服务端为每一个客户端维护客户信息，然后客户端通过`unique session id`来获取这些信息，当用户访问Web服务时，服务端通过以下三个步骤来创建一个session：

* 创建一个`unique session id`

* 为session分配一块存储空间

  > 通常我们将session存储在内存中，但是如果服务发生宕机等意外，所有的session都将丢失，如果这些信息比较重要，会造成严重的生产问题。为了解决这一问题，我们可以将session存储在数据库或文件中，持久化保存这些数据，这样也更方便于其他服务使用这些数据。

* 将`unique session id`返回给客户端

## Use Go to manage sessions

### Session management design

设计session management时，我们需要考虑以下功能：

* 全局session管理
* 保证session id的唯一性
* 每个用户有且仅有一个session
* session存储（内存，文件或数据库）
* 处理过期session

为此我们使用Manager来管理全局session，使用类似于简单工厂的方法，实现各种session存储方式的Provider实例。

当前支持的存储模式：

| 模式 | Provider Type |
| ---- | ------------- |
| 内存 | Memory        |

### Example

使用内存存储session示例：

```golang
package main

import (
   "fmt"
   "github.com/zjtjkj/syndra/utils/session"
)

func main() {
   // init manager
   manager, err := session.NewManager(session.Memory, 172800)
   if err != nil {
      panic(err)
   }
   // create new session
   sess1, err := manager.SessionStart("")
   if err != nil {
      panic(err)
   }
   fmt.Printf("sess1 id: %s, name: %s, age: %d\n", sess1.SessionID(), sess1.Get("name"), sess1.Get("age"))
   _ = sess1.Set("name", "Tom")
   _ = sess1.Set("age", 18)
   // get session
   sess2, err := manager.SessionStart(sess1.SessionID())
   if err != nil {
      panic(err)
   }
   fmt.Printf("sess2 id: %s, name: %s, age: %d\n", sess2.SessionID(), sess2.Get("name"), sess2.Get("age"))
   // destroy session
   err = manager.SessionDestroy(sess2.SessionID())
   if err != nil {
      panic(err)
   }
   // get session again
   sess3, err := manager.SessionStart(sess2.SessionID())
   if err != nil {
      panic(err)
   }
   fmt.Printf("sess3 id: %s, name: %s, age: %d\n", sess3.SessionID(), sess3.Get("name"), sess3.Get("age"))
}
```
