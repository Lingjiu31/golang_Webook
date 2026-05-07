import React, { useState, useEffect } from 'react';
import axios from "@/axios/axios";
import { message } from "antd";

function Page() {
    const [isLoading, setLoading] = useState(false)

    useEffect(() => {
        setLoading(true)
        axios.get('/oauth2/wechat/authurl')
            .then((res) => res.data)
            .then((data) => {
                setLoading(false)
                if(data && data.data) {
                    window.location.href = data.data
                } else {
                    message.error("获取微信授权链接失败")
                }
            })
            .catch((err) => {
                setLoading(false)
                message.error(err?.response?.data?.msg || "网络异常，请重试")
            })
    }, [])

    if (isLoading) return <p>Loading...</p>

    return (
        <div className="min-h-screen flex items-center justify-center bg-[#ffffff] px-4">
            <p>跳转中...</p>
        </div>
    )
}

export default Page
