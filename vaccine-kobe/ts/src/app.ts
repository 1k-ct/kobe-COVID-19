import axios from "axios";
import { ResArticles } from "./api/articles";
import * as fs from "fs";

export async function FetchItem<T>(url: string): Promise<T> {
  const { data } = await axios.get<T>(url);
  return data;
}

const main = async () => {
  const baseUrl = "https://api-cache.vaccines.sciseed.jp/public/281000/";
  interface URL {
    url: string;
    fileName: string;
  }
  const URLs: URL[] = [
    { url: baseUrl + "articles/", fileName: "./data/articles.json" },
    { url: baseUrl + "department/", fileName: "./data/department.json" },
    {
      url: baseUrl + "available_department/",
      fileName: "./data/available_department.json",
    },
    { url: baseUrl + "item/", fileName: "./data/item.json" },
    {
      url:
        baseUrl +
        "available_date/?department_id=8769&item_id=3&year=2021&month=8",
      fileName: "./data/available_date_1_8.json",
    },
    {
      url:
        baseUrl +
        "available_date/?department_id=8770&item_id=3&year=2021&month=8",
      fileName: "./data/available_date_2_8.json",
    },
    {
      url:
        baseUrl +
        "available_date/?department_id=8769&item_id=3&year=2021&month=9",
      fileName: "./data/available_date_1_9.json",
    },
    {
      url:
        baseUrl +
        "available_date/?department_id=8770&item_id=3&year=2021&month=9",
      fileName: "./data/available_date_2_9.json",
    },
    {
      url:
        baseUrl +
        "reservation_frame/?department_id=8769&item_id=3&start_date_after=2021-07-14&start_date_before=2022-10-10",
      fileName: "./data/reservation_frame_1.json",
    },
    {
      url:
        baseUrl +
        "reservation_frame/?department_id=8770&item_id=3&start_date_after=2021-07-14&start_date_before=2022-10-10",
      fileName: "./data/reservation_frame_2.json",
    },
  ];

  for (const i of URLs) {
    const vaccine = FetchItem<any>(i.url);
    try {
      const v = await vaccine;
      const jsonData = JSON.stringify(v, null, 4);
      fs.writeFileSync(i.fileName, jsonData);
    } catch (err) {
      console.log(err);
    }
  }
};

main();
