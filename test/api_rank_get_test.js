import http from "k6/http";
import {
  check,
  sleep
} from "k6";

const baseURL = "http://sea.robber.happygod.cn/api";

export let options = {
  vus: 100,
  duration: "10s"
};

export default function () {
  let res = http.get(baseURL + '/rank')
  check(res, {
    "get status was 200": (r) => r.status == 200,
    "get transaction time OK": (r) => r.timings.duration < 500
  });
};