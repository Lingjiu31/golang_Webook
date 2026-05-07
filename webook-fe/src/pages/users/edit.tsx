import React, {useEffect, useState} from 'react';
import {Button, DatePicker, Form, Input, message} from 'antd';
import axios from "@/axios/axios";
import moment from 'moment';
import router from "next/router";

const { TextArea } = Input;

function EditForm() {
    const [data, setData] = useState<Profile>({ Email: "", Phone: "", Nickname: "", Birthday: "", AboutMe: "" })
    const [isLoading, setLoading] = useState(true)

    useEffect(() => {
        setLoading(true)
        axios.get('/users/profile')
            .then((res) => res.data)
            .then((data) => {
                setData(data.data || { Email: "", Phone: "", Nickname: "", Birthday: "", AboutMe: "" } as Profile)
                setLoading(false)
            })
            .catch((err) => {
                setLoading(false)
                message.error(err?.response?.data?.msg || "获取个人信息失败")
            })
    }, [])

    const onFinish = (values: any) => {
        if (values.birthday) {
            values.birthday = moment(values.birthday).format("YYYY-MM-DD")
        }
        axios.post("/users/edit", values)
            .then((res) => {
                if(res.status != 200) {
                    message.error(res.statusText);
                    return
                }
                if(typeof res.data == 'string') {
                    if (res.data.includes("成功")) {
                        message.success(res.data);
                        router.push('/users/profile')
                    } else {
                        message.error(res.data);
                    }
                    return
                }
                if (res.data?.code == 0) {
                    message.success(res.data?.msg || "保存成功");
                    router.push('/users/profile')
                    return
                }
                message.error(res.data?.msg || "系统错误");
            }).catch((err) => {
                message.error(err?.response?.data?.msg || "网络异常，请重试");
        })
    };

    const onFinishFailed = () => {
        message.warning("请检查输入")
    };

    if (isLoading) return <p>Loading...</p>
    return (
    <div className="min-h-screen flex items-center justify-center bg-[#ffffff] px-4">
        <div className="w-full max-w-[480px] bg-white rounded-[32px] p-8 md:p-[32px] border border-[#dee3e9]">
            <h1 className="text-heading-sm font-optimistic text-ink-deep mb-2">编辑资料</h1>
            <p className="text-body-sm text-steel mb-xxl">修改你的个人信息</p>
            <Form
                name="basic"
                layout="vertical"
                initialValues={{
                    birthday: data.Birthday ? moment(data.Birthday, 'YYYY-MM-DD') : undefined,
                    nickname: data.Nickname,
                    aboutMe: data.AboutMe
                }}
                onFinish={onFinish}
                onFinishFailed={onFinishFailed}
                autoComplete="off"
                size="large"
            >
                <Form.Item
                    label={<span className="text-body-sm-bold text-ink">昵称</span>}
                    name="nickname"
                >
                    <Input placeholder="请输入昵称" />
                </Form.Item>

                <Form.Item
                    label={<span className="text-body-sm-bold text-ink">生日</span>}
                    name="birthday"
                >
                    <DatePicker format={"YYYY-MM-DD"} placeholder={"选择日期"} className="w-full" />
                </Form.Item>

                <Form.Item
                    label={<span className="text-body-sm-bold text-ink">关于我</span>}
                    name="aboutMe"
                >
                    <TextArea rows={4} placeholder="介绍一下自己..." />
                </Form.Item>

                <Form.Item>
                    <Button type="primary" htmlType="submit" block size="large" className="h-[44px]">
                        提交
                    </Button>
                </Form.Item>
            </Form>
        </div>
    </div>
    )
}

export default EditForm;
