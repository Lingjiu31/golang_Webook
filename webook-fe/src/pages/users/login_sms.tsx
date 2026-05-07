import React, { useState } from 'react';
import { Button, Form, Input, message } from 'antd';
import axios from "@/axios/axios";
import router from "next/router";

const LoginFormSMS: React.FC = () => {
    const [form] = Form.useForm();
    const [countdown, setCountdown] = useState(0);

    const onFinish = (values: any) => {
        axios.post("/users/login_sms", values)
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
                if (res.data.code == 0) {
                    message.success(res.data.msg || "登录成功");
                    router.push('/users/profile')
                    return;
                }
                message.error(res.data.msg || "登录失败")
            }).catch((err) => {
                message.error(err?.response?.data?.msg || "网络异常，请重试");
        })
    };

    const onFinishFailed = () => {
        message.warning("请检查输入")
    };

    const sendCode = () => {
        const phone = form.getFieldValue("phone")
        if (!phone) {
            message.warning("请先输入手机号码");
            return;
        }
        setCountdown(60);
        const timer = setInterval(() => {
            setCountdown(prev => {
                if (prev <= 1) {
                    clearInterval(timer);
                    return 0;
                }
                return prev - 1;
            });
        }, 1000);
        axios.post("/users/login_sms/code/send", {"phone": phone} )
            .then((res) => {
                if(res.status != 200) {
                    message.error(res.statusText);
                    return
                }
                message.success(res.data?.msg || "验证码已发送")
            }).catch((err) => {
                message.error(err?.response?.data?.msg || "发送失败，请重试");
        })
    }

    return (
    <div className="min-h-screen flex items-center justify-center bg-[#ffffff] px-4">
        <div className="w-full max-w-[440px] bg-white rounded-[32px] p-8 md:p-[32px] border border-[#dee3e9]">
            <h1 className="text-heading-sm font-optimistic text-ink-deep mb-2">手机号登录</h1>
            <p className="text-body-sm text-steel mb-xxl">使用手机号和验证码登录</p>
            <Form
                name="basic"
                layout="vertical"
                initialValues={{ remember: true }}
                onFinish={onFinish}
                onFinishFailed={onFinishFailed}
                autoComplete="off"
                form={form}
                size="large"
            >
                <Form.Item
                    label={<span className="text-body-sm-bold text-ink">手机号码</span>}
                    name="phone"
                    rules={[{ required: true, message: '请输入手机号码' }]}
                >
                    <Input placeholder="请输入手机号码" />
                </Form.Item>

                <Form.Item
                    label={<span className="text-body-sm-bold text-ink">验证码</span>}
                    name="code"
                    rules={[{ required: true, message: '请输入验证码' }]}
                >
                    <Input placeholder="请输入验证码" />
                </Form.Item>

                <Form.Item>
                    <Button type="default" onClick={sendCode} disabled={countdown > 0} block size="large"
                        className="h-[44px] border-2 border-ink-deep !text-ink-deep font-bold !rounded-[100px]">
                        {countdown > 0 ? `${countdown}秒后重发` : '发送验证码'}
                    </Button>
                </Form.Item>

                <Form.Item>
                    <Button type="primary" htmlType="submit" block size="large" className="h-[44px]">
                        登录/注册
                    </Button>
                </Form.Item>
            </Form>
        </div>
    </div>
)};

export default LoginFormSMS;
