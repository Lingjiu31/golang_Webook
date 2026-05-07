import React from 'react';
import { Button, Form, Input, message } from 'antd';
import axios from "@/axios/axios";
import Link from "next/link";
import router from "next/router";

const SignupForm: React.FC = () => {
    const [form] = Form.useForm();

    const onFinish = (values: any) => {
        axios.post("/users/signup", values)
            .then((res) => {
                if(res.status != 200) {
                    message.error(res.statusText);
                    return
                }
                if(typeof res.data == 'string') {
                    if (res.data.includes("成功")) {
                        message.success(res.data);
                        router.push('/users/login')
                    } else {
                        message.error(res.data);
                    }
                } else {
                    const msg = res.data?.msg || JSON.stringify(res.data)
                    if(res.data.code == 0) {
                        message.success(msg);
                        router.push('/users/login')
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
            <h1 className="text-heading-sm font-optimistic text-ink-deep mb-2">注册</h1>
            <p className="text-body-sm text-steel mb-xxl">创建你的小微书账号</p>
            <Form
                form={form}
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
                    rules={[{ required: true, message: '请输入邮箱' }, { type: 'email', message: '请输入有效的邮箱地址' }]}
                >
                    <Input placeholder="请输入邮箱地址" />
                </Form.Item>

                <Form.Item
                    label={<span className="text-body-sm-bold text-ink">密码</span>}
                    name="password"
                    rules={[{ required: true, message: '请输入密码' }, { min: 6, message: '密码至少6位' }]}
                >
                    <Input.Password placeholder="请输入密码" />
                </Form.Item>

                <Form.Item
                    label={<span className="text-body-sm-bold text-ink">确认密码</span>}
                    name="confirmPassword"
                    dependencies={['password']}
                    rules={[
                        { required: true, message: '请确认密码' },
                        ({ getFieldValue }) => ({
                            validator(_, value) {
                                if (!value || getFieldValue('password') === value) {
                                    return Promise.resolve();
                                }
                                return Promise.reject(new Error('两次输入的密码不一致'));
                            },
                        }),
                    ]}
                >
                    <Input.Password placeholder="请再次输入密码" />
                </Form.Item>

                <Form.Item>
                    <Button type="primary" htmlType="submit" block size="large" className="h-[44px]">
                        注册
                    </Button>
                </Form.Item>
            </Form>

            <div className="flex justify-center pt-base border-t border-hairline-soft">
                <span className="text-body-sm text-steel">已有账号？</span>
                <Link href={"/users/login"} className="text-meta-link font-bold no-underline ml-1">
                    登录
                </Link>
            </div>
        </div>
    </div>
)};

export default SignupForm;
