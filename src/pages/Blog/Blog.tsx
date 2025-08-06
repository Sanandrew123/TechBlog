import React, { useState } from 'react';
import BlogCard from '../../components/BlogCard/BlogCard';
import { useBlog } from '../../hooks/useBlog';
import styles from './Blog.module.css';

const Blog: React.FC = () => {
  const { posts, loading } = useBlog();
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedTag, setSelectedTag] = useState('');

  const allTags = Array.from(
    new Set(posts.flatMap(post => post.tags))
  ).sort();

  const filteredPosts = posts.filter(post => {
    const matchesSearch = post.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
                         post.excerpt.toLowerCase().includes(searchTerm.toLowerCase());
    const matchesTag = selectedTag === '' || post.tags.includes(selectedTag);
    return matchesSearch && matchesTag;
  });

  return (
    <div className={styles.blog}>
      <div className="container">
        <div className={styles.header}>
          <h1 className={styles.title}>技术博客</h1>
          <p className={styles.subtitle}>
            探索前沿技术，分享编程智慧与实践经验
          </p>
        </div>

        <div className={styles.filters}>
          <div className={styles.searchBox}>
            <input
              type="text"
              placeholder="搜索文章..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className={styles.searchInput}
            />
          </div>
          
          <div className={styles.tagFilter}>
            <select
              value={selectedTag}
              onChange={(e) => setSelectedTag(e.target.value)}
              className={styles.tagSelect}
            >
              <option value="">所有标签</option>
              {allTags.map(tag => (
                <option key={tag} value={tag}>{tag}</option>
              ))}
            </select>
          </div>
        </div>

        {loading ? (
          <div className={styles.loading}>
            <div className={styles.spinner}></div>
            <p>加载中...</p>
          </div>
        ) : (
          <>
            <div className={styles.resultsInfo}>
              <p className={styles.resultsCount}>
                找到 {filteredPosts.length} 篇文章
                {selectedTag && (
                  <span className={styles.activeFilter}>
                    · 标签: {selectedTag}
                  </span>
                )}
                {searchTerm && (
                  <span className={styles.activeFilter}>
                    · 搜索: "{searchTerm}"
                  </span>
                )}
              </p>
              
              {(searchTerm || selectedTag) && (
                <button
                  onClick={() => {
                    setSearchTerm('');
                    setSelectedTag('');
                  }}
                  className={styles.clearFilters}
                >
                  清除筛选
                </button>
              )}
            </div>

            {filteredPosts.length === 0 ? (
              <div className={styles.noResults}>
                <h3>没有找到匹配的文章</h3>
                <p>尝试调整搜索关键词或选择不同的标签</p>
              </div>
            ) : (
              <div className={styles.postsGrid}>
                {filteredPosts.map((post) => (
                  <BlogCard key={post.id} post={post} />
                ))}
              </div>
            )}
          </>
        )}
      </div>
    </div>
  );
};

export default Blog;