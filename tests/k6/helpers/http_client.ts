import http, { RefinedResponse, ResponseType } from "k6/http";
import { check } from "k6";

const BASE_URL: string = __ENV.BASE_URL || "http://localhost:8080";

interface RequestOptions {
    headers?: Record<string, string>;
    tags?: Record<string, string>;
}

const defaultHeaders: Record<string, string> = {
    "Content-Type": "application/json",
    Accept: "application/json",
};

export function get(
    path: string,
    options: RequestOptions = {},
): RefinedResponse<ResponseType> {
    const url = `${BASE_URL}${path}`;
    return http.get(url, {
        headers: { ...defaultHeaders, ...options.headers },
        tags: options.tags,
    });
}

export function post<T>(
    path: string,
    body: T,
    options: RequestOptions = {},
): RefinedResponse<ResponseType> {
    const url = `${BASE_URL}${path}`;
    return http.post(url, JSON.stringify(body), {
        headers: { ...defaultHeaders, ...options.headers },
        tags: options.tags,
    });
}

export function put<T>(
    path: string,
    body: T,
    options: RequestOptions = {},
): RefinedResponse<ResponseType> {
    const url = `${BASE_URL}${path}`;
    return http.put(url, JSON.stringify(body), {
        headers: { ...defaultHeaders, ...options.headers },
        tags: options.tags,
    });
}

export function del(
    path: string,
    options: RequestOptions = {},
): RefinedResponse<ResponseType> {
    const url = `${BASE_URL}${path}`;
    return http.del(url, null, {
        headers: { ...defaultHeaders, ...options.headers },
        tags: options.tags,
    });
}

// Helper untuk check response
export function checkStatus(
    res: RefinedResponse<ResponseType>,
    expectedStatus: number,
    label: string,
): boolean {
    return check(res, {
        [`${label} - status ${expectedStatus}`]: (r) =>
            r.status === expectedStatus,
        [`${label} - response time < 500ms`]: (r) => r.timings.duration < 500,
    });
}
