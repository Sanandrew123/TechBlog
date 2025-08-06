import { useState, useEffect } from 'react';
import { BlogPost } from '../types/blog';
import { blogPosts } from '../data/blogData';

export const useBlog = () => {
  const [posts, setPosts] = useState<BlogPost[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const loadPosts = async () => {
      try {
        await new Promise(resolve => setTimeout(resolve, 500));
        setPosts(blogPosts);
      } catch (error) {
        console.error('Failed to load blog posts:', error);
      } finally {
        setLoading(false);
      }
    };

    loadPosts();
  }, []);

  const getPostById = (id: string): BlogPost | undefined => {
    return posts.find(post => post.id === id);
  };

  const getPostsByTag = (tag: string): BlogPost[] => {
    return posts.filter(post => post.tags.includes(tag));
  };

  const getFeaturedPosts = (count: number = 3): BlogPost[] => {
    return posts.slice(0, count);
  };

  return {
    posts,
    loading,
    getPostById,
    getPostsByTag,
    getFeaturedPosts
  };
};