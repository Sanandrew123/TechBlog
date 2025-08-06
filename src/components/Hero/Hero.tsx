import React from 'react';
import { Link } from 'react-router-dom';
import styles from './Hero.module.css';

const Hero: React.FC = () => {
  return (
    <section className={styles.hero}>
      <div className="container">
        <div className={styles.heroContent}>
          <div className={styles.textContent}>
            <h1 className={styles.title}>
              探索前沿技术
              <span className={styles.highlight}>分享编程智慧</span>
            </h1>
            <p className={styles.subtitle}>
              专注于人工智能、后端开发与前沿技术分享。
              在这里，我们一起学习Go、Python、React等现代技术栈，
              探讨最新的AI发展趋势和实践经验。
            </p>
            <div className={styles.ctaButtons}>
              <Link to="/blog" className="btn btn-primary">
                阅读博客
              </Link>
              <Link to="/about" className="btn btn-secondary">
                了解更多
              </Link>
            </div>
          </div>
          <div className={styles.visualContent}>
            <div className={styles.codeBlock}>
              <div className={styles.codeHeader}>
                <div className={styles.dots}>
                  <span></span>
                  <span></span>
                  <span></span>
                </div>
                <span className={styles.filename}>main.go</span>
              </div>
              <div className={styles.codeBody}>
                <pre>
                  <code>
{`package main

import (
    "fmt"
    "context"
)

func main() {
    fmt.Println("Hello, Tech World! 🚀")
    
    // 探索无限可能
    ctx := context.Background()
    exploreAI(ctx)
}`}
                  </code>
                </pre>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default Hero;