//https://golangbyexample.com/using-context-in-golang-complete-guide/

Following is a list of best practices that you can follow while using a context.
1. Do not store a context within a struct type

2. Context should flow through your program. For example, in case of an HTTP request, a new context can be created for each incoming request which can be used to hold a request_id or put some common information in the context like currently logged in user which might be useful for that particular request.
Always pass context as the first argument to a function.

3. Whenever you are not sure whether to use the context or not, it is better to use the context.ToDo() as a placeholder.

4. Only the parent goroutine or function should the cancel context. Therefore do not pass the cancelFunc to downstream goroutines or functions. Golang will allow you to pass the cancelFunc around to child goroutines but it is not a recommended practice