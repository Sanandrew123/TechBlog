/*
开发心理过程：
1. 设计一个优雅的赞助页面，展示支持作者的价值
2. 提供多种赞助金额选项，让用户灵活选择
3. 集成微信支付，提供便捷的支付体验
4. 添加赞助者名单展示，增强社区感
*/

import React, { useState } from 'react';
import styles from './Sponsor.module.css';

interface SponsorAmount {
  id: string;
  amount: number;
  label: string;
  description: string;
  popular?: boolean;
}

const sponsorAmounts: SponsorAmount[] = [
  {
    id: 'coffee',
    amount: 10,
    label: '请我喝咖啡 ☕',
    description: '支持作者继续创作优质内容'
  },
  {
    id: 'lunch',
    amount: 30,
    label: '请我吃午餐 🍱',
    description: '感谢你对技术分享的认可',
    popular: true
  },
  {
    id: 'book',
    amount: 50,
    label: '买本技术书 📚',
    description: '帮助作者学习新技术'
  },
  {
    id: 'custom',
    amount: 0,
    label: '自定义金额 💝',
    description: '随心赞助，金额不限'
  }
];

const Sponsor: React.FC = () => {
  const [selectedAmount, setSelectedAmount] = useState<string>('lunch');
  const [customAmount, setCustomAmount] = useState<number>(0);
  const [sponsorName, setSponsorName] = useState<string>('');
  const [message, setMessage] = useState<string>('');
  const [isProcessing, setIsProcessing] = useState<boolean>(false);

  const getCurrentAmount = () => {
    if (selectedAmount === 'custom') {
      return customAmount;
    }
    return sponsorAmounts.find(item => item.id === selectedAmount)?.amount || 0;
  };

  const handleSponsor = async () => {
    const amount = getCurrentAmount();
    
    if (amount < 1) {
      alert('请输入有效的赞助金额');
      return;
    }

    setIsProcessing(true);

    try {
      // 调用后端API创建支付订单
      const response = await fetch('/api/v1/sponsor/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          amount,
          sponsorName: sponsorName || '匿名赞助者',
          message,
          paymentMethod: 'wechat'
        })
      });

      if (!response.ok) {
        throw new Error('创建支付订单失败');
      }

      const { qrCode, orderId } = await response.json();
      
      // 显示微信支付二维码
      showPaymentQRCode(qrCode, orderId, amount);
      
    } catch (error) {
      console.error('支付错误:', error);
      alert('支付创建失败，请稍后重试');
    } finally {
      setIsProcessing(false);
    }
  };

  const showPaymentQRCode = (qrCode: string, orderId: string, amount: number) => {
    // 创建支付弹窗
    const modal = document.createElement('div');
    modal.className = styles.paymentModal;
    modal.innerHTML = `
      <div class="${styles.paymentContent}">
        <div class="${styles.paymentHeader}">
          <h3>微信扫码支付</h3>
          <p>支付金额：¥${amount}</p>
        </div>
        <div class="${styles.qrCodeContainer}">
          <img src="${qrCode}" alt="微信支付二维码" />
          <p>请使用微信扫描二维码完成支付</p>
        </div>
        <div class="${styles.paymentActions}">
          <button onclick="this.closest('.${styles.paymentModal}').remove()">
            取消支付
          </button>
          <button onclick="checkPaymentStatus('${orderId}')">
            我已支付
          </button>
        </div>
      </div>
    `;
    
    document.body.appendChild(modal);
  };

  return (
    <div className={styles.sponsor}>
      <div className="container">
        <div className={styles.header}>
          <h1 className={styles.title}>赞助支持</h1>
          <p className={styles.subtitle}>
            您的支持是我创作的最大动力 ❤️
          </p>
        </div>

        <div className={styles.content}>
          <div className={styles.sponsorForm}>
            <div className={styles.section}>
              <h3 className={styles.sectionTitle}>选择赞助金额</h3>
              <div className={styles.amountGrid}>
                {sponsorAmounts.map((item) => (
                  <div
                    key={item.id}
                    className={`${styles.amountOption} ${
                      selectedAmount === item.id ? styles.selected : ''
                    } ${item.popular ? styles.popular : ''}`}
                    onClick={() => setSelectedAmount(item.id)}
                  >
                    {item.popular && (
                      <div className={styles.popularBadge}>推荐</div>
                    )}
                    <div className={styles.amountLabel}>{item.label}</div>
                    {item.amount > 0 && (
                      <div className={styles.amountValue}>¥{item.amount}</div>
                    )}
                    <div className={styles.amountDescription}>
                      {item.description}
                    </div>
                  </div>
                ))}
              </div>

              {selectedAmount === 'custom' && (
                <div className={styles.customAmount}>
                  <label>自定义金额</label>
                  <div className={styles.inputGroup}>
                    <span>¥</span>
                    <input
                      type="number"
                      min="1"
                      placeholder="输入金额"
                      value={customAmount || ''}
                      onChange={(e) => setCustomAmount(Number(e.target.value))}
                    />
                  </div>
                </div>
              )}
            </div>

            <div className={styles.section}>
              <h3 className={styles.sectionTitle}>赞助信息（可选）</h3>
              <div className={styles.inputGroup}>
                <label>您的昵称</label>
                <input
                  type="text"
                  placeholder="匿名赞助者"
                  value={sponsorName}
                  onChange={(e) => setSponsorName(e.target.value)}
                />
              </div>
              <div className={styles.inputGroup}>
                <label>留言</label>
                <textarea
                  placeholder="说点什么吧..."
                  value={message}
                  onChange={(e) => setMessage(e.target.value)}
                  rows={3}
                />
              </div>
            </div>

            <div className={styles.actions}>
              <button
                className={`${styles.sponsorButton} ${isProcessing ? styles.processing : ''}`}
                onClick={handleSponsor}
                disabled={isProcessing}
              >
                {isProcessing ? '创建支付中...' : `微信支付 ¥${getCurrentAmount()}`}
              </button>
            </div>
          </div>

          <div className={styles.sidebar}>
            <div className={styles.supportInfo}>
              <h3>为什么需要您的支持？</h3>
              <ul>
                <li>🚀 持续更新技术文章</li>
                <li>💻 维护开源项目</li>
                <li>📚 学习新技术和工具</li>
                <li>☁️ 服务器和域名费用</li>
                <li>⚡ 提供更好的用户体验</li>
              </ul>
            </div>

            <div className={styles.sponsorStats}>
              <h3>赞助统计</h3>
              <div className={styles.statItem}>
                <span>总赞助次数</span>
                <span>--</span>
              </div>
              <div className={styles.statItem}>
                <span>本月赞助</span>
                <span>--</span>
              </div>
            </div>
          </div>
        </div>

        <div className={styles.sponsorList}>
          <h3 className={styles.sectionTitle}>感谢赞助者</h3>
          <div className={styles.sponsorGrid}>
            <div className={styles.sponsorCard}>
              <div className={styles.sponsorAvatar}>💝</div>
              <div className={styles.sponsorInfo}>
                <div className={styles.sponsorName}>等待第一位赞助者</div>
                <div className={styles.sponsorMessage}>
                  成为第一个支持我的朋友吧！
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

// 全局函数用于检查支付状态
(window as any).checkPaymentStatus = async (orderId: string) => {
  try {
    const response = await fetch(`/api/v1/sponsor/status/${orderId}`);
    const { status } = await response.json();
    
    if (status === 'paid') {
      alert('支付成功！感谢您的支持！');
      // 关闭支付弹窗
      document.querySelector(`.${styles.paymentModal}`)?.remove();
      // 刷新页面显示最新的赞助者信息
      window.location.reload();
    } else {
      alert('支付尚未完成，请继续等待...');
    }
  } catch (error) {
    console.error('检查支付状态失败:', error);
    alert('检查支付状态失败，请稍后重试');
  }
};

export default Sponsor;