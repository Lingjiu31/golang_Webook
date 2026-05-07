import React from 'react';
import { Button, Form, Input, message } from 'antd';
import axios from "@/axios/axios";
import Link from "next/link";
import router from "next/router";

const LoginForm: React.FC = () => {

    const onFinish = (values: any) => {
        axios.post("/users/login", values)
            .then((res) => {
                if(res.status != 200) {
                    message.error(res.statusText);
                    return
                }
                if(typeof res.data == 'string') {
                    if (res.data.includes("成功")) {
                        message.success(res.data);
                        router.push('/articles/list')
                    } else {
                        message.error(res.data);
                    }
                } else {
                    const msg = res.data?.msg || JSON.stringify(res.data)
                    if(res.data.code == 0) {
                        message.success(msg);
                        router.push('/articles/list')
                    } else {
                        message.error(msg);
                    }
                }
            }).catch((err) => {
                message.error(err?.response?.data?.msg || "网络异常，请重试");
        })
    };

    const onFinishFailed = () => {
        message.warning("请检查输入")
    };

    return (
    <div className="min-h-screen flex items-center justify-center bg-[#ffffff] px-4">
        <div className="w-full max-w-[440px] bg-white rounded-[32px] p-8 md:p-[32px] border border-[#dee3e9]">
            <h1 className="text-heading-sm font-optimistic text-ink-deep mb-2">登录</h1>
            <p className="text-body-sm text-steel mb-xxl">欢迎回到小微书</p>
            <Form
                name="basic"
                layout="vertical"
                initialValues={{ remember: true }}
                onFinish={onFinish}
                onFinishFailed={onFinishFailed}
                autoComplete="off"
                size="large"
            >
                <Form.Item
                    label={<span className="text-body-sm-bold text-ink">邮箱</span>}
                    name="email"
                    rules={[{ required: true, message: '请输入邮箱' }]}
                >
                    <Input placeholder="请输入邮箱地址" />
                </Form.Item>

                <Form.Item
                    label={<span className="text-body-sm-bold text-ink">密码</span>}
                    name="password"
                    rules={[{ required: true, message: '请输入密码' }]}
                >
                    <Input.Password placeholder="请输入密码" />
                </Form.Item>

                <Form.Item>
                    <Button type="primary" htmlType="submit" block size="large" className="h-[44px]">
                        登录
                    </Button>
                </Form.Item>
            </Form>

            <div className="flex justify-center gap-4 text-body-sm text-steel pt-base border-t border-hairline-soft">
                <Link href={"/users/login_sms"} className="text-meta-link font-bold no-underline">
                    手机号登录
                </Link>
                <Link href={"/users/login_wechat"} className="text-meta-link font-bold no-underline">
                    微信扫码登录
                </Link>
                <Link href={"/users/signup"} className="text-meta-link font-bold no-underline">
                    注册
                </Link>
            </div>
        </div>
    </div>
)};

export default LoginForm;
