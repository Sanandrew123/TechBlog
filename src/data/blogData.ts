import { BlogPost } from '../types/blog';

export const blogPosts: BlogPost[] = [
  {
    id: '1',
    title: 'React 18新特性深度解析：并发渲染与Suspense',
    content: `
# React 18新特性深度解析

React 18引入了令人兴奋的新特性，其中最重要的是并发渲染。这篇文章将深入探讨这些新特性如何改变我们构建React应用的方式。

## 并发渲染的核心概念

并发渲染允许React在渲染过程中中断、暂停和恢复工作。这意味着React可以：

- 响应用户输入而不阻塞UI
- 优先处理更紧急的更新
- 在后台预渲染屏幕

## Suspense的进化

Suspense在React 18中得到了显著增强，现在支持服务端渲染和并发特性...
    `,
    excerpt: '深入探讨React 18的并发渲染、Suspense等新特性，以及它们如何改变现代前端开发。',
    author: '技术探索者',
    publishDate: '2024-01-15',
    tags: ['React', '前端开发', '并发渲染'],
    readTime: 8
  },
  {
    id: '2',
    title: 'Go语言在微服务架构中的最佳实践',
    content: `
# Go语言在微服务架构中的最佳实践

Go语言因其出色的并发性能和简洁的语法，成为构建微服务的首选语言之一。本文分享一些实践经验。

## 服务拆分策略

在设计微服务时，合理的服务边界划分至关重要...

## gRPC vs REST API

根据不同场景选择合适的通信协议...
    `,
    excerpt: '分享Go语言在微服务架构中的实践经验，包括服务拆分、通信协议选择等关键问题。',
    author: '技术探索者',
    publishDate: '2024-01-10',
    tags: ['Go', '微服务', '后端开发'],
    readTime: 12
  },
  {
    id: '3',
    title: '深度学习模型部署：从训练到生产的完整流程',
    content: `
# 深度学习模型部署完整指南

将训练好的深度学习模型部署到生产环境是一个复杂但关键的过程。本文将详细介绍整个流程。

## 模型优化

在部署前，需要对模型进行优化以提高推理速度...

## 容器化部署

使用Docker和Kubernetes进行模型部署...
    `,
    excerpt: '完整介绍深度学习模型从训练到生产部署的全流程，包括模型优化、容器化等关键步骤。',
    author: '技术探索者',
    publishDate: '2024-01-05',
    tags: ['人工智能', 'MLOps', 'Docker'],
    readTime: 15
  }
];