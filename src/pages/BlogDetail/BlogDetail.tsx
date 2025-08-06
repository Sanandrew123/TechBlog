import React from 'react';
import { useParams, Link, Navigate } from 'react-router-dom';
import { useBlog } from '../../hooks/useBlog';
import { formatDate } from '../../utils/dateUtils';
import styles from './BlogDetail.module.css';

const BlogDetail: React.FC = () => {
  const { slug } = useParams<{ slug: string }>();
  const { getPostBySlug, loading } = useBlog();
  
  if (!slug) {
    return <Navigate to="/blog" replace />;
  }

  const post = getPostBySlug(slug);

  if (loading) {
    return (
      <div className={styles.loading}>
        <div className={styles.spinner}></div>
        <p>加载中...</p>
      </div>
    );
  }

  if (!post) {
    return (
      <div className={styles.notFound}>
        <div className="container">
          <h1>文章未找到</h1>
          <p>抱歉，您查找的文章不存在。</p>
          <Link to="/blog" className="btn btn-primary">
            返回博客列表
          </Link>
        </div>
      </div>
    );
  }

  return (
    <article className={styles.article}>
      <div className="container">
        <header className={styles.header}>
          <nav className={styles.breadcrumb}>
            <Link to="/" className={styles.breadcrumbLink}>首页</Link>
            <span className={styles.breadcrumbSeparator}>›</span>
            <Link to="/blog" className={styles.breadcrumbLink}>博客</Link>
            <span className={styles.breadcrumbSeparator}>›</span>
            <span className={styles.breadcrumbCurrent}>当前文章</span>
          </nav>
          
          <h1 className={styles.title}>{post.title}</h1>
          
          <div className={styles.meta}>
            <div className={styles.metaItem}>
              <span className={styles.metaLabel}>发布日期：</span>
              <time dateTime={post.createdAt}>
                {formatDate(post.createdAt)}
              </time>
            </div>
            <div className={styles.metaItem}>
              <span className={styles.metaLabel}>作者：</span>
              <span>{post.author}</span>
            </div>
            <div className={styles.metaItem}>
              <span className={styles.metaLabel}>阅读时间：</span>
              <span>{post.readTime} 分钟</span>
            </div>
          </div>
          
          <div className={styles.tags}>
            {JSON.parse(post.tags || '[]').map((tag: string) => (
              <span key={tag} className={styles.tag}>
                {tag}
              </span>
            ))}
          </div>
        </header>

        {post.coverImage && (
          <div className={styles.coverImage}>
            <img src={post.coverImage} alt={post.title} />
          </div>
        )}

        <div className={styles.content}>
          <div 
            className={styles.prose}
            dangerouslySetInnerHTML={{ 
              __html: post.content.replace(/\n/g, '<br />') 
            }} 
          />
        </div>

        <footer className={styles.footer}>
          <div className={styles.shareSection}>
            <h3>分享这篇文章</h3>
            <div className={styles.shareButtons}>
              <button className={styles.shareButton}>
                Twitter
              </button>
              <button className={styles.shareButton}>
                LinkedIn
              </button>
              <button className={styles.shareButton}>
                复制链接
              </button>
            </div>
          </div>
          
          <div className={styles.navigation}>
            <Link to="/blog" className="btn btn-secondary">
              ← 返回博客列表
            </Link>
          </div>
        </footer>
      </div>
    </article>
  );
};

export default BlogDetail;