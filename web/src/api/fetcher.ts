// web/src/api/fetcher.ts

const API_BASE = import.meta.env.VITE_API_BASE_URL;

type Method = "GET" | "POST" | "PUT" | "DELETE" | "PATCH";

interface RequestParams extends Omit<RequestInit, "body" | "method"> {
    body?: unknown;
    headers?: HeadersInit;
}

async function request<T>(method: Method, endpoint: string, { body, headers, ...rest }: RequestParams = {}): Promise<T> {
    const url = `${API_BASE}${endpoint}`;
    const res = await fetch(url, {
        method,
        headers: {
            "Content-Type": "application/json",
            ...headers,
        },
        body: body ? JSON.stringify(body) : undefined,
        ...rest,
    });

    if (!res.ok) {
        const message = await res.text();
        throw new Error(`[${method} ${endpoint}] ${res.status}: ${message}`);
    }

    return res.json();
}

// Shorthand exports
export const api = {
    get: <T>(endpoint: string, params?: RequestParams) => request<T>("GET", endpoint, params),
    post: <T>(endpoint: string, body?: unknown, params?: RequestParams) =>
        request<T>("POST", endpoint, { ...params, body }),
    put: <T>(endpoint: string, body?: unknown, params?: RequestParams) =>
        request<T>("PUT", endpoint, { ...params, body }),
    del: <T>(endpoint: string, params?: RequestParams) => request<T>("DELETE", endpoint, params),
    patch: <T>(endpoint: string, body?: unknown, params?: RequestParams) =>
        request<T>("PATCH", endpoint, { ...params, body }),
};
