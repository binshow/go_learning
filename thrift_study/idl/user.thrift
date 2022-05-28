namespace go Sample
namespace php Sample

# thrift -r --gen go user.thrift  生成对应的thrift文件


struct User {
    1:required i32 id;
    2:required string name;
    3:optional string address;
}

typedef map<string, string> Data

struct Response {
    1:required i32 errCode; //错误码
    2:required string errMsg; //错误信息
    3:required Data data;
}

//定义服务
service Greeter {
    Response SayHello(
        1:required User user
    )

    Response GetUser(
        1:required i32 uid
    )
}