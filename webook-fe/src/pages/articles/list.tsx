import {EditOutlined} from '@ant-design/icons';
import {ProLayout, ProList} from '@ant-design/pro-components';
import {Button, message, Tag} from 'antd';
import React, {useEffect, useState} from 'react';
import axios from "@/axios/axios";
import router from "next/router";

const IconText = ({ icon, text, onClick }: { icon: any; text: string, onClick: any}) => (
    <Button onClick={onClick} type="default" size="small" className="!rounded-[100px] !font-bold">
    {React.createElement(icon, { style: { marginInlineEnd: 8 } })}
        {text}
  </Button>
);

interface ArticleItem {
    id: bigint
    title: string
    status: number
    abstract: string
}

const ArticleList = () => {
    const [data, setData] = useState<Array<ArticleItem>>([])
    const [loading, setLoading] = useState<boolean>()
    useEffect(() => {
        setLoading(true)
        axios.post('/articles/list', {
            "offset": 0,
            "limit": 100,
        }).then((res) => res.data)
            .then((data) => {
                setData(data.data)
                setLoading(false)
            })
            .catch((err) => {
                setLoading(false)
                message.error(err?.response?.data?.msg || "获取文章列表失败")
            })
    }, [])
    return (
        <div className="min-h-screen bg-[#ffffff]">
            <ProLayout
                title="创作中心"
                logo={null}
                headerRender={false}
                menuRender={false}
            >
                <div className="max-w-[960px] mx-auto px-4 py-section">
                    <ProList<ArticleItem>
                        toolBarRender={() => {
                            return [
                                <Button key="3" type="primary" size="large" href={"/articles/edit"}>
                                    写作
                                </Button>,
                            ];
                        }}
                        itemLayout="vertical"
                        rowKey="id"
                        headerTitle={<span className="text-heading-sm font-optimistic text-ink-deep">文章列表</span>}
                        loading={loading}
                        dataSource={data}
                        metas={{
                            title: {
                                dataIndex: "title",
                                render: (_, record) => (
                                    <span className="text-body-md-bold text-ink-deep">{record.title}</span>
                                ),
                            },
                            description: {
                                render: (data, record, idx) => {
                                    switch (record.status) {
                                        case 1:
                                            return (
                                                <Tag color="processing" className="!rounded-[100px] !font-bold">未发表</Tag>
                                            )
                                        case 2:
                                            return (
                                                <Tag color="success" className="!rounded-[100px] !font-bold">已发表</Tag>
                                            )
                                        case 3:
                                            return (
                                                <Tag color="warning" className="!rounded-[100px] !font-bold">仅自己可见</Tag>
                                            )
                                        default:
                                            return (<></>)
                                    }
                                },
                            },
                            actions: {
                                render: (text, row) => [
                                    <IconText
                                        icon={EditOutlined}
                                        text="编辑"
                                        onClick={() => {
                                            router.push("/articles/edit?id=" + row.id.toString())
                                        }}
                                        key="list-vertical-edit-o"
                                    />,
                                ],
                            },
                            content: {
                                render: (node, record) => {
                                    return (
                                        <div
                                            className="text-body-sm text-charcoal"
                                            dangerouslySetInnerHTML={{__html: record.abstract}}>
                                        </div>
                                    )
                                }
                            },
                        }}
                    />
                </div>
            </ProLayout>
        </div>
    );
};

export default ArticleList;
