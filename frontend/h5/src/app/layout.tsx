import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import './globals.css';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'Happy - 多内容平台',
  description: '包含长短视频、短剧、漫剧、小说、图文等多种内容形态的多内容化平台',
  keywords: '视频,短剧,漫剧,小说,图文,内容平台',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="zh-CN">
      <body className={inter.className}>{children}</body>
    </html>
  );
}
