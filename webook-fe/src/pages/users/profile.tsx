import { ProDescriptions } from '@ant-design/pro-components';
import React, { useState, useEffect } from 'react';
import { Button, message } from 'antd';
import axios from "@/axios/axios";

function Page() {
    let p: Profile = {Email: "", Phone: "", Nickname: "", Birthday:"", AboutMe: ""}
    const [data, setData] = useState<Profile>(p)
    const [isLoading, setLoading] = useState(true)

    useEffect(() => {
        setLoading(true)
        axios.get('/users/profile')
            .then((res) => res.data)
            .then((data) => {
                setData(data.data)
                setLoading(false)
            })
            .catch((err) => {
                setLoading(false)
                message.error(err?.response?.data?.msg || "获取个人信息失败")
            })
    }, [])

    if (isLoading) return <p>Loading...</p>
    if (!data || (!data.Email && !data.Phone && !data.Nickname)) return <p>No profile data</p>

    return (
        <div className="min-h-screen bg-[#ffffff] py-section px-4">
            <div className="max-w-[640px] mx-auto bg-white rounded-[32px] p-xxl border border-[#dee3e9]">
                <ProDescriptions
                    column={1}
                    title={<span className="text-heading-sm text-ink-deep font-optimistic">个人信息</span>}
                >
                    <ProDescriptions.Item label={<span className="text-body-sm-bold text-ink">昵称</span>} valueType="text">
                        {data.Nickname}
                    </ProDescriptions.Item>
                    <ProDescriptions.Item
                        valueType="text"
                        label={<span className="text-body-sm-bold text-ink">邮箱</span>}
                    >{data.Email}
                    </ProDescriptions.Item>
                    <ProDescriptions.Item
                        valueType="text"
                        label={<span className="text-body-sm-bold text-ink">手机</span>}
                    >{data.Phone}
                    </ProDescriptions.Item>
                    <ProDescriptions.Item label={<span className="text-body-sm-bold text-ink">生日</span>} valueType="date">
                        {data.Birthday}
                    </ProDescriptions.Item>
                    <ProDescriptions.Item
                        valueType="text"
                        label={<span className="text-body-sm-bold text-ink">关于我</span>}
                    >
                        {data.AboutMe}
                    </ProDescriptions.Item>
                    <ProDescriptions.Item>
                        <Button href={"/users/edit"} type="primary" size="large">修改</Button>
                    </ProDescriptions.Item>
                </ProDescriptions>
            </div>
        </div>
    )
}

export default Page
