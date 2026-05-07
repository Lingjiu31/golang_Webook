import type { Metadata } from 'next'
import './globals.css'

export const metadata: Metadata = {
  title: '小微书',
  description: '你的第一个 Web 应用',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="zh-CN">
      <body>{children}</body>
    </html>
  )
}
