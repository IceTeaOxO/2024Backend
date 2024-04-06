import http from 'k6/http';
import { sleep } from 'k6';

export let options = {
    stages: [
        { duration: "5s", target: 10 },
        { duration: "5s", target: 100 },
        { duration: "5s", target: 1000 },
        { duration: "1s", target: 10000 },
      ],
    thresholds: {
        // 期望在整個測試執行過程中，錯誤率必須低於 5%
        http_req_failed: ["rate<0.05"],
        // 平均請求必須在 300ms 內完成，90% 的請求必須在 200ms 內完成
        http_req_duration: ["avg < 300", "p(90) < 100"],
      }
};

export default function () {
  const res = http.get('http://127.0.0.1:8080/api/v1/ad?offset=0&limit=10');
  
  if (res.status !== 200) {
    console.error(`Unexpected status code: ${res.status}`);
  }

  sleep(1);
}
// k6 run --vus 10000 --duration 10s stress_test.js
