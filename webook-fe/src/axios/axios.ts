import axios from "axios";
import router from "next/router";

const BACKEND_BASE_URL = "http://localhost:8082";

const instance = axios.create({
    baseURL: BACKEND_BASE_URL,
    withCredentials: true
})

export interface Result<T> {
    code: number,
    msg: string,
    data: T,
}

instance.interceptors.response.use(function (resp) {
    const newToken = resp.headers["x-jwt-token"]
    const newRefreshToken = resp.headers["x-refresh-token"]
    if (newToken) {
        localStorage.setItem("token", newToken)
    }
    if (newRefreshToken) {
        localStorage.setItem("refresh_token", newRefreshToken)
    }
    if (resp.status === 401) {
        router.push("/users/login")
    }
    return resp
}, (err) => {
    console.error(err)
    if (err.response?.status === 401) {
        router.push("/users/login")
    }
    return Promise.reject(err)
})

instance.interceptors.request.use((req) => {
    const token = localStorage.getItem("token")
    if (token) {
        req.headers.setAuthorization("Bearer " + token, true)
    }
    return req
}, (err) => {
    console.error(err)
    return Promise.reject(err)
})

export default instance