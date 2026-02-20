import { Options } from "k6/options";
import { sleep, check, group } from "k6";
import { Counter, Rate, Trend } from "k6/metrics";
import { get, post, put, del, checkStatus } from "./helpers/http-client.js";
import { CreateUserPayload, ApiResponse, User } from "./types/common.js";

// Custom metrics
const userCreated = new Counter("users_created");
const userCreationFailed = new Rate("user_creation_failed");
const userListDuration = new Trend("user_list_duration", true);

export const options: Options = {
    stages: [
        { duration: "1m", target: 20 }, // ramp up
        { duration: "3m", target: 20 }, // steady state
        { duration: "1m", target: 50 }, // peak
        { duration: "3m", target: 50 }, // steady at peak
        { duration: "1m", target: 0 }, // ramp down
    ],
    thresholds: {
        http_req_duration: ["p(95)<300", "p(99)<500"],
        http_req_failed: ["rate<0.01"],
        checks: ["rate>0.95"],
        users_created: ["count>100"],
        user_creation_failed: ["rate<0.05"],
    },
};

export default function (): void {
    group("User CRUD Operations", () => {
        // CREATE
        const payload: CreateUserPayload = {
            name: `user_${__VU}_${__ITER}`,
            email: `user${__VU}_${__ITER}@loadtest.com`,
        };

        const createRes = post<CreateUserPayload>("/api/v1/users", payload);
        const created = checkStatus(createRes, 201, "create user");

        if (created) {
            userCreated.add(1);

            const body = createRes.json() as unknown as ApiResponse<User>;
            const userId = body.data.id;

            // READ
            const getRes = get(`/api/v1/users/${userId}`);
            checkStatus(getRes, 200, "get user");

            // UPDATE
            const updatePayload: CreateUserPayload = {
                name: `updated_${__VU}_${__ITER}`,
                email: `updated${__VU}_${__ITER}@loadtest.com`,
            };
            const updateRes = put<CreateUserPayload>(
                `/api/v1/users/${userId}`,
                updatePayload,
            );
            checkStatus(updateRes, 200, "update user");

            // DELETE
            const deleteRes = del(`/api/v1/users/${userId}`);
            checkStatus(deleteRes, 204, "delete user");
        } else {
            userCreationFailed.add(1);
        }
    });

    group("List Users", () => {
        const listRes = get("/api/v1/users?page=1&limit=10");
        userListDuration.add(listRes.timings.duration);

        check(listRes, {
            "list returns 200": (r) => r.status === 200,
            "list has data": (r) => {
                const body = r.json() as unknown as ApiResponse<User[]>;
                return Array.isArray(body.data);
            },
        });
    });

    sleep(1);
}
