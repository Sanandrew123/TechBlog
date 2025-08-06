import React, { useState } from 'react';
import styles from './Contact.module.css';

interface FormData {
  name: string;
  email: string;
  subject: string;
  message: string;
}

const Contact: React.FC = () => {
  const [formData, setFormData] = useState<FormData>({
    name: '',
    email: '',
    subject: '',
    message: ''
  });
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [submitStatus, setSubmitStatus] = useState<'idle' | 'success' | 'error'>('idle');

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);
    
    try {
      await new Promise(resolve => setTimeout(resolve, 1500));
      console.log('Form submitted:', formData);
      setSubmitStatus('success');
      setFormData({ name: '', email: '', subject: '', message: '' });
    } catch (error) {
      setSubmitStatus('error');
    } finally {
      setIsSubmitting(false);
    }
  };

  const contactMethods = [
    {
      icon: '📧',
      title: '邮箱',
      description: 'tech.explorer@example.com',
      action: 'mailto:tech.explorer@example.com'
    },
    {
      icon: '💼',
      title: 'LinkedIn',
      description: '专业网络连接',
      action: '#'
    },
    {
      icon: '🐙',
      title: 'GitHub',
      description: '查看我的项目',
      action: '#'
    },
    {
      icon: '🐦',
      title: 'Twitter',
      description: '关注技术动态',
      action: '#'
    }
  ];

  return (
    <div className={styles.contact}>
      <div className="container">
        <header className={styles.header}>
          <h1 className={styles.title}>联系我</h1>
          <p className={styles.subtitle}>
            有任何问题或想法？我很乐意与你交流讨论！
          </p>
        </header>

        <div className={styles.content}>
          <div className={styles.formSection}>
            <h2 className={styles.sectionTitle}>发送消息</h2>
            
            {submitStatus === 'success' && (
              <div className={styles.successMessage}>
                <p>✅ 消息发送成功！我会尽快回复您。</p>
              </div>
            )}
            
            {submitStatus === 'error' && (
              <div className={styles.errorMessage}>
                <p>❌ 发送失败，请稍后重试或直接发邮件给我。</p>
              </div>
            )}

            <form onSubmit={handleSubmit} className={styles.form}>
              <div className={styles.formRow}>
                <div className={styles.formGroup}>
                  <label htmlFor="name" className={styles.label}>
                    姓名 *
                  </label>
                  <input
                    type="text"
                    id="name"
                    name="name"
                    value={formData.name}
                    onChange={handleChange}
                    required
                    className={styles.input}
                    placeholder="您的姓名"
                  />
                </div>
                <div className={styles.formGroup}>
                  <label htmlFor="email" className={styles.label}>
                    邮箱 *
                  </label>
                  <input
                    type="email"
                    id="email"
                    name="email"
                    value={formData.email}
                    onChange={handleChange}
                    required
                    className={styles.input}
                    placeholder="your@email.com"
                  />
                </div>
              </div>
              
              <div className={styles.formGroup}>
                <label htmlFor="subject" className={styles.label}>
                  主题 *
                </label>
                <select
                  id="subject"
                  name="subject"
                  value={formData.subject}
                  onChange={handleChange}
                  required
                  className={styles.select}
                >
                  <option value="">选择主题</option>
                  <option value="tech-discussion">技术讨论</option>
                  <option value="collaboration">合作机会</option>
                  <option value="blog-feedback">博客反馈</option>
                  <option value="other">其他</option>
                </select>
              </div>
              
              <div className={styles.formGroup}>
                <label htmlFor="message" className={styles.label}>
                  消息 *
                </label>
                <textarea
                  id="message"
                  name="message"
                  value={formData.message}
                  onChange={handleChange}
                  required
                  rows={6}
                  className={styles.textarea}
                  placeholder="请详细描述您想要讨论的内容..."
                />
              </div>
              
              <button
                type="submit"
                disabled={isSubmitting}
                className={`btn btn-primary ${styles.submitButton}`}
              >
                {isSubmitting ? '发送中...' : '发送消息'}
              </button>
            </form>
          </div>

          <div className={styles.infoSection}>
            <h2 className={styles.sectionTitle}>其他联系方式</h2>
            <div className={styles.contactMethods}>
              {contactMethods.map((method) => (
                <a
                  key={method.title}
                  href={method.action}
                  className={styles.contactMethod}
                  target={method.action.startsWith('http') ? '_blank' : undefined}
                  rel={method.action.startsWith('http') ? 'noopener noreferrer' : undefined}
                >
                  <div className={styles.methodIcon}>{method.icon}</div>
                  <div className={styles.methodInfo}>
                    <h3 className={styles.methodTitle}>{method.title}</h3>
                    <p className={styles.methodDescription}>{method.description}</p>
                  </div>
                </a>
              ))}
            </div>

            <div className={styles.responseInfo}>
              <h3 className={styles.responseTitle}>回复时间</h3>
              <p className={styles.responseText}>
                我通常会在24-48小时内回复邮件和消息。
                如果是紧急事务，请在邮件标题中注明"紧急"。
              </p>
              <p className={styles.responseText}>
                最佳联系时间：周一到周五 9:00 - 18:00 (GMT+8)
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Contact;