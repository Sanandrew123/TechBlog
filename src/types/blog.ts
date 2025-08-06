export interface BlogPost {
  id: number;
  title: string;
  content: string;
  excerpt: string;
  author: string;
  createdAt: string;
  tags: string;
  coverImage?: string;
  readTime: number;
  viewCount: number;
  published: boolean;
  slug: string;
}