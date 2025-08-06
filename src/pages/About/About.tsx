import React from 'react';
import { Link } from 'react-router-dom';
import styles from './About.module.css';

const About: React.FC = () => {
  const skills = [
    { category: '后端开发', items: ['Go', 'Python', 'Node.js', 'PostgreSQL', 'Redis'] },
    { category: '前端开发', items: ['React', 'TypeScript', 'Next.js', 'Vue.js', 'CSS'] },
    { category: '人工智能', items: ['TensorFlow', 'PyTorch', 'Scikit-learn', 'OpenAI API', 'Langchain'] },
    { category: '工具链', items: ['Docker', 'Kubernetes', 'Git', 'Linux', 'AWS'] }
  ];

  const timeline = [
    { year: '2024', event: '开始专注AI和机器学习应用开发' },
    { year: '2023', event: '深入学习Go语言和云原生技术' },
    { year: '2022', event: '转向全栈开发，掌握现代前端框架' },
    { year: '2021', event: '开始软件开发之旅，专注后端技术' }
  ];

  return (
    <div className={styles.about}>
      <div className="container">
        <header className={styles.header}>
          <div className={styles.profileSection}>
            <div className={styles.avatar}>
              <div className={styles.avatarImage}>
                <span className={styles.avatarEmoji}>👨‍💻</span>
              </div>
            </div>
            <div className={styles.intro}>
              <h1 className={styles.title}>关于我</h1>
              <p className={styles.subtitle}>
                热爱技术的全栈开发者，专注于人工智能和现代web开发
              </p>
            </div>
          </div>
        </header>

        <section className={styles.story}>
          <h2 className={styles.sectionTitle}>我的故事</h2>
          <div className={styles.storyContent}>
            <p>
              你好！我是一名充满热情的技术学习者和实践者。从后端开发起步，
              逐渐扩展到全栈开发，现在专注于人工智能和前沿技术的探索。
            </p>
            <p>
              我相信技术的力量能够改变世界，通过这个博客，我希望分享自己在
              Go、Python、React等技术栈上的实践经验，以及对人工智能发展的思考和见解。
            </p>
            <p>
              当我不在编程的时候，我喜欢阅读技术文章、参与开源项目，
              并且持续关注行业的最新发展趋势。我相信持续学习是作为技术人员最重要的品质。
            </p>
          </div>
        </section>

        <section className={styles.skills}>
          <h2 className={styles.sectionTitle}>技能专长</h2>
          <div className={styles.skillsGrid}>
            {skills.map((skillGroup) => (
              <div key={skillGroup.category} className={styles.skillCategory}>
                <h3 className={styles.categoryTitle}>{skillGroup.category}</h3>
                <div className={styles.skillsList}>
                  {skillGroup.items.map((skill) => (
                    <span key={skill} className={styles.skillTag}>
                      {skill}
                    </span>
                  ))}
                </div>
              </div>
            ))}
          </div>
        </section>

        <section className={styles.timeline}>
          <h2 className={styles.sectionTitle}>技术之路</h2>
          <div className={styles.timelineList}>
            {timeline.map((item) => (
              <div key={item.year} className={styles.timelineItem}>
                <div className={styles.timelineYear}>{item.year}</div>
                <div className={styles.timelineEvent}>{item.event}</div>
              </div>
            ))}
          </div>
        </section>

        <section className={styles.contact}>
          <h2 className={styles.sectionTitle}>让我们连接</h2>
          <p className={styles.contactText}>
            如果您对我的文章感兴趣，或者想要讨论技术话题，欢迎与我联系！
          </p>
          <div className={styles.contactActions}>
            <Link to="/contact" className="btn btn-primary">
              联系我
            </Link>
            <Link to="/blog" className="btn btn-secondary">
              阅读博客
            </Link>
          </div>
        </section>
      </div>
    </div>
  );
};

export default About;