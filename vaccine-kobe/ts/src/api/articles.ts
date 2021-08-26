export interface ResArticles {
  articles: Article[];
}

interface Article {
  id: number;
  category?: any;
  header: string;
  description: string;
  url?: any;
  body: string;
  information: Information;
  created_at: string;
  updated_at: string;
}

interface Information {
  important: boolean;
}
