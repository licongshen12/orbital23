namespace go demo3

struct apiReq {
    1: string Name (api.query="name");
}

struct apiResp {
    1: string RespBody;
}

struct OtherReq {
    1: string Other (api.body="other");
}

struct OtherResp {
    1: string Resp;
}


service apiService {
    apiResp HelloMethod(1: apiReq request) (api.get="/hello");
    OtherResp OtherMethod(1: OtherReq request) (api.post="/other");
}

service NewService {
    apiResp NewMethod(1: apiReq request) (api.get="/new");
}