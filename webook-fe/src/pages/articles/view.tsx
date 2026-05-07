import React, {useState, useEffect} from 'react';
import axios, {Result} from "@/axios/axios";
import {Button, Modal, QRCode, Typography, message} from "antd";
import {ProLayout} from "@ant-design/pro-components";
import {EyeOutlined, LikeOutlined, MoneyCollectOutlined, StarOutlined} from "@ant-design/icons";
import {useRouter} from "next/router";

export const dynamic = 'force-dynamic'

interface CodeURL {
    codeURL: string
    rid: number
}


function Page(){
    const router = useRouter()
    const [data, setData] = useState<Article>()
    const [openQRCode, setOpenQRCode] = useState(false)
    const [codeURL, setCodeURL] = useState('')
    const [isLoading, setLoading] = useState(false)
    const [rid, setRid] = useState(0)
    const artID = (router.query.id as string) || '1'
    useEffect(() => {
        setLoading(true)
        axios.get('/articles/pub/'+artID)
            .then((res) => res.data)
            .then((data) => {
                setData(data.data)
                setLoading(false)
            })
            .catch((err) => {
                setLoading(false)
                message.error(err?.response?.data?.msg || "获取文章失败")
            })
    }, [artID])

    if (isLoading) return <p>Loading...</p>
    if (!data) return <p>No data</p>

    const like = () => {
        axios.post('/articles/pub/like', {
            id: parseInt(artID),
            like: !data.liked
        })
            .then((res) => res.data)
            .then((res) => {
                if(res.code == 0) {
                    setData(prev => prev ? {
                        ...prev,
                        liked: !prev.liked,
                        likeCnt: prev.liked ? prev.likeCnt - 1 : prev.likeCnt + 1
                    } : prev)
                }
            })
            .catch((err) => {
                message.error(err?.response?.data?.msg || "操作失败")
            })
    }

    const collect = () => {
        if (data.collected) {
            return
        }
        axios.post('/articles/pub/collect', {
            id: parseInt(artID),
            cid: 0,
        })
            .then((res) => res.data)
            .then((res) => {
                if(res.code == 0) {
                    setData(prev => prev ? {
                        ...prev,
                        collected: !prev.collected,
                        collectCnt: prev.collectCnt + 1
                    } : prev)
                }
            })
            .catch((err) => {
                message.error(err?.response?.data?.msg || "操作失败")
            })
    }
    const reward = function () {
        axios.post<Result<CodeURL>>('/articles/pub/reward', {
            id: parseInt(artID),
            amt: 1,
        })
            .then((res) => res.data)
            .then((res) => {
                setCodeURL(res.data.codeURL)
                setRid(res.data.rid)
                setOpenQRCode(true)
            })
            .catch((err) => {
                message.error(err?.response?.data?.msg || "操作失败")
            })
    }

    const closeModal = () => {
        setOpenQRCode(false)
        const currentRid = rid
        if(currentRid > 0) {
            axios.post<Result<string>>('/reward/detail', {
                rid: currentRid,
            }).then((res) => res.data)
                .then((res) => {
                    if(res.data == 'RewardStatusPayed') {
                        message.success("打赏成功")
                    } else {
                        console.log(res.data)
                    }
                })
                .catch((err) => {
                    message.error(err?.response?.data?.msg || "查询打赏状态失败")
                })
        }
    }

    return (
        <div className="min-h-screen bg-[#ffffff]">
            <ProLayout pure={true} headerRender={false} menuRender={false}>
                <div className="max-w-[800px] mx-auto px-4 py-section">
                    <div className="bg-white rounded-[32px] p-xxl border border-[#dee3e9] mb-xl">
                        <Typography>
                            <Typography.Title className="!text-heading-lg !font-optimistic !font-medium !text-ink-deep !mb-6">
                                {data.title}
                            </Typography.Title>
                            <Typography.Paragraph>
                                <div className="text-body-md text-ink leading-6"
                                    dangerouslySetInnerHTML={{__html: data.content}}></div>
                            </Typography.Paragraph>
                        </Typography>
                    </div>

                    <div className="flex flex-wrap gap-3">
                        <Button icon={<EyeOutlined />} type="default" className="!rounded-[100px] !font-bold">
                            {data.readCnt}
                        </Button>
                        <Button onClick={reward} icon={<MoneyCollectOutlined />} type="primary" className="!rounded-[100px]">
                            打赏一分钱
                        </Button>
                        <Button onClick={like} icon={<LikeOutlined style={data.liked? {color: "#e41e3f"}:{}}/>}
                            type="default" className="!rounded-[100px] !font-bold">
                            {data.likeCnt}
                        </Button>
                        <Button onClick={collect} icon={<StarOutlined style={data.collected? {color: "#f7b928"}:{}}/>}
                            type="default" className="!rounded-[100px] !font-bold">
                            {data.collectCnt}
                        </Button>
                    </div>

                    <Modal title="扫描二维码" open={openQRCode} onCancel={closeModal} onOk={closeModal}>
                        <QRCode value={codeURL} size={128} />
                    </Modal>
                </div>
            </ProLayout>
        </div>
    )
}

export default Page
