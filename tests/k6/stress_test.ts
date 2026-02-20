import { Options } from "k6/options";
import { sleep, check } from "k6";
import { Rate } from "k6/metrics";
import { get, post } from "./helpers/http-client.js";
import { CreateUserPayload } from "./types/common.js";

const errorRate = new Rate("errors");

export const options: Options = {
    stages: [
        { duration: "2m", target: 100 },
        { duration: "5m", target: 100 },
        { duration: "2m", target: 200 },
        { duration: "5m", target: 200 },
        { duration: "2m", target: 300 },
        { duration: "5m", target: 300 },
        { duration: "2m", target: 400 },
        { duration: "5m", target: 400 },
        { duration: "5m", target: 0 },
    ],
    thresholds: {
        http_req_duration: ["p(99)<1500"],
        errors: ["rate<0.1"],
    },
};

export default function (): void {
    const readRes = get("/api/v1/users");
    const readSuccess = check(readRes, {
        "GET status is 200": (r) => r.status === 200,
        "GET response time < 1000ms": (r) => r.timings.duration < 1000,
    });
    errorRate.add(!readSuccess);

    const payload: CreateUserPayload = {
        name: `stress_${__VU}_${__ITER}`,
        email: `stress${__VU}_${__ITER}@test.com`,
    };

    const writeRes = post<CreateUserPayload>("/api/v1/users", payload);
    const writeSuccess = check(writeRes, {
        "POST status is 201": (r) => r.status === 201,
        "POST response time < 1500ms": (r) => r.timings.duration < 1500,
    });
    errorRate.add(!writeSuccess);

    sleep(0.5);
}
