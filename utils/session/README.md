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

