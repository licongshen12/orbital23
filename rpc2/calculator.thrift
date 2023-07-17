namespace go calculator

struct Request {
    1: i32 num1
    2: i32 num2
}

struct Response {
    1: i32 result
}

service Calculator {
    Response Add(1: Request num1, 2: Request num2)
}