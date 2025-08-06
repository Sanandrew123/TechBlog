# 🚀 TechBlog - 个人技术博客

一个现代化的个人博客应用，专注于前沿技术分享、人工智能和后端开发。

## ✨ 特性

- 🎨 **现代化设计** - 采用响应式设计，支持各种设备
- 📝 **博客系统** - 完整的文章展示、分类和搜索功能
- 🏷️ **标签管理** - 支持标签筛选和分类
- 📱 **移动友好** - 完美适配移动端和平板设备
- ⚡ **高性能** - 基于React和TypeScript构建
- 🎯 **SEO优化** - 良好的搜索引擎优化
- 🌙 **优雅动画** - 流畅的交互动画效果

## 🛠️ 技术栈

- **前端框架**: React 18 + TypeScript
- **路由**: React Router v6
- **样式**: CSS Modules + 现代CSS
- **构建工具**: Vite
- **开发工具**: ESLint + TypeScript

## 📁 项目结构

```
personal-blog/
├── public/                 # 静态资源
├── src/
│   ├── components/         # 可复用组件
│   │   ├── Header/         # 导航头部
│   │   ├── Footer/         # 页脚
│   │   ├── BlogCard/       # 博客卡片
│   │   ├── Hero/           # 首页英雄区
│   │   └── Layout/         # 布局容器
│   ├── pages/              # 页面组件
│   │   ├── Home/           # 首页
│   │   ├── Blog/           # 博客列表
│   │   ├── BlogDetail/     # 博客详情
│   │   ├── About/          # 关于页面
│   │   └── Contact/        # 联系页面
│   ├── data/               # 数据文件
│   ├── types/              # TypeScript类型定义
│   ├── hooks/              # 自定义Hooks
│   ├── utils/              # 工具函数
│   └── styles/             # 全局样式
├── package.json
├── tsconfig.json
├── vite.config.ts
└── README.md
```

## 🚀 快速开始

### 环境要求
- Node.js >= 16
- npm 或 yarn

### 安装依赖
```bash
npm install
```

### 启动开发服务器
```bash
npm run dev
```

### 构建生产版本
```bash
npm run build
```

### 预览生产版本
```bash
npm run preview
```

## 📖 页面介绍

- **首页** - 展示精选文章和个人简介
- **博客** - 文章列表页，支持搜索和标签筛选
- **博客详情** - 文章详情页，支持分享和导航
- **关于** - 个人介绍和技能展示
- **联系** - 联系表单和联系方式

## 🎨 设计特色

- **渐变色彩** - 使用现代渐变色彩方案
- **卡片式布局** - 清晰的信息层次结构
- **响应式网格** - 自适应不同屏幕尺寸
- **微交互** - 细腻的悬停和点击效果
- **代码展示** - 优雅的代码块样式

## 🔧 自定义配置

### 修改博客数据
编辑 `src/data/blogData.ts` 文件来添加或修改博客文章。

### 样式自定义
在 `src/styles/variables.css` 中修改CSS变量来自定义颜色主题。

### 添加新页面
1. 在 `src/pages/` 中创建新的页面组件
2. 在 `src/App.tsx` 中添加路由配置
3. 在导航组件中添加链接

## 📦 部署

### Vercel (推荐)
```bash
npm install -g vercel
vercel --prod
```

### Netlify
```bash
npm run build
# 上传 dist 文件夹到 Netlify
```

### GitHub Pages
```bash
npm run build
# 配置 GitHub Actions 自动部署
```

## 🤝 贡献

欢迎提交 Issues 和 Pull Requests！

## 📄 许可证

MIT License

## 👨‍💻 作者

技术探索者 - 专注于前沿技术分享和实践

---

⭐ 如果这个项目对你有帮助，请给个 Star！