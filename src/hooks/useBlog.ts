import { useState, useEffect } from 'react';
import { BlogPost } from '../types/blog';

const API_BASE_URL = '/api/v1';

export const useBlog = () => {
  const [posts, setPosts] = useState<BlogPost[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const loadPosts = async () => {
      try {
        const response = await fetch(`${API_BASE_URL}/posts`);
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const result = await response.json();
        setPosts(result.data || []);
      } catch (error) {
        console.error('Failed to load blog posts:', error);
        setPosts([]);
      } finally {
        setLoading(false);
      }
    };

    loadPosts();
  }, []);

  const getPostBySlug = (slug: string): BlogPost | undefined => {
    if (!Array.isArray(posts)) {
      return undefined;
    }
    return posts.find(post => post.slug === slug);
  };

  const getPostsByTag = (tag: string): BlogPost[] => {
    if (!Array.isArray(posts)) {
      return [];
    }
    return posts.filter(post => {
      const tags = JSON.parse(post.tags || '[]');
      return tags.includes(tag);
    });
  };

  const getFeaturedPosts = (count: number = 3): BlogPost[] => {
    if (!Array.isArray(posts)) {
      return [];
    }
    return posts.slice(0, count);
  };

  return {
    posts,
    loading,
    getPostBySlug,
    getPostsByTag,
    getFeaturedPosts
  };
};