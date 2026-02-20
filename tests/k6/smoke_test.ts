import { Options } from "k6/options";
import { sleep } from "k6";
import { get, checkStatus } from "./helpers/http_client.js";

export const options: Options = {
    vus: 1,
    duration: "10s",
    thresholds: {
        http_req_duration: ["p(95)<500"],
        http_req_failed: ["rate<0.01"],
        checks: ["rate>0.99"],
    },
};

export default function (): void {
    const healthRes = get("/api/v1/health");
    checkStatus(healthRes, 200, "health check");

    const usersRes = get("/api/v1/users");
    checkStatus(usersRes, 200, "list users");

    sleep(1);
}
