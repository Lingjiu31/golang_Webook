import dynamic  from 'next/dynamic'
import {Button, Form, Input, message} from "antd";
import {useEffect, useState} from "react";
import axios from "@/axios/axios";
import {useRouter} from "next/router";
import {ProLayout} from "@ant-design/pro-components";
const WangEditor = dynamic(
    () => import('../../components/editor'),
    {ssr: false},
)

function Page() {
    const router = useRouter()
    const [form] = Form.useForm()
    const [html, setHtml] = useState()
    const artID = router.query.id as string | undefined
    const onFinish = (values: any) => {
        if(artID) {
            values.id = parseInt(artID)
        }
        values.content = html
        axios.post("/articles/edit", values)
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
                    return
                }
                if (res.data?.code == 0) {
                    message.success(res.data?.msg || "保存成功");
                    router.push('/articles/list')
                    return
                }
                message.error(res.data?.msg || "系统错误");
            }).catch((err) => {
                message.error(err?.response?.data?.msg || "网络异常，请重试");
        })
    };
    const publish = () => {
        const values = form.getFieldsValue()
        if (artID) {
            values.id = parseInt(artID)
        }
        values.content = html
        axios.post("/articles/publish", values)
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
                    return
                }
                if (res.data?.code == 0) {
                    message.success(res.data?.msg || "发表成功");
                    router.push('/articles/view?id='+res.data.data)
                    return
                }
                message.error(res.data?.msg || "系统错误");
            }).catch((err) => {
                message.error(err?.response?.data?.msg || "网络异常，请重试");
        })
    }

    useEffect(() => {
        if (!artID) {
            return
        }
        axios.get('/articles/detail/'+artID)
            .then((res) => res.data)
            .then((data) => {
                form.setFieldsValue(data.data)
                setHtml(data.data.content)
            })
            .catch((err) => {
                message.error(err?.response?.data?.msg || "获取文章详情失败")
            })
    }, [form, artID])

    return <div className="min-h-screen bg-[#ffffff]">
        <ProLayout
            title="创作中心"
            logo={null}
            headerRender={false}
            menuRender={false}
        >
            <div className="max-w-[960px] mx-auto px-4 py-section">
                <Form onFinish={onFinish}
                form={form}>
                    <Form.Item name={"title"}
                               rules={[{ required: true, message: '请输入标题' }]}
                    >
                        <Input placeholder={"请输入标题"} size="large" className="!text-heading-sm !font-optimistic !font-medium" />
                    </Form.Item>
                    <WangEditor html={html} setHtmlFn={setHtml}/>
                    <Form.Item>
                        <br/>
                        <Button type="primary" htmlType="submit" size="large">保存</Button>
                        <Button type="default" onClick={publish} size="large"
                            className="ml-3 border-2 border-ink-deep !text-ink-deep !font-bold !rounded-[100px]">
                            发表
                        </Button>
                    </Form.Item>
                </Form>
            </div>
        </ProLayout>
    </div>
}
export default Page
