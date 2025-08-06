import React from 'react';
import { Link } from 'react-router-dom';
import Hero from '../../components/Hero/Hero';
import BlogCard from '../../components/BlogCard/BlogCard';
import { useBlog } from '../../hooks/useBlog';
import styles from './Home.module.css';

const Home: React.FC = () => {
  const { posts, getFeaturedPosts, loading } = useBlog();
  const featuredPosts = getFeaturedPosts(3);

  return (
    <div className={styles.home}>
      <Hero />
      
      <section className={styles.featuredSection}>
        <div className="container">
          <div className={styles.sectionHeader}>
            <h2 className={styles.sectionTitle}>精选文章</h2>
            <p className={styles.sectionSubtitle}>
              最新的技术分享和深度思考
            </p>
          </div>
          
          {loading ? (
            <div className={styles.loading}>
              <div className={styles.spinner}></div>
              <p>加载中...</p>
            </div>
          ) : (
            <div className={styles.postsGrid}>
              {featuredPosts.map((post) => (
                <BlogCard key={post.id} post={post} />
              ))}
            </div>
          )}
          
          <div className={styles.viewAll}>
            <Link to="/blog" className="btn btn-primary">
              查看所有文章
            </Link>
          </div>
        </div>
      </section>

      <section className={styles.aboutSection}>
        <div className="container">
          <div className={styles.aboutContent}>
            <div className={styles.aboutText}>
              <h2 className={styles.aboutTitle}>关于这个博客</h2>
              <p className={styles.aboutDescription}>
                这里是一个专注于前沿技术分享的个人博客。我们探讨人工智能的最新发展，
                分享Go、Python、React等技术栈的实践经验，以及对软件工程的深度思考。
              </p>
              <p className={styles.aboutDescription}>
                无论你是初学者还是经验丰富的开发者，都能在这里找到有价值的内容。
                让我们一起在技术的海洋中探索前行。
              </p>
              <a href="/about" className="btn btn-secondary">
                了解更多
              </a>
            </div>
            <div className={styles.aboutVisual}>
              <div className={styles.statsGrid}>
                <div className={styles.statItem}>
                  <div className={styles.statNumber}>{Array.isArray(posts) ? posts.length : 0}</div>
                  <div className={styles.statLabel}>技术文章</div>
                </div>
                <div className={styles.statItem}>
                  <div className={styles.statNumber}>{Array.isArray(posts) ? posts.reduce((sum, post) => sum + post.viewCount, 0) : 0}</div>
                  <div className={styles.statLabel}>阅读量</div>
                </div>
                <div className={styles.statItem}>
                  <div className={styles.statNumber}>3</div>
                  <div className={styles.statLabel}>主要技术栈</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
};

export default Home;