import axios, { AxiosError, AxiosResponse } from "axios";

const URL = "https://api.";

interface Resp {
  articles: Article[] | null;
}
interface Article {
  id: number;
  category: null | string;
  header: string;
  description: string;
  url: null | string;
  body: string;
  information: any;
  created_at: string;
  updated_at: string;
}

async function FetchItem<T>(url: string): Promise<T> {
  try {
    const { data } = await axios.get<T>(url);
    return data;
  } catch (err) {
    return err;
  }
}

const main = async () => {
  const url = URL + "/articles/";
  const vaccine = FetchItem(url);

  console.info(await vaccine);
};

main();
