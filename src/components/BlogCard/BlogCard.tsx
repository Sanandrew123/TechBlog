import React from 'react';
import { Link } from 'react-router-dom';
import { BlogPost } from '../../types/blog';
import { getTimeAgo } from '../../utils/dateUtils';
import styles from './BlogCard.module.css';

interface BlogCardProps {
  post: BlogPost;
}

const BlogCard: React.FC<BlogCardProps> = ({ post }) => {
  return (
    <article className={styles.card}>
      {post.coverImage && (
        <div className={styles.imageContainer}>
          <img 
            src={post.coverImage} 
            alt={post.title}
            className={styles.image}
          />
        </div>
      )}
      
      <div className={styles.content}>
        <div className={styles.meta}>
          <time className={styles.date} dateTime={post.publishDate}>
            {getTimeAgo(post.publishDate)}
          </time>
          <span className={styles.readTime}>{post.readTime} 分钟阅读</span>
        </div>
        
        <h3 className={styles.title}>
          <Link to={`/blog/${post.id}`} className={styles.titleLink}>
            {post.title}
          </Link>
        </h3>
        
        <p className={styles.excerpt}>
          {post.excerpt}
        </p>
        
        <div className={styles.footer}>
          <div className={styles.tags}>
            {post.tags.map((tag) => (
              <span key={tag} className={styles.tag}>
                {tag}
              </span>
            ))}
          </div>
          
          <Link to={`/blog/${post.id}`} className={styles.readMore}>
            阅读更多 →
          </Link>
        </div>
      </div>
    </article>
  );
};

export default BlogCard;