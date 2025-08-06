import React from 'react';
import styles from './Footer.module.css';

const Footer: React.FC = () => {
  return (
    <footer className={styles.footer}>
      <div className="container">
        <div className={styles.footerContent}>
          <div className={styles.footerSection}>
            <h3 className={styles.footerTitle}>TechBlog</h3>
            <p className={styles.footerDescription}>
              专注于前沿技术分享，人工智能与后端开发
            </p>
            <div className={styles.socialLinks}>
              <a href="#" className={styles.socialLink} aria-label="GitHub">
                <span>GitHub</span>
              </a>
              <a href="#" className={styles.socialLink} aria-label="Twitter">
                <span>Twitter</span>
              </a>
              <a href="#" className={styles.socialLink} aria-label="LinkedIn">
                <span>LinkedIn</span>
              </a>
            </div>
          </div>
          
          <div className={styles.footerSection}>
            <h4 className={styles.sectionTitle}>快速链接</h4>
            <ul className={styles.linkList}>
              <li><a href="/blog" className={styles.footerLink}>最新文章</a></li>
              <li><a href="/about" className={styles.footerLink}>关于我</a></li>
              <li><a href="/contact" className={styles.footerLink}>联系方式</a></li>
            </ul>
          </div>
          
          <div className={styles.footerSection}>
            <h4 className={styles.sectionTitle}>技术栈</h4>
            <ul className={styles.linkList}>
              <li><span className={styles.techItem}>React & TypeScript</span></li>
              <li><span className={styles.techItem}>Go & Python</span></li>
              <li><span className={styles.techItem}>AI & Machine Learning</span></li>
            </ul>
          </div>
        </div>
        
        <div className={styles.footerBottom}>
          <p className={styles.copyright}>
            © {new Date().getFullYear()} TechBlog. 保留所有权利。
          </p>
          <p className={styles.builtWith}>
            用 ❤️ 和 React 构建
          </p>
        </div>
      </div>
    </footer>
  );
};

export default Footer;