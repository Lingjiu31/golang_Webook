'use client';

import React from 'react';
import Link from "next/link";

const App = () => {
    return (
        <div className="min-h-screen flex items-center justify-center bg-[#ffffff]">
            <div className="text-center">
                <h1 className="text-heading-lg font-optimistic text-ink-deep mb-4">小微书</h1>
                <p className="text-subtitle-md text-steel mb-xxl">你的第一个 Web 应用</p>
                <div className="flex gap-4 justify-center">
                    <Link href="/users/login" className="no-underline">
                        <span className="inline-block bg-ink-button text-on-ink-button text-button-md rounded-full py-[14px] px-[30px] font-bold">
                            登录
                        </span>
                    </Link>
                    <Link href="/users/signup" className="no-underline">
                        <span className="inline-block border-2 border-ink-deep text-ink-deep text-button-md rounded-full py-[12px] px-[28px] font-bold">
                            注册
                        </span>
                    </Link>
                </div>
            </div>
        </div>
    );
}

export default App;
