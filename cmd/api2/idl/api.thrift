namespace go api

struct AdditionRequest {
    1: i32 Num1 (api.query="num1");
    2: i32 Num2 (api.query="num2");
}

struct AdditionResponse {
    1: i32 result;
}

service Addition {
    AdditionResponse add(1: AdditionRequest req) (api.get="calculator/add")
}